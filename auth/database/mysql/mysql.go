package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/model"
)

var (
	// Database var
	Database  Mysql  = &msql{}
	namespace string = ""
	err       error
	once      sync.Once
)

// Mysql interface
type Mysql interface {
	Connect(config *conf.GlobalConfiguration, log logrus.FieldLogger) error
	GetDatabase() *gorm.DB
	AutoMigrate() error
}

type msql struct {
	DB *gorm.DB
}

func (m *msql) Connect(config *conf.GlobalConfiguration, log logrus.FieldLogger) error {
	once.Do(func() {
		if config.DB.Namespace != "" {
			namespace = config.DB.Namespace
		}

		datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Schema,
		)

		m.DB, err = gorm.Open(config.DB.Driver, datasource)
		if err != nil {
			log.Fatal(errors.Wrap(err, "opening database connection"))
		}

		m.DB.SetLogger(model.NewDBLogger(log))
		m.DB.LogMode(true)

		err = m.DB.DB().Ping()
		if err != nil {
			log.Fatal(errors.Wrap(err, "checking database connection"))
		}

		if config.DB.Automigrate {
			migDB := m.DB.New()
			migDB.SetLogger(model.NewDBLogger(log.WithField("task", "migration")))
			if err := m.AutoMigrate(); err != nil {
				log.Fatal(errors.Wrap(err, "database automigrate"))
			}
		}

	})

	return nil
}

func (m *msql) AutoMigrate() error {
	sql := m.DB.AutoMigrate(
		model.User{},
		model.Role{},
		model.Permission{},
	)
	return sql.Error
}

func (m *msql) GetDatabase() *gorm.DB {
	return m.DB
}
