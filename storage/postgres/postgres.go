package postgres

import (
	"database/sql"
	"fmt"

	// This loads the postgres drivers.
	_ "github.com/lib/pq"

	"github.com/rafaelgarbinatto/ursho/base62"
	"github.com/rafaelgarbinatto/ursho/storage"
)

// New returns a postgres backed storage service.
func New(host, port, user, password, dbName string) (storage.Service, error) {
	// Connect postgres
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	// Ping to connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	strQuery := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";CREATE EXTENSION IF NOT EXISTS \"pgcrypto\";CREATE TABLE shortener( uuid text default crypt(uuid_generate_v4()::text, 'domain.com.br') not null, url varchar not null, visited boolean default false, count integer default 0 );"

	_, err = db.Exec(strQuery)
	if err != nil {
		return nil, err
	}
	return &postgres{db}, nil
}

type postgres struct{ db *sql.DB }

func (p *postgres) Save(url string) (string, error) {
	var id string
	err := p.db.QueryRow("INSERT INTO shortener(url,visited,count) VALUES($1,$2,$3) returning uuid;", url, false, 0).Scan(&id)
	if err != nil {
		return "", err
	}
	return base62.Encode(id), nil
}

func (p *postgres) Load(code string) (string, error) {
	id, err := base62.Decode(code)
	if err != nil {
		return "", err
	}

	var url string
	err = p.db.QueryRow("update shortener set visited=true, count = count + 1 where uuid=$1 RETURNING url", id).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (p *postgres) LoadInfo(code string) (*storage.Item, error) {
	id, err := base62.Decode(code)
	if err != nil {
		return nil, err
	}

	var item storage.Item
	err = p.db.QueryRow("SELECT url, visited, count FROM shortener where uuid=$1 limit 1", id).
		Scan(&item.URL, &item.Visited, &item.Count)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (p *postgres) Close() error { return p.db.Close() }
