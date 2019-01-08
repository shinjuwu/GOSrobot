package lobbyservice

import (
	"fmt"
	"testing"
)

func TestSendIPGet(t *testing.T) {
	LSurl = "http://127.0.0.1:8888/LoadBalance"
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendIPGet()
			if err != nil {
				t.Errorf("SendIPGet() inner error")
			}
			if got.Code != 0 {
				t.Errorf("SendIPGet() error %s", got.Message)
			}
			fmt.Println(got.Data)
		})
	}
}

func TestSendServerAdd(t *testing.T) {
	LSurl = "http://192.168.2.131:30001/LoadBalance"
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendServerAdd()
			if err != nil {
				t.Errorf("SendServerAdd() inner error")
			}
			if got.Code != 0 {
				t.Errorf("SendServerAdd() error : %s", got.Message)
			}

		})
	}
}

func TestSendMemberLogin(t *testing.T) {
	LSurl = "http://192.168.2.131:30001/LoadBalance"
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendMemberLogin()
			if err != nil {
				t.Errorf("SendMemberLogin() inner failed")
			}
			if got.Code != 0 {
				t.Errorf("SendMemberLogin() error: %s ", got.Message)
			}
			aa := got.Data.(map[string]interface{})
			thrID := aa["thirdPartyUserId"].(float64)
			fmt.Println(thrID)
			fmt.Println(int64(thrID))
		})
	}
}

func TestSendMemberLogout(t *testing.T) {
	LSurl = "http://192.168.2.131:30001/LoadBalance"
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendMemberLogout()
			if err != nil {
				t.Errorf("SendMemberLogout() inner error")
			}
			if got.Code != 0 {
				t.Errorf(" SendMemberLogout() error %s", got.Message)
			}
		})
	}
}
