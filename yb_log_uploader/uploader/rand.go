package uploader

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func CreateClientSecret() {

	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 32)
	rand.Read(token)
	fmt.Println(token)
	data := base64.URLEncoding.EncodeToString(token)
	fmt.Println(data)
}
