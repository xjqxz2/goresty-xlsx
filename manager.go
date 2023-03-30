package shex

import (
	"github.com/xuri/excelize/v2"
)

var xlsx *XLSXFile

// CreateXLSXFile 创建一个 XLSX 文件
// 输入文件名、默认工作表名（选填），即可在内存中创建一个 XLSX文件
// 当创建成功时，系统会返回一个 ResourceId 此ID即为文件句柄 (废除）
// 接下来可使用该句柄来执行 Excel 对应的操作
func CreateXLSXFile(name, defaultSheetName string) (handle int) {

	//	生成一个 Excel 对象
	xlsx = &XLSXFile{
		File:      excelize.NewFile(),
		Name:      name,
		StylePool: NewStylePool(),
	}

	//	设置工作表名称，也可以不填写工作表名称
	//	当工作表名称未填写时，则使用 Sheet1 来代替
	if defaultSheetName == "" {
		//	Select Sheet
		xlsx.Select("Sheet1", false)
	} else {
		xlsx.Select(defaultSheetName, true)
	}

	//	一切OK，返回一个 XLSX 文件句柄
	return 0
}

// SearchXLSFile 通过 文件句柄 搜索XLSX在内存中的映射
// 若未找到文件则会返回一个错误
// 否则会返回一个 XLSXFile 对象
func SearchXLSFile(resourceId int) (*XLSXFile, error) {
	return xlsx, nil
}

// Release 释放文件句柄
func Release(resourceId int) {

}
