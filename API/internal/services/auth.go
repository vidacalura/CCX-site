// auth.go lida com toda a lógica de autenticação de login e sessão de usuários
// no sistema.
package services

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vidacalura/AFEB/internal/models"
	"github.com/vidacalura/AFEB/internal/utils"
)

func Login(c *gin.Context) {
	var user models.Usuario
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Dados para login inválidos.",
		})
		return
	}

	user.Senha = utils.CriptografarSenha(user.Senha)

	valido := user.ValidarLogin()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Nome de usuário ou senha inválidos.",
		})
		return
	}

	statusCode, errMsg := user.GetUsuario(user.Username)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": errMsg})
		return
	}

	tokenString := CriarJWT(user.CodUsu, user.Username, user.Adm)
	if tokenString == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar token de login",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"username": user.Username,
		"token":    tokenString,
	})
}

// Middleware de validação de tokens de sessão.
func ValidarSessaoUsuario(c *gin.Context) {
	// Permitir visualização de dados sem conta.
	if c.Request.Method == "GET" || c.Request.URL.Path == "/api/auth/login" {
		c.Next()
		return
	}

	if len(c.Request.Header["Authorization"]) == 0 {
		c.AbortWithStatus(403)
		return
	}

	key, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}

	token := strings.ReplaceAll(c.Request.Header["Authorization"][0], "Bearer ", "")
	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.Next()
}

// Cria um token JWT em string, ou retorna uma string vazia em caso de erro.
func CriarJWT(codUsu []byte, username string, adm bool) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"codUsu":     codUsu,
			"username":   username,
			"adm":        adm,
			"expireTime": time.Now().Add(24 * time.Hour),
		})

	key, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println(err)
		return ""
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}
