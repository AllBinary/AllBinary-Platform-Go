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

type MainFeatureFactory struct {
    FULL_SCREEN IFeature
    TITLE_BAR IFeature
    STATIC IFeature
    LOAD_ONDEMAND IFeature
    LOAD_ALL IFeature
}

var MainFeatureFactory_Instance *MainFeatureFactory = NewMainFeatureFactory()

func (this *MainFeatureFactory) initFields() {
    this.FULL_SCREEN = NewFeature("Full Screen")
    this.TITLE_BAR = NewFeature("Title Bar")
    this.STATIC = NewFeature("Static Dependent Feature")
    this.LOAD_ONDEMAND = NewFeature("Load OnDemand")
    this.LOAD_ALL = NewFeature("Load All")
}

func NewMainFeatureFactory() *MainFeatureFactory {
    this := &MainFeatureFactory{}
    this.initFields()
    return this
}


func MainFeatureFactory_GetInstance() *MainFeatureFactory {
        return MainFeatureFactory_Instance
}

