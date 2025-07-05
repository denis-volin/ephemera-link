package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	cfg *Config
	storage *Storage
	r *gin.Engine
}

func NewApp(cfg *Config, storage *Storage) *App {
	return &App{cfg: cfg, storage: storage, r: gin.Default()}
}

func (a *App) Run() {
	a.r.Use(gin.Recovery())
	a.r.LoadHTMLGlob("templates/*")
	a.r.Static("/static", "static")
	a.r.GET("/", a.Main)
	a.r.GET("/c/:id/:token", a.OpenSecret)
	a.r.POST("/", a.SaveSecret)
	a.r.POST("/retrieve", a.RetrieveSecret)
	err := a.r.Run(fmt.Sprintf(":%d", a.cfg.ListenPort))
	if err != nil {
		log.Fatalf("Can't start server: %v", err)
	}
}

func (a *App) Main(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func (a *App) SaveSecret(c *gin.Context) {
	secret := c.PostForm("secret")
	id, key, err := a.storage.SaveSecret(secret)
	expire := a.cfg.SecretsExpire
	if err != nil {
		c.Error(err)
		c.HTML(500, "error.tmpl", gin.H{
			"error": "Can't save secret.",
		})
		return
	}
	c.HTML(200, "saved.tmpl", gin.H{
		"link": "c/" + id + "/" + key,
		"expire": expire,
	})
}

func (a *App) OpenSecret(c *gin.Context){
	id := c.Param("id")
	token := c.Param("token")
	c.HTML(200, "view.tmpl", gin.H{
		"id": id,
		"token": token,
	})
}

func (a *App) RetrieveSecret(c *gin.Context) {
	id := c.PostForm("id")
	token := c.PostForm("token")
	data, err := a.storage.GetSecret(id, token)
	if err != nil {
		c.Error(err)
		c.HTML(500, "error.tmpl", gin.H{
			"error": "This secret has already been viewed or the link is invalid.",
		})
		return
	}
	c.HTML(200, "retrieve.tmpl", gin.H{
		"secret": data,
	})
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.r.ServeHTTP(w, r)
}
