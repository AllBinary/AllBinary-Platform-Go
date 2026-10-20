/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type DecalAnimationInterfaceFactory struct {
}

var DecalAnimationInterfaceFactory_Instance *DecalAnimationInterfaceFactory = NewDecalAnimationInterfaceFactory()

func NewDecalAnimationInterfaceFactory() *DecalAnimationInterfaceFactory {
    this := &DecalAnimationInterfaceFactory{}
    return this
}


func DecalAnimationInterfaceFactory_GetInstance() *DecalAnimationInterfaceFactory {
        return DecalAnimationInterfaceFactory_Instance
}


// throws Exception
func (this *DecalAnimationInterfaceFactory) GetAnimationFactory(basicGameResources IBasicGameResources, basicColor *BasicColor) AnimationInterfaceFactoryInterface {
        var animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface; _ = animationInterfaceFactoryInterface
        if basicColor == BasicColorFactory_GetInstance().RED {
            animationInterfaceFactoryInterface = FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(basicGameResources.BasicGameResources_().RESOURCE)
        } else {
            animationInterfaceFactoryInterface = NullRotationAnimationFactory_GetFactoryInstance()
        }
        return animationInterfaceFactoryInterface
}

