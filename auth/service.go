package auth

import (
	"errors"
	"strings"
	
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
)

func VerifySignature(address, nonce, signature string) (bool, error) {
	if !strings.HasPrefix(signature, "0x") {
		return false, errors.New("invalid signature format")
	}
	sig = common.FromHex(signature)
	if len(sig) != 65 {
		return false, errors.New("invalid signature length")
	}
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	msg := accounts.TextHash([]byte(nonce))
	pubKey, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false, err
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return strings.EqualFold(recoveredAddr.Hex(), address), nil


}
