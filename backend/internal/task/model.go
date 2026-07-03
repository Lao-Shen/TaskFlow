package task

type Task struct {
	ID   uint64 `gorm:"primary_key;auto_increment"`
	Name string `gorm:"type:varchar(255);not null"`
}
