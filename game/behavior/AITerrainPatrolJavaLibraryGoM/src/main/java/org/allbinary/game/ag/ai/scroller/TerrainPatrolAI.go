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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type TerrainPatrolAI struct {
    PacePatrolAI
    TerrainEventListener *TerrainEventListener
    DOWN IAngle
    CLIFF *BasicTerrainInfo
}


func (this *TerrainPatrolAI) initFields() {
    this.TerrainEventListener = NewTerrainEventListener()
    this.DOWN = AngleFactory_GetInstance().DOWN
    this.CLIFF = NewBasicTerrainInfo(this.DOWN)
}


func NewTerrainPatrolAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *TerrainPatrolAI {
    this := &TerrainPatrolAI{}
        this.PacePatrolAI = *NewPacePatrolAI(hashtable, ownerLayerInterface, gameInput)
        this.initFields()
        TerrainEventHandler_GetInstance(ownerLayerInterface).AddListener(this.TerrainEventListener)
        this.TerrainEventListener.OnTerrainEvent(TerrainEventCircularStaticPool_GetInstance().GetNext(this.CLIFF))
    return this
}


func (this *TerrainPatrolAI) Update() {
        this.PacePatrolAI.Update()
        this.ChangeDirectionIfCliffReached()
}


func (this *TerrainPatrolAI) ChangeDirectionIfCliffReached() {
        var list IBasicArrayList = this.TerrainEventListener.GetList(); _ = list
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var terrainEvent *TerrainEvent = any(list.RemoveAt(index)).(*TerrainEvent); _ = terrainEvent
            var basicTerrainInfo *BasicTerrainInfo = terrainEvent.GetBasicTerrainInfo(); _ = basicTerrainInfo
            var angle IAngle = basicTerrainInfo.GetAngle(); _ = angle
            if angle == this.DOWN {
                this.NextDirection()
                if !this.IsFollowLimitedByTerrain {
                    this.LogUtil.PutF("Following Limited", this, "onTerrainEvent")
                    this.IsFollowLimitedByTerrain = true
                }
            }
        }
}

