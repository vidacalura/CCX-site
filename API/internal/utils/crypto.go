// crypto.go contém as funcionalidades usadas para
// criptografia de dados do sistema da Connect
package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CriptografarSenha(senha string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	hash := sha512.New()
	hash.Write([]byte(senha))

	return base64.URLEncoding.EncodeToString(hash.Sum([]byte(os.Getenv("key"))))
}
