package logicpipeline

import (
	iservice "BDRV_300_Gonets/internal/interface/service"
	model "BDRV_300_Gonets/internal/models"
)

type DataProcessor struct {
	cacheService iservice.ICacheService
}

func NewDataProcessor(cacheService iservice.ICacheService) *DataProcessor {
	return &DataProcessor{
		cacheService: cacheService,
	}
}

func (processor *DataProcessor) ProcessIntMessage(data *model.DataFrameValue[int]) (bool, error) {
	if data == nil {
		return false, nil
	}

	if val, ok := processor.cacheService.GetTag(data.PrId); ok {
		println(val.Class)
		return true, nil
	}

	return false, nil
}

func (processor *DataProcessor) ProcessStrMessage(data *model.DataFrameValue[string]) (bool, error) {
	if data == nil {
		return false, nil
	}

	if val, ok := processor.cacheService.GetTag(data.PrId); ok {
		println(val.Class)
		return true, nil
	}

	return false, nil
}
