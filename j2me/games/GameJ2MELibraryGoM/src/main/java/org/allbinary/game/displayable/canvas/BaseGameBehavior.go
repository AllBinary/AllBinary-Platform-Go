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

type BaseGameBehavior struct {
    DemoGameBehavior
}

var BaseGameBehavior_InstanceC *BaseGameBehavior = NewBaseGameBehavior()

func NewBaseGameBehavior() *BaseGameBehavior {
    this := &BaseGameBehavior{}
    this.DemoGameBehavior = *NewDemoGameBehavior()
    return this
}


func BaseGameBehavior_GetInstance() IDemoGameBehavior {
        return BaseGameBehavior_InstanceC
}


// throws Exception
func (this *BaseGameBehavior) Init() {
        var gameAdState *GameAdState = GameAdStateFactory_GetInstance().GetCurrentInstance(); _ = gameAdState
        gameAdState.PlayingAdState()
}


func (this *BaseGameBehavior) Pause(allBinaryGameCanvas IAllBinaryGameCanvas) {
        var features *Features = Features_GetInstance(); _ = features
        if features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL_AS_GAME_THREAD) {
            var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
            currentDisplayableFactory.SetRunnable(allBinaryGameCanvas.AllBinaryGameCanvas_().GamePauseRunnable)
        }
}


func (this *BaseGameBehavior) UnPause(allBinaryGameCanvas IAllBinaryGameCanvas) {
        if Features_GetInstance().IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL_AS_GAME_THREAD) {
            var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
            currentDisplayableFactory.SetRunnable(allBinaryGameCanvas.AllBinaryGameCanvas_().GameRunnable)
            currentDisplayableFactory.SetDisplayable(allBinaryGameCanvas)
        }
}


// throws Exception
func (this *BaseGameBehavior) UpdateTouch(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.UpdateTouch2()
}


// throws Exception
func (this *BaseGameBehavior) UpdateScreenButtonPaintable(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.UpdateScreenButtonPaintable2()
}


// throws Exception
func (this *BaseGameBehavior) SetGameState(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.UpdateGameState()
}


func (this *BaseGameBehavior) RemoveAllGameKeyInputListeners(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.RemoveAllGameKeyInputListeners2()
}


// throws Exception
func (this *BaseGameBehavior) UpdateEndGameProcessor(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.UpdateEndGameProcessor2()
}


// throws Exception
func (this *BaseGameBehavior) BuildGame(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.BuildGame2()
}


// throws Exception
func (this *BaseGameBehavior) Run(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.Run2()
}


// throws Exception
func (this *BaseGameBehavior) SetHighScore(abeClientInformation AbeClientInformationInterface, allBinaryGameCanvas IAllBinaryGameCanvas, name string, score int64, autoSubmit bool, isLast bool) {
        allBinaryGameCanvas.SetHighScore2(abeClientInformation, name, score, autoSubmit, isLast)
}

