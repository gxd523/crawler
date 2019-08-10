package model

import (
	"log"
)

type UserInfo struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Birthplace string
	House      string
	Car        string
	Xinzuo     string
}

func (u *UserInfo) Print() {
	log.Printf(
		"{姓名:%s,性别:%s,年龄:%d,身高:%dcm,体重:%dkg,收入:%s,婚姻状况:%s,教育状况:%s,工作地%s,籍贯%s,房子:%s,车子:%s,星座:%s}",
		u.Name, u.Gender, u.Age, u.Height, u.Weight, u.Income, u.Marriage, u.Education, u.Occupation, u.Birthplace, u.House, u.Car, u.Xinzuo,
	)
}
