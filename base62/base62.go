package base62

import (
	//"fmt"
	//"strings"
)

// All characters
const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length   = string(len(alphabet))
)

// Encode number to base62.
func Encode(n string) string {
	/*if n == nil {
		return nil
	}*/

	s := ""
	/*for ; n > 0; n = n / length {
		s = string(alphabet[n%length]) + s
	}*/
	return s
}

// Decode converts a base62 token to int.
func Decode(key string) (string, error) {
	var n string
	/*for _, c := range []byte(key) {
		i := strings.IndexByte(alphabet, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		n = length*n + string(i)
	}*/
	return n, nil
}
