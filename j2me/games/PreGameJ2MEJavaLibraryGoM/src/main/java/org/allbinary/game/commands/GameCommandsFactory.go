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

type GameCommandsFactory struct {
    TOGGLE_FULLSCREEN ICommand
    TOGGLE_KEYBOARD ICommand
    SET_MENU_DISPLAYABLE ICommand
    APPLICATION_UPDATE ICommand
    DISPLAY_ABOUT ICommand
    OPEN_WEB_URL ICommand
    CLOSE_OPTIONS ICommand
    DISPLAY_OPTIONS ICommand
    START_TRACE ICommand
    STOP_TRACE ICommand
    DEFAULT_OPTIONS ICommand
    SHOW_GAME_CANVAS ICommand
    CLOSE_AND_SHOW_GAME_CANVAS ICommand
    QUIT_COMMAND ICommand
    RESTART_COMMAND ICommand
    DISPLAY_SAVE_FORM ICommand
    SAVE ICommand
    DISPLAY_LOAD_FORM ICommand
    LOAD_FILE ICommand
    DELETE_FILE ICommand
    EXIT_COMMAND ICommand
    EXIT_WITHOUT_PROGRESS_COMMAND ICommand
    START_COMMAND ICommand
    CONTINUE_COMMAND ICommand
    SELECT_COMMAND ICommand
    FEATURE_UPDATE ICommand
    BUY_COMMAND ICommand
}

var GameCommandsFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameCommandsFactory) initFields() {
    this.TOGGLE_FULLSCREEN = NewCommand("Toggle FullScreen", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.TOGGLE_KEYBOARD = NewCommand("Keyboard", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.SET_MENU_DISPLAYABLE = NewCommand("Set Menu Displayable", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.APPLICATION_UPDATE = NewCommand(CommonStrings_GetInstance().UPDATE, StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.DISPLAY_ABOUT = NewCommand(CanvasStrings_GetInstance().ABOUT, StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 2)
    this.OPEN_WEB_URL = NewCommand("Open Web URL", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 2)
    this.CLOSE_OPTIONS = NewCommand("Back", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.DISPLAY_OPTIONS = NewCommand(CanvasStrings_GetInstance().OPTIONS, StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 2)
    this.START_TRACE = NewCommand("Start Trace", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.STOP_TRACE = NewCommand("Stop Trace", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.DEFAULT_OPTIONS = NewCommand("Default", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.SHOW_GAME_CANVAS = NewCommand("Back to Game", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.CLOSE_AND_SHOW_GAME_CANVAS = NewCommand("Back", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.QUIT_COMMAND = NewCommand("Quit", StringUtil_GetInstance().EMPTY_STRING, Command_STOP, 1)
    this.RESTART_COMMAND = NewCommand("Restart", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 0)
    this.DISPLAY_SAVE_FORM = NewCommand("Save To", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 3)
    this.SAVE = NewCommand("Save", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 3)
    this.DISPLAY_LOAD_FORM = NewCommand("Load", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 3)
    this.LOAD_FILE = NewCommand("Load File", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 3)
    this.DELETE_FILE = NewCommand("Delete File", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.EXIT_COMMAND = NewCommand("Exit", StringUtil_GetInstance().EMPTY_STRING, Command_EXIT, 2)
    this.EXIT_WITHOUT_PROGRESS_COMMAND = NewCommand("Exit Without Progress", StringUtil_GetInstance().EMPTY_STRING, Command_EXIT, 2)
    this.START_COMMAND = NewCommand(CommonStrings_GetInstance().START, StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.CONTINUE_COMMAND = NewCommand("Continue", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.SELECT_COMMAND = NewCommand("Select", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.FEATURE_UPDATE = NewCommand("Feature Update", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
    this.BUY_COMMAND = NewCommand("BUY", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
}


func GameCommandsFactory_GetInstance() *GameCommandsFactory {
        if GameCommandsFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameCommandsFactory_Instance = NewGameCommandsFactory()
        }
        return any(GameCommandsFactory_Instance).(*GameCommandsFactory)
}


func NewGameCommandsFactory() *GameCommandsFactory {
    this := &GameCommandsFactory{}
    this.initFields()
    return this
}

