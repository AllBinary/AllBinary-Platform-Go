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

type InGameOptionsFormFactory struct {
    LogUtil *LogUtil
    SINGLETON ICommandForm
}

var InGameOptionsFormFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *InGameOptionsFormFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.SINGLETON = CommandForm_GetNullCommandForm()
}

func NewInGameOptionsFormFactory() *InGameOptionsFormFactory {
    this := &InGameOptionsFormFactory{}
    this.initFields()
    return this
}


func InGameOptionsFormFactory_GetInstance() *InGameOptionsFormFactory {
        if InGameOptionsFormFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            InGameOptionsFormFactory_Instance = NewInGameOptionsFormFactory()
        }
        return any(InGameOptionsFormFactory_Instance).(*InGameOptionsFormFactory)
}


func (this *InGameOptionsFormFactory) Init(commandListener CommandListener, initInterface IInit, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.INIT, e)
                }
            }()
            if this.SINGLETON == CommandForm_GetNullCommandForm() {
                initInterface.Init()
                this.SINGLETON = NewInGameOptionsForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InGameOptionsFormFactory) Get() ICommandForm {
        return this.SINGLETON
}

