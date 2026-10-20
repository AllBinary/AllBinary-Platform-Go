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

type GameCommandCanvas struct {
    MyCanvas
    RepaintProcessor IProcessor
    GameInputStrings *GameInputStrings
    ForegroundBasicColor *BasicColor
    BackgroundBasicColor *BasicColor
    InputToGameKeyMapping *InputToGameKeyMapping
    GameKeyFactory *GameKeyFactory
    GameKeyEventFactory *GameKeyEventFactory
    DownGameKeyEventHandler *DownGameKeyEventHandler
    UpGameKeyEventHandler *UpGameKeyEventHandler
    RepaintBehavior IRepaintBehavior
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    ForegroundColor int
    BackgroundColor int
    MenuInputProcessor IBasicMenuInputProcessor
    MenuPaintable IPaintable
    MenuForm IPaintableForm
    IsSingleKeyRepeatableProcessing bool
    FontHeight int
}

// Implements MenuListener, DisplayChangeEventListener, UpdateMyFontInterface

// The Java class GameCommandCanvas has subclasses so its variables are this interface
type IGameCommandCanvas interface {
    IMyCanvas
    GameCommandCanvas_() *GameCommandCanvas
    UpdateMeasurement(graphics IGraphics)
    OnEvent(eventObject IAllBinaryEventObject)
    OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent)
    InitCommands(cmdListener CommandListener)
    InitMenu()
    CreateForm() IScrollSelectionForm
    CreateRectangle(size int) *Rectangle
    Open()
    Close()
    Update()
    GetSourceId() int
    KeyPressed(keyCode int)
    KeyReleased(keyCode int)
    KeyRepeated(keyCode int)
    KeyPressedByDevice(keyCode int, deviceId int)
    KeyReleasedByDevice(keyCode int, deviceId int)
    KeyRepeatedByDevice(keyCode int, deviceId int)
    AddGameKeyEvent(keyCode int, deviceId int, repeated bool)
    RemoveGameKeyEvent(keyCode int, deviceId int, repeated bool)
    Paint(graphics IGraphics)
    SetMenuInputProcessor(menuInputProcessor IBasicMenuInputProcessor)
    GetMenuInputProcessor() IBasicMenuInputProcessor
}

func (this *GameCommandCanvas) GameCommandCanvas_() *GameCommandCanvas {
    return this
}

var GameCommandCanvas_Id int = 0

func (this *GameCommandCanvas) initFields() {
    this.RepaintProcessor = ScreenRepaintProcessorFactory_GetInstance().Create(this)
    this.GameInputStrings = GameInputStrings_GetInstance()
    
    
    this.InputToGameKeyMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().GetInputMapping()
    this.GameKeyFactory = GameKeyFactory_GetInstance()
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    this.DownGameKeyEventHandler = DownGameKeyEventHandler_GetInstance()
    this.UpGameKeyEventHandler = UpGameKeyEventHandler_GetInstance()
    
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    
    
    this.MenuInputProcessor = NoMenuInputProcessor_GetInstance()
    this.MenuPaintable = NullPaintable_GetInstance()
    this.MenuForm = PaintableForm_GetNullPaintableForm()
    this.IsSingleKeyRepeatableProcessing = Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS)
    this.FontHeight = 0
}


func NewGameCommandCanvas(cmdListener CommandListener, name string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *GameCommandCanvas {
    this := &GameCommandCanvas{}
        this.MyCanvas = *NewMyCanvas(name, CanvasStrings_GetInstance().EMPTY_CHILD_NAME_LIST)
        this.initFields()
        this.RepaintBehavior = RepaintBehavior_GetInstance()
        this.ForegroundBasicColor = foregroundBasicColor
        this.BackgroundBasicColor = backgroundBasicColor
        this.ForegroundColor = this.ForegroundBasicColor.IntValue()
        this.BackgroundColor = this.BackgroundBasicColor.IntValue()
        this.InitCommands(cmdListener)
        if cmdListener != NullCommandListener_NULL_COMMAND_LISTENER {
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                panic(errors.New(fmt.Sprint()))
                    }
                }()
                this.InitMenu()
                this.RepaintProcessor.Process()
                return
            }(); tryControl1 == 1 {
                return
            }
        }
    return this
}


func (this *GameCommandCanvas) UpdateMeasurement(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT, e)
                }
            }()
            var font Font = graphics.GetFont(); _ = font
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Append(this.CanvasStrings.FD_WIDTH).Appendint(MyFontProcessor_DefaultCharWidth(font)).Append(this.CanvasStrings.FD_HEIGHT).Appendint(font.GetHeight())), this, this.CanvasStrings.UPDATE_MEASUREMENT)
            this.FontHeight = font.GetHeight()
            var rectangle *Rectangle = this.CreateRectangle(this.MenuForm.Size()); _ = rectangle
            this.MenuForm.Init(rectangle, FormTypeFactory_GetInstance().VERTICAL_CENTER_FORM)
            this.Update()
            return
        }(); tryControl1 == 1 {
            return
        }
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *GameCommandCanvas) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *GameCommandCanvas) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Append(fmt.Sprint(DisplayInfoSingleton_GetInstance()))), this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT)
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *GameCommandCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.AddCommand(GameCommandsFactory_GetInstance().CLOSE_AND_SHOW_GAME_CANVAS)
        this.SetCommandListener(cmdListener)
}


// throws Exception
func (this *GameCommandCanvas) InitMenu() {
        var form IScrollSelectionForm = this.CreateForm(); _ = form
        this.MenuForm = form
        if form != ScrollSelectionFormNoneFactory_GetInstance() {
            this.SetMenuInputProcessor(NewImmediateCommandFormInputProcessor(NewBasicArrayListD(), -1, this, form))
            this.MenuPaintable = NewFormPaintable(form)
        }
        this.RepaintBehavior.OnChangeRepaint(this)
}


// throws Exception
func (this *GameCommandCanvas) CreateForm() IScrollSelectionForm {
        var items []IABCustomItem = NewCommandTextItemArrayFactory(NewAllCommandsVisitor()).GetInstance(this.GetCommandStack(), this.BackgroundBasicColor, this.ForegroundBasicColor); _ = items
        var rectangle *Rectangle = this.CreateRectangle(len(items)); _ = rectangle
        return CommandCurrentSelectionFormFactory_GetInstance(StringUtil_GetInstance().EMPTY_STRING, items, rectangle, FormTypeFactory_GetInstance().VERTICAL_CENTER_FORM, 15, false, this.BackgroundBasicColor, this.ForegroundBasicColor)
}


func (this *GameCommandCanvas) CreateRectangle(size int) *Rectangle {
        var height int = size * this.FontHeight; _ = height
        var startY int = (this.DisplayInfo.GetLastHeight() * 2 / 3) - height; _ = startY
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var rectangle *Rectangle = NewRectangle(pointFactory.CreateXY(30, startY), this.DisplayInfo.GetLastWidth() - 30, startY); _ = rectangle
        return rectangle
}


func (this *GameCommandCanvas) Open() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "open")
        BasicMotionGesturesHandler_GetInstance().AddListenerInterface(this.GetMenuInputProcessor())
        GameKeyEventHandler_GetInstance().AddListener(this.GetMenuInputProcessor())
        DisplayChangeEventHandler_GetInstance().AddListenerInterface(this)
}


// throws Exception
func (this *GameCommandCanvas) Close() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CLOSE)
        BasicMotionGesturesHandler_GetInstance().RemoveListener(this.GetMenuInputProcessor())
        GameKeyEventHandler_GetInstance().RemoveListener(this.GetMenuInputProcessor())
        DisplayChangeEventHandler_GetInstance().RemoveListener(this)
}


// throws Exception
func (this *GameCommandCanvas) Update() {
        this.RepaintProcessor.Process()
}


func (this *GameCommandCanvas) GetSourceId() int {
        return GameCommandCanvas_Id
}


func (this *GameCommandCanvas) KeyPressed(keyCode int) {
        this.KeyPressedByDevice(keyCode, 0)
}


func (this *GameCommandCanvas) KeyReleased(keyCode int) {
        this.KeyReleasedByDevice(keyCode, 0)
}


func (this *GameCommandCanvas) KeyRepeated(keyCode int) {
        this.KeyRepeatedByDevice(keyCode, 0)
}


func (this *GameCommandCanvas) KeyPressedByDevice(keyCode int, deviceId int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Appendint(keyCode)), this, this.GameInputStrings.KEY_PRESSED)
        this.AddGameKeyEvent(keyCode, 0, false)
}


func (this *GameCommandCanvas) KeyReleasedByDevice(keyCode int, deviceId int) {
        this.RemoveGameKeyEvent(keyCode, deviceId, false)
}


func (this *GameCommandCanvas) KeyRepeatedByDevice(keyCode int, deviceId int) {
        if this.IsSingleKeyRepeatableProcessing {
            this.AddGameKeyEvent(keyCode, deviceId, true)
        }
}


func (this *GameCommandCanvas) AddGameKeyEvent(keyCode int, deviceId int, repeated bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.ADD_KEY_EVENT, e)
                }
            }()
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(this, keyCode); _ = gameKey
            if gameKey != this.GameKeyFactory.NONE {
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this, gameKey); _ = gameKeyEvent
                this.DownGameKeyEventHandler.FireEvent(gameKeyEvent)
                this.DownGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
            } else {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.ADD_KEY_EVENT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameCommandCanvas) RemoveGameKeyEvent(keyCode int, deviceId int, repeated bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.REMOVE_KEY_EVENT, e)
                }
            }()
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(this, keyCode); _ = gameKey
            if gameKey != this.GameKeyFactory.NONE {
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this, gameKey); _ = gameKeyEvent
                this.UpGameKeyEventHandler.FireEvent(gameKeyEvent)
                this.UpGameKeyEventHandler.GetInstanceForDevice(deviceId).FireEvent(gameKeyEvent)
            } else {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.REMOVE_KEY_EVENT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameCommandCanvas) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        this.MenuPaintable.Paint(graphics)
        this.RepaintBehavior.Repaint(this)
}


func (this *GameCommandCanvas) SetMenuInputProcessor(menuInputProcessor IBasicMenuInputProcessor) {
        this.MenuInputProcessor = menuInputProcessor
}


func (this *GameCommandCanvas) GetMenuInputProcessor() IBasicMenuInputProcessor {
        return this.MenuInputProcessor
}

