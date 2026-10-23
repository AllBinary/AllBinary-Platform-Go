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

type InitDbHelperFactory struct {
}

var InitDbHelperFactory_CLASSNAME string = "org.allbinary.business.installer.InitDbCrypted"


func NewInitDbHelperFactory() *InitDbHelperFactory {
    this := &InitDbHelperFactory{}
    return this
}


// throws LicensingException
func InitDbHelperFactory_GetInstance_1(abeClientInformation AbeClientInformationInterface, databaseConnectionInfoInterface DatabaseConnectionInfoInterface) any {
        var params []any = make([]any, 1); _ = params
        var classes []reflect.Type = make([]reflect.Type, 1); _ = classes
        classes[0] = reflect.TypeOf((*DatabaseConnectionInfoInterface)(nil)).Elem()
        params[0] = any(databaseConnectionInfoInterface).(any)
        var anyType any = AbeFactory_GetInstance().GetInstance(abeClientInformation, InitDbHelperFactory_CLASSNAME, classes, params); _ = anyType
        return anyType
}

// Java overloads of InitDbHelperFactory.getInstance
func InitDbHelperFactory_GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[DatabaseConnectionInfoInterface](args[1], pass) {
            return InitDbHelperFactory_GetInstance_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[DatabaseConnectionInfoInterface](args[1]))
        }
        if len(args) == 1 && Java_IsArg[AbeClientInformationInterface](args[0], pass) {
            return InitDbHelperFactory_GetInstance_2(Java_Arg[AbeClientInformationInterface](args[0]))
        }
    }
    panic("No Java overload of InitDbHelperFactory.getInstance for the arguments")
}


// throws LicensingException
func InitDbHelperFactory_GetInstance_2(abeClientInformation AbeClientInformationInterface) any {
        var anyType any = AbeFactory_GetInstance().GetInstance(abeClientInformation, InitDbHelperFactory_CLASSNAME, nil, nil); _ = anyType
        return anyType
}

