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

type KeylessBasicArrayListCache struct {
    IndexedBasicArrayListCache
    LogUtil *LogUtil
    BasicArrayListUtil *BasicArrayListUtil
}

// The Java class KeylessBasicArrayListCache has subclasses so its variables are this interface
type IKeylessBasicArrayListCache interface {
    IIndexedBasicArrayListCache
    KeylessBasicArrayListCache_() *KeylessBasicArrayListCache
    Get() IBasicArrayList
    Clear()
    Log() string
}

func (this *KeylessBasicArrayListCache) KeylessBasicArrayListCache_() *KeylessBasicArrayListCache {
    return this
}


func (this *KeylessBasicArrayListCache) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewKeylessBasicArrayListCache() *KeylessBasicArrayListCache {
    this := &KeylessBasicArrayListCache{}
    this.IndexedBasicArrayListCache = *NewIndexedBasicArrayListCache()
    this.initFields()
    return this
}


func (this *KeylessBasicArrayListCache) Get() IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET, e)
            tryControl1, tryResult1 = 1, this.BasicArrayListUtil.GetImmutableInstance(); return
                }
            }()
            if this.Index >= this.List.Size() {
                this.AddDefault()
            }
            var list IBasicArrayList = any(this.GetAt(Java_PostIncrement(&this.Index))).(IBasicArrayList); _ = list
            tryControl1, tryResult1 = 1, list; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *KeylessBasicArrayListCache) Clear() {
        this.IndexedBasicArrayListCache.Clear()
        var basicArrayList IBasicArrayList; _ = basicArrayList
        for index := int(this.List.Size() - 1); index >= 0; index-- {
            basicArrayList = (any(this.List.BasicArrayList_().ObjectArray[index]).(IBasicArrayList))
            basicArrayList.Clear()
        }
}


func (this *KeylessBasicArrayListCache) Log() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("S: ")
        stringBuffer.Appendint(this.List.Size())
        var S_LABEL string = " s: "; _ = S_LABEL
        var basicArrayList IBasicArrayList; _ = basicArrayList
        for index := int(this.List.Size() - 1); index >= 0; index-- {
            stringBuffer.Append(S_LABEL)
            basicArrayList = (any(this.List.BasicArrayList_().ObjectArray[index]).(IBasicArrayList))
            stringBuffer.Appendint(basicArrayList.Size())
        }
        return fmt.Sprint(stringBuffer)
}

