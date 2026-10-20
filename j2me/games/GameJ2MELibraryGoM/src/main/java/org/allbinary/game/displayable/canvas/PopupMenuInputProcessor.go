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

type PopupMenuInputProcessor struct {
    BasicMenuInputProcessor
    RectangleCollisionUtil *RectangleCollisionUtil
    CLICK_DELAY int
    ClickTimeHelper ITimeDelayHelper
    Rectangle *Rectangle
}


func (this *PopupMenuInputProcessor) initFields() {
    this.RectangleCollisionUtil = RectangleCollisionUtil_GetInstance()
    this.CLICK_DELAY = 120
    this.ClickTimeHelper = NewTimeDelayHelper(this.CLICK_DELAY)
    
}


func NewPopupMenuInputProcessor(gameKeyEventList IBasicArrayList, playerInputId int, gameCanvas IMyCanvas, rectangle *Rectangle) *PopupMenuInputProcessor {
    this := &PopupMenuInputProcessor{}
        this.BasicMenuInputProcessor = *NewBasicMenuInputProcessor(gameKeyEventList, playerInputId, gameCanvas)
        this.initFields()
        this.Rectangle = rectangle
    return this
}


func (this *PopupMenuInputProcessor) Init(rectangle *Rectangle) {
        this.Rectangle = rectangle
}


// throws Exception
func (this *PopupMenuInputProcessor) ProcessInput(key int) int {
        if key == Canvas_KEY_STAR {
            var gameCanvas IAllBinaryGameCanvas = any(this.GetCanvas()).(IAllBinaryGameCanvas); _ = gameCanvas
            gameCanvas.ToggleMenu()
            return 1
        }
        return 0
}


// throws Exception
func (this *PopupMenuInputProcessor) ProcessInputList() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, GameInputStrings_GetInstance().PROCESS_INPUT, e)
            tryControl1, tryResult1 = 1, -1; return
                }
            }()
            var motionInputsIndex int = this.ProcessMotionInputs(); _ = motionInputsIndex
            var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
            var size int = list.Size(); _ = size
            var key int = 0; _ = key
            var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
            for index := int(0); index < size; index++ {
                gameKeyEvent = any(list.BasicArrayList_().ObjectArray[index]).(*GameKeyEvent)
                if gameKeyEvent != nil {
                    key = gameKeyEvent.GetKey()
                    if this.ProcessInput(key) == 1 {
                        break
                    }
                }
            }
            this.Clear()
            if size > 0 || motionInputsIndex >= 0 {
                tryControl1, tryResult1 = 1, 1; return
            } else {
                tryControl1, tryResult1 = 1, -1; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *PopupMenuInputProcessor) ProcessMotionInputs() int {
        var lastIndex int = this.MotionGestureEventList.Size() - 1; _ = lastIndex
        if lastIndex >= 0 {
            var motionGestureEvent *MotionGestureEvent = any(this.MotionGestureEventList.BasicArrayList_().ObjectArray[lastIndex]).(*MotionGestureEvent); _ = motionGestureEvent
            this.ProcessMotionInput(motionGestureEvent)
        }
        this.MotionGestureEventList.Clear()
        return lastIndex
}


// throws Exception
func (this *PopupMenuInputProcessor) ProcessMotionInput(motionGestureEvent *MotionGestureEvent) {
        if motionGestureEvent == nil {
            this.LogUtil.PutF("Exception: Bug", this, this.GameInputStrings.PROCESS_MOTION_INPUT)
            return
        }
        var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
        var motionGestureInput *MotionGestureInput = motionGestureEvent.GetMotionGesture(); _ = motionGestureInput
        if motionGestureInput == touchMotionGestureFactory.RELEASED {
            var point IGPoint = motionGestureEvent.GetCurrentPoint(); _ = point
            var rectPoint IGPoint = this.Rectangle.GetPoint(); _ = rectPoint
            if this.RectangleCollisionUtil.IsInside(rectPoint.GetX(), rectPoint.GetY(), this.Rectangle.GetMaxX() + 20, this.Rectangle.GetMaxY(), point.GetX(), point.GetY()) {
                if this.ClickTimeHelper.IsTimeTNT() {
                    var gameCanvas IAllBinaryGameCanvas = any(this.GetCanvas()).(IAllBinaryGameCanvas); _ = gameCanvas
                    gameCanvas.ToggleMenu()
                }
            }
        }
}

