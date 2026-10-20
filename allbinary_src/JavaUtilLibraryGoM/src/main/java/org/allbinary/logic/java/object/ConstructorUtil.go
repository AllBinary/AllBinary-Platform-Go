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

type ConstructorUtil struct {
}



func NewConstructorUtil() *ConstructorUtil {
    this := &ConstructorUtil{}
    return this
}


func ConstructorUtil_ViewAll(myClass reflect.Type, lineBreak string) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var constructor []*Constructor = myClass.GetConstructors(); _ = constructor
        stringBuffer.Append("Constructors: ")
        stringBuffer.Append(lineBreak)
        for index := int(0); index < len(constructor); index++ {
            stringBuffer.Append(ConstructorUtil_View(constructor[index], lineBreak))
        }
        return fmt.Sprint(stringBuffer)
}


func ConstructorUtil_View(constructor *Constructor, lineBreak string) string {
        if constructor != nil {
            var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(constructor.GetName())
            var classes []reflect.Type = constructor.GetParameterTypes(); _ = classes
            for index := int(0); index < len(classes); index++ {
                stringBuffer.Append(commonSeps.SPACE)
                stringBuffer.Append(classes[index].GetName())
                stringBuffer.Append(commonSeps.COLON)
                stringBuffer.Appendint(index)
            }
            stringBuffer.Append(lineBreak)
            return fmt.Sprint(stringBuffer)
        } else {
            return "Constructor Null"
        }
}

