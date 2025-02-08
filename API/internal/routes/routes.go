// routes.go contém todas as rotas da API.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vidacalura/AFEB/internal/models"
	"github.com/vidacalura/AFEB/internal/services"
	"github.com/vidacalura/AFEB/internal/utils"
)

func NewRouter() *gin.Engine {
	env := &models.Env{
		DB: utils.ConectarBD(),
	}

	models.E = env

	r := gin.Default()

	r.Use(CORSMiddleFunc())
	r.Use(services.ValidarSessaoUsuario)

	v1 := r.Group("/api")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", services.Login)
	}

	jog := v1.Group("/jogadores")
	{
		jog.GET("/ranking", services.MostrarRankingJogadores)
		jog.GET("/:codJog", services.MostrarJogador)
		jog.POST("", services.CadastrarJogador)
		jog.PUT("", services.EditarCadastroJogador)
		jog.DELETE("/:codJog", services.ExcluirJogador)
	}

	notc := v1.Group("/noticias")
	{
		notc.GET("", services.MostrarTodasNoticias)
		notc.GET("/feed", services.MostrarFeedNoticias)
		notc.GET("/:codNotc", services.MostrarNoticia)
		notc.POST("", services.CriarNoticia)
		notc.PUT("", services.EditarNoticia)
		notc.DELETE("/:codNotc", services.ExcluirNoticia)
	}

	torn := v1.Group("/torneios")
	{
		torn.GET("", services.MostrarTodosTorneios)
		torn.GET("/:codTorn", services.MostrarTorneio)
		torn.POST("", services.CriarTorneio)
		torn.PUT("", services.EditarTorneio)
		torn.DELETE("/:codTorn", services.ExcluirTorneio)
	}

	trof := v1.Group("/trofeus")
	{
		trof.GET("/jogador/:codJog", services.MostrarTrofeusJogador)
		trof.GET("/:codTrof", services.MostrarDadosTrofeu)
		trof.POST("", services.CriarTrofeu)
		trof.PUT("", services.EditarTrofeu)
		trof.DELETE("/:codTrof", services.ExcluirTrofeu)
	}

	usu := v1.Group("/usuarios")
	{
		usu.GET("", services.MostrarTodosUsuarios)
		usu.GET("/:username", services.MostrarUsuario)
		usu.POST("", services.CriarUsuario)
		usu.PUT("", services.EditarUsuario)
		usu.DELETE("/:username", services.ExcluirUsuario)
	}

	return r
}

func CORSMiddleFunc() gin.HandlerFunc {
	return func(c *gin.Context)	{
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://afeb-xadrez.web.app")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
