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

type ABHashMap[K any, V any] struct {
    HashMap
}


func NewABHashMap[K any, V any]() *ABHashMap[K, V] {
    this := &ABHashMap[K, V]{}
    this.HashMap = *NewHashMap()
    return this
}


func (this *ABHashMap[K, V]) Size() int {
        return this.HashMap.Size()
}


func (this *ABHashMap[K, V]) IsEmpty() bool {
        return this.HashMap.IsEmpty()
}


func (this *ABHashMap[K, V]) Get(key any) IV {
        return this.HashMap.Get(key)
}


func (this *ABHashMap[K, V]) ContainsKey(key any) bool {
        return this.HashMap.ContainsKey(key)
}


func (this *ABHashMap[K, V]) Put(key K, value IV) IV {
        return this.HashMap.Put(key, value)
}


func (this *ABHashMap[K, V]) PutAll(map_ Map[any, any]) {
        this.HashMap.PutAll(map_)
}


func (this *ABHashMap[K, V]) Remove(key any) IV {
        return this.HashMap.Remove(key)
}


func (this *ABHashMap[K, V]) Clear() {
        this.HashMap.Clear()
}


func (this *ABHashMap[K, V]) ContainsValue(value any) bool {
        return this.HashMap.ContainsValue(value)
}


func (this *ABHashMap[K, V]) KeySet() Set[K] {
        return this.HashMap.KeySet()
}


func (this *ABHashMap[K, V]) EntrySet() Set[*Entry[K, IV]] {
        return this.HashMap.EntrySet()
}


func (this *ABHashMap[K, V]) GetOrDefault(key any, defaultValue IV) IV {
        return this.HashMap.GetOrDefault(key, defaultValue)
}


func (this *ABHashMap[K, V]) PutIfAbsent(key K, value IV) IV {
        return this.HashMap.PutIfAbsent(key, value)
}


func (this *ABHashMap[K, V]) Clone() any {
        return this.HashMap.Clone()
}

