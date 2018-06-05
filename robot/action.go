package robot

import "fmt"

func (this *Work) Login() {
	msg, err := this.Request("Login/HD_Login", []byte(""))

	if err != nil {
		return
	}
	fmt.Println(msg.Topic(), string(msg.Payload()))
}
