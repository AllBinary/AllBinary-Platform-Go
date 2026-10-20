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

type CompleteMotionGestureInputToGameMotionGestureInput struct {
    LogUtil *LogUtil
}

var CompleteMotionGestureInputToGameMotionGestureInput_Instance *CompleteMotionGestureInputToGameMotionGestureInput = NewCompleteMotionGestureInputToGameMotionGestureInput()

func (this *CompleteMotionGestureInputToGameMotionGestureInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewCompleteMotionGestureInputToGameMotionGestureInput() *CompleteMotionGestureInputToGameMotionGestureInput {
    this := &CompleteMotionGestureInputToGameMotionGestureInput{}
    this.initFields()
    return this
}


func CompleteMotionGestureInputToGameMotionGestureInput_GetInstance() *CompleteMotionGestureInputToGameMotionGestureInput {
        return CompleteMotionGestureInputToGameMotionGestureInput_Instance
}


func (this *CompleteMotionGestureInputToGameMotionGestureInput) Init() {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.INIT, e)
                }
            }()
            PreLogUtil_Put("Compound Motion Gestures", this, commonStrings.INIT)
            var motionGestureConfiguration *MotionGestureConfiguration = MotionGestureConfigurationFactory_GetInstance(); _ = motionGestureConfiguration
            var touchMotionGestureFactory *TouchMotionGestureFactory = TouchMotionGestureFactory_GetInstance(); _ = touchMotionGestureFactory
            var list IBasicArrayList = NewBasicArrayListD(); _ = list
            list.Add(touchMotionGestureFactory.DIAGONAL_DOWN_LEFT)
            motionGestureConfiguration.AddMotionGestureAction(list, DiagonalDownLeftTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.DIAGONAL_DOWN_RIGHT)
            motionGestureConfiguration.AddMotionGestureAction(list, DiagonalDownRightTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.DIAGONAL_UP_LEFT)
            motionGestureConfiguration.AddMotionGestureAction(list, DiagonalUpLeftTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.DIAGONAL_UP_RIGHT)
            motionGestureConfiguration.AddMotionGestureAction(list, DiagonalUpRightTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.LEFT)
            motionGestureConfiguration.AddMotionGestureAction(list, LeftTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.RIGHT)
            motionGestureConfiguration.AddMotionGestureAction(list, RightTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.DOWN)
            motionGestureConfiguration.AddMotionGestureAction(list, DownTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.UP)
            motionGestureConfiguration.AddMotionGestureAction(list, UpTouchInputToGameKeyEventAction_GetInstance())
            list = NewBasicArrayListD()
            list.Add(touchMotionGestureFactory.RELEASED)
            motionGestureConfiguration.AddMotionGestureAction(list, ReleaseTouchInputToGameKeyEventAction_GetInstance())
            return
        }(); tryControl1 == 1 {
            return
        }
}

