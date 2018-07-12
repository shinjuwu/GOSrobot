package robot

import (
	"encoding/json"
	"fmt"
)

func (this *Work) Login() {
	// "data" : {
	// 	"token": string , //server token
	// 	"gameCode": string , //遊戲id
	// 	"clientType":string , //遊玩平台(web,android..)
	// 	"platformID":  int //登入平台(1:帳密 2:第三方平台)
	// 	}
	data := map[string]interface{}{
		"token":      "67e59e24e131abcd23f94a8aa33dabd5",
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
		"platformID": 1,
		"lobbyID":    1,
		"gameID":     1,
		"userID":     23445,
		"balance_ci": 0,
	}

	byteData, _ := json.Marshal(data)
	msg, err := this.Request("Game/HD_EnterGame", byteData)

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
		"BetLines":    12,
	}
	byteData, err := json.Marshal(data)
	msg, err := this.Request("Game/HD_Spin", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}
