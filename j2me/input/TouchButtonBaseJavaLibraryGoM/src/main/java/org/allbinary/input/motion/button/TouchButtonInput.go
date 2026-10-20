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

type TouchButtonInput struct {
    Input
    LogUtil *LogUtil
    GameKeyEventFactory *GameKeyEventFactory
    TOUCH_BUTTON_SOURCE_ID int
    GameKeyEvent *GameKeyEvent
}

// Implements GameKeyEventSourceInterface


func (this *TouchButtonInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    this.TOUCH_BUTTON_SOURCE_ID = this.GameKeyEventFactory.TOUCH_BUTTON_SOURCE_ID
    this.GameKeyEvent = GameKeyEvent_NONE
}


func NewTouchButtonInput(id int, name string) *TouchButtonInput {
    this := &TouchButtonInput{}
        this.Input = *NewInput(id, name)
        this.initFields()
        var inputFactory *InputFactory = InputFactory_GetInstance(); _ = inputFactory
        inputFactory.Add(this.GetId(), this)
    return this
}


func (this *TouchButtonInput) Update(inputToGameKeyMapping *InputToGameKeyMapping) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.UPDATE, e)
                }
            }()
            var gameKey *GameKey = inputToGameKeyMapping.GetInstance(this.GetId()); _ = gameKey
            this.GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this, gameKey)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *TouchButtonInput) GetSourceId() int {
        return this.TOUCH_BUTTON_SOURCE_ID
}


func (this *TouchButtonInput) GetGameKeyEvent() *GameKeyEvent {
        return this.GameKeyEvent
}

