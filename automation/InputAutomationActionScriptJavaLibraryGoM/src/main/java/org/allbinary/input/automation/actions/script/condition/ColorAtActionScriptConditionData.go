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

type ColorAtActionScriptConditionData struct {
}

var ColorAtActionScriptConditionData_NAME string = "COLOR_AT_ACTION_CONDITION_NAME"
var ColorAtActionScriptConditionData_LOCATION_X string = "COLOR_AT_ACTION_CONDITION_LOCATION_X"
var ColorAtActionScriptConditionData_LOCATION_Y string = "COLOR_AT_ACTION_CONDITION_LOCATION_Y"
var ColorAtActionScriptConditionData_MIN_RED string = "COLOR_AT_ACTION_CONDITION_MIN_RED"
var ColorAtActionScriptConditionData_MAX_RED string = "COLOR_AT_ACTION_CONDITION_MAX_RED"
var ColorAtActionScriptConditionData_MIN_GREEN string = "COLOR_AT_ACTION_CONDITION_MIN_GREEN"
var ColorAtActionScriptConditionData_MAX_GREEN string = "COLOR_AT_ACTION_CONDITION_MAX_GREEN"
var ColorAtActionScriptConditionData_MIN_BLUE string = "COLOR_AT_ACTION_CONDITION_MIN_BLUE"
var ColorAtActionScriptConditionData_MAX_BLUE string = "COLOR_AT_ACTION_CONDITION_MAX_BLUE"


func NewColorAtActionScriptConditionData() *ColorAtActionScriptConditionData {
    this := &ColorAtActionScriptConditionData{}
    return this
}

