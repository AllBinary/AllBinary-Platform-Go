/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

type ABSystemWrapper struct {
}

var ABSystemWrapper_Instance *ABSystemWrapper = NewABSystemWrapper()

func NewABSystemWrapper() *ABSystemWrapper {
    this := &ABSystemWrapper{}
    return this
}


func ABSystemWrapper_GetInstance() *ABSystemWrapper {
        return ABSystemWrapper_Instance
}


func (this *ABSystemWrapper) Arraycopy(src any, srcPos int, dest any, destPos int, length int) {
        reflect.Copy(reflect.ValueOf(dest).Slice(destPos, destPos + length), reflect.ValueOf(src).Slice(srcPos, srcPos + length))
}


func (this *ABSystemWrapper) CurrentTimeMillis() int64 {
        return time.Now().UnixMilli()
}

