package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/urlShortener/internal/util"
)

type UrlHandler struct {
	urlStore urlStore
}

type urlStore interface {
	AddUrl(url string, shortenedUrl string) (string, error)
	GetUrl(url string) (string, error)
}

func (h UrlHandler) AddUrl(c *gin.Context) {
	var url = c.Query("value")
	hashedValue := util.GetMd5Hash(url)
	v, err := h.urlStore.AddUrl(url, hashedValue)

	if err != nil {
		println(err)
	}

	c.JSON(http.StatusOK, gin.H{"shortenedUrl": v})
}
func (h UrlHandler) GetUrl(c *gin.Context) {
	var url = c.Query("value")
	v, err := h.urlStore.GetUrl(url)

	if err != nil {
		println(err)
	}

	c.JSON(http.StatusOK, gin.H{"Url": v})
}

func NewHandler(urlStore urlStore) *UrlHandler {
	return &UrlHandler{
		urlStore: urlStore,
	}
}
