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

type MotionGestureRecognizer struct {
    LogUtil *LogUtil
    J2seMath *J2SEMath
    Origin IGPoint
    Previous IGPoint
    Intermediate IGPoint
    Line *Line
    MotionGesturesHandler *BasicMotionGesturesHandler
    MovedMotionGesturesHandler IBasicEventHandler
    ScrolledMotionGesturesHandler IBasicEventHandler
    MotionEventCircularPool *MotionEventCircularPool
}


func (this *MotionGestureRecognizer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.J2seMath = J2SEMath_GetInstance()
    this.Origin = PointFactory_GetInstance().ZERO_ZERO
    this.Previous = this.Origin
    this.Intermediate = this.Origin
    this.Line = NewLine(this.Origin, this.Origin)
    
    
    
    
}


func NewMotionGestureRecognizer(id int) *MotionGestureRecognizer {
    this := &MotionGestureRecognizer{}
    this.initFields()
        this.MotionEventCircularPool = MotionEventCircularPool_CreatePool(id)
        var motionGesturesHandler IBasicEventHandler = NewBasicEventHandler(); _ = motionGesturesHandler
        var movedMotionGesturesHandler IBasicEventHandler = motionGesturesHandler; _ = movedMotionGesturesHandler
        var scrolledMotionGesturesHandler IBasicEventHandler = motionGesturesHandler; _ = scrolledMotionGesturesHandler
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            motionGesturesHandler = BasicMotionGesturesHandler_GetInstance()
            movedMotionGesturesHandler = MovedMotionGesturesHandler_GetInstance()
            scrolledMotionGesturesHandler = ScrolledMotionGesturesHandler_GetInstance()
            return
        }(); tryControl1 == 1 {
            return
        }
        this.MotionGesturesHandler = any(motionGesturesHandler).(*BasicMotionGesturesHandler)
        this.MovedMotionGesturesHandler = movedMotionGesturesHandler
        this.ScrolledMotionGesturesHandler = scrolledMotionGesturesHandler
    return this
}


// throws Exception
func (this *MotionGestureRecognizer) ProcessPressedMotionEvent(current IGPoint, deviceId int, button int) bool {
        this.Intermediate = this.Origin
        this.Previous = this.Origin
        var event *MotionGestureEvent = this.MotionEventCircularPool.GetInstance(TouchMotionGestureFactory_GetInstance().PRESSED); _ = event
        event.SetPreviousPoint(this.Previous)
        event.SetCurrentPoint(current)
        this.MotionGesturesHandler.FireEvent(event)
        return true
}


// throws Exception
func (this *MotionGestureRecognizer) ProcessReleasedMotionEvent(current IGPoint, deviceId int, button int) bool {
        var event *MotionGestureEvent = this.MotionEventCircularPool.GetInstance(TouchMotionGestureFactory_GetInstance().RELEASED); _ = event
        event.SetPreviousPoint(this.Previous)
        event.SetCurrentPoint(current)
        this.MotionGesturesHandler.FireEvent(event)
        return true
}


// throws Exception
func (this *MotionGestureRecognizer) ProcessDraggedMotionEvent(current IGPoint, deviceId int, buttonMask int) {
        if this.Previous == this.Origin || this.Intermediate == this.Origin {
            this.Previous = current
            this.Intermediate = current
            return
        }
        this.Line.SetP1(this.Previous)
        this.Line.SetP2(current)
        var minimumMotionGesture int = MotionGestureConfigurationFactory_GetInstance().GetMinimumMotionGesture(); _ = minimumMotionGesture
        if this.J2seMath.Abs(float32(this.Line.GetDeltaX())) < float32(minimumMotionGesture) && this.J2seMath.Abs(float32(this.Line.GetDeltaY())) < float32(minimumMotionGesture) {
            this.Intermediate = current
            return
        }
        var gradient float64 = this.Line.GetGradient(); _ = gradient
        var absGradient float64 = float64(this.J2seMath.Abs(float32(gradient))); _ = absGradient
        var conf *MotionGestureConfiguration = MotionGestureConfigurationFactory_GetInstance(); _ = conf
        var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
        var newMotionGesture *MotionGestureInput = touchMotionGestureFactory.NO_MOTION; _ = newMotionGesture
        var diagonalToleranceHigher float64 = 12.0; _ = diagonalToleranceHigher
        var diagonalToleranceLower float64 = 12.0; _ = diagonalToleranceLower
        if conf.IsDiagonalMotionGestureAllowed() {
            diagonalToleranceHigher = float64((90 - conf.GetDiagonalTolerance()))
            diagonalToleranceLower = float64(conf.GetDiagonalTolerance())
        }
        if absGradient > math.Tan((float64(diagonalToleranceHigher) * math.Pi / 180)) {
            if this.Line.GetDeltaY() > 0.0 {
                newMotionGesture = touchMotionGestureFactory.UP
            } else {
                newMotionGesture = touchMotionGestureFactory.DOWN
            }
        } else {
            if absGradient < math.Tan((float64(diagonalToleranceLower) * math.Pi / 180)) {
                if this.Line.GetDeltaX() > 0.0 {
                    newMotionGesture = touchMotionGestureFactory.LEFT
                } else {
                    newMotionGesture = touchMotionGestureFactory.RIGHT
                }
            } else {
                if gradient > 0.0 {
                    if this.Line.GetDeltaX() > 0.0 {
                        newMotionGesture = touchMotionGestureFactory.DIAGONAL_UP_LEFT
                    } else {
                        newMotionGesture = touchMotionGestureFactory.DIAGONAL_DOWN_RIGHT
                    }
                } else {
                    if this.Line.GetDeltaX() > 0.0 {
                        newMotionGesture = touchMotionGestureFactory.DIAGONAL_DOWN_LEFT
                    } else {
                        newMotionGesture = touchMotionGestureFactory.DIAGONAL_UP_RIGHT
                    }
                }
            }
        }
        this.Previous = current
        this.Intermediate = current
        var event *MotionGestureEvent = this.MotionEventCircularPool.GetInstance(newMotionGesture); _ = event
        event.SetPreviousPoint(this.Previous)
        event.SetCurrentPoint(current)
        this.MotionGesturesHandler.FireEvent(event)
        return
}


// throws Exception
func (this *MotionGestureRecognizer) ProcessMovedMotionEvent(current IGPoint, deviceId int, button int) bool {
        var event *MotionGestureEvent = this.MotionEventCircularPool.GetInstance(TouchMotionGestureFactory_GetInstance().NO_MOTION); _ = event
        event.SetPreviousPoint(this.Previous)
        event.SetCurrentPoint(current)
        this.MovedMotionGesturesHandler.FireEvent(event)
        return true
}


// throws Exception
func (this *MotionGestureRecognizer) ProcessScrolledMotionEvent(current IGPoint, deviceId int, button int) bool {
        var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
        var newMotionGesture *MotionGestureInput = touchMotionGestureFactory.NO_MOTION; _ = newMotionGesture
        if button > 0 {
            newMotionGesture = touchMotionGestureFactory.SCROLL_UP
        } else if button < 0 {
            newMotionGesture = touchMotionGestureFactory.SCROLL_DOWN
        }
        var event *MotionGestureEvent = this.MotionEventCircularPool.GetInstance(newMotionGesture); _ = event
        event.SetPreviousPoint(this.Previous)
        event.SetCurrentPoint(current)
        this.ScrolledMotionGesturesHandler.FireEvent(event)
        return true
}


func (this *MotionGestureRecognizer) GetMotionGesturesHandler() *BasicMotionGesturesHandler {
        return this.MotionGesturesHandler
}

