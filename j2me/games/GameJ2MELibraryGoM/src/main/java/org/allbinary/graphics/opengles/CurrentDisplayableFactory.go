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

type CurrentDisplayableFactory struct {
    CommonStrings *CommonStrings
    CanvasStrings *CanvasStrings
    StringUtil *StringUtil
    SET_DISPLAYABLE string
    SET_RUNNABLE string
    CLEAR_RUNNABLE string
    SET_USED_RUNNABLE string
    RUNNABLE string
    RUNNABLE_WAS_ALREADY_IN_USE string
    DEFAULT_CHILD_NAME_LIST IBasicArrayList
    Displayable IDisplayable
    OpenGlReadydisplayable IDisplayable
    Runnable IGameRunnable
    UsedRunnable IGameRunnable
}

var CurrentDisplayableFactory_Instance *CurrentDisplayableFactory = NewCurrentDisplayableFactory()

func (this *CurrentDisplayableFactory) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.CanvasStrings = CanvasStrings_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.SET_DISPLAYABLE = "setDisplayable"
    this.SET_RUNNABLE = "setRunnable"
    this.CLEAR_RUNNABLE = "clearRunnable"
    this.SET_USED_RUNNABLE = "setUsedRunnable"
    this.RUNNABLE = "Runnable: "
    this.RUNNABLE_WAS_ALREADY_IN_USE = "Runnable was already in use: "
    this.DEFAULT_CHILD_NAME_LIST = NewBasicArrayListD()
    this.Displayable = NullDisplayable_GetInstance()
    this.OpenGlReadydisplayable = NullDisplayable_GetInstance()
    this.Runnable = NullWaitGameRunnable_GetInstance()
    this.UsedRunnable = NullWaitGameRunnable_GetInstance()
}


func CurrentDisplayableFactory_GetInstance() *CurrentDisplayableFactory {
        return CurrentDisplayableFactory_Instance
}


func NewCurrentDisplayableFactory() *CurrentDisplayableFactory {
    this := &CurrentDisplayableFactory{}
    this.initFields()
        this.DEFAULT_CHILD_NAME_LIST.Add(this.CanvasStrings.ABOUT)
        this.DEFAULT_CHILD_NAME_LIST.Add(this.CanvasStrings.OPTIONS)
        this.DEFAULT_CHILD_NAME_LIST.Add(GameInputMappingCanvas_NAME)
        this.DEFAULT_CHILD_NAME_LIST.Add(GameInputMappingInstructionsCanvas_NAME)
        this.DEFAULT_CHILD_NAME_LIST.Add(HighScoresCanvas_NAME)
    return this
}


func (this *CurrentDisplayableFactory) SetMyCanvas(displayable IMyCanvas) {
        if func() bool { _, ok := any(this.Displayable).(IMyCanvas); return ok }() {
            if displayable.HasChild(any(this.Displayable).(IMyCanvas)) {
                PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Child Displayable is already set for: ").Append(this.StringUtil.ToString(displayable))), this, this.SET_DISPLAYABLE)
                return
            }
        }
        this.SetDisplayable(displayable)
}


func (this *CurrentDisplayableFactory) SetDisplayable(displayable IDisplayable) {
        this.Displayable = displayable
        this.SetOpenGlReadydisplayable(displayable)
}


func (this *CurrentDisplayableFactory) SetOpenGlReadydisplayable(openGlReadydisplayable IDisplayable) {
        this.OpenGlReadydisplayable = openGlReadydisplayable
        this.Update()
}


func (this *CurrentDisplayableFactory) GetOpenGlReadydisplayable() IDisplayable {
        return this.OpenGlReadydisplayable
}


func (this *CurrentDisplayableFactory) SetRunnable(runnable IGameRunnable) {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.RUNNABLE).Append(this.StringUtil.ToString(runnable))), this, this.SET_RUNNABLE)
        this.Runnable = runnable
        if this.UsedRunnable != this.Runnable {
            this.Update()
        } else {
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.RUNNABLE_WAS_ALREADY_IN_USE).Append(this.StringUtil.ToString(runnable))), this, this.SET_RUNNABLE)
        }
}


func (this *CurrentDisplayableFactory) ClearRunnable() {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.RUNNABLE).Append(this.StringUtil.ToString(NullRunnable_GetInstance()))), this, this.CLEAR_RUNNABLE)
        this.SetUsedRunnable(NullWaitGameRunnable_GetInstance())
}


func (this *CurrentDisplayableFactory) Update() {
        { // synchronized (this)
            var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
            var features *Features = Features_GetInstance(); _ = features
            var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
            stringMaker.Append(this.StringUtil.ToString(this.OpenGlReadydisplayable))
            if SWTUtil_IsSWT && !features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
                PreLogUtil_Put(fmt.Sprint(stringMaker.Append(commonSeps.SPACE).Append(SWTUtil_SWT).Append(commonSeps.SPACE).Append(this.RUNNABLE).Append(this.StringUtil.ToString(NullRunnable_GetInstance()))), this, this.CommonStrings.UPDATE)
                this.SetUsedRunnable(NullWaitGameRunnable_GetInstance())
            } else if J2MEUtil_IsHTML() {
                PreLogUtil_Put(fmt.Sprint(stringMaker.Append(commonSeps.SPACE).Append(GraphicsStrings_GetInstance().HTML).Append(commonSeps.SPACE).Append(this.RUNNABLE).Append(this.StringUtil.ToString(this.Runnable))), this, this.CommonStrings.UPDATE)
                this.SetUsedRunnable(this.Runnable)
            } else if func() bool { _, ok := any(this.OpenGlReadydisplayable).(*DemoCanvas); return ok }() || func() bool { _, ok := any(this.OpenGlReadydisplayable).(IAllBinaryGameCanvas); return ok }() || features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
                PreLogUtil_Put(fmt.Sprint(stringMaker.Append(commonSeps.SPACE).Append(fmt.Sprint(OpenGLFeatureFactory_GetInstance().OPENGL)).Append(commonSeps.SPACE).Append(this.RUNNABLE).Append(this.StringUtil.ToString(this.Runnable))), this, this.CommonStrings.UPDATE)
                this.SetUsedRunnable(this.Runnable)
            } else {
                PreLogUtil_Put(fmt.Sprint(stringMaker.Append(this.RUNNABLE).Append(this.StringUtil.ToString(NullRunnable_GetInstance()))), this, this.CommonStrings.UPDATE)
                this.SetUsedRunnable(NullWaitGameRunnable_GetInstance())
            }
        }
}


func (this *CurrentDisplayableFactory) GetDisplayable() IDisplayable {
        return this.Displayable
}


func (this *CurrentDisplayableFactory) GetUsedRunnable() IGameRunnable {
        return this.UsedRunnable
}


func (this *CurrentDisplayableFactory) SetUsedRunnable(usedRunnable IGameRunnable) {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.RUNNABLE).Append(this.StringUtil.ToString(this.Runnable)).Append(CommonSeps_GetInstance().COLON_SEP).Append(this.StringUtil.ToString(usedRunnable))), this, this.SET_USED_RUNNABLE)
        this.UsedRunnable = usedRunnable
}

