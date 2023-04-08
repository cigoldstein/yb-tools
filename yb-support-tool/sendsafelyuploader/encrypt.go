package sendsafelyuploader

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"golang.org/x/crypto/pbkdf2"
)

func createClientSecret() string {

	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		panic(err)

	}

	return base64.RawURLEncoding.EncodeToString(token)

}

func (u *Uploader) CreateChecksum() {

	//packageCode := []byte("og089z0ja3Ti6mTFCHIrrR3EXErmC01e0ukrA0EaWu0")
	//clientSecret := []byte("JoGe9M6DRXcvdhfjK3ggQLvNZKsE3b1kgGP6dAEmJlM")

	dk := pbkdf2.Key([]byte(u.ClientSecret), []byte(u.PackageInfo.PackageCode), 1024, 32, sha256.New)
	u.Checksum = hex.EncodeToString(dk)
}

func Encrypt(passphrase []byte, message []byte) ([]byte, error) {

	// configuration for file encryption
	var encryptConfig packet.Config
	encryptConfig.DefaultCipher = packet.CipherAES256
	encryptConfig.DefaultCompressionAlgo = packet.CompressionNone
	encryptConfig.DefaultHash = crypto.SHA256

	// Create buffer to write output to
	buf := new(bytes.Buffer)

	// Create encoder
	encoderWriter, err := armor.Encode(buf, "Message", make(map[string]string))
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating OpenPGP armor: %v", err)
	}

	// Create encryptor with encoder
	encryptorWriter, err := openpgp.SymmetricallyEncrypt(encoderWriter, passphrase, nil, &encryptConfig)
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating entity for encryption: %v", err)
	}

	// Create compressor with encryptor
	// Switching this to "NoCompression since it's required by SendSafely, but leaving it so that it's easy to change later
	compressorWriter, err := gzip.NewWriterLevel(encryptorWriter, gzip.NoCompression)
	if err != nil {
		return []byte{}, fmt.Errorf("Invalid compression level: %v", err)
	}

	// Write message to compressor
	messageReader := bytes.NewReader(message)
	_, err = io.Copy(compressorWriter, messageReader)
	if err != nil {
		return []byte{}, fmt.Errorf("Error writing data to compressor: %v", err)
	}

	compressorWriter.Close()
	encryptorWriter.Close()
	encoderWriter.Close()

	// Return buffer output - an encoded, encrypted, and compressed message
	return buf.Bytes(), nil
}

func EncryptFileParts(serverSecret, clientSecret string, unencryptedFilePart []uint8) (io.Reader, error) {

	passphrase := serverSecret + clientSecret

	log.Print("Encrypting file part with passphrase")
	encryptedArmoredFilePart, err := Encrypt([]byte(passphrase), unencryptedFilePart)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(encryptedArmoredFilePart)
	encryptedBlockFilePart, err := armor.Decode(r)
	if err != nil {
		return nil, err
	}
	log.Print("Encrypted file part to block.")

	return encryptedBlockFilePart.Body, nil
}