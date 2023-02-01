package line

const (
	replyURL string = "https://api.line.me/v2/bot/message/reply"
)

type LineWebhook struct {
	Destination string      `json:"destination"`
	Events      []LineEvent `json:"events"`
}

type LineEventType string

const LineEventTypeMessage LineEventType = "message"

type LineEvent struct {
	Type            LineEventType    `json:"type"` //values: "message"
	Mode            string           `json:"mode"` //value: "active", "standby"
	Timestamp       int64            `json:"timestamp"`
	Source          *LineEventSource `json:"source"`
	WebhookEventID  string           `json:"webhookEventId"`
	DeliveryContext *DeliveryContext `json:"deliveryContext,omitempty"`

	// Message event
	ReplyToken string           `json:"replyToken,omitempty"`
	Message    *LineMessage     `json:"message,omitempty"`
	Postback   *PostbackMessage `json:"postback,omitempty"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type PostbackMessage struct {
	Data   string                 `json:"data,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type LineMessageType string

const LineMessageTypeText LineMessageType = "text"

type LineMessage struct {
	Type LineMessageType `json:"type,omitempty"` //values: "text"
	ID   string          `json:"id,omitempty"`

	//Text message
	Text    string              `json:"text,omitempty"`
	Emojis  []*LineMessageEmoji `json:"emojis,omitempty"`
	Mention *LineMessageMention `json:"mention,omitempty"`

	//Flex message
	AltText  string         `json:"altText,omitempty"`
	Contents *FlexContainer `json:"contents,omitempty"`
}

type LineMessageMention struct {
	Mentionees []*LineMessageMentionee `json:"mentionees"`
}

type LineMessageMentionee struct {
	Index  int    `json:"index"`
	Length int    `json:"length"`
	UserID string `json:"userId,omitempty"`
}

type LineMessageEmoji struct {
	Index     int    `json:"index"`
	Length    int    `json:"length,omitempty"`
	ProductID string `json:"productId,omitempty"`
	EmojiID   string `json:"emojiId,omitempty"`
}

type LineEventSourceType string

const LineEventSourceTypeUser LineEventSourceType = "user"
const LineEventSourceTypeRoom LineEventSourceType = "room"
const LineEventSourceTypeGroup LineEventSourceType = "group"

type LineEventSource struct {
	Type    LineEventSourceType `json:"type"` //values: "user", "group", "room"
	UserID  string              `json:"userId,omitempty"`
	GroupID string              `json:"groupId,omitempty"`
	RoomID  string              `json:"roomId,omitempty"`
}

type LineReplyMessage struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []LineMessage `json:"messages"`
}
