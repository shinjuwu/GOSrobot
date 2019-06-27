package robot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func (this *Work) InitialGameData() {
	Index++
	this.account = "test" + strconv.Itoa(Index)
}

func (this *Work) Login() {
	data := map[string]interface{}{
		"token":    "6cc49274dz1v8jtn1lllzp895", //pk02
		"gameCode": "gp",
		"gameID":   2012,
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
	//this.closeSig = true
	// time.Sleep(1 * time.Second)
	//this.EnterGame()
	this.Enter()
	//this.Stake()
	//this.getOpenHistory()
}

func (this *Work) EnterGame() {
	// 	>1. platformID - int 登入平台(1:帳密 2:第三方平台)
	// >2. lobbyID - int  大廳編號
	// >3. gameID - int  遊戲編號
	// >4. userID - int 玩家編號
	// >5. balance_ci - int 帶入金額
	data := map[string]interface{}{
		"gameID": 2013,
	}
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}
	bodyByte, _ := json.Marshal(body)
	msg, err := this.Request("Simpleslot/HD_EnterGame", bodyByte)
	//msg, err := this.Request("Sanheap/HD_EnterGame", bodyByte)
	//msg, err := this.Request("Game/HD_EnterGame", bodyByte)

	if err != nil {
		return
	}
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))

	//time.Sleep(4 * time.Second)
	this.Spin()
	//this.Lottery()
	//this.LotteryDemo()
	//this.SpinDemo()
}

func (this *Work) Spin() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	data := map[string]interface{}{
		"BetMultiple": 1.8,
		"BetKey":      1,
		"BetLines":    20,
		"Choose":      10,
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
	msg, err := this.Request("Simpleslot/HD_Spin", byteData)
	//msg, err := this.Request("Sanheap/HD_Spin", byteData)
	//msg, err := this.Request("Infinitystone/HD_Spin", byteData)
	//msg, err := this.Request("Game/HD_Spin", byteData)
	if err != nil {
		return
	}
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))
	// time.Sleep(4 * time.Second)
	//this.StartSpin()
}

func (this *Work) Lottery() {
	data := map[string]interface{}{
		"BetKey": 1,
	}
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}
	byteData, err := json.Marshal(body)
	msg, err := this.Request("Game/HD_Lottery", byteData)
	if err != nil {
		return
	}
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//time.Sleep(4 * time.Second)
	//this.StartSpin()
}

func (this *Work) StartSpin() {
	ticker := time.NewTicker(3 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())

		data := map[string]interface{}{
			"BetMultiple": 1.8,
			"BetKey":      1,
			"BetLines":    20,
			"Choose":      10,
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
	//betKey := rand.Intn(9)
	// data := map[string]interface{}{
	// 	"BetMultiple": 1.8,
	// 	"BetLines":    20,
	// 	"NGDramaNo":   25,
	// 	"BGDramaNo":   -1,
	// 	"IsFreeGame":  false,
	// 	"BetKey":      betKey,
	// }
	data := map[string]interface{}{
		"BetLines": 10,
		"DramaNO":  1,
		"OpenW":    true,
		"BetKey":   1,
	}
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Simpleslot/HD_SpinDemo", byteData)
	//msg, err := this.Request("Sanheap/HD_SpinDemo", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	//fmt.Println(msg.Topic(), res)
	//this.StartSpin()
}

func (this *Work) LotteryDemo() {
	// 	>1. BetMultiple  int  壓注倍率
	// >2. BetLines: int  壓注線
	data := map[string]interface{}{
		"cheatCode": 2,
	}
	dataStr, _ := json.Marshal(data)
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Game/HD_LotteryDemo", byteData)

	if err != nil {
		return
	}
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//this.StartSpin()
}

func (this *Work) Enter() {

	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     "",
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Gamepoker/HD_Enter", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	this.Stake()
}

func (this *Work) Stake() {

	data := map[string]interface{}{
		"SpadeBets":   10,
		"HeartBets":   0,
		"DiamondBets": 0,
		"ClubBets":    0,
		"JokerBets":   10,
	}
	// data := map[string]interface{}{
	// 	"key":      62,
	// 	"paymode":  2,
	// 	"number":   "2202",
	// 	"spbet":    []int{100, 0},
	// 	"bet_list": [15]int64{},
	// }
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}
	byteData, err := json.Marshal(body)
	msg, err := this.Request("Gamepoker/HD_Stake", byteData)
	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}

func (this *Work) getOpenHistory() {
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     "",
	}
	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Ladder/HD_GetOpenHistory", byteData)
	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}
