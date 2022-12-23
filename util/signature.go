package util

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

//overwrite string form
func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
