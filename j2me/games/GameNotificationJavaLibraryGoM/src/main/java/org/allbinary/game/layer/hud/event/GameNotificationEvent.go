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

type GameNotificationEvent struct {
    AllBinaryEventObject
    String string
    Seconds int
    BasicColor *BasicColor
    Removeable bool
}


func (this *GameNotificationEvent) initFields() {
    this.String = StringUtil_GetInstance().EMPTY_STRING
    
    this.BasicColor = BasicColorFactory_GetInstance().BLACK
    
}


func NewGameNotificationEvent(anyType any, string_ string, seconds int, basicColor *BasicColor, permanent bool) *GameNotificationEvent {
    this := &GameNotificationEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(anyType)
        this.initFields()
        this.SetString(string_)
        this.Seconds = seconds
        this.SetBasicColorP(basicColor)
        this.Removeable = permanent
    return this
}


func (this *GameNotificationEvent) GetString() string {
        return this.String
}


func (this *GameNotificationEvent) GetSeconds() int {
        return this.Seconds
}


func (this *GameNotificationEvent) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *GameNotificationEvent) GetPermanent() bool {
        return this.Removeable
}


func (this *GameNotificationEvent) SetString(string_ string) {
        this.String = string_
}


func (this *GameNotificationEvent) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}

