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

type InGameFeatureChoiceGroups struct {
    Hashtable *ABHashtable[any, any]
}

var InGameFeatureChoiceGroups_GameFeatureChoiceGroupsExclusive *InGameFeatureChoiceGroups = NewInGameFeatureChoiceGroups()
var InGameFeatureChoiceGroups_GameFeatureChoiceGroupsMultiple *InGameFeatureChoiceGroups = NewInGameFeatureChoiceGroups()

func (this *InGameFeatureChoiceGroups) initFields() {
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
}


func NewInGameFeatureChoiceGroups() *InGameFeatureChoiceGroups {
    this := &InGameFeatureChoiceGroups{}
    this.initFields()
    return this
}


func InGameFeatureChoiceGroups_GetExclusiveInstance() *InGameFeatureChoiceGroups {
        return InGameFeatureChoiceGroups_GameFeatureChoiceGroupsExclusive
}


func InGameFeatureChoiceGroups_GetMultipleInstance() *InGameFeatureChoiceGroups {
        return InGameFeatureChoiceGroups_GameFeatureChoiceGroupsMultiple
}


func (this *InGameFeatureChoiceGroups) Get() *ABHashtable[any, any] {
        return this.Hashtable
}


func (this *InGameFeatureChoiceGroups) Add(name string, list IBasicArrayList) {
        this.Hashtable.Put(name, list)
}

