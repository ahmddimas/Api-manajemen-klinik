package models

type Doctor struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Name      string `json:"name"`
    Specialty string `json:"specialty"`
    Schedule  string `json:"schedule"`
}
