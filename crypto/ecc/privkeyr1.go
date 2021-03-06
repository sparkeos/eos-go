package ecc

import (
	"fmt"
)

type innerR1PrivateKey struct {
}

func (k *innerR1PrivateKey) publicKey() PublicKey {
	//TODO: Finish R1 support here - for now we  return bogus key
	//var pubKeyData []byte
	//pubKeyData = append(pubKeyData, byte(1))
	//pubKeyData = append(pubKeyData, bytes.Repeat([]byte{0}, 33)...)
	var content [33]byte
	content[0] = 1
	return PublicKey{Curve: CurveR1, Content: content, inner: &innerR1PublicKey{}}
}

func (k *innerR1PrivateKey) sign(hash []byte) (out Signature, err error) {
	return out, fmt.Errorf("R1 not supported")
}

func (k *innerR1PrivateKey) string() string {
	return "PVT_R1_PLACE_HOLDER"
}
func (k *innerR1PrivateKey) Serialize() []byte {
	return nil
}
