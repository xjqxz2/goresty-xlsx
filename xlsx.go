package shex

import (
	"encoding/json"
	"sync"

	"github.com/xuri/excelize/v2"
)

type XLSXFile struct {
	ResourceId int
	Sheet      string
	Name       string
	File       *excelize.File
	Sheets     sync.Map
	StylePool  *StylePool
}

// Select
// a sheet
func (p *XLSXFile) Select(sheet string, replaceDefaultSheetName bool) *XLSXFile {
	index := 0

	if value, ok := p.Sheets.Load(sheet); ok {
		index = value.(int)
	} else {
		index, _ = p.File.NewSheet(sheet)
		p.Sheets.Store(sheet, index)

		if replaceDefaultSheetName {
			p.File.DeleteSheet("Sheet1")
		}
	}

	p.File.SetActiveSheet(index)
	p.Sheet = sheet

	return p
}

// Cell 设置一个单元格值
func (p *XLSXFile) Cell(index, value string) *XLSXFile {
	p.File.SetCellValue(p.Sheet, index, value)
	return p
}

func (p *XLSXFile) Merge(start, end string) *XLSXFile {
	p.File.MergeCell(p.Sheet, start, end)
	return p
}

// 注册一个新样式
func (p *XLSXFile) RegisterStyle(ss string) (styleId int, err error) {
	var style excelize.Style

	//	解析样式
	if err := json.Unmarshal([]byte(ss), &style); err != nil {
		return 0, err
	}

	//	注册样式并生成一个 StyleId
	styleId, err = p.File.NewStyle(&style)
	if err != nil {
		return 0, err
	}

	//	并将样式信息保存到内存中
	go p.StylePool.Push(styleId, &style)

	//	返回样式信息
	return styleId, nil
}

func (p *XLSXFile) SetCellStyle(cellX, cellY string, styleId int) {
	_ = p.File.SetCellStyle(p.Sheet, cellX, cellY, styleId)
}

func (p *XLSXFile) SetColWidth(startCol, endCol string, width float64) {
	_ = p.File.SetColWidth(p.Sheet, startCol, endCol, width)
}

func (p *XLSXFile) SetRowHeight(row int, height float64) {
	_ = p.File.SetRowHeight(p.Sheet, row, height)
}

func (p *XLSXFile) Save() bool {
	return p.File.SaveAs(p.Name) == nil
}

func (p *XLSXFile) InsertPageBreak(cell string) {
	_ = p.File.InsertPageBreak(p.Sheet, cell)
}

func (p *XLSXFile) SetColStyle(columns string, styleId int) {
	_ = p.File.SetColStyle(p.Sheet, columns, styleId)
}

func (p *XLSXFile) GetCellStyle(axis string) int {
	styleId, _ := p.File.GetCellStyle(p.Sheet, axis)
	return styleId
}

// 追加一个样式
// 它会生成一个新的样式ID ，并且会把新的样式返回给调用者
func (p *XLSXFile) AppendBoardStyle(style, axis string) int {
	//	首先先把 Border 的样式先解析出来
	var border excelize.Border
	json.Unmarshal([]byte(style), &border)

	//	先获取单元格原样式
	styleId := p.GetCellStyle(axis)

	//	再在样式池中查找这个样式ID对应的样式表
	origin := p.StylePool.GetStyleInfo(styleId)

	if origin.Border == nil {
		origin.Border = []excelize.Border{}
	}
	//	将线条样式追加至原样式中
	origin.Border = append(origin.Border, border)

	//	重新注册样式
	//	注册样式并生成一个 StyleId
	newStyleId, err := p.File.NewStyle(origin)
	if err != nil {
		return 0
	}

	//	并将样式信息保存到内存中
	go p.StylePool.Push(newStyleId, origin)

	//	设置单元格样式
	_ = p.File.SetCellStyle(p.Sheet, axis, axis, newStyleId)

	//	返回样式信息
	return newStyleId
}

func (p *XLSXFile) SetPageMargins(top, left, right, bottom, header, footer float64) {
	var opts excelize.PageLayoutMarginsOptions

	if top > -1 {
		opts.Top = &top
	}

	if left > -1 {
		opts.Left = &left
	}

	if right > -1 {
		opts.Right = &right
	}

	if bottom > -1 {
		opts.Bottom = &bottom
	}

	if header > -1 {
		opts.Header = &header
	}

	if footer > -1 {
		opts.Footer = &footer
	}

	_ = p.File.SetPageMargins(p.Sheet, &opts)
}
