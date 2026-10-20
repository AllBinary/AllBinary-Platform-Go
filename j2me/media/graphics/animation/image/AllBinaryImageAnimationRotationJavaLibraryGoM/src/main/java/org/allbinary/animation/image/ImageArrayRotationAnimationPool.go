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

type ImageArrayRotationAnimationPool struct {
    BasicArrayListResetablePool
}

var ImageArrayRotationAnimationPool_IMAGE_ROTATION_ANIMATION_POOL PoolInterface = NewImageArrayRotationAnimationPool(NewImageArrayRotationAnimationCacheableInterfaceFactory())


func NewImageArrayRotationAnimationPool(cacheableInterfaceFactoryInterface CacheableInterfaceFactoryInterface) *ImageArrayRotationAnimationPool {
    this := &ImageArrayRotationAnimationPool{}
        this.BasicArrayListResetablePool = *NewBasicArrayListResetablePool(cacheableInterfaceFactoryInterface)
    return this
}


func ImageArrayRotationAnimationPool_GetInstance() PoolInterface {
        return ImageArrayRotationAnimationPool_IMAGE_ROTATION_ANIMATION_POOL
}


// throws Exception
func (this *ImageArrayRotationAnimationPool) Remove(key any) CacheableInterface {
        var cacheableInterface CacheableInterface = this.BasicArrayListResetablePool.Remove(key); _ = cacheableInterface
        var allBinaryImageRotationAnimationInfo *ImageArrayRotationAnimationInfo = any(key).(*ImageArrayRotationAnimationInfo); _ = allBinaryImageRotationAnimationInfo
        var imageArrayRotationAnimationCacheable *ImageArrayRotationAnimationCacheable = (any(cacheableInterface).(*ImageArrayRotationAnimationCacheable)); _ = imageArrayRotationAnimationCacheable
        imageArrayRotationAnimationCacheable.SetImageArray(allBinaryImageRotationAnimationInfo.GetImageArray())
        return cacheableInterface
}

