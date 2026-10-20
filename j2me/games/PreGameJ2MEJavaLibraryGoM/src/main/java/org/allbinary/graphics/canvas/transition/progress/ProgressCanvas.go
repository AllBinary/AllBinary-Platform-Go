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

type ProgressCanvas struct {
    RunnableCanvas
    HasPainted bool
    BackgroundBasicColor *BasicColor
    GAUGE_PAINTABLE IPaintable
    MaxValue float32
    Gauge *ABCustomGaugeItem
    TEXT string
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    AllbinaryMidlet IAllBinaryMidlet
    Value float32
    Text string
    Background bool
    Paintable PaintableInterface
    InProgress bool
    IN_GAME_PROCESSOR IProcessor
    InGameProcessor IProcessor
    ProgressEndListener *ProgressEndListener
    BackgroundLabel string
    START_BACKGROUND string
    ADD_PORTION string
    ADD_EARLY_PORTION string
}

// Implements PaintableInterface, UpdateMyFontInterface

// The Java class ProgressCanvas has subclasses so its variables are this interface
type IProgressCanvas interface {
    IRunnableCanvas
    ProgressCanvas_() *ProgressCanvas
    UpdateMeasurement(graphics IGraphics)
    Init(gameMidlet IAllBinaryMidlet)
    Update(graphics IGraphics)
    InitCommands(cmdListener CommandListener)
    GetMaxValue() float32
    Start()
    StartBackground(background bool)
    EndActual()
    InGame()
    End()
    EndFromInitialLazyLoadingComplete()
    IsInGame() bool
    EndIfPaintedSinceStart()
    AddEarlyPortion(value int, text string, index int)
    AddPortion(value int, text string, index int)
    AddNormalPortion(value int, text string)
    SetValue(value int)
    Paint(graphics IGraphics)
    Paint2(graphics IGraphics)
    PaintThreed(graphics IGraphics)
    GetValue() float32
    SetText(text string)
    GetText() string
    SetBackground(background bool)
    IsBackground() bool
}

func (this *ProgressCanvas) ProgressCanvas_() *ProgressCanvas {
    return this
}


func (this *ProgressCanvas) initFields() {
    
    
    this.GAUGE_PAINTABLE = NewProgressPaintable(this)
    this.MaxValue = 100.0
    
    this.TEXT = this.CommonStrings.LOADING
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.AllbinaryMidlet = AllBinaryMidlet_GetNullInstance()
    
    this.Text = this.TEXT
    this.Background = true
    this.Paintable = this.GAUGE_PAINTABLE
    this.InProgress = false
    this.IN_GAME_PROCESSOR = /* Java anonymous class body: implement it as a named type
        this.PathFindingThreadPool = PathFindingThreadPool_GetInstance()
        

    // throws Exception
    func (this *ProgressCanvas) Process() {
            this.PathFindingThreadPool.RunAPriorityTask()
    }
*/ NewProcessor()
    this.InGameProcessor = Processor_GetInstance()
    this.ProgressEndListener = ProgressEndListener_Instance
    this.BackgroundLabel = "Background AI Game Loading..."
    this.START_BACKGROUND = "startBackground"
    this.ADD_PORTION = "addPortion"
    this.ADD_EARLY_PORTION = "addEarlyPortion"
}


func NewProgressCanvas(title string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *ProgressCanvas {
    this := &ProgressCanvas{}
        this.RunnableCanvas = *NewRunnableCanvas(NullCommandListener_NULL_COMMAND_LISTENER, CanvasStrings_GetInstance().EMPTY_CHILD_NAME_LIST, false)
        this.initFields()
        this.BackgroundBasicColor = backgroundBasicColor
        this.Gauge = NewABCustomGaugeItem(StringUtil_GetInstance().EMPTY_STRING, int(this.MaxValue), 0, backgroundBasicColor, foregroundBasicColor)
    return this
}


func (this *ProgressCanvas) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.Gauge.SetHeight(font.GetHeight() + 2)
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *ProgressCanvas) Init(gameMidlet IAllBinaryMidlet) {
        this.AllbinaryMidlet = gameMidlet
}


// throws Exception
func (this *ProgressCanvas) Update(graphics IGraphics) {
}


func (this *ProgressCanvas) InitCommands(cmdListener CommandListener) {
}


func (this *ProgressCanvas) GetMaxValue() float32 {
        return this.MaxValue
}


func (this *ProgressCanvas) Start() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.START_METHOD_NAME)
        this.SetBackground(true)
        this.Gauge.SetHeight(30)
        this.Gauge.SetLabel(this.CommonStrings.PLEASE_WAIT)
        this.SetText(this.TEXT)
        this.SetValue(0)
        this.InGameProcessor = Processor_GetInstance()
        this.Paintable = this.GAUGE_PAINTABLE
        this.InProgress = true
}


func (this *ProgressCanvas) StartBackground(background bool) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.START_BACKGROUND)
        this.SetBackground(background)
        this.MyFontProcessor = this.UpdateMyFontProcessor
        this.Gauge.SetLabel(this.BackgroundLabel)
        this.SetText(this.TEXT)
        this.SetValue(0)
        this.InGameProcessor = Processor_GetInstance()
        this.Paintable = this.GAUGE_PAINTABLE
}


func (this *ProgressCanvas) EndActual() {
        this.AllbinaryMidlet.CommandAction(GameCommandsFactory_GetInstance().SHOW_GAME_CANVAS, NullCanvas_NULL_CANVAS)
        this.InProgress = false
        this.InGame()
}


func (this *ProgressCanvas) InGame() {
        this.ProgressEndListener.OnEndEvent()
        this.InGameProcessor = this.IN_GAME_PROCESSOR
}


func (this *ProgressCanvas) End() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.END_METHOD_NAME)
        this.Gauge.SetValue(this.GetMaxValue())
        this.EndActual()
        this.Paintable = NullPaintable_GetInstance()
}


func (this *ProgressCanvas) EndFromInitialLazyLoadingComplete() {
        this.ProgressEndListener.OnEndEvent()
        this.Gauge.SetValue(this.GetMaxValue())
        this.InGameProcessor = this.IN_GAME_PROCESSOR
}


func (this *ProgressCanvas) IsInGame() bool {
        if this.InGameProcessor == this.IN_GAME_PROCESSOR {
            return true
        }
        return false
}


func (this *ProgressCanvas) EndIfPaintedSinceStart() {
}


func (this *ProgressCanvas) AddEarlyPortion(value int, text string, index int) {
        this.SetText(fmt.Sprint(NewStringMaker().Append(text).Append(fmt.Sprint(SmallIntegerSingletonFactory_GetInstance().GetAt(index)))))
        this.Gauge.SetValue(this.Gauge.GetValue() + this.GetMaxValue() / float32(value))
}


func (this *ProgressCanvas) AddPortion(value int, text string, index int) {
        this.SetText(fmt.Sprint(NewStringMaker().Append(text).Append(fmt.Sprint(SmallIntegerSingletonFactory_GetInstance().GetAt(index)))))
        PreLogUtil_Put(this.Text, this, this.ADD_PORTION)
        this.Gauge.SetValue(this.Gauge.GetValue() + this.GetMaxValue() / float32(value))
}


func (this *ProgressCanvas) AddNormalPortion(value int, text string) {
        if this.Text != text {
            PreLogUtil_Put(text, this, this.ADD_PORTION)
        }
        this.SetText(text)
        this.Gauge.SetValue(this.Gauge.GetValue() + this.GetMaxValue() / float32(value))
}


func (this *ProgressCanvas) SetValue(value int) {
        this.Value = float32(value)
        this.Gauge.SetValue(float32(value))
}


func (this *ProgressCanvas) Paint(graphics IGraphics) {
        this.Paintable.Paint(graphics)
}


func (this *ProgressCanvas) Paint2(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        graphics.SetColor(this.BackgroundBasicColor.IntValue())
        graphics.FillRect(0, 0, this.DisplayInfo.GetLastWidth(), this.DisplayInfo.GetLastHeight())
        this.Gauge.PaintXY(graphics, 0, 0)
        this.HasPainted = true
}


func (this *ProgressCanvas) PaintThreed(graphics IGraphics) {
}


func (this *ProgressCanvas) GetValue() float32 {
        return this.Value
}


func (this *ProgressCanvas) SetText(text string) {
        this.Text = text
}


func (this *ProgressCanvas) GetText() string {
        return this.Text
}


func (this *ProgressCanvas) SetBackground(background bool) {
        this.Background = background
}


func (this *ProgressCanvas) IsBackground() bool {
        return this.Background
}

type ProgressPaintable struct {
    Paintable
    ProgressCanvas IProgressCanvas
}


func (this *ProgressPaintable) initFields() {
        
}
    

    func NewProgressCanvas(progressCanvas IProgressCanvas) *ProgressCanvas {
        this := &ProgressCanvas{}
        this.Paintable = *NewPaintable()
        this.initFields()
            this.ProgressCanvas = progressCanvas
        return this
    }
    

    func (this *ProgressCanvas) Paint(graphics IGraphics) {
            this.ProgressCanvas.Paint2(graphics)
    }

