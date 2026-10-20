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

type PathFinderGraphHackVisitor[V any, E any] struct {
    BasePathFinderGraphVisitor
    LogUtil *LogUtil
}


func (this *PathFinderGraphHackVisitor[V, E]) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPathFinderGraphHackVisitor[V any, E any](geographicMapInterface IBasicGeographicMap, edgeMinimum int, minPathWeight int, maxPathWeight int) *PathFinderGraphHackVisitor[V, E] {
    this := &PathFinderGraphHackVisitor[V, E]{}
        this.BasePathFinderGraphVisitor = *NewBasePathFinderGraphVisitor(geographicMapInterface, edgeMinimum, minPathWeight, maxPathWeight)
        this.initFields()
    return this
}


// throws Exception
func (this *PathFinderGraphHackVisitor[V, E]) Visit(graph *SimpleWeightedGraph, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList) {
        this.FixStart(graph, startPathFindingNodeList)
        this.FixEnd(graph, endPathFindingNodeList)
        this.FixOverPassEdges(graph)
}


// throws Exception
func (this *PathFinderGraphHackVisitor[V, E]) FixStart(graph *SimpleWeightedGraph[ICellPosition, *DefaultWeightedEdge], startPathFindingNodeList IBasicArrayList) {
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = this.GeographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var tiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        for index := int(startPathFindingNodeList.Size() - 1); index >= 0; index-- {
            var startPathFindingNode IPathFindingNode = any(startPathFindingNodeList.Get(index)).(IPathFindingNode); _ = startPathFindingNode
            var geographicMapCellPosition IGeographicMapCellPosition = startPathFindingNode.PathFindingNode_().GeographicMapCellPosition; _ = geographicMapCellPosition
            var column int = geographicMapCellPosition.GetColumn(); _ = column
            var row int = geographicMapCellPosition.GetRow(); _ = row
            var nextRow int = row + 1; _ = nextRow
            if tiledLayer.IsOnTileLayer(column, nextRow) {
                var geographicMapCellPositionNeighbor IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(column, nextRow); _ = geographicMapCellPositionNeighbor
                graph.RemoveEdge(geographicMapCellPosition, geographicMapCellPositionNeighbor)
            }
        }
}


// throws Exception
func (this *PathFinderGraphHackVisitor[V, E]) FixEnd(graph *SimpleWeightedGraph[ICellPosition, *DefaultWeightedEdge], endPathFindingNodeList IBasicArrayList) {
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = this.GeographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var tiledLayer IAllBinaryTiledLayer = this.GeographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        for index := int(endPathFindingNodeList.Size() - 1); index >= 0; index-- {
            var endPathFindingNode IPathFindingNode = any(endPathFindingNodeList.Get(index)).(IPathFindingNode); _ = endPathFindingNode
            var geographicMapCellPosition IGeographicMapCellPosition = endPathFindingNode.PathFindingNode_().GeographicMapCellPosition; _ = geographicMapCellPosition
            graph.AddVertex(geographicMapCellPosition)
            var column int = geographicMapCellPosition.GetColumn(); _ = column
            var row int = geographicMapCellPosition.GetRow(); _ = row
            var nextRow int = row + 1; _ = nextRow
            if tiledLayer.IsOnTileLayer(column, nextRow) {
                var geographicMapCellPositionNeighbor IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(column, nextRow); _ = geographicMapCellPositionNeighbor
                graph.AddEdge(geographicMapCellPosition, geographicMapCellPositionNeighbor)
            }
        }
}


// throws Exception
func (this *PathFinderGraphHackVisitor[V, E]) FixOverPassEdges(graph *SimpleWeightedGraph[ICellPosition, *DefaultWeightedEdge]) {
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = this.GeographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var raceTrackGeographicMap *RaceTrackGeographicMap = any(this.GeographicMapInterface).(*RaceTrackGeographicMap); _ = raceTrackGeographicMap
        var CellPositionArray []ICellPosition = raceTrackGeographicMap.GetRaceTrackData().GetOverPassGeographicMapCellPositionArray(); _ = CellPositionArray
        for index := int(len(CellPositionArray) - 1); index >= 0; index-- {
            var overPassGeographicMapCellPosition ICellPosition = CellPositionArray[index]; _ = overPassGeographicMapCellPosition
            this.LogUtil.PutF("Fixing Over Pass: " + fmt.Sprint(overPassGeographicMapCellPosition.GetColumn()) + ", " + fmt.Sprint(overPassGeographicMapCellPosition.GetRow()), this, "fixOverPassEdges")
            var underPassGeographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(overPassGeographicMapCellPosition.GetColumn(), overPassGeographicMapCellPosition.GetRow()); _ = underPassGeographicMapCellPosition
            var rightUnderPassGeographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(overPassGeographicMapCellPosition.GetColumn() + 1, overPassGeographicMapCellPosition.GetRow()); _ = rightUnderPassGeographicMapCellPosition
            var leftUnderPassGeographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(overPassGeographicMapCellPosition.GetColumn() - 1, overPassGeographicMapCellPosition.GetRow()); _ = leftUnderPassGeographicMapCellPosition
            graph.RemoveEdge(underPassGeographicMapCellPosition, rightUnderPassGeographicMapCellPosition)
            graph.RemoveEdge(underPassGeographicMapCellPosition, leftUnderPassGeographicMapCellPosition)
            graph.AddVertex(overPassGeographicMapCellPosition)
            graph.AddEdge(overPassGeographicMapCellPosition, rightUnderPassGeographicMapCellPosition)
            graph.AddEdge(overPassGeographicMapCellPosition, leftUnderPassGeographicMapCellPosition)
        }
}


// throws Exception
func (this *PathFinderGraphHackVisitor[V, E]) FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList) {
        for index := int(startPathFindingNodeList.Size() - 1); index >= 0; index-- {
            var endPathFindingNode IPathFindingNode = any(endPathFindingNodeList.Get(index)).(IPathFindingNode); _ = endPathFindingNode
            var startPathFindingNode IPathFindingNode = any(startPathFindingNodeList.Get(index)).(IPathFindingNode); _ = startPathFindingNode
            if BasicGeographicMapUtil_GetInstance().IsSameCellPosition(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition, endPathFindingNode.PathFindingNode_().GeographicMapCellPosition) {
                if pathList.Remove(endPathFindingNode.PathFindingNode_().GeographicMapCellPosition) {
                    pathList.Remove(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition)
                    pathList.AddAt(0, startPathFindingNode.PathFindingNode_().GeographicMapCellPosition)
                }
            }
        }
        this.RemoveOverPassEdges(pathList)
}


// throws Exception
func (this *PathFinderGraphHackVisitor[V, E]) RemoveOverPassEdges(pathList IBasicArrayList) {
        var geographicMapCellPositionFactory IBasicGeographicMapCellPositionFactory = this.GeographicMapInterface.GetGeographicMapCellPositionFactory(); _ = geographicMapCellPositionFactory
        var raceTrackGeographicMap *RaceTrackGeographicMap = any(this.GeographicMapInterface).(*RaceTrackGeographicMap); _ = raceTrackGeographicMap
        var CellPositionArray []ICellPosition = raceTrackGeographicMap.GetRaceTrackData().GetOverPassGeographicMapCellPositionArray(); _ = CellPositionArray
        for index := int(len(CellPositionArray) - 1); index >= 0; index-- {
            var overPassGeographicMapCellPosition ICellPosition = CellPositionArray[index]; _ = overPassGeographicMapCellPosition
            var underPassGeographicMapCellPosition IGeographicMapCellPosition = geographicMapCellPositionFactory.GetAt(overPassGeographicMapCellPosition.GetColumn(), overPassGeographicMapCellPosition.GetRow()); _ = underPassGeographicMapCellPosition
            var indexOf int = pathList.IndexOf(overPassGeographicMapCellPosition); _ = indexOf
            if indexOf != -1 {
                pathList.Set(indexOf, underPassGeographicMapCellPosition)
            }
        }
}


func (this *PathFinderGraphHackVisitor[V, E]) IsValid(graphPath IGraphPath) bool {
        if graphPath.GetEdgeList().Size() > this.EdgeMinimum && graphPath.GetWeight() < this.MaxPathWeight && graphPath.GetWeight() > this.MinPathWeight {
            return true
        } else {
            return false
        }
}

