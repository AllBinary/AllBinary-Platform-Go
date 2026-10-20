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

type ServiceReferenceMap struct {
    ServiceReference ServiceReference
}

// Implements Map


func (this *ServiceReferenceMap) initFields() {
    
}


func NewServiceReferenceMap(serviceReference ServiceReference) *ServiceReferenceMap {
    this := &ServiceReferenceMap{}
    this.initFields()
        this.ServiceReference = serviceReference
    return this
}


func (this *ServiceReferenceMap) Size() int {
        return -1
}


func (this *ServiceReferenceMap) IsEmpty() bool {
        return false
}


func (this *ServiceReferenceMap) ContainsKey(key any) bool {
        return false
}


func (this *ServiceReferenceMap) ContainsValue(value any) bool {
        return false
}


func (this *ServiceReferenceMap) Get(key any) any {
        return this.ServiceReference.GetProperty(any(key).(string))
}


func (this *ServiceReferenceMap) Put(key any, value any) any {
        return nil
}


func (this *ServiceReferenceMap) Remove(key any) any {
        return nil
}


func (this *ServiceReferenceMap) PutAll(m Map) {
}


func (this *ServiceReferenceMap) Clear() {
}


func (this *ServiceReferenceMap) KeySet() Set {
        var keyStringArray []string = this.ServiceReference.GetPropertyKeys(); _ = keyStringArray
        return nil
}


func (this *ServiceReferenceMap) Values() Collection[any] {
        return nil
}


func (this *ServiceReferenceMap) EntrySet() Set {
        return nil
}


func ServiceReferenceMap_ToString(map_ Map) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        return fmt.Sprint(stringBuffer)
}

