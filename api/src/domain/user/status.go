package user

type Status struct {
	UserUID       string
	EnteredRoomID *int
	State         State
}

func NewStatus(u *User) *Status {
	return &Status{
		UserUID:       u.UID,
		EnteredRoomID: nil,
		State:         StateNormal,
	}
}

func (s *Status) ChangeEnteredRoom(id int) {
	s.EnteredRoomID = &id
}

func (s *Status) LeaveRoom() {
	s.EnteredRoomID = nil
}

type State string

const (
	StateNormal State = "normal"
)
