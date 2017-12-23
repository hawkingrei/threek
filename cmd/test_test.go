package main

import (
	"testing"

	"github.com/WindomZ/testify/assert"
	"github.com/hawkingrei/threek/conf"
)

func TestDatabase(t *testing.T) {
	var conf conf.Configure
	conf.DbDriver = "mysql"
	conf.DbURL = "root:123123@/threek?charset=utf8&parseTime=True&loc=Local"
	conf.Debug = true

	_, err := setupStore(&conf)
	assert.NotEqual(t, nil, err)
}
