package logicpipeline

type DataPRocessor struct {
	_tagById map[parId]TagInfo
}

func (processor *DataPRocessor) ProcessIntMessage(data *DataFrameValue[int], val TagInfo) (bool, error) {
	if data == nil {
		return false, nil
	}

	if val, ok := processor._tagById[data.prId]; ok {
		println(val.clas)
		return true, nil
	}

	return false, nil
}

func (processor *DataPRocessor) ProcessStrMessage(data *DataFrameValue[string], val TagInfo) (bool, error) {
	if data == nil {
		return false, nil
	}

	if val, ok := processor._tagById[data.prId]; ok {
		println(val.clas)
		return true, nil
	}

	return false, nil
}
