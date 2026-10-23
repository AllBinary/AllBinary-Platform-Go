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

type AndroidHardware struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ComponentInterfaceVector IBasicArrayList
    PROC_BUS_INPUT_DIRECTORY string
    DEVICES string
    MINHARDWARE int
}

// Implements HardwareInterface


func (this *AndroidHardware) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ComponentInterfaceVector = NewBasicArrayListD()
    this.PROC_BUS_INPUT_DIRECTORY = "/proc/bus/input/"
    this.DEVICES = this.PROC_BUS_INPUT_DIRECTORY + "devices/"
    this.MINHARDWARE = 3
}


func NewAndroidHardware() *AndroidHardware {
    this := &AndroidHardware{}
    this.initFields()
        this.Init(this.DEVICES)
        var size int = this.ComponentInterfaceVector.Size(); _ = size
        if size < this.MINHARDWARE {
            panic(errors.New(fmt.Sprint("Not Enough Data For A Valid License On Linux")))
        }
    return this
}


// throws Exception
func (this *AndroidHardware) Init(filePath string) {
        var lineNumberReader *Closeable = NullCloseable_NULL_CLOSEABLE; _ = lineNumberReader
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Hardware Data: " + this.ToString(), this, this.CommonStrings.INIT, e)
            panic(e)
                }
            if lineNumberReader != nil {
                lineNumberReader.Close()
            }
            }()
            lineNumberReader = this.Get(filePath)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *AndroidHardware) Get(filePath string) *LineNumberReader {
        this.ComponentInterfaceVector = NewBasicArrayListD()
        var pciFile *FileReader = NewFileReader(filePath); _ = pciFile
        var lineNumberReader *LineNumberReader = NewLineNumberReader(pciFile); _ = lineNumberReader
        this.LogUtil.PutF("File Found", this, this.CommonStrings.CONSTRUCTOR)
        var nextLine string = lineNumberReader.ReadLine(); _ = nextLine
        for nextLine != "" {
            nextLine = lineNumberReader.ReadLine()
            this.ComponentInterfaceVector.Add(NewUnknownHardware(nextLine))
        }
        return lineNumberReader
}


func (this *AndroidHardware) GetComponent(index int) HardwareComponentInterface {
        return any(this.ComponentInterfaceVector.Get(index)).(HardwareComponentInterface)
}


func (this *AndroidHardware) ToString() string {
        var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
        var size int = this.ComponentInterfaceVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var componentInterface HardwareComponentInterface = any(this.ComponentInterfaceVector.Get(index)).(HardwareComponentInterface); _ = componentInterface
            stringBuilder.Append(fmt.Sprint(componentInterface))
            stringBuilder.Append(CommonSeps_GetInstance().NEW_LINE)
        }
        return fmt.Sprint(stringBuilder)
}

func (this *AndroidHardware) String() string {
    return this.ToString()
}


func (this *AndroidHardware) CompareTo(hardwareInterface HardwareInterface) bool {
        return true
}


func (this *AndroidHardware) Difference(hardwareInterface HardwareInterface) *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}

