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

type GenericOperatingSystem struct {
    OsName string
    OsArch string
    OsVersion string
    CharacterSet *CharacterSet
    Scalable bool
}

// Implements OperatingSystemInterface

// The Java class GenericOperatingSystem has subclasses so its variables are this interface
type IGenericOperatingSystem interface {
    GenericOperatingSystem_() *GenericOperatingSystem
    GetName() string
    GetVersion() string
    GetArch() string
    GetCharacterSet() *CharacterSet
    IsOverScan() bool
    GetOverScanXPercent() int
    GetOverScanYPercent() int
    IsScalable() bool
    IsAutoHide() bool
    ToString() string
}

func (this *GenericOperatingSystem) GenericOperatingSystem_() *GenericOperatingSystem {
    return this
}


func (this *GenericOperatingSystem) initFields() {
    
    
    
    this.CharacterSet = NewCharacterSet()
    this.Scalable = false
}


func NewGenericOperatingSystem() *GenericOperatingSystem {
    this := &GenericOperatingSystem{}
    this.initFields()
        var systemProperties *SystemProperties = SystemProperties_GetInstance(); _ = systemProperties
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        var EMPTY_STRING string = StringUtil_GetInstance().EMPTY_STRING; _ = EMPTY_STRING
        if stringValidationUtil.IsEmpty(systemProperties.GetName()) {
            this.OsName = EMPTY_STRING
        } else {
            this.OsName = systemProperties.GetName()
        }
        if stringValidationUtil.IsEmpty(systemProperties.GetArch()) {
            this.OsArch = EMPTY_STRING
        } else {
            this.OsArch = systemProperties.GetArch()
        }
        if stringValidationUtil.IsEmpty(systemProperties.GetVersion()) {
            this.OsVersion = EMPTY_STRING
        } else {
            this.OsVersion = systemProperties.GetVersion()
        }
    return this
}


func (this *GenericOperatingSystem) GetName() string {
        return this.OsName
}


func (this *GenericOperatingSystem) GetVersion() string {
        return this.OsVersion
}


func (this *GenericOperatingSystem) GetArch() string {
        return this.OsArch
}


func (this *GenericOperatingSystem) GetCharacterSet() *CharacterSet {
        return this.CharacterSet
}


func (this *GenericOperatingSystem) IsOverScan() bool {
        return false
}


func (this *GenericOperatingSystem) GetOverScanXPercent() int {
        return 100
}


func (this *GenericOperatingSystem) GetOverScanYPercent() int {
        return 100
}


func (this *GenericOperatingSystem) IsScalable() bool {
        return this.Scalable
}


func (this *GenericOperatingSystem) IsAutoHide() bool {
        return false
}


func (this *GenericOperatingSystem) ToString() string {
        var NEW_LINE string = CommonSeps_GetInstance().NEW_LINE; _ = NEW_LINE
        var osBuffer *StringMaker = NewStringMaker(); _ = osBuffer
        osBuffer.Append("Operating System Info: \n")
        osBuffer.Append(this.GetName())
        osBuffer.Append(NEW_LINE)
        osBuffer.Append(this.GetArch())
        osBuffer.Append(NEW_LINE)
        osBuffer.Append(this.GetVersion())
        osBuffer.Append(NEW_LINE)
        osBuffer.Append(StringUtil_GetInstance().ToString(this.GetCharacterSet()))
        osBuffer.Append(NEW_LINE)
        return fmt.Sprint(osBuffer)
}

func (this *GenericOperatingSystem) String() string {
    return this.ToString()
}
