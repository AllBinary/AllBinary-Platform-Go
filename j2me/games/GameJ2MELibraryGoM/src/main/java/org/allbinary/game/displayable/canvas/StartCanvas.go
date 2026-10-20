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

type StartCanvas struct {
    RunnableCanvas
    BasicColorFactory *BasicColorFactory
    Features *Features
    OpenGLFeatureFactory *OpenGLFeatureFactory
    MyCommandsFactory *MyCommandsFactory
    GameAdStateFactory *GameAdStateFactory
    GameStateFactory *GameStateFactory
    GameInputStrings *GameInputStrings
    FormUtil *FormUtil
    BasicGameDemoPaintable IStatePaintable
    SpecialAnimationInterface ISpecialAnimation
    PaintedSpecialAnimationInterface IAnimation
    FullScreenUtil *FullScreenUtil
    HighScoresHelper *HighScoresHelper
    RealHighScoresPaintable *HighScoresPaintable
    HighScoresPaintable IPaintable
    State int
    TimeDelayHelper ITimeDelayHelper
    HighScoresFactoryInterface HighScoresFactoryInterface
    MenuInputProcessor IBasicMenuInputProcessor
    MenuForm IScrollSelectionForm
    GameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory
    Initialized bool
    DefaultPaintableInterface IPaintable
    PaintableInterface PaintableInterface
    OverlayPaintable IInitUpdatePaintable
    TempWait int
    InputToGameKeyMapping *InputToGameKeyMapping
    GameRunnable IGameRunnable
    ProgressPaintable PaintableInterface
    IsSingleKeyRepeatableProcessing bool
    NONE *GameKey
    GameKeyEventFactory *GameKeyEventFactory
    SET_STATE string
    BaseGameStatistics IBaseGameStatistics
}

// Implements GameCanvasRunnableInterface, MenuListener, DisplayChangeEventListener

var StartCanvas_Id int = 0
var StartCanvas_BOT_GAME_STATS string = "Bot Game Statistics: "

func (this *StartCanvas) initFields() {
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.Features = Features_GetInstance()
    this.OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance()
    this.MyCommandsFactory = MyCommandsFactory_GetInstance()
    this.GameAdStateFactory = GameAdStateFactory_GetInstance()
    this.GameStateFactory = GameStateFactory_GetInstance()
    this.GameInputStrings = GameInputStrings_GetInstance()
    this.FormUtil = FormUtil_GetInstance()
    this.BasicGameDemoPaintable = StatePaintableFactory_GetInstance()
    this.SpecialAnimationInterface = SpecialAnimation_GetInstance()
    this.PaintedSpecialAnimationInterface = SpecialAnimation_GetInstance()
    this.FullScreenUtil = FullScreenUtil_GetInstance()
    this.HighScoresHelper = NewHighScoresHelper()
    this.RealHighScoresPaintable = NewHighScoresPaintable()
    this.HighScoresPaintable = NullPaintable_GetInstance()
    this.State = 0
    this.TimeDelayHelper = NewTimeDelayHelper(6000)
    
    this.MenuInputProcessor = NoMenuInputProcessor_GetInstance()
    this.MenuForm = ScrollSelectionForm_GetNullScrollSelectionForm()
    
    
    this.DefaultPaintableInterface = NullPaintable_GetInstance()
    this.PaintableInterface = NullPaintable_GetInstance()
    
    this.TempWait = NullWaitGameRunnable_GetInstance().GameRunnable_().WAIT
    this.InputToGameKeyMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().GetInputMapping()
    this.GameRunnable = NullWaitGameRunnable_GetInstance()
    this.ProgressPaintable = ProgressCanvasFactory_GetLazyInstance()
    this.IsSingleKeyRepeatableProcessing = this.Features.IsFeature(InputFeatureFactory_GetInstance().SINGLE_KEY_REPEAT_PRESS)
    this.NONE = GameKeyFactory_GetInstance().NONE
    this.GameKeyEventFactory = GameKeyEventFactory_GetInstance()
    this.SET_STATE = "setState"
    this.BaseGameStatistics = GameStatisticsFactory_GetInstance()
}


func NewStartCanvas(abeClientInformation AbeClientInformationInterface, commandListener CommandListener, highScoresFactoryInterface HighScoresFactoryInterface, paintable IPaintable, overlayPaintable IInitUpdatePaintable, gameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory, isContinue bool) *StartCanvas {
    this := &StartCanvas{}
        this.RunnableCanvas = *NewRunnableCanvas(commandListener, CurrentDisplayableFactory_GetInstance().DEFAULT_CHILD_NAME_LIST, true)
        this.initFields()
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
        DisplayChangeEventHandler_GetInstance().AddListenerInterface(this)
    return this
}


func (this *StartCanvas) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *StartCanvas) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT, e)
                }
            }()
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


func (this *StartCanvas) GetCustomCommands() []any {
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


func (this *StartCanvas) InitCommands(cmdListener CommandListener) {
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
func (this *StartCanvas) InitPostPaint() {
}


// throws Exception
func (this *StartCanvas) MediaInit() {
        AllBinaryMediaManager_Init(EarlySoundsFactory_GetInstance())
}


func (this *StartCanvas) ItemStateChanged(item IItem) {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
}


// throws Exception
func (this *StartCanvas) InitMenu() {
        this.Close()
        this.Open()
}


func (this *StartCanvas) Open() {
        BasicMotionGesturesHandler_GetInstance().AddListenerInterface(this.GetMenuInputProcessor())
        GameKeyEventHandler_GetInstance().AddListener(this.GetMenuInputProcessor())
}


func (this *StartCanvas) Close() {
        BasicMotionGesturesHandler_GetInstance().RemoveListener(this.GetMenuInputProcessor())
        GameKeyEventHandler_GetInstance().RemoveListener(this.GetMenuInputProcessor())
}


func (this *StartCanvas) GetSourceId() int {
        return StartCanvas_Id
}


func (this *StartCanvas) KeyPressed(keyCode int) {
        this.KeyPressedByDevice(keyCode, 0)
}


func (this *StartCanvas) KeyReleased(keyCode int) {
        this.KeyReleasedByDevice(keyCode, 0)
}


func (this *StartCanvas) KeyRepeated(keyCode int) {
        this.KeyRepeatedByDevice(keyCode, 0)
}


func (this *StartCanvas) KeyPressedByDevice(keyCode int, deviceId int) {
        this.AddGameKeyEvent(keyCode, false)
}


func (this *StartCanvas) KeyReleasedByDevice(keyCode int, deviceId int) {
        this.RemoveGameKeyEvent(keyCode, false)
}


func (this *StartCanvas) KeyRepeatedByDevice(keyCode int, deviceId int) {
        if this.IsSingleKeyRepeatableProcessing {
            this.AddGameKeyEvent(keyCode, true)
        }
}


func (this *StartCanvas) AddGameKeyEvent(keyCode int, repeated bool) {
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
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GameInputStrings.NO_KEY).Appendint(keyCode)), this, this.GameInputStrings.ADD_KEY_EVENT)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *StartCanvas) RemoveGameKeyEvent(keyCode int, repeated bool) {
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
func (this *StartCanvas) Pause() {
        this.Close()
        this.SetPaused(true)
}


// synchronized
func (this *StartCanvas) UnPause() {
        this.Open()
        this.SetPaused(false)
}


func (this *StartCanvas) IsPausable() bool {
        if CurrentDisplayableFactory_GetInstance().GetUsedRunnable() == NullWaitGameRunnable_GetInstance() {
            return true
        } else {
            return false
        }
}


func (this *StartCanvas) IsGameOver() bool {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.NOT_IMPLEMENTED).Append(" since not a game")), this, "isGameOver")
        return false
}


// throws Exception
func (this *StartCanvas) SetLoadStateHashtable(hashtable *ABHashtable[any, any]) {
        this.LogUtil.PutF("Trying to continue a demo lol - only continue a game canvas not the demo", this, "setLoadStateHashtable")
}


// throws Exception
func (this *StartCanvas) GetLoadStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF("Trying to continue a demo lol - only continue a game canvas not the demo", this, "getLoadStateHashtable")
        return this.StdUtil.NULL_TABLE
}


// throws Exception
func (this *StartCanvas) GetCurrentStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF("Trying to save the AI lol", this, "getCurrentStateHashtable")
        return this.StdUtil.NULL_TABLE
}


func (this *StartCanvas) SetHighScoreSubmitted(isNotUsed bool) {
}


func (this *StartCanvas) Paint(graphics IGraphics) {
        this.PaintableInterface.Paint(graphics)
        this.PaintedSpecialAnimationInterface.PaintXY(graphics, 0, 0)
        this.HighScoresPaintable.Paint(graphics)
        this.GetBasicGameDemoPaintable().Paint(graphics)
        this.OverlayPaintable.Paint(graphics)
        this.ProgressPaintable.Paint(graphics)
}


func (this *StartCanvas) PaintThreed(graphics IGraphics) {
        this.PaintableInterface.PaintThreed(graphics)
        this.PaintedSpecialAnimationInterface.PaintThreedXYZ(graphics, 0, 0, 0)
}


// synchronized
func (this *StartCanvas) ProcessGameOver() {
        this.LogUtil.PutF("Not Implemented since not a game", this, "setGameOver")
}


func (this *StartCanvas) DemoStateChange() {
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


func (this *StartCanvas) UpdateDemoState() {
        PreLogUtil_Put(fmt.Sprint(SmallIntegerSingletonFactory_GetInstance().CreateInstance(this.State)), this, this.SET_STATE)
        this.GetBasicGameDemoPaintable().SetState(this.State)
        if this.State == 0 {
            this.HighScoresPaintable = NullPaintable_GetInstance()
            this.PaintedSpecialAnimationInterface = this.GetSpecialAnimationInterface()
            this.GetSpecialAnimationInterface().Reset()
        } else if this.State == 1 {
            this.PaintedSpecialAnimationInterface = SpecialAnimation_GetInstance()
        } else if this.State == 2 {
            this.HighScoresPaintable = this.GetRealHighScoresPaintable()
        }
        var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
        gameAdState.ProcessPageAdState()
}


// throws Exception
func (this *StartCanvas) GetNextRandom() int {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *StartCanvas) CreateGameLayerManager(randomValue int) IAllBinaryGameLayerManager {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *StartCanvas) CreateRunnable(randomLevel int) GameCanvasRunnableInterface {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *StartCanvas) CreateCanvas() {
        this.HighScoresPaintable = NullPaintable_GetInstance()
        var randomLevel int = this.GetNextRandom(); _ = randomLevel
}


// throws Exception
func (this *StartCanvas) Start() {
}


func (this *StartCanvas) PreDemoProcess() {
        this.OverlayPaintable.Update()
}


// throws Exception
func (this *StartCanvas) Process() {
        this.GetMenuInputProcessor().ProcessInputList()
        this.PreDemoProcess()
        if this.State == 0 {
            var indexedAnimationBehavior *IndexedAnimationBehavior = any(this.GetSpecialAnimationInterface().GetAnimationBehavior()).(*IndexedAnimationBehavior); _ = indexedAnimationBehavior
            if indexedAnimationBehavior.LoopIndex < 1 {
                this.TimeDelayHelper.SetStartTimeTNT()
            }
            var demoGameMidlet IDemoGameMidlet = any(this.GetCustomCommandListener()).(IDemoGameMidlet); _ = demoGameMidlet
        }
}


// throws Exception
func (this *StartCanvas) StartDemoGame() {
}


// throws Exception
func (this *StartCanvas) ProcessGame() {
        this.GameRunnable.Run()
        this.SpecialAnimationInterface.NextFrame()
        this.Process()
}


func (this *StartCanvas) Run() {
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
            progressCanvas.AddNormalPortion(50, "Demo Thread")
            this.SetCurrentThread()
            this.SetRunning(true)
            if this.Features.IsFeature(MainFeatureFactory_GetInstance().LOAD_ONDEMAND) {
                progressCanvas.End()
            } else {
                progressCanvas.AddNormalPortion(50, "Demo Thread Running")
            }
            this.FullScreenUtil.InitOnRun(this, this.GetCustomCommandListener())
            this.InitMenu()
            this.InitPostPaint()
            this.UpdateDemoState()
            if this.Features.IsDefault(this.OpenGLFeatureFactory.OPENGL_AS_GAME_THREAD) {
                OpenGLThreadUtil_GetInstance().OnResume()
            }
            if this.Features.IsDefault(this.OpenGLFeatureFactory.OPENGL_AS_GAME_THREAD) || J2MEUtil_IsHTML() {
            } else {
                for this.IsRunning() {
                    this.LoopTimeHelper.SetStartTimeTNT()
                    this.ProcessGame()
                    this.ProcessLoopSleep()
                }
                this.End()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
        this.LogUtil.PutF(this.CommonStrings.END_RUNNABLE, this, this.CommonStrings.RUN)
}


func (this *StartCanvas) SetRunning(running bool) {
        this.RunnableCanvas.SetRunning(running)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.SET_RUNNING, e)
                }
            }()
            if (this.Features.IsDefault(this.OpenGLFeatureFactory.OPENGL) || J2MEUtil_IsHTML()) && !running {
                var currentDisplayableFactory *CurrentDisplayableFactory = CurrentDisplayableFactory_GetInstance(); _ = currentDisplayableFactory
                currentDisplayableFactory.ClearRunnable()
                this.End()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *StartCanvas) End() {
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        this.BaseGameStatistics.Add(fmt.Sprint(NewStringMaker().Append(StartCanvas_BOT_GAME_STATS).Append(fmt.Sprint(this.BaseGameStatistics)).Append(CommonSeps_GetInstance().NEW_LINE)))
        this.BaseGameStatistics.Init()
        if this.Features.IsFeature(MainFeatureFactory_GetInstance().LOAD_ONDEMAND) {
            progressCanvas.Start()
        }
        this.LogUtil.PutF("Demo End", this, this.CommonStrings.RUN)
        this.Close()
        DisplayChangeEventHandler_GetInstance().RemoveListener(this)
}


func (this *StartCanvas) SetGameState(gameState *GameState) {
}


func (this *StartCanvas) GetGameState() *GameState {
        return this.GameStateFactory.PLAYING_GAME_STATE
}


func (this *StartCanvas) GetState() int {
        return this.State
}


func (this *StartCanvas) SetState(state int) {
        this.State = state
}


func (this *StartCanvas) IsHighScoreSubmitted() bool {
        this.LogUtil.PutF("Wow the AI got a high score!", this, "isHighScoreSubmitted")
        return false
}


func (this *StartCanvas) GetRealHighScoresPaintable() *HighScoresPaintable {
        return this.RealHighScoresPaintable
}


func (this *StartCanvas) SetSpecialAnimationInterface(specialAnimationInterface ISpecialAnimation) {
        specialAnimationInterface.SetFrame(0)
        this.SpecialAnimationInterface = specialAnimationInterface
}


func (this *StartCanvas) GetSpecialAnimationInterface() ISpecialAnimation {
        return this.SpecialAnimationInterface
}


func (this *StartCanvas) SetPaintableInterface(paintableInterface PaintableInterface) {
        this.PaintableInterface = paintableInterface
}


func (this *StartCanvas) GetPaintableInterface() PaintableInterface {
        return this.PaintableInterface
}


func (this *StartCanvas) SetDefaultPaintableInterface(defaultPaintableInterface IPaintable) {
        this.DefaultPaintableInterface = defaultPaintableInterface
}


func (this *StartCanvas) GetDefaultPaintableInterface() IPaintable {
        return this.DefaultPaintableInterface
}


func (this *StartCanvas) GetHighScoresFactoryInterface() HighScoresFactoryInterface {
        return this.HighScoresFactoryInterface
}


func (this *StartCanvas) SetMenuInputProcessor(menuInputProcessor IBasicMenuInputProcessor) {
        this.MenuInputProcessor = menuInputProcessor
}


func (this *StartCanvas) GetMenuInputProcessor() IBasicMenuInputProcessor {
        return this.MenuInputProcessor
}


func (this *StartCanvas) GetMenuForm() IScrollSelectionForm {
        return this.MenuForm
}


func (this *StartCanvas) SetMenuForm(menuForm IScrollSelectionForm) {
        this.MenuForm = menuForm
}


func (this *StartCanvas) IsInitialized() bool {
        return this.Initialized
}


func (this *StartCanvas) GetOverlayPaintable() IPaintable {
        return this.OverlayPaintable
}


func (this *StartCanvas) SetBasicGameDemoPaintable(basicGameDemoPaintable IStatePaintable) {
        this.BasicGameDemoPaintable = basicGameDemoPaintable
}


func (this *StartCanvas) GetBasicGameDemoPaintable() IStatePaintable {
        return this.BasicGameDemoPaintable
}


func (this *StartCanvas) SetTempWait(tempWait int) {
        this.TempWait = tempWait
}


func (this *StartCanvas) GetTempWait() int {
        return this.TempWait
}


func (this *StartCanvas) IsSingleThread() bool {
        return OpenGLFeatureUtil_GetInstance().IsAnyThreed() || SWTUtil_IsSWT
}


func (this *StartCanvas) IsRunningInAnotherThread() bool {
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        if features.IsDefault(openGLFeatureFactory.OPENGL_AS_GAME_THREAD) {
            return true
        } else {
            return this.IsRunning()
        }
}


func (this *StartCanvas) GetGameInitializationInterfaceFactoryInterface() *BasicBuildGameInitializerFactory {
        return this.GameInitializationInterfaceFactoryInterface
}

