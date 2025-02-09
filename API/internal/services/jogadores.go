// jogadores.go cuida de toda a lógica por trás de jogadores do CCX.
package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CCX-site/internal/models"
)

func MostrarRankingJogadores(c *gin.Context) {
	var ranking models.RankingJogadores

	statusCode, msgErro := ranking.GetTop50CCX()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"ranking": ranking,
		"message": "Ranking encontrado com sucesso!",
	})
}

func MostrarJogador(c *gin.Context) {
	codJogador := c.Param("codJog")

	var jogador models.Jogador
	statusCode, msgErro := jogador.GetJogador(codJogador)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"jogador": jogador,
		"message": "Jogador encontrado com sucesso!",
	})
}

func CadastrarJogador(c *gin.Context) {
	var novoJogador models.Jogador
	if err := c.BindJSON(&novoJogador); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de jogador inválidos."})
		return
	}

	valido, msgErro := novoJogador.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	statusCode, msgErro := novoJogador.RegistrarJogador()
	if statusCode != http.StatusCreated {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Jogador cadastrado com sucesso!",
	})
}

func EditarCadastroJogador(c *gin.Context) {
	var jogador models.Jogador
	if err := c.BindJSON(&jogador); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de jogador inválidos."})
		return
	}

	valido, msgErro := jogador.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	if !models.JogadorExiste(jogador.CodJog) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Jogador a ser editado não encontrado."})
		return
	}

	statusCode, msgErro := jogador.EditarJogador()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Jogador editado com sucesso!",
	})
}

func ExcluirJogador(c *gin.Context) {
	var jogador models.Jogador
	var err error

	jogador.CodJog, err = strconv.Atoi(c.Param("codJog"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Código de jogador inválido."})
		return
	}

	if !models.JogadorExiste(jogador.CodJog) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Jogador a ser excluído não encontrado."})
		return
	}

	statusCode, msgErro := jogador.ExcluirJogador()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Jogador excluído com sucesso!",
	})
}
