package robot

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/liangdas/armyant/task"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/liangdas/armyant/work"
)

var Index int

type Work struct {
	work.MqttWork
	manager  *Manager
	account  string
	QPS      int
	closeSig bool
}

func NewWork(maanger *Manager) *Work {
	this := new(Work)
	this.manager = maanger
	//TODO:Initialize game data
	this.InitialGameData()
	//TODO:改成外部設定ip
	//opts := this.GetDefaultOptions("ws://35.221.162.196:10011")
	opts := this.GetDefaultOptions("ws://127.0.0.1:10011")
	//opts := this.GetDefaultOptions("ws://192.168.2.131:10005")
	//opts := this.GetDefaultOptions("ws://18.179.219.16:10011")
	//opts := this.GetDefaultOptions("ws://wmslot-gs.apolloegame.net/")
	//opts := this.GetDefaultOptions("ws://35.189.179.93:10031")
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
	if err != nil {
		fmt.Println(err.Error())
	}

	//接到服務器端的訊息
	this.On("Lottery/Opening", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})
	this.On("Lottery/Idle", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})
	this.On("Lottery/Betting", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})
	this.On("Lottery/Settlement", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})

	this.On("Lottery/SettlementResult", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})

	this.On("Lottery/Bingo", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})

	this.On("Lottery/Target", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println(msg.Topic(), string(msg.Payload()))
	})
	return this
}

func (this *Work) UnmarshalResult(payload []byte) map[string]interface{} {
	rmasg := map[string]interface{}{}
	json.Unmarshal(payload, &rmasg)
	return rmasg["Result"].(map[string]interface{})
}

func (this *Work) Init(t task.Task) {
	this.Login()

	// this.QPS = 1000 //每一个并发平均每秒请求次数(限流)
	// this.closeSig = false
}

func (this *Work) Close(t task.Task) {
	this.closeSig = true
}

func (this *Work) RunWorker(t task.Task) {
	//this.EnterGame()
}
