package main

import "BDRV_300_Gonets/internal/logicpipeline"

func main() {
	dataProcessor := logicpipeline.NewDataProcessor(nil)
	dataProcessor.ProcessStrMessage(nil)
}
