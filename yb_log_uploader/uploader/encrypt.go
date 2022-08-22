package uploader

import (
	"github.com/jchavannes/go-pgp/pgp"
	"os"
)

func encryptFileParts(unencryptedFilePart []uint8) ([]byte, []byte) {

	//message := "Hello world!"

	pair, err := pgp.GenerateKeyPair("dropzoneUploadKeyPair", "", "support@yugabyte.com")
	if err != nil {
		logger.Error("Error generating key pair")
		os.Exit(1)
	}

	//logger.Infof("%+v\n", pair)

	//logger.Info("message: ", message)
	//logger.Info("Encrypt file: ", unencryptedFilePart)
	logger.Info("Encrypt test: START")
	pubEntity, err := pgp.GetEntity([]byte(pair.PublicKey), []byte{})

	if err != nil {
		// handle error
	}

	logger.Info("Created public key entity.")

	encryptedFilePart, err := pgp.Encrypt(pubEntity, []byte(unencryptedFilePart))
	if err != nil {
		// handle error
	}

	logger.Info("Encrypted test message with public key entity.")
	logger.Info("Encrypted message: ", encryptedFilePart)

	//decryptedMessage := string(decrypted)
	//if decryptedMessage != "Hello World!" {
	//	// handle error
	//}
	//
	//logger.Info(decryptedMessage)
	//logger.Info("Decrypted message equals original message.")
	//logger.Info("Entcrypt test: END\n")

	privEntity, err := pgp.GetEntity([]byte(pair.PublicKey), []byte(pair.PrivateKey))
	if err != nil {
		// handle error
	}
	logger.Info("Created private key entity.")

	decryptedFilePart, err := pgp.Decrypt(privEntity, encryptedFilePart)
	if err != nil {
		// handle error
	}
	logger.Info("Decrypted message with private key entity.")

	return encryptedFilePart, decryptedFilePart
}
