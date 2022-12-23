package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
)

func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := string128ToBigIn64Tuple(s)
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
}

func string128ToBigIn64Tuple(s string) (big.Int, big.Int) {
	bx, _ := hex.DecodeString(s[:64])
	by, _ := hex.DecodeString(s[64:])

	var bix, biy big.Int
	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)

	return bix, biy
}
