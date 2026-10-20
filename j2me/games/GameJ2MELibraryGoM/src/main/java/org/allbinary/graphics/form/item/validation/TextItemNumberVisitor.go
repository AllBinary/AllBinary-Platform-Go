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

type TextItemNumberVisitor struct {
    Visitor
    LogUtil *LogUtil
}


func (this *TextItemNumberVisitor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewTextItemNumberVisitor() *TextItemNumberVisitor {
    this := &TextItemNumberVisitor{}
    this.Visitor = *NewVisitor()
    this.initFields()
    return this
}


func (this *TextItemNumberVisitor) Visit(anyType any) any {
        var value string = any(anyType).(string); _ = value
        if len(value) < 2 {
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            if stringValidationUtil.IsNumber(value) {
                return BooleanFactory_GetInstance().TRUE
            } else {
                if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                    this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.VISIT, e)
                        }
                    }()
                    TextNotificationUtil_GetInstance().FireError("Numbers Only")
                    return
                }(); tryControl1 == 1 {
                    return tryResult1
                }
                return BooleanFactory_GetInstance().FALSE
            }
        } else {
            return BooleanFactory_GetInstance().FALSE
        }
}

