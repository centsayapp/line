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

	req, err := http.NewRequest("POST", replyURL, body)
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
