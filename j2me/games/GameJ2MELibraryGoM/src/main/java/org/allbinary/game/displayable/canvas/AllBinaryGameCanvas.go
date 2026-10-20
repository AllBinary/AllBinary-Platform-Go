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

type AllBinaryGameCanvas struct {
    RunnableCanvas
    BasicColorFactory *BasicColorFactory
    BasicSetColorUtil IBasicColorSetUtil
    TouchFeatureFactory *TouchFeatureFactory
    OpenGLFeatureFactory *OpenGLFeatureFactory
    TouchButtonFactory *TouchButtonFactory
    GameAdStateFactory *GameAdStateFactory
    GameStrings *GameStrings
    GameInputStrings *GameInputStrings
    GameStateFactory *GameStateFactory
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    GameTickDisplayInfoSingleton *GameTickDisplayInfoSingleton
    GameRunnable *GameCanvasRunnable
    GamePauseRunnable *GameCanvasPauseRunnable
    FormUtil *FormUtil
    MyFormUtil *MyFormUtil
    UpdateMyFontProcessor IMyFontProcessor
    MyFontProcessor IMyFontProcessor
    GameSpecificPaintable IPaintable
    SensorGameUpdateProcessor ISensorGameUpdateProcessor
    EndGameInfo IEndGameInfo
    StartIntermissionInterface IntermissionInterface
    EndLevelIntermissionInterface IntermissionInterface
    GameLayerManager IAllBinaryGameLayerManager
    GameState *GameState
    GameOver bool
    Initialized bool
    GameStateTimeHelper ITimeDelayHelper
    HighScoresHelper HighScoresHelperBaseInterface
    HighScoreSubmitted bool
    RealHighScoresPaintable *HighScoresPaintable
    HighScoresPaintable IPaintable
    LocalPlayerGameInputList IBasicArrayList
    IsCheating bool
    Hashtable *ABHashtable[any, any]
    IsSingleKeyRepeatableProcessing bool
    GameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory
    TouchButtonsPaintable IPaintable
    TouchPaintable IPaintable
    CheatProcessor IPlayerGameInput
    GameInputProcessor IProcessor
    EndGameProcessor IProcessor
    RealEndGameProcessor IProcessor
    StartIntermissionProcessor IProcessor
    RealStartIntermissionProcessor IProcessor
    EndGamePaintable IPaintable
    EndGameStatePaintable IPaintable
    NonBotPaintable IPaintable
    IntermissionPaintable IPaintable
    StartIntermissionPaintable IInitUpdatePaintable
    MainStateProcessor IProcessor
    ProcessGameProcessor IProcessor
    HighScoresFactoryInterface HighScoresFactoryInterface
    StartLevel int
    MainMenuInputProcessor IBasicMenuInputProcessor
    PopupMenuInputProcessor IBasicMenuInputProcessor
    MenuInputProcessor IBasicMenuInputProcessor
    MenuForm IScrollSelectionForm
    FormPaintable IPaintable
    OpenMenuPaintable IPaintable
    MenuPaintable IPaintable
    CurrentTouchInputFactory IBaseTouchInput
    ColorFillPaintable IColorFillBasePaintable
    BaseGameStatistics IBaseGameStatistics
    PrimaryPlayerQueue *PlayerQueue
    SecondaryPlayerQueue *PlayerQueue
    BUILD_GAME string
    GameTypeFactory *GameTypeFactory
    ScreenCapture IScreenCapture
    BasicMotionGesturesHandler *BasicMotionGesturesHandler
    GameKeyEventHandler *GameKeyEventHandler
    GameBehavior IDemoGameBehavior
    MenuBehavior IBaseMenuBehavior
    ProgressPaintable PaintableInterface
    FontHeightP int
    GameCanvasStartListener DemoPaintableInterface
    RawGameInputProcessor IInputProcessor
    RawInputProcessor IInputProcessor
    InputProcessor IInputProcessor
    YIELD_SLEEP int64
}

// Implements AllBinaryGameCanvasInterface, GameCanvasRunnableInterface, MenuListener, IntermissionCompositeInterface, IntermissionEnableListenerInterface, PopupMenuInterface, DisplayChangeEventListener, UpdateMyFontInterface

// The Java class AllBinaryGameCanvas has subclasses so its variables are this interface
type IAllBinaryGameCanvas interface {
    IRunnableCanvas
    AllBinaryGameCanvas_() *AllBinaryGameCanvas
    UpdateMeasurement(graphics IGraphics)
    GetInGameMenuBehavior() IBaseMenuBehavior
    SetCurrentThread()
    OnEvent(eventObject IAllBinaryEventObject)
    OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent)
    UpdateMenuFromEvent(displayChangeEvent *DisplayChangeEvent)
    ProcessSleep()
    InitPopupMenu()
    InitMenu()
    InitMenu2()
    UpdateMenu()
    UpdateMenu2()
    Pause()
    UnPause()
    IsPausable() bool
    PopupMenu()
    PopupMenu2()
    ToggleMenu()
    CloseMenu()
    CloseMenu2()
    Open()
    Close()
    ProcessorInit()
    SetProcessGameProcessorInit()
    InitSpecialPaint()
    Init(gameLayerManager IAllBinaryGameLayerManager, buffered bool)
    NotifyIntermission(enable bool)
    MediaInit()
    InitConfigurable(abeClientInformation AbeClientInformationInterface)
    InitApp(abeClientInformation AbeClientInformationInterface)
    InitTouch()
    UpdateCurrentTouchInputFactory(nextTouchInput IBaseTouchInput)
    UpdateTouch()
    UpdateTouch2()
    PostInitTouch()
    InitCommands(cmdListener CommandListener)
    AddCommands()
    ItemStateChanged(item IItem)
    UpdateScreenButtonPaintable()
    UpdateScreenButtonPaintable2()
    GetLayerManager() IAllBinaryGameLayerManager
    SetLayerManager(layerManager IAllBinaryGameLayerManager)
    IsGameOver() bool
    SetGameOver(gameOver bool)
    ProcessGameOver()
    IsHighScoreSubmitted() bool
    SetHighScoreSubmitted(highScoreSubmitted bool)
    GetGameState() *GameState
    SetGameState(gameState *GameState)
    UpdateGameState()
    UpdateGameKeyEventProcessor()
    RemoveAllGameKeyInputListenersOnBuild()
    RemoveAllGameKeyInputListeners()
    RemoveAllGameKeyInputListeners2()
    RemoveKeyInputListener(playerGameInput IPlayerGameInput)
    UpdateEndGameProcessor()
    UpdateEndGameProcessor2()
    BuildGameInit(isPortion bool)
    CleanupGame()
    LoadResourceLoadingLevel(resourceLoadingLevel *ResourceLoadingLevel)
    LoadResources(level int)
    UpdateColor()
    BuildGame(portion int)
    BuildGame2()
    AddKeyInputListener(playerGameInput IPlayerGameInput)
    SetGameCanvasStartListener(gameCanvasStartListener DemoPaintableInterface)
    LoadState()
    GetLoadStateHashtable() *ABHashtable[any, any]
    SetLoadStateHashtable(hashtable *ABHashtable[any, any])
    GetCurrentStateHashtable() *ABHashtable[any, any]
    PaintGameOver(graphics IGraphics)
    Draw(graphics IGraphics)
    Clear(graphics IGraphics)
    Paint(graphics IGraphics)
    PaintThreed(graphics IGraphics)
    ProcessEndLevelIntermissionGameState()
    NonBotPaint(graphics IGraphics)
    PaintIntermission(graphics IGraphics)
    KeyPressed(keyCode int)
    KeyReleased(keyCode int)
    KeyRepeated(keyCode int)
    KeyPressedByDevice(keyCode int, deviceId int)
    KeyRepeatedByDevice(keyCode int, deviceId int)
    KeyReleasedByDevice(keyCode int, deviceId int)
    HandleRawKey(keyCode int, deviceId int, repeated bool)
    EndProgress(isProgress bool) int
    ProcessPlayingGame()
    ThreadInit()
    ProcessGame()
    NotifyDonePainting()
    ShouldWait()
    Run()
    Run2()
    Run3()
    SetRunning(running bool)
    End()
    EndGameThread()
    CreateHighScore(score int64) *HighScore
    SetHighScore(abeClientInformation AbeClientInformationInterface, name string, score int64, autoSubmit bool, isLast bool)
    SetHighScore2(abeClientInformation AbeClientInformationInterface, name string, score int64, autoSubmit bool, isLast bool)
    GetGameStateTimeHelper() ITimeDelayHelper
    SetInitialized(initialized bool)
    IsInitialized() bool
    GetSourceId() int
    SetProgressPaintable(paintable PaintableInterface)
    GetStartIntermissionInterface() IntermissionInterface
    GetEndLevelIntermissionInterface() IntermissionInterface
    SetTouchPaintableP(paintable IPaintable)
    GetTouchPaintableP() IPaintable
    SetEndGamePaintable(endGamePaintable IPaintable)
    GetEndGamePaintable() IPaintable
    SetIntermissionPaintable(intermissionPaintable IPaintable)
    GetIntermissionPaintable() IPaintable
    GetEndGameInfoP() IEndGameInfo
    SetHighScoresPaintable(highScoresPaintable IPaintable)
    GetHighScoresPaintable() IPaintable
    GetRealHighScoresPaintable() *HighScoresPaintable
    ClearPlayerGameInputList()
    AddPlayerGameInput(playerGameInput IPlayerGameInput)
    SetMenuInputProcessor(menuInputProcessor IBasicMenuInputProcessor)
    GetMenuInputProcessor() IBasicMenuInputProcessor
    GetMenuForm() IScrollSelectionForm
    SetMenuForm(menuForm IScrollSelectionForm)
    GetStartLevel() int
    SetStartLevel(startLevel int)
    SetTouchButtonsPaintable(touchButtonsPaintable IPaintable)
    GetTouchButtonsPaintable() IPaintable
    SetGameInputProcessor(gameInputProcessor IProcessor)
    GetGameInputProcessor() IProcessor
    SetEndGameProcessor(endGameProcessor IProcessor)
    GetEndGameProcessor() IProcessor
    SetEndGameStatePaintable(endGameStatePaintable IPaintable)
    GetEndGameStatePaintable() IPaintable
    SetNonBotPaintableP(nonBotPaintable IPaintable)
    GetNonBotPaintableP() IPaintable
    SetStartIntermissionPaintable(startIntermissionPaintable IInitUpdatePaintable)
    GetStartIntermissionPaintable() IInitUpdatePaintable
    SetMainStateProcessor(mainStateProcessor IProcessor)
    GetMainStateProcessor() IProcessor
    SetProcessGameProcessor(processGameProcessor IProcessor)
    GetProcessGameProcessor() IProcessor
    SetOpenMenuPaintable(openMenuPaintable IPaintable)
    GetOpenMenuPaintable() IPaintable
    SetPopupMenuInputProcessor(popupMenuInputProcessor IBasicMenuInputProcessor)
    GetPopupMenuInputProcessor() IBasicMenuInputProcessor
    GetSensorGameUpdateProcessor() ISensorGameUpdateProcessor
    GetRawGameInputProcessor() IInputProcessor
    GetRawInputProcessor() IInputProcessor
    SetInputProcessor(inputProcessor IInputProcessor)
    GetInputProcessor() IInputProcessor
    SetMenuPaintable(menuPaintable IPaintable)
    GetMenuPaintable() IPaintable
    SetFormPaintable(formPaintable IPaintable)
    GetFormPaintable() IPaintable
    SetGameSpecificPaintableP(gameSpecificPaintable IPaintable)
    GetGameSpecificPaintableP() IPaintable
    IsSingleThread() bool
    IsRunningInAnotherThread() bool
    GetType() int
}

func (this *AllBinaryGameCanvas) AllBinaryGameCanvas_() *AllBinaryGameCanvas {
    return this
}

var AllBinaryGameCanvas_Id int = 0
var AllBinaryGameCanvas_TYPE int = 2

func (this *AllBinaryGameCanvas) initFields() {
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    this.TouchFeatureFactory = TouchFeatureFactory_GetInstance()
    this.OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance()
    this.TouchButtonFactory = TouchButtonFactory_GetInstance()
    this.GameAdStateFactory = GameAdStateFactory_GetInstance()
    this.GameStrings = GameStrings_GetInstance()
    this.GameInputStrings = GameInputStrings_GetInstance()
    this.GameStateFactory = GameStateFactory_GetInstance()
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.GameTickDisplayInfoSingleton = GameTickDisplayInfoSingleton_GetInstance()
    this.GameRunnable = NewGameCanvasRunnable(this)
    this.GamePauseRunnable = NewGameCanvasPauseRunnable(this)
    this.FormUtil = FormUtil_GetInstance()
    this.MyFormUtil = MyFormUtil_GetInstance()
    this.UpdateMyFontProcessor = NewUpdateMyFontProcessor(this)
    this.MyFontProcessor = this.UpdateMyFontProcessor
    this.GameSpecificPaintable = NullPaintable_GetInstance()
    this.SensorGameUpdateProcessor = NewSingleSensorGameUpdateProcessor()
    this.EndGameInfo = NewEndGameInfo()
    this.StartIntermissionInterface = NewIntermission()
    this.EndLevelIntermissionInterface = NewIntermission()
    this.GameLayerManager = AllBinaryGameLayerManager_GetNullInstance()
    this.GameState = this.GameStateFactory.NO_GAME_STATE
    
    
    this.GameStateTimeHelper = NewTimeDelayHelper(0)
    
    
    this.RealHighScoresPaintable = NewHighScoresPaintable()
    this.HighScoresPaintable = NullPaintable_GetInstance()
    this.LocalPlayerGameInputList = NewBasicArrayListD()
    
    this.Hashtable = this.StdUtil.NULL_TABLE
    
    this.GameInitializationInterfaceFactoryInterface = BasicBuildGameInitializerFactory_NULL_BASE_BUILD_GMAE_INITIALIZER_FACTORY
    this.TouchButtonsPaintable = NullPaintable_GetInstance()
    this.TouchPaintable = NullPaintable_GetInstance()
    this.CheatProcessor = NoPlayerGameInput_GetInstance()
    this.GameInputProcessor = Processor_GetInstance()
    this.EndGameProcessor = Processor_GetInstance()
    this.RealEndGameProcessor = Processor_GetInstance()
    this.StartIntermissionProcessor = Processor_GetInstance()
    this.RealStartIntermissionProcessor = Processor_GetInstance()
    this.EndGamePaintable = NullPaintable_GetInstance()
    this.EndGameStatePaintable = NullPaintable_GetInstance()
    this.NonBotPaintable = NullPaintable_GetInstance()
    this.IntermissionPaintable = NullPaintable_GetInstance()
    this.StartIntermissionPaintable = NullInitUpdatePaintable_GetInstance()
    this.MainStateProcessor = Processor_GetInstance()
    this.ProcessGameProcessor = Processor_GetInstance()
    
    
    this.MainMenuInputProcessor = NoMenuInputProcessor_GetInstance()
    this.PopupMenuInputProcessor = NoMenuInputProcessor_GetInstance()
    this.MenuInputProcessor = NoMenuInputProcessor_GetInstance()
    this.MenuForm = ScrollSelectionForm_GetNullScrollSelectionForm()
    this.FormPaintable = NullPaintable_GetInstance()
    this.OpenMenuPaintable = NullPaintable_GetInstance()
    this.MenuPaintable = NullPaintable_GetInstance()
    this.CurrentTouchInputFactory = NoButtonsTouchInputFactory_GetInstance()
    this.ColorFillPaintable = ColorFillPaintableFactory_GetInstance().GetInstance(this.BasicColorFactory.BLACK, true)
    this.BaseGameStatistics = GameStatisticsFactory_GetInstance()
    this.PrimaryPlayerQueue = PrimaryPlayerQueueFactory_GetInstance()
    this.SecondaryPlayerQueue = SecondaryPlayerQueueFactory_GetInstance()
    this.BUILD_GAME = "buildGame"
    this.GameTypeFactory = GameTypeFactory_GetInstance()
    this.ScreenCapture = ScreenCaptureFactory_GetInstance()
    this.BasicMotionGesturesHandler = BasicMotionGesturesHandler_GetInstance()
    this.GameKeyEventHandler = GameKeyEventHandler_GetInstance()
    
    
    this.ProgressPaintable = ProgressCanvasFactory_GetLazyInstance()
    
    this.GameCanvasStartListener = NullDemoPaintable_NULL_DEMO_PAINTABLE
    this.RawGameInputProcessor = NewGameCanvasInputProcessor(this)
    this.RawInputProcessor = NewFormInputProcessor(this)
    this.InputProcessor = this.GetRawGameInputProcessor()
    this.YIELD_SLEEP = 100
}


func NewAllBinaryGameCanvas(commandListener CommandListener, gameLayerManager IAllBinaryGameLayerManager, highScoresFactoryInterface HighScoresFactoryInterface, gameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory, buffered bool) *AllBinaryGameCanvas {
    this := &AllBinaryGameCanvas{}
        this.RunnableCanvas = *NewRunnableCanvas(commandListener, CanvasStrings_GetInstance().EMPTY_CHILD_NAME_LIST, true)
        this.initFields()
        this.HighScoresHelper = highScoresFactoryInterface.CreateHighScoresHelper()
        this.GameInitializationInterfaceFactoryInterface = gameInitializationInterfaceFactoryInterface
        this.Init(gameLayerManager, buffered)
        this.HighScoresFactoryInterface = highScoresFactoryInterface
        if this.GameLayerManager.GetGameInfo().GetGameType() == this.GameTypeFactory.BOT || this.GameLayerManager.GetGameInfo().GetGameType() == this.GameTypeFactory.NONE {
            this.GameBehavior = DemoGameBehavior_GetInstance()
            this.MenuBehavior = BaseMenuBehavior_GetInstance()
        } else {
            this.GameBehavior = BaseGameBehavior_GetInstance()
            this.MenuBehavior = this.GetInGameMenuBehavior()
        }
        this.InitPopupMenu()
        this.InitMenu()
        this.InitSpecialPaint()
        DisplayChangeEventHandler_GetInstance().AddListenerInterface(this)
    return this
}


func (this *AllBinaryGameCanvas) UpdateMeasurement(graphics IGraphics) {
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
            this.FontHeightP = font.GetHeight()
            this.MyFormUtil.UpdateMeasurement(graphics)
            var popupMenuRectangle *Rectangle = this.MyFormUtil.GetPopupMenuRectangle(); _ = popupMenuRectangle
            var basicPopupMenuPaintable *BasicPopupMenuPaintable = (any(this.GetOpenMenuPaintable()).(*BasicPopupMenuPaintable)); _ = basicPopupMenuPaintable
            basicPopupMenuPaintable.Init(popupMenuRectangle)
            if this.GetPopupMenuInputProcessor() != NoMenuInputProcessor_GetInstance() {
                var popupMenuInputProcessor *PopupMenuInputProcessor = (any(this.GetPopupMenuInputProcessor()).(*PopupMenuInputProcessor)); _ = popupMenuInputProcessor
                popupMenuInputProcessor.Init(popupMenuRectangle)
            }
            this.MyFontProcessor = MyFontProcessor_GetInstance()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AllBinaryGameCanvas) GetInGameMenuBehavior() IBaseMenuBehavior {
        return InGameMenuBehavior_GetInstance()
}


func (this *AllBinaryGameCanvas) SetCurrentThread() {
        if J2MEUtil_IsHTML() {
            this.RunnableCanvas.SetCurrentThreadFake()
        } else {
            this.RunnableCanvas.SetCurrentThread()
        }
}


func (this *AllBinaryGameCanvas) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *AllBinaryGameCanvas) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT, e)
                }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Append(fmt.Sprint(DisplayInfoSingleton_GetInstance()))), this, this.CanvasStrings.UPDATE_MEASUREMENT)
            this.MyFontProcessor = this.UpdateMyFontProcessor
            this.MenuBehavior.OnDisplayChangeEvent(this, displayChangeEvent)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateMenuFromEvent(displayChangeEvent *DisplayChangeEvent) {
        var formType *FormType = FormTypeFactory_GetInstance().GetFormType(); _ = formType
        var rectangle *Rectangle = this.FormUtil.CreateFormRectangle(); _ = rectangle
        this.MenuForm.Init(rectangle, formType)
        if this.CurrentTouchInputFactory != NoButtonsTouchInputFactory_GetInstance() {
            this.TouchButtonFactory.Toggle(this.IsPaused(), this.CurrentTouchInputFactory.GetList())
        }
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


// throws Exception
func (this *AllBinaryGameCanvas) ProcessSleep() {
        this.RunnableCanvas.ProcessSleep()
        if AllBinaryMediaManager_Update() {
            if !this.PrimaryPlayerQueue.Process() {
                this.SecondaryPlayerQueue.Process()
            }
        }
        if this.MenuInputProcessor.ProcessInputList() != -1 {
            this.RunnableCanvasRefreshHelper.Process()
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) InitPopupMenu() {
        if this.PopupMenuInputProcessor == NoMenuInputProcessor_GetInstance() {
            var features *Features = Features_GetInstance(); _ = features
            var popupMenuRectangle *Rectangle = this.MyFormUtil.GetPopupMenuRectangle(); _ = popupMenuRectangle
            if features.IsFeature(this.TouchFeatureFactory.TOUCH_ENABLED) {
                this.SetOpenMenuPaintable(NewBasicPopupMenuPaintable(popupMenuRectangle, this.GameLayerManager.GetBackgroundBasicColor(), this.GameLayerManager.GetForegroundBasicColor()))
                this.SetPopupMenuInputProcessor(NewPopupMenuInputProcessor(NewBasicArrayListD(), -1, this, popupMenuRectangle))
            }
        } else {
        }
}


func (this *AllBinaryGameCanvas) InitMenu() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "initMenu", e)
                }
            }()
            this.MenuBehavior.InitMenu(this)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) InitMenu2() {
        this.CloseMenu()
        var formType *FormType = FormTypeFactory_GetInstance().GetFormType(); _ = formType
        var gameLimitedCommandTextItemArrayFactory *GameLimitedCommandTextItemArrayFactory = GameLimitedCommandTextItemArrayFactory_GetInstance(); _ = gameLimitedCommandTextItemArrayFactory
        var commandTextItemArrayFactory *CommandTextItemArrayFactory = gameLimitedCommandTextItemArrayFactory.GetCommandTextItemArrayFactory(); _ = commandTextItemArrayFactory
        var items []IABCustomItem = commandTextItemArrayFactory.GetInstance(this.GetCommandStack(), this.GameLayerManager.GetBackgroundBasicColor(), this.GameLayerManager.GetForegroundBasicColor()); _ = items
        var rectangle *Rectangle = this.FormUtil.CreateFormRectangle(); _ = rectangle
        this.SetMenuForm(CommandCurrentSelectionFormFactory_GetInstance(StringUtil_GetInstance().EMPTY_STRING, items, rectangle, formType, 25, false, this.GameLayerManager.GetBackgroundBasicColor(), this.GameLayerManager.GetForegroundBasicColor()))
        var scrollSelectionForm IScrollSelectionForm = this.GetMenuForm(); _ = scrollSelectionForm
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(this.TouchFeatureFactory.TOUCH_ENABLED) {
            this.MainMenuInputProcessor = NewPopupCommandFormInputProcessor(NewBasicArrayListD(), -1, this, scrollSelectionForm, any(this.GetPopupMenuInputProcessor()).(*PopupMenuInputProcessor))
        } else {
        }
        this.SetMenuInputProcessor(this.GetPopupMenuInputProcessor())
        if scrollSelectionForm != ScrollSelectionFormNoneFactory_GetInstance() {
            this.SetFormPaintable(NewFormPaintable(scrollSelectionForm))
        } else {
        }
        this.CloseMenu()
}


func (this *AllBinaryGameCanvas) UpdateMenu() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "initMenu", e)
                }
            }()
            this.MenuBehavior.UpdateMenu(this)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateMenu2() {
        var scrollSelectionForm IScrollSelectionForm = this.GetMenuForm(); _ = scrollSelectionForm
        scrollSelectionForm.DeleteAll()
        var gameLimitedCommandTextItemArrayFactory *GameLimitedCommandTextItemArrayFactory = GameLimitedCommandTextItemArrayFactory_GetInstance(); _ = gameLimitedCommandTextItemArrayFactory
        var commandTextItemArrayFactory *CommandTextItemArrayFactory = gameLimitedCommandTextItemArrayFactory.GetCommandTextItemArrayFactory(); _ = commandTextItemArrayFactory
        var items []IABCustomItem = commandTextItemArrayFactory.GetInstance(this.GetCommandStack(), this.GameLayerManager.GetBackgroundBasicColor(), this.GameLayerManager.GetForegroundBasicColor()); _ = items
        var size int = len(items); _ = size
        for index := int(0); index < size; index++ {
            scrollSelectionForm.Append(items[index])
        }
        var formType *FormType = FormTypeFactory_GetInstance().GetFormType(); _ = formType
        var rectangle *Rectangle = this.FormUtil.CreateFormRectangle(); _ = rectangle
        scrollSelectionForm.Init(rectangle, formType)
}


// synchronized
func (this *AllBinaryGameCanvas) Pause() {
        this.GameBehavior.Pause(this)
        this.CloseMenu()
        this.RunnableCanvas.Pause()
        this.TouchButtonFactory.Toggle(this.IsPaused(), BasicArrayListUtil_GetInstance().GetImmutableInstance())
        this.TsUtil.Gc()
}


// synchronized
func (this *AllBinaryGameCanvas) UnPause() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.GameStrings.UNPAUSE)
        this.CloseMenu()
        this.TsUtil.Gc()
        this.RunnableCanvas.UnPause()
        this.TouchButtonFactory.Toggle(this.IsPaused(), BasicArrayListUtil_GetInstance().GetImmutableInstance())
        this.GameBehavior.UnPause(this)
}


func (this *AllBinaryGameCanvas) IsPausable() bool {
        if CurrentDisplayableFactory_GetInstance().GetUsedRunnable() == NullWaitGameRunnable_GetInstance() {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) PopupMenu() {
        this.MenuBehavior.PopupMenu(this)
}


// throws Exception
func (this *AllBinaryGameCanvas) PopupMenu2() {
        this.PrimaryPlayerQueue.Add(SelectSound_GetInstance())
        this.SetMenuPaintable(this.GetFormPaintable())
        this.SetMenuInputProcessor(this.MainMenuInputProcessor)
        this.BasicMotionGesturesHandler.AddListenerInterface(this.MainMenuInputProcessor)
        this.GameKeyEventHandler.AddListener(this.MainMenuInputProcessor)
}


// throws Exception
func (this *AllBinaryGameCanvas) ToggleMenu() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.GameStrings.TOGGLE_MENU)
        if this.GetMenuPaintable() == this.GetOpenMenuPaintable() {
            this.Pause()
            this.UpdateMenu()
            this.PopupMenu()
        } else if this.GetMenuPaintable() == this.GetFormPaintable() {
            this.PrimaryPlayerQueue.Add(SelectSound_GetInstance())
            this.UnPause()
        }
}


func (this *AllBinaryGameCanvas) CloseMenu() {
        this.MenuBehavior.CloseMenu(this)
}


func (this *AllBinaryGameCanvas) CloseMenu2() {
        this.SetMenuPaintable(this.GetOpenMenuPaintable())
        this.BasicMotionGesturesHandler.RemoveListener(this.MainMenuInputProcessor)
        this.GameKeyEventHandler.RemoveListener(this.MainMenuInputProcessor)
        this.SetMenuInputProcessor(this.GetPopupMenuInputProcessor())
}


func (this *AllBinaryGameCanvas) Open() {
        this.BasicMotionGesturesHandler.AddListenerInterface(this.MenuInputProcessor)
        this.GameKeyEventHandler.AddListener(this.MenuInputProcessor)
}


func (this *AllBinaryGameCanvas) Close() {
        this.BasicMotionGesturesHandler.RemoveListener(this.MenuInputProcessor)
        this.GameKeyEventHandler.RemoveListener(this.MenuInputProcessor)
        this.PrimaryPlayerQueue.Clear()
        this.SecondaryPlayerQueue.Clear()
}


// throws Exception
func (this *AllBinaryGameCanvas) ProcessorInit() {
        this.SetMainStateProcessor(Processor_GetInstance())
        this.SetProcessGameProcessorInit()
}


func (this *AllBinaryGameCanvas) SetProcessGameProcessorInit() {
        this.SetProcessGameProcessor(NewGameProcessor(this))
}


func (this *AllBinaryGameCanvas) InitSpecialPaint() {
        this.MenuBehavior.InitSpecialPaint(this)
}


// throws Exception
func (this *AllBinaryGameCanvas) Init(gameLayerManager IAllBinaryGameLayerManager, buffered bool) {
        this.GameStateTimeHelper.SetStartTimeTNT()
        this.SetHighScoreSubmitted(false)
        this.SetLayerManager(gameLayerManager)
        if !buffered {
            var BUFF_MESSAGE string = "XXX Not Buffering Causes Concurrency Issues XXX"; _ = BUFF_MESSAGE
            this.LogUtil.PutF(BUFF_MESSAGE, this, this.CommonStrings.CONSTRUCTOR)
        } else {
            panic(errors.New(fmt.Sprint("Buffering is disabled")))
        }
        this.SetGameInputProcessor(Processor_GetInstance())
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(GameFeatureFactory_GetInstance().CHEATING) {
            this.IsCheating = true
            this.CheatProcessor = NewCheatGameInputProcessor(this)
            this.GameKeyEventHandler.AddListener(this.CheatProcessor)
        } else {
            this.IsCheating = false
            this.CheatProcessor = NoPlayerGameInput_GetInstance()
        }
        this.RealEndGameProcessor = NewEndGameProcessor(this)
        this.SetEndGameProcessor(Processor_GetInstance())
        this.RealStartIntermissionProcessor = NewStartIntermissionProcessor(this)
        this.StartIntermissionProcessor = Processor_GetInstance()
        this.SetEndGameStatePaintable(NewEndGamePaintable(this))
        this.SetEndGamePaintable(NullPaintable_GetInstance())
        this.SetIntermissionPaintable(NullPaintable_GetInstance())
        this.GetStartIntermissionInterface().SetListener(this)
}


func (this *AllBinaryGameCanvas) NotifyIntermission(enable bool) {
        if enable {
            this.StartIntermissionProcessor = this.RealStartIntermissionProcessor
        } else {
            this.StartIntermissionProcessor = Processor_GetInstance()
        }
        if enable {
            this.SetIntermissionPaintable(this.GetStartIntermissionPaintable())
        } else {
            this.SetIntermissionPaintable(NullPaintable_GetInstance())
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) MediaInit() {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}


// throws Exception
// synchronized
func (this *AllBinaryGameCanvas) InitConfigurable(abeClientInformation AbeClientInformationInterface) {
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.AddNormalPortion(50, "Setting Configurables")
        GameInitializationUtil_GetInstance().InitGame(abeClientInformation, this, this.GameInitializationInterfaceFactoryInterface)
        var gameFeatureFactory *GameFeatureFactory = GameFeatureFactory_GetInstance(); _ = gameFeatureFactory
        var changedGameFeatureListener *ChangedGameFeatureListener = ChangedGameFeatureListener_GetInstance(); _ = changedGameFeatureListener
        if changedGameFeatureListener.IsChangedFeature(gameFeatureFactory.SOUND) {
            var features *Features = Features_GetInstance(); _ = features
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Sound Changing To: ").Appendboolean(features.IsFeature(gameFeatureFactory.SOUND))), this, "initConfigurable")
            this.MediaInit()
            changedGameFeatureListener.Remove(gameFeatureFactory.SOUND)
        }
        var sensorFeatureFactory *SensorFeatureFactory = SensorFeatureFactory_GetInstance(); _ = sensorFeatureFactory
        if changedGameFeatureListener.IsChangedFeature(sensorFeatureFactory.ORIENTATION_SENSORS) || changedGameFeatureListener.IsChangedFeature(sensorFeatureFactory.NO_ORIENTATION) || changedGameFeatureListener.IsChangedFeature(sensorFeatureFactory.SIMULATED_ORIENTATION_SENSORS) {
            changedGameFeatureListener.Remove(sensorFeatureFactory.NO_ORIENTATION)
            changedGameFeatureListener.Remove(sensorFeatureFactory.ORIENTATION_SENSORS)
            changedGameFeatureListener.Remove(sensorFeatureFactory.SIMULATED_ORIENTATION_SENSORS)
        }
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        this.LoadResourceLoadingLevel(resourceLoadingLevelFactory.LOAD_GAME)
}


// throws Exception
func (this *AllBinaryGameCanvas) InitApp(abeClientInformation AbeClientInformationInterface) {
        this.InitConfigurable(abeClientInformation)
        this.ProcessorInit()
        this.SensorGameUpdateProcessor.Process(this.GameLayerManager)
        this.SensorGameUpdateProcessor.SendNotifications(this.GameLayerManager)
        this.InitTouch()
}


// throws Exception
func (this *AllBinaryGameCanvas) InitTouch() {
        var gameInitializedEvent *GameInitializedEvent = GameInitializationUtil_GetInstance().EVENT; _ = gameInitializedEvent
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        gameInitializedEvent.SetResourceLoadingLevel(resourceLoadingLevelFactory.LOAD_TOUCH)
        GameInitializedEventHandler_GetInstance().FireEvent(gameInitializedEvent)
        this.TouchButtonFactory.DefaultList()
        this.UpdateTouch()
        this.PostInitTouch()
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateCurrentTouchInputFactory(nextTouchInput IBaseTouchInput) {
        if nextTouchInput != this.CurrentTouchInputFactory {
            this.CurrentTouchInputFactory = nextTouchInput
            PreLogUtil_Put(fmt.Sprint(this.CurrentTouchInputFactory), this, this.GameInputStrings.UDPATE_CURRENT_TOUCH_INPUT_FACTORY)
            this.TouchButtonFactory.Toggle(this.IsPaused(), this.CurrentTouchInputFactory.GetList())
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateTouch() {
        this.GameBehavior.UpdateTouch(this)
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateTouch2() {
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(this.TouchFeatureFactory.AUTO_HIDE_SHOW_SCREEN_BUTTONS) {
            if this.GameLayerManager.GetGameInfo().GetCurrentLevel() - this.GetStartLevel() == 1 {
                this.SetTouchPaintableP(NullPaintable_GetInstance())
            }
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) PostInitTouch() {
        this.SetTouchButtonsPaintable(TouchButtonsPaintableFactory_GetInstance(this.GameLayerManager.GetForegroundBasicColor()))
        this.UpdateScreenButtonPaintable()
}


func (this *AllBinaryGameCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.AddCommands()
        this.SetCommandListener(cmdListener)
}


func (this *AllBinaryGameCanvas) AddCommands() {
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        var myCommandsFactory *MyCommandsFactory = MyCommandsFactory_GetInstance(); _ = myCommandsFactory
        if DebugFactory_GetInstance() != NoDebug_GetInstance() {
            this.AddCommand(gameCommandsFactory.START_TRACE)
        }
        this.AddCommand(gameCommandsFactory.RESTART_COMMAND)
        this.AddCommand(myCommandsFactory.PAUSE_COMMAND)
        this.AddCommand(gameCommandsFactory.QUIT_COMMAND)
        var isOverScan bool = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance().IsOverScan(); _ = isOverScan
        if J2MEUtil_IsHTML() {
        } else if SWTUtil_IsSWT {
        } else if !isOverScan {
            if TouchScreenFactory_GetInstance().IsTouch() && NewInGameFeatures().IsAny() {
                this.AddCommand(InGameOptionsForm_DISPLAY)
            }
            this.AddCommand(gameCommandsFactory.SAVE)
            this.AddCommand(gameCommandsFactory.DISPLAY_LOAD_FORM)
        }
}


func (this *AllBinaryGameCanvas) ItemStateChanged(item IItem) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "itemStateChanged", e)
                }
            }()
            var itemLabel string = item.GetLabel(); _ = itemLabel
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonLabels.ITEM_LABEL).Append(itemLabel)), this, "itemStateChanged")
            if func() bool { _, ok := any(item).(IChoiceGroup); return ok }() {
                var gameFeatureUtil *GameFeatureUtil = GameFeatureUtil_GetInstance(); _ = gameFeatureUtil
                if gameFeatureUtil.IsExclusive(itemLabel) {
                    gameFeatureUtil.UpdateExclusiveForChoiceGroup(any(item).(IChoiceGroup))
                } else {
                    gameFeatureUtil.UpdateMultiple(any(item).(IChoiceGroup))
                }
            }
            this.SensorGameUpdateProcessor.Process(this.GameLayerManager)
            this.UpdateTouch()
            this.UpdateScreenButtonPaintable()
            this.SensorGameUpdateProcessor.SendNotifications(this.GameLayerManager)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateScreenButtonPaintable() {
        this.GameBehavior.UpdateScreenButtonPaintable(this)
}


func (this *AllBinaryGameCanvas) UpdateScreenButtonPaintable2() {
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(this.TouchFeatureFactory.AUTO_HIDE_SHOW_SCREEN_BUTTONS) {
            this.SetTouchPaintableP(this.GetTouchButtonsPaintable())
            this.SetStartLevel(this.GameLayerManager.GetGameInfo().GetCurrentLevel())
        } else if features.IsFeature(this.TouchFeatureFactory.SHOW_SCREEN_BUTTONS) {
            this.SetTouchPaintableP(this.GetTouchButtonsPaintable())
        } else if features.IsFeature(this.TouchFeatureFactory.HIDE_SCREEN_BUTTONS) {
            this.SetTouchPaintableP(NullPaintable_GetInstance())
        } else {
            this.SetTouchPaintableP(NullPaintable_GetInstance())
        }
}


func (this *AllBinaryGameCanvas) GetLayerManager() IAllBinaryGameLayerManager {
        return this.GameLayerManager
}


func (this *AllBinaryGameCanvas) SetLayerManager(layerManager IAllBinaryGameLayerManager) {
        this.GameLayerManager = layerManager
}


// synchronized
func (this *AllBinaryGameCanvas) IsGameOver() bool {
        return this.GameOver
}


// synchronized
func (this *AllBinaryGameCanvas) SetGameOver(gameOver bool) {
        this.GameOver = gameOver
}


// throws Exception
func (this *AllBinaryGameCanvas) ProcessGameOver() {
        PreLogUtil_Put(this.CommonStrings.START, this, this.GameStrings.SET_GAME_OVER)
        this.SetGameOver(true)
        this.RemovePauseCommand()
        this.SetGameState(this.GameStateFactory.SHOW_END_RESULT_GAME_STATE)
        this.SetEndGamePaintable(this.GetEndGameStatePaintable())
}


func (this *AllBinaryGameCanvas) IsHighScoreSubmitted() bool {
        return this.HighScoreSubmitted
}


// throws Exception
func (this *AllBinaryGameCanvas) SetHighScoreSubmitted(highScoreSubmitted bool) {
        this.HighScoreSubmitted = highScoreSubmitted
        if highScoreSubmitted {
            this.HighScoresHelper.SetSelectedHighScores(NullHighScoresSingletonFactory_GetInstance())
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("isHighScoreSubmitted: ").Appendboolean(highScoreSubmitted)), this, "setHighScoreSubmitted")
}


func (this *AllBinaryGameCanvas) GetGameState() *GameState {
        return this.GameState
}


// throws Exception
func (this *AllBinaryGameCanvas) SetGameState(gameState *GameState) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameStrings.GAME_STATE).Append(this.StringUtil.ToString(gameState))), this, this.GameStrings.SET_GAME_STATE)
        this.GameState = gameState
        this.GameStateTimeHelper.SetStartTimeTNT()
        this.UpdateEndGameProcessor()
        this.UpdateGameKeyEventProcessor()
        this.GameBehavior.SetGameState(this)
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateGameState() {
        var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
        gameAdState.ProcessAdState(this.GameState, this.GameLayerManager.GetGameInfo().GetGameType())
        if this.GameState != this.GameStateFactory.PLAYING_GAME_STATE {
            gameAdState.ProcessPageAdState()
        }
}


func (this *AllBinaryGameCanvas) UpdateGameKeyEventProcessor() {
        if this.GetGameState() != this.GameStateFactory.PLAYING_GAME_STATE || this.IsCheating {
        } else {
        }
}


func (this *AllBinaryGameCanvas) RemoveAllGameKeyInputListenersOnBuild() {
        this.RemoveAllGameKeyInputListeners()
}


func (this *AllBinaryGameCanvas) RemoveAllGameKeyInputListeners() {
        this.GameBehavior.RemoveAllGameKeyInputListeners(this)
}


func (this *AllBinaryGameCanvas) RemoveAllGameKeyInputListeners2() {
        this.LogUtil.PutF("Remove PlayerInput Listeners", this, "removeAllGameKeyInputListeners")
        for index := int(this.LocalPlayerGameInputList.Size() - 1); index >= 0; index-- {
            var playerGameInput IPlayerGameInput = any(this.LocalPlayerGameInputList.Get(index)).(IPlayerGameInput); _ = playerGameInput
            this.RemoveKeyInputListener(playerGameInput)
            playerGameInput.RemoveNonAIInputGameKeyEvents()
        }
}


func (this *AllBinaryGameCanvas) RemoveKeyInputListener(playerGameInput IPlayerGameInput) {
        this.GameKeyEventHandler.RemoveListener(playerGameInput)
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateEndGameProcessor() {
        this.GameBehavior.UpdateEndGameProcessor(this)
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateEndGameProcessor2() {
        if this.GetGameState() == this.GameStateFactory.SHOW_END_RESULT_GAME_STATE || this.GetGameState() == this.GameStateFactory.SHOW_HIGH_SCORE_GAME_STATE {
            this.SetEndGameProcessor(this.RealEndGameProcessor)
        } else {
            this.SetEndGameProcessor(Processor_GetInstance())
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) BuildGameInit(isPortion bool) {
}


// throws Exception
func (this *AllBinaryGameCanvas) CleanupGame() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CLEANUP)
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.AddNormalPortion(10, "Cleaning Up")
        this.PrimaryPlayerQueue.Clear()
        this.SecondaryPlayerQueue.Clear()
        GameLevelDisplayChangeEventListenersFactory_GetInstance().Clear()
}


// throws Exception
func (this *AllBinaryGameCanvas) LoadResourceLoadingLevel(resourceLoadingLevel *ResourceLoadingLevel) {
        var gameInitializedEvent *GameInitializedEvent = GameInitializationUtil_GetInstance().EVENT; _ = gameInitializedEvent
        gameInitializedEvent.SetResourceLoadingLevel(resourceLoadingLevel)
        GameInitializedEventHandler_GetInstance().FireEvent(gameInitializedEvent)
}


// throws Exception
func (this *AllBinaryGameCanvas) LoadResources(level int) {
        var gameInitializedEvent *GameInitializedEvent = GameInitializationUtil_GetInstance().EVENT; _ = gameInitializedEvent
        gameInitializedEvent.SetLevel(level)
        GameInitializedEventHandler_GetInstance().FireEvent(gameInitializedEvent)
}


// throws Exception
func (this *AllBinaryGameCanvas) UpdateColor() {
}


// throws Exception
func (this *AllBinaryGameCanvas) BuildGame(portion int) {
        this.ScreenCapture.EndRecording()
        this.ScreenCapture.StartRecording()
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        var features *Features = Features_GetInstance(); _ = features
        progressCanvas.AddNormalPortion(portion, "Generic Build")
        this.RemoveAllGameKeyInputListenersOnBuild()
        this.UpdateTouch()
        progressCanvas.AddNormalPortion(portion, "High Scores")
        this.HighScoresFactoryInterface.FetchHighScores(this.GameLayerManager.GetGameInfo(), this.HighScoresHelper)
        this.SetHighScoresPaintable(NullPaintable_GetInstance())
        progressCanvas.AddNormalPortion(portion, "Finishing.")
        this.GetEndGameInfoP().Update(this.GetLayerManager().GetGameInfo(), this)
        this.IsSingleKeyRepeatableProcessing = features.IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS)
        this.SetGameOver(false)
        progressCanvas.AddNormalPortion(portion, "Finishing..")
        this.UpdateColor()
        this.GetStartIntermissionPaintable().Update()
        if !this.IsRunningInAnotherThread() {
            return
        }
        progressCanvas.AddNormalPortion(portion, "Finishing...")
        if this.IsMainCanvas() || !this.IsInitialized() {
            if features.IsFeature(MainFeatureFactory_GetInstance().LOAD_ONDEMAND) {
                if this.GetCustomCommandListener() != nil {
                    progressCanvas.End()
                }
            } else {
                progressCanvas.End()
            }
            this.SetInitialized(true)
        }
        if this.GetCustomCommandListener() == NullCommandListener_NULL_COMMAND_LISTENER {
            this.LogUtil.PutF("Show Game Paintable in DemoCanvas Thread", this, this.BUILD_GAME)
            this.GameCanvasStartListener.ShowGamePaintable()
        } else {
            this.LogUtil.PutF("No GameCanvasStartListener", this, this.BUILD_GAME)
        }
        this.ColorFillPaintable.SetBasicColorP(this.GameLayerManager.GetBackgroundBasicColor())
        this.GameBehavior.BuildGame(this)
}


func (this *AllBinaryGameCanvas) BuildGame2() {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.ENABLE_PLAYER_GAME_INPUTS).Appendint(this.LocalPlayerGameInputList.Size())), this, this.BUILD_GAME)
        var playerGameInput IPlayerGameInput; _ = playerGameInput
        for index := int(this.LocalPlayerGameInputList.Size() - 1); index >= 0; index-- {
            playerGameInput = any(this.LocalPlayerGameInputList.Get(index)).(IPlayerGameInput)
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.ENABLE_PLAYER_GAME_INPUT).Append(fmt.Sprint(playerGameInput))), this, this.BUILD_GAME)
            playerGameInput.RemoveNonAIInputGameKeyEvents()
            this.AddKeyInputListener(playerGameInput)
        }
}


func (this *AllBinaryGameCanvas) AddKeyInputListener(playerGameInput IPlayerGameInput) {
        GameKeyEventHandler_GetInstance().AddListenerForPlayer(playerGameInput, playerGameInput.GetPlayerInputId())
}


func (this *AllBinaryGameCanvas) SetGameCanvasStartListener(gameCanvasStartListener DemoPaintableInterface) {
        this.GameCanvasStartListener = gameCanvasStartListener
}


// throws Exception
func (this *AllBinaryGameCanvas) LoadState() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.LOAD)
        var hashtable *ABHashtable[any, any] = this.GetLoadStateHashtable(); _ = hashtable
        if hashtable != nil && hashtable.Size() > 0 {
            var levelAsString string = any(hashtable.Get(any(GameInfo_LEVEL_NAME).(any))).(string); _ = levelAsString
            var level int = int(fmt.Sprint(levelAsString)); _ = level
            var gameInfo *GameInfo = this.GameLayerManager.GetGameInfo(); _ = gameInfo
            gameInfo.SetCurrentLevel(level)
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) GetLoadStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonLabels.START_LABEL).Append(this.StringUtil.ToString(this.Hashtable))), this, "getLoadStateHashtable")
        return this.Hashtable
}


func (this *AllBinaryGameCanvas) SetLoadStateHashtable(hashtable *ABHashtable[any, any]) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonLabels.START_LABEL).Append(this.StringUtil.ToString(hashtable))), this, "setLoadStateHashtable")
        this.Hashtable = hashtable
}


func (this *AllBinaryGameCanvas) GetCurrentStateHashtable() *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = this.StdUtil.CreateHashtable(); _ = hashtable
        var level int = this.GameLayerManager.GetGameInfo().GetCurrentLevel(); _ = level
        hashtable.Put(GameInfo_LEVEL_NAME, fmt.Sprint(level))
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("End: ").Append(this.StringUtil.ToString(hashtable))), this, "getCurrentStateHashtable")
        return hashtable
}


func (this *AllBinaryGameCanvas) PaintGameOver(graphics IGraphics) {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}


func (this *AllBinaryGameCanvas) Draw(graphics IGraphics) {
        this.ColorFillPaintable.Paint(graphics)
        this.BasicSetColorUtil.SetBasicColorP(graphics, this.GameLayerManager.GetForegroundBasicColor())
        this.GameSpecificPaintable.Paint(graphics)
}


func (this *AllBinaryGameCanvas) Clear(graphics IGraphics) {
        this.ColorFillPaintable.Paint(graphics)
}


func (this *AllBinaryGameCanvas) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        this.BaseGameStatistics.NextRefresh()
        this.Draw(graphics)
        this.MenuPaintable.Paint(graphics)
        this.ProgressPaintable.Paint(graphics)
}


func (this *AllBinaryGameCanvas) PaintThreed(graphics IGraphics) {
}


// throws Exception
func (this *AllBinaryGameCanvas) ProcessEndLevelIntermissionGameState() {
}


func (this *AllBinaryGameCanvas) NonBotPaint(graphics IGraphics) {
        this.EndGamePaintable.Paint(graphics)
        this.IntermissionPaintable.Paint(graphics)
        this.PaintIntermission(graphics)
        this.TouchPaintable.Paint(graphics)
}


func (this *AllBinaryGameCanvas) PaintIntermission(graphics IGraphics) {
}


func (this *AllBinaryGameCanvas) KeyPressed(keyCode int) {
        this.KeyPressedByDevice(keyCode, 0)
}


func (this *AllBinaryGameCanvas) KeyReleased(keyCode int) {
        this.KeyReleasedByDevice(keyCode, 0)
}


func (this *AllBinaryGameCanvas) KeyRepeated(keyCode int) {
        this.KeyRepeatedByDevice(keyCode, 0)
}


func (this *AllBinaryGameCanvas) KeyPressedByDevice(keyCode int, deviceId int) {
        this.InputProcessor.KeyPressedByDevice(keyCode, deviceId)
}


func (this *AllBinaryGameCanvas) KeyRepeatedByDevice(keyCode int, deviceId int) {
        if this.IsSingleKeyRepeatableProcessing {
            this.InputProcessor.KeyPressedByDevice(keyCode, deviceId)
        }
}


func (this *AllBinaryGameCanvas) KeyReleasedByDevice(keyCode int, deviceId int) {
        this.InputProcessor.KeyReleasedByDevice(this, keyCode, deviceId)
}


// throws Exception
func (this *AllBinaryGameCanvas) HandleRawKey(keyCode int, deviceId int, repeated bool) {
}


func (this *AllBinaryGameCanvas) EndProgress(isProgress bool) int {
        var portion int = 30; _ = portion
        if isProgress && this.IsMainCanvas() {
            var myCommandsFactory *MyCommandsFactory = MyCommandsFactory_GetInstance(); _ = myCommandsFactory
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            progressCanvas.Start()
            this.GetCustomCommandListener().CommandAction(myCommandsFactory.SET_DISPLAYABLE, progressCanvas)
            portion = 4
        }
        return portion
}


// throws Exception
func (this *AllBinaryGameCanvas) ProcessPlayingGame() {
        this.CheatProcessor.Update()
        this.SensorGameUpdateProcessor.GetInputSensor().Update()
        this.GameInputProcessor.Process()
        this.GameLayerManager.Process()
        this.StartIntermissionProcessor.Process()
}


// throws Exception
func (this *AllBinaryGameCanvas) ThreadInit() {
}


// throws Exception
func (this *AllBinaryGameCanvas) ProcessGame() {
        this.RunnableCanvas.Process()
        if AllBinaryMediaManager_Update() {
            if !this.PrimaryPlayerQueue.Process() {
                this.SecondaryPlayerQueue.Process()
            }
        }
        this.MainStateProcessor.Process()
        this.MenuInputProcessor.ProcessInputList()
        this.EndGameProcessor.Process()
        this.BaseGameStatistics.NextFrame()
}


func (this *AllBinaryGameCanvas) NotifyDonePainting() {
        { // synchronized (this)
            this.ThreadObjectUtil.NotifyObject(this)
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) ShouldWait() {
        var features *Features = Features_GetInstance(); _ = features
        if AndroidUtil_IsAndroid() {
        } else if features.IsFeature(this.OpenGLFeatureFactory.OPENGL_2D) {
        } else {
            if features.IsDefault(this.OpenGLFeatureFactory.OPENGL) {
                var openGLImageSpecificFactory *OpenGLImageSpecificFactory = OpenGLImageSpecificFactory_GetInstance(); _ = openGLImageSpecificFactory
                for openGLImageSpecificFactory.Updating {
                    time.Sleep(time.Duration(this.YIELD_SLEEP) * time.Millisecond)
                }
            }
        }
}


func (this *AllBinaryGameCanvas) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START_RUNNABLE, this, this.CommonStrings.RUN)
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            progressCanvas.AddNormalPortion(50, "Game Thread")
            this.SetCurrentThread()
            this.SetRunning(true)
            if this.GetCustomCommandListener() == NullCommandListener_NULL_COMMAND_LISTENER {
                time.Sleep(time.Duration(this.YIELD_SLEEP) * time.Millisecond)
            }
            this.ShouldWait()
            this.ThreadInit()
            if this.GetCustomCommandListener() == NullCommandListener_NULL_COMMAND_LISTENER {
                time.Sleep(time.Duration(this.YIELD_SLEEP) * time.Millisecond)
            }
            this.Open()
            var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
            gameAdState.Init()
            gameAdState.SetGameIsReady(true)
            this.GameBehavior.Run(this)
            this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) Run2() {
        var features *Features = Features_GetInstance(); _ = features
        if features.IsDefault(this.OpenGLFeatureFactory.OPENGL_AS_GAME_THREAD) || J2MEUtil_IsHTML() {
            if features.IsDefault(this.OpenGLFeatureFactory.OPENGL_AS_GAME_THREAD) {
                this.LogUtil.PutF(this.OpenGLFeatureFactory.OPENGL_AS_GAME_THREAD.GetName(), this, this.CommonStrings.RUN)
            }
            if J2MEUtil_IsHTML() {
                this.LogUtil.PutF(GraphicsStrings_GetInstance().HTML, this, this.CommonStrings.RUN)
            }
            var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
            currentDisplayableFactory.SetRunnable(this.GameRunnable)
            currentDisplayableFactory.SetDisplayable(this)
            OpenGLThreadUtil_GetInstance().OnResume()
        } else if features.IsDefault(this.OpenGLFeatureFactory.OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS) {
            this.LogUtil.PutF(this.OpenGLFeatureFactory.OPENGL_AND_GAME_HAVE_DIFFERENT_THREADS.GetName(), this, this.CommonStrings.RUN)
            OpenGLThreadUtil_GetInstance().OnResume()
            for this.IsRunning() {
                this.Run3()
            }
            this.End()
        } else {
            this.LogUtil.PutF("this thread", this, this.CommonStrings.RUN)
            for this.IsRunning() {
                this.Run3()
            }
            this.End()
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) Run3() {
        this.LoopTimeHelper.SetStartTime(this.GameTickTimeDelayHelper.SetStartTime())
        this.GameTickDisplayInfoSingleton.Update()
        this.ProcessGame()
        this.ProcessLoopSleep()
}


func (this *AllBinaryGameCanvas) SetRunning(running bool) {
        this.RunnableCanvas.SetRunning(running)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.SET_RUNNING, e)
                }
            }()
            var features *Features = Features_GetInstance(); _ = features
            if running {
            } else {
                if features.IsDefault(this.OpenGLFeatureFactory.OPENGL) || SWTUtil_IsSWT {
                    if this.GameLayerManager.GetGameInfo().GetGameType() != this.GameTypeFactory.BOT {
                        var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
                        currentDisplayableFactory.ClearRunnable()
                    }
                    this.End()
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AllBinaryGameCanvas) End() {
        this.ScreenCapture.EndRecording()
        this.BaseGameStatistics.Add(fmt.Sprint(NewStringMaker().Append(fmt.Sprint(this.BaseGameStatistics)).Append(CommonSeps_GetInstance().NEW_LINE)))
        this.BaseGameStatistics.Init()
        this.GameKeyEventHandler.RemoveListener(this.CheatProcessor)
        this.Close()
        this.RemoveAllGameKeyInputListeners()
        this.EndGameThread()
}


// throws Exception
func (this *AllBinaryGameCanvas) EndGameThread() {
        DisplayChangeEventHandler_GetInstance().RemoveListener(this)
        GameLevelDisplayChangeEventListenersFactory_GetInstance().Clear()
        var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
        gameAdState.NotPlayingAdState()
}


func (this *AllBinaryGameCanvas) CreateHighScore(score int64) *HighScore {
        var gameInfo *GameInfo = this.GameLayerManager.GetGameInfo(); _ = gameInfo
        return NewHighScore(0, "NONE", NewGameInfo(gameInfo.GetGameType(), gameInfo.GetGameMode(), PlayerTypesFactory_GetInstance().PLAYER_TYPE_ONE, 0, 0), score)
}


// throws Exception
func (this *AllBinaryGameCanvas) SetHighScore(abeClientInformation AbeClientInformationInterface, name string, score int64, autoSubmit bool, isLast bool) {
        this.GameBehavior.SetHighScore(abeClientInformation, this, name, score, autoSubmit, isLast)
}


// throws Exception
func (this *AllBinaryGameCanvas) SetHighScore2(abeClientInformation AbeClientInformationInterface, name string, score int64, autoSubmit bool, isLast bool) {
        var highScore *HighScore = this.CreateHighScore(score); _ = highScore
        var textBox *HighScoreTextBox = NewHighScoreTextBox(this.HighScoresFactoryInterface, this.HighScoresHelper, abeClientInformation, this.GameLayerManager.GetGameInfo(), this.GetCustomCommandListener(), name, highScore, this.GameLayerManager.GetBackgroundBasicColor(), this.GameLayerManager.GetForegroundBasicColor()); _ = textBox
        if isLast {
            this.GetCustomCommandListener().CommandAction(GameCommandsFactory_GetInstance().SET_MENU_DISPLAYABLE, textBox)
        }
        if autoSubmit {
            /* Java local class: move it out of the function
type SaveHighScoreRunnable struct {
    ProgressCanvas IProgressCanvas
}

// Implements Runnable


func (this *SaveHighScoreRunnable) initFields() {
                this.ProgressCanvas = ProgressCanvasFactory_GetInstance()
}

func NewSaveHighScoreRunnable() *SaveHighScoreRunnable {
    this := &SaveHighScoreRunnable{}
    this.initFields()
    return this
}
            

            func (this *AllBinaryGameCanvas) Run() {
                    var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                    if tryControl1 := func() (tryControl1 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                        logUtil.Put(commonStrings.EXCEPTION, this, "run", e)
                        this.ProgressCanvas.End()
                            }
                        }()
                        if this.IsLast {
                            this.TextBox.Submit()
                        } else {
                            this.TextBox.SaveHighScore()
                        }
                        return
                    }(); tryControl1 == 1 {
                        return
                    }
            }
*/
            SecondaryThreadPool_GetInstance().RunTask(NewSaveHighScoreRunnable())
        }
}


func (this *AllBinaryGameCanvas) GetGameStateTimeHelper() ITimeDelayHelper {
        return this.GameStateTimeHelper
}


func (this *AllBinaryGameCanvas) SetInitialized(initialized bool) {
        this.Initialized = initialized
}


func (this *AllBinaryGameCanvas) IsInitialized() bool {
        return this.Initialized
}


func (this *AllBinaryGameCanvas) GetSourceId() int {
        return AllBinaryGameCanvas_Id
}


func (this *AllBinaryGameCanvas) SetProgressPaintable(paintable PaintableInterface) {
        this.ProgressPaintable = paintable
}


func (this *AllBinaryGameCanvas) GetStartIntermissionInterface() IntermissionInterface {
        return this.StartIntermissionInterface
}


func (this *AllBinaryGameCanvas) GetEndLevelIntermissionInterface() IntermissionInterface {
        return this.EndLevelIntermissionInterface
}


func (this *AllBinaryGameCanvas) SetTouchPaintableP(paintable IPaintable) {
        this.TouchPaintable = paintable
}


func (this *AllBinaryGameCanvas) GetTouchPaintableP() IPaintable {
        return this.TouchPaintable
}


func (this *AllBinaryGameCanvas) SetEndGamePaintable(endGamePaintable IPaintable) {
        this.EndGamePaintable = endGamePaintable
}


func (this *AllBinaryGameCanvas) GetEndGamePaintable() IPaintable {
        return this.EndGamePaintable
}


func (this *AllBinaryGameCanvas) SetIntermissionPaintable(intermissionPaintable IPaintable) {
        this.IntermissionPaintable = intermissionPaintable
}


func (this *AllBinaryGameCanvas) GetIntermissionPaintable() IPaintable {
        return this.IntermissionPaintable
}


func (this *AllBinaryGameCanvas) GetEndGameInfoP() IEndGameInfo {
        return this.EndGameInfo
}


func (this *AllBinaryGameCanvas) SetHighScoresPaintable(highScoresPaintable IPaintable) {
        this.HighScoresPaintable = highScoresPaintable
}


func (this *AllBinaryGameCanvas) GetHighScoresPaintable() IPaintable {
        return this.HighScoresPaintable
}


func (this *AllBinaryGameCanvas) GetRealHighScoresPaintable() *HighScoresPaintable {
        return this.RealHighScoresPaintable
}


func (this *AllBinaryGameCanvas) ClearPlayerGameInputList() {
        var playerGameInput IPlayerGameInput; _ = playerGameInput
        for index := int(this.LocalPlayerGameInputList.Size() - 1); index >= 0; index-- {
            playerGameInput = any(this.LocalPlayerGameInputList.Get(index)).(IPlayerGameInput)
            this.GameKeyEventHandler.RemoveListener(playerGameInput)
        }
        this.LocalPlayerGameInputList.Clear()
}


func (this *AllBinaryGameCanvas) AddPlayerGameInput(playerGameInput IPlayerGameInput) {
        this.LocalPlayerGameInputList.Add(playerGameInput)
}


func (this *AllBinaryGameCanvas) SetMenuInputProcessor(menuInputProcessor IBasicMenuInputProcessor) {
        this.MenuInputProcessor = menuInputProcessor
}


func (this *AllBinaryGameCanvas) GetMenuInputProcessor() IBasicMenuInputProcessor {
        return this.MenuInputProcessor
}


func (this *AllBinaryGameCanvas) GetMenuForm() IScrollSelectionForm {
        return this.MenuForm
}


func (this *AllBinaryGameCanvas) SetMenuForm(menuForm IScrollSelectionForm) {
        this.MenuForm = menuForm
}


func (this *AllBinaryGameCanvas) GetStartLevel() int {
        return this.StartLevel
}


func (this *AllBinaryGameCanvas) SetStartLevel(startLevel int) {
        this.StartLevel = startLevel
}


func (this *AllBinaryGameCanvas) SetTouchButtonsPaintable(touchButtonsPaintable IPaintable) {
        this.TouchButtonsPaintable = touchButtonsPaintable
}


func (this *AllBinaryGameCanvas) GetTouchButtonsPaintable() IPaintable {
        return this.TouchButtonsPaintable
}


func (this *AllBinaryGameCanvas) SetGameInputProcessor(gameInputProcessor IProcessor) {
        this.GameInputProcessor = gameInputProcessor
}


func (this *AllBinaryGameCanvas) GetGameInputProcessor() IProcessor {
        return this.GameInputProcessor
}


func (this *AllBinaryGameCanvas) SetEndGameProcessor(endGameProcessor IProcessor) {
        this.EndGameProcessor = endGameProcessor
}


func (this *AllBinaryGameCanvas) GetEndGameProcessor() IProcessor {
        return this.EndGameProcessor
}


func (this *AllBinaryGameCanvas) SetEndGameStatePaintable(endGameStatePaintable IPaintable) {
        this.EndGameStatePaintable = endGameStatePaintable
}


func (this *AllBinaryGameCanvas) GetEndGameStatePaintable() IPaintable {
        return this.EndGameStatePaintable
}


func (this *AllBinaryGameCanvas) SetNonBotPaintableP(nonBotPaintable IPaintable) {
        this.NonBotPaintable = nonBotPaintable
}


func (this *AllBinaryGameCanvas) GetNonBotPaintableP() IPaintable {
        return this.NonBotPaintable
}


func (this *AllBinaryGameCanvas) SetStartIntermissionPaintable(startIntermissionPaintable IInitUpdatePaintable) {
        this.StartIntermissionPaintable = startIntermissionPaintable
}


func (this *AllBinaryGameCanvas) GetStartIntermissionPaintable() IInitUpdatePaintable {
        return this.StartIntermissionPaintable
}


func (this *AllBinaryGameCanvas) SetMainStateProcessor(mainStateProcessor IProcessor) {
        this.MainStateProcessor = mainStateProcessor
}


func (this *AllBinaryGameCanvas) GetMainStateProcessor() IProcessor {
        return this.MainStateProcessor
}


func (this *AllBinaryGameCanvas) SetProcessGameProcessor(processGameProcessor IProcessor) {
        this.ProcessGameProcessor = processGameProcessor
}


func (this *AllBinaryGameCanvas) GetProcessGameProcessor() IProcessor {
        return this.ProcessGameProcessor
}


func (this *AllBinaryGameCanvas) SetOpenMenuPaintable(openMenuPaintable IPaintable) {
        this.OpenMenuPaintable = openMenuPaintable
}


func (this *AllBinaryGameCanvas) GetOpenMenuPaintable() IPaintable {
        return this.OpenMenuPaintable
}


func (this *AllBinaryGameCanvas) SetPopupMenuInputProcessor(popupMenuInputProcessor IBasicMenuInputProcessor) {
        this.PopupMenuInputProcessor = popupMenuInputProcessor
}


func (this *AllBinaryGameCanvas) GetPopupMenuInputProcessor() IBasicMenuInputProcessor {
        return this.PopupMenuInputProcessor
}


func (this *AllBinaryGameCanvas) GetSensorGameUpdateProcessor() ISensorGameUpdateProcessor {
        return this.SensorGameUpdateProcessor
}


func (this *AllBinaryGameCanvas) GetRawGameInputProcessor() IInputProcessor {
        return this.RawGameInputProcessor
}


func (this *AllBinaryGameCanvas) GetRawInputProcessor() IInputProcessor {
        return this.RawInputProcessor
}


func (this *AllBinaryGameCanvas) SetInputProcessor(inputProcessor IInputProcessor) {
        this.InputProcessor = inputProcessor
}


func (this *AllBinaryGameCanvas) GetInputProcessor() IInputProcessor {
        return this.InputProcessor
}


func (this *AllBinaryGameCanvas) SetMenuPaintable(menuPaintable IPaintable) {
        this.MenuPaintable = menuPaintable
}


func (this *AllBinaryGameCanvas) GetMenuPaintable() IPaintable {
        return this.MenuPaintable
}


func (this *AllBinaryGameCanvas) SetFormPaintable(formPaintable IPaintable) {
        this.FormPaintable = formPaintable
}


func (this *AllBinaryGameCanvas) GetFormPaintable() IPaintable {
        return this.FormPaintable
}


func (this *AllBinaryGameCanvas) SetGameSpecificPaintableP(gameSpecificPaintable IPaintable) {
        this.GameSpecificPaintable = gameSpecificPaintable
}


func (this *AllBinaryGameCanvas) GetGameSpecificPaintableP() IPaintable {
        return this.GameSpecificPaintable
}


func (this *AllBinaryGameCanvas) IsSingleThread() bool {
        return OpenGLFeatureUtil_GetInstance().IsAnyThreed() || SWTUtil_IsSWT
}


func (this *AllBinaryGameCanvas) IsRunningInAnotherThread() bool {
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        if features.IsDefault(openGLFeatureFactory.OPENGL_AS_GAME_THREAD) {
            return true
        } else {
            return this.IsRunning()
        }
}


func (this *AllBinaryGameCanvas) GetType() int {
        return AllBinaryGameCanvas_TYPE
}

