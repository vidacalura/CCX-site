// torneios.go cuida de toda a lógica relacionada aos torneios do CCX.
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CCX-site/internal/models"
)

func MostrarTodosTorneios(c *gin.Context) {
	var torneios models.Torneios

	statusCode, msgErro := torneios.GetTorneios()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"torneios": torneios,
		"message":  "Torneios encontrados com sucesso!",
	})
}

func MostrarTorneio(c *gin.Context) {
	codTorneio := c.Param("codTorn")

	var torneio models.Torneio
	statusCode, msgErro := torneio.GetTorneio(codTorneio)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"torneio": torneio,
		"message": "Torneio encontrado com sucesso!",
	})
}

func CriarTorneio(c *gin.Context) {
	var torneio models.Torneio
	if err := c.BindJSON(&torneio); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de torneio inválidos."})
		return
	}

	valido, msgErro := torneio.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	statusCode, msgErro := torneio.RegistrarTorneio()
	if statusCode != http.StatusCreated {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Torneio criado com sucesso!",
	})
}

func EditarTorneio(c *gin.Context) {
	var torneio models.Torneio
	if err := c.BindJSON(&torneio); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de torneio inválidos."})
		return
	}

	valido, msgErro := torneio.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	if !models.TorneioExiste(torneio.CodTorn) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Torneio a ser editado não encontrado."})
	}

	statusCode, msgErro := torneio.EditarTorneio()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Torneio editado com sucesso!",
	})
}

func ExcluirTorneio(c *gin.Context) {
	var torneio models.Torneio
	var err error

	torneio.CodTorn, err = strconv.Atoi(c.Param("codTorn"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Código de torneio a ser excluído inválido."})
		return
	}

	if !models.TorneioExiste(torneio.CodTorn) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Torneio a ser excluído não encontrado."})
	}

	statusCode, msgErro := torneio.ExcluirTorneio()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Torneio excluído com sucesso!",
	})
}
