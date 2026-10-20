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

type MirrorMotionGestureEventUtil struct {
    MotionEventCircularPool *MotionEventCircularPool
}

var MirrorMotionGestureEventUtil_Instance *MirrorMotionGestureEventUtil = NewMirrorMotionGestureEventUtil()

func (this *MirrorMotionGestureEventUtil) initFields() {
    
}


func NewMirrorMotionGestureEventUtil() *MirrorMotionGestureEventUtil {
    this := &MirrorMotionGestureEventUtil{}
    this.initFields()
        this.MotionEventCircularPool = MotionEventCircularPool_CreatePool(Integer_MAX_VALUE)
    return this
}


// throws Exception
func (this *MirrorMotionGestureEventUtil) MirrorVerticle(motionGestureEvent *MotionGestureEvent, halfWidth int) *MotionGestureEvent {
        var event *MotionGestureEvent = this.MotionEventCircularPool.GetInstance(motionGestureEvent.GetMotionGesture()); _ = event
        var previousNewX int = this.GetNewX(motionGestureEvent.GetPreviousPoint().GetX(), halfWidth); _ = previousNewX
        var currentNewX int = this.GetNewX(motionGestureEvent.GetCurrentPoint().GetX(), halfWidth); _ = currentNewX
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var mirroredPreviousPoint IGPoint = pointFactory.CreateXY(previousNewX, motionGestureEvent.GetPreviousPoint().GetY()); _ = mirroredPreviousPoint
        var mirroredCurrentPoint IGPoint = pointFactory.CreateXY(currentNewX, motionGestureEvent.GetCurrentPoint().GetY()); _ = mirroredCurrentPoint
        event.SetPreviousPoint(mirroredPreviousPoint)
        event.SetCurrentPoint(mirroredCurrentPoint)
        return event
}


func (this *MirrorMotionGestureEventUtil) GetNewX(x int, halfWidth int) int {
        var newX int; _ = newX
        if x > halfWidth {
            newX = halfWidth - (x - halfWidth)
        } else {
            newX = halfWidth + (halfWidth - x)
        }
        return newX
}


func MirrorMotionGestureEventUtil_GetInstance() *MirrorMotionGestureEventUtil {
        return MirrorMotionGestureEventUtil_Instance
}

