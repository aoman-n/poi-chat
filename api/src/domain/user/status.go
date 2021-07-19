package user

type Status struct {
	UserID        int
	UserUID       string
	EnteredRoomID *int
	State         State
}

func NewStatus(u *User) *Status {
	return &Status{
		UserID:        u.ID,
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
