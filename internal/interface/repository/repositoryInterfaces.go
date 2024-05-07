package RepositoryInterface

import model "BDRV_300_Gonets/internal/models"

type IDataContext interface {
}

type ITagRepository interface {
	GetAllTags() []model.TagInfo
}

type IClassRepository interface {
	GetAllClassDestinations() []model.ClassDestination
}
