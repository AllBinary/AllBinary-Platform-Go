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

type GameKeyEventFactory struct {
    LogUtil *LogUtil
    TOUCH_BUTTON_SOURCE_ID int
    MOTION_GESTURE_SOURCE_ID int
    MAX_SOURCES int
    ARRAY [][]*GameKeyEvent
}

var GameKeyEventFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameKeyEventFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TOUCH_BUTTON_SOURCE_ID = 2
    this.MOTION_GESTURE_SOURCE_ID = 3
    this.MAX_SOURCES = 4
    this.ARRAY = func() [][]*GameKeyEvent { array := make([][]*GameKeyEvent, this.MAX_SOURCES); for index := range array { array[index] = make([]*GameKeyEvent, InputFactory_GetInstance().MAX) }; return array }()
}


func GameKeyEventFactory_GetInstance() *GameKeyEventFactory {
        if GameKeyEventFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameKeyEventFactory_Instance = NewGameKeyEventFactory()
        }
        return any(GameKeyEventFactory_Instance).(*GameKeyEventFactory)
}


func NewGameKeyEventFactory() *GameKeyEventFactory {
    this := &GameKeyEventFactory{}
    this.initFields()
    return this
}


func (this *GameKeyEventFactory) Init() {
        var nullUtil *NullUtil = NullUtil_GetInstance(); _ = nullUtil
        var size int = InputFactory_GetInstance().MAX; _ = size
        for index := int(this.MAX_SOURCES); Java_PreDecrement(&index) >= 0;  {
            for index2 := int(size); Java_PreDecrement(&index2) >= 0;  {
                this.ARRAY[index][index2] = GameKeyEvent_CreateEvent(nullUtil.NULL_OBJECT, index, index2)
            }
        }
}


// throws Exception
func (this *GameKeyEventFactory) GetInstanceForKey(anyType GameKeyEventSourceInterface, key int) *GameKeyEvent {
        var gameKeyEvent *GameKeyEvent = this.ARRAY[anyType.GetSourceId()][key]; _ = gameKeyEvent
        return gameKeyEvent
}


// throws Exception
func (this *GameKeyEventFactory) GetInstanceForInput(anyType GameKeyEventSourceInterface, input IInput) *GameKeyEvent {
        var gameKeyEvent *GameKeyEvent = this.ARRAY[anyType.GetSourceId()][input.GetId()]; _ = gameKeyEvent
        return gameKeyEvent
}

