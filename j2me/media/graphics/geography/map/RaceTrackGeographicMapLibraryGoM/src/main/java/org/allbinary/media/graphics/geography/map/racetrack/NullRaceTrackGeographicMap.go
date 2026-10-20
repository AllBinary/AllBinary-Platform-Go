/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type NullRaceTrackGeographicMap struct {
}

// Implements RaceTrackGeographicMapInterface

var NullRaceTrackGeographicMap_NULL_RACE_TRACK_GEOGRAPHIC_MAP *NullRaceTrackGeographicMap = NewNullRaceTrackGeographicMap()

func NewNullRaceTrackGeographicMap() *NullRaceTrackGeographicMap {
    this := &NullRaceTrackGeographicMap{}
    return this
}


func (this *NullRaceTrackGeographicMap) GetId() int {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetName() string {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) Reset() {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetAllBinaryTiledLayer() IAllBinaryTiledLayer {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRaceTrackGeographicMap) GetCellPositionsAt(layer ILayer, currentCellPositionArray [][]IGeographicMapCellPosition, cellPositionArray [][]IGeographicMapCellPosition) bool {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRaceTrackGeographicMap) GetCellPositionAtXY(x int, y int) IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRaceTrackGeographicMap) GetCellPositionAtXYNoThrow(x int, y int) IGeographicMapCellPosition {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetCellTypeFromMapCellTypeInt(cellTypeId int) int {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRaceTrackGeographicMap) GetCellTypeAtXY(x int, y int) IGeographicMapCellType {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *NullRaceTrackGeographicMap) GetCellTypeAt(geographicMapCellPosition IGeographicMapCellPosition) IGeographicMapCellType {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetGeographicMapCellPositionFactory() IBasicGeographicMapCellPositionFactory {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetGeographicMapCellPositionFactoryInterface() GeographicMapCellPositionFactoryInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetForegroundBasicColor() *BasicColor {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetBackgroundBasicColor() *BasicColor {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetRaceTrackInfo() *RaceTrackInfo {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) SetRaceTrackInfo(aRaceTrackInfo *RaceTrackInfo) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) GetRaceTrackData() *RaceTrackData {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullRaceTrackGeographicMap) SetRaceTrackData(raceTrackData *RaceTrackData) {
        panic(errors.New(fmt.Sprint()))
}

