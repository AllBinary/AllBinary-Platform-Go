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

type DiveAndDirectionalTrackingAI struct {
    BasicAI
    MathUtil *MathUtil
    DirectionFactory *DirectionFactory
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    TimeDelayHelper ITimeDelayHelper
    List IBasicArrayList
    DirectionalInterface DirectionalInterface
    VelocityInterface IBasicVelocityProperties
    DirectionToKeyMap []int
    InitialDropped bool
    DiveP bool
    DirectionOfTarget *Direction
    MIN_DISTANCE int
    LastTrackingLayerInterface IAllBinaryLayer
    ArtificialIntelligenceInterface ArtificialIntelligenceInterface
    AiVistor IVisitor
    LastDirection *Direction
}

// Implements TrackingEventListenerInterface, DestroyedEventListenerInterface


func (this *DiveAndDirectionalTrackingAI) initFields() {
    this.MathUtil = MathUtil_GetInstance()
    this.DirectionFactory = DirectionFactory_GetInstance()
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.TimeDelayHelper = NewTimeDelayHelper(500)
    this.List = NewBasicArrayListD()
    
    
    this.DirectionToKeyMap = []int{Canvas_LEFT, Canvas_DOWN, Canvas_UP, Canvas_RIGHT}
    
    
    this.DirectionOfTarget = Direction_GetInstance(0)
    this.MIN_DISTANCE = 40
    this.LastTrackingLayerInterface = AllBinaryLayer_NULL_ALLBINARY_LAYER
    
    
    this.LastDirection = DirectionFactory_GetInstance().NO_DIRECTION
}


func NewDiveAndDirectionalTrackingAI(ownerLayerInterface IAllBinaryLayer, artificialIntelligenceInterface ArtificialIntelligenceInterface, gameInput IGameInput, visitor IVisitor) *DiveAndDirectionalTrackingAI {
    this := &DiveAndDirectionalTrackingAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.AiVistor = visitor
        this.ArtificialIntelligenceInterface = artificialIntelligenceInterface
        var directionalCompositeInterface DirectionalCompositeInterface = any(this.GetOwnerLayerInterface()).(DirectionalCompositeInterface); _ = directionalCompositeInterface
        this.DirectionalInterface = directionalCompositeInterface.GetDirectionalInterface()
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(VelocityInterfaceCompositeInterface); _ = velocityInterfaceCompositeInterface
        this.VelocityInterface = velocityInterfaceCompositeInterface.GetVelocityProperties()
        DestroyedEventHandler_GetInstance().AddListenerInterface(this)
        this.Init()
    return this
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.IsBeyondTarget() {
            this.Init()
            this.VelocityInterface.Zero()
        }
        if !this.DiveP {
            this.Target(allBinaryLayerManager)
        } else {
            this.Dive()
            this.Attack()
        }
        this.List.Clear()
}


func (this *DiveAndDirectionalTrackingAI) Init() {
        this.DiveP = false
        this.DirectionOfTarget = DirectionFactory_GetInstance().NOT_BORDERED_WITH
        TrackingEventHandler_GetInstance().AddListenerInterface(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) Target(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.List.Size() == 0 {
            return
        }
        var lastTrackingEvent *TrackingEvent = any(this.List.RemoveAt(0)).(*TrackingEvent); _ = lastTrackingEvent
        this.LastTrackingLayerInterface = lastTrackingEvent.GetLayerInterface()
        var layerInterface IAllBinaryLayer = this.LastTrackingLayerInterface; _ = layerInterface
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        if layerInterface.GetGroupInterface() == ownerLayerInterface.GetGroupInterface() {
            return
        }
        var x int = this.LastTrackingLayerInterface.GetXP(); _ = x
        var y int = this.LastTrackingLayerInterface.GetYP(); _ = y
        var yDistance int = ownerLayerInterface.GetYP() - y - ownerLayerInterface.GetHeight(); _ = yDistance
        var xDistance int = ownerLayerInterface.GetXP() - x - ownerLayerInterface.GetWidth(); _ = xDistance
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if this.MathUtil.Abs(yDistance) > this.MathUtil.Abs(xDistance) {
            if yDistance > this.MIN_DISTANCE {
                this.DirectionOfTarget = directionFactory.UP
            } else if yDistance < -this.MIN_DISTANCE {
                this.DirectionOfTarget = directionFactory.DOWN
            }
        } else {
            if xDistance > this.MIN_DISTANCE {
                this.DirectionOfTarget = directionFactory.LEFT
            } else if xDistance < -this.MIN_DISTANCE {
                this.DirectionOfTarget = directionFactory.RIGHT
            }
        }
        if !this.InitialDropped && this.DirectionOfTarget != directionFactory.NOT_BORDERED_WITH {
            this.Drop()
        }
        if this.InitialDropped {
            if this.DirectionOfTarget == directionFactory.UP || this.DirectionOfTarget == directionFactory.DOWN {
                this.VerticalTargeting(x)
            } else if this.DirectionOfTarget == directionFactory.LEFT || this.DirectionOfTarget == directionFactory.RIGHT {
                this.HorizontalTargeting(y)
            } else {
                this.ArtificialIntelligenceInterface.ProcessAI(allBinaryLayerManager)
            }
        }
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) VerticalTargeting(x int) {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        if ownerLayerInterface.GetXP() < x {
            this.MoveRight()
        }
        if ownerLayerInterface.GetXP() > x {
            this.MoveLeft()
        }
        if ownerLayerInterface.GetXP() > x - 3 && ownerLayerInterface.GetXP() < x + 3 {
            this.SetDive()
        }
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) HorizontalTargeting(y int) {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        if ownerLayerInterface.GetYP() < y {
            this.MoveDown()
        }
        if ownerLayerInterface.GetYP() > y {
            this.MoveUp()
        }
        if ownerLayerInterface.GetYP() > y - 3 && ownerLayerInterface.GetYP() < y + 3 {
            this.SetDive()
        }
}


func (this *DiveAndDirectionalTrackingAI) SetDive() {
        this.DiveP = true
        this.VelocityInterface.Zero()
        TrackingEventHandler_GetInstance().RemoveListener(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) MoveRight() {
        this.SetLastDirection(this.DirectionFactory.RIGHT)
        this.DirectionalInterface.SetFrameByDirection(this.LastDirection)
        this.AiVistor.Visit(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) MoveLeft() {
        this.SetLastDirection(this.DirectionFactory.LEFT)
        this.DirectionalInterface.SetFrameByDirection(this.LastDirection)
        this.AiVistor.Visit(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) MoveDown() {
        this.SetLastDirection(this.DirectionFactory.DOWN)
        this.DirectionalInterface.SetFrameByDirection(this.LastDirection)
        this.AiVistor.Visit(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) MoveUp() {
        this.SetLastDirection(this.DirectionFactory.UP)
        this.DirectionalInterface.SetFrameByDirection(this.LastDirection)
        this.AiVistor.Visit(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) Dive() {
        this.DirectionalInterface.SetFrameByDirection(this.DirectionOfTarget)
        this.AiVistor.Visit(this)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) Attack() {
        this.BasicAI.ProcessKeyAI(Canvas_KEY_NUM1)
}


// throws Exception
func (this *DiveAndDirectionalTrackingAI) Drop() {
        if this.TimeDelayHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            this.InitialDropped = true
            this.AiVistor.Visit(this)
        }
}


func (this *DiveAndDirectionalTrackingAI) IsBeyondTarget() bool {
        if this.DirectionOfTarget == this.DirectionFactory.DOWN {
            if this.GetOwnerLayerInterface().GetYP() > this.LastTrackingLayerInterface.GetYP() + this.MIN_DISTANCE {
                return true
            } else {
                return false
            }
        } else if this.DirectionOfTarget == this.DirectionFactory.UP {
            if this.GetOwnerLayerInterface().GetYP() < this.LastTrackingLayerInterface.GetYP() + this.MIN_DISTANCE {
                return true
            } else {
                return false
            }
        } else if this.DirectionOfTarget == this.DirectionFactory.RIGHT {
            if this.GetOwnerLayerInterface().GetXP() > this.LastTrackingLayerInterface.GetXP() + this.MIN_DISTANCE {
                return true
            } else {
                return false
            }
        } else if this.DirectionOfTarget == this.DirectionFactory.LEFT {
            if this.GetOwnerLayerInterface().GetXP() < this.LastTrackingLayerInterface.GetXP() + this.MIN_DISTANCE {
                return true
            } else {
                return false
            }
        } else {
            return false
        }
}


func (this *DiveAndDirectionalTrackingAI) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *DiveAndDirectionalTrackingAI) OnMovement(trackingEvent *TrackingEvent) {
        this.List.Add(trackingEvent)
}


func (this *DiveAndDirectionalTrackingAI) OnDestroyed(destroyedEvent *DestroyedEvent) {
        if this.GetOwnerLayerInterface() == destroyedEvent.GetLayerInterface() {
            TrackingEventHandler_GetInstance().RemoveListener(this)
            DestroyedEventHandler_GetInstance().RemoveListener(this)
        }
}


func (this *DiveAndDirectionalTrackingAI) SetLastDirection(lastDirection *Direction) {
        this.LastDirection = lastDirection
        var value int = this.GetLastDirection().GetValue(); _ = value
        if value < 4 {
            this.BasicAI.SetLastKey(this.DirectionToKeyMap[value])
        }
}


func (this *DiveAndDirectionalTrackingAI) SetLastKey(lastKey int) {
        this.BasicAI.SetLastKey(lastKey)
        if this.GetLastKey() == Canvas_LEFT {
            this.LastDirection = DirectionFactory_GetInstance().LEFT
        } else if this.GetLastKey() == Canvas_DOWN {
            this.LastDirection = DirectionFactory_GetInstance().DOWN
        } else if this.GetLastKey() == Canvas_UP {
            this.LastDirection = DirectionFactory_GetInstance().UP
        } else if this.GetLastKey() == Canvas_RIGHT {
            this.LastDirection = DirectionFactory_GetInstance().RIGHT
        }
}


func (this *DiveAndDirectionalTrackingAI) GetLastDirection() *Direction {
        return this.LastDirection
}


func (this *DiveAndDirectionalTrackingAI) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.BasicAI.ToString())
        stringBuffer.Append(" LastDirection: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.GetLastDirection()))
        return fmt.Sprint(stringBuffer)
}

func (this *DiveAndDirectionalTrackingAI) String() string {
    return this.ToString()
}

