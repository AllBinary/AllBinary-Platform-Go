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

type PCComponentFactory struct {
    LogUtil *LogUtil
    ComponentFactory *ComponentFactory
}

var PCComponentFactory_Instance *PCComponentFactory = NewPCComponentFactory()

func (this *PCComponentFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.ComponentFactory = ComponentFactory_GetInstance()
}


func PCComponentFactory_GetInstance() *PCComponentFactory {
        return PCComponentFactory_Instance
}


func NewPCComponentFactory() *PCComponentFactory {
    this := &PCComponentFactory{}
    this.initFields()
    return this
}


// Java overloaded method: Go needs a different name
func (this *PCComponentFactory) GetInstance(component string, componentData string) HardwareComponentInterface {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 HardwareComponentInterface) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().FACTORYERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET_INSTANCE, e)
            }
            tryControl1, tryResult1 = 1, NewUnknown(commonStrings.EXCEPTION); return
                }
            }()
            if component != "" {
                if strings.Compare(component, fmt.Sprint(this.ComponentFactory.ETHERNET)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewEthernet(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.BRIDGE)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewBridge(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.IDE)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewHardDriveController(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.MULTIMEDIA)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewMedia(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.SCSI)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewHardDriveController(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.USB)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewUsb(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.VGA)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewVideo(componentData)).(HardwareComponentInterface); return
                } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.PCI)) == 0 {
                    tryControl1, tryResult1 = 1, any(NewPCI(componentData)).(HardwareComponentInterface); return
                }
            }
            tryControl1, tryResult1 = 1, any(NewUnknown(componentData)).(HardwareComponentInterface); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *PCComponentFactory) GetComponentType(component string) string {
        if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.ETHERNET))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.ETHERNET)
        } else if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.BRIDGE))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.BRIDGE)
        } else if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.IDE))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.IDE)
        } else if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.MULTIMEDIA))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.MULTIMEDIA)
        } else if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.SCSI))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.SCSI)
        } else if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.USB))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.USB)
        } else if strings.Index(component, string(fmt.Sprint(this.ComponentFactory.VGA))) >= 0 {
            return fmt.Sprint(this.ComponentFactory.VGA)
        } else if strings.Compare(component, fmt.Sprint(this.ComponentFactory.PCI)) == 0 {
            return fmt.Sprint(this.ComponentFactory.PCI)
        }
        return fmt.Sprint(this.ComponentFactory.UNKNOWN)
}

