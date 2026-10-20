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

type DemoCanvas struct {
    RunnableCanvas
    BasicColorFactory *BasicColorFactory
    MyCommandsFactory *MyCommandsFactory
    GameAdStateFactory *GameAdStateFactory
    GameInputStrings *GameInputStrings
    GameStateFactory *GameStateFactory
    ThreadFactoryUtil *ThreadFactoryUtil
    FormUtil *FormUtil
    FullscreenPaintable IPaintable
    BasicGameDemoPaintable IStatePaintable
    SpecialAnimationInterface ISpecialAnimation
    PaintedSpecialAnimationInterface IAnimation
    FullScreenUtil *FullScreenUtil
    GameCanvas IAllBinaryGameCanvas
    HighScoresHelper IHighScoresHelperBase
    RealHighScoresPaintable *HighScoresPaintable
    HighScoresPaintable IPaintable
    State int
    CanvasThread IThread
    TimeDelayHelper ITimeDelayHelper
    BasicColor *BasicColor
    AbeClientInformation AbeClientInformationInterface
    HighScoresFactoryInterface HighScoresFactoryInterface
    MenuInputProcessor IBasicMenuInputProcessor
    MenuForm IScrollSelectionForm
    GameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory
    Initialized bool
    DemoGameRunnable *DemoGameStartupRunnable
    DefaultPaintableInterface IPaintable
    PaintableInterface PaintableInterface
    OverlayPaintable IInitUpdatePaintable
    ProgressPaintable PaintableInterface
    TempWait int
    InputToGameKeyMapping *InputToGameKeyMapping
    GameRunnable IGameRunnable
    IsSingleKeyRepeatableProcessing bool
    NONE *GameKey
    GameKeyEventFactory *GameKeyEventFactory
    LastKeyNotMapped int
    SET_STATE string
    BaseGameStatistics IBaseGameStatistics
}

// Implements GameCanvasRunnableInterface, MenuListener, DisplayChangeEventListener, DemoPaintableInterface

var DemoCanvas_Id int = 0
var DemoCanvas_BOT_GAME_STATS string = "Bot Game Statistics: "
var DemoCanvas_TYPE int = 3

func (this *DemoCanvas) initFields() {
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.MyCommandsFactory = MyCommandsFactory_GetInstance()
    this.GameAdStateFactory = GameAdStateFactory_GetInstance()
    this.GameInputStrings = GameInputStrings_GetInstance()
    this.GameStateFactory = GameStateFactory_GetInstance()
    this.ThreadFactoryUtil = ThreadFactoryUtil_GetInstance()
    this.FormUtil = FormUtil_GetInstance()
    this.FullscreenPaintable = FullScreenPaintableFactory_GetInstance().Paintable
    this.BasicGameDemoPaintable = StatePaintableFactory_GetInstance()
    this.SpecialAnimationInterface = SpecialAnimation_GetInstance()
    this.PaintedSpecialAnimationInterface = SpecialAnimation_GetInstance()
    this.FullScreenUtil = FullScreenUtil_GetInstance()
    this.GameCanvas = NullGameCanvas_GetInstance()
    this.HighScoresHelper = NewHighScoresHelper()
    this.RealHighScoresPaintable = NewHighScoresPaintable()
    this.HighScoresPaintable = NullPaintable_GetInstance()
    this.State = 0
    this.CanvasThread = NullThread_NULL_THREAD
    this.TimeDelayHelper = NewTimeDelayHelper(6000)
    this.BasicColor = this.BasicColorFactory.RED
    
    
    this.MenuInputProcessor = NoMenuInputProcessor_GetInstance()
    this.MenuForm = ScrollSelectionForm_GetNullScrollSelectionForm()
    
    
    
    this.DefaultPaintableInterface = NullPaintable_GetInstance()
    this.PaintableInterface = NullPaintable_GetInstance()
    
    this.ProgressPaintable = ProgressCanvasFactory_GetLazyInstance()
    this.TempWait = NullWaitGameRunnable_GetInstance().GameRunnable_().WAIT
    this.InputToGameKeyMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().GetInputMapping()
    this.GameRunnable = NullWaitGameRunnable_GetInstance()
    this.IsSingleKeyRepeatableProcessing = Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS)
    this.NONE = GameKeyFactory_GetInstance().NONE
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    this.LastKeyNotMapped = -1
    this.SET_STATE = "setState"
    this.BaseGameStatistics = GameStatisticsFactory_GetInstance()
}


func NewDemoCanvas(abeClientInformation AbeClientInformationInterface, commandListener CommandListener, highScoresFactoryInterface HighScoresFactoryInterface, paintable IPaintable, overlayPaintable IInitUpdatePaintable, gameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory, isContinue bool) *DemoCanvas {
    this := &DemoCanvas{}
        this.RunnableCanvas = *NewRunnableCanvas(commandListener, CurrentDisplayableFactory_GetInstance().DEFAULT_CHILD_NAME_LIST, true)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.SetWait(NullWaitGameRunnable_GetInstance().GameRunnable_().WAIT)
        this.GameInitializationInterfaceFactoryInterface = gameInitializationInterfaceFactoryInterface
        GameInitializationUtil_GetInstance().InitDemo(abeClientInformation, this, gameInitializationInterfaceFactoryInterface)
        ResizableListenerHandler_GetInstance().FireEvent(false)
        this.OverlayPaintable = overlayPaintable
        this.HighScoresFactoryInterface = highScoresFactoryInterface
        this.SetDefaultPaintableInterface(paintable)
        this.SetPaintableInterface(this.GetDefaultPaintableInterface())
        if isContinue {
            this.AddCommand(GameCommandsFactory_GetInstance().CONTINUE_COMMAND)
        }
        if ChangedGameFeatureListener_GetInstance().IsChangedFeature(GameFeatureFactory_GetInstance().SOUND) {
            this.MediaInit()
        }
        this.DemoGameRunnable = NewDemoGameStartupRunnable(this)
        DisplayChangeEventHandler_GetInstance().AddListenerInterface(this)
    return this
}


func (this *DemoCanvas) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *DemoCanvas) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT, e)
                }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonLabels.START_LABEL).Append(fmt.Sprint(this.DisplayInfo))), this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT)
            var scrollSelectionForm IScrollSelectionForm = this.GetMenuForm(); _ = scrollSelectionForm
            if scrollSelectionForm != nil {
                var formType *FormType = FormTypeFactory_GetInstance().GetFormType(); _ = formType
                var rectangle *Rectangle = this.FormUtil.CreateFormRectangle(); _ = rectangle
                scrollSelectionForm.Init(rectangle, formType)
            }
            this.OverlayPaintable.Init()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DemoCanvas) GetCustomCommands() []any {
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        if J2MEUtil_IsHTML() {
            var commandArray []any = []any{gameCommandsFactory.START_COMMAND, HighScoreCommands_GetInstance().DISPLAY, GameInputMappingCanvas_DISPLAY, gameCommandsFactory.DISPLAY_ABOUT}; _ = commandArray
            return commandArray
        } else {
            var commandList IBasicArrayList = NewBasicArrayListD(); _ = commandList
            commandList.Add(gameCommandsFactory.START_COMMAND)
            var inApplicationPurchaseFactory *InApplicationPurchaseFactory = InApplicationPurchaseFactory_GetInstance(); _ = inApplicationPurchaseFactory
            if inApplicationPurchaseFactory.IsEnabled() {
                var list IBasicArrayList = LockableFeatureFactory_GetInstance().GetList(); _ = list
                if list.Size() > 0 && !inApplicationPurchaseFactory.IsPurchased(any(list.Get(0)).(*LockableFeature)) {
                    commandList.Add(gameCommandsFactory.BUY_COMMAND)
                }
            }
            commandList.Add(HighScoreCommands_GetInstance().DISPLAY)
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []any) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                    }
                }()
                var isOverScan bool = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance().IsOverScan(); _ = isOverScan
                if SWTUtil_IsSWT {
                    commandList.Add(GameInputMappingCanvas_DISPLAY)
                } else if !isOverScan {
                    commandList.Add(gameCommandsFactory.DISPLAY_OPTIONS)
                    commandList.Add(gameCommandsFactory.DISPLAY_LOAD_FORM)
                    commandList.Add(GameInputMappingCanvas_DISPLAY)
                }
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            commandList.Add(gameCommandsFactory.DISPLAY_ABOUT)
            var commandArray []any = commandList.ToArray(); _ = commandArray
            return commandArray
        }
}


func (this *DemoCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        var commandArray []any = this.GetCustomCommands(); _ = commandArray
        var size int = len(commandArray); _ = size
        for index := int(0); index < size; index++ {
            this.AddCommand(any(commandArray[index]).(ICommand))
        }
        CustomGameMenuUtil_Add(this)
        this.SetCommandListener(cmdListener)
}


// throws Exception
func (this *DemoCanvas) InitPostPaint() {
}


// throws Exception
func (this *DemoCanvas) MediaInit() {
        AllBinaryMediaManager_Init(EarlySoundsFactory_GetInstance())
}


func (this *DemoCanvas) ItemStateChanged(item IItem) {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}


// throws Exception
func (this *DemoCanvas) InitMenu() {
        this.Close()
        var commandTextItemArrayFactory *CommandTextItemArrayFactory = DemoLimitedCommandTextItemArrayFactory_GetInstance().GetCommandTextItemArrayFactory(); _ = commandTextItemArrayFactory
        var items []IABCustomItem = commandTextItemArrayFactory.GetInstance(this.GetCommandStack(), this.BasicColorFactory.BLACK, this.BasicColorFactory.WHITE); _ = items
        var formType *FormType = FormTypeFactory_GetInstance().GetFormType(); _ = formType
        var rectangle *Rectangle = this.FormUtil.CreateFormRectangle(); _ = rectangle
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.CommonLabels.START_LABEL).Append(fmt.Sprint(this.DisplayInfo))), this, "initMenu")
        var scrollSelectionForm IScrollSelectionForm = CommandCurrentSelectionFormFactory_GetInstance(StringUtil_GetInstance().EMPTY_STRING, items, rectangle, formType, 15, true, this.BasicColorFactory.BLACK, this.BasicColorFactory.WHITE); _ = scrollSelectionForm
        this.SetMenuForm(scrollSelectionForm)
        var formType2 *FormType = FormTypeFactory_GetInstance().GetFormType(); _ = formType2
        var rectangle2 *Rectangle = this.FormUtil.CreateFormRectangle(); _ = rectangle2
        scrollSelectionForm.Init(rectangle2, formType2)
        if this.GetMenuForm() != ScrollSelectionFormNoneFactory_GetInstance() {
            this.SetMenuInputProcessor(NewCommandFormInputProcessor(NewBasicArrayListD(), -1, this, this.GetMenuForm()))
        }
        this.Open()
}


func (this *DemoCanvas) Open() {
        BasicMotionGesturesHandler_GetInstance().AddListenerInterface(this.GetMenuInputProcessor())
        GameKeyEventHandler_GetInstance().AddListener(this.GetMenuInputProcessor())
}


func (this *DemoCanvas) Close() {
        BasicMotionGesturesHandler_GetInstance().RemoveListener(this.GetMenuInputProcessor())
        GameKeyEventHandler_GetInstance().RemoveListener(this.GetMenuInputProcessor())
}


func (this *DemoCanvas) GetSourceId() int {
        return DemoCanvas_Id
}


func (this *DemoCanvas) KeyPressed(keyCode int) {
        this.KeyPressedByDevice(keyCode, 0)
}


func (this *DemoCanvas) KeyReleased(keyCode int) {
        this.KeyReleasedByDevice(keyCode, 0)
}


func (this *DemoCanvas) KeyRepeated(keyCode int) {
        this.KeyRepeatedByDevice(keyCode, 0)
}


func (this *DemoCanvas) KeyPressedByDevice(keyCode int, deviceId int) {
        this.AddGameKeyEvent(keyCode, false)
}


func (this *DemoCanvas) KeyReleasedByDevice(keyCode int, deviceId int) {
        this.RemoveGameKeyEvent(keyCode, false)
}


func (this *DemoCanvas) KeyRepeatedByDevice(keyCode int, deviceId int) {
        if this.IsSingleKeyRepeatableProcessing {
            this.AddGameKeyEvent(keyCode, true)
        }
}


func (this *DemoCanvas) AddGameKeyEvent(keyCode int, repeated bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.ADD_KEY_EVENT, e)
                }
            }()
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(this, keyCode); _ = gameKey
            if gameKey != this.NONE {
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this, gameKey); _ = gameKeyEvent
                DownGameKeyEventHandler_GetInstance().FireEvent(gameKeyEvent)
            } else {
                if this.LastKeyNotMapped != keyCode {
                    this.LastKeyNotMapped = keyCode
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.ADD_KEY_EVENT)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DemoCanvas) RemoveGameKeyEvent(keyCode int, repeated bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.REMOVE_KEY_EVENT, e)
                }
            }()
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(this, keyCode); _ = gameKey
            if gameKey != this.NONE {
                var gameKeyEvent *GameKeyEvent = this.GameKeyEventFactory.GetInstanceForInput(this, gameKey); _ = gameKeyEvent
                UpGameKeyEventHandler_GetInstance().FireEvent(gameKeyEvent)
            } else {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.REMOVE_KEY_EVENT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *DemoCanvas) Pause() {
        this.Close()
        this.SetPaused(true)
        this.GameRunnable = NullWaitGameRunnable_GetInstance()
        this.GameCanvas.Pause()
}


// synchronized
func (this *DemoCanvas) UnPause() {
        this.Open()
        this.GameCanvas.UnPause()
        this.GameRunnable = this.GameCanvas.AllBinaryGameCanvas_().GameRunnable
        this.SetPaused(false)
}


func (this *DemoCanvas) IsPausable() bool {
        if CurrentDisplayableFactory_GetInstance().GetUsedRunnable() == NullWaitGameRunnable_GetInstance() {
            return true
        } else {
            return false
        }
}


func (this *DemoCanvas) IsGameOver() bool {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.NOT_IMPLEMENTED).Append(" since not a game")), this, "isGameOver")
        return false
}


// throws Exception
func (this *DemoCanvas) SetLoadStateHashtable(hashtable *ABHashtable[any, any]) {
        this.LogUtil.PutF("Trying to continue a demo lol - only continue a game canvas not the demo", this, "setLoadStateHashtable")
}


// throws Exception
func (this *DemoCanvas) GetLoadStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF("Trying to continue a demo lol - only continue a game canvas not the demo", this, "getLoadStateHashtable")
        return this.StdUtil.NULL_TABLE
}


// throws Exception
func (this *DemoCanvas) GetCurrentStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF("Trying to save the AI lol", this, "getCurrentStateHashtable")
        return this.StdUtil.NULL_TABLE
}


func (this *DemoCanvas) SetHighScoreSubmitted(isNotUsed bool) {
}


func (this *DemoCanvas) Paint(graphics IGraphics) {
        this.PaintableInterface.Paint(graphics)
        this.PaintedSpecialAnimationInterface.PaintXY(graphics, 0, 0)
        this.HighScoresPaintable.Paint(graphics)
        this.GetBasicGameDemoPaintable().Paint(graphics)
        this.OverlayPaintable.Paint(graphics)
        this.FullscreenPaintable.Paint(graphics)
        this.ProgressPaintable.Paint(graphics)
}


func (this *DemoCanvas) PaintThreed(graphics IGraphics) {
        this.PaintableInterface.PaintThreed(graphics)
        this.PaintedSpecialAnimationInterface.PaintThreedXYZ(graphics, 0, 0, 0)
}


// synchronized
func (this *DemoCanvas) ProcessGameOver() {
        this.LogUtil.PutF("Not Implemented since not a game", this, "setGameOver")
}


func (this *DemoCanvas) DemoStateChange() {
        var newState int = this.State + 1; _ = newState
        if newState > 2 {
            newState = 0
        } else if newState == 2 {
            if !this.HighScoresHelper.IsAnyHighScores() {
                newState = 0
            } else {
                var highScores IHighScores = this.HighScoresHelper.GetNextHighScores(); _ = highScores
                if highScores == NullHighScoresSingletonFactory_GetInstance() {
                    newState = 0
                } else {
                    this.GetRealHighScoresPaintable().SetHighScores(highScores)
                }
            }
        }
        this.SetState(newState)
        this.UpdateDemoState()
}


func (this *DemoCanvas) UpdateDemoState() {
        PreLogUtil_Put(fmt.Sprint(SmallIntegerSingletonFactory_GetInstance().CreateInstance(this.State)), this, this.SET_STATE)
        this.GetBasicGameDemoPaintable().SetState(this.State)
        if this.State == 0 {
            this.HighScoresPaintable = NullPaintable_GetInstance()
            this.PaintedSpecialAnimationInterface = this.GetSpecialAnimationInterface()
            if !this.DemoGameRunnable.IsRunning() && this.GameCanvas.IsInitialized() {
                this.GetSpecialAnimationInterface().Reset()
            }
        } else if this.State == 1 {
            this.PaintedSpecialAnimationInterface = SpecialAnimation_GetInstance()
        } else if this.State == 2 {
            this.HighScoresPaintable = this.GetRealHighScoresPaintable()
        }
        var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
        gameAdState.ProcessPageAdState()
}


// throws Exception
func (this *DemoCanvas) GetNextRandom() int {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *DemoCanvas) CreateGameLayerManager(randomValue int) IAllBinaryGameLayerManager {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *DemoCanvas) CreateRunnable(randomLevel int) GameCanvasRunnableInterface {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *DemoCanvas) Create() {
        PreLogUtil_Put(this.CommonStrings.START, this, "create")
        this.HighScoresPaintable = NullPaintable_GetInstance()
        var randomLevel int = this.GetNextRandom(); _ = randomLevel
        this.GameCanvas = any(this.CreateRunnable(randomLevel)).(IAllBinaryGameCanvas)
        this.BasicColor = this.GameCanvas.GetLayerManager().GetForegroundBasicColor()
        this.GetRealHighScoresPaintable().SetBasicColorP(this.BasicColor)
        this.GameCanvas.SetGameCanvasStartListener(this)
        var gameInfo *GameInfo = this.GameCanvas.GetLayerManager().GetGameInfo(); _ = gameInfo
        this.GetHighScoresFactoryInterface().FetchHighScores(gameInfo, this.HighScoresHelper)
}


// throws Exception
func (this *DemoCanvas) Start() {
        var gameCanvas IAllBinaryGameCanvas = this.GameCanvas; _ = gameCanvas
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Game Thread in DemoCanvas: ").Append(this.StringUtil.ToString(gameCanvas))), this, this.CommonStrings.START)
        this.CanvasThread = this.ThreadFactoryUtil.GetInstanceGameCanvasRunnable(gameCanvas)
        this.GameCanvas.SetThread(this.CanvasThread)
        this.ThreadFactoryUtil.Start(this.CanvasThread)
        if this.GetWait() == NullWaitGameRunnable_GetInstance().GameRunnable_().WAIT {
            this.SetWait(this.GetTempWait())
        }
}


func (this *DemoCanvas) PreDemoProcess() {
        if !this.GameCanvas.IsInitialized() || this.GameCanvas.GetTitle() == NullGameCanvas_NO_GAME {
            if AllBinaryMediaManager_Update() {
                if !PrimaryPlayerQueueFactory_GetInstance().Process() {
                    SecondaryPlayerQueueFactory_GetInstance().Process()
                }
            }
        }
        this.OverlayPaintable.Update()
}


// throws Exception
func (this *DemoCanvas) Process() {
        this.RunnableCanvas.Process()
        this.GetMenuInputProcessor().ProcessInputList()
        this.PreDemoProcess()
        if this.State == 0 {
            var indexedAnimationBehavior *IndexedAnimationBehavior = (any(this.GetSpecialAnimationInterface().GetAnimationBehavior()).(*IndexedAnimationBehavior)); _ = indexedAnimationBehavior
            if indexedAnimationBehavior.LoopIndex < 1 {
                this.TimeDelayHelper.SetStartTimeTNT()
            }
            var demoGameMidlet IDemoGameMidlet = any(this.GetCustomCommandListener()).(IDemoGameMidlet); _ = demoGameMidlet
            if this.GameCanvas != NullGameCanvas_GetInstance() && this.GameCanvas.IsGameOver() {
                this.StopGameDemo()
                var randomLevel int = this.GetNextRandom(); _ = randomLevel
                var gameInfo *GameInfo = this.GameCanvas.GetLayerManager().GetGameInfo(); _ = gameInfo
                gameInfo.SetCurrentLevel(randomLevel)
                this.GameCanvas.SetGameOver(false)
                this.Start()
            } else if this.GameCanvas == NullGameCanvas_GetInstance() && demoGameMidlet.IsReady() {
                if !this.DemoGameRunnable.IsRunning() {
                    this.StartDemoGame()
                    this.DemoGameRunnable.SetRunning(true)
                    var thread IThread = this.ThreadFactoryUtil.GetInstanceForRunnable(this.DemoGameRunnable); _ = thread
                    this.DemoGameRunnable.SetThread(thread)
                    this.ThreadFactoryUtil.Start(thread)
                }
            }
        }
}


// throws Exception
func (this *DemoCanvas) StartDemoGame() {
        DemoCanvasProgressUtil_ShowProgress(this)
}


// throws Exception
func (this *DemoCanvas) StopGameDemo() {
        if this.GameCanvas != NullGameCanvas_GetInstance() {
            this.LogUtil.PutF("Set Running False", this, "stopGameDemo")
            this.GameCanvas.SetRunning(false)
        }
        ThreadUtil_GetInstance().Join(this.CanvasThread)
}


func (this *DemoCanvas) ShowGamePaintable() {
        var METHOD_NAME string = "showGamePaintable"; _ = METHOD_NAME
        PreLogUtil_Put(this.CommonStrings.START, this, METHOD_NAME)
        var isDefault bool = J2MEUtil_IsHTML(); _ = isDefault
        if this.GameCanvas != NullGameCanvas_GetInstance() && (this.GameCanvas.IsRunning() || isDefault || SWTUtil_IsSWT) && !(this.GameCanvas.GetType() == NullGameCanvas_TYPE) {
            this.GameRunnable = this.GameCanvas.AllBinaryGameCanvas_().GameRunnable
            PreLogUtil_Put("Showing Game", this, METHOD_NAME)
            this.SetPaintableInterface(this.GameCanvas)
        } else {
            this.GameRunnable = NullWaitGameRunnable_GetInstance()
            PreLogUtil_Put("Not Showing Game", this, METHOD_NAME)
            this.SetPaintableInterface(this.GetDefaultPaintableInterface())
        }
}


func (this *DemoCanvas) IsReadyForStateChange() bool {
        return !this.DemoGameRunnable.IsRunning() && this.GameCanvas.IsInitialized()
}


// throws Exception
func (this *DemoCanvas) ProcessGame() {
        this.GameRunnable.Run()
        if !this.SpecialAnimationInterface.IsComplete() && this.IsReadyForStateChange() {
            this.SpecialAnimationInterface.NextFrame()
        }
        if this.TimeDelayHelper.IsTimeTNT() && this.IsReadyForStateChange() {
            this.DemoStateChange()
        } else {
            this.Process()
        }
}


func (this *DemoCanvas) Run() {
        this.LogUtil.PutF(this.CommonStrings.START_RUNNABLE, this, this.CommonStrings.RUN)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            var features *Features = Features_GetInstance(); _ = features
            var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
            progressCanvas.AddNormalPortion(50, "Demo Thread")
            this.SetCurrentThread()
            this.SetRunning(true)
            if features.IsFeature(MainFeatureFactory_GetInstance().LOAD_ONDEMAND) {
                progressCanvas.End()
            } else {
                progressCanvas.AddNormalPortion(50, "Demo Thread Running")
            }
            this.FullScreenUtil.InitOnRun(this, this.GetCustomCommandListener())
            this.InitMenu()
            this.InitPostPaint()
            this.UpdateDemoState()
            if features.IsDefault(openGLFeatureFactory.OPENGL_AS_GAME_THREAD) {
                for this.GameCanvas == NullGameCanvas_GetInstance() || !this.GameCanvas.IsInitialized() {
                    this.LoopTimeHelper.SetStartTimeTNT()
                    this.ProcessGame()
                    this.ProcessLoopSleep()
                }
                var demoGameRunnable *DemoGameRunnable = NewDemoGameRunnable(this); _ = demoGameRunnable
                var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
                currentDisplayableFactory.SetRunnable(demoGameRunnable)
                currentDisplayableFactory.SetMyCanvas(this)
                OpenGLThreadUtil_GetInstance().OnResume()
            }
            if features.IsDefault(openGLFeatureFactory.OPENGL_AS_GAME_THREAD) || J2MEUtil_IsHTML() {
                var demoGameRunnable *DemoGameRunnable = NewDemoGameRunnable(this); _ = demoGameRunnable
                var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
                currentDisplayableFactory.SetRunnable(demoGameRunnable)
                currentDisplayableFactory.SetMyCanvas(this)
            } else {
                for this.IsRunning() {
                    this.Run3()
                }
                this.End()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
        this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
}


// throws Exception
func (this *DemoCanvas) Run3() {
        this.LoopTimeHelper.SetStartTimeTNT()
        this.ProcessGame()
        this.ProcessLoopSleep()
}


func (this *DemoCanvas) SetRunning(running bool) {
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
            var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
            if running {
            } else {
                if (features.IsDefault(openGLFeatureFactory.OPENGL) || J2MEUtil_IsHTML()) || SWTUtil_IsSWT {
                    var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
                    currentDisplayableFactory.ClearRunnable()
                    this.End()
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *DemoCanvas) End() {
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        var features *Features = Features_GetInstance(); _ = features
        this.BaseGameStatistics.Add(fmt.Sprint(NewStringMaker().Append(DemoCanvas_BOT_GAME_STATS).Append(fmt.Sprint(this.BaseGameStatistics)).Append(CommonSeps_GetInstance().NEW_LINE)))
        this.BaseGameStatistics.Init()
        if features.IsFeature(MainFeatureFactory_GetInstance().LOAD_ONDEMAND) {
            progressCanvas.Start()
        }
        this.LogUtil.PutF("Demo End", this, this.CommonStrings.RUN)
        this.Close()
        DisplayChangeEventHandler_GetInstance().RemoveListener(this)
        this.StopGameDemo()
}


func (this *DemoCanvas) SetGameState(gameState *GameState) {
}


func (this *DemoCanvas) GetGameState() *GameState {
        return this.GameStateFactory.PLAYING_GAME_STATE
}


func (this *DemoCanvas) GetGameCanvasRunnableInterface() IAllBinaryGameCanvas {
        return this.GameCanvas
}


func (this *DemoCanvas) IsDemoLoading() bool {
        var gameCanvas IAllBinaryGameCanvas = this.GameCanvas; _ = gameCanvas
        if gameCanvas == NullGameCanvas_GetInstance() {
            return true
        } else if gameCanvas.IsInitialized() {
            return false
        }
        return true
}


func (this *DemoCanvas) GetState() int {
        return this.State
}


func (this *DemoCanvas) SetState(state int) {
        this.State = state
}


func (this *DemoCanvas) IsHighScoreSubmitted() bool {
        this.LogUtil.PutF("Wow the AI got a high score!", this, "isHighScoreSubmitted")
        return false
}


func (this *DemoCanvas) GetRealHighScoresPaintable() *HighScoresPaintable {
        return this.RealHighScoresPaintable
}


func (this *DemoCanvas) SetSpecialAnimationInterface(specialAnimationInterface ISpecialAnimation) {
        specialAnimationInterface.SetFrame(0)
        this.SpecialAnimationInterface = specialAnimationInterface
}


func (this *DemoCanvas) GetSpecialAnimationInterface() ISpecialAnimation {
        return this.SpecialAnimationInterface
}


func (this *DemoCanvas) SetPaintableInterface(paintableInterface PaintableInterface) {
        this.PaintableInterface = paintableInterface
}


func (this *DemoCanvas) GetPaintableInterface() PaintableInterface {
        return this.PaintableInterface
}


func (this *DemoCanvas) SetDefaultPaintableInterface(defaultPaintableInterface IPaintable) {
        this.DefaultPaintableInterface = defaultPaintableInterface
}


func (this *DemoCanvas) GetDefaultPaintableInterface() IPaintable {
        return this.DefaultPaintableInterface
}


func (this *DemoCanvas) GetHighScoresFactoryInterface() HighScoresFactoryInterface {
        return this.HighScoresFactoryInterface
}


func (this *DemoCanvas) SetMenuInputProcessor(menuInputProcessor IBasicMenuInputProcessor) {
        this.MenuInputProcessor = menuInputProcessor
}


func (this *DemoCanvas) GetMenuInputProcessor() IBasicMenuInputProcessor {
        return this.MenuInputProcessor
}


func (this *DemoCanvas) GetMenuForm() IScrollSelectionForm {
        return this.MenuForm
}


func (this *DemoCanvas) SetMenuForm(menuForm IScrollSelectionForm) {
        this.MenuForm = menuForm
}


func (this *DemoCanvas) IsInitialized() bool {
        return this.Initialized
}


func (this *DemoCanvas) GetOverlayPaintable() IPaintable {
        return this.OverlayPaintable
}


func (this *DemoCanvas) SetBasicGameDemoPaintable(basicGameDemoPaintable IStatePaintable) {
        this.BasicGameDemoPaintable = basicGameDemoPaintable
}


func (this *DemoCanvas) GetBasicGameDemoPaintable() IStatePaintable {
        return this.BasicGameDemoPaintable
}


func (this *DemoCanvas) SetTempWait(tempWait int) {
        this.TempWait = tempWait
}


func (this *DemoCanvas) GetTempWait() int {
        return this.TempWait
}


func (this *DemoCanvas) IsSingleThread() bool {
        return OpenGLFeatureUtil_GetInstance().IsAnyThreed() || SWTUtil_IsSWT
}


func (this *DemoCanvas) IsRunningInAnotherThread() bool {
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        if features.IsDefault(openGLFeatureFactory.OPENGL_AS_GAME_THREAD) {
            return true
        } else {
            return this.IsRunning()
        }
}


func (this *DemoCanvas) GetGameInitializationInterfaceFactoryInterface() *BasicBuildGameInitializerFactory {
        return this.GameInitializationInterfaceFactoryInterface
}


func (this *DemoCanvas) GetType() int {
        return DemoCanvas_TYPE
}

