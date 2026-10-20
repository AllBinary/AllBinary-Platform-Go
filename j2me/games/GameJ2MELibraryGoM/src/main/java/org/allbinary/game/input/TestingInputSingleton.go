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

type TestingInputSingleton struct {
    AllBinaryGameLayer
    InputProcessorArray []IGameInputProcessor
    Paintable PaintableInterface
}

// Implements GameInputInterface

var TestingInputSingleton_Instance *TestingInputSingleton = NewTestingInputSingleton()

func (this *TestingInputSingleton) initFields() {
    this.InputProcessorArray = make([]IGameInputProcessor, InputFactory_GetInstance().MAX)
    this.Paintable = NullPaintable_GetInstance()
}


func TestingInputSingleton_GetInstance() *TestingInputSingleton {
        return TestingInputSingleton_Instance
}


func NewTestingInputSingleton() *TestingInputSingleton {
    this := &TestingInputSingleton{}
        this.AllBinaryGameLayer = *NewAllBinaryGameLayer(StringUtil_GetInstance().EMPTY_STRING, RectangleFactory_SINGLETON, ViewPosition_GetInstanceD())
        this.initFields()
        GameInputProcessorUtil_Init(this.InputProcessorArray)
    return this
}


func (this *TestingInputSingleton) InitInputProcessors() {
}


func (this *TestingInputSingleton) InitInputProcessorsForTestInput(aTestInputInterface TestInputInterface) {
        var testInputInterface TestInputInterface = aTestInputInterface; _ = testInputInterface
        this.Paintable = testInputInterface
        this.InputProcessorArray[Canvas_UP] = /* Java anonymous class body: implement it as a named type
            

        // throws Exception
        func (this *TestingInputSingleton) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
                testInputInterface.Up()
        }
*/ NewGameInputProcessor()
        this.InputProcessorArray[Canvas_DOWN] = /* Java anonymous class body: implement it as a named type
            

        // throws Exception
        func (this *TestingInputSingleton) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
                testInputInterface.Down()
        }
*/ NewGameInputProcessor()
        this.InputProcessorArray[Canvas_LEFT] = /* Java anonymous class body: implement it as a named type
            

        // throws Exception
        func (this *TestingInputSingleton) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
                testInputInterface.Left()
        }
*/ NewGameInputProcessor()
        this.InputProcessorArray[Canvas_RIGHT] = /* Java anonymous class body: implement it as a named type
            

        // throws Exception
        func (this *TestingInputSingleton) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
                testInputInterface.Right()
        }
*/ NewGameInputProcessor()
        this.InputProcessorArray[Canvas_KEY_NUM7] = /* Java anonymous class body: implement it as a named type
            

        // throws Exception
        func (this *TestingInputSingleton) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
                testInputInterface.StrafeLeft()
        }
*/ NewGameInputProcessor()
        this.InputProcessorArray[Canvas_KEY_NUM9] = /* Java anonymous class body: implement it as a named type
            

        // throws Exception
        func (this *TestingInputSingleton) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
                testInputInterface.StrafeRight()
        }
*/ NewGameInputProcessor()
        GameInputProcessorUtil_Init(this.InputProcessorArray)
}


// throws Exception
// synchronized
func (this *TestingInputSingleton) ProcessInputList(allbinaryLayerManager IAllBinaryLayerManager, list IBasicArrayList) {
        var key int = 0; _ = key
        var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            gameKeyEvent = (any(list.BasicArrayList_().ObjectArray[index]).(*GameKeyEvent))
            key = gameKeyEvent.GetKey()
            this.InputProcessorArray[key].ProcessEvent(allbinaryLayerManager, GameKeyEvent_NONE)
        }
}


// throws Exception
// synchronized
func (this *TestingInputSingleton) ProcessInput(allbinaryLayerManager IAllBinaryLayerManager) {
        this.ProcessInputList(allbinaryLayerManager, this.GetGameKeyEventList())
        this.GetGameKeyEventList().Clear()
}


func (this *TestingInputSingleton) Paint(graphics IGraphics) {
        this.Paintable.Paint(graphics)
}


func (this *TestingInputSingleton) ImplmentsGameInputInterface() bool {
        return true
}

