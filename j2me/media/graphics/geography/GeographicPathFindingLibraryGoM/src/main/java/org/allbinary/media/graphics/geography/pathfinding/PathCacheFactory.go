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

type PathCacheFactory struct {
    BasicArrayListUtil *BasicArrayListUtil
    Hashtable *ABHashtable[any, any]
}

var PathCacheFactory_PathFactory *PathCacheFactory = NewPathCacheFactory()

func (this *PathCacheFactory) initFields() {
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
}


func NewPathCacheFactory() *PathCacheFactory {
    this := &PathCacheFactory{}
    this.initFields()
    return this
}


func PathCacheFactory_GetInstance() *PathCacheFactory {
        return PathCacheFactory_PathFactory
}


func (this *PathCacheFactory) GetSize() int {
        return this.Hashtable.Size()
}


func (this *PathCacheFactory) Add(pathId int, list IBasicArrayList) {
        this.Hashtable.Put(pathId, list)
}


func (this *PathCacheFactory) Remove(pathId int) {
        this.Hashtable.Remove(pathId)
}


// throws Exception
func (this *PathCacheFactory) RemoveAll() {
        this.Hashtable.Clear()
        if this.Hashtable.Size() > 0 {
            panic(errors.New(fmt.Sprint("Did not clear")))
        }
}


// throws Exception
func (this *PathCacheFactory) GetOrCreate(pathIdInteger int) IBasicArrayList {
        var listCanBeNull any = this.Hashtable.Get(pathIdInteger); _ = listCanBeNull
        if listCanBeNull == nil {
            listCanBeNull = this.BasicArrayListUtil.GetImmutableInstance()
        }
        return any(listCanBeNull).(IBasicArrayList)
}

