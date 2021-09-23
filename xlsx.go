package shex

import (
	"sync"

	"github.com/xuri/excelize/v2"
)

type XLSXFile struct {
	ResourceId int
	Sheet      string
	Name       string
	File       *excelize.File
	Sheets     sync.Map
}

//	Select a sheet
func (p *XLSXFile) Select(sheet string) *XLSXFile {
	index := 0

	if value, ok := p.Sheets.Load(sheet); ok {
		index = value.(int)
		p.File.SetActiveSheet(index)
	} else {
		index = p.File.NewSheet(sheet)
		p.Sheets.Store(sheet, index)
	}

	p.File.SetActiveSheet(index)
	p.Sheet = sheet

	return p
}

//
func (p *XLSXFile) Cell(index, value string) *XLSXFile {
	p.File.SetCellValue(p.Sheet, index, value)
	return p
}

func (p *XLSXFile) Merge(start, end string) *XLSXFile {
	p.File.MergeCell(p.Sheet, start, end)
	return p
}

func (p *XLSXFile) Save() bool {
	return p.File.SaveAs(p.Name) == nil
}
