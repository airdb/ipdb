package ipipmod

import (
	"fmt"
	"ipdb/modules/config"
)

type DB struct {
	DB *City
}

func NewDB() *DB {
	db, err := NewCity(config.GlobalConfig.IPIP.Path)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return &DB{
		DB: db,
	}
}
