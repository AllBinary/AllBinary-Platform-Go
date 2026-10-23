/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type InventoryColumnUtil struct {
    LogUtil *LogUtil
    BasicItemData *BasicItemData
}

var InventoryColumnUtil_Instance *InventoryColumnUtil = NewInventoryColumnUtil()

func (this *InventoryColumnUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
}

func NewInventoryColumnUtil() *InventoryColumnUtil {
    this := &InventoryColumnUtil{}
    this.initFields()
    return this
}


func InventoryColumnUtil_GetInstance() *InventoryColumnUtil {
        return InventoryColumnUtil_Instance
}


func (this *InventoryColumnUtil) GetColumnWhereLike(inventoryEntity *InventoryEntity, category string, column string) IBasicArrayList {
        this.LogUtil.PutF("Start Category: " + category, this, "getKeywords")
        var keywords IBasicArrayList = NewBasicArrayListD(); _ = keywords
        var vectorOfHashMaps IBasicArrayList = inventoryEntity.GetAllRows(); _ = vectorOfHashMaps
        var size int = vectorOfHashMaps.Size(); _ = size
        for index := int(0); index < size; index++ {
            var hashMap *HashMap = any(vectorOfHashMaps.Get(index)).(*HashMap); _ = hashMap
            var categoryName string = any(hashMap.Get(this.BasicItemData.CATEGORY)).(string); _ = categoryName
            if !StringValidationUtil_GetInstance().IsEmpty(categoryName) && strings.HasPrefix(categoryName, category) {
                keywords.Add(hashMap.Get(column))
            }
        }
        return keywords
}

