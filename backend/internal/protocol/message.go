package protocol

type MessageType string

const (
	MsgTypeMove  MessageType = "move"  // 位置同期用
	MsgTypeImage MessageType = "image" // 画像バイナリ用
)

type Message struct {
	Type      MessageType `json:"type"`
	PlayerID  string      `json:"playerId,omitempty"`
	X         float64     `json:"x,omitempty"`
	Y         float64     `json:"y,omitempty"`
	Timestamp int64       `json:"ts"`
	// バイナリ部分は後続のバイト列として送出
}
