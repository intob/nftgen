package util

import (
	"crypto/rand"
	"math/big"
)

func RandInt(max int) (int, error) {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(i.Int64()), nil
}
