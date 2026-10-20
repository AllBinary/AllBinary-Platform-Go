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

type PathFindingInfo struct {
    PathFinder IGeographicPathFinderBase
    StartPathFindingNodeList IBasicArrayList
    EndPathFindingNodeList IBasicArrayList
    PathFindingNodeCostInfoFactoryInterface IPathFindingNodeCostInfoFactoryBase
}

// Implements PathFindingInfoInterface

var PathFindingInfo_NULL_PATH_FINDING_INFO *PathFindingInfo = NewPathFindingInfo(PathFindingNodeCostInfoFactoryBase_NULL_PATH_FINDING_NODE_COST_INFO_FACTORY_BASE, BasicArrayListUtil_GetInstance().GetImmutableInstance(), BasicArrayListUtil_GetInstance().GetImmutableInstance())

func (this *PathFindingInfo) initFields() {
    this.PathFinder = GeographicPathFinderBase_NULL_GEOGRAPHIC_PATH_FINDER_BASE
    
    
    
}


func NewPathFindingInfo(pathFindingNodeCostInfoFactoryInterface IPathFindingNodeCostInfoFactoryBase, startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList) *PathFindingInfo {
    this := &PathFindingInfo{}
    this.initFields()
        this.PathFindingNodeCostInfoFactoryInterface = pathFindingNodeCostInfoFactoryInterface
        this.StartPathFindingNodeList = startPathFindingNodeList
        this.EndPathFindingNodeList = endPathFindingNodeList
    return this
}


func (this *PathFindingInfo) GetStartPathFindingNodeList() IBasicArrayList {
        return this.StartPathFindingNodeList
}


func (this *PathFindingInfo) GetEndPathFindingNodeList() IBasicArrayList {
        return this.EndPathFindingNodeList
}


func (this *PathFindingInfo) GetPathFinder() IGeographicPathFinderBase {
        return this.PathFinder
}


func (this *PathFindingInfo) AddStartPathFindingNode(startPathFindingNode IPathFindingNode) {
        this.StartPathFindingNodeList.Add(startPathFindingNode)
}


func (this *PathFindingInfo) AddEndPathFindingNode(endPathFindingNode IPathFindingNode) {
        this.EndPathFindingNodeList.Add(endPathFindingNode)
}


func (this *PathFindingInfo) SetPathFinder(pathFinder IGeographicPathFinderBase) {
        this.PathFinder = pathFinder
}


func (this *PathFindingInfo) GetPathFindingNodeCostInfoFactoryInterface() IPathFindingNodeCostInfoFactoryBase {
        return this.PathFindingNodeCostInfoFactoryInterface
}

