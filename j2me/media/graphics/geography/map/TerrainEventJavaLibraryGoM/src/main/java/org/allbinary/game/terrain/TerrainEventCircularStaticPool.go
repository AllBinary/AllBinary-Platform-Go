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

type TerrainEventCircularStaticPool struct {
    EVENT_POOL *AllBinaryEventCircularPool
}

var TerrainEventCircularStaticPool_Instance *TerrainEventCircularStaticPool = NewTerrainEventCircularStaticPool()

func (this *TerrainEventCircularStaticPool) initFields() {
    this.EVENT_POOL = NewAllBinaryEventCircularPool(20)
}

func NewTerrainEventCircularStaticPool() *TerrainEventCircularStaticPool {
    this := &TerrainEventCircularStaticPool{}
    this.initFields()
    return this
}


func TerrainEventCircularStaticPool_GetInstance() *TerrainEventCircularStaticPool {
        return TerrainEventCircularStaticPool_Instance
}


func (this *TerrainEventCircularStaticPool) Init() {
        this.EVENT_POOL.InitAllBinaryEventCircularPool(NewTerrainEventFactory())
}


// throws Exception
// synchronized
func (this *TerrainEventCircularStaticPool) GetNext(basicTerrainInfo *BasicTerrainInfo) *TerrainEvent {
        var trackingEvent *TerrainEvent = any(this.EVENT_POOL.GetNextInstance()).(*TerrainEvent); _ = trackingEvent
        trackingEvent.SetBasicTerrainInfoForCircularStaticPool(basicTerrainInfo)
        return trackingEvent
}

