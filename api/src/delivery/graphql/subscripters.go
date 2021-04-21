package graphql

import (
	"sync"

	"github.com/laster18/poi/api/graph/model"
)

// Subscripter 各roomのSubscriptionデータの送受信を管理する
type Subscripter struct {
	// chan: map[userId]chan xxx
	messageChan    map[string]chan *model.Message
	userEventChan  map[string]chan model.UserEvent
	userStatusChan map[string]chan model.UserStatus
	mutex          sync.Mutex
}

func newSubscripter() *Subscripter {
	return &Subscripter{
		messageChan:    make(map[string]chan *model.Message),
		userEventChan:  make(map[string]chan model.UserEvent),
		userStatusChan: make(map[string]chan model.UserStatus),
		mutex:          sync.Mutex{},
	}
}

func (s *Subscripter) AddMessageChan(userID string, ch chan *model.Message) {
	s.mutex.Lock()
	s.messageChan[userID] = ch
	s.mutex.Unlock()
}

func (s *Subscripter) DeleteMessageChan(userID string) {
	s.mutex.Lock()
	delete(s.messageChan, userID)
	s.mutex.Unlock()
}

func (s *Subscripter) AddUserEventChan(userID string, ch chan model.UserEvent) {
	s.mutex.Lock()
	s.userEventChan[userID] = ch
	s.mutex.Unlock()
}

func (s *Subscripter) DeleteUserEventChan(userID string) {
	s.mutex.Lock()
	delete(s.userEventChan, userID)
	s.mutex.Unlock()
}

func (s *Subscripter) AddUserStatusChan(userID string, ch chan model.UserStatus) {
	s.mutex.Lock()
	s.userStatusChan[userID] = ch
	s.mutex.Unlock()
}

func (s *Subscripter) DeleteUserStatusChan(userID string) {
	s.mutex.Lock()
	delete(s.userStatusChan, userID)
	s.mutex.Unlock()
}

// PublishMessage ルーム内すべてのユーザーにメッセージを送信する
func (s *Subscripter) PublishMessage(msg *model.Message) {
	for _, c := range s.messageChan {
		c <- msg
	}
}

// PublishUserEvent ルーム内のすべてのユーザーにユーザーイベントを送信する
func (s *Subscripter) PublishUserEvent(msg model.UserEvent) {
	for _, c := range s.userEventChan {
		c <- msg
	}
}

// Subscripter 各roomのSubscripterを管理する
type Subscripters struct {
	// Data map[roomId]*Subscripter
	store map[string]*Subscripter
}

func NewSubscripters() *Subscripters {
	return &Subscripters{
		store: make(map[string]*Subscripter),
	}
}

func (s *Subscripters) Add(roomID string) {
	s.store[roomID] = newSubscripter()
}

func (s *Subscripters) Delete(roomID string) {
	delete(s.store, roomID)
}

func (s *Subscripters) Get(roomID string) (*Subscripter, bool) {
	v, ok := s.store[roomID]
	return v, ok
}
