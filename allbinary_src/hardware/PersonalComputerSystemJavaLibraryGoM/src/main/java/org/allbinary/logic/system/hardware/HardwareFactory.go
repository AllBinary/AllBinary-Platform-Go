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

type HardwareFactory struct {
}

var HardwareFactory_Instance *HardwareFactory = NewHardwareFactory()


func HardwareFactory_GetInstance_1() *HardwareFactory {
        return HardwareFactory_Instance
}

// Java overloads of HardwareFactory.getInstance
func HardwareFactory_GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return HardwareFactory_GetInstance_1()
        }
        if len(args) == 1 && Java_IsArg[IGenericOperatingSystem](args[0], pass) {
            return HardwareFactory_GetInstance_2(Java_Arg[IGenericOperatingSystem](args[0]))
        }
    }
    panic("No Java overload of HardwareFactory.getInstance for the arguments")
}


func NewHardwareFactory() *HardwareFactory {
    this := &HardwareFactory{}
    return this
}


// throws Exception
func HardwareFactory_GetInstance_2(os IGenericOperatingSystem) HardwareInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 HardwareInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var operatingSystems *OperatingSystems = OperatingSystems_GetInstance(); _ = operatingSystems
            if strings.Compare(os.GetName(), operatingSystems.LINUX) == 0 {
                tryControl1, tryResult1 = 1, any(NewLinuxHardware()).(HardwareInterface); return
            } else if strings.Index(os.GetName(), string(operatingSystems.WINDOWS)) >= 0 {
                tryControl1, tryResult1 = 1, any(NewWindowsHardware()).(HardwareInterface); return
            }
            panic(errors.New(fmt.Sprint("No Hardware Imp for: " + os.GetName())))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}
