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

type InterfaceUtil struct {
}



func NewInterfaceUtil() *InterfaceUtil {
    this := &InterfaceUtil{}
    return this
}


func InterfaceUtil_ViewAll(myClass reflect.Type, lineBreak string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var interfaces []reflect.Type = myClass.GetInterfaces(); _ = interfaces
        stringBuffer.Append(lineBreak)
        stringBuffer.Append("Interfaces: ")
        stringBuffer.Append(lineBreak)
        for index := int(0); index < len(interfaces); index++ {
            stringBuffer.Append("Interface: ")
            stringBuffer.Append(interfaces[index].GetName())
            stringBuffer.Append(lineBreak)
        }
        return fmt.Sprint(stringBuffer)
}


func InterfaceUtil_GetClass(interfaceName string, interfaces []reflect.Type) reflect.Type {
        if interfaceName != "" && interfaces != nil && len(interfaces) > 0 {
            for index := int(0); index < len(interfaces); index++ {
                if strings.Index(interfaces[index].GetName(), string(interfaceName)) >= 0 {
                    return interfaces[index]
                }
            }
        }
        return nil
}


func InterfaceUtil_IsImplemented(clazz reflect.Type, anyType any) bool {
        var classes []reflect.Type = reflect.TypeOf(anyType).GetInterfaces(); _ = classes
        for index := int(0); index < len(classes); index++ {
            if strings.Compare(clazz.String(), classes[index].GetName()) == 0 {
                return true
            }
        }
        return false
}


// Java overloaded method: Go needs a different name
func InterfaceUtil_IsImplemented(className string, anyType any) bool {
        var classes []reflect.Type = reflect.TypeOf(anyType).GetInterfaces(); _ = classes
        for index := int(0); index < len(classes); index++ {
            if strings.Index(classes[index].GetName(), string(className)) == 0 {
                return true
            }
        }
        return false
}


func InterfaceUtil_IsImplementedView(clazz reflect.Type, anyType any) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var classes []reflect.Type = reflect.TypeOf(anyType).GetInterfaces(); _ = classes
        stringBuffer.Append("isImplementedView: \n")
        for index := int(0); index < len(classes); index++ {
            stringBuffer.Append(clazz.String())
            stringBuffer.Append(" should be = ")
            stringBuffer.Append(classes[index].GetName())
            stringBuffer.Append(classes[index].GetName())
            stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        }
        return fmt.Sprint(stringBuffer)
}

