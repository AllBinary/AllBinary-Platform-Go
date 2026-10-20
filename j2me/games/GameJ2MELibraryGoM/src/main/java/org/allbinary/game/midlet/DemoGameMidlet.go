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

type DemoGameMidlet struct {
    GameMidlet
}

// The Java class DemoGameMidlet has subclasses so its variables are this interface
type IDemoGameMidlet interface {
    IGameMidlet
    DemoGameMidlet_() *DemoGameMidlet
    IsContinue() bool
    GetHighestLevel() int
    CreateDemoGameCanvasRunnableInterface() GameCanvasRunnableInterface
    CreateGameCanvasRunnable(allBinaryGameLayerManager IAllBinaryGameLayerManager) GameCanvasRunnableInterface
    CreateGameCanvasRunnableInterface() GameCanvasRunnableInterface
    DemoSetup()
    PostDemoSetup()
    SetDemo()
    CreateGame()
    IsReady() bool
}

func (this *DemoGameMidlet) DemoGameMidlet_() *DemoGameMidlet {
    return this
}



func NewDemoGameMidlet(clientInformationFactory *ClientInformationFactory) *DemoGameMidlet {
    this := &DemoGameMidlet{}
        this.GameMidlet = *NewGameMidlet(clientInformationFactory)
    return this
}


// throws Exception
func (this *DemoGameMidlet) IsContinue() bool {
        var hashtable *ABHashtable[any, any] = this.GetStartStateHashtable(); _ = hashtable
        if hashtable != nil && hashtable.Size() > 0 {
            return true
        } else {
            return false
        }
}


func (this *DemoGameMidlet) GetHighestLevel() int {
        return 1
}


// throws Exception
func (this *DemoGameMidlet) CreateDemoGameCanvasRunnableInterface() GameCanvasRunnableInterface {
        this.LogUtil.PutF(this.CommonStrings.NOT_IMPLEMENTED, this, this.CommonStrings.CREATE)
        return NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE
}


// throws Exception
func (this *DemoGameMidlet) CreateGameCanvasRunnable(allBinaryGameLayerManager IAllBinaryGameLayerManager) GameCanvasRunnableInterface {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *DemoGameMidlet) CreateGameCanvasRunnableInterface() GameCanvasRunnableInterface {
        return this.CreateGameCanvasRunnable(this.CreateGameLayerManager())
}


func (this *DemoGameMidlet) DemoSetup() {
}


func (this *DemoGameMidlet) PostDemoSetup() {
}


// throws Exception
// synchronized
func (this *DemoGameMidlet) SetDemo() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "setDemo")
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.Start()
        PrimaryThreadPool_GetInstance().RunTask(NewDemoRunnable(this))
}


// throws Exception
// synchronized
func (this *DemoGameMidlet) CreateGame() {
        var CREATE_GAME string = "createGame"; _ = CREATE_GAME
        this.LogUtil.PutF(this.CommonStrings.START, this, CREATE_GAME)
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.Start()
        var hashtable *ABHashtable[any, any] = this.GetStartStateHashtable(); _ = hashtable
        this.SetStartStateHashtable(StdUtil_GetInstance().NULL_TABLE)
        PrimaryThreadPool_GetInstance().RunTask(NewCreateGameRunnable(this, hashtable))
        this.LogUtil.PutF(this.CommonStrings.END, this, CREATE_GAME)
}


func (this *DemoGameMidlet) IsReady() bool {
        return true
}

