package robot

import (
	"GOSrobot/lobbyservice"
	"GOSrobot/tool"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func (this *Work) InitialGameData() {
	Index++
	this.account = "test" + strconv.Itoa(Index)
}

func (this *Work) Login() {
	lobbyservice.LSurl = "http://192.168.2.131:30001/LoadBalance"
	ipGetResp, _ := lobbyservice.SendIPGet("")
	ipGetInfo := ipGetResp.Data.(map[string]interface{})
	serverID := int(ipGetInfo["serverId"].(float64))
	gameTicket := ipGetInfo["gameTicket"].(string)
	ip := ipGetInfo["ip"].(string)
	fmt.Println(ip)
	data := map[string]interface{}{
		"token":      "0d15c68a91651c5c0ea5cbfe5a63e570",
		"gameCode":   "ragnarok5x20",
		"clientType": "web",
		"platformID": 2,
		"gameTicket": gameTicket,
		"serverID":   serverID,
		"account":    "",
		"password":   "",
	}
	dataStr, _ := json.Marshal(data)
	// aesData, _ := tool.MsgEncrypt(string(dataStr))
	// fmt.Println(aesData)
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
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//this.closeSig = true
	// time.Sleep(1 * time.Second)
	// this.EnterGame()
}

func (this *Work) EnterGame() {
	// 	>1. platformID - int 登入平台(1:帳密 2:第三方平台)
	// >2. lobbyID - int  大廳編號
	// >3. gameID - int  遊戲編號
	// >4. userID - int 玩家編號
	// >5. balance_ci - int 帶入金額
	data := map[string]interface{}{
		"gameID": 2010,
	}
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
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
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))

	time.Sleep(4 * time.Second)
	this.Spin()
}

func (this *Work) Spin() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	data := map[string]interface{}{
		"BetMultiple": 1.8,
		"BetKey":      1,
		"BetLines":    20,
		"Choose":      1,
		"MonsterID":   1,
	}
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
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
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))
	time.Sleep(4 * time.Second)
	this.StartSpin()
}

func (this *Work) StartSpin() {
	ticker := time.NewTicker(3 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())

		data := map[string]interface{}{
			"BetMultiple": 1.8,
			"BetKey":      1,
			"BetLines":    20,
			"Choose":      1,
		}
		dataStr, _ := json.Marshal(data)
		//aesData, _ := tool.MsgEncrypt(string(dataStr))
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
		//res, _ := tool.MsgDecrypt(string(msg.Payload()))
		fmt.Println(msg.Topic(), string(msg.Payload()))
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
		"IsFreeGame":  false,
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
