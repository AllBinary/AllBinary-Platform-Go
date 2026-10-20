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

type LinuxHardware struct {
}

// Implements HardwareInterface



func NewLinuxHardware() *LinuxHardware {
    this := &LinuxHardware{}
    return this
}


func (this *LinuxHardware) GetMultimedia(index int) MediaInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetBridge(index int) BridgeInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetCpu(index int) CpuInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetEthernet(index int) EthernetInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetFireWire(index int) FireWireInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetHardDriveController(index int) HardDriveControllerInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetHardDrive(index int) HardDriveInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetMachineAccessControlAddress(index int) MachineAccessControlAddressInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetMonitor(index int) MonitorInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetUsb(index int) UsbInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetVideo(index int) VideoInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) GetComponent(index int) HardwareComponentInterface {
        panic(errors.New(fmt.Sprint()))
}


func (this *LinuxHardware) ToString() string {
        return StringUtil_GetInstance().EMPTY_STRING
}

func (this *LinuxHardware) String() string {
    return this.ToString()
}


func (this *LinuxHardware) CompareTo(hardwareInterface HardwareInterface) bool {
        return true
}


func (this *LinuxHardware) Difference(hardwareInterface HardwareInterface) *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}


func (this *LinuxHardware) IsNextHardware(nextLine string) bool {
        return false
}
