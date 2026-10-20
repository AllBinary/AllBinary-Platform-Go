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

type GeographicMapCellPathHistoryInfo struct {
    PreviousGeographicMapCellPosition IGeographicMapCellPosition
    PreviousOnPathGeographicMapCellPosition IGeographicMapCellPosition
    NextUnvisitedOnPathGeographicMapCellPosition IGeographicMapCellPosition
    NextOnPathGeographicMapCellPosition IGeographicMapCellPosition
    NextChosenOnPathGeographicMapCellPosition IGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) initFields() {
    this.PreviousGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.PreviousOnPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.NextUnvisitedOnPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.NextOnPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
    this.NextChosenOnPathGeographicMapCellPosition = SimpleGeographicMapCellPositionFactory_NULL_GEOGRAPHIC_MAP_CELL_POSITION
}

func NewGeographicMapCellPathHistoryInfo() *GeographicMapCellPathHistoryInfo {
    this := &GeographicMapCellPathHistoryInfo{}
    this.initFields()
    return this
}


func (this *GeographicMapCellPathHistoryInfo) GetPreviousGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.PreviousGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) SetPreviousGeographicMapCellPosition(previousGeographicMapCellPosition IGeographicMapCellPosition) {
        this.PreviousGeographicMapCellPosition = previousGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) GetPreviousOnPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.PreviousOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) SetPreviousOnPathGeographicMapCellPosition(previousOnPathGeographicMapCellPosition IGeographicMapCellPosition) {
        this.PreviousOnPathGeographicMapCellPosition = previousOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) GetNextUnvisitedOnPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.NextUnvisitedOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) SetNextUnvisitedOnPathGeographicMapCellPosition(nextUnvisitedOnPathGeographicMapCellPosition IGeographicMapCellPosition) {
        this.NextUnvisitedOnPathGeographicMapCellPosition = nextUnvisitedOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) GetNextOnPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.NextOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) SetNextOnPathGeographicMapCellPosition(nextOnPathGeographicMapCellPosition IGeographicMapCellPosition) {
        this.NextOnPathGeographicMapCellPosition = nextOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) GetNextChosenOnPathGeographicMapCellPosition() IGeographicMapCellPosition {
        return this.NextChosenOnPathGeographicMapCellPosition
}


func (this *GeographicMapCellPathHistoryInfo) SetNextChosenOnPathGeographicMapCellPosition(nextChosenOnPathGeographicMapCellPosition IGeographicMapCellPosition) {
        this.NextChosenOnPathGeographicMapCellPosition = nextChosenOnPathGeographicMapCellPosition
}

