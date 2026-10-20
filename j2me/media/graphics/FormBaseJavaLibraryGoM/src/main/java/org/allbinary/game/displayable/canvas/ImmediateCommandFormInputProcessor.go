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

type ImmediateCommandFormInputProcessor struct {
    CommandFormInputProcessor
    TimeHelper ITimeDelayHelper
    GameCommandCanvas IGameCommandCanvas
}


func (this *ImmediateCommandFormInputProcessor) initFields() {
    this.TimeHelper = NewTimeDelayHelper(25)
    
}


func NewImmediateCommandFormInputProcessor(gameKeyEventList IBasicArrayList, playerInputId int, gameCommandCanvas IGameCommandCanvas, form IScrollSelectionForm) *ImmediateCommandFormInputProcessor {
    this := &ImmediateCommandFormInputProcessor{}
        this.CommandFormInputProcessor = *NewCommandFormInputProcessor(gameKeyEventList, playerInputId, gameCommandCanvas, form)
        this.initFields()
        this.GameCommandCanvas = gameCommandCanvas
    return this
}


// synchronized
func (this *ImmediateCommandFormInputProcessor) OnPressGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.ON_PRESS_GAME_KEY, e)
                }
            }()
            this.CommandFormInputProcessor.OnPressGameKeyEvent(gameKeyEvent)
            if this.TimeHelper.IsTimeTNT() {
                this.ProcessInputList()
                this.GameCommandCanvas.GameCommandCanvas_().RepaintBehavior.OnChangeRepaint(this.GetCanvas())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *ImmediateCommandFormInputProcessor) OnDownGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.ON_DOWN_GAME_KEY, e)
                }
            }()
            this.CommandFormInputProcessor.OnDownGameKeyEvent(gameKeyEvent)
            if this.TimeHelper.IsTimeTNT() {
                this.ProcessInputList()
                this.GameCommandCanvas.GameCommandCanvas_().RepaintBehavior.OnChangeRepaint(this.GetCanvas())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *ImmediateCommandFormInputProcessor) OnUpGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.GameInputStrings.ON_UP_GAME_KEY, e)
                }
            }()
            this.CommandFormInputProcessor.OnUpGameKeyEvent(gameKeyEvent)
            if this.TimeHelper.IsTimeTNT() {
                this.ProcessInputList()
                this.GameCommandCanvas.GameCommandCanvas_().RepaintBehavior.OnChangeRepaint(this.GetCanvas())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *ImmediateCommandFormInputProcessor) OnMotionGestureEvent(motionGestureEvent *MotionGestureEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "onMotionGestureEvent", e)
                }
            }()
            this.CommandFormInputProcessor.OnMotionGestureEvent(motionGestureEvent)
            if this.TimeHelper.IsTimeTNT() {
                this.ProcessInputList()
                this.GameCommandCanvas.GameCommandCanvas_().RepaintBehavior.OnChangeRepaint(this.GetCanvas())
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

