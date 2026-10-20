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

type AbstractArrayListPool struct {
    CommonStrings *CommonStrings
    Buffers IBasicArrayList
    CacheableInterfaceFactoryInterface CacheableInterfaceFactoryInterface
}

// Implements PoolInterface

// The Java class AbstractArrayListPool has subclasses so its variables are this interface
type IAbstractArrayListPool interface {
    AbstractArrayListPool_() *AbstractArrayListPool
    Clear()
    Remove(key any) CacheableInterface
    Add(cacheableInterface CacheableInterface)
    AddArray(cacheableInterfaces []CacheableInterface)
}

func (this *AbstractArrayListPool) AbstractArrayListPool_() *AbstractArrayListPool {
    return this
}


func (this *AbstractArrayListPool) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.Buffers = NewBasicArrayListD()
    
}


func NewAbstractArrayListPool(cacheableInterfaceFactoryInterface CacheableInterfaceFactoryInterface) *AbstractArrayListPool {
    this := &AbstractArrayListPool{}
    this.initFields()
        this.CacheableInterfaceFactoryInterface = cacheableInterfaceFactoryInterface
    return this
}


// throws Exception
func (this *AbstractArrayListPool) Clear() {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *AbstractArrayListPool) Remove(key any) CacheableInterface {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *AbstractArrayListPool) Add(cacheableInterface CacheableInterface) {
        this.Buffers.Add(cacheableInterface)
}


// throws Exception
func (this *AbstractArrayListPool) AddArray(cacheableInterfaces []CacheableInterface) {
        for index := int(0); index < len(cacheableInterfaces); index++ {
            this.Buffers.Add(cacheableInterfaces[index])
        }
}

