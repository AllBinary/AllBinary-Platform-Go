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

type ByteUtil struct {
}

var ByteUtil_Instance *ByteUtil = NewByteUtil()


func ByteUtil_GetInstance() *ByteUtil {
        return ByteUtil_Instance
}


func NewByteUtil() *ByteUtil {
    this := &ByteUtil{}
    return this
}


func (this *ByteUtil) Xor(array []int8, mask int8) []int8 {
        var bytes []int8 = make([]int8, len(array)); _ = bytes
        var size int = len(array); _ = size
        var index int = 0; _ = index
        var currentByte int8; _ = currentByte
        for index < size {
            currentByte = array[index]
            bytes[index] = int8((int(currentByte) ^ int(mask)))
            index++
        }
        return bytes
}


func (this *ByteUtil) Not(array []int8) []int8 {
        var bytes []int8 = make([]int8, len(array)); _ = bytes
        var size int = len(array); _ = size
        var index int = 0; _ = index
        for index < size {
            bytes[index] = int8(^array[index])
            index++
        }
        return bytes
}


func (this *ByteUtil) XorByte(array []int8, index int) []int8 {
        var mask int8 = 0; _ = mask
        if index == 0 {
            mask = 1
        }
        if index == 1 {
            mask = 2
        }
        if index == 2 {
            mask = 4
        }
        if index == 3 {
            mask = 8
        }
        if index == 4 {
            mask = 16
        }
        if index == 5 {
            mask = 32
        }
        if index == 6 {
            mask = 64
        }
        if index == 7 {
            mask = int8(-127)
        }
        return this.Xor(array, mask)
}


func (this *ByteUtil) Compare(newByteArray []int8, oldByteArray []int8) bool {
        if len(newByteArray) != len(oldByteArray) {
            return false
        } else {
        }
        var size int = len(newByteArray); _ = size
        for index := int(0); index < size; index++ {
            if int(newByteArray[index]) != int(oldByteArray[index]) {
                return false
            }
        }
        return true
}


func (this *ByteUtil) CompareWithStats(newByteArray []int8, newSize int, oldByteArray []int8, stats []int) bool {
        if newSize != len(oldByteArray) {
            stats[0] = -1
            stats[2] = newSize
            stats[3] = len(oldByteArray)
            return false
        } else {
            stats[0] = newSize
        }
        stats[1] = 0
        var size int = newSize; _ = size
        for index := int(0); index < size; index++ {
            if int(newByteArray[index]) != int(oldByteArray[index]) {
                return false
            }
            stats[1]++
        }
        return true
}

