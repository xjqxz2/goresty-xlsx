package main

import (
	"fmt"
	"log"

	"peon.top/shex"
)

func main() {
	const CTestFileName = "xlsx_test.xlsx"

	fileId := shex.CreateXLSXFile(CTestFileName, "")
	xlsx, _ := shex.SearchXLSFile(fileId)

	styleRAW := `{"border": [{"type": "left","color":"0000FF","style":1}]}`
	styleId, _ := xlsx.RegisterStyle(styleRAW)

	log.Printf("Generator StyleId: %d", styleId)

	//	Set to Cell
	xlsx.Cell("A1", "Hello")
	xlsx.SetCellStyle("A1", "A3", styleId)

	//	No Style
	cellStyleId := xlsx.GetCellStyle("B1")
	log.Printf("Cell B1 StyleId: %d", cellStyleId)

	//	Have Style
	cellStyleId = xlsx.GetCellStyle("A1")
	log.Printf("Cell A1 StyleId: %d", cellStyleId)

	//	Have Style
	cellStyleId = xlsx.GetCellStyle("A2")
	log.Printf("Cell A2 StyleId: %d", cellStyleId)

	//	获取一个样式信息
	styleInfo := xlsx.StylePool.GetStyleInfo(cellStyleId)
	fmt.Println(styleInfo)

	//	Save file
	xlsx.Save()
}
