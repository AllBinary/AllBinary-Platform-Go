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

type RendererStrings struct {
    ON_SURFACE_CREATED string
    ON_SURFACE_CHANGED string
    ON_DRAW_FRAME string
    ON_DRAW_FRAME_THREED string
    ON_DRAW_FRAME_2D string
}

var RendererStrings_Instance *RendererStrings = NewRendererStrings()

func (this *RendererStrings) initFields() {
    this.ON_SURFACE_CREATED = "onSurfaceCreated"
    this.ON_SURFACE_CHANGED = "onSurfaceChanged"
    this.ON_DRAW_FRAME = "onDrawFrame"
    this.ON_DRAW_FRAME_THREED = "onDrawFrameThreed"
    this.ON_DRAW_FRAME_2D = "onDrawFrame2d"
}

func NewRendererStrings() *RendererStrings {
    this := &RendererStrings{}
    this.initFields()
    return this
}


func RendererStrings_GetInstance() *RendererStrings {
        return RendererStrings_Instance
}

