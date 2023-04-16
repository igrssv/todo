package handler

import (
	"net/http"
	"strconv"
	"todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorRespons(c, http.StatusBadGateway, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id")) // return int for param request ":id"
	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorRespons(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateItem(userId, listId, input)
	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
