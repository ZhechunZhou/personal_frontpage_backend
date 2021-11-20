package common

import "github.com/gin-gonic/gin"

type Serializer interface {
	response() interface{}
}

type DefaultSerializer struct {
	C *gin.Context
}

type Validator interface {
	bind(c *gin.Context) error
}
