// usuarios.go cuida de toda a lógica ligada a login e autenticação de usuários
// no sistema.
package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vidacalura/AFEB/internal/models"
)

func MostrarTodosUsuarios(c *gin.Context) {
	var usuarios models.Usuarios

	statusCode, msgErro := usuarios.GetUsuarios()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"usuarios": usuarios,
		"message":  "Usuários encontrados com sucesso!",
	})
}

func MostrarUsuario(c *gin.Context) {
	username := c.Param("username")

	var usu models.Usuario
	statusCode, msgErro := usu.GetUsuario(username)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"usuario": usu,
		"message": "Usuário encontrado com sucesso!",
	})
}

func CriarUsuario(c *gin.Context) {
	var usu models.Usuario
	if err := c.BindJSON(&usu); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de usuário inválidos."})
		return
	}

	valido, msgErro := usu.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	statusCode, msgErro := usu.CriarUsuario()
	if statusCode != http.StatusCreated {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Usuário cadastrado com sucesso!",
	})
}

func EditarUsuario(c *gin.Context) {
	var usu models.Usuario
	if err := c.BindJSON(&usu); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de usuário inválidos."})
		return
	}

	valido, msgErro := usu.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	if !models.UsuarioExiste(usu.CodUsu) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Usuário a ser alterado não encontrado."})
		return
	}

	statusCode, msgErro := usu.EditarUsuario()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Usuário atualizado com sucesso!",
	})
}

func ExcluirUsuario(c *gin.Context) {
	var usu models.Usuario
	usu.Username = c.Param("username")

	statusCode, msgErro := usu.ExcluirUsuario()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Usuário excluído com sucesso!",
	})
}
