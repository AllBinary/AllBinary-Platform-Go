/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type CapitalEventHandlerFactory struct {
}

var CapitalEventHandlerFactory_CapitalEventHandlerArray []*CapitalEventHandler

func NewCapitalEventHandlerFactory() *CapitalEventHandlerFactory {
    this := &CapitalEventHandlerFactory{}
    return this
}


func init() {
        CapitalEventHandlerFactory_CapitalEventHandlerArray = make([]*CapitalEventHandler, 10)
        for index := int(len(CapitalEventHandlerFactory_CapitalEventHandlerArray) - 1); index >= 0; index-- {
            CapitalEventHandlerFactory_CapitalEventHandlerArray[index] = NewCapitalEventHandler()
        }
}


func CapitalEventHandlerFactory_GetInstance(groupInterface *Group) *CapitalEventHandler {
        return CapitalEventHandlerFactory_CapitalEventHandlerArray[int(groupInterface.GetGroupId())]
}


func CapitalEventHandlerFactory_RemoveAll() {
        for index := int(len(CapitalEventHandlerFactory_CapitalEventHandlerArray) - 1); index >= 0; index-- {
            CapitalEventHandlerFactory_CapitalEventHandlerArray[index].RemoveAllListeners()
        }
}

