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

type OnDemandResources struct {
    LogUtil *LogUtil
}

var OnDemandResources_NULL_ON_DEMAND_RESOURCES *OnDemandResources = NewOnDemandResources()

func (this *OnDemandResources) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewOnDemandResources() *OnDemandResources {
    this := &OnDemandResources{}
    this.initFields()
    return this
}


// throws Exception
func (this *OnDemandResources) WaitFor() {
}


// throws Exception
func (this *OnDemandResources) Init() {
}


// throws Exception
func (this *OnDemandResources) InitAt(imageCache IImageCache, resourceAnimationInterfaceFactoryInterfaceFactory IBaseResourceAnimationInterfaceFactoryInterfaceFactory, portion int, loadingString string, index int) int {
        return 0
}

