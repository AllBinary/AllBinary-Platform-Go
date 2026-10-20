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

//Warning you must have sql root access
type InitDb struct {
    AnyType any
    METHOD_CALL_GET_HOST_NAME string
    METHOD_CALL_ADD_USERS string
    METHOD_CALL_ADD_DATABASES string
    METHOD_CALL_ADD_TABLES string
    METHOD_CALL_TEMP_MAIN_PATH string
    METHOD_CALL_MAIN_PATH string
}


func (this *InitDb) initFields() {
    
    this.METHOD_CALL_GET_HOST_NAME = "getHostName"
    this.METHOD_CALL_ADD_USERS = "addUsers"
    this.METHOD_CALL_ADD_DATABASES = "addDatabases"
    this.METHOD_CALL_ADD_TABLES = "addTables"
    this.METHOD_CALL_TEMP_MAIN_PATH = "useTemporaryMainPath"
    this.METHOD_CALL_MAIN_PATH = "useNormalMainPath"
}


func NewInitDb() *InitDb {
    this := &InitDb{}
    this.initFields()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewInitDb(dbConnectionInfo *DbConnectionInfo) *InitDb {
    this := &InitDb{}
    this.initFields()
    return this
}


func (this *InitDb) SetHelper(anyType any) {
        this.AnyType = anyType
}


func (this *InitDb) GetHelper() any {
        return this.AnyType
}


func (this *InitDb) GetHostName() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var dynamicClass reflect.Type = reflect.TypeOf(this.AnyType); _ = dynamicClass
            var method *Method = dynamicClass.GetMethod(this.METHOD_CALL_GET_HOST_NAME, nil); _ = method
            var result string = any(method.Invoke(this.AnyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDb) AddUsers() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var dynamicClass reflect.Type = reflect.TypeOf(this.AnyType); _ = dynamicClass
            var method *Method = dynamicClass.GetMethod(this.METHOD_CALL_ADD_USERS, nil); _ = method
            var result bool = any(method.Invoke(this.AnyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result.BooleanValue(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDb) AddDatabases() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var dynamicClass reflect.Type = reflect.TypeOf(this.AnyType); _ = dynamicClass
            var method *Method = dynamicClass.GetMethod(this.METHOD_CALL_ADD_DATABASES, nil); _ = method
            var result bool = any(method.Invoke(this.AnyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result.BooleanValue(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDb) AddTables() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var dynamicClass reflect.Type = reflect.TypeOf(this.AnyType); _ = dynamicClass
            var method *Method = dynamicClass.GetMethod(this.METHOD_CALL_ADD_TABLES, nil); _ = method
            var result bool = any(method.Invoke(this.AnyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result.BooleanValue(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDb) UseTemporaryMainPath() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var dynamicClass reflect.Type = reflect.TypeOf(this.AnyType); _ = dynamicClass
            var method *Method = dynamicClass.GetMethod(this.METHOD_CALL_TEMP_MAIN_PATH, nil); _ = method
            var result bool = any(method.Invoke(this.AnyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result.BooleanValue(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitDb) UseNormalMainPath() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var dynamicClass reflect.Type = reflect.TypeOf(this.AnyType); _ = dynamicClass
            var method *Method = dynamicClass.GetMethod(this.METHOD_CALL_MAIN_PATH, nil); _ = method
            var result bool = any(method.Invoke(this.AnyType, nil)).(bool); _ = result
            tryControl1, tryResult1 = 1, result.BooleanValue(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

