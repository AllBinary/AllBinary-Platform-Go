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

type GeographicMapCellHistoryProgress struct {
    GeographicMapCellHistoryProgressBase
    TotalVisitedArray []int
    GeographicMapCellHistoryArray []*GeographicMapCellHistory
}


func (this *GeographicMapCellHistoryProgress) initFields() {
    this.TotalVisitedArray = NullUtil_GetInstance().NULL_INT_ARRAY
    this.GeographicMapCellHistoryArray = GeographicMapCellHistory_NULL_GEOGRPAHIC_MAP_HISTORY_ARRAY
}


func NewGeographicMapCellHistoryProgress(geographicMapCellHistoryArray []*GeographicMapCellHistory) *GeographicMapCellHistoryProgress {
    this := &GeographicMapCellHistoryProgress{}
    this.GeographicMapCellHistoryProgressBase = *NewGeographicMapCellHistoryProgressBase()
    this.initFields()
        this.GeographicMapCellHistoryArray = geographicMapCellHistoryArray
        this.Init()
    return this
}


func (this *GeographicMapCellHistoryProgress) Init() {
        this.TotalVisitedArray = make([]int, len(this.GeographicMapCellHistoryArray))
}


func (this *GeographicMapCellHistoryProgress) IsAnyProgress() bool {
        var isProgressing bool = false; _ = isProgressing
        var size int = len(this.GeographicMapCellHistoryArray); _ = size
        var currentTotalVisited int; _ = currentTotalVisited
        var totalVisited int; _ = totalVisited
        for index := int(0); index < size; index++ {
            totalVisited = this.GeographicMapCellHistoryArray[index].GetTotalVisited()
            currentTotalVisited = this.TotalVisitedArray[index]
            if currentTotalVisited < totalVisited {
                isProgressing = true
                this.TotalVisitedArray[index] = totalVisited
            }
        }
        return isProgressing
}

