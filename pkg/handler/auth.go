package handler

import (
	"net/http"
	"todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User

	//parse input
	if err := c.BindJSON(&input); err != nil {
		newErrorRespons(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {

}
