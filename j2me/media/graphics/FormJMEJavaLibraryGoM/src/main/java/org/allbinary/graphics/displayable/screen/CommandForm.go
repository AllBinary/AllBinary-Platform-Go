/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
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

type CommandForm struct {
    Form
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    RepaintProcessor IProcessor
    CommandStack *ABStack[any]
}

// Implements MyCommandInterface, MenuListener

// The Java class CommandForm has subclasses so its variables are this interface
type ICommandForm interface {
    IForm
    CommandForm_() *CommandForm
    InitCommands(cmdListener CommandListener)
    Open()
    Close()
    Update()
    GetSourceId() int
    AddCommand(command ICommand)
    RemoveAllCommands()
}

func (this *CommandForm) CommandForm_() *CommandForm {
    return this
}

var CommandForm_NULL_COMMAND_FORM any = NullUtil_GetInstance().NULL_OBJECT

func (this *CommandForm) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.RepaintProcessor = ScreenRepaintProcessorFactory_GetInstance().Create(this)
    
}


func CommandForm_GetNullCommandForm() ICommandForm {
        if CommandForm_NULL_COMMAND_FORM == NullUtil_GetInstance().NULL_OBJECT {
            CommandForm_NULL_COMMAND_FORM = NewCommandForm(NullCommandListener_NULL_COMMAND_LISTENER, StringUtil_GetInstance().EMPTY_STRING, BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE)
        }
        return any(CommandForm_NULL_COMMAND_FORM).(ICommandForm)
}


func NewCommandForm(commandListener CommandListener, formTitle string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *CommandForm {
    this := &CommandForm{}
        this.Form = *NewForm(formTitle)
        this.initFields()
        this.CommandStack = StdUtil_GetInstance().CreateStack()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
                }
            }()
            this.RepaintProcessor.Process()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *CommandForm) InitCommands(cmdListener CommandListener) {
}


func (this *CommandForm) Open() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "open")
}


// throws Exception
func (this *CommandForm) Close() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CLOSE)
}


// throws Exception
func (this *CommandForm) Update() {
        this.RepaintProcessor.Process()
}


func (this *CommandForm) GetSourceId() int {
        return 0
}


func (this *CommandForm) AddCommand(command ICommand) {
        this.CommandStack.Push(command)
        this.Form.AddCommand(command)
}


func (this *CommandForm) RemoveAllCommands() {
        var size int = this.CommandStack.Size(); _ = size
        for index := int(0); index < size; index++ {
            this.Form.RemoveCommand(any(this.CommandStack.Pop()).(ICommand))
        }
}

