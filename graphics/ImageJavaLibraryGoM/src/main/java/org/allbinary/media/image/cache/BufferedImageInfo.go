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

type BufferedImageInfo struct {
    Width int
    Height int
    Type int
}


func (this *BufferedImageInfo) initFields() {
    
    
    
}


func NewBufferedImageInfo(width int, height int, type_ int) *BufferedImageInfo {
    this := &BufferedImageInfo{}
    this.initFields()
        this.Width = width
        this.Height = height
        this.Type = type_
    return this
}


func (this *BufferedImageInfo) GetWidth() int {
        return this.Width
}


func (this *BufferedImageInfo) SetWidth(width int) {
        this.Width = width
}


func (this *BufferedImageInfo) GetHeight() int {
        return this.Height
}


func (this *BufferedImageInfo) SetHeight(height int) {
        this.Height = height
}


func (this *BufferedImageInfo) GetType() int {
        return this.Type
}


func (this *BufferedImageInfo) SetType(type_ int) {
        this.Type = type_
}


func (this *BufferedImageInfo) ToString() string {
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("BufferedImageInfo -")
        stringBuffer.Append(commonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(this.GetWidth())
        stringBuffer.Append(commonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(this.GetHeight())
        stringBuffer.Append(" Type: ")
        stringBuffer.Appendint(this.GetType())
        return fmt.Sprint(stringBuffer)
}

func (this *BufferedImageInfo) String() string {
    return this.ToString()
}

