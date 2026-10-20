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
}

// Implements HardwareInterface



func NewWindowsHardware() *WindowsHardware {
    this := &WindowsHardware{}
    return this
}


func (this *WindowsHardware) GetMultimedia(index int) MediaInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetBridge(index int) BridgeInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetCpu(index int) CpuInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetEthernet(index int) EthernetInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetFireWire(index int) FireWireInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetHardDriveController(index int) HardDriveControllerInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetHardDrive(index int) HardDriveInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetMachineAccessControlAddress(index int) MachineAccessControlAddressInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetMonitor(index int) MonitorInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetUsb(index int) UsbInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetVideo(index int) VideoInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) GetComponent(index int) HardwareComponentInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *WindowsHardware) ToString() string {
        return StringUtil_GetInstance().EMPTY_STRING
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
