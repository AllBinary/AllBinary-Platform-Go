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

type CapturedImageWorkerResultsEvent struct {
    AllBinaryEventObject
    Frame int64
    BufferedImage *BufferedImage
}


func (this *CapturedImageWorkerResultsEvent) initFields() {
    
    
}


func NewCapturedImageWorkerResultsEvent(anyType any, frame int64, bufferedImage *BufferedImage) *CapturedImageWorkerResultsEvent {
    this := &CapturedImageWorkerResultsEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(anyType)
        this.initFields()
        this.Frame = frame
        this.SetBufferedImage(bufferedImage)
    return this
}


func (this *CapturedImageWorkerResultsEvent) GetBufferedImage() *BufferedImage {
        return this.BufferedImage
}


func (this *CapturedImageWorkerResultsEvent) SetBufferedImage(bufferedImage *BufferedImage) {
        this.BufferedImage = bufferedImage
}


func (this *CapturedImageWorkerResultsEvent) GetFrame() int64 {
        return this.Frame
}

