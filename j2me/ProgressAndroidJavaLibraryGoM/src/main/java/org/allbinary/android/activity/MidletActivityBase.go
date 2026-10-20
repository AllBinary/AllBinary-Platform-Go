/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type MidletActivityBase struct {
    Activity
}

// Implements GameFeatureListenerInterface, GameInitializedListenerInterface, SimpleProgressActivityInterface, KeyInputViewCompositeInterface, GenericMotionEventVerifyInterface

// The Java class MidletActivityBase has subclasses so its variables are this interface
type IMidletActivityBase interface {
    MidletActivityBase_() *MidletActivityBase
    OnGameFeatureChange(gameFeatureEvent *GameFeatureEvent)
    OnEvent(eventObject IAllBinaryEventObject)
    OnGameInitialized(gameInitializedEvent *GameInitializedEvent)
    OnSetProgress(value int, text string)
    OnDismissProgress(animationSequence []int)
    OnShowProgress(indeterminate bool, animationSequence []int)
    GetRootView() *View
    GetView() *View
    SetKeyInputView(keyInputView *View)
    OnGenericMotionEvent(event *MotionEvent) bool
}

func (this *MidletActivityBase) MidletActivityBase_() *MidletActivityBase {
    return this
}


func NewMidletActivityBase() *MidletActivityBase {
    this := &MidletActivityBase{}
    this.Activity = *NewActivity()
    return this
}


func (this *MidletActivityBase) OnGameFeatureChange(gameFeatureEvent *GameFeatureEvent) {
}


func (this *MidletActivityBase) OnEvent(eventObject IAllBinaryEventObject) {
}


func (this *MidletActivityBase) OnGameInitialized(gameInitializedEvent *GameInitializedEvent) {
}


// throws Exception
func (this *MidletActivityBase) OnSetProgress(value int, text string) {
}


// throws Exception
func (this *MidletActivityBase) OnDismissProgress(animationSequence []int) {
}


// throws Exception
func (this *MidletActivityBase) OnShowProgress(indeterminate bool, animationSequence []int) {
}


func (this *MidletActivityBase) GetRootView() *View {
        panic(errors.New(fmt.Sprint()))
}


func (this *MidletActivityBase) GetView() *View {
        panic(errors.New(fmt.Sprint()))
}


func (this *MidletActivityBase) SetKeyInputView(keyInputView *View) {
}


func (this *MidletActivityBase) OnGenericMotionEvent(event *MotionEvent) bool {
        panic(errors.New(fmt.Sprint()))
}

