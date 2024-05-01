package models

import "time"

type DataFrame struct {
	prId     uint
	srcId    uint
	dt       time.Time
	qlt      uint8
	clas     DataClass
	updateAt time.Time
}

type DataClass int64

const (
	analog   DataClass = iota
	discrete           = iota
	string             = iota
	event              = iota
)
