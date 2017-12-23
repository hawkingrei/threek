package store

import (
	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hawkingrei/threek/model"
)

type Store interface {
	// GetUser gets a user by unique Username.
	GetUser(string) (*model.User, error)
	CreateUser(username string)
	HasUser(string) bool
	CreateTable(models ...interface{})

	CreateRoom() (*model.Room, error)
	GetRoom(rid int) (*model.Room, error)
	JoinRoom(rid int64, username string) (*model.Room, error)
	Close()
}

func GetUser(c context.Context, username string) (*model.User, error) {
	return FromContext(c).GetUser(username)
}

func CreateUser(c context.Context, username string) {
	FromContext(c).CreateUser(username)
}

func HasUser(c context.Context, username string) bool {
	return FromContext(c).HasUser(username)
}

func CreateRoom(c context.Context) (*model.Room, error) {
	return FromContext(c).CreateRoom()
}

func GetRoom(c context.Context, rid int) (*model.Room, error) {
	return FromContext(c).GetRoom(rid)
}

func JoinRoom(c context.Context, rid int64, username string) (*model.Room, error) {
	return FromContext(c).JoinRoom(rid, username)
}
