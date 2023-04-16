package handler

import (
	"net/http"
	"todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	//search user for id
	userId, err := getUserId(c)
	if err != nil {
		newErrorRespons(c, http.StatusBadGateway, err.Error())
		return
	}

	// create todo list
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorRespons(c, http.StatusBadGateway, err.Error())
		return
	}

	// call service
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
	}

	// return id create list
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
