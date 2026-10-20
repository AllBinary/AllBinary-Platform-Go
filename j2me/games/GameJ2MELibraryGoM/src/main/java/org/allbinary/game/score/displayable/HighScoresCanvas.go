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

type HighScoresCanvas struct {
    GameCommandCanvas
    Paintable IPaintable
    HighScoreCommandsFactory *HighScoreCommandsFactory
    WaitPaintable *SimpleTextPaintable
    HighScoresPaintable *HighScoresPaintable
    HighScoresFactoryInterface HighScoresFactoryInterface
    ColorFillPaintable IColorFillBasePaintable
    HighScoresHelper IHighScoresHelperBase
    GameInfo *GameInfo
    HighScoresCanvasInputProcessor IHighScoresCanvasInputProcessor
    CurrentCommand ICommand
    HasPainted bool
}

// Implements HighScoresResultsListener

var HighScoresCanvas_NAME string = "HighScoresCanvas"

func (this *HighScoresCanvas) initFields() {
    this.Paintable = NullPaintable_GetInstance()
    this.HighScoreCommandsFactory = HighScoreCommandsFactory_GetInstance()
    this.WaitPaintable = NewSimpleTextPaintable(this.CommonStrings.PLEASE_WAIT_FOR_SERVER, BasicColorFactory_GetInstance().WHITE)
    
    
    
    this.HighScoresHelper = NewHighScoresHelperBase()
    
    
    this.CurrentCommand = this.HighScoreCommandsFactory.HIGH_SCORE_COMMANDS[0]
    this.HasPainted = false
}


func NewHighScoresCanvas(commandListener CommandListener, allBinaryGameLayerManager IAllBinaryGameLayerManager, gameInfo *GameInfo, paintable *HighScoresPaintable, highScoresFactoryInterface HighScoresFactoryInterface, highScoresCanvasInputProcessorFactoryInterface HighScoresCanvasInputProcessorFactoryInterface) *HighScoresCanvas {
    this := &HighScoresCanvas{}
        this.GameCommandCanvas = *NewGameCommandCanvas(commandListener, HighScoresCanvas_NAME, allBinaryGameLayerManager.GetBackgroundBasicColor(), allBinaryGameLayerManager.GetForegroundBasicColor())
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.HighScoresPaintable = paintable
        this.HighScoresFactoryInterface = highScoresFactoryInterface
        this.HighScoresCanvasInputProcessor = highScoresCanvasInputProcessorFactoryInterface.GetInstance(this)
        this.GameInfo = gameInfo
        this.WaitPaintable.SetBasicColorP(allBinaryGameLayerManager.GetForegroundBasicColor())
        this.GetHighScoresPaintable().SetBasicColorP(allBinaryGameLayerManager.GetForegroundBasicColor())
        this.ColorFillPaintable = ColorFillPaintableFactory_GetInstance().GetInstance(allBinaryGameLayerManager.GetBackgroundBasicColor(), false)
        if this.HighScoresHelper.GetHighScoresArray() == NoHighScoresFactory_GetInstance().NO_HIGH_SCORES {
            this.SetPaintable(this.WaitPaintable)
        } else {
            this.LogUtil.PutF("Show HighScores that are already loaded", this, this.CommonStrings.CONSTRUCTOR)
            this.UpdateCommand(this.CurrentCommand)
            this.SetPaintable(this.GetHighScoresPaintable())
        }
        /* Java local class: move it out of the function
type HighScoreRunnable struct {
    ARunnable
    HighScoresCanvas *HighScoresCanvas
}


func (this *HighScoreRunnable) initFields() {
            
}
        

        func NewHighScoresCanvas(highScoresCanvas *HighScoresCanvas) *HighScoresCanvas {
            this := &HighScoresCanvas{}
            this.ARunnable = *NewARunnable()
            this.initFields()
                this.HighScoresCanvas = highScoresCanvas
            return this
        }
        

        func (this *HighScoresCanvas) Run() {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                var highScoresCanvas *HighScoresCanvas = this.HighScoresCanvas; _ = highScoresCanvas
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                        }
                    }()
                    var isHTML bool = J2MEUtil_IsHTML(); _ = isHTML
                    if !isHTML {
                        for !highScoresCanvas.HasPainted {
                        }
                        highScoresCanvas.HasPainted = false
                    }
                    var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
                    var systemWrapper *ABSystemWrapper = ABSystemWrapper_GetInstance(); _ = systemWrapper
                    var currentTimeMillis int64 = systemWrapper.CurrentTimeMillis(); _ = currentTimeMillis
                    logUtil.PutF(fmt.Sprint(stringMaker.Append("HighScoresCanvas - Request repaint to be sure: ").Appendlong(currentTimeMillis)), this, commonStrings.RUN)
                    highScoresCanvas.RepaintBehavior.OnChangeRepaint(highScoresCanvas)
                    if !isHTML {
                        for !highScoresCanvas.HasPainted {
                        }
                    }
                    stringMaker.Delete(0, stringMaker.Length())
                    logUtil.PutF(fmt.Sprint(stringMaker.Append("HighScoresCanvas - Now that the canvas has completed repaint go ahead and fetch the scores: ").Appendlong(currentTimeMillis)), this, commonStrings.RUN)
                    highScoresCanvas.ExecuteUpdate()
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/
        SecondaryThreadPool_GetInstance().RunTask(NewHighScoreRunnable(this))
    return this
}


func (this *HighScoresCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.AddCommand(GameCommandsFactory_GetInstance().CLOSE_AND_SHOW_GAME_CANVAS)
        this.SetCommandListener(cmdListener)
}


func (this *HighScoresCanvas) Open() {
        this.GameCommandCanvas.Open()
        this.HighScoresCanvasInputProcessor.Open()
}


// throws Exception
func (this *HighScoresCanvas) Close() {
        this.GameCommandCanvas.Close()
        this.HighScoresCanvasInputProcessor.Close()
}


func (this *HighScoresCanvas) Paint(graphics IGraphics) {
        this.ColorFillPaintable.Paint(graphics)
        this.Paintable.Paint(graphics)
        if this.WaitPaintable != this.Paintable {
            this.HighScoresCanvasInputProcessor.Paint(graphics)
        }
        this.GameCommandCanvas.Paint(graphics)
        this.HasPainted = true
}


func (this *HighScoresCanvas) ExecuteUpdate() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.UPDATE, e)
                }
            }()
            this.HighScoresFactoryInterface.FetchHighScores(this.GetGameInfo(), this)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *HighScoresCanvas) GetHighScoresPaintable() *HighScoresPaintable {
        return this.HighScoresPaintable
}


func (this *HighScoresCanvas) SetHighScoresArray(highScoresArray []IHighScores) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.UPDATE, e)
                }
            }()
            if highScoresArray != nil {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Appendint(len(highScoresArray))), this, "setHighScoresArray")
            } else {
                this.LogUtil.PutF(this.CommonStrings.START, this, "setHighScoresArray")
            }
            this.HighScoresHelper.SetHighScoresArray(highScoresArray)
            this.UpdateCommand(this.CurrentCommand)
            this.SetPaintable(this.GetHighScoresPaintable())
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *HighScoresCanvas) UpdateCommand(command ICommand) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Append(this.StringUtil.ToString(command))), this, this.CommonStrings.UPDATE)
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        if this.HighScoreCommandsFactory.IsHighScoreCommand(command) {
            var index int = this.HighScoreCommandsFactory.GetIndex(command); _ = index
            var nextIndex int = index + 1; _ = nextIndex
            var highScoresArray []IHighScores = this.HighScoresHelper.GetHighScoresArray(); _ = highScoresArray
            if nextIndex >= len(highScoresArray) {
                nextIndex = 0
            }
            if len(highScoresArray) > 0 {
                this.GetHighScoresPaintable().SetHighScores(highScoresArray[index])
            } else {
                this.GetHighScoresPaintable().SetHighScores(NullHighScoresSingletonFactory_GetInstance())
            }
            if index != nextIndex {
                this.RemoveAllCommands()
                this.AddCommand(gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS)
                this.AddCommand(this.HighScoreCommandsFactory.HIGH_SCORE_COMMANDS[nextIndex])
            }
        }
        this.CurrentCommand = command
        this.Close()
        this.GameCommandCanvas.InitMenu()
        this.Open()
        this.RepaintBehavior.OnChangeRepaint(this)
}


func (this *HighScoresCanvas) GetGameInfo() *GameInfo {
        return this.GameInfo
}


func (this *HighScoresCanvas) SetPaintable(paintable IPaintable) {
        this.Paintable = paintable
        this.RepaintBehavior.OnChangeRepaint(this)
}

