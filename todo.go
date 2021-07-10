package myapp

type TodoLost struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
type TodoItem struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}
type ListItem struct {
	Id     int
	ItemId int
	ListId int
}
