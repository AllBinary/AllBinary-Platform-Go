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

type DirectionalAnalogLocationInputProcessor struct {
    AnalogLocationInputProcessor
    LogUtil *LogUtil
    InputProcessorArray []IGameInputProcessor
    GameKeyFactory *GameKeyFactory
    LeftGameKeyEvent *GameKeyEvent
    RightGameKeyEvent *GameKeyEvent
    UpGameKeyEvent *GameKeyEvent
    DownGameKeyEvent *GameKeyEvent
    LeftTriggerGameKeyEvent *GameKeyEvent
    RightTriggerGameKeyEvent *GameKeyEvent
}

// Implements GameKeyEventSourceInterface


func (this *DirectionalAnalogLocationInputProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.GameKeyFactory = GameKeyFactory_GetInstance()
    this.LeftGameKeyEvent = GameKeyEvent_NONE
    this.RightGameKeyEvent = GameKeyEvent_NONE
    this.UpGameKeyEvent = GameKeyEvent_NONE
    this.DownGameKeyEvent = GameKeyEvent_NONE
    this.LeftTriggerGameKeyEvent = GameKeyEvent_NONE
    this.RightTriggerGameKeyEvent = GameKeyEvent_NONE
}


func NewDirectionalAnalogLocationInputProcessor(inputProcessorArray []IGameInputProcessor) *DirectionalAnalogLocationInputProcessor {
    this := &DirectionalAnalogLocationInputProcessor{}
    this.AnalogLocationInputProcessor = *NewAnalogLocationInputProcessor()
    this.initFields()
        this.InputProcessorArray = inputProcessorArray
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            this.LeftGameKeyEvent = GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.GameKeyFactory.LEFT)
            this.RightGameKeyEvent = GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.GameKeyFactory.RIGHT)
            this.UpGameKeyEvent = GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.GameKeyFactory.UP)
            this.DownGameKeyEvent = GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.GameKeyFactory.DOWN)
            this.LeftTriggerGameKeyEvent = GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.GameKeyFactory.KEY_NUM0)
            this.RightTriggerGameKeyEvent = GameKeyEventFactory_GetInstance().GetInstanceForInput(this, this.GameKeyFactory.KEY_NUM5)
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *DirectionalAnalogLocationInputProcessor) Process(allbinaryLayerManager IAllBinaryLayerManager, analogLocationInput *AnalogLocationInput) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put("Unable to process analog input", this, commonStrings.PROCESS, e)
                }
            }()
            var customGPoint *CustomGPoint; _ = customGPoint
            customGPoint = analogLocationInput.GetCustomGPoint()
            var x int = customGPoint.GetX(); _ = x
            var y int = customGPoint.GetY(); _ = y
            var leftTrigger int = analogLocationInput.GetLeftTrigger(); _ = leftTrigger
            var rightTrigger int = analogLocationInput.GetRightTrigger(); _ = rightTrigger
            if x < 0 {
                this.InputProcessorArray[this.LeftGameKeyEvent.GetKey()].ProcessAnalog(allbinaryLayerManager, this.LeftGameKeyEvent, x)
            } else if x > 0 {
                this.InputProcessorArray[this.RightGameKeyEvent.GetKey()].ProcessAnalog(allbinaryLayerManager, this.RightGameKeyEvent, x)
            }
            if y < 0 {
                this.InputProcessorArray[this.DownGameKeyEvent.GetKey()].ProcessAnalog(allbinaryLayerManager, this.DownGameKeyEvent, y)
            } else if y > 0 {
                this.InputProcessorArray[this.UpGameKeyEvent.GetKey()].ProcessAnalog(allbinaryLayerManager, this.UpGameKeyEvent, y)
            }
            if leftTrigger > 0 {
                this.InputProcessorArray[this.LeftTriggerGameKeyEvent.GetKey()].ProcessAnalog(allbinaryLayerManager, this.LeftTriggerGameKeyEvent, leftTrigger)
            }
            if rightTrigger > 0 {
                this.InputProcessorArray[this.RightTriggerGameKeyEvent.GetKey()].ProcessAnalog(allbinaryLayerManager, this.RightTriggerGameKeyEvent, rightTrigger)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DirectionalAnalogLocationInputProcessor) GetSourceId() int {
        return 0
}

