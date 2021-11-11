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
	xlsx.Cell("C3", "Hello")
	xlsx.SetCellStyle("C3", "A3", styleId)

	//	No Style
	cellStyleId := xlsx.GetCellStyle("B1")
	log.Printf("Cell B1 StyleId: %d", cellStyleId)

	//	Have Style
	cellStyleId = xlsx.GetCellStyle("C3")
	log.Printf("Cell C3 StyleId: %d", cellStyleId)

	//	Have Style
	cellStyleId = xlsx.GetCellStyle("A2")
	log.Printf("Cell A2 StyleId: %d", cellStyleId)

	//	获取一个样式信息
	styleInfo := xlsx.StylePool.GetStyleInfo(cellStyleId)
	fmt.Println(styleInfo)

	//	追加一个边框样式
	styleRightRAW := `{"type": "bottom","color":"F1223F","style":1}`
	newStyleId := xlsx.AppendBoardStyle(styleRightRAW, "C3")
	fmt.Printf("新样式: %d\n", newStyleId)

	//	Save file
	xlsx.Save()
}
