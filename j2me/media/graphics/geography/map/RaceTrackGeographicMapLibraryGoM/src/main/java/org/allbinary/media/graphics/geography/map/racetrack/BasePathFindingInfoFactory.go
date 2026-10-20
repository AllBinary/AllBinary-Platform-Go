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

type BasePathFindingInfoFactory struct {
    LogUtil *LogUtil
}

// The Java class BasePathFindingInfoFactory has subclasses so its variables are this interface
type IBasePathFindingInfoFactory interface {
    BasePathFindingInfoFactory_() *BasePathFindingInfoFactory
    Init(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int)
    AddStartPathFindingNode(pathFindingInfo *PathFindingInfo, startGeographicMapCellPosition IGeographicMapCellPosition)
    AddEndPathFindingNode(pathFindingInfo *PathFindingInfo, endGeographicMapCellPosition IGeographicMapCellPosition)
    GetInstancePathFindingInfo(geographicMapInterface IBasicGeographicMap, graphArray [][]int) *PathFindingInfo
}

func (this *BasePathFindingInfoFactory) BasePathFindingInfoFactory_() *BasePathFindingInfoFactory {
    return this
}


func (this *BasePathFindingInfoFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewBasePathFindingInfoFactory() *BasePathFindingInfoFactory {
    this := &BasePathFindingInfoFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *BasePathFindingInfoFactory) Init(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int) {
        /* Java local class: move it out of the function
type RaceTrackGeographicMapCellPositionFactoryInitVisitor struct {
    StartLineId int
    FinishLineId int
}

// Implements GeographicMapCellPositionFactoryInitVisitorInterface


func (this *RaceTrackGeographicMapCellPositionFactoryInitVisitor) initFields() {
            
            
}
        

        func NewBasePathFindingInfoFactory() *BasePathFindingInfoFactory {
            this := &BasePathFindingInfoFactory{}
            this.initFields()
                var raceTrackGeographicMapCellTypeFactory IGeographicMapCellTypeFactory = this.GeographicMapInterface.GetGeographicMapCellTypeFactory(); _ = raceTrackGeographicMapCellTypeFactory
                this.StartLineId = raceTrackGeographicMapCellTypeFactory.GetStartType()
                this.FinishLineId = raceTrackGeographicMapCellTypeFactory.GetEndType()
            return this
        }
        

        // throws Exception
        func (this *BasePathFindingInfoFactory) Visit(tiledLayer IAllBinaryTiledLayer, cellPosition IGeographicMapCellPosition) {
                var row int = cellPosition.GetRow(); _ = row
                var column int = cellPosition.GetColumn(); _ = column
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
                    this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append(commonSeps.BRACKET_OPEN).Appendint(row).Append(commonSeps.BRACKET_CLOSE).Append(commonSeps.BRACKET_OPEN).Appendint(column).Append("] in [").Appendint(len(this.MapArray)).Append(commonSeps.BRACKET_CLOSE).Append(commonSeps.BRACKET_OPEN).Appendint(len(this.MapArray[0])).Append(commonSeps.BRACKET_CLOSE)), this, "visit", e)
                    panic(e)
                        }
                    }()
                    var cellTypeId int = this.MapArray[row][column]; _ = cellTypeId
                    var geographicCellType int = this.GeographicMapInterface.GetCellTypeFromMapCellTypeInt(cellTypeId); _ = geographicCellType
                    if geographicCellType == this.StartLineId {
                        this.AddStartPathFindingNode(this.PathFindingInfo, cellPosition)
                    }
                    if geographicCellType == this.FinishLineId {
                        if geographicCellType == this.StartLineId {
                            this.AddEndPathFindingNode(this.PathFindingInfo, this.GeographicMapInterface.GetGeographicMapCellPositionFactoryInterface().GetInstance(this.GeographicMapInterface, cellPosition.GetColumn(), cellPosition.GetRow(), tiledLayer.GetColumns(), tiledLayer.GetRows(), tiledLayer.GetCellWidth(), tiledLayer.GetCellHeight()))
                        } else {
                            this.AddEndPathFindingNode(this.PathFindingInfo, cellPosition)
                        }
                    }
                    var raceTrackGeographicMapCellTypeFactory IGeographicMapCellTypeFactory = this.GeographicMapInterface.GetGeographicMapCellTypeFactory(); _ = raceTrackGeographicMapCellTypeFactory
                    var geographicMapCellTypeFactory IGeographicMapCellTypeFactory = GeographicMapCellTypeFactory_GetInstance(); _ = geographicMapCellTypeFactory
                    if raceTrackGeographicMapCellTypeFactory.IsPath(geographicMapCellTypeFactory.Get(geographicCellType)) {
                        RaceTrackRoadsGeographicMapCellHistoryFactory_GetInstance().Track(cellPosition)
                    } else {
                    }
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/
        geographicMapInterface.GetGeographicMapCellPositionFactory().Visit(NewRaceTrackGeographicMapCellPositionFactoryInitVisitor())
}


// throws Exception
func (this *BasePathFindingInfoFactory) AddStartPathFindingNode(pathFindingInfo *PathFindingInfo, startGeographicMapCellPosition IGeographicMapCellPosition) {
        pathFindingInfo.AddStartPathFindingNode(NewPathFindingNode(NullUtil_GetInstance().NULL_OBJECT, startGeographicMapCellPosition))
}


// throws Exception
func (this *BasePathFindingInfoFactory) AddEndPathFindingNode(pathFindingInfo *PathFindingInfo, endGeographicMapCellPosition IGeographicMapCellPosition) {
        pathFindingInfo.AddEndPathFindingNode(NewPathFindingNode(NullUtil_GetInstance().NULL_OBJECT, endGeographicMapCellPosition))
}


// throws Exception
func (this *BasePathFindingInfoFactory) GetInstancePathFindingInfo(geographicMapInterface IBasicGeographicMap, graphArray [][]int) *PathFindingInfo {
        panic(errors.New(fmt.Sprint()))
}

