package persist

import "crawler/model"

func ItemSaver() chan model.UserInfo {
	itemChan := make(chan model.UserInfo)
	go func() {
		userInfoCount := 0
		for userInfo := range itemChan {
			userInfoCount++
			userInfo.Print(userInfoCount)
		}
	}()
	return itemChan
}
