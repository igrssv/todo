package todo

type TodoList struct {
	ID          int    `json:"id"`
	Titel       string `json:"titel" binding:"required"`
	Descriprion string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Titel       string `json:"Titel"`
	Descriprion string `json:"descriprion"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemID int
}
