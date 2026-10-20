/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2022 AllBinary
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

type WebCommandProcessor struct {
    LogUtil *LogUtil
    AboutCommandProcessor *AboutCommandProcessor
    List IBasicArrayList
}

var WebCommandProcessor_Instance *WebCommandProcessor = NewWebCommandProcessor()

func (this *WebCommandProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AboutCommandProcessor = AboutCommandProcessor_GetInstance()
    this.List = NewBasicArrayListD()
}

func NewWebCommandProcessor() *WebCommandProcessor {
    this := &WebCommandProcessor{}
    this.initFields()
    return this
}


func WebCommandProcessor_GetInstance() *WebCommandProcessor {
        return WebCommandProcessor_Instance
}


func (this *WebCommandProcessor) Process(midletCommandListener CommandListener, command ICommand, canvas ICanvas) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.PROCESS, e)
                }
            }()
            PlayN_OpenURL(any(this.List.RemoveAt(0)).(string))
            return
        }(); tryControl1 == 1 {
            return
        }
}

