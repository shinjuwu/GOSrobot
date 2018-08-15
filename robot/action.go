package robot

import (
	"GOSrobot/tool"
	"encoding/json"
	"fmt"
	"math/rand"
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
		"token":      "5bfcde0a151b37b9480a0b23c83ae249",
		"gameCode":   "ragnarok5x20",
		"clientType": "web",
		"platform":   1,
	}
	dataStr, _ := json.Marshal(data)
	aesData, _ := tool.MsgEncrypt(string(dataStr))
	fmt.Println(aesData)
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     aesData,
	}
	bodyByte, _ := json.Marshal(body)
	msg, err := this.Request("Login/HD_Login", bodyByte)

	if err != nil {
		return
	}
	res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), res)
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
	aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     aesData,
	}
	bodyByte, _ := json.Marshal(body)
	msg, err := this.Request("Game/HD_EnterGame", bodyByte)

	if err != nil {
		return
	}
	res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), res)
}

func (this *Work) Spin() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	data := map[string]interface{}{
		"BetMultiple": 1.8,
		"BetKey":      5,
		"BetLines":    20,
	}
	dataStr, _ := json.Marshal(data)
	aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     aesData,
	}
	byteData, err := json.Marshal(body)
	msg, err := this.Request("Game/HD_Spin", byteData)

	if err != nil {
		return
	}
	res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), res)
	this.StartSpin()
}

func (this *Work) StartSpin() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
		betKey := rand.Intn(9)
		data := map[string]interface{}{
			"BetMultiple": 1.8,
			"BetLines":    20,
			"BetKey":      betKey,
		}
		dataStr, _ := json.Marshal(data)
		aesData, _ := tool.MsgEncrypt(string(dataStr))
		body := map[string]interface{}{
			"sn":       "",
			"isEncode": false,
			"data":     aesData,
		}
		byteData, err := json.Marshal(body)
		msg, err := this.Request("Game/HD_Spin", byteData)

		if err != nil {
			return
		}
		res, _ := tool.MsgDecrypt(string(msg.Payload()))
		fmt.Println(msg.Topic(), res)
	}
}

func (this *Work) SpinDemo() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	betKey := rand.Intn(9)
	data := map[string]interface{}{
		"BetMultiple": 1.8,
		"BetLines":    20,
		"NGDramaNo":   25,
		"BGDramaNo":   -1,
		"IsFreeGame":  true,
		"BetKey":      betKey,
	}
	dataStr, _ := json.Marshal(data)
	aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     aesData,
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Game/HD_SpinDemo", byteData)

	if err != nil {
		return
	}
	res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), res)
	//this.StartSpin()
}
