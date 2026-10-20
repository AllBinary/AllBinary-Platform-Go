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
type BasePathFinderGraphVisitor struct {
    PathFinderGraphVisitorBase
    GeographicMapInterface IBasicGeographicMap
    EdgeMinimum int
    MinPathWeight int
    MaxPathWeight int
}

// The Java class BasePathFinderGraphVisitor has subclasses so its variables are this interface
type IBasePathFinderGraphVisitor interface {
    IPathFinderGraphVisitorBase
    BasePathFinderGraphVisitor_() *BasePathFinderGraphVisitor
    Visit(graph any, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList)
    FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList)
    IsValid(graphPath any) bool
    GetInvalidReason(graphPath any) string
}

func (this *BasePathFinderGraphVisitor) BasePathFinderGraphVisitor_() *BasePathFinderGraphVisitor {
    return this
}


func (this *BasePathFinderGraphVisitor) initFields() {
    
    
    
    
}


func NewBasePathFinderGraphVisitor(geographicMapInterface IBasicGeographicMap, edgeMinimum int, minPathWeight int, maxPathWeight int) *BasePathFinderGraphVisitor {
    this := &BasePathFinderGraphVisitor{}
    this.PathFinderGraphVisitorBase = *NewPathFinderGraphVisitorBase()
    this.initFields()
        this.GeographicMapInterface = geographicMapInterface
        this.EdgeMinimum = edgeMinimum
        this.MinPathWeight = minPathWeight
        this.MaxPathWeight = maxPathWeight
    return this
}


// throws Exception
func (this *BasePathFinderGraphVisitor) Visit(graph any, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList) {
}


// throws Exception
func (this *BasePathFinderGraphVisitor) FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList) {
        var startPathFindingNode IPathFindingNode; _ = startPathFindingNode
        for index := int(startPathFindingNodeList.Size() - 1); index >= 0; index-- {
            startPathFindingNode = any(startPathFindingNodeList.Get(index)).(IPathFindingNode)
            pathList.Remove(startPathFindingNode.PathFindingNode_().GeographicMapCellPosition)
            pathList.AddAt(0, startPathFindingNode.PathFindingNode_().GeographicMapCellPosition)
        }
}


func (this *BasePathFinderGraphVisitor) IsValid(graphPath any) bool {
        return false
}


func (this *BasePathFinderGraphVisitor) GetInvalidReason(graphPath any) string {
        return StringUtil_GetInstance().EMPTY_STRING
}

