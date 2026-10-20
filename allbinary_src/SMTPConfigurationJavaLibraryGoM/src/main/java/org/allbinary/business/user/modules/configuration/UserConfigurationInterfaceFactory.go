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

type UserConfigurationInterfaceFactory struct {
}

var UserConfigurationInterfaceFactory_ConfigurationName string = "DefaultUserConfiguration"


func NewUserConfigurationInterfaceFactory() *UserConfigurationInterfaceFactory {
    this := &UserConfigurationInterfaceFactory{}
    return this
}


// throws Exception
func UserConfigurationInterfaceFactory_GetInstance(document *Document) UserConfigurationInterface {
        var configurationNode *Node = DomSearchHelper_GetNode(UserConfigurationData_NAME, document.GetChildNodes()); _ = configurationNode
        return any(NewUserConfigurationView(configurationNode).GetUserConfigurationInterface()).(UserConfigurationInterface)
}


// throws Exception
// Java overloaded method: Go needs a different name
func UserConfigurationInterfaceFactory_GetInstance(userRole *UserRole) UserConfigurationInterface {
        var abPath *AbPath = NewAbPath(URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().USERCONFIGURATIONPATH, fmt.Sprint(UserConfigurationInterfaceFactory_ConfigurationName) + fmt.Sprint(userRole) + fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + fmt.Sprint(UserConfigurationData_UNCRYPTED_EXTENSION)); _ = abPath
        var documentString string = NewCryptFileReader(UserConfigurationData_UNCRYPTED_EXTENSION, UserConfigurationData_ENCRYPTED_EXTENSION).Get(abPath); _ = documentString
        var document *Document = DomDocumentHelper_Create(documentString); _ = document
        return UserConfigurationInterfaceFactory_GetInstance(document)
}


// throws Exception
// Java overloaded method: Go needs a different name
func UserConfigurationInterfaceFactory_GetInstance() UserConfigurationInterface {
        return NewUserConfiguration()
}

