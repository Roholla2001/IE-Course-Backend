package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadInt64(ctx *gin.Context, key string) (int64, bool) {
	s := ctx.Query(key)
	if s == "" {
		ctx.Error(fmt.Errorf("invalid argument: %s", key))
		return 0, false
	}

	var n int64
	var err error
	if n, err = strconv.ParseInt(s, 10, 64); err != nil {
		ctx.Error(fmt.Errorf("invalid argument: %s", key))
		return 0, false
	}

	return n, true
}

func ReadFromJSON(ctx *gin.Context, any interface{}) bool {
	req, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Error(err)
		return false
	}

	err = json.Unmarshal(req, any)
	if err != nil {
		ctx.Error(err)
		return false
	}

	return true
}

func WriteToJSON(ctx *gin.Context, any interface{}, err error) bool {
	if err != nil {
		ctx.Error(err)
		return false
	}

	json, err := json.Marshal(any)
	if err != nil {
		ctx.Error(err)
		return false
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", json)
	return false
}
