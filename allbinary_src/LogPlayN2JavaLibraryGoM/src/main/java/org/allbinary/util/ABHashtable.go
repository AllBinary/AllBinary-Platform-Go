/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

type ABHashtable[K any, V any] struct {
    Hashtable
}



func NewABHashtable[K any, V any]() *ABHashtable[K, V] {
    this := &ABHashtable[K, V]{}
        this.Hashtable = *NewHashtable()
    return this
}


func (this *ABHashtable[K, V]) Size() int {
        return this.Hashtable.Size()
}


func (this *ABHashtable[K, V]) IsEmpty() bool {
        return this.Hashtable.IsEmpty()
}


func (this *ABHashtable[K, V]) Elements() Enumeration {
        return this.Hashtable.Elements()
}


func (this *ABHashtable[K, V]) Contains(value any) bool {
        return this.Hashtable.Contains(value)
}


func (this *ABHashtable[K, V]) ContainsValue(value any) bool {
        return this.Hashtable.ContainsValue(value)
}


func (this *ABHashtable[K, V]) ContainsKey(key any) bool {
        return this.Hashtable.ContainsKey(key)
}


func (this *ABHashtable[K, V]) Get(key any) V {
        return this.Hashtable.Get(key)
}


func (this *ABHashtable[K, V]) Clear() {
        this.Hashtable.Clear()
}


func (this *ABHashtable[K, V]) Clone() any {
        return this.Hashtable.Clone()
}


func (this *ABHashtable[K, V]) ToString() string {
        return this.Hashtable.ToString()
}

func (this *ABHashtable[K, V]) String() string {
    return this.ToString()
}


func (this *ABHashtable[K, V]) Equals(obj any) bool {
        return this.Hashtable.Equals(obj)
}


func (this *ABHashtable[K, V]) HashCode() int {
        return this.Hashtable.HashCode()
}

