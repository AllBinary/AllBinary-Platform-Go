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

type BasicColorUtil struct {
    ZERO_ARRAY []*BasicColor
    ALPHA int
    IsAlpha bool
    FfOpaque bool
}

var BasicColorUtil_Instance *BasicColorUtil = NewBasicColorUtil()

func (this *BasicColorUtil) initFields() {
    this.ZERO_ARRAY = make([]*BasicColor, 0)
    this.ALPHA = 0xFF000000
    this.IsAlpha = true
    this.FfOpaque = true
}

func NewBasicColorUtil() *BasicColorUtil {
    this := &BasicColorUtil{}
    this.initFields()
    return this
}


func BasicColorUtil_GetInstance() *BasicColorUtil {
        return BasicColorUtil_Instance
}


func (this *BasicColorUtil) GetRGB(red int, green int, blue int) int {
        var value int = 0; _ = value
        if this.IsAlpha {
            if this.FfOpaque {
                value = this.ALPHA
            }
        }
        value = (red << 16) | value
        value = (green << 8) | value
        value = blue | value
        return value
}


func (this *BasicColorUtil) GetARGB(alphaValue int, r int, g int, b int) int {
        return ((alphaValue << 24) & this.ALPHA) + ((r << 16) & 0x00FF0000) + ((g << 8) & 0x0000FF00) + (b & 0x000000FF)
}


func (this *BasicColorUtil) Get(alphaValue int, color int) int {
        return ((alphaValue << 24) & this.ALPHA) + color
}


func (this *BasicColorUtil) Invert(color int) int {
        return (0xFFFFFF - color) | this.ALPHA
}

