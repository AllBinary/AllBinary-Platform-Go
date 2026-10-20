/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type DemoGameBehavior struct {
    LogUtil *LogUtil
}

// The Java class DemoGameBehavior has subclasses so its variables are this interface
type IDemoGameBehavior interface {
    DemoGameBehavior_() *DemoGameBehavior
    Init()
    Pause(allBinaryGameCanvas IAllBinaryGameCanvas)
    UnPause(allBinaryGameCanvas IAllBinaryGameCanvas)
    UpdateTouch(allBinaryGameCanvas IAllBinaryGameCanvas)
    UpdateScreenButtonPaintable(allBinaryGameCanvas IAllBinaryGameCanvas)
    SetGameState(allBinaryGameCanvas IAllBinaryGameCanvas)
    RemoveAllGameKeyInputListeners(allBinaryGameCanvas IAllBinaryGameCanvas)
    UpdateEndGameProcessor(allBinaryGameCanvas IAllBinaryGameCanvas)
    BuildGame(allBinaryGameCanvas IAllBinaryGameCanvas)
    Run(allBinaryGameCanvas IAllBinaryGameCanvas)
    SetHighScore(abeClientInformation AbeClientInformationInterface, allBinaryGameCanvas IAllBinaryGameCanvas, name string, score int64, autoSubmit bool, isLast bool)
}

func (this *DemoGameBehavior) DemoGameBehavior_() *DemoGameBehavior {
    return this
}

var DemoGameBehavior_Instance IDemoGameBehavior = NewDemoGameBehavior()

func (this *DemoGameBehavior) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewDemoGameBehavior() *DemoGameBehavior {
    this := &DemoGameBehavior{}
    this.initFields()
    return this
}


func DemoGameBehavior_GetInstance() IDemoGameBehavior {
        return DemoGameBehavior_Instance
}


// throws Exception
func (this *DemoGameBehavior) Init() {
}


func (this *DemoGameBehavior) Pause(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


func (this *DemoGameBehavior) UnPause(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *DemoGameBehavior) UpdateTouch(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *DemoGameBehavior) UpdateScreenButtonPaintable(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.SetTouchPaintableP(NullPaintable_GetInstance())
}


// throws Exception
func (this *DemoGameBehavior) SetGameState(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


func (this *DemoGameBehavior) RemoveAllGameKeyInputListeners(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *DemoGameBehavior) UpdateEndGameProcessor(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *DemoGameBehavior) BuildGame(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *DemoGameBehavior) Run(allBinaryGameCanvas IAllBinaryGameCanvas) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(GameTypeFactory_GetInstance().BOT), this, commonStrings.RUN)
}


// throws Exception
func (this *DemoGameBehavior) SetHighScore(abeClientInformation AbeClientInformationInterface, allBinaryGameCanvas IAllBinaryGameCanvas, name string, score int64, autoSubmit bool, isLast bool) {
}

