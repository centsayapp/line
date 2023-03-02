package line

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LineHandler struct {
	Environment   string
	ChannelToken  string
	ChannelSecret string
}

func (h *LineHandler) Debugf(format string, a ...interface{}) {
	if h.Environment == "dev" {
		log.Printf(format, a...)
	}
}

func (l *LineHandler) SignLineRequest(body string) string {
	mac := hmac.New(sha256.New, []byte(l.ChannelSecret))
	mac.Write([]byte(body))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return signature
}

func (l *LineHandler) GetUserProfile(uid string) (*LineUserProfile, error) {
	req, err := http.NewRequest("GET", getUserProfileAPI+"/"+uid, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+l.ChannelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling Line get user profile API: %s", err)
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("cannot get user profile: %s", string(respBody))
	}

	profile := new(LineUserProfile)
	err = json.Unmarshal(respBody, profile)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal user profile: %s", err)
	}

	return profile, nil
}

func (l *LineHandler) SendReplyMessage(token string, messages LineMessage) error {
	replyMsg := &LineReplyMessage{
		ReplyToken: token,
		Messages:   []LineMessage{messages},
	}

	s, err := json.Marshal(replyMsg)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(s)

	req, err := http.NewRequest("POST", replyMessageAPI, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.ChannelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot send Line message: %s", string(respBody))
	}

	return nil
}

func (l *LineHandler) SendPushMessage(to string, messages LineMessage) error {
	pushMsg := &LinePushMessage{
		To:       to,
		Messages: []LineMessage{messages},
	}

	s, err := json.Marshal(pushMsg)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(s)

	req, err := http.NewRequest("POST", pushMessageAPI, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.ChannelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot send Line push message: %s", string(respBody))
	}

	return nil
}

func (l *LineHandler) SendMulticastMessage(to []string, messages LineMessage) error {
	pushMsg := &LineMulticastMessage{
		To:       to,
		Messages: []LineMessage{messages},
	}

	s, err := json.Marshal(pushMsg)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(s)

	req, err := http.NewRequest("POST", multicastMessageAPI, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.ChannelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot send Line multicast message: %s", string(respBody))
	}

	return nil
}
