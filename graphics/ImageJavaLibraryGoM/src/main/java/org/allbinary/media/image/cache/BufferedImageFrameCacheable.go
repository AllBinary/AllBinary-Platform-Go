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

type BufferedImageFrameCacheable struct {
    Frame int64
    BufferedImage *BufferedImage
}

// Implements CacheableInterface


func (this *BufferedImageFrameCacheable) initFields() {
    
    
}


func NewBufferedImageFrameCacheable(bufferedImage *BufferedImage, frame int64) *BufferedImageFrameCacheable {
    this := &BufferedImageFrameCacheable{}
    this.initFields()
        this.Frame = frame
        this.SetBufferedImage(bufferedImage)
    return this
}


func (this *BufferedImageFrameCacheable) GetKey() any {
        return this.Frame
}


func (this *BufferedImageFrameCacheable) GetBufferedImage() *BufferedImage {
        return this.BufferedImage
}


func (this *BufferedImageFrameCacheable) SetBufferedImage(bufferedImage *BufferedImage) {
        this.BufferedImage = bufferedImage
}


func (this *BufferedImageFrameCacheable) GetFrame() int64 {
        return this.Frame
}


func (this *BufferedImageFrameCacheable) SetFrame(frame int64) {
        this.Frame = frame
}


func (this *BufferedImageFrameCacheable) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(" Frame: ")
        stringBuffer.Appendlong(this.GetFrame())
        return fmt.Sprint(stringBuffer)
}

func (this *BufferedImageFrameCacheable) String() string {
    return this.ToString()
}

