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

type TrueTypeFontUtilBase struct {
    ShortPattern string
    PatternAsBytes []int8
    Pattern string
    CharArray []int16
    _characterWidth []int
    CharacterArray []rune
    Size int
    LastCapIndex int
    CurrentFont Font
    Scale int
    CELLS_PER_ROW int
    FontSize int
    BaseCharWidth int
    CellSize int
    TextureSize int
    ActualCellsPerRow int
}

// The Java class TrueTypeFontUtilBase has subclasses so its variables are this interface
type ITrueTypeFontUtilBase interface {
    TrueTypeFontUtilBase_() *TrueTypeFontUtilBase
    GetAsTextureSize(textureSize int) int
    GetCellSize(cellSize int) int
    GetYOffset() int
}

func (this *TrueTypeFontUtilBase) TrueTypeFontUtilBase_() *TrueTypeFontUtilBase {
    return this
}


func (this *TrueTypeFontUtilBase) initFields() {
    
    this.PatternAsBytes = []int8{32, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 32, 32, 32, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 32, 32, 32, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 32, 32, 32, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 32, 32, 32, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 46, 63, 32, 32, 32, 33, 36, 37, 96, int8(-62), int8(-84), int8(-62), int8(-84), 34, int8(-62), int8(-93), int8(-62), int8(-93), 94, 38, 42, 40, 32, 32, 32, 41, 95, 43, 45, 61, 91, 93, 123, 125, 59, 39, 35, 58, 32, 32, 32, 64, 126, 44, 47, 60, 62, 92, 124, int8(-62), int8(-82), int8(-62), int8(-82), int8(-62), int8(-87), int8(-62), int8(-87)}
    this.Pattern = string(this.PatternAsBytes)
    this.CharArray = []int16{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6f, 0x50, 0x56, 0x6b, 0x51, 0x52, 0x5a, 0x6a, 0x5c, 0x60, 0x5b, 0x62, 0x72, 0x63, 0x4b, 0x73, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0x6c, 0x69, 0x74, 0x64, 0x75, 0x4c, 0x70, 0xb, 0xc, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x65, 0x76, 0x66, 0x59, 0x61, 0x53, 0x2b, 0x2c, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x67, 0x77, 0x68, 0x71, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x58, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7b, 0x0, 0x0, 0x55, 0x0, 0x79, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
    this._characterWidth = make([]int, len(this.Pattern))
    this.CharacterArray = make([]rune, 1)
    this.Size = len(this.Pattern)
    this.LastCapIndex = strings.Index(this.Pattern, string('Z'))
    this.CurrentFont = Font_GetDefaultFont()
    
    
    
    
    
    
    
}


func NewTrueTypeFontUtilBase(scale int) *TrueTypeFontUtilBase {
    this := &TrueTypeFontUtilBase{}
    this.initFields()
        var shortPattern string = StringUtil_GetInstance().EMPTY_STRING; _ = shortPattern
        if AvianUtil_IsAvian() {
            var shortPatterAsByteArray []int8 = []int8{32, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 46, 63, 33, 36, 37, 96, int8(-62), int8(-84), int8(-62), int8(-84), 34, int8(-62), int8(-93), int8(-62), int8(-93), 94, 38, 42, 40, 41, 95, 43, 45, 61, 91, 93, 123, 125, 59, 39, 35, 58, 64, 126, 44, 47, 60, 62, 92, 124, int8(-62), int8(-82), int8(-62), int8(-82), int8(-62), int8(-87), int8(-62), int8(-87)}; _ = shortPatterAsByteArray
            shortPattern = string(shortPatterAsByteArray)
        } else {
            var shortPatterAsByteArray []int8 = []int8{32, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 46, 63, 33, 36, 37, 96, int8(-62), int8(-84), 34, int8(-62), int8(-93), 94, 38, 42, 40, 41, 95, 43, 45, 61, 91, 93, 123, 125, 59, 39, 35, 58, 64, 126, 44, 47, 60, 62, 92, 124, int8(-62), int8(-82), int8(-62), int8(-87)}; _ = shortPatterAsByteArray
            shortPattern = string(shortPatterAsByteArray)
        }
        this.ShortPattern = shortPattern
        this.Scale = scale
        this.CELLS_PER_ROW = 16
        this.FontSize = (20 + 6) * this.Scale
        this.BaseCharWidth = this.FontSize + (6 * this.Scale)
        this.CellSize = this.FontSize + (6 * this.Scale)
        this.TextureSize = this.GetAsTextureSize(this.CELLS_PER_ROW * this.CellSize)
        this.ActualCellsPerRow = this.TextureSize / this.CellSize
    return this
}


func (this *TrueTypeFontUtilBase) GetAsTextureSize(textureSize int) int {
        if textureSize <= 64 {
            textureSize = 64
        } else if textureSize <= 128 {
            textureSize = 128
        } else if textureSize <= 256 {
            textureSize = 256
        } else if textureSize <= 512 {
            textureSize = 512
        } else if textureSize <= 1024 {
            textureSize = 1024
        }
        return textureSize
}


func (this *TrueTypeFontUtilBase) GetCellSize(cellSize int) int {
        if cellSize <= 16 {
            cellSize = 16
        } else if cellSize <= 32 {
            cellSize = 32
        } else if cellSize <= 64 {
            cellSize = 64
        } else if cellSize <= 128 {
            cellSize = 128
        } else if cellSize <= 256 {
            cellSize = 256
        }
        return cellSize
}


func (this *TrueTypeFontUtilBase) GetYOffset() int {
        return ((this.Scale - 1) * this.CellSize)
}

