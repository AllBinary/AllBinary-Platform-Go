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

type PlatformKeyFactory struct {
}

var PlatformKeyFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func NewPlatformKeyFactory() *PlatformKeyFactory {
    this := &PlatformKeyFactory{}
    return this
}


func PlatformKeyFactory_GetInstance() *PlatformKeyFactory {
        if PlatformKeyFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            PlatformKeyFactory_Instance = NewPlatformKeyFactory()
        }
        return any(PlatformKeyFactory_Instance).(*PlatformKeyFactory)
}


func (this *PlatformKeyFactory) GetString(keyCode int) string {
        return PCGameKey_GetString(keyCode)
}


func (this *PlatformKeyFactory) IsSubmission(input IInput) bool {
        return PCKeyFactory_GetInstance().IsSubmission(input)
}


func (this *PlatformKeyFactory) IsDelete(input IInput) bool {
        return PCKeyFactory_GetInstance().IsDelete(input)
}


func (this *PlatformKeyFactory) IsBackSpace(input IInput) bool {
        return PCKeyFactory_GetInstance().IsBackSpace(input)
}


func (this *PlatformKeyFactory) IsUp(input IInput) bool {
        return PCKeyFactory_GetInstance().IsUp(input)
}


func (this *PlatformKeyFactory) IsDown(input IInput) bool {
        return PCKeyFactory_GetInstance().IsDown(input)
}


func (this *PlatformKeyFactory) IsEnter(input IInput) bool {
        return PCKeyFactory_GetInstance().IsEnter(input)
}


func (this *PlatformKeyFactory) IsLeft(input IInput) bool {
        return PCKeyFactory_GetInstance().IsLeft(input)
}


func (this *PlatformKeyFactory) IsRight(input IInput) bool {
        return PCKeyFactory_GetInstance().IsRight(input)
}

