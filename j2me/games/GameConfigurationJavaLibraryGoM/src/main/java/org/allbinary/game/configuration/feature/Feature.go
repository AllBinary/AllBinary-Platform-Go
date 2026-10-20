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

type Feature struct {
    Name string
}

// The Java class Feature has subclasses so its variables are this interface
type IFeature interface {
    Feature_() *Feature
    ToString() string
    GetName() string
}

func (this *Feature) Feature_() *Feature {
    return this
}

var Feature_Hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable()

func (this *Feature) initFields() {
    
}


func NewFeature(name string) *Feature {
    this := &Feature{}
    this.initFields()
        this.Name = name
        Feature_Add(name, this)
        ChangedGameFeatureListener_GetInstance().Add(this)
    return this
}


func Feature_GetInstance(name string) IFeature {
        return any(Feature_Hashtable.Get(name)).(IFeature)
}


func Feature_Add(name string, gameFeature IFeature) {
        Feature_Hashtable.Put(name, gameFeature)
}


func (this *Feature) ToString() string {
        return this.GetName()
}

func (this *Feature) String() string {
    return this.ToString()
}


func (this *Feature) GetName() string {
        return this.Name
}

