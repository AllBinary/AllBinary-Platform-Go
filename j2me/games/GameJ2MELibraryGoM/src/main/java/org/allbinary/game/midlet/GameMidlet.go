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

type GameMidlet struct {
    ProgressMidlet
    EnumerationUtil *EnumerationUtil
    BasicColorFactory *BasicColorFactory
    Features *Features
    MidletStrings *MidletStrings
    MyCommandsFactory *MyCommandsFactory
    GameStrings *GameStrings
    GameAdStateFactory *GameAdStateFactory
    GameStateFactory *GameStateFactory
    TsUtil *TsUtil
    DISPLAYABLE string
    COMMAND_NAME string
    NO_COMMAND string
    NO_DISPLAYABLE string
    COMMAND_ACTION string
    PAUSE_APP_BACKGROUND string
    UN_PAUSE_APP_BACKGROUND string
    AboutCommandProcessor *AboutCommandProcessor
    WebCommandProcessor *WebCommandProcessor
    GameMidletStateFactory *GameMidletStateFactory
    GameStartTimeHelper ITimeDelayHelper
    FullScreenUtil *FullScreenUtil
    DebugInterface DebugInterface
    AllbinaryGameCanvasRunnableInterface GameCanvasRunnableInterface
    Thread IThread
    LoadGameForm ICommandForm
    IsFullScreen bool
    Resized bool
    StartedBefore bool
}

// Implements CommandListener

// The Java class GameMidlet has subclasses so its variables are this interface
type IGameMidlet interface {
    IProgressMidlet
    GameMidlet_() *GameMidlet
    Init()
    SetDemo()
    CreateGame()
    MediaShutdown()
    StopAll()
    PauseApp()
    PauseAppBackground(background bool)
    UnPauseApp()
    UnPauseAppBackground(background bool)
    DestroyAppInRunnable(unconditional bool, isProgress bool)
    DestroyApp(unconditional bool)
    StartApp()
    CommandAction(command ICommand, displayable IDisplayable)
    UpdateFullScreen()
    OnEvent(eventObject IAllBinaryEventObject)
    GetInputMappingInstructionsCanvas() IMyCanvas
    GetInputMappingCanvas() IMyCanvas
    GetAboutCanvas() IMyCanvas
    GetHelpPaintable() IHelpPaintable
    CreateHighScoresCanvas() *HighScoresCanvas
    StartGameCanvasRunnableInterface()
    StopGameCanvasRunnableInterface()
    GetGameCanvasRunnableInterface() GameCanvasRunnableInterface
    SetGameCanvasRunnableInterface(gameCanvasRunnableInterface GameCanvasRunnableInterface)
    CreateGameLayerManager() IAllBinaryGameLayerManager
    GetGameOptionsForm() ICommandForm
    Save()
    GetCurrentStateHashtable() *ABHashtable[any, any]
    SetLoadGameForm(loadGameForm *LoadGameForm)
    GetLoadGameForm() ICommandForm
    SetResized(resized bool)
    IsResized() bool
    IsDemoLoading() bool
}

func (this *GameMidlet) GameMidlet_() *GameMidlet {
    return this
}


func (this *GameMidlet) initFields() {
    this.EnumerationUtil = EnumerationUtil_GetInstance()
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.Features = Features_GetInstance()
    this.MidletStrings = MidletStrings_GetInstance()
    this.MyCommandsFactory = MyCommandsFactory_GetInstance()
    this.GameStrings = GameStrings_GetInstance()
    this.GameAdStateFactory = GameAdStateFactory_GetInstance()
    this.GameStateFactory = GameStateFactory_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.DISPLAYABLE = " Displayable: "
    this.COMMAND_NAME = "command Name/Label: "
    this.NO_COMMAND = "No Command"
    this.NO_DISPLAYABLE = "No Displayable"
    this.COMMAND_ACTION = fmt.Sprint(NewStringMaker().Append("GameMidlet::").Append(this.MidletStrings.COMMAND_ACTION))
    this.PAUSE_APP_BACKGROUND = "pauseAppBackground"
    this.UN_PAUSE_APP_BACKGROUND = "unPauseAppBackground"
    this.AboutCommandProcessor = AboutCommandProcessor_GetInstance()
    this.WebCommandProcessor = WebCommandProcessor_GetInstance()
    this.GameMidletStateFactory = GameMidletStateFactory_GetInstance()
    this.GameStartTimeHelper = NewTimeDelayHelper(240)
    this.FullScreenUtil = FullScreenUtil_GetInstance()
    
    this.AllbinaryGameCanvasRunnableInterface = NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE
    this.Thread = NullThread_NULL_THREAD
    
    
    
    this.StartedBefore = false
}


func NewGameMidlet(clientInformationFactory *ClientInformationFactory) *GameMidlet {
    this := &GameMidlet{}
        this.ProgressMidlet = *NewProgressMidlet(clientInformationFactory)
        this.initFields()
        SmallIntegerSingletonFactory_GetInstance().InitWithRange(0x291, 6)
        this.LoadGameForm = CommandForm_GetNullCommandForm()
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.Init(this)
        GameFeatureEventHandler_GetInstance().AddListenerInterface(ChangedGameFeatureListener_GetInstance())
        GamePersistanceSingleton_GetInstance().Clear()
        this.DebugInterface = DebugFactory_GetInstance()
        this.Init()
    return this
}


func (this *GameMidlet) Init() {
}


// throws Exception
func (this *GameMidlet) SetDemo() {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *GameMidlet) CreateGame() {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *GameMidlet) MediaShutdown() {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *GameMidlet) StopAll() {
}


func (this *GameMidlet) PauseApp() {
        this.PauseAppBackground(true)
        var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
        gameAdState.GetAdvertisements().StopAll()
}


func (this *GameMidlet) PauseAppBackground(background bool) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.PAUSE_APP_BACKGROUND)
        if this.AllbinaryGameCanvasRunnableInterface != NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE {
            this.AllbinaryGameCanvasRunnableInterface.Pause()
        } else {
            this.LogUtil.PutF("<<<<<< Null", this, this.PAUSE_APP_BACKGROUND)
        }
        AllBinarySensorManager_GetInstance().Shutdown()
}


func (this *GameMidlet) UnPauseApp() {
        this.UnPauseAppBackground(true)
        var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
        gameAdState.GetAdvertisements().StartAll()
}


func (this *GameMidlet) UnPauseAppBackground(background bool) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.UN_PAUSE_APP_BACKGROUND)
        AllBinarySensorManager_GetInstance().Init()
        var gameCanvasRunnableInterface GameCanvasRunnableInterface = this.AllbinaryGameCanvasRunnableInterface; _ = gameCanvasRunnableInterface
        if gameCanvasRunnableInterface != NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE {
            gameCanvasRunnableInterface.UnPause()
        } else {
            this.LogUtil.PutF("<<<<<< Null", this, this.UN_PAUSE_APP_BACKGROUND)
        }
}


func (this *GameMidlet) DestroyAppInRunnable(unconditional bool, isProgress bool) {
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        if isProgress {
            progressCanvas.Start()
            this.CommandAction(this.MyCommandsFactory.SET_DISPLAYABLE, progressCanvas)
        }
        this.DestroyApp(unconditional)
        if isProgress {
            progressCanvas.End()
        }
}


func (this *GameMidlet) DestroyApp(unconditional bool) {
        var METHOD_NAME string = "GameMidlet::destroyApp"; _ = METHOD_NAME
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, METHOD_NAME, e)
                }
            }()
            PreLogUtil_Put(fmt.Sprint(GameStatisticsFactory_GetInstance()), this, METHOD_NAME)
            var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
            gameAdState.GetAdvertisements().StopAll()
            if !this.IsDestroyed() {
                this.StopGameCanvasRunnableInterface()
                this.SetGameCanvasRunnableInterface(NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE)
                this.MediaShutdown()
            } else {
                this.LogUtil.PutF("Midlet Managment Error: Midlet Should Only Be Destroyed Once", this, METHOD_NAME)
            }
            this.ProgressMidlet.DestroyApp(true)
            PreLogUtil_Put(this.CommonStrings.END, this, METHOD_NAME)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.LogUtil.PutF(this.CommonStrings.END, this, METHOD_NAME)
}


// throws MIDletStateChangeException
func (this *GameMidlet) StartApp() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "startApp", e)
            this.DestroyApp(false)
            this.NotifyDestroyed()
                }
            }()
            var gameAdState *GameAdState = this.GameAdStateFactory.GetCurrentInstance(); _ = gameAdState
            gameAdState.GetAdvertisements().StartAll()
            var START_APP string = "startApp"; _ = START_APP
            this.LogUtil.PutF(this.CommonStrings.START, this, START_APP)
            var gameCanvasRunnableInterface any = this.AllbinaryGameCanvasRunnableInterface; _ = gameCanvasRunnableInterface
            if gameCanvasRunnableInterface == NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE {
                this.GameMidletStateFactory.SetCurrentGameState(this.GameStateFactory.NO_GAME_STATE)
                this.SetDemo()
            } else {
                this.LogUtil.PutF("GameCanvasRunnableInterface is available", this, START_APP)
                if gameCanvasRunnableInterface == this.GetCurrentDisplayable() {
                    this.UnPauseAppBackground(false)
                } else {
                    this.LogUtil.PutF("GameCanvasRunnableInterface is not current displayable", this, START_APP)
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *GameMidlet) CommandAction(command ICommand, displayable IDisplayable) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.EXCEPTION_LABEL).Append(command.GetLabel())), this, this.MidletStrings.COMMAND_ACTION, e)
            if command != GameCommandsFactory_GetInstance().EXIT_COMMAND {
                this.ExitProgress(false)
            }
                }
            }()
            var displayableAsString string = this.NO_DISPLAYABLE; _ = displayableAsString
            if displayable != nil {
                displayableAsString = fmt.Sprint(displayable)
            }
            var label string = this.NO_COMMAND; _ = label
            if command != nil {
                label = command.GetLabel()
            }
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.COMMAND_NAME).Append(label).Append(this.DISPLAYABLE).Append(displayableAsString)), this, this.COMMAND_ACTION)
            var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
            var allbinaryGameCanvasRunnableInterface GameCanvasRunnableInterface = this.AllbinaryGameCanvasRunnableInterface; _ = allbinaryGameCanvasRunnableInterface
            var gameCanvasRunnableInterface any = allbinaryGameCanvasRunnableInterface; _ = gameCanvasRunnableInterface
            if command == gameCommandsFactory.SHOW_GAME_CANVAS {
                if this.GetDisplay().GetCurrent() != gameCanvasRunnableInterface && allbinaryGameCanvasRunnableInterface.GetType() != NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE.GetType() {
                    this.SetDisplay(any(gameCanvasRunnableInterface).(IDisplayable))
                }
                this.UnPauseAppBackground(false)
            } else if command == gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS {
                var menuListener MenuListener = any(displayable).(MenuListener); _ = menuListener
                menuListener.Close()
                this.SetDisplay(any(gameCanvasRunnableInterface).(IDisplayable))
                this.UnPauseAppBackground(false)
            } else if command == gameCommandsFactory.EXIT_COMMAND {
                if GameMidletEventHandler_GetInstance().GetEventListenerInterfaceListP().Size() == 0 {
                    this.ExitProgress(true)
                }
                GameMidletEventHandler_GetInstance().FireEvent(NewDemoGameMidletEvent(this, DemoGameMidletStateFactory_GetInstance().NONE))
            } else if command == gameCommandsFactory.EXIT_WITHOUT_PROGRESS_COMMAND {
                this.ExitProgress(false)
            } else if command == gameCommandsFactory.START_COMMAND || command == gameCommandsFactory.RESTART_COMMAND || command == gameCommandsFactory.CONTINUE_COMMAND {
                if this.GameMidletStateFactory.GetCurrentGameState() != this.GameStateFactory.PLAYING_GAME_STATE || command == gameCommandsFactory.RESTART_COMMAND {
                    if this.GameStartTimeHelper.IsTimeTNT() {
                        if command == gameCommandsFactory.START_COMMAND && this.IsDemoLoading() {
                            this.LogUtil.PutF("Trying to Start Game Before Loading Complete", this, this.MidletStrings.COMMAND_ACTION)
                        } else {
                            this.StartedBefore = true
                            this.CreateGame()
                            this.GameMidletStateFactory.SetCurrentGameState(this.GameStateFactory.PLAYING_GAME_STATE)
                        }
                    } else {
                        this.LogUtil.PutF("Starting Game Too Often", this, this.MidletStrings.COMMAND_ACTION)
                    }
                } else {
                    this.LogUtil.PutF("Already in playing state", this, this.COMMAND_ACTION)
                }
            } else if command == gameCommandsFactory.BUY_COMMAND {
                var inApplicationPurchaseFactory *InApplicationPurchaseFactory = InApplicationPurchaseFactory_GetInstance(); _ = inApplicationPurchaseFactory
                var list IBasicArrayList = LockableFeatureFactory_GetInstance().GetList(); _ = list
                var lockableFeature *LockableFeature = any(list.Get(0)).(*LockableFeature); _ = lockableFeature
                if list.Size() > 0 && !inApplicationPurchaseFactory.IsPurchased(lockableFeature) {
                    inApplicationPurchaseFactory.Purchase(lockableFeature)
                } else {
                    TextNotificationUtil_GetInstance().FireSuccess("Already Purchased")
                }
            } else if command == gameCommandsFactory.QUIT_COMMAND {
                if this.GameStartTimeHelper.IsTimeTNT() {
                    if func() bool { _, ok := any(displayable).(*HighScoreTextBox); return ok }() {
                        var menuListener MenuListener = any(displayable).(MenuListener); _ = menuListener
                        menuListener.Close()
                    }
                    this.StopGameCanvasRunnableInterface()
                    this.GameMidletStateFactory.SetCurrentGameState(this.GameStateFactory.NO_GAME_STATE)
                    this.SetDemo()
                }
            } else if command == this.MyCommandsFactory.RESUME_COMMAND {
                this.UnPauseAppBackground(false)
            } else if command == this.MyCommandsFactory.PAUSE_COMMAND {
                this.PauseAppBackground(false)
            } else if command == this.MyCommandsFactory.SET_DISPLAYABLE {
                this.PauseAppBackground(false)
                this.SetDisplay(displayable)
            } else if command == gameCommandsFactory.SET_MENU_DISPLAYABLE {
                this.PauseAppBackground(false)
                var menuListener MenuListener = any(displayable).(MenuListener); _ = menuListener
                menuListener.Open()
                this.SetDisplay(displayable)
            } else if command == gameCommandsFactory.START_TRACE {
                this.DebugInterface.Start()
                var gameCanvas IAllBinaryGameCanvas = any(allbinaryGameCanvasRunnableInterface).(IAllBinaryGameCanvas); _ = gameCanvas
                gameCanvas.AddCommand(gameCommandsFactory.STOP_TRACE)
                gameCanvas.RemoveCommand(gameCommandsFactory.START_TRACE)
            } else if command == gameCommandsFactory.STOP_TRACE {
                this.PauseAppBackground(false)
                this.DebugInterface.Stop()
                this.UnPauseAppBackground(false)
                var gameCanvas IAllBinaryGameCanvas = any(allbinaryGameCanvasRunnableInterface).(IAllBinaryGameCanvas); _ = gameCanvas
                gameCanvas.AddCommand(gameCommandsFactory.START_TRACE)
                gameCanvas.RemoveCommand(gameCommandsFactory.STOP_TRACE)
            } else if command == gameCommandsFactory.DEFAULT_OPTIONS {
                var tempDisplayable IDisplayable = displayable; _ = tempDisplayable
                if func() bool { _, ok := any(tempDisplayable).(*GameOptionsForm); return ok }() {
                    GameFeatureFormUtil_GetInstance().SetDefault(any(tempDisplayable).(ICommandForm))
                }
            } else if command == GameInputMappingInstructionsCanvas_DISPLAY {
                var tempDisplayable IDisplayable = displayable; _ = tempDisplayable
                if func() bool { _, ok := any(tempDisplayable).(*GameInputMappingCanvas); return ok }() {
                    var gameInputMappingCanvas *GameInputMappingCanvas = any(tempDisplayable).(*GameInputMappingCanvas); _ = gameInputMappingCanvas
                    gameInputMappingCanvas.Update()
                }
                this.CommandAction(gameCommandsFactory.SET_MENU_DISPLAYABLE, this.GetInputMappingInstructionsCanvas())
            } else if HighScoreCommandsFactory_GetInstance().IsHighScoreCommand(command) {
                var tempDisplayable IDisplayable = displayable; _ = tempDisplayable
                if func() bool { _, ok := any(tempDisplayable).(*HighScoresCanvas); return ok }() {
                    var highScoresCanvas *HighScoresCanvas = any(tempDisplayable).(*HighScoresCanvas); _ = highScoresCanvas
                    highScoresCanvas.UpdateCommand(command)
                }
            } else if command == HighScoreCommands_GetInstance().DISPLAY {
                this.CommandAction(gameCommandsFactory.SET_MENU_DISPLAYABLE, this.CreateHighScoresCanvas())
            } else if command == GameInputMappingInstructionsCanvas_CLOSE {
                var menuListener MenuListener = any(displayable).(MenuListener); _ = menuListener
                menuListener.Close()
                this.CommandAction(GameInputMappingCanvas_DISPLAY, NullCanvas_NULL_CANVAS)
            } else if command == GameInputMappingCanvas_DISPLAY {
                if J2MEUtil_IsHTML() {
                    this.PauseAppBackground(false)
                }
                this.CommandAction(gameCommandsFactory.SET_MENU_DISPLAYABLE, this.GetInputMappingCanvas())
            } else if command == GameInputMappingCanvas_DEFAULT {
                var tempDisplayable IDisplayable = displayable; _ = tempDisplayable
                if func() bool { _, ok := any(tempDisplayable).(*GameInputMappingCanvas); return ok }() {
                    var gameInputMappingCanvas *GameInputMappingCanvas = any(tempDisplayable).(*GameInputMappingCanvas); _ = gameInputMappingCanvas
                    gameInputMappingCanvas.SetDefault()
                }
            } else if command == InGameOptionsForm_DISPLAY {
                this.PauseAppBackground(false)
                var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
                progressCanvas.AddNormalPortion(50, "In Game Options")
                var layerManager IAllBinaryGameLayerManager = this.CreateGameLayerManager(); _ = layerManager
                var inGameOptionsFormFactory *InGameOptionsFormFactory = InGameOptionsFormFactory_GetInstance(); _ = inGameOptionsFormFactory
                inGameOptionsFormFactory.Init(this, NewInGameFeatures(), "Options In Game", layerManager.GetBackgroundBasicColor(), layerManager.GetForegroundBasicColor())
                var inGameOptionsForm ICommandForm = inGameOptionsFormFactory.Get(); _ = inGameOptionsForm
                inGameOptionsForm.SetItemStateListener(allbinaryGameCanvasRunnableInterface)
                this.CommandAction(gameCommandsFactory.SET_MENU_DISPLAYABLE, inGameOptionsForm)
            } else if command == InGameOptionsForm_DEFAULT {
                var inGameOptionsForm *InGameOptionsForm = any(InGameOptionsFormFactory_GetInstance().Get()).(*InGameOptionsForm); _ = inGameOptionsForm
                GameFeatureFormUtil_GetInstance().SetDefault(inGameOptionsForm)
            } else if command == gameCommandsFactory.DISPLAY_OPTIONS {
                AllBinaryMediaManager_SetMuted(true)
                this.StopAll()
                var mainFeatureFactory *MainFeatureFactory = MainFeatureFactory_GetInstance(); _ = mainFeatureFactory
                this.IsFullScreen = this.Features.IsFeature(mainFeatureFactory.FULL_SCREEN)
                ResizableListenerHandler_GetInstance().FireEvent(true)
                this.SetResized(false)
                this.CommandAction(gameCommandsFactory.SET_MENU_DISPLAYABLE, this.GetGameOptionsForm())
            } else if command == gameCommandsFactory.CLOSE_OPTIONS {
                ResizableListenerHandler_GetInstance().FireEvent(false)
                this.CommandAction(gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS, displayable)
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Close isFullScreen/change: ")
                stringBuffer.Appendboolean(this.IsFullScreen)
                stringBuffer.Appendboolean(this.FullScreenUtil.IsScreenChange(this.IsFullScreen))
                stringBuffer.Append(" isResized: ")
                stringBuffer.Appendboolean(this.IsResized())
                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, this.COMMAND_ACTION)
                if this.IsResized() || this.FullScreenUtil.IsScreenChange(this.IsFullScreen) {
                    this.UpdateFullScreen()
                }
                AllBinaryMediaManager_SetMuted(false)
            } else if command == gameCommandsFactory.DISPLAY_LOAD_FORM {
                this.PauseAppBackground(false)
                var keyValuePersistance IKeyValuePersistance = GamePersistanceSingleton_GetInstance(); _ = keyValuePersistance
                keyValuePersistance.Clear()
                keyValuePersistance.LoadAll(this.AbeClientInformation)
                if this.GetLoadGameForm() == CommandForm_GetNullCommandForm() {
                    var layerManager IAllBinaryGameLayerManager = this.CreateGameLayerManager(); _ = layerManager
                    this.SetLoadGameForm(NewLoadGameForm(this, "Load Game", layerManager.GetBackgroundBasicColor(), layerManager.GetForegroundBasicColor()))
                } else {
                    this.GetLoadGameForm().Update()
                }
                this.CommandAction(gameCommandsFactory.SET_MENU_DISPLAYABLE, this.GetLoadGameForm())
            } else if command == gameCommandsFactory.LOAD_FILE {
                var loadGameForm *LoadGameForm = any(this.GetLoadGameForm()).(*LoadGameForm); _ = loadGameForm
                var index int = loadGameForm.GetSelectedId(); _ = index
                if index != -1 {
                    var keyValuePersistance IKeyValuePersistance = GamePersistanceSingleton_GetInstance(); _ = keyValuePersistance
                    this.SetStartStateHashtable(keyValuePersistance.Get(index))
                    var menuListener MenuListener = this.GetLoadGameForm(); _ = menuListener
                    menuListener.Close()
                    PreLogUtil_Put(fmt.Sprint(BasicMotionGesturesHandler_GetInstance()), this, this.COMMAND_ACTION)
                    this.CommandAction(gameCommandsFactory.START_COMMAND, NullCanvas_NULL_CANVAS)
                    PreLogUtil_Put(fmt.Sprint(BasicMotionGesturesHandler_GetInstance()), this, this.COMMAND_ACTION)
                }
            } else if command == gameCommandsFactory.DELETE_FILE {
                var loadGameForm *LoadGameForm = any(this.GetLoadGameForm()).(*LoadGameForm); _ = loadGameForm
                var index int = loadGameForm.GetSelectedText(); _ = index
                if index != -1 {
                    var keyValuePersistance IKeyValuePersistance = GamePersistanceSingleton_GetInstance(); _ = keyValuePersistance
                    keyValuePersistance.Delete(this.AbeClientInformation, index)
                    keyValuePersistance.Clear()
                    keyValuePersistance.LoadAllSize(this.AbeClientInformation, 1)
                    this.GetLoadGameForm().Update()
                }
            } else if command == gameCommandsFactory.SAVE {
                var keyValuePersistance IKeyValuePersistance = GamePersistanceSingleton_GetInstance(); _ = keyValuePersistance
                this.PauseAppBackground(false)
                keyValuePersistance.Save(this.AbeClientInformation, this.GetCurrentStateHashtable())
                this.UnPauseAppBackground(false)
            } else if command == HighScoreUtil_SUBMIT_TEXTBOX_COMMAND {
                this.LogUtil.PutF("Submitted Score", this, this.COMMAND_ACTION)
                if func() bool { _, ok := any(displayable).(*HighScoreTextBox); return ok }() {
                    var highScoreTextBox *HighScoreTextBox = any(displayable).(*HighScoreTextBox); _ = highScoreTextBox
                    highScoreTextBox.Submitted = true
                }
                allbinaryGameCanvasRunnableInterface.SetHighScoreSubmitted(true)
                this.CommandAction(gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS, displayable)
            } else if command == gameCommandsFactory.DISPLAY_ABOUT {
                this.AboutCommandProcessor.Process(this, gameCommandsFactory.SET_MENU_DISPLAYABLE, this.GetAboutCanvas())
            } else if command == gameCommandsFactory.OPEN_WEB_URL {
                this.WebCommandProcessor.Process(this, gameCommandsFactory.OPEN_WEB_URL, NullCanvas_NULL_CANVAS)
            } else if command == gameCommandsFactory.TOGGLE_KEYBOARD {
                var virtualKeyboardEventHandler *VirtualKeyboardEventHandler = VirtualKeyboardEventHandler_GetInstance(); _ = virtualKeyboardEventHandler
                virtualKeyboardEventHandler.FireEvent(virtualKeyboardEventHandler.SHOW_EVENT)
            } else if this.TsUtil.CompareTo(command.GetLabel(), gameCommandsFactory.TOGGLE_FULLSCREEN.GetLabel()) == 0 {
                var mainFeatureFactory *MainFeatureFactory = MainFeatureFactory_GetInstance(); _ = mainFeatureFactory
                var isFullScreen bool = this.Features.IsFeature(mainFeatureFactory.FULL_SCREEN); _ = isFullScreen
                if isFullScreen {
                    this.Features.RemoveDefault(mainFeatureFactory.FULL_SCREEN)
                } else {
                    this.Features.AddDefault(mainFeatureFactory.FULL_SCREEN)
                }
                this.UpdateFullScreen()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameMidlet) UpdateFullScreen() {
        var mainFeatureFactory *MainFeatureFactory = MainFeatureFactory_GetInstance(); _ = mainFeatureFactory
        var isFullScreen bool = this.Features.IsFeature(mainFeatureFactory.FULL_SCREEN); _ = isFullScreen
        var displayable IDisplayable = this.GetDisplay().GetCurrent(); _ = displayable
        if func() bool { _, ok := any(displayable).(ICanvas); return ok }() {
            var canvas ICanvas = any(displayable).(ICanvas); _ = canvas
            canvas.SetFullScreenMode(isFullScreen)
        }
}


func (this *GameMidlet) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


// throws Exception
func (this *GameMidlet) GetInputMappingInstructionsCanvas() IMyCanvas {
        return NewGameInputMappingInstructionsCanvas(this, this.CreateGameLayerManager())
}


// throws Exception
func (this *GameMidlet) GetInputMappingCanvas() IMyCanvas {
        return NewGameInputMappingCanvas(this.AbeClientInformation, this, this.CreateGameLayerManager(), this.GetHelpPaintable())
}


// throws Exception
func (this *GameMidlet) GetAboutCanvas() IMyCanvas {
        return NewBasicPaintablesCanvas(this, CanvasStrings_GetInstance().ABOUT, this.CreateGameLayerManager(), AboutPaintableFactory_GetInstance().PaintableArray)
}


// throws Exception
func (this *GameMidlet) GetHelpPaintable() IHelpPaintable {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *GameMidlet) CreateHighScoresCanvas() *HighScoresCanvas {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *GameMidlet) StartGameCanvasRunnableInterface() {
        var threadFactoryUtil *ThreadFactoryUtil = ThreadFactoryUtil_GetInstance(); _ = threadFactoryUtil
        this.Thread = threadFactoryUtil.GetInstanceGameCanvasRunnable(this.AllbinaryGameCanvasRunnableInterface)
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Thread Priority: ").Appendint(int(this.Thread.GetPriority()))), this, "startGameCanvasRunnableInterface")
        this.AllbinaryGameCanvasRunnableInterface.SetThread(this.Thread)
        threadFactoryUtil.Start(this.Thread)
}


// throws Exception
func (this *GameMidlet) StopGameCanvasRunnableInterface() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.GameStrings.STOP_GAME_CANVAS_RUNNABLE_INTERFACE)
        GameNotificationEventHandler_GetInstance().RemoveAllListeners()
        ColorChangeEventHandler_GetInstance().RemoveAllListeners()
        GameEventHandlerUtil_RemoveAllListeners()
        var gameCanvasRunnableInterface GameCanvasRunnableInterface = this.AllbinaryGameCanvasRunnableInterface; _ = gameCanvasRunnableInterface
        if gameCanvasRunnableInterface != NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Set Running False: ").Append(fmt.Sprint(gameCanvasRunnableInterface))), this, this.GameStrings.STOP_GAME_CANVAS_RUNNABLE_INTERFACE)
            gameCanvasRunnableInterface.SetRunning(false)
        } else {
            this.LogUtil.PutF("StopGame - Could Not Stop", this, this.GameStrings.STOP_GAME_CANVAS_RUNNABLE_INTERFACE)
        }
        ThreadUtil_GetInstance().Join(this.Thread)
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        if this.Features.IsFeature(MainFeatureFactory_GetInstance().LOAD_ALL) {
            progressCanvas.AddNormalPortion(50, "Stopped Game Runnable")
        } else {
            progressCanvas.AddNormalPortion(50, "Stopped Main Runnable")
        }
        this.LogUtil.PutF(this.CommonStrings.END, this, this.GameStrings.STOP_GAME_CANVAS_RUNNABLE_INTERFACE)
}


func (this *GameMidlet) GetGameCanvasRunnableInterface() GameCanvasRunnableInterface {
        return this.AllbinaryGameCanvasRunnableInterface
}


func (this *GameMidlet) SetGameCanvasRunnableInterface(gameCanvasRunnableInterface GameCanvasRunnableInterface) {
        this.AllbinaryGameCanvasRunnableInterface = gameCanvasRunnableInterface
}


func (this *GameMidlet) CreateGameLayerManager() IAllBinaryGameLayerManager {
        var gameInfo *GameInfo = NewGameInfo(GameTypeFactory_GetInstance().SINGLE_PLAYER, GameMode_SERVER, PlayerTypesFactory_GetInstance().PLAYER_TYPE_ONE, 1, 1); _ = gameInfo
        return NewAllBinaryGameLayerManager(this.BasicColorFactory.BLACK, this.BasicColorFactory.WHITE, gameInfo)
}


func (this *GameMidlet) GetGameOptionsForm() ICommandForm {
        var layerManager IAllBinaryGameLayerManager = this.CreateGameLayerManager(); _ = layerManager
        return GameOptionsFormFactory_GetInstance().Init(this, "Game Options", layerManager.GetBackgroundBasicColor(), layerManager.GetForegroundBasicColor())
}


// throws Exception
func (this *GameMidlet) Save() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.SAVE)
        var hashtable *ABHashtable[any, any] = this.GetCurrentStateHashtable(); _ = hashtable
        GamePersistanceSingleton_GetInstance().Save(this.AbeClientInformation, hashtable)
}


// throws Exception
func (this *GameMidlet) GetCurrentStateHashtable() *ABHashtable[any, any] {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getCurrentStateHashtable")
        var hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = hashtable
        if this.AllbinaryGameCanvasRunnableInterface != NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE {
            var currentHashtable *ABHashtable[any, any] = this.AllbinaryGameCanvasRunnableInterface.GetCurrentStateHashtable(); _ = currentHashtable
            var enumeration Enumeration = currentHashtable.Keys(); _ = enumeration
            var key any; _ = key
            for this.EnumerationUtil.HasMoreElements(enumeration) {
                key = this.EnumerationUtil.NextElement(enumeration)
                hashtable.Put(key, currentHashtable.Get(key))
            }
        }
        return hashtable
}


func (this *GameMidlet) SetLoadGameForm(loadGameForm *LoadGameForm) {
        this.LoadGameForm = loadGameForm
}


func (this *GameMidlet) GetLoadGameForm() ICommandForm {
        return this.LoadGameForm
}


func (this *GameMidlet) SetResized(resized bool) {
        this.Resized = resized
}


func (this *GameMidlet) IsResized() bool {
        return this.Resized
}


func (this *GameMidlet) IsDemoLoading() bool {
        if this.StartedBefore {
            return false
        }
        var displayable IDisplayable = this.GetDisplay().GetCurrent(); _ = displayable
        if func() bool { _, ok := any(displayable).(*DemoCanvas); return ok }() {
            var demoCanvas *DemoCanvas = any(displayable).(*DemoCanvas); _ = demoCanvas
            if demoCanvas.IsDemoLoading() {
                return true
            }
        }
        return false
}

