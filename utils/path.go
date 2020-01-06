package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IndexOfWithString(arr []string, value string) int {
	if len(arr) == 0 {
		return -1
	}
	index := -1
	for i, v := range arr {
		if v == value {
			index = i
			break
		}
	}
	return index
}

func SessionWithList() []string { //列表有什么作用呢？？？
	list := []string{
		"ping",
		"name",
		"admin/user/create",
		"admin/user/login",
		"article/upload/cb", //？？？cb是指什么
		"article/info",
		"article/list",
		"tag/list",
		"tag/info",
		"image/upload/cb",
		"image/upload/ue/cb",
	}

	return list
}

func RelativePath(c *gin.Context, basePath string) string {
	baseLen := len(basePath)

	if basePath[baseLen-1] != '/' {
		basePath = fmt.Sprintf("%s/", basePath)
		baseLen += 1
	}

	absolutePath := c.Request.URL.Path
	if absolutePath[:baseLen] != basePath { //说明绝对路径比basePath短或者不是一个？？
		return absolutePath //为什么这里传绝对路径？？？
	}
	return absolutePath[baseLen:]
}
