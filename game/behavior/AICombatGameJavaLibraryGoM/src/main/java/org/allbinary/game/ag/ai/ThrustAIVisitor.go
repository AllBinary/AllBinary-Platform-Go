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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type ThrustAIVisitor struct {
    Visitor
    LogUtil *LogUtil
    TRUE bool
}


func (this *ThrustAIVisitor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TRUE = BooleanFactory_GetInstance().TRUE
}


func NewThrustAIVisitor() *ThrustAIVisitor {
    this := &ThrustAIVisitor{}
    this.Visitor = *NewVisitor()
    this.initFields()
    return this
}


func (this *ThrustAIVisitor) Visit(anyType any) any {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.VISIT, e)
                }
            }()
            var basicAI IBasicAI = any(anyType).(IBasicAI); _ = basicAI
            basicAI.ProcessKeyAI(Canvas_UP)
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return this.TRUE
}

