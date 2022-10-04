package openpgp

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"fmt"
	"github.com/ProtonMail/go-crypto/openpgp/packet"
	"io"
	"main/log"
	"os"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
)

var logger = log.Log()

func Encrypt(passphrase []byte, message []byte) ([]byte, error) {
	// Create buffer to write output to
	buf := new(bytes.Buffer)

	// Create encoder
	encoderWriter, err := armor.Encode(buf, "Message", make(map[string]string))
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating OpenPGP armor: %v", err)
	}

	var encryptConfig packet.Config
	encryptConfig.DefaultCipher = packet.CipherAES256
	encryptConfig.DefaultCompressionAlgo = packet.CompressionNone
	encryptConfig.DefaultHash = crypto.SHA256
	encryptConfig.S2KCount = 65535

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

func EncryptFileParts(unencryptedFilePart []uint8) (io.Reader, []byte) {

	//message := "Hello world!"

	pair, err := generateKeyPair("dropzoneUploadKeyPair", "", "support@yugabyte.com")
	if err != nil {
		logger.Error("Error generating key pair")
		os.Exit(1)
	}

	clientSecret := "JoGe9M6DRXcvdhfjK3ggQLvNZKsE3b1kgGP6dAEmJlM"
	serverSecret := "ABSEtmWY2RavL57pcyoXBm4pZ-qHCcPAwg"
	passphrase := serverSecret + clientSecret

	//logger.Infof("%+v\n", pair)

	//logger.Info("message: ", message)
	//logger.Info("Encrypt file: ", unencryptedFilePart)
	logger.Info("Encrypt test: START")
	//pubEntity, err := getEntity([]byte(pair.PublicKey), []byte{})

	if err != nil {
		// handle error
	}

	logger.Info("Created public key entity.")

	logger.Info("Encrypting test message with public key entity.")
	encryptedArmoredFilePart, err := Encrypt([]byte(passphrase), unencryptedFilePart)
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
	//logger.Info("Encrypt test: END\n")

	r := bytes.NewReader(encryptedArmoredFilePart)
	encryptedBlockFilePart, err := armor.Decode(r)
	if err != nil {
		// handle error
	}
	logger.Info("Encrypted file part to block.")

	privEntity, err := getEntity([]byte(pair.PublicKey), []byte(pair.PrivateKey))
	if err != nil {
		// handle error
	}
	logger.Info("Created private key entity.")

	decryptedArmoredFilePart, err := Decrypt(privEntity, encryptedArmoredFilePart)
	if err != nil {
		// handle error
	}
	logger.Info("Decrypted message with private key entity.")

	return encryptedBlockFilePart.Body, decryptedArmoredFilePart
}
