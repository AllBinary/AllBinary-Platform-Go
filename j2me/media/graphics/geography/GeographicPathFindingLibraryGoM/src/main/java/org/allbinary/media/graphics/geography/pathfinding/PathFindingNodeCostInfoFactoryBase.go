/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type PathFindingNodeCostInfoFactoryBase struct {
}

// Implements PathFindingNodeCostInfoFactoryBaseInterface

// The Java class PathFindingNodeCostInfoFactoryBase has subclasses so its variables are this interface
type IPathFindingNodeCostInfoFactoryBase interface {
    PathFindingNodeCostInfoFactoryBase_() *PathFindingNodeCostInfoFactoryBase
    Create(geographicMapInterface IBasicGeographicMap, comingFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition, costFromStart int64, costToEnd int64)
    GetTotalCost(geographicMapInterface IBasicGeographicMap, comingFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition) int64
}

func (this *PathFindingNodeCostInfoFactoryBase) PathFindingNodeCostInfoFactoryBase_() *PathFindingNodeCostInfoFactoryBase {
    return this
}

var PathFindingNodeCostInfoFactoryBase_NULL_PATH_FINDING_NODE_COST_INFO_FACTORY_BASE IPathFindingNodeCostInfoFactoryBase = NewPathFindingNodeCostInfoFactoryBase()

func NewPathFindingNodeCostInfoFactoryBase() *PathFindingNodeCostInfoFactoryBase {
    this := &PathFindingNodeCostInfoFactoryBase{}
    return this
}


// throws Exception
func (this *PathFindingNodeCostInfoFactoryBase) Create(geographicMapInterface IBasicGeographicMap, comingFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition, costFromStart int64, costToEnd int64) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *PathFindingNodeCostInfoFactoryBase) GetTotalCost(geographicMapInterface IBasicGeographicMap, comingFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition) int64 {
        panic(errors.New(fmt.Sprint()))
}

