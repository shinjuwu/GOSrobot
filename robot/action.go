package robot

import (
	"encoding/json"
	"fmt"
	"time"
)

func (this *Work) Login() {
	// "data" : {
	// 	"token": string , //server token
	// 	"gameCode": string , //遊戲id
	// 	"clientType":string , //遊玩平台(web,android..)
	// 	"platformID":  int //登入平台(1:帳密 2:第三方平台)
	// 	}
	data := map[string]interface{}{
		"token":      "20f9d57bb7311b68968128ae1bf85671",
		"gameCode":   "ragnarok5x20",
		"clientType": "web",
		"platform":   1,
	}
	dataStr, _ := json.Marshal(data)
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}
	bodyByte, _ := json.Marshal(body)
	msg, err := this.Request("Login/HD_Login", bodyByte)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}

func (this *Work) EnterGame() {
	// 	>1. platformID - int 登入平台(1:帳密 2:第三方平台)
	// >2. lobbyID - int  大廳編號
	// >3. gameID - int  遊戲編號
	// >4. userID - int 玩家編號
	// >5. balance_ci - int 帶入金額
	data := map[string]interface{}{
		"gameID": 99,
	}
	dataStr, _ := json.Marshal(data)
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}
	bodyByte, _ := json.Marshal(body)
	msg, err := this.Request("Game/HD_EnterGame", bodyByte)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}

func (this *Work) Spin() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	data := map[string]interface{}{
		"BetMultiple": 1.8,
		"BetLines":    20,
	}
	dataStr, _ := json.Marshal(data)
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}
	byteData, err := json.Marshal(body)
	msg, err := this.Request("Game/HD_Spin", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	this.StartSpin()
}

func (this *Work) StartSpin() {
	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
		data := map[string]interface{}{
			"BetMultiple": 1.8,
			"BetLines":    20,
		}
		dataStr, _ := json.Marshal(data)
		body := map[string]interface{}{
			"sn":       "",
			"isEncode": false,
			"data":     string(dataStr),
		}
		byteData, err := json.Marshal(body)
		msg, err := this.Request("Game/HD_Spin", byteData)

		if err != nil {
			return
		}
		fmt.Println(msg.Topic(), string(msg.Payload()))
	}
}

func (this *Work) SpinDemo() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	data := map[string]interface{}{
		"BetMultiple": 1.8,
		"BetLines":    20,
		"NGDramaNo":   25,
		"BGDramaNo":   -1,
		"IsFreeGame":  false,
	}
	dataStr, _ := json.Marshal(data)
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}

	byteData, err := json.Marshal(body)
	msg, err := this.Request("Game/HD_SpinDemo", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//this.StartSpin()
}
