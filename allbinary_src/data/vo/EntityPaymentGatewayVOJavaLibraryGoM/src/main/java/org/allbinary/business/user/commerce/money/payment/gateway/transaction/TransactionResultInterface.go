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

type TransactionResultInterface interface {
    ToString() string
    GetValues() IBasicArrayList
    GetResult() string
    GetPnRef() string
    GetRespMsg() string
    GetAvsAddr() string
    GetAvsZip() string
    GetOrigResult() string
    GetStatus() string
    GetFraudCode() string
    GetFraudMsg() string
    GetErrCode() string
    GetScore() string
    GetReason1() string
    GetReason2() string
    GetReason3() string
    GetException1() string
    GetException2() string
    GetException3() string
    GetException4() string
    GetException5() string
    GetException6() string
    GetException7() string
}


