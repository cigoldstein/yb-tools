package uploader

import (
	"github.com/jchavannes/go-pgp/pgp"
)

func generateKeyPair() {

	message := "Hello world!"

	pair, err := pgp.GenerateKeyPair("dropzoneUploadKeyPair", "", "support@yugabyte.com")
	if err != nil {
		return
	}

	logger.Infof("%+v\n", pair)

	logger.Info("message: ", message)
	logger.Info("Encrypt test: START")
	pubEntity, err := pgp.GetEntity([]byte(pair.PublicKey), []byte{})

	if err != nil {
		// handle error
	}

	logger.Info("Created public key entity.")

	encrypted, err := pgp.Encrypt(pubEntity, []byte(message))
	if err != nil {
		// handle error
	}

	logger.Info("Encrypted test message with public key entity.")
	logger.Info("Encrypted message: ", encrypted)

	privEntity, err := pgp.GetEntity([]byte(pair.PublicKey), []byte(pair.PrivateKey))
	if err != nil {
		// handle error
	}
	logger.Info("Created private key entity.")

	decrypted, err := pgp.Decrypt(privEntity, encrypted)
	if err != nil {
		// handle error
	}
	logger.Info("Decrypted message with private key entity.")

	decryptedMessage := string(decrypted)
	if decryptedMessage != "Hello World!" {
		// handle error
	}

	logger.Info(decryptedMessage)
	logger.Info("Decrypted message equals original message.")
	logger.Info("Entcrypt test: END\n")

}
