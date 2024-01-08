package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/urlShortener/internal/database"
	"github.com/schoolboybru/urlShortener/internal/util"
)

type handler struct {
	store database.Store
}

type UrlHandler interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
}

func (h *handler) Add(c *gin.Context) {
	var url = c.Query("value")
	hashedValue := util.GetMd5Hash(url)
	v, err := h.store.AddUrl(url, hashedValue)

	if err != nil {
		println(err)
	}

	c.JSON(http.StatusOK, gin.H{"shortenedUrl": v.Short})
}
func (h *handler) Get(c *gin.Context) {
	var url = c.Query("value")
	v, err := h.store.GetUrl(url)

	if err != nil {
		println(err)
	}

	c.Redirect(http.StatusPermanentRedirect, v)
}

func NewHandler(store database.Store) UrlHandler {
	return &handler{
		store: store,
	}
}
