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

type KeyValueDrawString struct {
    MyFontProcessor IMyFontProcessor
    LABEL string
    LabelX int
    ValueX int
    Value string
}

// Implements UpdateMyFontInterface


func (this *KeyValueDrawString) initFields() {
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    
    
    
    this.Value = StringUtil_GetInstance().EMPTY_STRING
}


func NewKeyValueDrawString(label string, x int) *KeyValueDrawString {
    this := &KeyValueDrawString{}
    this.initFields()
        this.LABEL = label
        this.LabelX = x
    return this
}


func (this *KeyValueDrawString) UpdateMeasurement(graphics IGraphics) {
        var labelWidth int = graphics.GetFont().StringWidth(this.LABEL) - len(this.LABEL); _ = labelWidth
        this.ValueX = labelWidth + this.LabelX
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *KeyValueDrawString) Paint(graphics IGraphics, y int) {
        this.MyFontProcessor.Process(graphics)
        graphics.DrawString(this.LABEL, this.LabelX, y, 0)
        graphics.DrawString(this.Value, this.ValueX, y, 0)
}


func (this *KeyValueDrawString) Update(value string) {
        if value != "" {
            this.Value = value
        }
}

