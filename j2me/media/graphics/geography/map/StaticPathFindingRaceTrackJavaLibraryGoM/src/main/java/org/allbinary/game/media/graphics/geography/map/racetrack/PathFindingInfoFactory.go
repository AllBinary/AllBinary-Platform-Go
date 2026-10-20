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

type PathFindingInfoFactory struct {
    BasePathFindingInfoFactory
}

var PathFindingInfoFactory_Instance *PathFindingInfoFactory = NewPathFindingInfoFactory()


func PathFindingInfoFactory_GetInstance() *PathFindingInfoFactory {
        return PathFindingInfoFactory_Instance
}


func NewPathFindingInfoFactory() *PathFindingInfoFactory {
    this := &PathFindingInfoFactory{}
    this.BasePathFindingInfoFactory = *NewBasePathFindingInfoFactory()
    return this
}


// throws Exception
func (this *PathFindingInfoFactory) GetInstancePathFindingInfo(geographicMapInterface IBasicGeographicMap, graphArray [][]int) *PathFindingInfo {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        PreLogUtil_Put(commonStrings.START, this, commonStrings.GET_INSTANCE)
        var pathFindingInfo *PathFindingInfo = NewPathFindingInfo(NewPathFindingNodeCostInfoFactoryBase(), NewBasicArrayListS(1), NewBasicArrayListS(1)); _ = pathFindingInfo
        RaceTrackRoadsGeographicMapCellHistoryFactory_GetInstance().Init()
        this.Init(geographicMapInterface, pathFindingInfo, graphArray)
        return pathFindingInfo
}

