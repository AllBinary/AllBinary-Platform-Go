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

type WeblisketSessionData struct {
}

var WeblisketSessionData_INVALIDATESESSION string = "SESSION_INVALIDATE_SESSION"
var WeblisketSessionData_ATTEMPTS string = "SESSION_ATTEMPTS"
var WeblisketSessionData_TIMEOUT string = "SESSION_TIMEOUT"
var WeblisketSessionData_REMOVABLEUSERNAME string = "SESSION_REMOVABLE_USER_NAME"
var WeblisketSessionData_REMOVABLEPASSWORD string = "SESSION_REMOVABLE_PASSWORD"
var WeblisketSessionData_REMOVABLENEWPASSWORD string = "SESSION_REMOVABLE_NEW_PASSWORD"
var WeblisketSessionData_REMOVABLEREENTERNEWPASSWORD string = "SESSION_REMOVABLE_REENTER_NEW_PASSWORD"
var WeblisketSessionData_AUTHENTICATED string = "SESSION_AUTHENTICATED"


func NewWeblisketSessionData() *WeblisketSessionData {
    this := &WeblisketSessionData{}
    return this
}

