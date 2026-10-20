/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2006 AllBinary
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

type PathFindingNodeCostInfoFactory struct {
    PathFindingNodeCostInfoFactoryBase
    PathFindingNodeCostInfoAdjacencyList [][]*PathFindingNodeCostInfo
}


func (this *PathFindingNodeCostInfoFactory) initFields() {
    
}


func NewPathFindingNodeCostInfoFactory(max_ int) *PathFindingNodeCostInfoFactory {
    this := &PathFindingNodeCostInfoFactory{}
    this.PathFindingNodeCostInfoFactoryBase = *NewPathFindingNodeCostInfoFactoryBase()
    this.initFields()
        this.PathFindingNodeCostInfoAdjacencyList = func() [][]*PathFindingNodeCostInfo { array := make([][]*PathFindingNodeCostInfo, max_); for index := range array { array[index] = make([]*PathFindingNodeCostInfo, max_) }; return array }()
    return this
}


// throws Exception
func (this *PathFindingNodeCostInfoFactory) Create(geographicMapInterface IBasicGeographicMap, goingToGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition, costFromStart int64, costToEnd int64) {
        this.GetInstance(goingToGeographicMapCellPosition, geographicMapCellPosition, costFromStart, costToEnd)
}


// throws Exception
func (this *PathFindingNodeCostInfoFactory) GetInstance_1(goingToGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition, costFromStart int64, costToEnd int64) *PathFindingNodeCostInfo {
        var pathFindingNodeCostInfo *PathFindingNodeCostInfo = this.GetInstance(goingToGeographicMapCellPosition, geographicMapCellPosition); _ = pathFindingNodeCostInfo
        if pathFindingNodeCostInfo == nil {
            pathFindingNodeCostInfo = NewPathFindingNodeCostInfo(costFromStart, costToEnd)
            pathFindingNodeCostInfo.SetCalculatedTotalCost()
            this.PathFindingNodeCostInfoAdjacencyList[geographicMapCellPosition.GetId()][goingToGeographicMapCellPosition.GetId()] = pathFindingNodeCostInfo
        } else {
            pathFindingNodeCostInfo.SetCostFromStart(costFromStart)
            pathFindingNodeCostInfo.SetCostToEnd(costToEnd)
            pathFindingNodeCostInfo.SetCalculatedTotalCost()
        }
        return pathFindingNodeCostInfo
}

// Java overloads of PathFindingNodeCostInfoFactory.getInstance
func (this *PathFindingNodeCostInfoFactory) GetInstance(args ...any) *PathFindingNodeCostInfo {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[IGeographicMapCellPosition](args[0], pass) && Java_IsArg[IGeographicMapCellPosition](args[1], pass) && Java_IsArg[int64](args[2], pass) && Java_IsArg[int64](args[3], pass) {
            return this.GetInstance_1(Java_Arg[IGeographicMapCellPosition](args[0]), Java_Arg[IGeographicMapCellPosition](args[1]), Java_Arg[int64](args[2]), Java_Arg[int64](args[3]))
        }
        if len(args) == 2 && Java_IsArg[IGeographicMapCellPosition](args[0], pass) && Java_IsArg[IGeographicMapCellPosition](args[1], pass) {
            return this.GetInstance_2(Java_Arg[IGeographicMapCellPosition](args[0]), Java_Arg[IGeographicMapCellPosition](args[1]))
        }
    }
    panic("No Java overload of PathFindingNodeCostInfoFactory.getInstance for the arguments")
}


func (this *PathFindingNodeCostInfoFactory) GetInstance_2(goingToFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition) *PathFindingNodeCostInfo {
        return this.PathFindingNodeCostInfoAdjacencyList[geographicMapCellPosition.GetId()][goingToFromGeographicMapCellPosition.GetId()]
}


func (this *PathFindingNodeCostInfoFactory) GetTotalCost(geographicMapInterface IBasicGeographicMap, comingFromGeographicMapCellPosition IGeographicMapCellPosition, geographicMapCellPosition IGeographicMapCellPosition) int64 {
        return this.GetInstance(comingFromGeographicMapCellPosition, geographicMapCellPosition).TotalCostP
}

