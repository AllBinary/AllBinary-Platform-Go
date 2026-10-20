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

type BasicArrayListPool struct {
    AbstractArrayListPool
}



func NewBasicArrayListPool(cacheableInterfaceFactoryInterface CacheableInterfaceFactoryInterface) *BasicArrayListPool {
    this := &BasicArrayListPool{}
        this.AbstractArrayListPool = *NewAbstractArrayListPool(cacheableInterfaceFactoryInterface)
    return this
}


func (this *BasicArrayListPool) Clear() {
        this.Buffers.Clear()
}


// throws Exception
func (this *BasicArrayListPool) Remove(key any) CacheableInterface {
        var size int = this.Buffers.Size(); _ = size
        if size > 0 {
            return any(this.Buffers.RemoveAt(size - 1)).(CacheableInterface)
        } else {
            return this.CacheableInterfaceFactoryInterface.GetInstance(key)
        }
}


// throws Exception
func (this *BasicArrayListPool) AddAll(usedList IBasicArrayList) {
        for index := int(usedList.Size()); Java_PreDecrement(&index) >= 0;  {
            var anyType CacheableInterface = any(usedList.BasicArrayList_().ObjectArray[index]).(CacheableInterface); _ = anyType
            this.Add(anyType)
        }
        usedList.Clear()
}


// throws Exception
func (this *BasicArrayListPool) ReleaseUsedBackToPool(usedList IBasicArrayList) {
        for index := int(usedList.Size()); Java_PreDecrement(&index) >= 0;  {
            var list *BasicArrayListCacheable = any(usedList.BasicArrayList_().ObjectArray[index]).(*BasicArrayListCacheable); _ = list
            if list.Size() == 0 {
                usedList.RemoveAt(index)
                this.Add(list)
            }
        }
}

