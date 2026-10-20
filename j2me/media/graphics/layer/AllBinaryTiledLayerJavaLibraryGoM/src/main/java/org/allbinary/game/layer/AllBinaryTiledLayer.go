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

type AllBinaryTiledLayer struct {
    Layer
    CellWidth int
    CellHeight int
    HalfWidth int
    HalfHeight int
    HalfCellWidth int
    HalfCellHeight int
    DataId int
}

// Implements NamedInterface, PaintableInterface

// The Java class AllBinaryTiledLayer has subclasses so its variables are this interface
type IAllBinaryTiledLayer interface {
    ILayer
    AllBinaryTiledLayer_() *AllBinaryTiledLayer
    SetCells(mapTwoDArray [][]int)
    UpdateCells(mapTwoDArray [][]int, fromTileId int, toTileId int)
    GetName() string
    Paint(graphics IGraphics)
    PaintThreed(graphics IGraphics)
    GetX2() int
    GetY2() int
    GetZ2() int
    GetCellWidth() int
    GetCellHeight() int
    GetCell(col int, row int) int
    GetColumns() int
    GetRows() int
    SetCell(col int, row int, tileIndex int)
    GetHalfWidth() int
    GetHalfHeight() int
    GetDataId() int
    GetHalfCellWidth() int
    GetHalfCellHeight() int
    IsOnTileLayer(column int, row int) bool
    GetAnimatedTile(animationTileIndex int) int
}

func (this *AllBinaryTiledLayer) AllBinaryTiledLayer_() *AllBinaryTiledLayer {
    return this
}


func (this *AllBinaryTiledLayer) initFields() {
    
    
    
    
    
    
    
}


func NewAllBinaryTiledLayer(dataId int, width int, height int, cellWidth int, cellHeight int) *AllBinaryTiledLayer {
    this := &AllBinaryTiledLayer{}
        this.Layer = *NewLayer(0, 0)
        this.initFields()
        this.DataId = dataId
        this.SetLayerWidth(width)
        this.SetLayerHeight(height)
        this.Layer.SetPosition(0, 0, 0)
        this.CellWidth = cellWidth
        this.CellHeight = cellHeight
        this.HalfWidth = (this.GetWidth() >> 1)
        this.HalfHeight = (this.GetHeight() >> 1)
        this.HalfCellWidth = (cellWidth >> 1)
        this.HalfCellHeight = (cellHeight >> 1)
    return this
}


func (this *AllBinaryTiledLayer) SetCells(mapTwoDArray [][]int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Start: r: ")
        stringBuffer.Appendint(len(mapTwoDArray))
        stringBuffer.Append(" c: ")
        stringBuffer.Appendint(len(mapTwoDArray[0]))
        stringBuffer.Append(" rows: ")
        stringBuffer.Appendint(this.GetRows())
        stringBuffer.Append(" columns: ")
        stringBuffer.Appendint(this.GetColumns())
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "setCells")
        var rows int = this.GetRows(); _ = rows
        var columns int = this.GetColumns(); _ = columns
        for col := int(0); col < columns; col++ {
            for row := int(0); row < rows; row++ {
                this.SetCell(col, row, mapTwoDArray[row][col])
            }
        }
}


func (this *AllBinaryTiledLayer) UpdateCells(mapTwoDArray [][]int, fromTileId int, toTileId int) {
        var rows int = this.GetRows(); _ = rows
        var columns int = this.GetColumns(); _ = columns
        for col := int(0); col < columns; col++ {
            for row := int(0); row < rows; row++ {
                if fromTileId == mapTwoDArray[row][col] {
                    mapTwoDArray[row][col] = toTileId
                    this.SetCell(col, row, mapTwoDArray[row][col])
                }
            }
        }
}


func (this *AllBinaryTiledLayer) GetName() string {
        return reflect.TypeOf(this).String()
}


func (this *AllBinaryTiledLayer) Paint(graphics IGraphics) {
}


func (this *AllBinaryTiledLayer) PaintThreed(graphics IGraphics) {
}


func (this *AllBinaryTiledLayer) GetX2() int {
        return this.X + this.GetWidth()
}


func (this *AllBinaryTiledLayer) GetY2() int {
        return this.Y + this.GetHeight()
}


func (this *AllBinaryTiledLayer) GetZ2() int {
        return int((int64(this.Z) + this.GetDepth()))
}


func (this *AllBinaryTiledLayer) GetCellWidth() int {
        return this.CellWidth
}


func (this *AllBinaryTiledLayer) GetCellHeight() int {
        return this.CellHeight
}


func (this *AllBinaryTiledLayer) GetCell(col int, row int) int {
        return -1
}


func (this *AllBinaryTiledLayer) GetColumns() int {
        return -1
}


func (this *AllBinaryTiledLayer) GetRows() int {
        return -1
}


func (this *AllBinaryTiledLayer) SetCell(col int, row int, tileIndex int) {
}


func (this *AllBinaryTiledLayer) GetHalfWidth() int {
        return this.HalfWidth
}


func (this *AllBinaryTiledLayer) GetHalfHeight() int {
        return this.HalfHeight
}


func (this *AllBinaryTiledLayer) GetDataId() int {
        return this.DataId
}


func (this *AllBinaryTiledLayer) GetHalfCellWidth() int {
        return this.HalfCellWidth
}


func (this *AllBinaryTiledLayer) GetHalfCellHeight() int {
        return this.HalfCellHeight
}


func (this *AllBinaryTiledLayer) IsOnTileLayer(column int, row int) bool {
        if this.GetColumns() > column && this.GetRows() > row && row >= 0 && column >= 0 {
            return true
        } else {
            return false
        }
}


func (this *AllBinaryTiledLayer) GetAnimatedTile(animationTileIndex int) int {
        return animationTileIndex
}

