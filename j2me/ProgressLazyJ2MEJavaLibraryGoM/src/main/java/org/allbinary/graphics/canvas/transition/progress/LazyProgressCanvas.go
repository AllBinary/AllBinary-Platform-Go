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

type LazyProgressCanvas struct {
    ProgressCanvas
}



func NewLazyProgressCanvas(title string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *LazyProgressCanvas {
    this := &LazyProgressCanvas{}
        this.ProgressCanvas = *NewProgressCanvas(title, backgroundBasicColor, foregroundBasicColor)
    return this
}


func (this *LazyProgressCanvas) Start() {
        this.ProgressCanvas.Start()
        this.HasPainted = false
}


func (this *LazyProgressCanvas) End() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION, this, this.CommonStrings.END_METHOD_NAME)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.END_METHOD_NAME)
            this.EndActual()
            this.Paintable = this.GAUGE_PAINTABLE
            ImageCacheFactory_GetInstance().RunTask()
            ImageCacheFactory_GetInstance().ProgressEnded()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LazyProgressCanvas) InGame() {
        this.InGameProcessor = Processor_GetInstance()
}


func (this *LazyProgressCanvas) EndFromInitialLazyLoadingComplete() {
        this.ProgressCanvas.EndFromInitialLazyLoadingComplete()
        this.Paintable = NullPaintable_GetInstance()
}


func (this *LazyProgressCanvas) EndIfPaintedSinceStart() {
        if this.Paintable == this.GAUGE_PAINTABLE && this.HasPainted {
            this.EndFromInitialLazyLoadingComplete()
        }
}

