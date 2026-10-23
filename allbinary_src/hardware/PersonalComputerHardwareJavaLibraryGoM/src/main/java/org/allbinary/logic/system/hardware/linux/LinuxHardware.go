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
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ComponentInterfaceVector IBasicArrayList
    VideoInterfaceVector IBasicArrayList
    HardDriveControllerInterfaceVector IBasicArrayList
    CpuInterfaceVector IBasicArrayList
    UsbInterfaceVector IBasicArrayList
    EthernetInterfaceVector IBasicArrayList
    MultimediaInterfaceVector IBasicArrayList
    FireWireInterfaceVector IBasicArrayList
    BridgeInterfaceVector IBasicArrayList
    HardDriveInterfaceVector IBasicArrayList
    MacInterfaceVector IBasicArrayList
    MonitorInterfaceVector IBasicArrayList
    NAME string
    MINHARDWARE int
}

// Implements HardwareInterface

var LinuxHardware_PCIFILE string = "/proc/pci"

func (this *LinuxHardware) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    this.NAME = "Linux Hardware Profile"
    this.MINHARDWARE = 5
}




// Java overloads of new LinuxHardware
func NewLinuxHardware(args ...any) *LinuxHardware {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewLinuxHardware_1(Java_Arg[string](args[0]))
        }
        if len(args) == 0 {
            return NewLinuxHardware_2()
        }
    }
    panic("No Java overload of new LinuxHardware for the arguments")
}
func NewLinuxHardware_1(path string) *LinuxHardware {
    this := &LinuxHardware{}
    this.initFields()
        this.Init(path)
    return this
}


func NewLinuxHardware_2() *LinuxHardware {
    this := &LinuxHardware{}
    this.initFields()
        this.Init(LinuxHardware_PCIFILE)
        if this.ComponentInterfaceVector.Size() < this.MINHARDWARE {
            panic(errors.New(fmt.Sprint("Not Enough Data For A Valid License On Linux")))
        }
        var cpu *Cpu = NewCpu(); _ = cpu
        if cpu != nil {
            this.CpuInterfaceVector.Add(cpu)
            this.ComponentInterfaceVector.Add(cpu)
        }
        if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
            this.LogUtil.PutF("Hardware Data: " + this.ToString(), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *LinuxHardware) Init_1(filePath string) {
        var lineNumberReader *LineNumberReader = nil; _ = lineNumberReader
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
                this.LogUtil.Put("Hardware Data: " + this.ToString(), this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(e)
                }
            }()
            this.Init(lineNumberReader, filePath)
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of LinuxHardware.init
func (this *LinuxHardware) Init(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            this.Init_1(Java_Arg[string](args[0]))
            return
        }
        if len(args) == 2 && Java_IsArg[*LineNumberReader](args[0], pass) && Java_IsArg[string](args[1], pass) {
            this.Init_2(Java_Arg[*LineNumberReader](args[0]), Java_Arg[string](args[1]))
            return
        }
    }
    panic("No Java overload of LinuxHardware.init for the arguments")
}


// throws Exception
func (this *LinuxHardware) Init_2(lineNumberReader *LineNumberReader, filePath string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            if lineNumberReader != nil {
                lineNumberReader.Close()
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
            var pciFile *FileReader = NewFileReader(filePath); _ = pciFile
            lineNumberReader = NewLineNumberReader(pciFile)
            if lineNumberReader == nil {
                var fileVector IBasicArrayList = SubDirectory_GetInstance().Search(filePath, AbFile_CreateAbFile(FilePathData_GetInstance().PATH_START)); _ = fileVector
                if !fileVector.IsEmpty() {
                    var file *AbFile = any(fileVector.Get(0)).(*AbFile); _ = file
                    lineNumberReader = NewLineNumberReader(NewFileReader(file.GetPath()))
                }
                if lineNumberReader == nil {
                    panic(errors.New(fmt.Sprint("No Linux File Data")))
                }
            }
            if lineNumberReader != nil {
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
                    this.LogUtil.PutF("PCI File Found", this, this.CommonStrings.CONSTRUCTOR)
                }
                var nextLine string = lineNumberReader.ReadLine(); _ = nextLine
                for lineNumberReader != nil && nextLine != "" {
                    if this.IsNextHardware(nextLine) {
                        break
                    }
                    nextLine = lineNumberReader.ReadLine()
                }
                for this.IsNextHardware(nextLine) {
                    var componentData *StringMaker = NewStringMaker(); _ = componentData
                    if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
                        this.LogUtil.PutF("Found Hardware Device: " + fmt.Sprint(this.ComponentInterfaceVector.Size()), this, this.CommonStrings.CONSTRUCTOR)
                    }
                    componentData.Append(nextLine)
                    componentData.Append(CommonSeps_GetInstance().NEW_LINE)
                    nextLine = lineNumberReader.ReadLine()
                    var componentType string = PCComponentFactory_GetInstance().GetComponentType(nextLine); _ = componentType
                    for lineNumberReader != nil {
                        componentData.Append(nextLine)
                        componentData.Append(CommonSeps_GetInstance().NEW_LINE)
                        nextLine = lineNumberReader.ReadLine()
                        if nextLine == "" || this.IsNextHardware(nextLine) {
                            break
                        }
                    }
                    var componentInterface HardwareComponentInterface = PCComponentFactory_GetInstance().GetInstance(componentType, fmt.Sprint(componentData)); _ = componentInterface
                    if componentInterface != nil {
                        this.ComponentInterfaceVector.Add(componentInterface)
                    }
                }
            } else {
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
                    this.LogUtil.PutF("Could not load PCI File", this, this.CommonStrings.CONSTRUCTOR)
                }
            }
            lineNumberReader.Close()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *LinuxHardware) GetMultimedia(index int) MediaInterface {
        return any(this.MultimediaInterfaceVector.Get(index)).(MediaInterface)
}


func (this *LinuxHardware) GetBridge(index int) BridgeInterface {
        return any(this.BridgeInterfaceVector.Get(index)).(BridgeInterface)
}


func (this *LinuxHardware) GetCpu(index int) CpuInterface {
        return any(this.CpuInterfaceVector.Get(index)).(CpuInterface)
}


func (this *LinuxHardware) GetEthernet(index int) EthernetInterface {
        return any(this.EthernetInterfaceVector.Get(index)).(EthernetInterface)
}


func (this *LinuxHardware) GetFireWire(index int) FireWireInterface {
        return any(this.FireWireInterfaceVector.Get(index)).(FireWireInterface)
}


func (this *LinuxHardware) GetHardDriveController(index int) HardDriveControllerInterface {
        return any(this.HardDriveControllerInterfaceVector.Get(index)).(HardDriveControllerInterface)
}


func (this *LinuxHardware) GetHardDrive(index int) HardDriveInterface {
        return any(this.HardDriveInterfaceVector.Get(index)).(HardDriveInterface)
}


func (this *LinuxHardware) GetMachineAccessControlAddress(index int) MachineAccessControlAddressInterface {
        return any(this.MacInterfaceVector.Get(index)).(MachineAccessControlAddressInterface)
}


func (this *LinuxHardware) GetMonitor(index int) MonitorInterface {
        return any(this.MonitorInterfaceVector.Get(index)).(MonitorInterface)
}


func (this *LinuxHardware) GetUsb(index int) UsbInterface {
        return any(this.UsbInterfaceVector.Get(index)).(UsbInterface)
}


func (this *LinuxHardware) GetVideo(index int) VideoInterface {
        return any(this.VideoInterfaceVector.Get(index)).(VideoInterface)
}


func (this *LinuxHardware) GetComponent(index int) HardwareComponentInterface {
        return any(this.ComponentInterfaceVector.Get(index)).(HardwareComponentInterface)
}


func (this *LinuxHardware) ToString() string {
        var hardwareBuffer *StringMaker = NewStringMaker(); _ = hardwareBuffer
        var size int = this.ComponentInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var componentInterface HardwareComponentInterface = any(this.ComponentInterfaceVector.Get(index)).(HardwareComponentInterface); _ = componentInterface
            hardwareBuffer.Append("Component ")
            hardwareBuffer.Appendint(index)
            hardwareBuffer.Append(": \n")
            hardwareBuffer.Append(fmt.Sprint(componentInterface))
            hardwareBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        }
        return fmt.Sprint(hardwareBuffer)
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
        if nextLine != "" {
            var index int = strings.Index(nextLine, string("Bus")); _ = index
            if index >= 0 && index < 4 {
                return true
            }
        }
        return false
}

