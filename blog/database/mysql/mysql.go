package mysql

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/model"
)

var (
	// Database var
	Database  Mysqli = &msql{}
	namespace string = ""
	err       error
	once      sync.Once
)

// Mysqli interface
type Mysqli interface {
	Connect(config *conf.GlobalConfiguration, log logrus.FieldLogger) error
	GetDatabase() *gorm.DB
	AutoMigrate() error
}

type msql struct {
	DB *gorm.DB
}

func (m *msql) Connect(config *conf.GlobalConfiguration, log logrus.FieldLogger) error {
	once.Do(func() {
		if config.MYSQL.Namespace != "" {
			namespace = config.MYSQL.Namespace
		}

		conf := &gorm.Config{}

		if config.MYSQL.Logger {
			conf.Logger = model.NewDBLogger()
		}

		datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.MYSQL.Username,
			config.MYSQL.Password,
			config.MYSQL.Host,
			config.MYSQL.Schema,
		)

		m.DB, err = gorm.Open(mysql.Open(datasource), conf)
		if err != nil {
			log.Fatal(errors.Wrap(err, "opening database connection"))
			return
		}

		if config.MYSQL.Automigrate {
			if err := m.AutoMigrate(); err != nil {
				log.Fatal(errors.Wrap(err, "database automigrate"))
				return
			}
		}

	})

	return err
}

func (m *msql) AutoMigrate() error {
	sql := m.DB.AutoMigrate(
		model.Post{},
		model.Tag{},
		model.Param{},
		model.Media{},
	)
	return sql
}

func (m *msql) GetDatabase() *gorm.DB {
	return m.DB
}
