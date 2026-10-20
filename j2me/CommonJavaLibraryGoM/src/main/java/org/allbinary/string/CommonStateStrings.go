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

type CommonStateStrings struct {
    CONTEXT string
    FINISH string
    CREATE string
    START string
    RESTART string
    RESUME string
    PAUSE string
    STOP string
    DESTROY string
    BIND string
    ON_SAVE_INSTANCE_STATE string
    ON_START_COMMAND string
}

var CommonStateStrings_Instance *CommonStateStrings = NewCommonStateStrings()

func (this *CommonStateStrings) initFields() {
    this.CONTEXT = "Context: "
    this.FINISH = "finish"
    this.CREATE = "onCreate"
    this.START = "onStart"
    this.RESTART = "onRestart"
    this.RESUME = "onResume"
    this.PAUSE = "onPause"
    this.STOP = "onStop"
    this.DESTROY = "onDestroy"
    this.BIND = "onBind"
    this.ON_SAVE_INSTANCE_STATE = "onSaveInstanceState"
    this.ON_START_COMMAND = "onStartCommand"
}

func NewCommonStateStrings() *CommonStateStrings {
    this := &CommonStateStrings{}
    this.initFields()
    return this
}


func CommonStateStrings_GetInstance() *CommonStateStrings {
        return CommonStateStrings_Instance
}
