package subscriber

import "fmt"

// TODO: このファイルはresolverで使っているものと共通化する

func makeRoomUserID(userUID string) string {
	return fmt.Sprintf("RoomUser:%s", userUID)
}

func makeMessageID(id int) string {
	return fmt.Sprintf("Messaage:%d", id)
}

func makeUserID(userUID string) string {
	return fmt.Sprintf("User:%s", userUID)
}
