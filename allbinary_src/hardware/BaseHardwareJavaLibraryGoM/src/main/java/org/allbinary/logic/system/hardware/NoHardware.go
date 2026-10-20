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

type NoHardware struct {
}

// Implements HardwareInterface

var NoHardware_Instance HardwareInterface = NewNoHardware()

func NewNoHardware() *NoHardware {
    this := &NoHardware{}
    return this
}


func NoHardware_GetInstance() HardwareInterface {
        return NoHardware_Instance
}


func (this *NoHardware) GetComponent(index int) HardwareComponentInterface {
        return NoHardwareComponent_GetInstance()
}


func (this *NoHardware) ToString() string {
        return "No Hardware"
}

func (this *NoHardware) String() string {
    return this.ToString()
}


func (this *NoHardware) CompareTo(hardwareInterface HardwareInterface) bool {
        return true
}


func (this *NoHardware) Difference(hardwareInterface HardwareInterface) *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}
