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

type BaseScalable struct {
}

// Implements ScalableListener

// The Java class BaseScalable has subclasses so its variables are this interface
type IBaseScalable interface {
    BaseScalable_() *BaseScalable
    GetRatio(ratio float32) float32
    GetLeft(xOffset int) int
    GetTop(yOffset int) int
    Scale(ratio float32)
}

func (this *BaseScalable) BaseScalable_() *BaseScalable {
    return this
}


func NewBaseScalable() *BaseScalable {
    this := &BaseScalable{}
    return this
}


func (this *BaseScalable) GetRatio(ratio float32) float32 {
        return 1.0
}


func (this *BaseScalable) GetLeft(xOffset int) int {
        return -1
}


func (this *BaseScalable) GetTop(yOffset int) int {
        return -1
}


func (this *BaseScalable) Scale(ratio float32) {
}

