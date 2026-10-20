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

type GameNotificationHud struct {
    BasicHud
    METHOD_NAME string
    LastGameNotificationEvent *GameNotificationEvent
}

// Implements GameNotificationListenerInterface

// The Java class GameNotificationHud has subclasses so its variables are this interface
type IGameNotificationHud interface {
    IBasicHud
    GameNotificationHud_() *GameNotificationHud
    OnEvent(eventObject IAllBinaryEventObject)
    OnGameNotificationEvent(gameNotificationEvent *GameNotificationEvent)
    Add(string_ string, seconds int, basicColor *BasicColor, permanent bool)
    ProcessTick()
    Paint(graphics IGraphics)
    Clear()
}

func (this *GameNotificationHud) GameNotificationHud_() *GameNotificationHud {
    return this
}

var GameNotificationHud_NULL_GAME_NOTIFICATION IGameNotificationHud = NewGameNotificationHud(BasicHudFactory_GetInstance().TOPCENTER, BasicHudFactory_GetInstance().HORIZONTAL, 0, BasicColorFactory_GetInstance().RED)

func (this *GameNotificationHud) initFields() {
    this.METHOD_NAME = "onGameNotificationEvent"
    this.LastGameNotificationEvent = GameNotification_NULL_GAME_NOTIFICATION_EVENT
}


func NewGameNotificationHud(location int, direction int, bufferZone int, basicColor *BasicColor) *GameNotificationHud {
    this := &GameNotificationHud{}
        this.BasicHud = *NewBasicHud(location, direction, bufferZone, basicColor)
        this.initFields()
    return this
}


func (this *GameNotificationHud) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *GameNotificationHud) OnGameNotificationEvent(gameNotificationEvent *GameNotificationEvent) {
        if this.LastGameNotificationEvent != gameNotificationEvent {
            this.LastGameNotificationEvent = gameNotificationEvent
            this.LogUtil.PutF(gameNotificationEvent.GetString(), this, this.METHOD_NAME)
        }
        this.Add(gameNotificationEvent.GetString(), gameNotificationEvent.GetSeconds(), gameNotificationEvent.GetBasicColorP(), gameNotificationEvent.GetPermanent())
}


func (this *GameNotificationHud) Add(string_ string, seconds int, basicColor *BasicColor, permanent bool) {
}


// throws Exception
func (this *GameNotificationHud) ProcessTick() {
}


func (this *GameNotificationHud) Paint(graphics IGraphics) {
}


func (this *GameNotificationHud) Clear() {
}

