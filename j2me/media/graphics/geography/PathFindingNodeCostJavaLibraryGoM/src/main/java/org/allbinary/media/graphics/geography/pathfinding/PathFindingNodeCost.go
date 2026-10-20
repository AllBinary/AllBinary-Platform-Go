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

type PathFindingNodeCost struct {
    PathFindingNode
    CommonSeps *CommonSeps
    PathFindingNodeCostInfoP *PathFindingNodeCostInfo
}

// Implements Comparable[*PathFindingNodeCost]

var PathFindingNodeCost_NULL_PATH_FINDING_NODE_COST_ARRAY_ARRAY [][]*PathFindingNodeCost = func() [][]*PathFindingNodeCost { array := make([][]*PathFindingNodeCost, 0); for index := range array { array[index] = make([]*PathFindingNodeCost, 0) }; return array }()

func (this *PathFindingNodeCost) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
    
}


func NewPathFindingNodeCost(parent any, geographicMapCellPosition IGeographicMapCellPosition, pathFindingNodeCostInfo *PathFindingNodeCostInfo) *PathFindingNodeCost {
    this := &PathFindingNodeCost{}
        this.PathFindingNode = *NewPathFindingNode(parent, geographicMapCellPosition)
        this.initFields()
        this.PathFindingNodeCostInfoP = pathFindingNodeCostInfo
        if this.GeographicMapCellPosition == nil {
            panic(errors.New(fmt.Sprint("No GeographicMapCellPosition")))
        }
        if this.GetPathFindingNodeCostInfo() == nil {
            panic(errors.New(fmt.Sprint("No PathFindingNodeCostInfo")))
        }
    return this
}


func (this *PathFindingNodeCost) GetPathFindingNodeCostInfo() *PathFindingNodeCostInfo {
        return this.PathFindingNodeCostInfoP
}


func (this *PathFindingNodeCost) SetPathFindingNodeCostInfo(pathFindingNodeCostInfo *PathFindingNodeCostInfo) {
        this.PathFindingNodeCostInfoP = pathFindingNodeCostInfo
}


func (this *PathFindingNodeCost) CompareTo(pathFindingNodeCost *PathFindingNodeCost) int {
        return this.PathFindingNodeCostInfoP.CompareTo(pathFindingNodeCost.PathFindingNodeCostInfoP)
}


func (this *PathFindingNodeCost) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(CommonLabels_GetInstance().COLON_SEP)
        stringBuffer.Append(fmt.Sprint(this.GetPathFindingNodeCostInfo()))
        stringBuffer.Append(" Path: ")
        stringBuffer.Append(fmt.Sprint(this.GeographicMapCellPosition))
        var pathFindingNode IPathFindingNode = any(this.Parent).(IPathFindingNode); _ = pathFindingNode
        for pathFindingNode != nil {
            stringBuffer.Append(fmt.Sprint(pathFindingNode.PathFindingNode_().GeographicMapCellPosition))
            stringBuffer.Append(this.CommonSeps.SPACE)
            pathFindingNode = any(pathFindingNode.PathFindingNode_().Parent).(IPathFindingNode)
        }
        return fmt.Sprint(stringBuffer)
}

func (this *PathFindingNodeCost) String() string {
    return this.ToString()
}

