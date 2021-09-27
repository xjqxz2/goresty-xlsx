package shex

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

const MAX_RANDOM_LENGTH = 999999999

var (
	resources = make(map[int]*XLSXFile)
	rmu       sync.RWMutex
)

//	创建一个 XLSX 文件
//	输入文件名、默认工作表名（选填），即可在内存中创建一个 XLSX文件
//	当创建成功时，系统会返回一个 ResourceId 此ID即为文件句柄
//	接下来可使用该句柄来执行 Excel 对应的操作
func CreateXLSXFile(name string, defaultSheetName string) (handle int) {
	rmu.Lock()
	defer rmu.Unlock()

	//	随机生成一个不重复的 ResrouceId
	//	若生成失败则返回0
	resourceId, err := randomInt()
	if err != nil {
		return 0
	}

	//	生成一个 Excel 对象
	xlsx := &XLSXFile{
		ResourceId: resourceId,
		File:       excelize.NewFile(),
		Name:       name,
	}

	//	设置工作表名称，也可以不填写工作表名称
	//	当工作表名称未填写时，则使用 Sheet1 来代替
	if defaultSheetName == "" {
		//	Select Sheet
		xlsx.Select("Sheet1", false)
	} else {
		xlsx.Select(defaultSheetName, true)
	}

	//	将文件句柄与文件压入到句柄管理器中
	//	以便后期可以通过句柄查找到对应内存中的文件
	resources[resourceId] = xlsx

	//	一切OK，返回一个 XLSX 文件句柄
	return resourceId
}

//	通过 文件句柄 搜索XLSX在内存中的映射
//	若未找到文件则会返回一个错误
//	否则会返回一个 XLSXFile 对象
func SearchXLSFile(resourceId int) (*XLSXFile, error) {
	rmu.RLock()
	defer rmu.RUnlock()

	xlsx, ok := resources[resourceId]
	if !ok {
		return nil, errors.New("Not found xlsx")
	}

	return xlsx, nil
}

func Release(resourceId int) {
	rmu.Lock()
	defer rmu.Unlock()
	delete(resources, resourceId)
}

//	generate handle id
func randomInt() (int, error) {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))

	//	retry 99 count
	for i := 0; i < 99; i++ {
		val := r.Intn(MAX_RANDOM_LENGTH) + 1

		if _, ok := resources[val]; !ok {
			return val, nil
		}
	}

	return 0, errors.New("Create ResrouceId Failed")
}
