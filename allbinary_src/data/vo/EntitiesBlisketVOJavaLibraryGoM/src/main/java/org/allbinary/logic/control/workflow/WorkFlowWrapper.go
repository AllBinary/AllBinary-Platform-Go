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

type WorkFlowWrapper struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AnyType any
}

// Implements WorkFlowInterface


func (this *WorkFlowWrapper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewWorkFlowWrapper(anyType any) *WorkFlowWrapper {
    this := &WorkFlowWrapper{}
    this.initFields()
        this.AnyType = anyType
    return this
}


func (this *WorkFlowWrapper) GetWorkFlowObject() any {
        return this.AnyType
}


// throws Exception
func (this *WorkFlowWrapper) GetName() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Reflection Exception"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getName()", e)
            }
            panic(errors.New(fmt.Sprint(error_)))
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.AnyType); _ = helperClass
            var method *Method = helperClass.GetMethod("getName", nil); _ = method
            var result string = any(method.Invoke(this.AnyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *WorkFlowWrapper) GetStoreName() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Reflection Exception"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getStoreName()", e)
            }
            panic(errors.New(fmt.Sprint(error_)))
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.AnyType); _ = helperClass
            var method *Method = helperClass.GetMethod("getStoreName", nil); _ = method
            var result string = any(method.Invoke(this.AnyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *WorkFlowWrapper) Process() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Reflection Exception"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
            }
            panic(errors.New(fmt.Sprint(error_)))
                }
            }()
            var helperClass reflect.Type = reflect.TypeOf(this.AnyType); _ = helperClass
            var method *Method = helperClass.GetMethod(this.CommonStrings.PROCESS, nil); _ = method
            var result int = any(method.Invoke(this.AnyType, nil)).(int); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

