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

//	Create a xls file handle in resource space
func CreateXLSXFile(name string, defaultSheetName string) (handle int) {
	rmu.Lock()
	defer rmu.Unlock()

	//	Create resourceId ,if err != nil ,maybe the resource space
	//	is full, return 0 is no
	resourceId, err := randomInt()
	if err != nil {
		return 0
	}

	xlsx := &XLSXFile{
		ResourceId: resourceId,
		File:       excelize.NewFile(),
		Name:       name,
	}

	//	Set default sheet name ,if [defaultSeheetName] is empty
	//	The sheet name is sheet1
	if defaultSheetName == "" {
		//	Select Sheet
		xlsx.Select("Sheet1")
	} else {
		xlsx.Select(defaultSheetName)
	}

	//	Save XLSFile in resource space
	resources[resourceId] = xlsx

	//	return a handle
	return resourceId
}

//	Search XLSX File by resourceId
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
