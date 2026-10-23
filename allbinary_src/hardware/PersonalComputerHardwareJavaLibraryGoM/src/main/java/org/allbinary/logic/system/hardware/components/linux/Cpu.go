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

type Cpu struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CPUFILE string
    CpuHashMap *HashMap
}

// Implements CpuInterface, HardwareComponentInterface


func (this *Cpu) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CPUFILE = "/proc/cpuinfo"
    
}




// Java overloads of new Cpu
func NewCpu(args ...any) *Cpu {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewCpu_1(Java_Arg[string](args[0]))
        }
        if len(args) == 0 {
            return NewCpu_2()
        }
    }
    panic("No Java overload of new Cpu for the arguments")
}
func NewCpu_1(filePath string) *Cpu {
    this := &Cpu{}
    this.initFields()
        this.Init(filePath)
    return this
}


func NewCpu_2() *Cpu {
    this := &Cpu{}
    this.initFields()
        this.Init(this.CPUFILE)
    return this
}


// throws Exception
func (this *Cpu) Init_1(filePath string) {
        var lineNumberReader *LineNumberReader = nil; _ = lineNumberReader
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
                this.LogUtil.Put(fmt.Sprint(NewStringMaker().Append("Cpu Data: ").Append(this.ToString())), this, this.CommonStrings.CONSTRUCTOR, e)
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

// Java overloads of Cpu.init
func (this *Cpu) Init(args ...any) {
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
    panic("No Java overload of Cpu.init for the arguments")
}


// throws Exception
func (this *Cpu) Init_2(lineNumberReader *LineNumberReader, filePath string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            if lineNumberReader != nil {
                lineNumberReader.Close()
            }
            }()
            var fileReader *FileReader = NewFileReader(filePath); _ = fileReader
            lineNumberReader = NewLineNumberReader(fileReader)
            this.CpuHashMap = StdUtil_GetInstance().CreateHashMap()
            if lineNumberReader == nil {
                var fileVector IBasicArrayList = SubDirectory_GetInstance().Search(filePath, AbFile_CreateAbFile(FilePathData_GetInstance().PATH_START)); _ = fileVector
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().OS) {
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Cpu File BasicArrayList Size: ").Appendint(fileVector.Size())), this, this.CommonStrings.CONSTRUCTOR)
                }
                if !fileVector.IsEmpty() {
                    var file *AbFile = any(fileVector.Get(0)).(*AbFile); _ = file
                    lineNumberReader = NewLineNumberReader(NewFileReader(file.GetPath()))
                }
            }
            if lineNumberReader != nil {
                var nextLine string = lineNumberReader.ReadLine(); _ = nextLine
                for lineNumberReader != nil && nextLine != "" {
                    var index int = strings.Index(nextLine, string(':')); _ = index
                    if index >= 0 {
                        var key string = nextLine[0:index]; _ = key
                        var valueIndex int = index + 1; _ = valueIndex
                        var value string = nextLine[valueIndex:len(nextLine)]; _ = value
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_PROCESSOR))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_PROCESSOR, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_VENDORID))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_VENDORID, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_CPUFAMILY))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_CPUFAMILY, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_MODEL))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_MODEL, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_MODELNAME))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_MODELNAME, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_STEPPING))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_STEPPING, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_CPUMHZ))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_CPUMHZ, value)
                        }
                        if strings.Index(key, string(fmt.Sprint(CpuInfo_CACHESIZE))) >= 0 {
                            this.CpuHashMap.Put(CpuInfo_CACHESIZE, value)
                        }
                    }
                    nextLine = lineNumberReader.ReadLine()
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *Cpu) GetProcessor() string {
        return any(this.CpuHashMap.Get(CpuInfo_PROCESSOR)).(string)
}


func (this *Cpu) GetVendorId() string {
        return any(this.CpuHashMap.Get(CpuInfo_VENDORID)).(string)
}


func (this *Cpu) GetCpuFamily() string {
        return any(this.CpuHashMap.Get(CpuInfo_CPUFAMILY)).(string)
}


func (this *Cpu) GetModel() string {
        return any(this.CpuHashMap.Get(CpuInfo_MODEL)).(string)
}


func (this *Cpu) GetModelName() string {
        return any(this.CpuHashMap.Get(CpuInfo_MODELNAME)).(string)
}


func (this *Cpu) GetStepping() string {
        return any(this.CpuHashMap.Get(CpuInfo_STEPPING)).(string)
}


func (this *Cpu) GetCpuSpeed() string {
        return any(this.CpuHashMap.Get(CpuInfo_CPUMHZ)).(string)
}


func (this *Cpu) GetCacheSize() string {
        return any(this.CpuHashMap.Get(CpuInfo_CACHESIZE)).(string)
}


func (this *Cpu) CompareTo_1(cpuInterface CpuInterface) bool {
        return true
}

// Java overloads of Cpu.compareTo
func (this *Cpu) CompareTo(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[CpuInterface](args[0], pass) {
            return this.CompareTo_1(Java_Arg[CpuInterface](args[0]))
        }
        if len(args) == 1 && Java_IsArg[HardwareComponentInterface](args[0], pass) {
            return this.CompareTo_2(Java_Arg[HardwareComponentInterface](args[0]))
        }
    }
    panic("No Java overload of Cpu.compareTo for the arguments")
}


func (this *Cpu) CompareTo_2(componentInterface HardwareComponentInterface) bool {
        return true
}


func (this *Cpu) ToString() string {
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        stringMaker.Append(fmt.Sprint(CpuInfo_PROCESSOR)).Append(commonLabels.COLON_SEP).Append(this.GetProcessor())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_VENDORID)).Append(commonLabels.COLON_SEP).Append(this.GetVendorId())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_CPUFAMILY)).Append(commonLabels.COLON_SEP).Append(this.GetCpuFamily())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_MODEL)).Append(commonLabels.COLON_SEP).Append(this.GetModel())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_MODELNAME)).Append(commonLabels.COLON_SEP).Append(this.GetModelName())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_STEPPING)).Append(commonLabels.COLON_SEP).Append(this.GetStepping())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_CPUMHZ)).Append(commonLabels.COLON_SEP).Append(this.GetCpuSpeed())
        stringMaker.Append(commonSeps.NEW_LINE)
        stringMaker.Append(fmt.Sprint(CpuInfo_CACHESIZE)).Append(commonLabels.COLON_SEP).Append(this.GetCacheSize())
        return fmt.Sprint(stringMaker)
}

func (this *Cpu) String() string {
    return this.ToString()
}

