/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type PlayerOwnedRTSLayers struct {
    RtsLayerList IBasicArrayList
}

// The Java class PlayerOwnedRTSLayers has subclasses so its variables are this interface
type IPlayerOwnedRTSLayers interface {
    PlayerOwnedRTSLayers_() *PlayerOwnedRTSLayers
    ToString() string
}

func (this *PlayerOwnedRTSLayers) PlayerOwnedRTSLayers_() *PlayerOwnedRTSLayers {
    return this
}


func (this *PlayerOwnedRTSLayers) initFields() {
    this.RtsLayerList = NewBasicArrayListD()
}

func NewPlayerOwnedRTSLayers() *PlayerOwnedRTSLayers {
    this := &PlayerOwnedRTSLayers{}
    this.initFields()
    return this
}


func (this *PlayerOwnedRTSLayers) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("All: ")
        stringBuffer.Appendint(this.RtsLayerList.Size())
        return fmt.Sprint(stringBuffer)
}

func (this *PlayerOwnedRTSLayers) String() string {
    return this.ToString()
}

