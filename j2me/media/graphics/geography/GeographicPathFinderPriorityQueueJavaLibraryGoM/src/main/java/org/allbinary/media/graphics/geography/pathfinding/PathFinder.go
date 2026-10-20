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

type PathFinder struct {
    GeographicPathFinderBase
    LogUtil *LogUtil
    BasicArrayListUtil *BasicArrayListUtil
    MathUtil *MathUtil
    OpenPriorityQueue *PriorityQueue[*PathFindingNodeCost]
    ClosedSet *HashSet[*PathFindingNodeCost]
    GeographicMapInterface any
    CostArray [][]*PathFindingNodeCost
}


func (this *PathFinder) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    this.MathUtil = MathUtil_GetInstance()
    this.OpenPriorityQueue = NewPriorityQueue[*PathFindingNodeCost]()
    this.ClosedSet = NewHashSet[*PathFindingNodeCost]()
    this.GeographicMapInterface = NullUtil_GetInstance().NULL_OBJECT
    this.CostArray = PathFindingNodeCost_NULL_PATH_FINDING_NODE_COST_ARRAY_ARRAY
}

func NewPathFinder() *PathFinder {
    this := &PathFinder{}
    this.GeographicPathFinderBase = *NewGeographicPathFinderBase()
    this.initFields()
    return this
}


// throws Exception
func (this *PathFinder) Init(geographicMapInterface IBasicGeographicMap) {
        this.GeographicMapInterface = geographicMapInterface
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        this.CostArray = func() [][]*PathFindingNodeCost { array := make([][]*PathFindingNodeCost, tiledLayer.GetColumns()); for index := range array { array[index] = make([]*PathFindingNodeCost, tiledLayer.GetRows()) }; return array }()
        var basicGeographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = basicGeographicMapCellPositionFactory
        var node *PathFindingNodeCost; _ = node
        var sizeX int = len(this.CostArray); _ = sizeX
        var sizeY int = len(this.CostArray[0]); _ = sizeY
        for column := int(0); column < sizeX; column++ {
            for row := int(0); row < sizeY; row++ {
                var geographicMapCellType IGeographicMapCellType = geographicMapInterface.GetCellTypeAt(basicGeographicMapCellPositionFactory.GetAt(column, row)); _ = geographicMapCellType
                var raceTrackGeographicMapCellType *RaceTrackGeographicMapCellType = any(geographicMapCellType).(*RaceTrackGeographicMapCellType); _ = raceTrackGeographicMapCellType
                node = NewPathFindingNodeCost(NullUtil_GetInstance().NULL_OBJECT, basicGeographicMapCellPositionFactory.GetAt(column, row), NewPathFindingNodeCostInfo(int64(raceTrackGeographicMapCellType.GetTravelCost()), int64(-1)))
                this.CostArray[column][row] = node
            }
        }
}


func (this *PathFinder) SearchTotalPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, totalPaths int) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, "search", e)
            tryControl1, tryResult1 = 1, this.BasicArrayListUtil.GetImmutableInstance(); return
                }
            }()
            tryControl1, tryResult1 = 1, this.Search(any(startPathFindingNodeList.Get(0)).(IPathFindingNode), any(endPathFindingNodeList.Get(0)).(IPathFindingNode)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *PathFinder) SearchTotalPathN(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, totalPaths int, multipassState *MultipassState) IBasicArrayList {
        if multipassState.Step == 0 {
            multipassState.Step++
        } else if multipassState.Step == 1 {
            this.SearchStart(any(startPathFindingNodeList.Get(0)).(IPathFindingNode), any(endPathFindingNodeList.Get(0)).(IPathFindingNode), multipassState)
        } else if multipassState.Step == 2 {
            return this.SearchN(any(startPathFindingNodeList.Get(0)).(IPathFindingNode), any(endPathFindingNodeList.Get(0)).(IPathFindingNode), multipassState)
        }
        return this.BasicArrayListUtil.GetImmutableInstance()
}


// throws Exception
func (this *PathFinder) Search(startPathFindingNode IPathFindingNode, endPathFindingNode IPathFindingNode) IBasicArrayList {
        var list IBasicArrayList = this.FindPath(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition, endPathFindingNode.PathFindingNode_().GeographicMapCellPosition); _ = list
        var pathList IBasicArrayList = NewBasicArrayListD(); _ = pathList
        pathList.Add(list)
        return pathList
}


// throws Exception
func (this *PathFinder) SearchStart(startPathFindingNode IPathFindingNode, endPathFindingNode IPathFindingNode, multipassState *MultipassState) {
        this.FindPathStart(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition, endPathFindingNode.PathFindingNode_().GeographicMapCellPosition, multipassState)
}


// throws Exception
func (this *PathFinder) SearchN(startPathFindingNode IPathFindingNode, endPathFindingNode IPathFindingNode, multipassState *MultipassState) IBasicArrayList {
        var list IBasicArrayList = this.FindPathEnd(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition, endPathFindingNode.PathFindingNode_().GeographicMapCellPosition, multipassState); _ = list
        if list == nil {
            return this.BasicArrayListUtil.GetImmutableInstance()
        }
        var pathList IBasicArrayList = NewBasicArrayListD(); _ = pathList
        pathList.Add(list)
        multipassState.Iteration = 0
        multipassState.Step = 0
        return pathList
}


// throws Exception
func (this *PathFinder) FindPath(start IGeographicMapCellPosition, target IGeographicMapCellPosition) IBasicArrayList {
        this.OpenPriorityQueue.Clear()
        this.ClosedSet.Clear()
        var discoveryCalculation int64; _ = discoveryCalculation
        var node *PathFindingNodeCost; _ = node
        var targetColumn int = target.GetColumn(); _ = targetColumn
        var targetRow int = target.GetRow(); _ = targetRow
        var sizeX int = len(this.CostArray); _ = sizeX
        var sizeY int = len(this.CostArray[0]); _ = sizeY
        for column := int(0); column < sizeX; column++ {
            for row := int(0); row < sizeY; row++ {
                discoveryCalculation = int64(this.MathUtil.Abs(column - targetColumn)) + int64(this.MathUtil.Abs(row - targetRow))
                node = this.CostArray[column][row]
                node.PathFindingNodeCostInfoP.TotalCostP = 0
                node.PathFindingNodeCostInfoP.CostToEndP = discoveryCalculation
            }
        }
        var startNode *PathFindingNodeCost = this.CostArray[start.GetColumn()][start.GetRow()]; _ = startNode
        this.OpenPriorityQueue.Add(startNode)
        var geographicMapInterface IBasicGeographicMap = any(this.GeographicMapInterface).(IBasicGeographicMap); _ = geographicMapInterface
        var basicGeographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = basicGeographicMapCellPositionFactory
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var targetNode *PathFindingNodeCost = this.CostArray[target.GetColumn()][target.GetRow()]; _ = targetNode
        var current *PathFindingNodeCost; _ = current
        for {
            current = this.OpenPriorityQueue.Poll()
            this.ClosedSet.Add(current)
            if Java_Equals(current, targetNode) {
                return this.ExtractPath(start, current)
            }
            var neighbor *PathFindingNodeCost; _ = neighbor
            var neighborInfo *PathFindingNodeCostInfo; _ = neighborInfo
            var calculatedCost int64; _ = calculatedCost
            for column := int(current.GeographicMapCellPosition.GetColumn() - 1); column < current.GeographicMapCellPosition.GetColumn() + 2; column++ {
                for row := int(current.GeographicMapCellPosition.GetRow() - 1); row < current.GeographicMapCellPosition.GetRow() + 2; row++ {
                    if column > 0 && row > 0 && column < allBinaryTiledLayer.GetColumns() && row < allBinaryTiledLayer.GetRows() && geographicMapInterface.IsOnMap(basicGeographicMapCellPositionFactory.GetAt(column, row)) {
                        neighbor = this.CostArray[column][row]
                        if this.ClosedSet.Contains(neighbor) {
                            continue
                        }
                        neighborInfo = neighbor.PathFindingNodeCostInfoP
                        calculatedCost = neighborInfo.CostToEndP + neighborInfo.CostFromStartP + current.PathFindingNodeCostInfoP.TotalCostP
                        if calculatedCost < neighborInfo.TotalCostP || !this.OpenPriorityQueue.Contains(neighbor) {
                            neighborInfo.TotalCostP = calculatedCost
                            neighbor.Parent = current
                            if !this.OpenPriorityQueue.Contains(neighbor) {
                                if geographicMapInterface.IsOfFourDirections(current.GeographicMapCellPosition, neighbor.GeographicMapCellPosition) {
                                    this.OpenPriorityQueue.Add(neighbor)
                                }
                            }
                        }
                    }
                }
            }
            if !(!this.OpenPriorityQueue.IsEmpty()) {
                break
            }
        }
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *PathFinder) FindPathStart(start IGeographicMapCellPosition, target IGeographicMapCellPosition, multipassState *MultipassState) {
        this.OpenPriorityQueue.Clear()
        this.ClosedSet.Clear()
        var discoveryCalculation int64; _ = discoveryCalculation
        var node *PathFindingNodeCost; _ = node
        var targetColumn int = target.GetColumn(); _ = targetColumn
        var targetRow int = target.GetRow(); _ = targetRow
        var sizeX int = len(this.CostArray); _ = sizeX
        var sizeY int = len(this.CostArray[0]); _ = sizeY
        for column := int(0); column < sizeX; column++ {
            for row := int(0); row < sizeY; row++ {
                discoveryCalculation = int64(this.MathUtil.Abs(column - targetColumn)) + int64(this.MathUtil.Abs(row - targetRow))
                node = this.CostArray[column][row]
                node.PathFindingNodeCostInfoP.TotalCostP = 0
                node.PathFindingNodeCostInfoP.CostToEndP = discoveryCalculation
            }
        }
        var startNode *PathFindingNodeCost = this.CostArray[start.GetColumn()][start.GetRow()]; _ = startNode
        this.OpenPriorityQueue.Add(startNode)
        multipassState.Step++
}


// throws Exception
func (this *PathFinder) FindPathEnd(start IGeographicMapCellPosition, target IGeographicMapCellPosition, multipassState *MultipassState) IBasicArrayList {
        var geographicMapInterface IBasicGeographicMap = any(this.GeographicMapInterface).(IBasicGeographicMap); _ = geographicMapInterface
        var basicGeographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = geographicMapInterface.GetGeographicMapCellPositionFactory(); _ = basicGeographicMapCellPositionFactory
        var allBinaryTiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = allBinaryTiledLayer
        var targetNode *PathFindingNodeCost = this.CostArray[target.GetColumn()][target.GetRow()]; _ = targetNode
        var current *PathFindingNodeCost; _ = current
        var total int = 0; _ = total
        for {
            current = this.OpenPriorityQueue.Poll()
            this.ClosedSet.Add(current)
            if Java_Equals(current, targetNode) {
                return this.ExtractPath(start, current)
            }
            var neighbor *PathFindingNodeCost; _ = neighbor
            var neighborInfo *PathFindingNodeCostInfo; _ = neighborInfo
            var calculatedCost int64; _ = calculatedCost
            for column := int(current.GeographicMapCellPosition.GetColumn() - 1); column < current.GeographicMapCellPosition.GetColumn() + 2; column++ {
                for row := int(current.GeographicMapCellPosition.GetRow() - 1); row < current.GeographicMapCellPosition.GetRow() + 2; row++ {
                    if column > 0 && row > 0 && column < allBinaryTiledLayer.GetColumns() && row < allBinaryTiledLayer.GetRows() && geographicMapInterface.IsOnMap(basicGeographicMapCellPositionFactory.GetAt(column, row)) {
                        neighbor = this.CostArray[column][row]
                        if this.ClosedSet.Contains(neighbor) {
                            continue
                        }
                        neighborInfo = neighbor.PathFindingNodeCostInfoP
                        calculatedCost = neighborInfo.CostToEndP + neighborInfo.CostFromStartP + current.PathFindingNodeCostInfoP.TotalCostP
                        if calculatedCost < neighborInfo.TotalCostP || !this.OpenPriorityQueue.Contains(neighbor) {
                            neighborInfo.TotalCostP = calculatedCost
                            neighbor.Parent = current
                            if !this.OpenPriorityQueue.Contains(neighbor) {
                                if geographicMapInterface.IsOfFourDirections(current.GeographicMapCellPosition, neighbor.GeographicMapCellPosition) {
                                    this.OpenPriorityQueue.Add(neighbor)
                                }
                            }
                        }
                    }
                }
            }
            total++
            if total > 10 {
                return this.BasicArrayListUtil.GetImmutableInstance()
            }
            if !(!this.OpenPriorityQueue.IsEmpty()) {
                break
            }
        }
        panic(errors.New(fmt.Sprint()))
}


func (this *PathFinder) ExtractPath(start IGeographicMapCellPosition, current *PathFindingNodeCost) IBasicArrayList {
        var path IBasicArrayList = NewBasicArrayListD(); _ = path
        for current.Parent != NullUtil_GetInstance().NULL_OBJECT {
            path.Add(current.GeographicMapCellPosition)
            current = any(current.Parent).(*PathFindingNodeCost)
        }
        path.Add(start)
        this.BasicArrayListUtil.Reverse(path)
        return path
}

