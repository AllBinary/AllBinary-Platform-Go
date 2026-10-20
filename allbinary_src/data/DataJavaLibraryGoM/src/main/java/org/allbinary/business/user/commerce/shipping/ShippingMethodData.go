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

type ShippingMethodData struct {
}

var ShippingMethodData_DESCRIPTION string = "SHIPPINGMETHOD_DESCRIPTION"
var ShippingMethodData_DEFAULT string = "SHIPPINGMETHOD_DEFAULT"
var ShippingMethodData_NAME string = "SHIPPINGMETHOD_NAME"
var ShippingMethodData_GROUP string = "SHIPPINGMETHOD_GROUP"
var ShippingMethodData_COST string = "SHIPPINGMETHOD_COST"


func NewShippingMethodData() *ShippingMethodData {
    this := &ShippingMethodData{}
    return this
}

