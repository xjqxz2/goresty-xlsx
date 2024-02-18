package shex

import (
	"github.com/xuri/excelize/v2"
)

const DefaultSheetName = "Sheet1"

var excelFile *XLSXFile

func CreateXLSXFile(name, sheet string) error {
	//	生成一个 Excel 对象
	excelFile = &XLSXFile{
		File:      excelize.NewFile(),
		Name:      name,
		StylePool: NewStylePool(),
	}

	//	设置工作表名称，也可以不填写工作表名称
	//	当工作表名称未填写时，则使用 Sheet1 来代替
	if sheet == "" {
		//	Select Sheet
		excelFile.Select(DefaultSheetName, false)
	} else {
		excelFile.Select(sheet, true)
	}

	return nil
}

// SearchXLSFile 通过 文件句柄 搜索XLSX在内存中的映射
// 若未找到文件则会返回一个错误
// 否则会返回一个 XLSXFile 对象
func SearchXLSFile() (*XLSXFile, error) {
	return excelFile, nil
}

// Release 释放文件句柄
func Release() {
	if err := excelFile.File.Close(); err != nil {
		return
	}

	excelFile = nil
}
