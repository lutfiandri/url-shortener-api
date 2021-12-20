package entity

type Url struct {
	Id      string `gorm:"unique;primaryKey;not null"`
	LongUrl string `gorm:"not null"`
	UserId  string
}
