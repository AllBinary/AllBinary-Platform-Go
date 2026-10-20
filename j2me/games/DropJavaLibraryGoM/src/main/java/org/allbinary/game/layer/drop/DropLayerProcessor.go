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

type DropLayerProcessor struct {
    BasicLayerProcessor
}

var DropLayerProcessor_LAYER_PROCESSOR IBasicLayerProcessor = NewDropLayerProcessor()


func NewDropLayerProcessor() *DropLayerProcessor {
    this := &DropLayerProcessor{}
    this.BasicLayerProcessor = *NewBasicLayerProcessor()
    return this
}


func DropLayerProcessor_GetInstance() IBasicLayerProcessor {
        return DropLayerProcessor_LAYER_PROCESSOR
}


// throws Exception
func (this *DropLayerProcessor) Process(allBinaryLayerManager IAllBinaryLayerManager) {
        var list IBasicArrayList = this.GetList(); _ = list
        var size int = list.Size(); _ = size
        var layerInterface IAllBinaryLayer; _ = layerInterface
        var dropLayerInterface DropLayerInterface; _ = dropLayerInterface
        for index := int(0); index < size; index++ {
            layerInterface = any(list.BasicArrayList_().ObjectArray[index]).(IAllBinaryLayer)
            dropLayerInterface = any(layerInterface).(DropLayerInterface)
            layerInterface = any(dropLayerInterface.GetDroppedLayer()).(IAllBinaryLayer)
            allBinaryLayerManager.Append(layerInterface)
        }
        list.Clear()
}

