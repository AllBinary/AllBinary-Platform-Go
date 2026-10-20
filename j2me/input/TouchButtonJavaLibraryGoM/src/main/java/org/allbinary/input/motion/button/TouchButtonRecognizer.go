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

type TouchButtonRecognizer struct {
    LogUtil *LogUtil
    RectangleCollisionUtil *RectangleCollisionUtil
    UpGameKeyEventHandler *UpGameKeyEventHandler
    DownGameKeyEventHandler *DownGameKeyEventHandler
    LastPressedTouchButtonInput *TouchButtonInput
    CurrentlyPressedTouchButtonSingleton ICurrentlyPressedTouchButtonSingleton
    ReleaseHelper IReleaseHelper
}


func (this *TouchButtonRecognizer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.RectangleCollisionUtil = RectangleCollisionUtil_GetInstance()
    this.UpGameKeyEventHandler = UpGameKeyEventHandler_GetInstance()
    this.DownGameKeyEventHandler = DownGameKeyEventHandler_GetInstance()
    this.LastPressedTouchButtonInput = BasicTouchInputFactory_GetInstance().NONE
    this.CurrentlyPressedTouchButtonSingleton = CurrentlyPressedTouchButtonSingleton_GetInstance()
    
}


func NewTouchButtonRecognizer() *TouchButtonRecognizer {
    this := &TouchButtonRecognizer{}
    this.initFields()
        if TouchScreenFactory_GetInstance().IsMultiTouch() {
            this.ReleaseHelper = NewMultitouchReleaseHelper(this)
        } else {
            this.ReleaseHelper = NewReleaseHelper(this)
        }
    return this
}


// throws Exception
func (this *TouchButtonRecognizer) ProcessRelease(touchButtonInput *TouchButtonInput, deviceId int) {
        this.ReleaseHelper.Release(touchButtonInput, deviceId)
        var gameKeyEvent *GameKeyEvent = touchButtonInput.GetGameKeyEvent(); _ = gameKeyEvent
        this.UpGameKeyEventHandler.FireEvent(gameKeyEvent)
        this.UpGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
        this.CurrentlyPressedTouchButtonSingleton.Remove(touchButtonInput)
}


// throws Exception
// synchronized
func (this *TouchButtonRecognizer) ReleaseTouchButtonInput(x int, y int, deviceId int) bool {
        var touchButtonInput *TouchButtonInput = this.LastPressedTouchButtonInput; _ = touchButtonInput
        if this.CurrentlyPressedTouchButtonSingleton.Contains(touchButtonInput) {
            this.LastPressedTouchButtonInput = BasicTouchInputFactory_GetInstance().NONE
            this.ProcessRelease(touchButtonInput, deviceId)
            return true
        } else {
            var list IBasicArrayList = TouchButtonFactory_GetInstance().GetList(); _ = list
            var rectangle *Rectangle; _ = rectangle
            var point IGPoint; _ = point
            var touchButton ITouchButton; _ = touchButton
            for index := int(list.Size() - 1); index >= 0; index-- {
                touchButton = any(list.BasicArrayList_().ObjectArray[index]).(ITouchButton)
                rectangle = touchButton.GetRectangle()
                point = rectangle.GetPoint()
                if this.RectangleCollisionUtil.IsInside(point.GetX(), point.GetY(), rectangle.GetMaxX(), rectangle.GetMaxY(), x, y) {
                    touchButtonInput = touchButton.GetTouchButtonInput()
                    this.ProcessRelease(touchButtonInput, deviceId)
                    return true
                }
            }
        }
        return false
}


// throws Exception
// synchronized
func (this *TouchButtonRecognizer) PressTouchButtonInput(x int, y int, deviceId int) bool {
        var list IBasicArrayList = TouchButtonFactory_GetInstance().GetList(); _ = list
        var rectangle *Rectangle; _ = rectangle
        var point IGPoint; _ = point
        var touchButton ITouchButton; _ = touchButton
        var touchButtonInput *TouchButtonInput; _ = touchButtonInput
        var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
        for index := int(list.Size() - 1); index >= 0; index-- {
            touchButton = any(list.BasicArrayList_().ObjectArray[index]).(ITouchButton)
            rectangle = touchButton.GetRectangle()
            point = rectangle.GetPoint()
            if this.RectangleCollisionUtil.IsInside(point.GetX(), point.GetY(), rectangle.GetMaxX(), rectangle.GetMaxY(), x, y) {
                touchButtonInput = touchButton.GetTouchButtonInput()
                if !this.CurrentlyPressedTouchButtonSingleton.Contains(touchButtonInput) {
                    this.ReleaseHelper.Release(touchButtonInput, deviceId)
                    this.LastPressedTouchButtonInput = touchButtonInput
                    this.CurrentlyPressedTouchButtonSingleton.Add(touchButtonInput)
                    gameKeyEvent = touchButtonInput.GetGameKeyEvent()
                    this.DownGameKeyEventHandler.FireEvent(gameKeyEvent)
                    this.DownGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
                }
                return true
            }
        }
        return false
}

//CurrentlyPressedTouchButtonSingletonDebug.getInstance();
type ReleaseHelper struct {
    TouchButtonRecognizer *TouchButtonRecognizer
}

// The Java class ReleaseHelper has subclasses so its variables are this interface
type IReleaseHelper interface {
    ReleaseHelper_() *ReleaseHelper
    Release(touchButtonInput *TouchButtonInput, deviceId int)
}

func (this *ReleaseHelper) ReleaseHelper_() *ReleaseHelper {
    return this
}


func (this *ReleaseHelper) initFields() {
        
}
    

    func NewTouchButtonRecognizer(touchButtonRecognizer *TouchButtonRecognizer) *TouchButtonRecognizer {
        this := &TouchButtonRecognizer{}
        this.initFields()
            this.TouchButtonRecognizer = touchButtonRecognizer
        return this
    }
    

    // throws Exception
    func (this *TouchButtonRecognizer) Release(touchButtonInput *TouchButtonInput, deviceId int) {
            var nextTouchButtonInput *TouchButtonInput; _ = nextTouchButtonInput
            var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
            for index := int(this.TouchButtonRecognizer.CurrentlyPressedTouchButtonSingleton.Size() - 1); index >= 0; index-- {
                nextTouchButtonInput = this.TouchButtonRecognizer.CurrentlyPressedTouchButtonSingleton.Get(index)
                gameKeyEvent = nextTouchButtonInput.GetGameKeyEvent()
                this.TouchButtonRecognizer.UpGameKeyEventHandler.FireEvent(gameKeyEvent)
                this.TouchButtonRecognizer.UpGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
                this.TouchButtonRecognizer.CurrentlyPressedTouchButtonSingleton.RemoveAt(index)
            }
    }

type MultitouchReleaseHelper struct {
    ReleaseHelper
}

    

    func NewTouchButtonRecognizer(touchButtonRecognizer *TouchButtonRecognizer) *TouchButtonRecognizer {
        this := &TouchButtonRecognizer{}
            
        return this
    }
    

    // throws Exception
    func (this *TouchButtonRecognizer) Release(touchButtonInput *TouchButtonInput, deviceId int) {
            var cancelTouchButtonInput *TouchButtonInput = CancelTouchButtonInputFactory_GetInstance().GetCancel(touchButtonInput); _ = cancelTouchButtonInput
            var nextTouchButtonInput *TouchButtonInput; _ = nextTouchButtonInput
            var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
            for index := int(this.TouchButtonRecognizer.CurrentlyPressedTouchButtonSingleton.Size() - 1); index >= 0; index-- {
                nextTouchButtonInput = this.TouchButtonRecognizer.CurrentlyPressedTouchButtonSingleton.Get(index)
                if cancelTouchButtonInput == nextTouchButtonInput {
                    gameKeyEvent = nextTouchButtonInput.GetGameKeyEvent()
                    this.TouchButtonRecognizer.UpGameKeyEventHandler.FireEvent(gameKeyEvent)
                    this.TouchButtonRecognizer.UpGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
                    this.TouchButtonRecognizer.CurrentlyPressedTouchButtonSingleton.RemoveAt(index)
                }
            }
    }

