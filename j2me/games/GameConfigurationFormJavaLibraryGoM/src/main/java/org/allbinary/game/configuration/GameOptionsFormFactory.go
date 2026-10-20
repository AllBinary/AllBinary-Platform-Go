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

type GameOptionsFormFactory struct {
    LogUtil *LogUtil
}

var GameOptionsFormFactory_Instance *GameOptionsFormFactory = NewGameOptionsFormFactory()

func (this *GameOptionsFormFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewGameOptionsFormFactory() *GameOptionsFormFactory {
    this := &GameOptionsFormFactory{}
    this.initFields()
    return this
}


func GameOptionsFormFactory_GetInstance() *GameOptionsFormFactory {
        return GameOptionsFormFactory_Instance
}


func (this *GameOptionsFormFactory) Init(commandListener CommandListener, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) ICommandForm {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 ICommandForm) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.INIT, e)
            tryControl1, tryResult1 = 1, CommandForm_GetNullCommandForm(); return
                }
            }()
            tryControl1, tryResult1 = 1, NewGameOptionsForm(commandListener, title, backgrounBasicColor, foregroundBasicColor); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

