/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type ThreedObjResources struct {
    Hashtable *ABHashtable[any, any]
}

var ThreedObjResources_Instance *ThreedObjResources = NewThreedObjResources()

func (this *ThreedObjResources) initFields() {
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
}

func NewThreedObjResources() *ThreedObjResources {
    this := &ThreedObjResources{}
    this.initFields()
    return this
}


func ThreedObjResources_GetInstance() *ThreedObjResources {
        return ThreedObjResources_Instance
}


func (this *ThreedObjResources) Get(resourceName string) string {
        var resourceName2 string = any(this.Hashtable.Get(resourceName)).(string); _ = resourceName2
        if resourceName2 != "" {
            return resourceName2
        }
        return resourceName
}

