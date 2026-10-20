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

type BasePatrolAI struct {
    BasicAI
    MaxDistance int
    CurrentDistance int
    XTotalDistance int
    LastKeyDirection int
    KeyDirection int
}

// The Java class BasePatrolAI has subclasses so its variables are this interface
type IBasePatrolAI interface {
    IBasicAI
    BasePatrolAI_() *BasePatrolAI
    Update()
    UpdateTotalDistance()
    ChangeDirectionIfReachedPacingAreaMax()
    NextDirection()
    Disable()
}

func (this *BasePatrolAI) BasePatrolAI_() *BasePatrolAI {
    return this
}

var BasePatrolAI_MAX_DISTANCE int = SmallIntegerSingletonFactory_GetInstance().GetAt(2)
var BasePatrolAI_DIRECTION_ARRAY int = SmallIntegerSingletonFactory_GetInstance().GetAt(3)
var BasePatrolAI_FIRING_DISTANCE int = SmallIntegerSingletonFactory_GetInstance().GetAt(4)
var BasePatrolAI_SPEED int = SmallIntegerSingletonFactory_GetInstance().GetAt(5)

func (this *BasePatrolAI) initFields() {
    
    
    
    this.LastKeyDirection = Canvas_RIGHT
    this.KeyDirection = Canvas_RIGHT
}


func NewBasePatrolAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *BasePatrolAI {
    this := &BasePatrolAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.MaxDistance = any(hashtable.Get(BasePatrolAI_MAX_DISTANCE)).(int)
        this.CurrentDistance = int(this.MaxDistance)
    return this
}


func (this *BasePatrolAI) Update() {
        this.UpdateTotalDistance()
        this.ChangeDirectionIfReachedPacingAreaMax()
}


func (this *BasePatrolAI) UpdateTotalDistance() {
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(VelocityInterfaceCompositeInterface); _ = velocityInterfaceCompositeInterface
        var velocityProperties IBasicVelocityProperties = velocityInterfaceCompositeInterface.GetVelocityProperties(); _ = velocityProperties
        this.XTotalDistance += velocityProperties.GetVelocityXBasicDecimalP().GetScaled()
}


func (this *BasePatrolAI) ChangeDirectionIfReachedPacingAreaMax() {
        var mathUtil *MathUtil = MathUtil_GetInstance(); _ = mathUtil
        var totalDistance int = mathUtil.Abs(this.XTotalDistance); _ = totalDistance
        if totalDistance > this.CurrentDistance {
            this.NextDirection()
        }
}


func (this *BasePatrolAI) NextDirection() {
        if this.LastKeyDirection == Canvas_LEFT {
            this.KeyDirection = Canvas_RIGHT
            this.LastKeyDirection = this.KeyDirection
            this.XTotalDistance = 0
        } else if this.LastKeyDirection == Canvas_RIGHT {
            this.KeyDirection = Canvas_LEFT
            this.LastKeyDirection = this.KeyDirection
            this.XTotalDistance = 0
        }
}


func (this *BasePatrolAI) Disable() {
}

