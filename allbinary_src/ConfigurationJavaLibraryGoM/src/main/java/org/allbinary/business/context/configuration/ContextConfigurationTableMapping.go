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

type ContextConfigurationTableMapping struct {
    ContextConfigurationInterface ContextConfigurationInterface
}

// Implements TableMappingInterface


func (this *ContextConfigurationTableMapping) initFields() {
    
}


func NewContextConfigurationTableMapping(contextConfigurationInterface ContextConfigurationInterface) *ContextConfigurationTableMapping {
    this := &ContextConfigurationTableMapping{}
    this.initFields()
        this.ContextConfigurationInterface = contextConfigurationInterface
    return this
}


func (this *ContextConfigurationTableMapping) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        return hashMap
}


// throws Exception
func (this *ContextConfigurationTableMapping) GetKey() any {
        return ContextConfigurationData_GetInstance().NAME
}


// throws Exception
func (this *ContextConfigurationTableMapping) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        return vector
}

