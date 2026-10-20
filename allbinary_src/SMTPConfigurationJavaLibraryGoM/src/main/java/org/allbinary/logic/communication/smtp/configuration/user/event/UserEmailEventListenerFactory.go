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

type UserEmailEventListenerFactory struct {
}



func NewUserEmailEventListenerFactory() *UserEmailEventListenerFactory {
    this := &UserEmailEventListenerFactory{}
    return this
}


// throws Exception
func UserEmailEventListenerFactory_GetInstance(abeClientInformation AbeClientInformationInterface, userEmailEventConfigurationInterface UserEmailEventConfigurationInterface, userInterface UserInterface) UserEmailEventListenerInterface {
        var params []any = make([]any, 1); _ = params
        var classes []reflect.Type = make([]reflect.Type, 1); _ = classes
        classes[0] = reflect.TypeOf((*UserInterface)(nil)).Elem()
        params[0] = any(userInterface).(any)
        return any(AbeFactory_GetInstance().GetInstance(abeClientInformation, userEmailEventConfigurationInterface.GetEventListenerClassPath(), classes, params)).(UserEmailEventListenerInterface)
}

