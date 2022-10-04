package uploader

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func main() {

	packageCode := []byte("og089z0ja3Ti6mTFCHIrrR3EXErmC01e0ukrA0EaWu0")
	clientSecret := []byte("JoGe9M6DRXcvdhfjK3ggQLvNZKsE3b1kgGP6dAEmJlM")

	dk := pbkdf2.Key(clientSecret, packageCode, 1024, 32, sha256.New)
	fmt.Println(hex.EncodeToString(dk))
}
