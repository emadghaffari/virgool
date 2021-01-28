package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/emadghaffari/seeder/seeder"
	"github.com/emadghaffari/virgool/blog/conf"
	"github.com/emadghaffari/virgool/blog/database/mysql"
	"github.com/emadghaffari/virgool/blog/env"
	"github.com/emadghaffari/virgool/blog/model"
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
	if err := env.LoadGlobalConfiguration(dir + "/config.yaml"); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	conf.GlobalConfigs.MYSQL.Automigrate = false
	conf.GlobalConfigs.MYSQL.Logger = false

	if err := mysql.Database.Connect(&conf.GlobalConfigs, logrus.New()); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	db := mysql.Database.GetDatabase()
	seedPost(db)
}

func seedTags(db *gorm.DB) (tags []model.Tag) {
	for i := 0; i < 40; i++ {
		tx := db.Begin()

		tag := model.Tag{
			Name: seeder.Color() + " - #" + strconv.Itoa(rand.Intn(888888)+111111),
		}
		s := tx.Create(&tag)

		if s.Error == nil {
			tags = append(tags, tag)
			tx.Commit()
		} else {
			tx.Rollback()
		}

	}

	return
}

func seedParams(db *gorm.DB) (params []model.Param) {
	for i := 0; i < 40; i++ {
		tx := db.Begin()

		param := model.Param{
			Query: model.Query{
				Name:  "car",
				Value: seeder.Car(),
			},
		}
		s := tx.Create(&param)

		if s.Error == nil {
			params = append(params, param)
			tx.Commit()
		} else {
			tx.Rollback()
		}

	}

	return
}

func seedMedia(db *gorm.DB) (medias []model.Media) {
	for i := 0; i < 40; i++ {
		tx := db.Begin()
		title := seeder.Name()
		desc := seeder.Address()
		media := model.Media{
			URL:         seeder.Avatar(),
			Type:        "IMAGE",
			Title:       &title,
			Description: &desc,
		}
		sx := tx.Create(&media)

		if sx.Error == nil {
			medias = append(medias, media)
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}

	return
}

func seedPost(db *gorm.DB) {
	tags := seedTags(db)
	params := seedParams(db)
	medias := seedMedia(db)
	for i := 0; i < 20; i++ {
		tx := db.Begin()

		t := []*model.Tag{}
		p := []*model.Param{}
		m := []*model.Media{}

		if rand.Intn(2) == 1 {
			t = append(t, &tags[i])
		}
		if rand.Intn(2) == 1 {
			p = append(p, &params[i])
		}
		if rand.Intn(2) == 1 {
			m = append(m, &medias[i])
		}

		var status model.StatusPost

		switch rand.Intn(4) {
		case 3:
			status = model.Pending
		case 1:
			status = model.Published
		case 2:
			status = model.Deleted
		default:
			status = model.Pending
		}

		title := seeder.Title()

		post := model.Post{
			UserID:      rand.Uint64(),
			Title:       title,
			Slug:        strings.Replace(title, " ", "-", -1),
			Description: seeder.Text(),
			Text:        seeder.Text(),
			Params:      p,
			Media:       m,
			Tags:        t,
			Status:      status,
			Rate:        uint8(rand.Intn(5)),
			PublishedAT: time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		sx := tx.Create(&post)
		if sx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}
