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

type OpenGLImageUtilBase struct {
}

// The Java class OpenGLImageUtilBase has subclasses so its variables are this interface
type IOpenGLImageUtilBase interface {
    OpenGLImageUtilBase_() *OpenGLImageUtilBase
    Paint(args ...any) *OpenGLBitmap
    CreateImage(name string) *OpenGLBitmap
    CreateImageFromInputStream(inputStream IInputStream) *OpenGLBitmap
}

func (this *OpenGLImageUtilBase) OpenGLImageUtilBase_() *OpenGLImageUtilBase {
    return this
}


func NewOpenGLImageUtilBase() *OpenGLImageUtilBase {
    this := &OpenGLImageUtilBase{}
    return this
}


func (this *OpenGLImageUtilBase) Paint_1(alphaBitmap *OpenGLBitmap, originalBitmap *OpenGLBitmap, index int) *OpenGLBitmap {
        panic(errors.New(fmt.Sprint()))
}

// Java overloads of OpenGLImageUtilBase.paint
func (this *OpenGLImageUtilBase) Paint(args ...any) *OpenGLBitmap {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[*OpenGLBitmap](args[0], pass) && Java_IsArg[*OpenGLBitmap](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return this.Paint_1(Java_Arg[*OpenGLBitmap](args[0]), Java_Arg[*OpenGLBitmap](args[1]), Java_Arg[int](args[2]))
        }
        if len(args) == 4 && Java_IsArg[*OpenGLBitmap](args[0], pass) && Java_IsArg[*OpenGLBitmap](args[1], pass) && Java_IsArg[int](args[2], pass) && Java_IsArg[*BasicColor](args[3], pass) {
            return this.Paint_2(Java_Arg[*OpenGLBitmap](args[0]), Java_Arg[*OpenGLBitmap](args[1]), Java_Arg[int](args[2]), Java_Arg[*BasicColor](args[3]))
        }
    }
    panic("No Java overload of OpenGLImageUtilBase.paint for the arguments")
}


func (this *OpenGLImageUtilBase) Paint_2(colorBitmap *OpenGLBitmap, originalBitmap *OpenGLBitmap, alpha int, basicColor *BasicColor) *OpenGLBitmap {
        panic(errors.New(fmt.Sprint()))
}


// throws IOException
func (this *OpenGLImageUtilBase) CreateImage(name string) *OpenGLBitmap {
        panic(errors.New(fmt.Sprint()))
}


// throws IOException
func (this *OpenGLImageUtilBase) CreateImageFromInputStream(inputStream IInputStream) *OpenGLBitmap {
        panic(errors.New(fmt.Sprint()))
}

