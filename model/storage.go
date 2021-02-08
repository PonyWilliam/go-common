package model
//库房信息表
type Storage struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Name string `json:"name"`
	Level int64 `json:"level"`
	ContactID int64 `json:"Concat_id"`
}
