package ip2locationmod

import (
	"fmt"
	"ipdb/modules/config"

	"github.com/ip2location/ip2location-go/v9"
)

func NewIPdb() *ip2location.DB {
	db, err := ip2location.OpenDB(config.GlobalConfig.IP2Location.Path)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return db
}
