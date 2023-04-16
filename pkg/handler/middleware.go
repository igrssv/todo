package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorization = "Authorization"
	userCtx       = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	//get header != empty
	header := c.GetHeader(authorization)
	if header == "" {
		newErrorRespons(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	// split get header
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorRespons(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//parse token
	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorRespons(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

// convert id in idInt
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorRespons(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorRespons(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil

}
