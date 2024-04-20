package app

type Task struct {
	UUID        string `json:"id" xorm:"varchar(255) not null pk 'id'"`
	Title       string `json:"title" xorm:"varchar(255) not null 'title'"`
	Description string `json:"description" xorm:"varchar(255) not null 'description'"`
}
