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

type DestroyEventCircularStaticPool struct {
    EVENT_POOL *AllBinaryEventCircularPool
}

var DestroyEventCircularStaticPool_Instance *DestroyEventCircularStaticPool = NewDestroyEventCircularStaticPool()

func (this *DestroyEventCircularStaticPool) initFields() {
    this.EVENT_POOL = NewAllBinaryEventCircularPool(20)
}


func DestroyEventCircularStaticPool_GetInstance() *DestroyEventCircularStaticPool {
        return DestroyEventCircularStaticPool_Instance
}


func NewDestroyEventCircularStaticPool() *DestroyEventCircularStaticPool {
    this := &DestroyEventCircularStaticPool{}
    this.initFields()
    return this
}


func (this *DestroyEventCircularStaticPool) Init(combatGameCanvas IAllBinaryGameCanvas) {
        this.EVENT_POOL.InitAllBinaryEventCircularPool(NewDestroyEventFactory(combatGameCanvas))
}


// throws Exception
func (this *DestroyEventCircularStaticPool) GetInstanceForLayer(layerInterface IAllBinaryLayer) *DestroyedEvent {
        var destroyedEvent *DestroyedEvent = any(this.EVENT_POOL.GetNextInstance()).(*DestroyedEvent); _ = destroyedEvent
        destroyedEvent.SetLayerInterfaceForCircularStaticPool(layerInterface)
        return destroyedEvent
}

