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

type HashtableUtil struct {
    EnumerationUtil *EnumerationUtil
}

var HashtableUtil_Instance *HashtableUtil = NewHashtableUtil()

func (this *HashtableUtil) initFields() {
    this.EnumerationUtil = EnumerationUtil_GetInstance()
}

func NewHashtableUtil() *HashtableUtil {
    this := &HashtableUtil{}
    this.initFields()
    return this
}


func HashtableUtil_GetInstance() *HashtableUtil {
        return HashtableUtil_Instance
}


func (this *HashtableUtil) PutAll(fromHashtable *ABHashtable[any, any], hashtable *ABHashtable[any, any]) {
        var enumeration Enumeration = fromHashtable.Keys(); _ = enumeration
        var keyCanBeNull any; _ = keyCanBeNull
        var valueCanBeNull any; _ = valueCanBeNull
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            keyCanBeNull = this.EnumerationUtil.NextElement(enumeration)
            valueCanBeNull = hashtable.Get(keyCanBeNull)
            if valueCanBeNull == nil {
            } else {
                hashtable.Put(keyCanBeNull, valueCanBeNull)
            }
        }
}


func (this *HashtableUtil) GetKeysAsArray(hashtable *ABHashtable[any, any]) []any {
        var objectArray []any = make([]any, hashtable.Size()); _ = objectArray
        var index int = 0; _ = index
        var enumeration Enumeration = hashtable.Keys(); _ = enumeration
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            objectArray[Java_PostIncrement(&index)] = this.EnumerationUtil.NextElement(enumeration)
        }
        return objectArray
}


func (this *HashtableUtil) GetKeysAsList(hashtable *ABHashtable[any, any]) IBasicArrayList {
        var list IBasicArrayList = NewBasicArrayListS(hashtable.Size()); _ = list
        var enumeration Enumeration = hashtable.Keys(); _ = enumeration
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            list.Add(this.EnumerationUtil.NextElement(enumeration))
        }
        return list
}

