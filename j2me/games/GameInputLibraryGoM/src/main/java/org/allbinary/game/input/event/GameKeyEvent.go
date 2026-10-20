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

type GameKeyEvent struct {
    AllBinaryEventObject
    SourceId int
    Repeated bool
    Key int
    GameActionKey int
    OriginalKey int
}

var GameKeyEvent_NONE *GameKeyEvent = NewGameKeyEvent(NullUtil_GetInstance().NULL_OBJECT, -1, -1, -1, -1, false)

func (this *GameKeyEvent) initFields() {
    
    
    
    
    
}


func GameKeyEvent_CreateEvent(anyType any, sourceId int, key int) *GameKeyEvent {
        return NewGameKeyEvent(anyType, sourceId, 0, 0, key, false)
}


func NewGameKeyEvent(anyType any, sourceId int, originalKey int, gameActionKey int, key int, repeated bool) *GameKeyEvent {
    this := &GameKeyEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(anyType)
        this.initFields()
        this.SourceId = sourceId
        this.SetOriginalKey(originalKey)
        this.Key = key
        this.SetGameActionKey(gameActionKey)
        this.SetRepeatEvents(repeated)
    return this
}


func (this *GameKeyEvent) Init(anyType any) {
        this.SetSource(anyType)
}


func (this *GameKeyEvent) SetRepeatEvents(repeated bool) {
        this.Repeated = repeated
}


func (this *GameKeyEvent) HasRepeatEvents() bool {
        return this.Repeated
}


func (this *GameKeyEvent) GetKey() int {
        return this.Key
}


func (this *GameKeyEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("\nKey: ")
        stringBuffer.Appendint(this.GetKey())
        return fmt.Sprint(stringBuffer)
}

func (this *GameKeyEvent) String() string {
    return this.ToString()
}


func (this *GameKeyEvent) GetOriginalKey() int {
        return this.OriginalKey
}


func (this *GameKeyEvent) SetOriginalKey(originalKey int) {
        this.OriginalKey = originalKey
}


func (this *GameKeyEvent) GetGameActionKey() int {
        return this.GameActionKey
}


func (this *GameKeyEvent) SetGameActionKey(gameActionKey int) {
        this.GameActionKey = gameActionKey
}


func (this *GameKeyEvent) GetSourceId() int {
        return this.SourceId
}

