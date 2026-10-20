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

type IIOImageCacheable struct {
    Key string
    IioImage *IIOImage
}

// Implements CacheableInterface


func (this *IIOImageCacheable) initFields() {
    
    
}


func NewIIOImageCacheable(key string) *IIOImageCacheable {
    this := &IIOImageCacheable{}
    this.initFields()
        this.Key = key
    return this
}


func (this *IIOImageCacheable) GetKey() any {
        return this.Key
}


func (this *IIOImageCacheable) SetBufferedImage(bufferedImage *BufferedImage) {
        this.SetIioImage(NewIIOImage(bufferedImage, nil, nil))
}


func (this *IIOImageCacheable) GetIioImage() *IIOImage {
        return this.IioImage
}


func (this *IIOImageCacheable) SetIioImage(iioImage *IIOImage) {
        this.IioImage = iioImage
}


func (this *IIOImageCacheable) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(reflect.TypeOf(this).String())
        stringBuffer.Append(" Reusable IIOImage")
        return fmt.Sprint(stringBuffer)
}

func (this *IIOImageCacheable) String() string {
    return this.ToString()
}

