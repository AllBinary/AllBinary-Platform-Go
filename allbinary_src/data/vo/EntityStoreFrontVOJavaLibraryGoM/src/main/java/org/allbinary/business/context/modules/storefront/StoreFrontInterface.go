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

type StoreFrontInterface interface {
    TableMappingInterface
    IsNameValid() bool
    IsValid() bool
    NameValidationInfo() string
    ValidationInfo() string
    GetTestHtmlPath() string
    GetCurrentHostName() string
    GetCurrentHostNamePath() string
    GetCurrentHomeHostName() string
    GetCurrentHomeHostNamePath() string
    GetName() string
    GetBasketName() string
    GetHomeHostName() string
    GetHomeHostNamePath() string
    GetHostName() string
    GetHostNamePath() string
    GetTestHomeHostName() string
    GetTestHomeHostNamePath() string
    GetTestHostName() string
    GetTestHostNamePath() string
    GetStaticPath() string
    GetCategoryPath() string
    GetSubStores() *BasicArrayList
    GetTagLocation() string
    GetPackageLocation() string
    GetInventoryControl() string
    GetContextConfigurationInterface() ContextConfigurationInterface
    SetContextConfigurationInterface(contextConfigurationInterface ContextConfigurationInterface)
    GetFtp() string
    GetFtpPath() string
    GetFtpUserName() string
    GetFtpPassword() string
    GetTestFtp() string
    GetTestFtpPath() string
    GetTestFtpUserName() string
    GetTestFtpPassword() string
    GetTimeCreated() string
    GetLastModified() string
    SetName(value string)
    SetBasketName(value string)
    SetHomeHostName(value string)
    SetHomeHostNamePath(value string)
    SetHostName(value string)
    SetHostNamePath(value string)
    SetTestHomeHostName(value string)
    SetTestHomeHostNamePath(value string)
    SetTestHostName(value string)
    SetTestHostNamePath(value string)
    SetImagePath(value string)
    SetStaticPath(value string)
    SetCategoryPath(value string)
    SetPackageLocation(value string)
    SetInventoryControl(value string)
    SetFtp(value string)
    SetFtpPath(value string)
    SetFtpUserName(value string)
    SetFtpPassword(value string)
    SetTestFtp(value string)
    SetTestFtpPath(value string)
    SetTestFtpUserName(value string)
    SetTestFtpPassword(value string)
    SetTimeCreated(value string)
    SetLastModified(value string)
    Install(current int, total int)
}


