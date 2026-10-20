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

type OpenGLESStrings struct {
}

// The Java class OpenGLESStrings has subclasses so its variables are this interface
type IOpenGLESStrings interface {
    OpenGLESStrings_() *OpenGLESStrings
    Set(gl GL10)
    InitFont(gl GL10, width int, height int)
    Update(gl GL10, width int, height int)
    DrawChar(gl GL10, graphics IGraphics, character rune, x int, y int, anchor int)
    DrawChars(gl GL10, graphics IGraphics, data []rune, offset int, length int, x int, y int, anchor int)
    DrawString(gl GL10, graphics IGraphics, string_ string, x int, y int, anchor int)
    DrawSubstring(gl GL10, graphics IGraphics, string_ string, offset int, len_ int, x int, y int, anchor int)
    CharWidth(ch rune) int
    CharsWidth(ch []rune, offset int, length int) int
    StringWidth(args ...any) int
}

func (this *OpenGLESStrings) OpenGLESStrings_() *OpenGLESStrings {
    return this
}


func NewOpenGLESStrings() *OpenGLESStrings {
    this := &OpenGLESStrings{}
    return this
}


// throws Exception
func (this *OpenGLESStrings) Set(gl GL10) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) InitFont(gl GL10, width int, height int) {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *OpenGLESStrings) Update(gl GL10, width int, height int) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) DrawChar(gl GL10, graphics IGraphics, character rune, x int, y int, anchor int) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) DrawChars(gl GL10, graphics IGraphics, data []rune, offset int, length int, x int, y int, anchor int) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) DrawString(gl GL10, graphics IGraphics, string_ string, x int, y int, anchor int) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) DrawSubstring(gl GL10, graphics IGraphics, string_ string, offset int, len_ int, x int, y int, anchor int) {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) CharWidth(ch rune) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) CharsWidth(ch []rune, offset int, length int) int {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLESStrings) StringWidth_1(string_ string) int {
        panic(errors.New(fmt.Sprint()))
}

// Java overloads of OpenGLESStrings.stringWidth
func (this *OpenGLESStrings) StringWidth(args ...any) int {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.StringWidth_1(Java_Arg[string](args[0]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return this.StringWidth_2(Java_Arg[string](args[0]), Java_Arg[int](args[1]), Java_Arg[int](args[2]))
        }
    }
    panic("No Java overload of OpenGLESStrings.stringWidth for the arguments")
}


func (this *OpenGLESStrings) StringWidth_2(string_ string, offset int, len_ int) int {
        panic(errors.New(fmt.Sprint()))
}

