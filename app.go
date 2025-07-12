package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg     *Config
	storage *Storage
	r       *gin.Engine
}

type Retrieve struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Link string `json:"link"`
}

func NewApp(cfg *Config, storage *Storage) *App {
	return &App{cfg: cfg, storage: storage, r: gin.Default()}
}

func (a *App) Run() {
	a.r.Use(gin.Recovery())
	a.r.LoadHTMLGlob("templates/*")
	a.r.Static("/static", "static")
	a.r.Static("/favicon.ico", "static/favicon.ico")
	a.r.GET("/", a.Main)
	a.r.GET("/c/:id/:token", a.OpenSecret)
	a.r.POST("/saved", a.SaveSecret)
	a.r.POST("/retrieve", a.RetrieveSecret)
	a.r.POST("/api/create", a.SaveSecretAPI)
	a.r.POST("/api/retrieve", a.RetrieveSecretAPI)
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
		"link":   a.cfg.URI + "c/" + id + "/" + key,
		"expire": expire,
	})
}

func (a *App) OpenSecret(c *gin.Context) {
	id := c.Param("id")
	token := c.Param("token")
	c.HTML(200, "view.tmpl", gin.H{
		"id":    id,
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

func (a *App) SaveSecretAPI(c *gin.Context) {
	secret, err := c.GetRawData()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	id, key, err := a.storage.SaveSecret(string(secret))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create secret",
		})
		return
	}
	expire := a.cfg.SecretsExpire
	expiryDate := time.Now().Add(time.Second * time.Duration(expire))

	c.JSON(http.StatusOK, gin.H{
		"id":                 id,
		"key":                key,
		"expires_in_seconds": fmt.Sprintf("%d", expire),
		"expires_at":         fmt.Sprint(expiryDate.Format(time.RFC3339)),
		"link":               a.cfg.URI + "c/" + id + "/" + key,
	})
}

func (a *App) RetrieveSecretAPI(c *gin.Context) {
	var json Retrieve
	if err := c.ShouldBindJSON(&json); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	if len(json.ID) != 0 && len(json.Key) != 0 {
		data, err := a.storage.GetSecret(json.ID, json.Key)
		if err != nil {
			c.Error(err)
			c.JSON(500, gin.H{
				"error": "This secret has already been viewed or the id/key is invalid.",
			})
			return
		}

		c.String(http.StatusOK, data+"\n")
	} else if len(json.Link) != 0 {
		link, err := url.Parse(json.Link)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error parsing link",
			})
			return
		} else if link.Scheme+"://"+link.Host+"/" != a.cfg.URI {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error parsing link",
			})
			return
		}

		segments := strings.Split(strings.Trim(link.Path, "/"), "/")
		if len(segments) == 3 {
			id := segments[1]
			key := segments[2]
			data, err := a.storage.GetSecret(id, key)
			if err != nil {
				c.Error(err)
				c.JSON(500, gin.H{
					"error": "This secret has already been viewed or the id/key is invalid.",
				})
				return
			}

			c.String(http.StatusOK, data+"\n")
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide id and key or link",
		})
		return
	}
}
