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
func UserConfigurationInterfaceFactory_GetInstance_1(document *Document) UserConfigurationInterface {
        var configurationNode *Node = DomSearchHelper_GetNode(UserConfigurationData_NAME, document.GetChildNodes()); _ = configurationNode
        return any(NewUserConfigurationView(configurationNode).GetUserConfigurationInterface()).(UserConfigurationInterface)
}

// Java overloads of UserConfigurationInterfaceFactory.getInstance
func UserConfigurationInterfaceFactory_GetInstance(args ...any) UserConfigurationInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*Document](args[0], pass) {
            return UserConfigurationInterfaceFactory_GetInstance_1(Java_Arg[*Document](args[0]))
        }
        if len(args) == 1 && Java_IsArg[IUserRole](args[0], pass) {
            return UserConfigurationInterfaceFactory_GetInstance_2(Java_Arg[IUserRole](args[0]))
        }
        if len(args) == 0 {
            return UserConfigurationInterfaceFactory_GetInstance_3()
        }
    }
    panic("No Java overload of UserConfigurationInterfaceFactory.getInstance for the arguments")
}


// throws Exception
func UserConfigurationInterfaceFactory_GetInstance_2(userRole IUserRole) UserConfigurationInterface {
        var abPath IAbPath = NewAbPath(URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().USERCONFIGURATIONPATH, fmt.Sprint(UserConfigurationInterfaceFactory_ConfigurationName) + fmt.Sprint(userRole) + fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + fmt.Sprint(UserConfigurationData_UNCRYPTED_EXTENSION)); _ = abPath
        var documentString string = NewCryptFileReader(UserConfigurationData_UNCRYPTED_EXTENSION, UserConfigurationData_ENCRYPTED_EXTENSION).Get(abPath); _ = documentString
        var document *Document = DomDocumentHelper_Create(documentString); _ = document
        return UserConfigurationInterfaceFactory_GetInstance(document)
}


// throws Exception
func UserConfigurationInterfaceFactory_GetInstance_3() UserConfigurationInterface {
        return NewUserConfiguration()
}

