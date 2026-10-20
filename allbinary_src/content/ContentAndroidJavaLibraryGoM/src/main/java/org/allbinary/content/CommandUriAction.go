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
    Hashtable *ABHashtable[any, any]
}

var CommandUriAction_Instance *CommandUriAction = NewCommandUriAction()

func (this *CommandUriAction) initFields() {
    this.LogUtil = LogUtil_GetInstance()
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
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.PROCESS, e)
                }
            }()
            var intent *Intent = this.GetIntent(command); _ = intent
            ResourceUtil_GetInstance().GetContext().StartActivity(intent)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *CommandUriAction) GetIntent(command ICommand) *Intent {
        var url string = any(this.Hashtable.Get(command)).(string); _ = url
        var uri *Uri = Uri_Parse(url); _ = uri
        var intent *Intent = NewIntent(Intent_ACTION_VIEW, uri); _ = intent
        return intent
}

