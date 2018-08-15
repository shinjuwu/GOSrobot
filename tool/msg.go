package tool

import "encoding/json"

type ResData struct {
	Error  string `json:"Error"`
	Result string `json:"Result"`
}

func MsgEncrypt(data string) (string, error) {

	res, err := NeoEnc(data)
	if err != nil {
		return "", err
	}
	return res, nil
}
func MsgDecrypt(data string) (string, error) {
	dataDecode := ResData{}
	err := json.Unmarshal([]byte(data), &dataDecode)
	if dataDecode.Result == "" {
		res := "error code:" + dataDecode.Error
		return res, nil
	}
	res, err := NeoDec(dataDecode.Result)
	if err != nil {
		return "", err
	}
	return res, nil

}
