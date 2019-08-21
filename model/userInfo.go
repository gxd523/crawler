package model

import (
	"encoding/json"
	"log"
)

type UserInfo struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Birthplace string
	House      string
	Car        string
	Xinzuo     string
}

func (u *UserInfo) Print(url string, index int) {
	log.Printf("%s\n%d...%+v", url, index, u)
}

func FromJsonObj(obj interface{}) (UserInfo, error) {
	var userInfo UserInfo
	bytes, err := json.Marshal(obj)
	if err != nil {
		return userInfo, err
	}

	err = json.Unmarshal(bytes, &userInfo)
	return userInfo, err
}
