package shex

import (
	"sync"

	"github.com/xuri/excelize/v2"
)

type StylePool struct {
	sync.RWMutex
	styles map[int]*excelize.Style
}

//	构建一个样式管理器
func NewStylePool() *StylePool {
	return &StylePool{
		styles: make(map[int]*excelize.Style),
	}
}

//	将样式数据压入样式库中
func (p *StylePool) Push(styleId int, styleInfo *excelize.Style) *StylePool {
	p.Lock()
	defer p.Unlock()

	p.styles[styleId] = styleInfo

	return p
}

//	获取样式信息
func (p *StylePool) GetStyleInfo(styleId int) *excelize.Style {
	p.RLock()
	defer p.RUnlock()

	//	获取 StyleInfo
	if styleInfo, ok := p.styles[styleId]; ok {
		return styleInfo
	}

	return &excelize.Style{}
}
