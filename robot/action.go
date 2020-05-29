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
		"token":    "660d9af29zcl9hhdimbcz1l2a2", //pk02
		"gameCode": "LTNEW",
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
	this.GetTableList()
	this.EnterTable()
	this.Stake()
	//this.JoinBanker()
	//this.LeftBanker()
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
	//msg, err := this.Request("Infinitystone/HD_EnterGame", bodyByte)
	//msg, err := this.Request("Spslot/HD_EnterGame", bodyByte)
	//msg, err := this.Request("Simpleslot/HD_EnterGame", bodyByte)
	//msg, err := this.Request("Sanheap/HD_EnterGame", bodyByte)
	msg, err := this.Request("Game/HD_EnterGame", bodyByte)
	//msg, err := this.Request("Eliminateslot/HD_EnterGame", bodyByte)

	if err != nil {
		return
	}
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))

	//time.Sleep(4 * time.Second)
	this.Spin()
	//this.Lottery()
	//this.LotteryDemo()
	///this.SpinDemo()
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
	//msg, err := this.Request("Eliminateslot/HD_Spin", byteData)
	//msg, err := this.Request("Sanheap/HD_Spin", byteData)
	//msg, err := this.Request("Infinitystone/HD_Spin", byteData)
	//msg, err := this.Request("Simpleslot/HD_Spin", byteData)
	//msg, err := this.Request("Spslot/HD_Spin", byteData)
	//time.Sleep(10 * time.Second)
	msg, err := this.Request("Game/HD_Spin", byteData)
	if err != nil {
		return
	}
	//res, _ := tool.MsgDecrypt(string(msg.Payload()))
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//time.Sleep(3 * time.Second)
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
	time.Sleep(3 * time.Second)
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
		msg, err := this.Request("Simpleslot/HD_Spin", byteData)
		//msg, err := this.Request("Game/HD_Spin", byteData)

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
	// data := map[string]interface{}{
	// 	"BetLines": 20,
	// 	"DramaNO":  66,
	// 	"OpenW":    true,
	// 	"BetKey":   1,
	// }

	data := map[string]interface{}{
		"BetLines": 0,
		"DramaNO":  1,
		"Split":    3,
		"BetKey":   0,
	}
	dataStr, _ := json.Marshal(data)
	//aesData, _ := tool.MsgEncrypt(string(dataStr))
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Spslot/HD_SpinDemo", byteData)
	//msg, err := this.Request("Simpleslot/HD_SpinDemo", byteData)
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
	//msg, err := this.Request("Ladder/HD_Enter", byteData)
	msg, err := this.Request("Lottery/HD_Enter", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//this.Stake()
}

func (this *Work) GetTableList() {
	data := map[string]interface{}{
		"roomID": 0,
	}
	dataStr, _ := json.Marshal(data)

	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Lottery/HD_GetTableList", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//this.Stake()
}

func (this *Work) EnterTable() {
	data := map[string]interface{}{
		"tableID": 1,
	}
	dataStr, _ := json.Marshal(data)

	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     string(dataStr),
	}

	byteData, _ := json.Marshal(body)
	msg, err := this.Request("Lottery/HD_EnterTable", byteData)

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
	//this.Stake()
}

func (this *Work) Stake() {
	data := map[string]interface{}{
		"target": 1,
	}
	// data := map[string]interface{}{
	// 	"key":      62,
	// 	"paymode":  2,
	// 	"number":   "2202",
	// 	"spbet":    []int{100, 100},
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
	msg, err := this.Request("Lottery/HD_Stake", byteData)
	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}

func (this *Work) JoinBanker() {
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     "",
	}
	byteData, err := json.Marshal(body)
	msg, err := this.Request("Lottery/HD_JoinBanker", byteData)
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

func (this *Work) LeftBanker() {
	body := map[string]interface{}{
		"sn":       "",
		"isEncode": false,
		"data":     "",
	}
	byteData, err := json.Marshal(body)
	msg, err := this.Request("Lottery/HD_LeftBanker", byteData)
	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}
