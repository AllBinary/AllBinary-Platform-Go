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

type NullGameCanvas struct {
    AllBinaryGameCanvas
    RunningNull bool
}

var NullGameCanvas_Instance any = NullUtil_GetInstance().NULL_OBJECT
var NullGameCanvas_NO_GAME string = "No Background Game"
var NullGameCanvas_TYPE int = 1

func (this *NullGameCanvas) initFields() {
    this.RunningNull = true
}


func NullGameCanvas_CreateNull() *NullGameCanvas {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *NullGameCanvas) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            LogUtil_GetInstance().Put(CommonStrings_GetInstance().EXCEPTION, "NullGameCanvas", CommonStrings_GetInstance().GET_INSTANCE, e)
            panic(errors.New(fmt.Sprint()))
                }
            }()
            tryControl1, tryResult1 = 1, NewNullGameCanvas(AllBinaryGameLayerManager_GetNullInstance()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func NullGameCanvas_GetInstance() *NullGameCanvas {
        if NullGameCanvas_Instance == NullUtil_GetInstance().NULL_OBJECT {
            NullGameCanvas_Instance = NullGameCanvas_CreateNull()
        }
        return any(NullGameCanvas_Instance).(*NullGameCanvas)
}


// throws Exception
func NullGameCanvas_CreateCanvas(gameLayerManager IAllBinaryGameLayerManager) *NullGameCanvas {
        var nullGameCanvas *NullGameCanvas = NewNullGameCanvas(gameLayerManager); _ = nullGameCanvas
        nullGameCanvas.SetInitialized(true)
        nullGameCanvas.SetTitle(NullGameCanvas_NO_GAME)
        return nullGameCanvas
}


func NewNullGameCanvas(gameLayerManager IAllBinaryGameLayerManager) *NullGameCanvas {
    this := &NullGameCanvas{}
        this.AllBinaryGameCanvas = *NewAllBinaryGameCanvas(NullCommandListener_NULL_COMMAND_LISTENER, gameLayerManager, NoHighScoresFactory_GetInstance(), BasicBuildGameInitializerFactory_NULL_BASE_BUILD_GMAE_INITIALIZER_FACTORY, false)
        this.initFields()
        this.AllBinaryGameCanvas.SetWait(1200)
    return this
}


// synchronized
func (this *NullGameCanvas) IsGameOver() bool {
        return false
}


func (this *NullGameCanvas) SetRunning(running bool) {
        this.RunningNull = running
}


// synchronized
func (this *NullGameCanvas) IsRunning() bool {
        return this.RunningNull
}


// synchronized
func (this *NullGameCanvas) Pause() {
        this.SetPaused(true)
}


// synchronized
func (this *NullGameCanvas) UnPause() {
        this.SetPaused(false)
}


func (this *NullGameCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.SetCommandListener(cmdListener)
}


// throws Exception
func (this *NullGameCanvas) BuildGameInit(isProgress bool) {
}


func (this *NullGameCanvas) Draw(graphics IGraphics) {
}


func (this *NullGameCanvas) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START_RUNNABLE, this, this.CommonStrings.RUN)
            this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *NullGameCanvas) GetType() int {
        return NullGameCanvas_TYPE
}

