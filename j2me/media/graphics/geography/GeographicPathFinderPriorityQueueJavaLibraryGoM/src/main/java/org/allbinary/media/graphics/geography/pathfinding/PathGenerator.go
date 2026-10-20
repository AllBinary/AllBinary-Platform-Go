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

type PathGenerator struct {
}

// Implements PathGeneratorInterface

var PathGenerator_SINGLETON *PathGenerator = NewPathGenerator()


func PathGenerator_GetInstance() *PathGenerator {
        return PathGenerator_SINGLETON
}


func NewPathGenerator() *PathGenerator {
    this := &PathGenerator{}
    return this
}


// throws Exception
func (this *PathGenerator) Init(geographicMapInterface any, totalPaths int) {
}


// throws Exception
func (this *PathGenerator) GetInstanceNoCache(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, totalPaths int) IBasicArrayList {
        var geographicMapCellPositionBasicArrayList IBasicArrayList = this.Create(geographicMapInterface, pathFindingInfo, totalPaths); _ = geographicMapCellPositionBasicArrayList
        return geographicMapCellPositionBasicArrayList
}


// throws Exception
func (this *PathGenerator) CreateList(geographicMapInterface IBasicGeographicMap, geographicMapCellHistory *GeographicMapCellHistory, pathFindingInfo *PathFindingInfo, totalPaths int) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *PathGenerator) Create(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, totalPaths int) IBasicArrayList {
        var startPathFindingNodeList IBasicArrayList = pathFindingInfo.GetStartPathFindingNodeList(); _ = startPathFindingNodeList
        var endPathFindingNodeList IBasicArrayList = pathFindingInfo.GetEndPathFindingNodeList(); _ = endPathFindingNodeList
        var geographicPathFinderInterface IGeographicPathFinderBase = pathFindingInfo.GetPathFinder(); _ = geographicPathFinderInterface
        var geographicMapCellPositionBasicArrayList IBasicArrayList = geographicPathFinderInterface.SearchTotalPath(startPathFindingNodeList, endPathFindingNodeList, totalPaths); _ = geographicMapCellPositionBasicArrayList
        return geographicMapCellPositionBasicArrayList
}


// throws Exception
func (this *PathGenerator) CreateN(geographicMapInterface IBasicGeographicMap, pathFindingInfo *PathFindingInfo, totalPaths int, multipassState *MultipassState) IBasicArrayList {
        var startPathFindingNodeList IBasicArrayList = pathFindingInfo.GetStartPathFindingNodeList(); _ = startPathFindingNodeList
        var endPathFindingNodeList IBasicArrayList = pathFindingInfo.GetEndPathFindingNodeList(); _ = endPathFindingNodeList
        var geographicPathFinderInterface IGeographicPathFinderBase = pathFindingInfo.GetPathFinder(); _ = geographicPathFinderInterface
        var geographicMapCellPositionBasicArrayList IBasicArrayList = geographicPathFinderInterface.SearchTotalPathN(startPathFindingNodeList, endPathFindingNodeList, totalPaths, multipassState); _ = geographicMapCellPositionBasicArrayList
        return geographicMapCellPositionBasicArrayList
}


// throws Exception
func (this *PathGenerator) IsValid(pathList IBasicArrayList, geographicMapCellHistory *GeographicMapCellHistory) bool {
        panic(errors.New(fmt.Sprint()))
}

