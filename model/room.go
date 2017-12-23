package model

type Room struct {
	Lord     User
	Loyalist []User
	Quisling []User
	Thief    []User
}

var roomMap map[int]Room

func GetRoomMap() map[int]Room {
	tmp := Room{}
	tmp.Loyalist = make([]User, 1)
	tmp.Quisling = make([]User, 1)
	tmp.Thief = make([]User, 1)
	roomMap[4] = tmp

	tmp = Room{}
	tmp.Loyalist = make([]User, 1)
	tmp.Quisling = make([]User, 1)
	tmp.Thief = make([]User, 2)
	roomMap[5] = tmp

	tmp = Room{}
	tmp.Loyalist = make([]User, 1)
	tmp.Quisling = make([]User, 1)
	tmp.Thief = make([]User, 3)
	roomMap[6] = tmp
	return roomMap
}
