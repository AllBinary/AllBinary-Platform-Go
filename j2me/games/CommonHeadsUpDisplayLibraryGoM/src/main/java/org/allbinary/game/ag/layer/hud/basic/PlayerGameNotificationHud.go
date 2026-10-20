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

type PlayerGameNotificationHud struct {
    GameNotificationHud
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    EMPTY_STRING string
    DisplayInfo *DisplayInfoSingleton
    TimeDelayHelper ITimeDelayHelper
    CircularIndexUtil *CircularIndexUtil
    GameNotification *GameNotification
    PermanentGameNotification *GameNotification
    SetAndRemoveProcessor IProcessor
    NextUnremoveableProcessor IProcessor
    Processor IProcessor
    PERMANENT_GAME_NOTIFICATION string
    LastString string
    String string
    Point *CustomGPoint
    Width int
}

// Implements UpdateMyFontInterface


func (this *PlayerGameNotificationHud) initFields() {
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
    this.DisplayInfo = DisplayInfoSingleton_GetInstance()
    this.TimeDelayHelper = NewTimeDelayHelper(0)
    
    this.GameNotification = NewGameNotification()
    this.PermanentGameNotification = NewGameNotification()
    this.SetAndRemoveProcessor = NewSetAndRemoveProcessor(this)
    this.NextUnremoveableProcessor = NewNextUnremoveableProcessor(this)
    this.Processor = Processor_GetInstance()
    this.PERMANENT_GAME_NOTIFICATION = "Permanent Game Notification: "
    this.LastString = StringUtil_GetInstance().EMPTY_STRING
    this.String = this.EMPTY_STRING
    this.Point = CustomGPoint_NULL_CUSTOM_POINT
    
}


func NewPlayerGameNotificationHud(location int, direction int, bufferZone int, basicColor *BasicColor) *PlayerGameNotificationHud {
    this := &PlayerGameNotificationHud{}
        this.GameNotificationHud = *NewGameNotificationHud(location, direction, bufferZone, basicColor)
        this.initFields()
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(0)
        var gameNotificationEventHandler *GameNotificationEventHandler = GameNotificationEventHandler_GetInstance(); _ = gameNotificationEventHandler
        gameNotificationEventHandler.RemoveAllListeners()
        gameNotificationEventHandler.AddListenerInterface(this)
    return this
}


func (this *PlayerGameNotificationHud) UpdateMeasurement(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.UPDATE, e)
                }
            }()
            this.GameNotificationHud.UpdateMeasurement(graphics)
            var font Font = graphics.GetFont(); _ = font
            this.Width = font.StringWidth(this.String)
            this.Processor.Process()
            this.Processor = Processor_GetInstance()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *PlayerGameNotificationHud) GetPoint(x int, y int) IGPoint {
        this.Point = CustomGPoint_GetInstance3(0, 0)
        this.Point.SetX(x)
        this.Point.SetY(y)
        return this.Point
}


func (this *PlayerGameNotificationHud) Add(string_ string, seconds int, basicColor *BasicColor, permanent bool) {
        if permanent {
            if this.LastString != string_ {
                this.LastString = string_
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.PERMANENT_GAME_NOTIFICATION).Append(string_)), this, this.CommonStrings.ADD)
            }
            this.PermanentGameNotification.Add(string_, seconds, basicColor)
            this.CircularIndexUtil.SetSize(this.PermanentGameNotification.GetSize())
        } else {
            this.GameNotification.Add(string_, seconds, basicColor)
        }
}


// throws Exception
func (this *PlayerGameNotificationHud) ProcessTick() {
        if this.TimeDelayHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            var gameAdState *GameAdState = GameAdStateFactory_GetInstance().GetCurrentInstance(); _ = gameAdState
            if gameAdState.IsShowingAtLocation(this.GetLocation()) {
                this.OffsetY = -54
            } else {
                this.OffsetY = 0
            }
            if this.GameNotification.GetSize() > 0 {
                this.SetAndRemove()
            } else if this.PermanentGameNotification.GetSize() > 0 {
                this.SetNextUnremoveable()
            } else {
                this.String = this.EMPTY_STRING
            }
        }
}


// throws Exception
func (this *PlayerGameNotificationHud) SetAndRemove() {
        this.String = any(this.GameNotification.StringList.RemoveAt(0)).(string)
        this.Processor = this.SetAndRemoveProcessor
}


// throws Exception
func (this *PlayerGameNotificationHud) SetAndRemoveProcess() {
        this.SetX((this.DisplayInfo.GetLastWidth() - this.Width) >> 1)
        this.Point.SetX(this.GetX())
        this.Point.SetY(this.GetY())
        var time int = any(this.GameNotification.TimeList.RemoveAt(0)).(int); _ = time
        var iTime int = int(time) * 1000; _ = iTime
        if iTime == 0 {
            iTime = 500
        }
        this.TimeDelayHelper.TimeDelayHelper_().Delay = iTime
        this.SetBasicColorP(any(this.GameNotification.ColorList.RemoveAt(0)).(*BasicColor))
}


// throws Exception
func (this *PlayerGameNotificationHud) SetNextUnremoveable() {
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        this.String = any(this.PermanentGameNotification.StringList.BasicArrayList_().ObjectArray[index]).(string)
        this.Processor = this.NextUnremoveableProcessor
}


// throws Exception
func (this *PlayerGameNotificationHud) SetNextUnremoveableProcess() {
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        this.SetX((this.DisplayInfo.GetLastWidth() - this.Width) >> 1)
        this.Point.SetX(this.GetX())
        this.Point.SetY(this.GetY())
        var time int = any(this.PermanentGameNotification.TimeList.BasicArrayList_().ObjectArray[index]).(int); _ = time
        this.TimeDelayHelper.TimeDelayHelper_().Delay = int(time) * 1000
        this.SetBasicColorP(any(this.PermanentGameNotification.ColorList.BasicArrayList_().ObjectArray[index]).(*BasicColor))
        this.CircularIndexUtil.Next()
}


func (this *PlayerGameNotificationHud) Clear() {
        this.GameNotification.Clear()
        this.PermanentGameNotification.Clear()
}


func (this *PlayerGameNotificationHud) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        this.GameNotificationHud.PaintString(graphics, this.String)
}

type SetAndRemoveProcessor struct {
    Processor
    UpdateMeasurementProcessor *PlayerGameNotificationHud
}


func (this *SetAndRemoveProcessor) initFields() {
        
}
    

    func NewPlayerGameNotificationHud(updateMeasurementProcessor *PlayerGameNotificationHud) *PlayerGameNotificationHud {
        this := &PlayerGameNotificationHud{}
        this.Processor = *NewProcessor()
        this.initFields()
            this.UpdateMeasurementProcessor = updateMeasurementProcessor
        return this
    }
    

    // throws Exception
    func (this *PlayerGameNotificationHud) Process() {
            this.UpdateMeasurementProcessor.SetAndRemoveProcess()
    }

type NextUnremoveableProcessor struct {
    Processor
    UpdateMeasurementProcessor *PlayerGameNotificationHud
}


func (this *NextUnremoveableProcessor) initFields() {
        
}
    

    func NewPlayerGameNotificationHud(updateMeasurementProcessor *PlayerGameNotificationHud) *PlayerGameNotificationHud {
        this := &PlayerGameNotificationHud{}
        this.Processor = *NewProcessor()
        this.initFields()
            this.UpdateMeasurementProcessor = updateMeasurementProcessor
        return this
    }
    

    // throws Exception
    func (this *PlayerGameNotificationHud) Process() {
            this.UpdateMeasurementProcessor.SetNextUnremoveableProcess()
    }

