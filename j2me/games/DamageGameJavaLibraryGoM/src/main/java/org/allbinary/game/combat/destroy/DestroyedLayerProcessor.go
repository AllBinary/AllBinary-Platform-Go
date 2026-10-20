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

type DestroyedLayerProcessor struct {
    BasicLayerProcessor
}

var DestroyedLayerProcessor_DESTROYED_LAYER_PROCESSOR IBasicLayerProcessor = BasicLayerProcessor_NULL_LAYER_PROCESSOR


func NewDestroyedLayerProcessor() *DestroyedLayerProcessor {
    this := &DestroyedLayerProcessor{}
    this.BasicLayerProcessor = *NewBasicLayerProcessor()
    return this
}


func DestroyedLayerProcessor_Init() {
        DestroyedLayerProcessor_DESTROYED_LAYER_PROCESSOR = NewDestroyedLayerProcessor()
}


func DestroyedLayerProcessor_GetInstance() IBasicLayerProcessor {
        return DestroyedLayerProcessor_DESTROYED_LAYER_PROCESSOR
}


// throws Exception
func (this *DestroyedLayerProcessor) Process(allBinaryLayerManager IAllBinaryLayerManager) {
        var list IBasicArrayList = this.GetList(); _ = list
        var size int = list.Size(); _ = size
        var destroyedEventHandler *DestroyedEventHandler = DestroyedEventHandler_GetInstance(); _ = destroyedEventHandler
        var destroyEventCircularStaticPool *DestroyEventCircularStaticPool = DestroyEventCircularStaticPool_GetInstance(); _ = destroyEventCircularStaticPool
        var destroyedEvent *DestroyedEvent; _ = destroyedEvent
        var layerInterface IAllBinaryLayer; _ = layerInterface
        for index := int(0); index < size; index++ {
            var layerInterfaceCanBeNull any = list.BasicArrayList_().ObjectArray[index]; _ = layerInterfaceCanBeNull
            if layerInterfaceCanBeNull != nil {
                layerInterface = any(layerInterfaceCanBeNull).(IAllBinaryLayer)
                allBinaryLayerManager.Remove(layerInterface)
                destroyedEvent = destroyEventCircularStaticPool.GetInstanceForLayer(layerInterface)
                destroyedEventHandler.FireEvent(destroyedEvent)
            } else {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            }
        }
        list.Clear()
}

