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
type PathFinderGraphVisitorBase struct {
}

// The Java class PathFinderGraphVisitorBase has subclasses so its variables are this interface
type IPathFinderGraphVisitorBase interface {
    PathFinderGraphVisitorBase_() *PathFinderGraphVisitorBase
    Visit(graph any, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList)
    FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList)
    IsValid(graphPath any) bool
    GetInvalidReason(graphPath any) string
}

func (this *PathFinderGraphVisitorBase) PathFinderGraphVisitorBase_() *PathFinderGraphVisitorBase {
    return this
}


func NewPathFinderGraphVisitorBase() *PathFinderGraphVisitorBase {
    this := &PathFinderGraphVisitorBase{}
    return this
}


// throws Exception
func (this *PathFinderGraphVisitorBase) Visit(graph any, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *PathFinderGraphVisitorBase) FixPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, pathList IBasicArrayList) {
        panic(errors.New(fmt.Sprint()))
}


func (this *PathFinderGraphVisitorBase) IsValid(graphPath any) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *PathFinderGraphVisitorBase) GetInvalidReason(graphPath any) string {
        panic(errors.New(fmt.Sprint()))
}

