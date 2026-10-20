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

type Vertex struct {
    X *BasicDecimal
    Y *BasicDecimal
    Z *BasicDecimal
}


func (this *Vertex) initFields() {
    this.X = NewBasicDecimal(0)
    this.Y = NewBasicDecimal(0)
    this.Z = NewBasicDecimal(0)
}


func NewVertex() *Vertex {
    this := &Vertex{}
    this.initFields()
    return this
}


func (this *Vertex) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var positionStrings *PositionStrings = PositionStrings_GetInstance(); _ = positionStrings
        stringBuffer.Append(positionStrings.X_LABEL)
        stringBuffer.Append(fmt.Sprint(this.X))
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(positionStrings.Y_LABEL)
        stringBuffer.Append(fmt.Sprint(this.Y))
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(positionStrings.Z_LABEL)
        stringBuffer.Append(fmt.Sprint(this.Z))
        return fmt.Sprint(stringBuffer)
}

func (this *Vertex) String() string {
    return this.ToString()
}

