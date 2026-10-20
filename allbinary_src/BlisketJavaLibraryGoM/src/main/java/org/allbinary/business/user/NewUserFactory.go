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

type NewUserFactory struct {
}



func NewNewUserFactory() *NewUserFactory {
    this := &NewUserFactory{}
    return this
}


// throws Exception
func NewUserFactory_GetInstance_1() UserInterface {
        return NewUser()
}

// Java overloads of NewUserFactory.getInstance
func NewUserFactory_GetInstance(args ...any) UserInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewUserFactory_GetInstance_1()
        }
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewUserFactory_GetInstance_2(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[*HttpServletRequest](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewUserFactory_GetInstance_3(Java_Arg[*HttpServletRequest](args[0]), Java_Arg[*HashMap](args[1]))
        }
    }
    panic("No Java overload of NewUserFactory.getInstance for the arguments")
}


// throws Exception
func NewUserFactory_GetInstance_2(transformInfoInterface TransformInfoInterface) UserInterface {
        var httpTransformInfoInterface TransformInfoHttpInterface = any(transformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
        var httpServletRequest *HttpServletRequest = any(httpTransformInfoInterface.GetPageContext().GetRequest()).(*HttpServletRequest); _ = httpServletRequest
        var hashMap *HashMap = NewRequestParams(httpServletRequest).ToHashMap(); _ = hashMap
        var role string = any(httpTransformInfoInterface.GetPropertiesHashMap().Get(UserRoleData_NAME)).(string); _ = role
        hashMap.Put(UserRoleData_NAME, role)
        return CreateUserFactory_GetInstance(hashMap)
}


// throws Exception
func NewUserFactory_GetInstance_3(httpServletRequest *HttpServletRequest, propertiesHashMap *HashMap) UserInterface {
        if propertiesHashMap != nil {
            var requestHashMap *HashMap = NewRequestParams(httpServletRequest).ToHashMap(); _ = requestHashMap
            var roleString string = any(propertiesHashMap.Get(UserRoleData_NAME)).(string); _ = roleString
            requestHashMap.Put(UserRoleData_NAME, roleString)
            return CreateUserFactory_GetInstance(requestHashMap)
        } else {
            panic(errors.New(fmt.Sprint("Null HashMap")))
        }
}

