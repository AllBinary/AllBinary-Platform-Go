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

type MethodUtil struct {
}



func NewMethodUtil() *MethodUtil {
    this := &MethodUtil{}
    return this
}


func MethodUtil_ViewAll(myClass reflect.Type, lineBreak string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var method []*Method = myClass.GetMethods(); _ = method
        stringBuffer.Append("Methods: ")
        stringBuffer.Append(lineBreak)
        for index := int(0); index < len(method); index++ {
            stringBuffer.Append(method[index].GetReturnType().GetName())
            stringBuffer.Append(" ")
            stringBuffer.Append(method[index].GetName())
            stringBuffer.Append("(")
            ParamsUtil_ViewParams(method[index].GetParameterTypes())
            stringBuffer.Append(")")
            stringBuffer.Append(lineBreak)
        }
        return fmt.Sprint(stringBuffer)
}

