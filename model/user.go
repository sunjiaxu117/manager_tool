package model

import "manger/db"

type Users struct {
	Id       uint `xorm:"pk autoincr"`
	Name     string
	Password string
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) GetByName(name string) error {
	_, err := db.Engine.Where("name = ?", name).Get(u)
	return err
}
