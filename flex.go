package line

type FlexContainer interface {
	FlexContainer()
}

type CarouselContainer struct {
	Type     string            `json:"type,omitempty"`
	Contents []BubbleContainer `json:"contents"`
}

func (t CarouselContainer) FlexContainer() {}

type BubbleContainer struct {
	Type   string        `json:"type,omitempty"`
	Size   string        `json:"size,omitempty"`
	Header *BoxComponent `json:"header,omitempty"`
	Hero   FlexComponent `json:"hero,omitempty"`
	Body   *BoxComponent `json:"body,omitempty"`
	Footer *BoxComponent `json:"footer,omitempty"`
}

func (t BubbleContainer) FlexContainer() {}

type BoxComponent struct {
	Type            string           `json:"type,omitempty"`
	Layout          string           `json:"layout,omitempty"`
	Margin          string           `json:"margin,omitempty"`
	Spacing         string           `json:"spacing,omitempty"`
	BackgroundColor string           `json:"backgroundColor,omitempty"`
	BorderWidth     string           `json:"borderWidth,omitempty"`
	CornerRadius    string           `json:"cornerRadius,omitempty"`
	Flex            int              `json:"flex,omitempty"`
	Width           string           `json:"width,omitempty"`
	Height          string           `json:"height,omitempty"`
	AlignItems      string           `json:"alignItems,omitempty"`
	Contents        []FlexComponent  `json:"contents"`
	Action          *ActionComponent `json:"action,omitempty"`

	// Padding
	PaddingAll    string `json:"paddingAll,omitempty"`
	PaddingTop    string `json:"paddingTop,omitempty"`
	PaddingBottom string `json:"paddingBottom,omitempty"`
	PaddingStart  string `json:"paddingStart,omitempty"`
	PaddingEnd    string `json:"paddingEnd,omitempty"`
}

func (t BoxComponent) FlexComponent() {}

type TextComponent struct {
	Type    string `json:"type,omitempty"`
	Text    string `json:"text,omitempty"`
	Weight  string `json:"weight,omitempty"`
	Color   string `json:"color,omitempty"`
	Size    string `json:"size,omitempty"`
	Flex    int    `json:"flex,omitempty"`
	Align   string `json:"align,omitempty"`
	Margin  string `json:"margin,omitempty"`
	Wrap    bool   `json:"wrap,omitempty"`
	Gravity string `json:"gravity,omitempty"`
}

func (t TextComponent) FlexComponent() {}

type ImageComponent struct {
	Type        string `json:"type,omitempty"`
	Size        string `json:"size,omitempty"`
	Flex        int    `json:"flex,omitempty"`
	Align       string `json:"align,omitempty"`
	Gravity     string `json:"gravity,omitempty"`
	AspectMode  string `json:"aspectMode,omitempty"`
	AspectRatio string `json:"aspectRatio,omitempty"`
	Margin      string `json:"margin,omitempty"`
	Wrap        bool   `json:"wrap,omitempty"`
	URL         string `json:"url,omitempty"`
}

func (t ImageComponent) FlexComponent() {}

type SeparatorComponent struct {
	Type   string `json:"type,omitempty"`
	Margin string `json:"margin,omitempty"`
}

func (t SeparatorComponent) FlexComponent() {}

type SpacerComponent struct {
	Type string `json:"type,omitempty"`
}

func (t SpacerComponent) FlexComponent() {}

type ButtonComponent struct {
	Type   string           `json:"type,omitempty"`
	Style  string           `json:"style,omitempty"`
	Margin string           `json:"margin,omitempty"`
	Height string           `json:"height,omitempty"`
	Action *ActionComponent `json:"action,omitempty"`
}

func (t ButtonComponent) FlexComponent() {}

type FlexComponent interface {
	FlexComponent()
}

type ActionComponent struct {
	Type        string `json:"type,omitempty"`
	Label       string `json:"label,omitempty"`
	Data        string `json:"data,omitempty"`
	DisplayText string `json:"displayText,omitempty"`
	URI         string `json:"uri,omitempty"`
}

func CreateFlexMessage(message FlexContainer, altText string) LineMessage {
	return LineMessage{
		Type:     "flex",
		AltText:  altText,
		Contents: &message,
	}
}
