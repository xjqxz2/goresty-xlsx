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
func createExcelFile(filename string) int {
	return shex.CreateXLSXFile(filename)
}

//export cell
func cell(resourceId int, tabIndex, val string) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.Cell(tabIndex, val)
}

//export selectSheet
func selectSheet(resourceId int, sheet string) {
	xlsx, _ := shex.SearchXLSFile(resourceId)
	xlsx.Select(sheet)
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

func main() {

}
