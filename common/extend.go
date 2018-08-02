package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// BindExtend 扩展熟悉绑定方法防止400异常
func BindExtend(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindWith(obj, binding.JSON)
	return err
}
