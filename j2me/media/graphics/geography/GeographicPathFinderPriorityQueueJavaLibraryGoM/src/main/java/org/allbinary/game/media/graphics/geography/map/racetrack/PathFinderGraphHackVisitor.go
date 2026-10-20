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
type PathFinderGraphHackVisitor struct {
    BasePathFinderGraphVisitor
}



func NewPathFinderGraphHackVisitor(geographicMapInterface IBasicGeographicMap, edgeMinimum int, minPathWeight int, maxPathWeight int) *PathFinderGraphHackVisitor {
    this := &PathFinderGraphHackVisitor{}
        this.BasePathFinderGraphVisitor = *NewBasePathFinderGraphVisitor(geographicMapInterface, edgeMinimum, minPathWeight, maxPathWeight)
    return this
}


// throws Exception
func (this *PathFinderGraphHackVisitor) Visit(graph any, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList) {
        this.FixStart(graph, startPathFindingNodeList)
        this.FixEnd(graph, endPathFindingNodeList)
        this.FixOverPassEdges(graph)
}


// throws Exception
func (this *PathFinderGraphHackVisitor) FixStart(graph any, startPathFindingNodeList IBasicArrayList) {
}


// throws Exception
func (this *PathFinderGraphHackVisitor) FixEnd(graph any, endPathFindingNodeList IBasicArrayList) {
}


// throws Exception
func (this *PathFinderGraphHackVisitor) FixOverPassEdges(graph any) {
}


// throws Exception
func (this *PathFinderGraphHackVisitor) FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList) {
}


// throws Exception
func (this *PathFinderGraphHackVisitor) RemoveOverPassEdges(pathList IBasicArrayList) {
}


func (this *PathFinderGraphHackVisitor) IsValid(graphPath any) bool {
        return false
}

