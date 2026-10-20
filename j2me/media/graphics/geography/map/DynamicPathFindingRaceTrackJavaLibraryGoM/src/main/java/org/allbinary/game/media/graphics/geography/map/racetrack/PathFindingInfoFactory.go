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
    MAX_DIRECTIONS int
    PathFindingInfo *PathFindingInfo
    BasicGeographicMapGraph *BasicGeographicMapGraph
    BasicGeographicMapPathFinder *BasicGeographicMapPathFinder
    PathFinderGraphVisitorFactoryInterface IPathFinderGraphVisitorFactoryBase
}

var PathFindingInfoFactory_Instance *PathFindingInfoFactory

func (this *PathFindingInfoFactory) initFields() {
    this.MAX_DIRECTIONS = 4
    
    
    
    
}


func PathFindingInfoFactory_GetInstance() *PathFindingInfoFactory {
        return PathFindingInfoFactory_Instance
}


func PathFindingInfoFactory_InitMax(max_ int) {
        PathFindingInfoFactory_Instance = NewPathFindingInfoFactory(NewSimplePathFinderGraphVisitorFactory(1, 1, Integer_MAX_VALUE), max_)
}


func NewPathFindingInfoFactory(pathFinderGraphVisitorFactoryInterface IPathFinderGraphVisitorFactoryBase, max_ int) *PathFindingInfoFactory {
    this := &PathFindingInfoFactory{}
    this.BasePathFindingInfoFactory = *NewBasePathFindingInfoFactory()
    this.initFields()
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        PreLogUtil_Put("Using Dynamic Path Finding", this, commonStrings.GET_INSTANCE)
        this.BasicGeographicMapPathFinder = NewBasicGeographicMapPathFinder(max_)
        this.PathFindingInfo = NewPathFindingInfo(NewPathFindingNodeCostInfoFactory(max_), NewBasicArrayListS(1), NewBasicArrayListS(1))
        this.BasicGeographicMapGraph = NewBasicGeographicMapGraph(any(this.PathFindingInfo.GetPathFindingNodeCostInfoFactoryInterface()).(*PathFindingNodeCostInfoFactory))
        this.PathFinderGraphVisitorFactoryInterface = pathFinderGraphVisitorFactoryInterface
    return this
}


// throws Exception
func (this *PathFindingInfoFactory) GetInstancePathFindingInfo(geographicMapInterface IBasicGeographicMap, mapArray [][]int) *PathFindingInfo {
        var pathFindingInfo *PathFindingInfo = NewPathFindingInfo(this.PathFindingInfo.GetPathFindingNodeCostInfoFactoryInterface(), NewBasicArrayListS(1), NewBasicArrayListS(1)); _ = pathFindingInfo
        RaceTrackRoadsGeographicMapCellHistoryFactory_GetInstance().Init()
        this.Init(geographicMapInterface, pathFindingInfo, mapArray)
        this.BasicGeographicMapPathFinder.Init(geographicMapInterface, this.BasicGeographicMapGraph, this.PathFinderGraphVisitorFactoryInterface.GetInstance(geographicMapInterface))
        pathFindingInfo.SetPathFinder(this.BasicGeographicMapPathFinder)
        return pathFindingInfo
}


// throws Exception
func (this *PathFindingInfoFactory) Init(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int) {
        this.BasePathFindingInfoFactory.Init(geographicMapInterface, pathFindingInfo, mapArray)
        this.BuildPathFindingNodes(geographicMapInterface, pathFindingInfo, mapArray)
}


// throws Exception
func (this *PathFindingInfoFactory) BuildPathFindingNodesForCellPosition(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, mapArray [][]int, cellPosition IGeographicMapCellPosition) {
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var cellTypeId int = mapArray[cellPosition.GetRow()][cellPosition.GetColumn()]; _ = cellTypeId
        var geographicMapCellType IGeographicMapCellType = GeographicMapCellTypeFactory_GetInstance().Get(geographicMapInterface.GetCellTypeFromMapCellTypeInt(cellTypeId)); _ = geographicMapCellType
        var raceTrackGeographicMapCellType *RaceTrackGeographicMapCellType = any(geographicMapCellType).(*RaceTrackGeographicMapCellType); _ = raceTrackGeographicMapCellType
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var pathFindingNodeCostInfoFactoryInterface IPathFindingNodeCostInfoFactoryBase = pathFindingInfo.GetPathFindingNodeCostInfoFactoryInterface(); _ = pathFindingNodeCostInfoFactoryInterface
        var basicGeographicMapUtil *BasicGeographicMapUtil = BasicGeographicMapUtil_GetInstance(); _ = basicGeographicMapUtil
        var row int; _ = row
        var column int; _ = column
        var goingToGeographicMapCellPosition IGeographicMapCellPosition; _ = goingToGeographicMapCellPosition
        var costFromStart int; _ = costFromStart
        var costToEnd int; _ = costToEnd
        for index := int(0); index < this.MAX_DIRECTIONS; index++ {
            row = basicGeographicMapUtil.GetBorderingRow(index, cellPosition)
            column = basicGeographicMapUtil.GetBorderingColumn(index, cellPosition)
            if row < allBinaryTiledLayer.GetRows() && column < allBinaryTiledLayer.GetColumns() && row >= 0 && column >= 0 {
                goingToGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(column, row)
                costFromStart = raceTrackGeographicMapCellType.GetTravelCost()
                costToEnd = costFromStart
                pathFindingNodeCostInfoFactoryInterface.Create(geographicMapInterface, goingToGeographicMapCellPosition, cellPosition, costFromStart, costToEnd)
            }
        }
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

