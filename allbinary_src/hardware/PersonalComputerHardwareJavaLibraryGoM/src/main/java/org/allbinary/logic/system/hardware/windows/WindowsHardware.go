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

type WindowsHardware struct {
    ComponentInterfaceVector *BasicArrayList
    VideoInterfaceVector *BasicArrayList
    HardDriveControllerInterfaceVector *BasicArrayList
    CpuInterfaceVector *BasicArrayList
    UsbInterfaceVector *BasicArrayList
    EthernetInterfaceVector *BasicArrayList
    MultimediaInterfaceVector *BasicArrayList
    FireWireInterfaceVector *BasicArrayList
    BridgeInterfaceVector *BasicArrayList
    HardDriveInterfaceVector *BasicArrayList
    MacInterfaceVector *BasicArrayList
    MonitorInterfaceVector *BasicArrayList
    NAME string
    MINHARDWARE int
    ComponentFactory *ComponentFactory
}

// Implements HardwareInterface


func (this *WindowsHardware) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    this.NAME = "Windows Hardware Profile"
    this.MINHARDWARE = 0
    this.ComponentFactory = ComponentFactory_GetInstance()
}


func NewWindowsHardware() *WindowsHardware {
    this := &WindowsHardware{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            this.ComponentInterfaceVector = NewBasicArrayListD()
            this.VideoInterfaceVector = NewBasicArrayListD()
            this.HardDriveControllerInterfaceVector = NewBasicArrayListD()
            this.CpuInterfaceVector = NewBasicArrayListD()
            this.UsbInterfaceVector = NewBasicArrayListD()
            this.EthernetInterfaceVector = NewBasicArrayListD()
            this.MultimediaInterfaceVector = NewBasicArrayListD()
            this.FireWireInterfaceVector = NewBasicArrayListD()
            this.BridgeInterfaceVector = NewBasicArrayListD()
            this.HardDriveInterfaceVector = NewBasicArrayListD()
            this.MacInterfaceVector = NewBasicArrayListD()
            this.MonitorInterfaceVector = NewBasicArrayListD()
            if this.ComponentInterfaceVector.Size() < this.MINHARDWARE {
                panic(errors.New(fmt.Sprint("Not Enough Data For A Valid License On Windows")))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *WindowsHardware) GetComponentType(component string) string {
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
        } else {
            return nil
        }
}


func (this *WindowsHardware) GetMultimedia(index int) MediaInterface {
        return any(this.MultimediaInterfaceVector.Get(index)).(MediaInterface)
}


func (this *WindowsHardware) GetBridge(index int) BridgeInterface {
        return any(this.BridgeInterfaceVector.Get(index)).(BridgeInterface)
}


func (this *WindowsHardware) GetCpu(index int) CpuInterface {
        return any(this.CpuInterfaceVector.Get(index)).(CpuInterface)
}


func (this *WindowsHardware) GetEthernet(index int) EthernetInterface {
        return any(this.EthernetInterfaceVector.Get(index)).(EthernetInterface)
}


func (this *WindowsHardware) GetFireWire(index int) FireWireInterface {
        return any(this.FireWireInterfaceVector.Get(index)).(FireWireInterface)
}


func (this *WindowsHardware) GetHardDriveController(index int) HardDriveControllerInterface {
        return any(this.HardDriveControllerInterfaceVector.Get(index)).(HardDriveControllerInterface)
}


func (this *WindowsHardware) GetHardDrive(index int) HardDriveInterface {
        return any(this.HardDriveInterfaceVector.Get(index)).(HardDriveInterface)
}


func (this *WindowsHardware) GetMachineAccessControlAddress(index int) MachineAccessControlAddressInterface {
        return any(this.MacInterfaceVector.Get(index)).(MachineAccessControlAddressInterface)
}


func (this *WindowsHardware) GetMonitor(index int) MonitorInterface {
        return any(this.MonitorInterfaceVector.Get(index)).(MonitorInterface)
}


func (this *WindowsHardware) GetUsb(index int) UsbInterface {
        return any(this.UsbInterfaceVector.Get(index)).(UsbInterface)
}


func (this *WindowsHardware) GetVideo(index int) VideoInterface {
        return any(this.VideoInterfaceVector.Get(index)).(VideoInterface)
}


func (this *WindowsHardware) GetComponent(index int) HardwareComponentInterface {
        return any(this.ComponentInterfaceVector.Get(index)).(HardwareComponentInterface)
}


func (this *WindowsHardware) ToString() string {
        var hardwareBuffer *StringMaker = NewStringMaker(); _ = hardwareBuffer
        var size int = this.ComponentInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var componentInterface HardwareComponentInterface = any(this.ComponentInterfaceVector.Get(index)).(HardwareComponentInterface); _ = componentInterface
            hardwareBuffer.Append(fmt.Sprint(componentInterface))
            hardwareBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        }
        return fmt.Sprint(hardwareBuffer)
}

func (this *WindowsHardware) String() string {
    return this.ToString()
}


func (this *WindowsHardware) CompareTo(hardwareInterface HardwareInterface) bool {
        return true
}


func (this *WindowsHardware) Difference(hardwareInterface HardwareInterface) *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}

