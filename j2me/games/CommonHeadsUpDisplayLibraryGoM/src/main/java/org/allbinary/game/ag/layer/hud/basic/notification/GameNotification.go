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

type GameNotification struct {
    StringList IBasicArrayList
    TimeList IBasicArrayList
    ColorList IBasicArrayList
}

var GameNotification_NULL_GAME_NOTIFICATION_EVENT *GameNotificationEvent = NewGameNotificationEvent(NullUtil_GetInstance().NULL_OBJECT, StringUtil_GetInstance().EMPTY_STRING, int(0), BasicColorFactory_GetInstance().RED, BooleanFactory_GetInstance().FALSE)

func (this *GameNotification) initFields() {
    this.StringList = NewBasicArrayListD()
    this.TimeList = NewBasicArrayListD()
    this.ColorList = NewBasicArrayListD()
}

func NewGameNotification() *GameNotification {
    this := &GameNotification{}
    this.initFields()
    return this
}


func (this *GameNotification) Add(string_ string, seconds int, basicColor *BasicColor) {
        if !this.StringList.Contains(string_) {
            this.StringList.Add(string_)
            this.TimeList.Add(seconds)
            this.ColorList.Add(basicColor)
        }
}


func (this *GameNotification) Clear() {
        this.StringList.Clear()
        this.TimeList.Clear()
        this.ColorList.Clear()
}


func (this *GameNotification) GetSize() int {
        return this.StringList.Size()
}

