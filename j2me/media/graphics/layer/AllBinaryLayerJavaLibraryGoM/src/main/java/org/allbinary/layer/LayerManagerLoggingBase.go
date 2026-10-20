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

type LayerManagerLoggingBase struct {
}

// The Java class LayerManagerLoggingBase has subclasses so its variables are this interface
type ILayerManagerLoggingBase interface {
    LayerManagerLoggingBase_() *LayerManagerLoggingBase
    Append(layerInterface IAllBinaryLayer)
    AppendAt(layerInterface IAllBinaryLayer, index int)
    Remove(layerInterface IAllBinaryLayer)
    RemoveResult(layerManager ILayerManager, layerInterface IAllBinaryLayer, result bool)
    Clear()
}

func (this *LayerManagerLoggingBase) LayerManagerLoggingBase_() *LayerManagerLoggingBase {
    return this
}


func NewLayerManagerLoggingBase() *LayerManagerLoggingBase {
    this := &LayerManagerLoggingBase{}
    return this
}


// throws Exception
func (this *LayerManagerLoggingBase) Append(layerInterface IAllBinaryLayer) {
}


func (this *LayerManagerLoggingBase) AppendAt(layerInterface IAllBinaryLayer, index int) {
}


func (this *LayerManagerLoggingBase) Remove(layerInterface IAllBinaryLayer) {
}


func (this *LayerManagerLoggingBase) RemoveResult(layerManager ILayerManager, layerInterface IAllBinaryLayer, result bool) {
}


func (this *LayerManagerLoggingBase) Clear() {
}

