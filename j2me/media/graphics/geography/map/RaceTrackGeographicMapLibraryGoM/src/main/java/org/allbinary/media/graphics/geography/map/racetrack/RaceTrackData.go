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

type RaceTrackData struct {
    TileSetImage IImage
    MiniTileSetImage IImage
    CellWidth int
    CellHeight int
    MiniCellWidth int
    MiniCellHeight int
    MapArray [][]int
    CellTypeIdToGeographicMapCellTypeArray []int
    NO_OVER_PASSES_ARRAY []ICellPosition
    OverPassGeographicMapCellPositionArray []ICellPosition
    Id int
}


func (this *RaceTrackData) initFields() {
    this.TileSetImage = NullImage_NULL_IMAGE
    this.MiniTileSetImage = NullImage_NULL_IMAGE
    this.CellWidth = 0
    this.CellHeight = 0
    this.MiniCellWidth = 0
    this.MiniCellHeight = 0
    this.MapArray = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY
    this.CellTypeIdToGeographicMapCellTypeArray = NullUtil_GetInstance().NULL_INT_ARRAY
    this.NO_OVER_PASSES_ARRAY = make([]ICellPosition, 0)
    this.OverPassGeographicMapCellPositionArray = this.NO_OVER_PASSES_ARRAY
    
}


func NewRaceTrackData(id int, cellWidth int, cellHeight int, miniCellWidth int, miniCellHeight int, cellTypeIdToGeographicMapCellTypeArray []int, mapArray [][]int) *RaceTrackData {
    this := &RaceTrackData{}
    this.initFields()
        this.Id = id
        this.CellWidth = cellWidth
        this.CellHeight = cellHeight
        this.MiniCellWidth = miniCellWidth
        this.MiniCellHeight = miniCellHeight
        this.MapArray = mapArray
        this.CellTypeIdToGeographicMapCellTypeArray = cellTypeIdToGeographicMapCellTypeArray
    return this
}


func (this *RaceTrackData) SetMapArray(mapArray [][]int) {
        this.MapArray = mapArray
}


func (this *RaceTrackData) GetMapArray() [][]int {
        return this.MapArray
}


func (this *RaceTrackData) SetCellTypeIdToGeographicMapCellTypeArray(cellTypeIdToGeographicMapCellTypeArray []int) {
        this.CellTypeIdToGeographicMapCellTypeArray = cellTypeIdToGeographicMapCellTypeArray
}


func (this *RaceTrackData) GetCellTypeIdToGeographicMapCellTypeArray() []int {
        return this.CellTypeIdToGeographicMapCellTypeArray
}


func (this *RaceTrackData) SetCellWidth(cellWidth int) {
        this.CellWidth = cellWidth
}


func (this *RaceTrackData) GetCellWidth() int {
        return this.CellWidth
}


func (this *RaceTrackData) SetCellHeight(cellHeight int) {
        this.CellHeight = cellHeight
}


func (this *RaceTrackData) GetCellHeight() int {
        return this.CellHeight
}


func (this *RaceTrackData) GetOverPassGeographicMapCellPositionArray() []ICellPosition {
        return this.OverPassGeographicMapCellPositionArray
}


func (this *RaceTrackData) SetOverPassGeographicMapCellPositionArray(overPassGeographicMapCellPositionArray []ICellPosition) {
        this.OverPassGeographicMapCellPositionArray = overPassGeographicMapCellPositionArray
}


func (this *RaceTrackData) GetTileSetImage() IImage {
        return this.TileSetImage
}


func (this *RaceTrackData) SetTileSetImage(tileSetImage IImage) {
        this.TileSetImage = tileSetImage
}


func (this *RaceTrackData) GetMiniTileSetImage() IImage {
        return this.MiniTileSetImage
}


func (this *RaceTrackData) SetMiniTileSetImage(miniTileSetImage IImage) {
        this.MiniTileSetImage = miniTileSetImage
}


func (this *RaceTrackData) GetMiniCellWidth() int {
        return this.MiniCellWidth
}


func (this *RaceTrackData) SetMiniCellWidth(miniCellWidth int) {
        this.MiniCellWidth = miniCellWidth
}


func (this *RaceTrackData) GetMiniCellHeight() int {
        return this.MiniCellHeight
}


func (this *RaceTrackData) SetMiniCellHeight(miniCellHeight int) {
        this.MiniCellHeight = miniCellHeight
}


func (this *RaceTrackData) SetId(id int) {
        this.Id = id
}


func (this *RaceTrackData) GetId() int {
        return this.Id
}

