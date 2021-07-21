package message

type Message struct {
	From            string `json:"from"`
	Msg             string `json:"msg"`
	UpdateName      bool   `json:"updateName"`
	FirstConnection bool   `jsob:"firstConnection"`
}
