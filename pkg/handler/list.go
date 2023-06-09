package handler

import (
	"net/http"
	"strconv"
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

// struct by response get all lists
type getAllResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorRespons(c, http.StatusBadGateway, err.Error())
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
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

	list, err := h.services.TodoList.GetById(userId, listId)

	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

func (h *Handler) updateList(c *gin.Context) {
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
	var input todo.UpdateListInput

	if err := c.BindJSON(&input); err != nil {
		newErrorRespons(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, listId, input); err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
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

	err = h.services.TodoList.DeleteList(userId, listId)

	if err != nil {
		newErrorRespons(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
