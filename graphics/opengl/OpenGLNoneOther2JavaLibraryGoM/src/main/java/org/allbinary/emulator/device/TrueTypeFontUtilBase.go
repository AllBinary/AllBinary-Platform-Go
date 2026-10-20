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
}

// The Java class TrueTypeFontUtilBase has subclasses so its variables are this interface
type ITrueTypeFontUtilBase interface {
    TrueTypeFontUtilBase_() *TrueTypeFontUtilBase
    GetAsTextureSize(textureSize int) int
    GetCellSize(cellSize int) int
}

func (this *TrueTypeFontUtilBase) TrueTypeFontUtilBase_() *TrueTypeFontUtilBase {
    return this
}



func NewTrueTypeFontUtilBase(scale int) *TrueTypeFontUtilBase {
    this := &TrueTypeFontUtilBase{}
    return this
}


func (this *TrueTypeFontUtilBase) GetAsTextureSize(textureSize int) int {
        return -1
}


func (this *TrueTypeFontUtilBase) GetCellSize(cellSize int) int {
        return -1
}

