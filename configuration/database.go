package configuration

import (
	"strconv"
	"sync"
	"time"

	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbLinksy *gorm.DB
	mutex    sync.Mutex
)

func DB() (*gorm.DB, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if dbLinksy != nil {
		return dbLinksy, nil
	}

	dsn, err := dataSourceName()
	if err != nil {
		return nil, err
	}

	connection, err := gorm.Open(mysql.Open(dsn.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbLinksy = connection

	return dbLinksy, nil
}

func dataSourceName() (driver.Config, error) {
	config, err := LoadOrCreate()
	if err != nil {
		return driver.Config{}, err
	}

	dsn := driver.Config{
		User:      config.Database.User,
		Passwd:    config.Database.Password,
		Net:       "tcp",
		Addr:      config.Database.Host + ":" + strconv.Itoa(config.Database.Port),
		DBName:    config.Database.Name,
		ParseTime: true,
		Loc:       time.Local,
	}

	return dsn, nil
}
