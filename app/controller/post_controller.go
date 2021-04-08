package controller

import (
	"goAPISample/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

var ps service.PostService

func (pc PostController) Index(c *gin.Context) {
	posts, err := ps.GetPosts()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

func (pc PostController) Show(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	post, err := ps.GetPostById(idInt)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, post)
	}
}

func (pc PostController) Create(c *gin.Context) {
	// 検索処理
	type Body struct {
		Content string `json:"content"`
	}

	var body Body
	c.BindJSON(&body)
	post, err := ps.CreatePost(body.Content)
	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusCreated, post)
	}
}

func (pc PostController) Update(c *gin.Context) {
	// 検索処理
	type Body struct {
		Content string `json:"content"`
	}

	var body Body

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	c.BindJSON(&body)
	post, err := ps.UpdatePostById(idInt, body.Content)
	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusCreated, post)
	}
}

func (pc PostController) Delete(c *gin.Context) {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	post, err := ps.DeletePostById(idInt)
	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, post)
	}
}
