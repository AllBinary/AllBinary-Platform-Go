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

type ImageComparisonResultFrameCacheable struct {
    Frame int64
    ImageComparisonResult *ImageComparisonResult
}

// Implements CacheableInterface


func (this *ImageComparisonResultFrameCacheable) initFields() {
    
    
}


func NewImageComparisonResultFrameCacheable(imageComparisonResult *ImageComparisonResult, frame int64) *ImageComparisonResultFrameCacheable {
    this := &ImageComparisonResultFrameCacheable{}
    this.initFields()
        this.SetImageComparisonResult(imageComparisonResult)
        this.Frame = frame
    return this
}


func (this *ImageComparisonResultFrameCacheable) GetKey() any {
        return this.Frame
}


func (this *ImageComparisonResultFrameCacheable) GetFrame() int64 {
        return this.Frame
}


func (this *ImageComparisonResultFrameCacheable) SetFrame(frame int64) {
        this.Frame = frame
}


func (this *ImageComparisonResultFrameCacheable) GetImageComparisonResult() *ImageComparisonResult {
        return this.ImageComparisonResult
}


func (this *ImageComparisonResultFrameCacheable) SetImageComparisonResult(imageComparisonResult *ImageComparisonResult) {
        this.ImageComparisonResult = imageComparisonResult
}


func (this *ImageComparisonResultFrameCacheable) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(" Frame: ")
        stringBuffer.Appendlong(this.GetFrame())
        return fmt.Sprint(stringBuffer)
}

func (this *ImageComparisonResultFrameCacheable) String() string {
    return this.ToString()
}

