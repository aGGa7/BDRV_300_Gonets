package serviceInterface

import model "BDRV_300_Gonets/internal/models"

type ICacheService interface {
	GetTag(id model.ParId) (*model.TagInfo, bool)
	GetClass(id model.ClassId) (*model.ClassDestination, bool)
}
