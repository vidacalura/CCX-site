// noticias.go cuida de toda a lógica relacionada às notícias do CCX.
package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CCX-site/internal/models"
)

func MostrarFeedNoticias(c *gin.Context) {
	var feed models.Feed

	statusCode, msgErro := feed.GetFeed()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"noticias": feed,
		"message":  "Notícias encontradas com sucesso!",
	})
}

func MostrarTodasNoticias(c *gin.Context) {
	var feed models.Feed

	statusCode, msgErro := feed.GetNoticias()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"noticias": feed,
		"message":  "Notícias encontradas com sucesso!",
	})
}

func MostrarNoticia(c *gin.Context) {
	codNotc := c.Param("codNotc")

	var noticia models.Noticia
	statusCode, msgErro := noticia.GetNoticia(codNotc)
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"noticia": noticia,
		"message": "Notícia encontrada com sucesso!",
	})
}

func CriarNoticia(c *gin.Context) {
	var noticia models.Noticia
	if err := c.BindJSON(&noticia); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de notícia inválidos."})
		return
	}

	valido, msgErro := noticia.IsValid()
	if !valido {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": msgErro})
		return
	}

	if !models.UsuarioExiste(noticia.CodAutor) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Autor não existe."})
		return
	}

	statusCode, msgErro := noticia.CriarNoticia()
	if statusCode != http.StatusCreated {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Notícia criada com sucesso!",
	})
}

func EditarNoticia(c *gin.Context) {
	var noticia models.Noticia
	if err := c.BindJSON(&noticia); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Dados de notícia inválidos."})
		return
	}

	if len(noticia.Titulo) < 4 || len(noticia.Titulo) > 255 {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Título de notícia inválido."})
		return
	}

	if noticia.Noticia == "" {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Corpo da notícia não pode estar vazio."})
		return
	}

	if !models.UsuarioExiste(noticia.CodAutor) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Autor não existe."})
		return
	}

	if !models.NoticiaExiste(noticia.CodNotc) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Notícia a ser editada não encontrada."})
		return
	}

	statusCode, msgErro := noticia.EditarNoticia()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Notícia editada com sucesso!",
	})
}

func ExcluirNoticia(c *gin.Context) {
	var noticia models.Noticia
	var err error

	noticia.CodNotc, err = strconv.Atoi(c.Param("codNotc"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": "Código de notícia a ser excluída inválido."})
		return
	}

	if !models.NoticiaExiste(noticia.CodNotc) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"error": "Notícia a ser excluída não encontrada."})
		return
	}

	statusCode, msgErro := noticia.ExcluirNoticia()
	if statusCode != http.StatusOK {
		c.IndentedJSON(statusCode, gin.H{"error": msgErro})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"message": "Notícia excluída com sucesso!",
	})
}
