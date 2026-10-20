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

type MouseActionScriptInputData struct {
}

var MouseActionScriptInputData_NAME string = "MOUSE_ACTION_INPUT_NAME"
var MouseActionScriptInputData_BUTTONS string = "MOUSE_ACTION_INPUT_BUTTONS"
var MouseActionScriptInputData_MOVE string = "MOUSE_ACTION_INPUT_MOVE"
var MouseActionScriptInputData_MOVE_X string = "MOUSE_ACTION_INPUT_MOVE_X"
var MouseActionScriptInputData_MOVE_Y string = "MOUSE_ACTION_INPUT_MOVE_Y"


func NewMouseActionScriptInputData() *MouseActionScriptInputData {
    this := &MouseActionScriptInputData{}
    return this
}

