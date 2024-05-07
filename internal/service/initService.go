package service

import repoInterface "BDRV_300_Gonets/internal/interface/repository"

type InitService struct {
	ClassRepository repoInterface.IClassRepository
	TagRepository   repoInterface.ITagRepository
}

func (serv *InitService) InitializeCacheServices() *CacheService {
	classDestinations := serv.ClassRepository.GetAllClassDestinations()
	tagsConfig := serv.TagRepository.GetAllTags()

	return NewCacheService(classDestinations, tagsConfig)
}
