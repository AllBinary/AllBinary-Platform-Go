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

type BufferedImageCacheable struct {
    Key *BufferedImageInfo
    BufferedImage *BufferedImage
}

// Implements CacheableInterface


func (this *BufferedImageCacheable) initFields() {
    
    
}


func NewBufferedImageCacheable(key *BufferedImageInfo) *BufferedImageCacheable {
    this := &BufferedImageCacheable{}
    this.initFields()
        this.Key = key
        this.SetBufferedImage(NewBufferedImage(key.GetWidth(), key.GetHeight(), key.GetType()))
    return this
}


func (this *BufferedImageCacheable) GetKey() any {
        return this.Key
}


func (this *BufferedImageCacheable) GetBufferedImage() *BufferedImage {
        return this.BufferedImage
}


func (this *BufferedImageCacheable) SetBufferedImage(bufferedImage *BufferedImage) {
        this.BufferedImage = bufferedImage
}


func (this *BufferedImageCacheable) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(" ")
        stringBuffer.Append(fmt.Sprint(this.GetKey()))
        return fmt.Sprint(stringBuffer)
}

func (this *BufferedImageCacheable) String() string {
    return this.ToString()
}

