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

type ClassUtil struct {
}



func NewClassUtil() *ClassUtil {
    this := &ClassUtil{}
    return this
}


func ClassUtil_ViewAll(anyType any, lineBreak string) string {
        if anyType != nil {
            return ClassUtil_ViewAll(reflect.TypeOf(anyType), lineBreak)
        } else {
            return "Object Null"
        }
}


// Java overloaded method: Go needs a different name
func ClassUtil_ViewAll(myClass reflect.Type, lineBreak string) string {
        if myClass != nil {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var classes []reflect.Type = myClass.GetClasses(); _ = classes
            stringBuffer.Append(lineBreak)
            stringBuffer.Append("Class: ")
            stringBuffer.Append(myClass.String())
            stringBuffer.Append(" uses ")
            stringBuffer.Appendint(len(classes))
            stringBuffer.Append(" other classes")
            stringBuffer.Append(lineBreak)
            stringBuffer.Append(lineBreak)
            stringBuffer.Append("Classes: ")
            stringBuffer.Append(lineBreak)
            for index := int(0); index < len(classes); index++ {
                stringBuffer.Append("Class: ")
                stringBuffer.Append(classes[index].GetName())
                stringBuffer.Append(lineBreak)
            }
            stringBuffer.Append(InterfaceUtil_ViewAll(myClass, lineBreak))
            stringBuffer.Append(lineBreak)
            stringBuffer.Append(ConstructorUtil_ViewAll(myClass, lineBreak))
            stringBuffer.Append(lineBreak)
            stringBuffer.Append(MethodUtil_ViewAll(myClass, lineBreak))
            return fmt.Sprint(stringBuffer)
        } else {
            return "Class Null"
        }
}

