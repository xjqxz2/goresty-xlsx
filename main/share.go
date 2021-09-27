package main

import "C"
import (
	"fmt"

	"peon.top/shex"
)

//export sum
func sum(a, b int) int {
	return a + b
}

//export println
func println(str string) {
	fmt.Println(str)
}

//export createExcelFile
func createExcelFile(filename string, defaultSheetName string) int {
	return shex.CreateXLSXFile(filename, defaultSheetName)
}

//export cell
func cell(resourceId int, tabIndex, val string) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.Cell(tabIndex, val)
}

//export selectSheet
func selectSheet(resourceId int, sheet string) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.Select(sheet, false)
}

//export merge
func merge(resourceId int, si, ei string) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.Merge(si, ei)
}

//export save
func save(resourceId int) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.Save()
	shex.Release(resourceId)
}

//export registerStyle
func registerStyle(resourceId int, style string) int {
	xlsx, _ := shex.SearchXLSFile(resourceId)

	styleId, err := xlsx.RegisterStyle(style)

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return styleId
}

//export setCellStyle
func setCellStyle(resourceId int, cellX, cellY string, styleId int) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.SetCellStyle(cellX, cellY, styleId)
}

//export setColWidth
func setColWidth(resourceId int, startCol, endCol string, width float64) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.SetColWidth(startCol, endCol, width)
}

//export setRowHeight
func setRowHeight(resourceId int, row int, height float64) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.SetRowHeight(row, height)
}

func main() {

}
