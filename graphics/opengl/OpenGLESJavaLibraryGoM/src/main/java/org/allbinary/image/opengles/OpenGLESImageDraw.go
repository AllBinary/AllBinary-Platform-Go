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

type OpenGLESImageDraw struct {
}

var OpenGLESImageDraw_Instance *OpenGLESImageDraw = NewOpenGLESImageDraw()

func NewOpenGLESImageDraw() *OpenGLESImageDraw {
    this := &OpenGLESImageDraw{}
    return this
}


func OpenGLESImageDraw_GetInstance() *OpenGLESImageDraw {
        return OpenGLESImageDraw_Instance
}


func (this *OpenGLESImageDraw) DrawRegion(gl GL10, viewHeight int, x_src float32, y_src float32, width float32, height float32, x int, y int, z int) {
}


func (this *OpenGLESImageDraw) Draw(gl GL10, x int, y int, z int) {
}

