package line

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Liner interface {
	Config(token string)
	Trigger(sendMsg string) (string, error)
}
type Notify struct {
	Token string
}

func (n *Notify) Config(token string) {
	n.Token = token
}
func (n *Notify) Trigger(sendMsg string) (string, error) {
	url := fmt.Sprintf("https://notify-api.line.me/api/notify?message=%s", url.QueryEscape(sendMsg))
	method := "POST"
	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", n.Token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	log.Printf("Line Notify Send %s", sendMsg)
	//log.Println("recv:", string(body))
	if string(body) != "{\"status\":200,\"message\":\"ok\"}" {
		return string(body), errors.New(string(body))
	}
	return string(body), err
}
func NewNotify() *Notify {
	return &Notify{}
}
