package task

type Task struct {
	ID     uint   `gorm: "primarykey"`
	UserId uint   `json: "userId"`
	Title  string `json:"title"`
}
