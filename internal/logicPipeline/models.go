package logicpipeline

import "time"

type parId uint
type srcId uint

type DataFrameInfo struct {
	prId     parId
	srcId    srcId
	dt       time.Time
	qlt      uint8
	clas     DataClass
	updateAt time.Time
}

type DataFrameValue[T int | string | bool] struct {
	value T
	DataFrameInfo
}

type DataClass int64

const (
	analog   DataClass = iota
	discrete           = iota
	text               = iota
	event              = iota
)

type TagInfo struct {
	parId    parId
	sourceId srcId
	clas     DataClass
	options  TagOption
}

type TagOption struct {
	deltaT  float32
	deltaV  float32
	forceT  *int
	scaling string
}
