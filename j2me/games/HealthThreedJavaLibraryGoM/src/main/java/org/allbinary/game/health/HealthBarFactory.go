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

type HealthBarFactory struct {
}


func NewHealthBarFactory() *HealthBarFactory {
    this := &HealthBarFactory{}
    return this
}


// throws Exception
func (this *HealthBarFactory) GetInstance(layerInterface IAllBinaryLayer, healthInterface *Health, location int, direction int) *HealthBar {
        if OpenGLFeatureUtil_GetInstance().IsAnyThreed() {
            var featuredAnimationInterfaceFactoryInterfaceFactory *FeaturedAnimationInterfaceFactoryInterfaceFactory = FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance(); _ = featuredAnimationInterfaceFactoryInterfaceFactory
            var healthBarAnimationInterfaceFactory *HealthBarAnimationInterfaceFactory = any(featuredAnimationInterfaceFactoryInterfaceFactory.Get(Health_NAME)).(*HealthBarAnimationInterfaceFactory); _ = healthBarAnimationInterfaceFactory
            return NewHealthBar(layerInterface, healthInterface, healthBarAnimationInterfaceFactory.GetInstanceByLocation(layerInterface, location), direction)
        } else {
            return NewHealthBar(layerInterface, healthInterface, NewHealthBarTwodAnimation(layerInterface, location), direction)
        }
}

