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

type InterfaceCastProxy struct {
    AnyType any
}

// Implements *InvocationHandler


func (this *InterfaceCastProxy) initFields() {
    
}


func NewInterfaceCastProxy(anyType any) *InterfaceCastProxy {
    this := &InterfaceCastProxy{}
    this.initFields()
        this.AnyType = anyType
    return this
}


func InterfaceCastProxy_NewInstance(anyType any) any {
        return Proxy_NewProxyInstance(reflect.TypeOf(anyType).GetClassLoader(), reflect.TypeOf(anyType).GetInterfaces(), NewInterfaceCastProxy(anyType))
}


// throws Throwable
func (this *InterfaceCastProxy) Invoke(proxyObject any, proxyMethod *Method, proxyArgs []any) any {
        var methodName string = proxyMethod.GetName(); _ = methodName
        var realMethod *Method = reflect.TypeOf(this.AnyType).GetMethod(methodName, proxyMethod.GetParameterTypes()); _ = realMethod
        if !realMethod.IsAccessible() {
            realMethod.SetAccessible(true)
        }
        return proxyMethod.Invoke(this.AnyType, proxyArgs)
}

