package openpgp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"main/log"
	"os"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
)

var logger = log.Log()

func Encrypt(entity *openpgp.Entity, message []byte) ([]byte, error) {
	// Create buffer to write output to
	buf := new(bytes.Buffer)

	// Create encoder
	encoderWriter, err := armor.Encode(buf, "Message", make(map[string]string))
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating OpenPGP armor: %v", err)
	}

	// Create encryptor with encoder
	encryptorWriter, err := openpgp.Encrypt(encoderWriter, []*openpgp.Entity{entity}, nil, nil, nil)
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

func EncryptFileParts(unencryptedFilePart []uint8) ([]byte, []byte) {

	//message := "Hello world!"

	pair, err := generateKeyPair("dropzoneUploadKeyPair", "", "support@yugabyte.com")
	if err != nil {
		logger.Error("Error generating key pair")
		os.Exit(1)
	}

	//logger.Infof("%+v\n", pair)

	//logger.Info("message: ", message)
	//logger.Info("Encrypt file: ", unencryptedFilePart)
	logger.Info("Encrypt test: START")
	pubEntity, err := getEntity([]byte(pair.PublicKey), []byte{})

	if err != nil {
		// handle error
	}

	logger.Info("Created public key entity.")

	logger.Info("Encrypting test message with public key entity.")
	encryptedFilePart, err := Encrypt(pubEntity, []byte(unencryptedFilePart))
	if err != nil {
		// handle error
	}

	//logger.Info("Encrypted message: ", encryptedFilePart)

	//decryptedMessage := string(decrypted)
	//if decryptedMessage != "Hello World!" {
	//	// handle error
	//}
	//
	//logger.Info(decryptedMessage)
	//logger.Info("Decrypted message equals original message.")
	//logger.Info("Entcrypt test: END\n")

	privEntity, err := getEntity([]byte(pair.PublicKey), []byte(pair.PrivateKey))
	if err != nil {
		// handle error
	}
	logger.Info("Created private key entity.")

	decryptedFilePart, err := Decrypt(privEntity, encryptedFilePart)
	if err != nil {
		// handle error
	}
	logger.Info("Decrypted message with private key entity.")

	return encryptedFilePart, decryptedFilePart
}
