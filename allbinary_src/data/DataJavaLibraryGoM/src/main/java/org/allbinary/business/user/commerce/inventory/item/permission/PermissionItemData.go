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

type PermissionItemData struct {
}

var PermissionItemData_WHO string = "PERMISSIONSITEM_WHO"
var PermissionItemData_WHAT string = "PERMISSIONSITEM_WHAT"
var PermissionItemData_TYPE string = "PERMISSIONSITEM_TYPE"
var PermissionItemData_START_TIME string = "PERMISSIONSITEM_START_TIME"
var PermissionItemData_END_TIME string = "PERMISSIONSITEM_END_TIME"
var PermissionItemData_REMOTE_ADDR string = "PERMISSIONSITEM_REMOTE_ADDR"

func NewPermissionItemData() *PermissionItemData {
    this := &PermissionItemData{}
    return this
}

