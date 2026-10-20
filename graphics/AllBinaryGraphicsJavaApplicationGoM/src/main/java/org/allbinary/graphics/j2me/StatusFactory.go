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

type StatusFactory struct {
    LogUtil *LogUtil
    SETSTATUS string
    STATUS_LABEL string
}

var StatusFactory_Instance *StatusFactory = NewStatusFactory()

func (this *StatusFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.SETSTATUS = "setStatus"
    this.STATUS_LABEL = "Status: "
}

func NewStatusFactory() *StatusFactory {
    this := &StatusFactory{}
    this.initFields()
    return this
}


func StatusFactory_GetInstance() *StatusFactory {
        return StatusFactory_Instance
}


func (this *StatusFactory) SetStatusNoLog(status string) {
        MyFrame_StatusJLabel.SetText(status)
}


func (this *StatusFactory) SetStatus(status string) {
        this.SetStatusNoLog(status)
        this.LogUtil.PutF(this.STATUS_LABEL + status, this, this.SETSTATUS)
}

