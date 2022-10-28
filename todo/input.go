package todo

type InputTodo struct {
	ActivityGroupId int    `json:"activity_group_id" binding:"required"`
	Title           string `json:"title" binding:"required"`
}
