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

type TextNotificationUtil struct {
    TWO int
    FALSE bool
    GameNotificationEventHandler *GameNotificationEventHandler
}

var TextNotificationUtil_Instance *TextNotificationUtil = NewTextNotificationUtil()

func (this *TextNotificationUtil) initFields() {
    this.TWO = SmallIntegerSingletonFactory_GetInstance().GetAt(2)
    this.FALSE = BooleanFactory_GetInstance().FALSE
    this.GameNotificationEventHandler = GameNotificationEventHandler_GetInstance()
}

func NewTextNotificationUtil() *TextNotificationUtil {
    this := &TextNotificationUtil{}
    this.initFields()
    return this
}


func TextNotificationUtil_GetInstance() *TextNotificationUtil {
        return TextNotificationUtil_Instance
}


// throws Exception
func (this *TextNotificationUtil) FireError(message string) {
        PrimaryPlayerQueueFactory_GetInstance().Add(ErrorSound_GetInstance())
        this.Fire(message)
}


// throws Exception
func (this *TextNotificationUtil) FireNew(message string) {
        PrimaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
        this.Fire(message)
}


// throws Exception
func (this *TextNotificationUtil) FireSuccess(message string) {
        PrimaryPlayerQueueFactory_GetInstance().Add(SelectSound_GetInstance())
        this.Fire(message)
}


// throws Exception
func (this *TextNotificationUtil) Fire(message string) {
        var gameNotificationEvent *GameNotificationEvent = NewGameNotificationEvent(this, message, this.TWO, BasicColorFactory_GetInstance().RED, this.FALSE); _ = gameNotificationEvent
        this.GameNotificationEventHandler.FireEvent(gameNotificationEvent)
}

