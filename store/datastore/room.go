package datastore

import (
	"errors"
	"math/rand"
	"time"

	"github.com/hawkingrei/threek/model"
)

func (ds datastore) GetRoom(rid int) (*model.Room, error) {
	var rm = new(model.Room)
	err := ds.Db.Where("Rid = ?", rid).Find(&rm).Error
	return rm, err
}

func (ds datastore) CreateRoom() (*model.Room, error) {
	var rm = new(model.Room)
	err := ds.Db.Create(&rm).Error
	return rm, err
}

func (ds datastore) JoinRoom(rid int64, username string) (*model.Room, error) {
	var rm = new(model.Room)
	err := ds.Db.Where("Rid = ?", rid).Find(&rm).Error
	if err != nil {
		return nil, err
	}
	if rm.Lord != username && rm.Loyalist != username && rm.Quisling != username && rm.Thief != username {
		rand.Seed(time.Now().UnixNano())
		var uids []int
		if rm.Lord == "" {
			uids = append(uids, 1)
		}
		if rm.Loyalist == "" {
			uids = append(uids, 2)
		}
		if rm.Quisling == "" {
			uids = append(uids, 3)
		}
		if rm.Thief == "" {
			uids = append(uids, 4)
		}
		index := rand.Intn(len(uids))
		switch index {
		case 1:
			rm.Lord = username
		case 2:
			rm.Loyalist = username
		case 3:
			rm.Quisling = username
		case 4:
			rm.Thief = username
		}
		ds.Db.Save(&rm)
		return rm, err
	}
	return rm, errors.New("has join")
}
