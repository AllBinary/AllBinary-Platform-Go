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

type GameKeyCompleteMotionGestureInputEventFactory struct {
    LogUtil *LogUtil
    EventList IBasicArrayList
}

var GameKeyCompleteMotionGestureInputEventFactory_SINGLETON *GameKeyCompleteMotionGestureInputEventFactory = NewGameKeyCompleteMotionGestureInputEventFactory()

func (this *GameKeyCompleteMotionGestureInputEventFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.EventList = NewBasicArrayListD()
}

func NewGameKeyCompleteMotionGestureInputEventFactory() *GameKeyCompleteMotionGestureInputEventFactory {
    this := &GameKeyCompleteMotionGestureInputEventFactory{}
    this.initFields()
    return this
}


func GameKeyCompleteMotionGestureInputEventFactory_GetInstance() *GameKeyCompleteMotionGestureInputEventFactory {
        return GameKeyCompleteMotionGestureInputEventFactory_SINGLETON
}


func (this *GameKeyCompleteMotionGestureInputEventFactory) Add(gameKeyCompleteMotionGestureInputEvent IGameKeyCompleteMotionGestureInputEvent) {
        this.EventList.Add(gameKeyCompleteMotionGestureInputEvent)
}


func (this *GameKeyCompleteMotionGestureInputEventFactory) UpdateAll() {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Appendint(this.EventList.Size())), this, "updateAll")
        for index := int(this.EventList.Size() - 1); index >= 0; index-- {
            var gameKeyCompleteMotionGestureInputEvent IGameKeyCompleteMotionGestureInputEvent = any(this.EventList.BasicArrayList_().ObjectArray[index]).(IGameKeyCompleteMotionGestureInputEvent); _ = gameKeyCompleteMotionGestureInputEvent
            gameKeyCompleteMotionGestureInputEvent.Update()
        }
}

