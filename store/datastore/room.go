package datastore

import "github.com/hawkingrei/threek/model"

func (ds datastore) GetRoom(rid int) (*model.Room, error) {
	var rm = new(model.Room)
	err := ds.Db.Where("roomid = ?", rid).Find(&rm).Error
	return rm, err
}

func (ds datastore) CreateRoom() (*model.Room, error) {
	var rm = new(model.Room)
	err := ds.Db.Create(&rm).Error
	return rm, err
}
