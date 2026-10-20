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

type TrueTypeFontUtil struct {
    TrueTypeFontUtilBase
    FontImage IOpenGLESImage
}

var TrueTypeFontUtil_Instance *TrueTypeFontUtil = NewTrueTypeFontUtil()

func (this *TrueTypeFontUtil) initFields() {
    this.FontImage = OpenGLESImage_NULL_OPENGL_IMAGE
}


func TrueTypeFontUtil_GetInstance() *TrueTypeFontUtil {
        return TrueTypeFontUtil_Instance
}


func NewTrueTypeFontUtil() *TrueTypeFontUtil {
    this := &TrueTypeFontUtil{}
        this.TrueTypeFontUtilBase = *NewTrueTypeFontUtilBase(1)
        this.initFields()
    return this
}


func (this *TrueTypeFontUtil) GetAsTextureSize(textureSize int) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *TrueTypeFontUtil) GetCellSize(cellSize int) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *TrueTypeFontUtil) GetFontBitmap(filename string, fontSize int, cellSize int, basicColor *BasicColor) IImage {
        panic(errors.New(fmt.Sprint()))
}


func (this *TrueTypeFontUtil) SaveFontAtlasAsFile() {
}


func (this *TrueTypeFontUtil) GetFontBitmapGL(gl GL10, filename string, cellSize int, basicColor *BasicColor) IImage {
        panic(errors.New(fmt.Sprint()))
}


func (this *TrueTypeFontUtil) GetFontWidths(filename string, fontSize int) []int {
        panic(errors.New(fmt.Sprint()))
}

