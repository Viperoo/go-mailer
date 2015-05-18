package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Mails struct {
	Id        int16
	Name      string
	Email     string
	Mailing   Mailing
	MailingId int16
	Send      int16
	SendAt    time.Time
}

type Mailing struct {
	Id        int16
	Subject   string
	From      string
	Template  string
	Status    int16
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Date string

var db gorm.DB

func conn() (db gorm.DB) {
	db, err := gorm.Open("sqlite3", workdir+"/go-mailer.db")

	if err != nil {
		logger.Critical(err.Error())
	}

	return db
}

func makeMigrate(db gorm.DB) {
	db.DropTableIfExists(&Mailing{})
	db.DropTableIfExists(&Mails{})

	db.CreateTable(&Mailing{})
	db.CreateTable(&Mails{})

	logger.Info("Migration complete")
}
