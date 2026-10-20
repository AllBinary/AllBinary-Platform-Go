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

//1.4.2
type PathFindingInfoFactory struct {
    BasePathFindingInfoFactory
    PathFindingInfo *PathFindingInfo
    PathFinder *PathFinder
}

var PathFindingInfoFactory_Instance *PathFindingInfoFactory = NewPathFindingInfoFactory()
var PathFindingInfoFactory_MAX int = 32768

func (this *PathFindingInfoFactory) initFields() {
    
    
}


func PathFindingInfoFactory_GetInstance() *PathFindingInfoFactory {
        return PathFindingInfoFactory_Instance
}


func PathFindingInfoFactory_InitMax(max_ int) {
        PathFindingInfoFactory_MAX = max_
}


func NewPathFindingInfoFactory() *PathFindingInfoFactory {
    this := &PathFindingInfoFactory{}
    this.BasePathFindingInfoFactory = *NewBasePathFindingInfoFactory()
    this.initFields()
        this.PathFindingInfo = NewPathFindingInfo(NewPathFindingNodeCostInfoFactory(PathFindingInfoFactory_MAX), NewBasicArrayListS(1), NewBasicArrayListS(1))
        this.PathFinder = NewPathFinder()
    return this
}


// throws Exception
func (this *PathFindingInfoFactory) GetInstancePathFindingInfo(geographicMapInterface IBasicGeographicMap, mapArray [][]int) *PathFindingInfo {
        var pathFindingInfo *PathFindingInfo = NewPathFindingInfo(this.PathFindingInfo.GetPathFindingNodeCostInfoFactoryInterface(), NewBasicArrayListS(1), NewBasicArrayListS(1)); _ = pathFindingInfo
        RaceTrackRoadsGeographicMapCellHistoryFactory_GetInstance().Init()
        this.Init(geographicMapInterface, pathFindingInfo, mapArray)
        this.PathFinder.Init(geographicMapInterface)
        pathFindingInfo.SetPathFinder(this.PathFinder)
        return pathFindingInfo
}


// throws Exception
func (this *PathFindingInfoFactory) Init(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int) {
        this.BasePathFindingInfoFactory.Init(geographicMapInterface, pathFindingInfo, mapArray)
        this.BuildPathFindingNodes(geographicMapInterface, pathFindingInfo, mapArray)
}


// throws Exception
func (this *PathFindingInfoFactory) BuildPathFindingNodesForCellPosition(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int, cellPosition IGeographicMapCellPosition) {
}


// throws Exception
func (this *PathFindingInfoFactory) BuildPathFindingNodes(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int) {
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var totalColumns int = allBinaryTiledLayer.GetColumns(); _ = totalColumns
        var totalRows int = allBinaryTiledLayer.GetRows(); _ = totalRows
        for column := int(0); column < totalColumns; column++ {
            for row := int(0); row < totalRows; row++ {
                this.BuildPathFindingNodesForCellPosition(geographicMapInterface, pathFindingInfo, mapArray, geographicMapCellPositionFactory.GetAt(column, row))
            }
        }
}

