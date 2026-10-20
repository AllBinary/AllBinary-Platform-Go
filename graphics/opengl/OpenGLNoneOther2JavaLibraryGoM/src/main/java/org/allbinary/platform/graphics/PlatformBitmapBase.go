/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type PlatformBitmapBase struct {
    Alpha int
    BasicColor *BasicColor
}

// The Java class PlatformBitmapBase has subclasses so its variables are this interface
type IPlatformBitmapBase interface {
    PlatformBitmapBase_() *PlatformBitmapBase
    GetWidth() int
    GetHeight() int
    GetImage() IImage
}

func (this *PlatformBitmapBase) PlatformBitmapBase_() *PlatformBitmapBase {
    return this
}

var PlatformBitmapBase_NULL_PLATFORM_BITMAP_BASE IPlatformBitmapBase = NewPlatformBitmapBase()

func (this *PlatformBitmapBase) initFields() {
    this.Alpha = 255
    this.BasicColor = BasicColorFactory_GetInstance().NULL_COLOR
}

func NewPlatformBitmapBase() *PlatformBitmapBase {
    this := &PlatformBitmapBase{}
    this.initFields()
    return this
}


func (this *PlatformBitmapBase) GetWidth() int {
        return 0
}


func (this *PlatformBitmapBase) GetHeight() int {
        return 0
}


func (this *PlatformBitmapBase) GetImage() IImage {
        return NullImage_NULL_IMAGE
}

