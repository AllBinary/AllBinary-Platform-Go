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

type PathFindingNodeCostInfo struct {
    PathFindingNodeCostInfoData *PathFindingNodeCostInfoData
    CostFromStartP int64
    CostToEndP int64
    TotalCostP int64
}

// Implements Comparable[*PathFindingNodeCostInfo]


func (this *PathFindingNodeCostInfo) initFields() {
    this.PathFindingNodeCostInfoData = PathFindingNodeCostInfoData_GetInstance()
    
    
    
}


func NewPathFindingNodeCostInfo(costFromStart int64, costToGoal int64) *PathFindingNodeCostInfo {
    this := &PathFindingNodeCostInfo{}
    this.initFields()
        this.CostFromStartP = costFromStart
        this.CostToEndP = costToGoal
        this.SetCalculatedTotalCost()
    return this
}


func (this *PathFindingNodeCostInfo) SetCostToEnd(costToEnd int64) {
        this.CostToEndP = costToEnd
}


func (this *PathFindingNodeCostInfo) SetTotalCost(totalCost int64) {
        this.TotalCostP = totalCost
}


// throws Exception
func (this *PathFindingNodeCostInfo) SetCalculatedTotalCost() {
        this.TotalCostP = this.CostFromStartP + this.CostToEndP
        if this.TotalCostP > int64(this.PathFindingNodeCostInfoData.MAX_NODE_COST) {
            panic(errors.New(fmt.Sprint("Max Cost Exceeded")))
        }
}


func (this *PathFindingNodeCostInfo) SetCostFromStart(cost int64) {
        this.CostFromStartP = cost
}


func (this *PathFindingNodeCostInfo) AddCostFromStart(cost int64) {
        this.CostFromStartP += cost
}


func (this *PathFindingNodeCostInfo) CompareTo(pathFindingNodeCostInfo *PathFindingNodeCostInfo) int {
        return java_Lang.Long.Compare(this.TotalCostP, pathFindingNodeCostInfo.TotalCostP)
}


func (this *PathFindingNodeCostInfo) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(CommonLabels_GetInstance().COLON_SEP)
        stringBuffer.Append("CostFromStart: ")
        stringBuffer.Appendlong(this.CostFromStartP)
        stringBuffer.Append(" CostToEnd: ")
        stringBuffer.Appendlong(this.CostToEndP)
        stringBuffer.Append(" TotalCost: ")
        stringBuffer.Appendlong(this.TotalCostP)
        return fmt.Sprint(stringBuffer)
}

func (this *PathFindingNodeCostInfo) String() string {
    return this.ToString()
}

