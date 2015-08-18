package love

// Love LoveDetector view representation
type Love struct {
	ID  int32             `json:"id"`
	Msg map[string]string `json:"msg"`
	Img string            `json:"img"`
}
