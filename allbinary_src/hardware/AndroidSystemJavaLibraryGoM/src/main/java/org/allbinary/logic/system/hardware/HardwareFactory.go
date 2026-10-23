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
    LogUtil *LogUtil
    TsUtil *TsUtil
}

var HardwareFactory_Instance *HardwareFactory = NewHardwareFactory()

func (this *HardwareFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
}


func HardwareFactory_GetInstance() *HardwareFactory {
        return HardwareFactory_Instance
}


func NewHardwareFactory() *HardwareFactory {
    this := &HardwareFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *HardwareFactory) GetInstance(os IGenericOperatingSystem) HardwareInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 HardwareInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put("Ignoring Exception Returning NoHardware", this, commonStrings.GET_INSTANCE, e)
            tryControl1, tryResult1 = 1, NewNoHardware(); return
                }
            }()
            var operatingSystems *OperatingSystems = OperatingSystems_GetInstance(); _ = operatingSystems
            if this.TsUtil.CompareTo(os.GetName(), operatingSystems.ANDROID) == 0 {
                tryControl1, tryResult1 = 1, NewAndroidHardware(); return
            } else if this.TsUtil.CompareTo(os.GetName(), operatingSystems.HARMONY) == 0 {
                tryControl1, tryResult1 = 1, NewAndroidHardware(); return
            }
            panic(errors.New(fmt.Sprint("No Hardware Imp for: " + os.GetName())))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

