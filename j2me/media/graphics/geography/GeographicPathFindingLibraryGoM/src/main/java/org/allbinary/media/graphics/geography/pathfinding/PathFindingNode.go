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

type PathFindingNode struct {
    Parent any
    GeographicMapCellPosition IGeographicMapCellPosition
}

// The Java class PathFindingNode has subclasses so its variables are this interface
type IPathFindingNode interface {
    PathFindingNode_() *PathFindingNode
    GetParentP() IPathFindingNode
    SetGeographicMapCellPositionP(geographicMapCellPosition IGeographicMapCellPosition)
    ToString() string
}

func (this *PathFindingNode) PathFindingNode_() *PathFindingNode {
    return this
}


func (this *PathFindingNode) initFields() {
    this.Parent = NullUtil_GetInstance().NULL_OBJECT
    this.GeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}


func NewPathFindingNode(parent any, geographicMapCellPosition IGeographicMapCellPosition) *PathFindingNode {
    this := &PathFindingNode{}
    this.initFields()
        this.Parent = parent
        this.SetGeographicMapCellPositionP(geographicMapCellPosition)
        if this.GeographicMapCellPosition == SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION {
            panic(errors.New(fmt.Sprint("No GeographicMapCellPosition")))
        }
    return this
}


func (this *PathFindingNode) GetParentP() IPathFindingNode {
        return any(this.Parent).(IPathFindingNode)
}


func (this *PathFindingNode) SetGeographicMapCellPositionP(geographicMapCellPosition IGeographicMapCellPosition) {
        this.GeographicMapCellPosition = geographicMapCellPosition
}


func (this *PathFindingNode) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(CommonLabels_GetInstance().COLON_SEP)
        stringBuffer.Append(" Path: ")
        stringBuffer.Append(fmt.Sprint(this.GeographicMapCellPosition))
        var pathFindingNode IPathFindingNode = this.GetParentP(); _ = pathFindingNode
        for pathFindingNode != nil {
            stringBuffer.Append(fmt.Sprint(pathFindingNode.PathFindingNode_().GeographicMapCellPosition))
            stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            pathFindingNode = pathFindingNode.GetParentP()
        }
        return fmt.Sprint(stringBuffer)
}

func (this *PathFindingNode) String() string {
    return this.ToString()
}

