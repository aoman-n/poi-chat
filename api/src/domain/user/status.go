package user

type Status struct {
	EnteredRoomID *int
	State         State
}

func NewStatus() *Status {
	return &Status{
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
