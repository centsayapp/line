package line

type LineUserProfile struct {
	UID           string `json:"userId"`
	DisplayName   string `json:"displayName"`
	Language      string `json:"language,omitempty"`
	PictureURL    string `json:"pictureUrl,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}
