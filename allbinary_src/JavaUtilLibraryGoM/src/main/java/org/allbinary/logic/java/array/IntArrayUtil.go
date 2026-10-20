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

type IntArrayUtil struct {
}



func NewIntArrayUtil() *IntArrayUtil {
    this := &IntArrayUtil{}
    return this
}


func IntArrayUtil_ToString(array []int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index := int(0); index < len(array); index++ {
            stringBuffer.Appendint(array[index])
            if index < len(array) {
                stringBuffer.Append(", ")
            }
        }
        return fmt.Sprint(stringBuffer)
}


// Java overloaded method: Go needs a different name
func IntArrayUtil_ToString(array [][]int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index := int(0); index < len(array); index++ {
            stringBuffer.Append("{")
            for index2 := int(0); index2 < len(array[index]); index2++ {
                stringBuffer.Appendint(int(array[index][index2]))
                if index2 < len(array[index]) - 1 {
                    stringBuffer.Append(", ")
                }
            }
            stringBuffer.Append("}")
            if index < len(array) - 1 {
                stringBuffer.Append(", \n")
            }
        }
        return fmt.Sprint(stringBuffer)
}


// Java overloaded method: Go needs a different name
func IntArrayUtil_ToString(array [][][]int) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        for index := int(0); index < len(array); index++ {
            stringBuffer.Append("{")
            for index2 := int(0); index2 < len(array[index]); index2++ {
                stringBuffer.Append("{")
                for index3 := int(0); index3 < len(array[index][index2]); index3++ {
                    stringBuffer.Appendint(int(array[index][index2][index3]))
                    if index3 < len(array[index][index2]) - 1 {
                        stringBuffer.Append(", ")
                    }
                }
                stringBuffer.Append("}")
                if index2 < len(array[index]) - 1 {
                    stringBuffer.Append(", ")
                }
            }
            stringBuffer.Append("}")
            if index < len(array) - 1 {
                stringBuffer.Append(", \n")
            }
        }
        return fmt.Sprint(stringBuffer)
}

