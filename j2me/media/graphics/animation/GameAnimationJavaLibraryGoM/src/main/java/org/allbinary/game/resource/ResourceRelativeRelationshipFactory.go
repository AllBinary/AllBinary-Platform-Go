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

type ResourceRelativeRelationshipFactory struct {
    BasicArrayListUtil *BasicArrayListUtil
    Hashtable *ABHashtable[any, any]
    Initialized bool
}

// Implements ResourceRelativeRelationshipFactoryInterface


func (this *ResourceRelativeRelationshipFactory) initFields() {
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    this.Hashtable = StdUtil_GetInstance().CreateHashtable()
    
}

func NewResourceRelativeRelationshipFactory() *ResourceRelativeRelationshipFactory {
    this := &ResourceRelativeRelationshipFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *ResourceRelativeRelationshipFactory) Init(level int) {
        this.SetInitialized(true)
}


// throws Exception
func (this *ResourceRelativeRelationshipFactory) GetResourceRelativeRelationshipList(resource string) IBasicArrayList {
        var listCanBeNull any = this.Hashtable.Get(resource); _ = listCanBeNull
        if listCanBeNull == nil {
            return this.BasicArrayListUtil.GetImmutableInstance()
        }
        return any(listCanBeNull).(IBasicArrayList)
}


// throws Exception
func (this *ResourceRelativeRelationshipFactory) AddResourceRelativeRelationship(resource string, hardPoint IRelativeRelationship) {
        var list IBasicArrayList = this.GetResourceRelativeRelationshipList(resource); _ = list
        if list == this.BasicArrayListUtil.GetImmutableInstance() {
            list = NewBasicArrayListD()
        }
        list.Add(hardPoint)
        this.Hashtable.Put(resource, list)
}


func (this *ResourceRelativeRelationshipFactory) IsLoadingLevel(level int) bool {
        return false
}


func (this *ResourceRelativeRelationshipFactory) IsFeature() bool {
        return false
}


func (this *ResourceRelativeRelationshipFactory) SetInitialized(initialized bool) {
        this.Initialized = initialized
}


func (this *ResourceRelativeRelationshipFactory) IsInitialized() bool {
        return this.Initialized
}

