package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	w := &Wallet{}
	prk, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = prk
	pbk := &w.privateKey.PublicKey
	w.publicKey = pbk

	//SHA256 hashing (result 32 bytes)
	enc1 := sha256.New()
	enc1.Write(w.publicKey.X.Bytes())
	enc1.Write(w.publicKey.Y.Bytes())
	dh1 := enc1.Sum(nil)
	//RIPEMD160 hashing on SHA256 (result 20 bytes)
	enc2 := ripemd160.New()
	enc2.Write(dh1)
	dh2 := enc2.Sum(nil)
	//Add version bytes of main network (version: 0x00)
	vb := make([]byte, 21)
	vb[0] = 0x00
	copy(vb[1:], dh2[:])
	//Add another SHA256 again
	enc4 := sha256.New()
	enc4.Write(vb)
	dh4 := enc4.Sum(nil)
	//Add another one
	enc5 := sha256.New()
	enc5.Write(dh4)
	dh5 := enc5.Sum(nil)
	//Get first 4 bytes for checksum
	chs := dh5[:4]
	//Add checksum bytes at the end of version byte digested hash (result 25 bytes)
	dh6 := make([]byte, 25)
	copy(dh6[:21], vb[:])
	copy(dh6[21:], chs[:])
	//Convert to base58 string
	adr := base58.Encode(dh6)
	w.blockchainAddress = adr

	return w
}

func (w *Wallet) GetPrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) GetPrivateKeyAsString() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) GetPublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) GetPublicKeyAsString() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) GetBlockchainAddress() string {
	return w.blockchainAddress
}
