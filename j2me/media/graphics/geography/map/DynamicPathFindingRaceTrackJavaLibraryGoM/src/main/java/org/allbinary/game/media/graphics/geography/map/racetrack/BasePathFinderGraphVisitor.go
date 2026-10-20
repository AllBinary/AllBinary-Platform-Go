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

type BasePathFinderGraphVisitor[V any, E any] struct {
    PathFinderGraphVisitorBase
    GeographicMapInterface IBasicGeographicMap
    EdgeMinimum int
    MinPathWeight int
    MaxPathWeight int
}

// The Java class BasePathFinderGraphVisitor has subclasses so its variables are this interface
type IBasePathFinderGraphVisitor[V any, E any] interface {
    IPathFinderGraphVisitorBase
    BasePathFinderGraphVisitor_() *BasePathFinderGraphVisitor[V, E]
    Visit(graph *SimpleWeightedGraph, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList)
    FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList)
    IsValid(graphPath IGraphPath) bool
    GetInvalidReason(graphPath IGraphPath) string
}

func (this *BasePathFinderGraphVisitor[V, E]) BasePathFinderGraphVisitor_() *BasePathFinderGraphVisitor[V, E] {
    return this
}


func (this *BasePathFinderGraphVisitor[V, E]) initFields() {
    
    
    
    
}


func NewBasePathFinderGraphVisitor[V any, E any](geographicMapInterface IBasicGeographicMap, edgeMinimum int, minPathWeight int, maxPathWeight int) *BasePathFinderGraphVisitor[V, E] {
    this := &BasePathFinderGraphVisitor[V, E]{}
    this.PathFinderGraphVisitorBase = *NewPathFinderGraphVisitorBase()
    this.initFields()
        this.GeographicMapInterface = geographicMapInterface
        this.EdgeMinimum = edgeMinimum
        this.MinPathWeight = minPathWeight
        this.MaxPathWeight = maxPathWeight
    return this
}


// throws Exception
func (this *BasePathFinderGraphVisitor[V, E]) Visit(graph *SimpleWeightedGraph, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList) {
}


// throws Exception
func (this *BasePathFinderGraphVisitor[V, E]) FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList) {
        for index := int(startPathFindingNodeList.Size() - 1); index >= 0; index-- {
            var startPathFindingNode IPathFindingNode = any(startPathFindingNodeList.Get(index)).(IPathFindingNode); _ = startPathFindingNode
            pathList.Remove(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition)
            pathList.AddAt(0, startPathFindingNode.PathFindingNode_().GeographicMapCellPosition)
        }
}


func (this *BasePathFinderGraphVisitor[V, E]) IsValid(graphPath IGraphPath) bool {
        if graphPath.GetEdgeList().Size() > this.EdgeMinimum && graphPath.GetWeight() < int64(this.MaxPathWeight) && graphPath.GetWeight() > int64(this.MinPathWeight) {
            return true
        } else {
            return false
        }
}


func (this *BasePathFinderGraphVisitor[V, E]) GetInvalidReason(graphPath IGraphPath) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if graphPath.GetEdgeList().Size() < this.EdgeMinimum {
            stringBuffer.Append("Ignoring Small: " + fmt.Sprint(graphPath.GetEdgeList().Size()))
        }
        if graphPath.GetWeight() <= int64(this.MinPathWeight) {
            stringBuffer.Append(" Low Weighted Path: " + fmt.Sprint(graphPath.GetWeight()))
        }
        if graphPath.GetWeight() >= int64(this.MaxPathWeight) {
            stringBuffer.Append(" To High Weighted Path: " + fmt.Sprint(graphPath.GetWeight()))
        }
        return fmt.Sprint(stringBuffer)
}

