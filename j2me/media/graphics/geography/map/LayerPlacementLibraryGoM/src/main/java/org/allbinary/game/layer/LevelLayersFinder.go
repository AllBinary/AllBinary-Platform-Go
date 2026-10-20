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

type LevelLayersFinder struct {
    EnumerationUtil *EnumerationUtil
}

var LevelLayersFinder_Instance *LevelLayersFinder = NewLevelLayersFinder()

func (this *LevelLayersFinder) initFields() {
    this.EnumerationUtil = EnumerationUtil_GetInstance()
}

func NewLevelLayersFinder() *LevelLayersFinder {
    this := &LevelLayersFinder{}
    this.initFields()
    return this
}


func LevelLayersFinder_GetInstance() *LevelLayersFinder {
        return LevelLayersFinder_Instance
}


func (this *LevelLayersFinder) Get(hashtable *ABHashtable[any, any]) IBasicArrayList {
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var list IBasicArrayList = NewBasicArrayListD(); _ = list
        var enumeration Enumeration = hashtable.Keys(); _ = enumeration
        var layerHashtableCanBeNull *ABHashtable[any, any]; _ = layerHashtableCanBeNull
        var integerCanBeNull int; _ = integerCanBeNull
        var cachedIntegerCanBeNull int; _ = cachedIntegerCanBeNull
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            layerHashtableCanBeNull = any(hashtable.Get(any(this.EnumerationUtil.NextElement(enumeration)).(any))).(*ABHashtable[any, any])
            integerCanBeNull = any(layerHashtableCanBeNull.Get(any(Layer_ID).(any))).(int)
            cachedIntegerCanBeNull = smallIntegerSingletonFactory.GetAt(int(integerCanBeNull))
            if !list.Contains(cachedIntegerCanBeNull) {
                list.Add(cachedIntegerCanBeNull)
            }
        }
        return list
}

