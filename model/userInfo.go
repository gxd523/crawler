package model

import (
	"encoding/json"
	"log"
)

type UserInfo struct {
	Name        string
	Gender      string
	Age         int
	Height      int
	Weight      int
	Birthplace  string
	Xinzuo      string
	City        string
	Nationality string
	Education   string
	Marriage    string
	Occupation  string
	Income      string
	House       string
	Car         string
	HaveChild   string
	WannaChild  string
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
