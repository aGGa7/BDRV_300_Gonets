package gonetzmodel

import "time"

type ParId uint
type SrcId uint
type ClassId uint

type DataFrameInfo struct {
	PrId     ParId
	SrcId    SrcId
	Dt       time.Time
	Qlt      uint8
	Class    DataClass
	UpdateAt time.Time
}

type DataFrameValue[T int | string | bool] struct {
	Value T
	DataFrameInfo
}

type DataClass int64

const (
	Analog   DataClass = iota
	Discrete           = iota
	Text               = iota
	Event              = iota
)

type TagInfo struct {
	ParId    ParId
	SourceId SrcId
	Class    DataClass
	Options  TagOption
}

type TagOption struct {
	DeltaT  float32
	DeltaV  float32
	ForceT  *int
	Scaling string
}

type ClassDestination struct {
	Id    ClassId
	Table string
	Class DataClass
}
