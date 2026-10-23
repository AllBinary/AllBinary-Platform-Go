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

type ComponentFactory struct {
    VGA *Component
    IDE *Component
    SCSI *Component
    BRIDGE *Component
    ETHERNET *Component
    USB *Component
    MULTIMEDIA *Component
    PCI *Component
    RAID *Component
    SMBUS *Component
    PIC *Component
    Class *Component
    UNKNOWN *Component
}

var ComponentFactory_Instance *ComponentFactory = NewComponentFactory()

func (this *ComponentFactory) initFields() {
    this.VGA = NewComponent("VGA")
    this.IDE = NewComponent("IDE")
    this.SCSI = NewComponent("SCSI")
    this.BRIDGE = NewComponent("bridge")
    this.ETHERNET = NewComponent("Ethernet")
    this.USB = NewComponent("USB")
    this.MULTIMEDIA = NewComponent("Multimedia")
    this.PCI = NewComponent("PCI")
    this.RAID = NewComponent("RAID")
    this.SMBUS = NewComponent("SMBus")
    this.PIC = NewComponent("PIC")
    this.Class = NewComponent("Class")
    this.UNKNOWN = NewComponent(CommonStrings_GetInstance().UNKNOWN)
}

func NewComponentFactory() *ComponentFactory {
    this := &ComponentFactory{}
    this.initFields()
    return this
}


func ComponentFactory_GetInstance() *ComponentFactory {
        return ComponentFactory_Instance
}

