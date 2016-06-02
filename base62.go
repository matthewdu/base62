package base62

import (
	"errors"
)

// characters used for conversion
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// converts n to base62
func Encode(n int64) string {
	chars := [12]byte{} // log base 62 (2^64) is around 10.7

	i := 12
	for n >= 62 {
		q := n / 62
		r := n % 62
		i--
		chars[i] = alphabet[r]
		n = q
	}
	i--
	chars[i] = alphabet[n]

	return string(chars[i:])
}

// converts base62 token to int
func Decode(token string) (int64, error) {
	var n uint64 = 0

	for i := 0; i < len(token); i++ {
		var v byte
		d := token[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 36
		default:
			return 0, errors.New("Invalid token string to decode")
		}
		n *= 62
		n += uint64(v)
	}

	return int64(n), nil
}
