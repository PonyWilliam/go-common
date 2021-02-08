package model
//存储表
type DeviceArea struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment;" json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}