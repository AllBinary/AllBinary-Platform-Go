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

type PaymentGatewayInterface interface {
    TableMappingInterface
    GetEnable() string
    GetStoreName() string
    GetName() string
    GetPaymentMethod() string
    GetMode() string
    GetTestProtocol() string
    GetTestServer() string
    GetTestPort() string
    GetTestPath() string
    GetServerProtocol() string
    GetServer() string
    GetServerPort() string
    GetServerPath() string
    GetUserName() string
    GetPassword() string
    GetProxyProtocol() string
    GetProxyServer() string
    GetProxyPath() string
    GetProxyPort() string
    GetProxyUserName() string
    GetProxyPassword() string
    GetTimeout() string
    GetProxyTimeout() string
    GetSpecial1() string
    GetSpecial2() string
    GetSpecial3() string
    GetSpecial4() string
    GetSpecial5() string
    GetSpecial6() string
    GetSpecial7() string
    GetSpecial8() string
    GetSpecial9() string
    GetLastModified() string
    GetTimeEntered() string
    SetEnable(value string)
    SetStoreName(value string)
    SetName(value string)
    SetPaymentMethod(value string)
    SetMode(value string)
    SetTestProtocol(value string)
    SetTestServer(value string)
    SetTestPort(value string)
    SetTestPath(value string)
    SetServerProtocol(value string)
    SetServer(value string)
    SetServerPort(value string)
    SetServerPath(value string)
    SetUserName(value string)
    SetPassword(value string)
    SetProxyProtocol(value string)
    SetProxyServer(value string)
    SetProxyPath(value string)
    SetProxyPort(value string)
    SetProxyUserName(value string)
    SetProxyPassword(value string)
    SetTimeout(value string)
    SetProxyTimeout(value string)
    SetSpecial1(value string)
    SetSpecial2(value string)
    SetSpecial3(value string)
    SetSpecial4(value string)
    SetSpecial5(value string)
    SetSpecial6(value string)
    SetSpecial7(value string)
    SetSpecial8(value string)
    SetSpecial9(value string)
    SetLastModified(value string)
    SetTimeEntered(value string)
    ToHashMap(isEncrypted bool) *HashMap
    ToVector(isEncrypted bool) IBasicArrayList
}


