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

type GenericConfigData struct {
}

var GenericConfigData_RELATIVE_PATH string = "./"
var GenericConfigData_CONFIG_PATH string = "modules/configs/"
var GenericConfigData_MODULE_RELATIVE_PATH string = fmt.Sprint(GenericConfigData_RELATIVE_PATH + GenericConfigData_CONFIG_PATH) + "allbinary/input/automation/module/"
var GenericConfigData_GAME_MODULE_RELATIVE_PATH string = fmt.Sprint(GenericConfigData_MODULE_RELATIVE_PATH) + "game/"
var GenericConfigData_MMO_GAME_MODULE_RELATIVE_PATH string = fmt.Sprint(GenericConfigData_GAME_MODULE_RELATIVE_PATH) + "mmog/"
var GenericConfigData_FREE_MMO_GAME_MODULE_RELATIVE_PATH string = "free/"


func NewGenericConfigData() *GenericConfigData {
    this := &GenericConfigData{}
    return this
}

