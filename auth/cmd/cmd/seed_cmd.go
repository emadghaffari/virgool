package cmd

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/emadghaffari/seeder/seeder"
	"github.com/emadghaffari/virgool/auth/conf"
	"github.com/emadghaffari/virgool/auth/database/mysql"
	"github.com/emadghaffari/virgool/auth/env"
	"github.com/emadghaffari/virgool/auth/model"
)

var seedCMD = cobra.Command{
	Use:  "seed database",
	Long: "seed database strucutures. This will seed tables",
	Run:  seed,
}

func seed(cmd *cobra.Command, args []string) {

	// Current working directory
	dir, err := os.Getwd()
	if err != nil {
		logrus.Warn(err.Error())
	}

	// read from file
	if err := env.LoadGlobalConfiguration(dir + "/auth/config.yaml"); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	conf.GlobalConfigs.MYSQL.Automigrate = false
	conf.GlobalConfigs.MYSQL.Logger = false

	if err := mysql.Database.Connect(&conf.GlobalConfigs, conf.Logger); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	db := mysql.Database.GetDatabase()
	ids := rolePermission(db)
	users(db, ids)
}

func rolePermission(db *gorm.DB) (ids []uint64) {
	for i := 0; i < 2; i++ {
		permissions := []*model.Permission{}
		for i := 0; i < 10; i++ {
			permissions = append(permissions, &model.Permission{
				Name:      seeder.Name(),
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			})
		}
		role := &model.Role{
			Name:        seeder.Name(),
			Permissions: permissions,
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		}
		db.Create(role)
		ids = append(ids, role.ID)
	}
	return ids
}

func users(db *gorm.DB, ids []uint64) {
	for i := 0; i < 200; i++ {
		pass := seeder.Password()
		db.Create(&model.User{
			Username:  seeder.Username(),
			Password:  &pass,
			Phone:     seeder.Phone(),
			Name:      seeder.Name(),
			LastName:  seeder.Name(),
			Email:     seeder.Email(),
			RoleID:    seeder.RandomArray(ids).(reflect.Value).Uint(),
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		})
	}
}
