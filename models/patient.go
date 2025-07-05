package models

type Patient struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Name   string `json:"name"`
    Gender string `json:"gender"`
    Age    int    `json:"age"`
}
