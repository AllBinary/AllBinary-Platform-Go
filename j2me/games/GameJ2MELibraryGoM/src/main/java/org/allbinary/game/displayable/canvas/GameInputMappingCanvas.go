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

type GameInputMappingCanvas struct {
    GameCommandCanvas
    Paintable IProcessPaintable
    HelpPaintable *InputMappingHelpPaintable
    ColorFillPaintable IColorFillBasePaintable
    InputMapping IPersistentInputMapping
    NONE *GameKey
    AbeClientInformation AbeClientInformationInterface
    SelectedGameKey *GameKey
    SelectedInput IInput
    InputFactory *InputFactory
}

// Implements InputMappingInterface

var GameInputMappingCanvas_NAME string = "GameInputMappingCanvas"
var GameInputMappingCanvas_DISPLAY ICommand = NewCommand("Controls", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 2)
var GameInputMappingCanvas_DEFAULT ICommand = NewCommand("Default", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 2)

func (this *GameInputMappingCanvas) initFields() {
    
    
    
    
    this.NONE = GameKeyFactory_GetInstance().NONE
    
    this.SelectedGameKey = this.NONE
    this.SelectedInput = this.NONE
    this.InputFactory = InputFactory_GetInstance()
}


func NewGameInputMappingCanvas(abeClientInformation AbeClientInformationInterface, commandListener CommandListener, allBinaryGameLayerManager IAllBinaryGameLayerManager, helpPaintable IHelpPaintable) *GameInputMappingCanvas {
    this := &GameInputMappingCanvas{}
        this.GameCommandCanvas = *NewGameCommandCanvas(commandListener, GameInputMappingCanvas_NAME, allBinaryGameLayerManager.GetBackgroundBasicColor(), allBinaryGameLayerManager.GetForegroundBasicColor())
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        if helpPaintable == nil {
            panic(errors.New(fmt.Sprint("Help Paintable Exception")))
        }
        this.AbeClientInformation = abeClientInformation
        this.HelpPaintable = any(helpPaintable).(*InputMappingHelpPaintable)
        this.InputMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance()
        this.Paintable = NewProcessPaintable()
        this.ColorFillPaintable = ColorFillPaintableFactory_GetInstance().GetInstance(allBinaryGameLayerManager.GetBackgroundBasicColor(), false)
    return this
}


// throws Exception
func (this *GameInputMappingCanvas) Close() {
        this.GameCommandCanvas.Close()
        this.Paintable.Process()
        this.SelectedGameKey = this.NONE
        this.SelectedInput = this.NONE
        this.Update()
}


func (this *GameInputMappingCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.AddCommand(GameCommandsFactory_GetInstance().CLOSE_AND_SHOW_GAME_CANVAS)
        this.AddCommand(GameInputMappingCanvas_DEFAULT)
        this.AddCommand(GameInputMappingInstructionsCanvas_DISPLAY)
        this.SetCommandListener(cmdListener)
}


func (this *GameInputMappingCanvas) KeyPressed(keyCode int) {
        this.KeyPressedByDevice(keyCode, 0)
}


func (this *GameInputMappingCanvas) KeyReleased(keyCode int) {
        this.KeyReleasedByDevice(keyCode, 0)
}


func (this *GameInputMappingCanvas) KeyRepeated(keyCode int) {
        this.KeyRepeatedByDevice(keyCode, 0)
}


func (this *GameInputMappingCanvas) KeyPressedByDevice(keyCode int, deviceId int) {
        this.AddGameKey(keyCode, false)
        this.GameCommandCanvas.KeyPressedByDevice(keyCode, 0)
}


func (this *GameInputMappingCanvas) AddGameKey(keyCode int, repeated bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Key Event Error", this, this.GameInputStrings.ADD_KEY_EVENT, e)
                }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Raw Device Key Code: ").Append(strconv.FormatInt(int64(keyCode), 16))), this, this.GameInputStrings.ADD_KEY_EVENT)
            var gameKey *GameKey = this.InputToGameKeyMapping.GetInstanceForCanvas(this, keyCode); _ = gameKey
            var input IInput = this.InputFactory.GetInstanceById(keyCode); _ = input
            this.ProcessInputMapping(gameKey, input)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *GameInputMappingCanvas) ProcessInputMapping(gameKey *GameKey, input IInput) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Start Passed GameKey: ")
        stringBuffer.Append(this.StringUtil.ToString(gameKey))
        stringBuffer.Append(" Input: ")
        stringBuffer.Append(this.StringUtil.ToString(input))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.PROCESS)
        if this.SelectedGameKey != this.NONE {
            this.GameActionCrud(gameKey, input)
        } else {
            this.SetSelectedAction(gameKey)
        }
}


func (this *GameInputMappingCanvas) SetSelectedAction(gameKey *GameKey) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Selected GameKey: ").Append(this.StringUtil.ToString(gameKey))), this, "setSelectedAction")
        this.SelectedGameKey = gameKey
        this.SelectedInput = this.NONE
        this.HelpPaintable.Update(this.SelectedGameKey, this.SelectedInput)
        this.RepaintBehavior.OnChangeRepaint(this)
}


// throws Exception
func (this *GameInputMappingCanvas) GameActionCrud(gameKey *GameKey, input IInput) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Start GameKey: ")
        stringBuffer.Append(this.StringUtil.ToString(this.SelectedGameKey))
        stringBuffer.Append(" Input: ")
        stringBuffer.Append(this.StringUtil.ToString(this.SelectedInput))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "gameActionCrud")
        if this.SelectedInput == this.NONE {
            var list IBasicArrayList = this.InputMapping.GetInputMapping().GetMappedInput(this.SelectedGameKey); _ = list
            var isInputAlreadyMappedToSelectedAction bool = list.Contains(input); _ = isInputAlreadyMappedToSelectedAction
            if isInputAlreadyMappedToSelectedAction {
                this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Already Mapped Input: ").Append(this.StringUtil.ToString(input))), this, "gameActionCrud")
                this.SelectedInput = input
                this.HelpPaintable.Update(this.SelectedGameKey, this.SelectedInput)
                this.RepaintBehavior.OnChangeRepaint(this)
            } else {
                this.AddNewMapping(gameKey, input)
            }
        } else if this.InputMapping.IsDelete(input) {
            this.DeleteCurrentMapping()
        } else {
            this.SetSelectedAction(gameKey)
        }
}


// throws Exception
func (this *GameInputMappingCanvas) AddNewMapping(gameKey *GameKey, input IInput) {
        var METHOD_NAME string = "addNewMapping"; _ = METHOD_NAME
        this.LogUtil.PutF(this.CommonStrings.START, this, METHOD_NAME)
        var isInputAlreadyMapped bool = this.InputMapping.GetInputMapping().IsMapped(input); _ = isInputAlreadyMapped
        if !isInputAlreadyMapped && !this.InputMapping.IsSystemInput(input) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Add Key Mapping : GameKey: ")
            stringBuffer.Append(this.StringUtil.ToString(this.SelectedGameKey))
            stringBuffer.Append(" Input: ")
            stringBuffer.Append(this.StringUtil.ToString(this.SelectedInput))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
            this.InputMapping.GetInputMapping().Add(this.SelectedGameKey, input)
            this.SelectedInput = input
            this.Update()
        } else {
            this.LogUtil.PutF("Unable to add Mapping since one already exists or is MENU, HOME, or BACK key and setting selected action to what it is already mapped to", this, METHOD_NAME)
            this.SetSelectedAction(gameKey)
        }
}


// throws Exception
func (this *GameInputMappingCanvas) DeleteCurrentMapping() {
        var METHOD_NAME string = "deleteCurrentMapping"; _ = METHOD_NAME
        var list IBasicArrayList = this.InputMapping.GetInputMapping().GetMappedInput(this.SelectedGameKey); _ = list
        if list.Size() > 1 {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Start GameKey: ")
            stringBuffer.Append("Remove Key Mapping: GameKey: ")
            stringBuffer.Append(this.StringUtil.ToString(this.SelectedGameKey))
            stringBuffer.Append(" Input: ")
            stringBuffer.Append(this.StringUtil.ToString(this.SelectedInput))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
            this.InputMapping.GetInputMapping().Remove(this.SelectedGameKey, this.SelectedInput)
            this.SelectedInput = this.NONE
            this.Update()
        } else {
            this.LogUtil.PutF("Can't Remove Last Key Mapping", this, METHOD_NAME)
        }
}


// throws Exception
func (this *GameInputMappingCanvas) SetDefault() {
        this.InputMapping.SetDefault(this.AbeClientInformation)
        this.HelpPaintable.Update(this.NONE, this.NONE)
        this.RepaintBehavior.OnChangeRepaint(this)
}


// throws Exception
func (this *GameInputMappingCanvas) Update() {
        this.InputMapping.Update(this.AbeClientInformation)
        this.HelpPaintable.Update(this.SelectedGameKey, this.SelectedInput)
        this.RepaintBehavior.OnChangeRepaint(this)
}


func (this *GameInputMappingCanvas) Paint(graphics IGraphics) {
        this.ColorFillPaintable.Paint(graphics)
        this.HelpPaintable.Paint(graphics)
        this.Paintable.Paint(graphics)
        this.GameCommandCanvas.Paint(graphics)
}

