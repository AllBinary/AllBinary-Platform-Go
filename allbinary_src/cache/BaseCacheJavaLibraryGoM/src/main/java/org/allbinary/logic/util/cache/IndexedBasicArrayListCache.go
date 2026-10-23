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

type IndexedBasicArrayListCache struct {
    BaseBasicArrayListCache
    CommonStrings *CommonStrings
    Size int
    Index int
}

// Implements CacheInterface

// The Java class IndexedBasicArrayListCache has subclasses so its variables are this interface
type IIndexedBasicArrayListCache interface {
    IBaseBasicArrayListCache
    IndexedBasicArrayListCache_() *IndexedBasicArrayListCache
    Init(size int)
    GetAt(index int) any
    AddDefault()
    AddObject(anyType any)
    Add(cacheableInterface CacheableInterface)
    AddArray(cacheableInterfaces []CacheableInterface)
    Clear()
}

func (this *IndexedBasicArrayListCache) IndexedBasicArrayListCache_() *IndexedBasicArrayListCache {
    return this
}


func (this *IndexedBasicArrayListCache) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.Size = 0
    this.Index = 0
}


// throws Exception
func IndexedBasicArrayListCache_CreateCache(size int) IIndexedBasicArrayListCache {
        var indexedBasicArrayListCache IIndexedBasicArrayListCache = NewIndexedBasicArrayListCache(); _ = indexedBasicArrayListCache
        indexedBasicArrayListCache.Init(size)
        return indexedBasicArrayListCache
}


func NewIndexedBasicArrayListCache() *IndexedBasicArrayListCache {
    this := &IndexedBasicArrayListCache{}
    this.BaseBasicArrayListCache = *NewBaseBasicArrayListCache()
    this.initFields()
    return this
}


// throws Exception
func (this *IndexedBasicArrayListCache) Init(size int) {
        if this.Size < size {
            this.List.EnsureCapacity(size)
            for index := int(size - 1); index >= this.Size; index-- {
                this.AddDefault()
            }
            this.Size = size
        }
}


func (this *IndexedBasicArrayListCache) GetAt(index int) any {
        return this.List.BasicArrayList_().ObjectArray[index]
}


// throws Exception
func (this *IndexedBasicArrayListCache) AddDefault() {
}


// throws Exception
func (this *IndexedBasicArrayListCache) AddObject(anyType any) {
        this.List.Add(anyType)
}


// throws Exception
func (this *IndexedBasicArrayListCache) Add(cacheableInterface CacheableInterface) {
        this.List.Add(cacheableInterface)
}


// throws Exception
func (this *IndexedBasicArrayListCache) AddArray(cacheableInterfaces []CacheableInterface) {
        panic(errors.New(fmt.Sprint("No Imple")))
}


func (this *IndexedBasicArrayListCache) Clear() {
        this.Index = 0
}

