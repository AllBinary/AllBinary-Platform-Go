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

type TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory struct {
    BaseResourceAnimationInterfaceFactoryInterfaceFactory
}

// The Java class TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory has subclasses so its variables are this interface
type ITouchButtonResourceAnimationInterfaceFactoryInterfaceFactory interface {
    IBaseResourceAnimationInterfaceFactoryInterfaceFactory
    TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory_() *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory
    Init(level int)
    InitImageCache(imageCache IImageCache, level int)
    IsLoadingLevel(level int) bool
    IsFeature() bool
}

func (this *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory) TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory_() *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory {
    return this
}



func TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory_CreateFactory() ITouchButtonResourceAnimationInterfaceFactoryInterfaceFactory {
        return NewTouchButtonResourceAnimationInterfaceFactoryInterfaceFactory("TouchButton Animations")
}


func NewTouchButtonResourceAnimationInterfaceFactoryInterfaceFactory(name string) *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory {
    this := &TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory{}
        this.BaseResourceAnimationInterfaceFactoryInterfaceFactory = *NewBaseResourceAnimationInterfaceFactoryInterfaceFactory(name, StdUtil_GetInstance().CreateHashtable(), StdUtil_GetInstance().CreateHashtable(), StdUtil_GetInstance().CreateHashtable())
    return this
}


// throws Exception
func (this *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory) Init(level int) {
        this.InitImageCache(ImageCacheFactory_GetInstance(), level)
}


// throws Exception
func (this *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory) InitImageCache(imageCache IImageCache, level int) {
        if this.IsInitialized() {
            return
        }
        this.BaseResourceAnimationInterfaceFactoryInterfaceFactory.Init(level)
}


func (this *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory) IsLoadingLevel(level int) bool {
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        if level == resourceLoadingLevelFactory.LOAD_TOUCH.GetLevel() {
            return true
        } else {
            return this.BaseResourceAnimationInterfaceFactoryInterfaceFactory.IsLoadingLevel(level)
        }
}


func (this *TouchButtonResourceAnimationInterfaceFactoryInterfaceFactory) IsFeature() bool {
        if !Features_GetInstance().IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
            return true
        } else {
            return false
        }
}

