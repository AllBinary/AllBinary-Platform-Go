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

type CommandUriAction struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Hashtable *ABHashtable[any, any]
}

var CommandUriAction_Instance *CommandUriAction = NewCommandUriAction()

func (this *CommandUriAction) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
}

func NewCommandUriAction() *CommandUriAction {
    this := &CommandUriAction{}
    this.initFields()
    return this
}


func CommandUriAction_GetInstance() *CommandUriAction {
        return CommandUriAction_Instance
}


func (this *CommandUriAction) Add(command ICommand, url string) {
        this.Hashtable.Put(command, url)
}


func (this *CommandUriAction) Process(command ICommand) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *CommandUriAction) GetIntent(command ICommand) *Intent {
        var url string = any(this.Hashtable.Get(command)).(string); _ = url
        return this.Intent
}

