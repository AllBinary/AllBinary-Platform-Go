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

type OperatingSystems struct {
    LINUX string
    WINDOWS string
    SOLARIS string
    WINDOWS2000 string
    WINDOWS_NT string
    WINDOWS_10 string
    WINDOWS_11 string
    WINDOWS_VISTA string
    WINDOWS_XP string
    ANDROID string
    HARMONY string
    UnknownSpecificOSAllowed bool
}

var OperatingSystems_Instance *OperatingSystems = NewOperatingSystems()

func (this *OperatingSystems) initFields() {
    this.LINUX = "Linux"
    this.WINDOWS = "Windows"
    this.SOLARIS = "Solaris"
    this.WINDOWS2000 = "Windows 2000"
    this.WINDOWS_NT = "Windows NT"
    this.WINDOWS_10 = "Windows 10"
    this.WINDOWS_11 = "Windows 11"
    this.WINDOWS_VISTA = "Windows Vista"
    this.WINDOWS_XP = "Windows XP"
    this.ANDROID = this.LINUX
    this.HARMONY = "HarmonyOS"
    this.UnknownSpecificOSAllowed = true
}


func NewOperatingSystems() *OperatingSystems {
    this := &OperatingSystems{}
    this.initFields()
    return this
}


func (this *OperatingSystems) IsUnknownSpecificOSAllowed() bool {
        return this.UnknownSpecificOSAllowed
}


func (this *OperatingSystems) SetUnknownSpecificOSAllowed(aUnknownSpecificOSAllowed bool) {
        this.UnknownSpecificOSAllowed = aUnknownSpecificOSAllowed
}


func OperatingSystems_GetInstance() *OperatingSystems {
        return OperatingSystems_Instance
}
