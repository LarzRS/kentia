package main

import (
	"html/template"
	"kentia/controlador"
	"math/rand"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	html     *template.Template
	servidor *gin.Engine
)

func init() {
	gin.SetMode(gin.DebugMode)
	rand.Seed(time.Now().UnixNano())
	servidor = gin.Default()

	store := sessions.NewCookieStore([]byte("ef7fbfd3d599befe7a86cbf37c8f05c814dcad918b8dbefb441de846c4f62ea3"))
	servidor.Use(sessions.Sessions("mysession", store))

	cargarTemplates()
	servidor.Use(static.Serve("/", static.LocalFile("./public", false)))
	servidor.StaticFile("/login", "./public/index.html")
	servidor.NoRoute(func(c *gin.Context) {
		html.ExecuteTemplate(c.Writer, "404.html", nil)
	})
}

func cargarTemplates() {
	html = template.Must(template.ParseFiles(
		"public/404.html",
		"templates/registroPrenda.html", "templates/combinacion.html"))
	servidor.SetHTMLTemplate(html)
}

func main() {
	servidor.GET("/combinacion", func(c *gin.Context) {
		html.ExecuteTemplate(c.Writer, "combinacion.html", nil)
	})
	servidor.GET("/registroPrenda", controlador.RegistroPrendaGET(html))
	servidor.GET("/", controlador.Index())
	servidor.POST("/login", controlador.Login(html))
	servidor.POST("/registroUsuario", controlador.RegistroUsuario())
	servidor.POST("/registroPrenda", controlador.RegistroPrendaPOST())
	servidor.GET("/generarCombinacion", controlador.GenerarCombinacionGET(html))
	servidor.Run(":3000")
}
