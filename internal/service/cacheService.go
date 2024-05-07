package service

import model "BDRV_300_Gonets/internal/models"

type CacheService struct {
	classDestinationById map[model.ClassId]model.ClassDestination
	tagConfigById        map[model.ParId]model.TagInfo
}

func NewCacheService(classes []model.ClassDestination, tags []model.TagInfo) *CacheService {
	return &CacheService{
		classDestinationById: createMap(classes, getClassKey),
		tagConfigById:        createMap(tags, getTagKey),
	}
}

func (service *CacheService) GetTag(id model.ParId) (*model.TagInfo, bool) {
	if val, ok := service.tagConfigById[id]; ok {
		return &val, true
	}
	return nil, false
}

func getTagKey(tag model.TagInfo) model.ParId {
	return tag.ParId
}

func getClassKey(class model.ClassDestination) model.ClassId {
	return class.Id
}

func createMap[T any, V comparable](src []T, key func(T) V) map[V]T {
	var result = make(map[V]T)
	for _, v := range src {
		result[key(v)] = v
	}
	return result
}
