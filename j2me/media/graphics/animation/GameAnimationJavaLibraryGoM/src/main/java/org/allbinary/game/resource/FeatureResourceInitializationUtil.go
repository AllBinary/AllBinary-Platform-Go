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

type FeatureResourceInitializationUtil struct {
    LogUtil *LogUtil
}

var FeatureResourceInitializationUtil_Instance *FeatureResourceInitializationUtil = NewFeatureResourceInitializationUtil()

func (this *FeatureResourceInitializationUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewFeatureResourceInitializationUtil() *FeatureResourceInitializationUtil {
    this := &FeatureResourceInitializationUtil{}
    this.initFields()
    return this
}


func FeatureResourceInitializationUtil_GetInstance() *FeatureResourceInitializationUtil {
        return FeatureResourceInitializationUtil_Instance
}


// throws Exception
func (this *FeatureResourceInitializationUtil) Init(level int) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(commonStrings.START, this, commonStrings.INIT)
        var featuredResourceFactory IFeaturedResourceFactory = FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance(); _ = featuredResourceFactory
        featuredResourceFactory.Init(level)
        FeaturedResourceRelativeRelationshipFactory_GetInstance().Init(level)
}

