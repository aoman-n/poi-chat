package domain

// for Redis
type RoomUser struct {
	ID          int    `json:"id"`
	RoomID      int    `json:"roomId"`
	Name        string `json:"name"`
	AvatarURL   string `json:"avatarUrl"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	LastMessage string `json:"lastMessage"`
	LastEvent   RoomUserEvent
}

type RoomUserEvent int

const (
	// Eventを保持していない場合にはNoneEventを入れる
	JoinEvent RoomUserEvent = iota + 1
	MoveEvent
	MessageEvent
)

// type RoomUserMoveEvent struct {
// 	RoomUserID int
// 	X          int
// 	Y          int
// }
// type RoomUserJoinEvent struct{}
// type RoomUserExitEvent struct{}
// type RoomUserMessageEvent struct{}
