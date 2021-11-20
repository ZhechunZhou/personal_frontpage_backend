package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

//Error handling

// Error type that will help return my customized Error info
// {"database": {"hello":"no such table", error: "not_exists"}}
type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// Warp the error info in a object
func NewError(key string, err error) Error {
	res := Error{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// NewValidatorError To handle the error returned by c.bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) Error {
	res := Error{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}
