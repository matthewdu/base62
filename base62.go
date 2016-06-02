package base62

import (
	"errors"
)

const (
	alphabet  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maxUint64 = (1<<64 - 1)
)

// converts n to base62
func Encode(n uint64) string {
	chars := [12]byte{} // log base 62 (2^64-1) is around 10.7

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
func Decode(token string) (uint64, error) {
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
			return 0, errors.New("invalid token")
		}

		n1 := n
		n *= 62
		if n < n1 {
			return maxUint64, errors.New("value out of range")
		}

		n1 = n + uint64(v)
		if n1 < n {
			return maxUint64, errors.New("value out of range")
		}
		n = n1
	}

	return n, nil
}
