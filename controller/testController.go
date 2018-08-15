package controller

import (
	linq "github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
)

func Test1(c *gin.Context) {
	books := []Book{
		Book{
			id:      1,
			title:   "百年孤独",
			authors: []string{"hh", "kk"},
		},
		Book{
			id:      1,
			title:   "悲惨世界",
			authors: []string{"gg", "mm"},
		},
	}
	authors := linq.From(books).SelectMany(func(b interface{}) linq.Query {
		return linq.From(b.(Book).authors)
	}).Results()
	c.JSON(200, gin.H{
		"result": authors,
	})
}

type Book struct {
	id      int
	title   string
	authors []string
}
