package robot

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/liangdas/armyant/work"
)

type Work struct {
	work.MqttWork
	manager *Manager
}

func NewWork(maanger *Manager) *Work {
	this := new(Work)
	this.manager = maanger
	//TODO:Initialize game data
	//
	//TODO:改成外部設定ip
	opts := this.GetDefaultOptions("ws://127.0.0.1:3653")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("ConnectionLost", err.Error())
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("OnConnectHandler")
	})

	caData, err := ioutil.ReadFile("./robot/caextract.pem")
	if err != nil {
		fmt.Println(err.Error())
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	config := &tls.Config{
		RootCAs:            pool,
		InsecureSkipVerify: true,
	}
	opts.SetTLSConfig(config)
	err = this.Connect(opts)
	if err != nil{
		fmt.Println(err.Error())
	}

	//接到服務器端的訊息
	this.On()
}

func (this *Work) UnmarshalResult(payload []byte) map[string]interface{}{
	rmasg := map[string]interface{}
	json.Unmarshal(payload,&rmasg)
	return rmasg["Result"].(map[string]interface{})
}


