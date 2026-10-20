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

type StaticPathGenerator struct {
    LogUtil *LogUtil
    BasicArrayListUtil *BasicArrayListUtil
}

// Implements PathGeneratorInterface


func (this *StaticPathGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewStaticPathGenerator() *StaticPathGenerator {
    this := &StaticPathGenerator{}
    this.initFields()
        PreLogUtil_Put("Using Static Path Finding", this, CommonStrings_GetInstance().CONSTRUCTOR)
    return this
}


// throws Exception
func (this *StaticPathGenerator) Init(geographicMapInterface any, totalPaths int) {
}


// throws Exception
func (this *StaticPathGenerator) GetGeographicMapCellPositionListFromBasicGeographicMapCellPositionList(geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory, pathList IBasicArrayList) IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var size int = pathList.Size(); _ = size
        var basicGeographicMapCellPosition ICellPosition; _ = basicGeographicMapCellPosition
        var geographicMapCellPosition IGeographicMapCellPosition; _ = geographicMapCellPosition
        for index := int(0); index < size; index++ {
            basicGeographicMapCellPosition = any(pathList.Get(index)).(ICellPosition)
            geographicMapCellPosition = geographicMapCellPositionFactory.GetAt(basicGeographicMapCellPosition.GetColumn(), basicGeographicMapCellPosition.GetRow())
            list.Add(geographicMapCellPosition)
        }
        return list
}


// throws Exception
func (this *StaticPathGenerator) CreateList(geographicMapInterface IBasicGeographicMap, geographicMapCellHistory *GeographicMapCellHistory, pathFindingInfo *PathFindingInfo, totalPaths int) IBasicArrayList {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            tryControl1, tryResult1 = 1, NewBasicArrayListD(); return
                }
            }()
            var pathCacheFactory *PathCacheFactory = PathCacheFactory_GetInstance(); _ = pathCacheFactory
            var mapIdInteger int = geographicMapInterface.GetAllBinaryTiledLayer().GetDataId(); _ = mapIdInteger
            var list IBasicArrayList = pathCacheFactory.GetOrCreate(mapIdInteger); _ = list
            if list == this.BasicArrayListUtil.GetImmutableInstance() {
                list = NewBasicArrayListD()
                var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
                var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
                var id int = PathData_GetInstance().OFFSET + int(mapIdInteger); _ = id
                var basicList IBasicArrayList = pathCacheFactory.GetOrCreate(smallIntegerSingletonFactory.GetAt(id)); _ = basicList
                var size int = basicList.Size(); _ = size
                var pathList IBasicArrayList; _ = pathList
                for index := int(0); index < size; index++ {
                    pathList = this.GetGeographicMapCellPositionListFromBasicGeographicMapCellPositionList(geographicMapCellPositionFactory, any(basicList.Get(index)).(IBasicArrayList))
                    list.Add(pathList)
                }
                pathCacheFactory.Add(mapIdInteger, list)
                pathCacheFactory.Remove(smallIntegerSingletonFactory.GetAt(id))
            }
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Using Cached Path(s): ").Append(StringUtil_GetInstance().ToString(list))), this, commonStrings.GET_INSTANCE)
            tryControl1, tryResult1 = 1, list; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

