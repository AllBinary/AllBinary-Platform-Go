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

type ImageUnifierProperties struct {
    Rows int
    Columns int
    ImageUnifierCell *ImageUnifierCell
}


func (this *ImageUnifierProperties) initFields() {
    
    
    
}


func NewImageUnifierProperties() *ImageUnifierProperties {
    this := &ImageUnifierProperties{}
    this.initFields()
    return this
}


func (this *ImageUnifierProperties) GetWidth() int {
        return int(this.GetColumns()) * int(this.GetImageUnifierCell().GetWidth())
}


func (this *ImageUnifierProperties) GetHeight() int {
        return int(this.GetRows()) * int(this.GetImageUnifierCell().GetHeight())
}


func (this *ImageUnifierProperties) GetRows() int {
        return this.Rows
}


func (this *ImageUnifierProperties) SetRows(rows int) {
        this.Rows = rows
}


func (this *ImageUnifierProperties) GetColumns() int {
        return this.Columns
}


func (this *ImageUnifierProperties) SetColumns(columns int) {
        this.Columns = columns
}


func (this *ImageUnifierProperties) GetImageUnifierCell() *ImageUnifierCell {
        return this.ImageUnifierCell
}


func (this *ImageUnifierProperties) SetImageUnifierCell(imageUnifierCell *ImageUnifierCell) {
        this.ImageUnifierCell = imageUnifierCell
}

