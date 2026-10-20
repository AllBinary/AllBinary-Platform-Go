/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

//Android
type TsUtil struct {
}

var TsUtil_Instance *TsUtil = NewTsUtil()

func NewTsUtil() *TsUtil {
    this := &TsUtil{}
    return this
}


func TsUtil_GetInstance() *TsUtil {
        return TsUtil_Instance
}


func (this *TsUtil) HashCode(anyType any) int {
        return Java_HashCode(anyType)
}


// throws InterruptedException
func (this *TsUtil) WaitFor(any_ any, timeoutMillis int64) {
        var anyType any = any(any_).(any); _ = anyType
        Java_Wait(anyType, int64(timeoutMillis))
}


// throws RecordStoreNotOpenException, InvalidRecordIDException, RecordStoreException
func (this *TsUtil) GetRecord(anyType any, recordId int) []int8 {
        var recordStore IRecordStore = any(anyType).(IRecordStore); _ = recordStore
        var data []int8; _ = data
        { // synchronized (this)
            data = make([]int8, recordStore.GetRecordSize(recordId))
            recordStore.GetRecord(recordId, data, 0).(int)
        }
        return func() any { if len(data) < 1 { return NullUtil_GetInstance().NULL_BYTE_ARRAY }; return data }()
}


func (this *TsUtil) CompareTo(a string, b string) int {
        return strings.Compare(a, b)
}


func (this *TsUtil) EqualIgnoreCase(a string, b string) bool {
        return strings.EqualFold(a, b)
}


func (this *TsUtil) EqualsNotstring(anyType any, object2 any) bool {
        return Java_Equals(anyType, object2)
}


func (this *TsUtil) GetClassClassLoader(anyType any) *ClassLoader {
        return reflect.TypeOf(anyType).GetClassLoader()
}


func (this *TsUtil) ToNumberFromChar(value rune) int {
        return int(value)
}


func (this *TsUtil) ToNumber(value int) int {
        return value
}


func (this *TsUtil) GetByteArray(string_ string) []int8 {
        return Java_GetBytes(string_)
}


func (this *TsUtil) IsItemListener(owner IDisplayable) bool {
        if func() bool { _, ok := any(owner).(IForm); return ok }() {
            return true
        } else {
            return false
        }
}


func (this *TsUtil) Gc() {
        runtime.GC()
}

