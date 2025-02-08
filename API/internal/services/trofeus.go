// trofeus.go lida com todas as premiações dadas a jogadores pela AFEB.
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vidacalura/AFEB/internal/models"
)

func MostrarTrofeusJogador(c *gin.Context) {
	codJog, err := strconv.Atoi(c.Param("codJog"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Código de jogador inválido."})
		return
	}

	if !models.JogadorExiste(codJog) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Jogador não encontrado."})
		return
	}

	var trofeus models.Trofeus
	statusCode, msgErro := trofeus.GetTrofeusJogador(codJog)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"trofeus": trofeus,
		"message": "Troféus encontrados com sucesso!",
	})
}

func MostrarDadosTrofeu(c *gin.Context) {
	codTrof := c.Param("codTrof")

	var trofeu models.Trofeu
	statusCode, msgErro := trofeu.GetDadosTrofeu(codTrof)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"trofeu":  trofeu,
		"message": "Troféu encontrado com sucesso!",
	})
}

func CriarTrofeu(c *gin.Context) {
	var trofeu models.Trofeu
	if err := c.BindJSON(&trofeu); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Request body inválida ou insuficiente."})
		return
	}

	valido, msgErro := trofeu.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	statusCode, msgErro := trofeu.CriarTrofeu()
	if statusCode != http.StatusCreated {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Troféu registrado com sucesso!",
	})
}

func EditarTrofeu(c *gin.Context) {
	var trofeu models.Trofeu
	if err := c.BindJSON(&trofeu); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Request body inválida ou insuficiente."})
		return
	}

	valido, msgErro := trofeu.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	if !models.TrofeuExiste(trofeu.CodTrof) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Troféu a ser editado não encontrado."})
		return
	}

	statusCode, msgErro := trofeu.EditarTrofeu()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Troféu atualizado com sucesso!",
	})
}

func ExcluirTrofeu(c *gin.Context) {
	var trofeu models.Trofeu
	var err error

	trofeu.CodTrof, err = strconv.Atoi(c.Param("codTrof"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Código de troféu inválido."})
		return
	}

	if !models.TrofeuExiste(trofeu.CodTrof) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Troféu a ser excluído não encontrado."})
		return
	}

	statusCode, msgErro := trofeu.ExcluirTrofeu()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Troféu excluído com sucesso!",
	})
}
