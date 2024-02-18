package main

import (
	"fmt"
	"peon.top/shex"
)

/*
struct XLSXFile {
}
*/

import "C"

//export createExcelFile
func createExcelFile(filename string, defaultSheetName string) uint8 {
	return BoolToUint8(shex.CreateXLSXFile(luaString(filename), luaString(defaultSheetName)) == nil)
}

//export cell
func cell(resourceId int, tabIndex, val string) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.Cell(luaString(tabIndex), luaString(val))
}

//export selectSheet
func selectSheet(resourceId int, sheet string) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.Select(luaString(sheet), false)
}

//export merge
func merge(resourceId int, si, ei string) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.Merge(luaString(si), luaString(ei))
}

//export save
func save(resourceId int) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.Save()
	shex.Release()
}

//export registerStyle
func registerStyle(resourceId int, style string) int {
	xlsx, _ := shex.SearchXLSFile()

	styleId, err := xlsx.RegisterStyle(luaString(style))

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return styleId
}

//export setCellStyle
func setCellStyle(resourceId int, cellX, cellY string, styleId int) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.SetCellStyle(luaString(cellX), luaString(cellY), styleId)
}

//export setColWidth
func setColWidth(resourceId int, startCol, endCol string, width float64) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.SetColWidth(luaString(startCol), luaString(endCol), width)
}

//export setRowHeight
func setRowHeight(resourceId int, row int, height float64) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.SetRowHeight(row, height)
}

//export insertPageBreak
func insertPageBreak(resourceId int, cell string) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.InsertPageBreak(luaString(cell))
}

//export setColStyle
func setColStyle(resourceId int, columns string, styleId int) {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.SetColStyle(luaString(columns), styleId)
}

//export getCellStyle
func getCellStyle(resourceId int, axis string) int {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.GetCellStyle(luaString(axis))

	return 0
}

//export appendBoardStyle
func appendBoardStyle(resourceId int, board, axis string) int {
	xlsx, _ := shex.SearchXLSFile()
	return xlsx.AppendBoardStyle(luaString(board), luaString(axis))
}

//export setPageMargins
func setPageMargins(resourceId int, top, left, right, bottom, header, footer float64) int {
	xlsx, _ := shex.SearchXLSFile()
	xlsx.SetPageMargins(top, left, right, bottom, header, footer)

	return 0
}

func luaString(str string) string {
	return fmt.Sprintf("%s", str)
}
