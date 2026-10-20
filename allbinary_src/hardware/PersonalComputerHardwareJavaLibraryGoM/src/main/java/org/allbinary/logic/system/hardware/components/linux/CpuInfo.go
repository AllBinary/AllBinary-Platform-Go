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

type CpuInfo struct {
    Name string
}

var CpuInfo_PROCESSOR *CpuInfo = NewCpuInfo("processor")
var CpuInfo_VENDORID *CpuInfo = NewCpuInfo("vendor_id")
var CpuInfo_CPUFAMILY *CpuInfo = NewCpuInfo("cpu family")
var CpuInfo_MODEL *CpuInfo = NewCpuInfo("model")
var CpuInfo_MODELNAME *CpuInfo = NewCpuInfo("model name")
var CpuInfo_STEPPING *CpuInfo = NewCpuInfo("stepping")
var CpuInfo_CPUMHZ *CpuInfo = NewCpuInfo("cpu MHz")
var CpuInfo_CACHESIZE *CpuInfo = NewCpuInfo("cache size")
var CpuInfo_FDIVBUG *CpuInfo = NewCpuInfo("fdiv_bug")
var CpuInfo_HLTBUG *CpuInfo = NewCpuInfo("hlt_bug")
var CpuInfo_F00FBUG *CpuInfo = NewCpuInfo("f00f_bug")
var CpuInfo_COMABUG *CpuInfo = NewCpuInfo("coma_bug")
var CpuInfo_FPU *CpuInfo = NewCpuInfo("fpu")
var CpuInfo_FPUEXCEPTION *CpuInfo = NewCpuInfo("fpu_exception")
var CpuInfo_CPUIDLEVEL *CpuInfo = NewCpuInfo("cache size")
var CpuInfo_WP *CpuInfo = NewCpuInfo("yes")
var CpuInfo_FLAGS *CpuInfo = NewCpuInfo("flags")
var CpuInfo_BOGOMIPS *CpuInfo = NewCpuInfo("bogomips")

func (this *CpuInfo) initFields() {
    
}


func NewCpuInfo(name string) *CpuInfo {
    this := &CpuInfo{}
    this.initFields()
        this.Name = name
    return this
}


func (this *CpuInfo) ToString() string {
        return this.Name
}

func (this *CpuInfo) String() string {
    return this.ToString()
}

