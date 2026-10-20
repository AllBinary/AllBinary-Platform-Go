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

type BasicColorFactory struct {
    NULL_COLOR *BasicColor
    TAN *BasicColor
    SKIN_PINK *BasicColor
    RED *BasicColor
    LIGHT_RED *BasicColor
    DARK_GREEN *BasicColor
    GREEN *BasicColor
    LIGHT_GREEN *BasicColor
    PURPLE *BasicColor
    BLUE *BasicColor
    LIGHT_BLUE *BasicColor
    VERY_LIGHT_BLUE *BasicColor
    PUCE *BasicColor
    AQUA *BasicColor
    YELLOW *BasicColor
    PINK *BasicColor
    BLACK *BasicColor
    GREY *BasicColor
    LIGHT_GREY *BasicColor
    DARK_GREY *BasicColor
    ORANGE *BasicColor
    BROWN *BasicColor
    WHITE *BasicColor
    TRANSPARENT_COLOR *BasicColor
    TRANSPARENT_GREY *BasicColor
    TRANSPARENT_BLACK *BasicColor
    TRANSPARENT_RED *BasicColor
    CLEAR_COLOR *BasicColor
    BORDER_COLOR *BasicColor
}

var BasicColorFactory_Instance *BasicColorFactory = NewBasicColorFactory()

func (this *BasicColorFactory) initFields() {
    this.NULL_COLOR = this.CreateInstanceARGB(255, 255, 255, 255, "null color")
    this.TAN = this.CreateInstance(0xE9D8B5)
    this.SKIN_PINK = this.CreateInstance(0xFDCECE)
    this.RED = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0xFF0000, "red")
    this.LIGHT_RED = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0xFFA0A0, "light red")
    this.DARK_GREEN = this.CreateInstance(0x00FF00)
    this.GREEN = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0x00DD00, "green")
    this.LIGHT_GREEN = this.CreateInstance(0x007700)
    this.PURPLE = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0x400040, "purple")
    this.BLUE = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0x0000FF, "blue")
    this.LIGHT_BLUE = this.CreateInstance(0x00aaff)
    this.VERY_LIGHT_BLUE = this.CreateInstance(0x000033)
    this.PUCE = this.CreateInstance(0xFF00AA)
    this.AQUA = this.CreateInstance(0x00FFFF)
    this.YELLOW = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0xFFFF00, "yellow")
    this.PINK = this.CreateInstance(0xFF00FF)
    this.BLACK = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0x000000, "black")
    this.GREY = this.CreateInstanceAN(BasicColorUtil_GetInstance().ALPHA, 0xAAAAAA, "grey")
    this.LIGHT_GREY = this.CreateInstance(0xC0C0C0)
    this.DARK_GREY = this.CreateInstance(0x404040)
    this.ORANGE = this.CreateInstance(0xe07718)
    this.BROWN = this.CreateInstance(0x956B00)
    this.WHITE = this.CreateInstance(0xFFFFFF)
    this.TRANSPARENT_COLOR = this.CreateInstanceAN(0x00000000, 0x000000, StringUtil_GetInstance().EMPTY_STRING)
    this.TRANSPARENT_GREY = this.CreateInstanceAN(0x52000000, 0xECECEC, StringUtil_GetInstance().EMPTY_STRING)
    this.TRANSPARENT_BLACK = this.CreateInstanceAN(0x52000000, 0x000000, StringUtil_GetInstance().EMPTY_STRING)
    this.TRANSPARENT_RED = this.CreateInstanceAN(0x52000000, 0xFF0000, StringUtil_GetInstance().EMPTY_STRING)
    this.CLEAR_COLOR = this.BLACK
    this.BORDER_COLOR = this.WHITE
}

func NewBasicColorFactory() *BasicColorFactory {
    this := &BasicColorFactory{}
    this.initFields()
    return this
}


func BasicColorFactory_GetInstance() *BasicColorFactory {
        return BasicColorFactory_Instance
}


func (this *BasicColorFactory) CreateInstance(value int) *BasicColor {
        return NewBasicColor(BasicColorUtil_GetInstance().ALPHA, value, StringUtil_GetInstance().EMPTY_STRING)
}


func (this *BasicColorFactory) CreateInstanceA(alphaValue int, value int) *BasicColor {
        return NewBasicColor(alphaValue, value, StringUtil_GetInstance().EMPTY_STRING)
}


func (this *BasicColorFactory) CreateInstanceAN(alphaValue int, value int, name string) *BasicColor {
        return NewBasicColor(alphaValue, value, name)
}


func (this *BasicColorFactory) CreateInstanceARGB(alphaValue int, r int, g int, b int, name string) *BasicColor {
        var ALPHA_MASK int = 0xFF000000; _ = ALPHA_MASK
        return NewBasicColor(alphaValue, ((alphaValue << 24) & ALPHA_MASK) + ((r << 16) & 0x00FF0000) + ((g << 8) & 0x0000FF00) + (b & 0x000000FF), StringUtil_GetInstance().EMPTY_STRING)
}
