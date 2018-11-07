package lobbyservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var LSurl string

func SendIPGet() (*LSRespLayer, error) {
	cmd := "loadBalanceHttpIpGet"
	ipGetData := LSIPGet{
		PlatformID:    2,
		MemberCode:    0,
		AgentID:       0,
		GameID:        2010,
		GameCode:      "ragnarok5x20",
		Token:         "a2d7e10e4057ff73132ee0defc7b15d0",
		Account:       "",
		Password:      "",
		ClientType:    "web",
		ClientVersion: "foo",
		GameTicket:    "",
		ServerIdx:     0,
	}

	v, err := packageResqData(cmd, false, &ipGetData)
	if err != nil {
		panic("json inner error")
	}

	resData := LSRespLayer{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendServerAdd() (*LSRespLayer, error) {
	cmd := "loadBalanceHttpServerAdd"
	serverAddData := LSserverAddData{
		ServerID:      10001,
		ServerVersion: "1.0.0",
		ServerMode:    "SimonMode",
		ReleaseMode:   0,
		PlatformID:    2,
		MemberCode:    1,
		AgentID:       1,
		IP:            "127.0.0.1",
		IP2:           "127.0.0.1",
		DBname:        "gos",
		DBport:        3306,
		ServerPort:    3653,
		HTTPPort:      9888,
	}
	v, err := packageResqData(cmd, false, &serverAddData)
	if err != nil {
		panic("json inner error")
	}
	resData := LSRespLayer{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendMemberLogin(memberLoginData *LSmemberLogin) (*LSRespLayer, error) {
	cmd := "loadBalanceHttpMemberLogin"
	v, err := packageResqData(cmd, false, &memberLoginData)
	if err != nil {
		panic("json inner error")
	}
	resData := LSRespLayer{}
	fmt.Println(v)
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil

}

func SendMemberLogout() (*LSRespLayer, error) {
	cmd := "loadBalanceHttpMemberLogout"
	memberLogoutData := LSMemberLogout{
		ServerID:            10001,
		PlatformID:          2,
		UserID:              23254,
		ThirdPartyUserID:    12264458,
		ThirdPartyUserIDStr: "",
		TotalUser:           12,
		Token:               "c25a27b78fbfa8bfcea1d7622935d990",
		GameTicket:          "d8d8c86e-c03a-4250-9373-4525d11b5e17",
		Account:             "",
		Password:            "",
		GameID:              2010,
		GameCode:            "ragnarok5x20",
		ClientType:          "web",
		ClientVersion:       "1.0.0",
		IsLogin:             true,
	}

	v, err := packageResqData(cmd, false, &memberLogoutData)
	if err != nil {
		panic("json inner error")
	}

	resData := LSRespLayer{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendPlatformInfoGet() (*PlatformInfoGetResponse, error) {
	cmd := "loadBalanceHttpPlatformInfoGet"
	platformInfoGetData := LSPlatformInfoGet{
		ServerID: 1,
	}

	v, err := packageResqData(cmd, false, &platformInfoGetData)
	if err != nil {
		panic("json inner error")
	}

	resData := PlatformInfoGetResponse{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendGameInfoGet() (*GameInfoGetResponse, error) {
	cmd := "loadBalanceHttpGameInfoGet"
	gameInfoGetData := LSGameInfoGet{
		ServerID: 1,
	}

	v, err := packageResqData(cmd, false, &gameInfoGetData)
	if err != nil {
		panic("json inner error")
	}

	resData := GameInfoGetResponse{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendLobbyInfoGet() (*LobbyInfoGetResponse, error) {
	cmd := "loadBalanceHttpLobbyInfoGet"
	lobbyInfoGetData := LSLobbyInfoGet{
		ServerID: 1,
	}
	v, err := packageResqData(cmd, false, &lobbyInfoGetData)
	if err != nil {
		panic("json inner error")
	}
	resData := LobbyInfoGetResponse{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendGameDataGet() (*GameDataGetResponse, error) {
	cmd := "loadBalanceHttpGameDataGet"
	gameDataGetData := LSGameDataGet{
		ServerID:   1,
		PlatformID: 7,
		LobbyID:    1,
		GameID:     12,
		Account:    "shinjuwu",
		UserID:     123,
	}
	v, err := packageResqData(cmd, false, &gameDataGetData)
	if err != nil {
		panic("json inner error")
	}

	resData := GameDataGetResponse{}
	resp, _ := sendData("POST", LSurl, v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func SendGameDataSet(gameDataSetData *GameDataSetData) (*GameDataSetResponse, error) {
	cmd := "loadBalanceHttpGameDataGet"
	v, err := packageResqData(cmd, false, gameDataSetData)
	if err != nil {
		panic("json inner error")
	}

	resData := GameDataSetResponse{}
	resp, err := sendData("POST", LSurl, v)
	fmt.Println(string(resp))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return nil, err
	}
	return &resData, nil
}

func packageResqData(cmd string, isEncode bool, data interface{}) (url.Values, error) {
	commandLayer := LSCommandLayer{
		Sys:      "system",
		Cmd:      cmd,
		Sn:       0,
		IsEncode: isEncode,
		Data:     data,
	}
	bytedata, err := json.Marshal(commandLayer)
	if err != nil {
		return nil, err
	}
	commandLayerStr := string(bytedata)
	v := url.Values{}
	v.Set("platformId", "7")
	v.Set("authKey", "112233")
	v.Set("sendData", commandLayerStr)
	return v, nil
}

func sendData(method string, url string, v url.Values) ([]byte, error) {
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := http.NewRequest(method, url, formBody)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
