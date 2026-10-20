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

type ImageUnifierCell struct {
    Width int
    Height int
}


func (this *ImageUnifierCell) initFields() {
    
    
}


func NewImageUnifierCell(width int, height int) *ImageUnifierCell {
    this := &ImageUnifierCell{}
    this.initFields()
        this.SetWidth(width)
        this.SetHeight(height)
    return this
}


func (this *ImageUnifierCell) GetWidth() int {
        return this.Width
}


func (this *ImageUnifierCell) SetWidth(width int) {
        this.Width = width
}


func (this *ImageUnifierCell) GetHeight() int {
        return this.Height
}


func (this *ImageUnifierCell) SetHeight(height int) {
        this.Height = height
}

