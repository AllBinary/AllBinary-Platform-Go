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

type PacePatrolAI struct {
    BasePatrolAI
    TrackingList IBasicArrayList
    FiringDistance int
    IsFollowLimitedByTerrain bool
}

// Implements TrackingEventListenerInterface

// The Java class PacePatrolAI has subclasses so its variables are this interface
type IPacePatrolAI interface {
    IBasePatrolAI
    PacePatrolAI_() *PacePatrolAI
    ProcessAI(allBinaryLayerManager IAllBinaryLayerManager)
    SetFiringDirectionForTargetIfInRange() *Direction
    Disable()
    OnEvent(eventObject IAllBinaryEventObject)
    OnMovement(trackingEvent *TrackingEvent)
    GetFiringDistance() int
    SetFiringDistance(firingDistance int)
}

func (this *PacePatrolAI) PacePatrolAI_() *PacePatrolAI {
    return this
}


func (this *PacePatrolAI) initFields() {
    
    
    this.IsFollowLimitedByTerrain = false
}


func NewPacePatrolAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *PacePatrolAI {
    this := &PacePatrolAI{}
        this.BasePatrolAI = *NewBasePatrolAI(hashtable, ownerLayerInterface, gameInput)
        this.initFields()
        TrackingEventHandler_GetInstance().AddListenerInterface(this)
        this.TrackingList = NewBasicArrayListD()
    return this
}


// throws Exception
func (this *PacePatrolAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        this.Update()
        var direction *Direction = this.SetFiringDirectionForTargetIfInRange(); _ = direction
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if direction == directionFactory.LEFT {
            if this.LastKeyDirection != this.KeyDirection || !this.IsFollowLimitedByTerrain {
                this.BasePatrolAI.ProcessKeyAI(Canvas_LEFT)
            }
            this.KeyDirection = Canvas_LEFT
            this.LastKeyDirection = this.KeyDirection
            this.XTotalDistance = 0
            this.BasePatrolAI.ProcessKeyAI(Canvas_KEY_NUM1)
        } else if direction == directionFactory.RIGHT {
            if this.LastKeyDirection != this.KeyDirection || !this.IsFollowLimitedByTerrain {
                this.BasePatrolAI.ProcessKeyAI(Canvas_RIGHT)
            }
            this.KeyDirection = Canvas_RIGHT
            this.LastKeyDirection = this.KeyDirection
            this.XTotalDistance = 0
            this.BasePatrolAI.ProcessKeyAI(Canvas_KEY_NUM1)
        } else {
            this.BasePatrolAI.ProcessKeyAI(this.KeyDirection)
        }
}


func (this *PacePatrolAI) SetFiringDirectionForTargetIfInRange() *Direction {
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var direction *Direction = directionFactory.NOT_BORDERED_WITH; _ = direction
        var lastTrackingEvent *TrackingEvent; _ = lastTrackingEvent
        var lastTrackingLayerInterface LayerInterface; _ = lastTrackingLayerInterface
        var directionCompositeInterface DirectionCompositeInterface; _ = directionCompositeInterface
        var size int = this.TrackingList.Size(); _ = size
        for index := int(0); index < size; index++ {
            lastTrackingEvent = any(this.TrackingList.Get(0)).(*TrackingEvent)
            lastTrackingLayerInterface = lastTrackingEvent.GetLayerInterface()
            var x int = lastTrackingLayerInterface.GetXP(); _ = x
            var y int = lastTrackingLayerInterface.GetYP(); _ = y
            var yDistance int = ownerLayerInterface.GetYP() - y - ownerLayerInterface.GetHeight(); _ = yDistance
            var xDistance int = ownerLayerInterface.GetXP() - x - ownerLayerInterface.GetWidth(); _ = xDistance
            var absXDistance int = max(xDistance, -(xDistance)); _ = absXDistance
            var absYDistance int = max(yDistance, -(yDistance)); _ = absYDistance
            if absYDistance <= 100 {
                if absXDistance < this.GetFiringDistance() / 2 {
                    directionCompositeInterface = any(this.GetOwnerLayerInterface()).(DirectionCompositeInterface)
                    if xDistance < 0 && directionCompositeInterface.GetDirection() == directionFactory.RIGHT {
                        direction = directionFactory.RIGHT
                    } else if xDistance > 0 && directionCompositeInterface.GetDirection() == directionFactory.LEFT {
                        direction = directionFactory.LEFT
                    }
                }
            }
        }
        return direction
}


func (this *PacePatrolAI) Disable() {
}


func (this *PacePatrolAI) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *PacePatrolAI) OnMovement(trackingEvent *TrackingEvent) {
        this.TrackingList.Clear()
        this.TrackingList.Add(trackingEvent)
}


func (this *PacePatrolAI) GetFiringDistance() int {
        return this.FiringDistance
}


func (this *PacePatrolAI) SetFiringDistance(firingDistance int) {
        this.FiringDistance = firingDistance
}

