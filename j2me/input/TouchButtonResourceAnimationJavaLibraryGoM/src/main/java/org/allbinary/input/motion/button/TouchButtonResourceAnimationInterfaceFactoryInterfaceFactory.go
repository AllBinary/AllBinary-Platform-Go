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
        var portion int = 140; _ = portion
        var loadingString string = fmt.Sprint(NewStringMaker().Append(this.ToString()).Append(" Loading: ")); _ = loadingString
        var index int = 0; _ = index
        var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        var touchButtonBlankResource ITouchButtonResource = TouchButtonBlankResource_GetInstance(); _ = touchButtonBlankResource
        var touchButtonGenericActionResource ITouchButtonResource = TouchButtonGenericActionResource_GetInstance(); _ = touchButtonGenericActionResource
        var touchButtonStartResource ITouchButtonResource = TouchButtonStartResource_GetInstance(); _ = touchButtonStartResource
        var touchButtonUpResource ITouchButtonResource = TouchButtonUpResource_GetInstance(); _ = touchButtonUpResource
        var touchButtonDownResource ITouchButtonResource = TouchButtonDownResource_GetInstance(); _ = touchButtonDownResource
        var touchButtonTurnLeftResource ITouchButtonResource = TouchButtonTurnLeftResource_GetInstance(); _ = touchButtonTurnLeftResource
        var touchButtonTurnRightResource ITouchButtonResource = TouchButtonTurnRightResource_GetInstance(); _ = touchButtonTurnRightResource
        var touchButtonStrafeLeftResource ITouchButtonResource = TouchButtonStrafeLeftResource_GetInstance(); _ = touchButtonStrafeLeftResource
        var touchButtonStrafeRightResource ITouchButtonResource = TouchButtonStrafeRightResource_GetInstance(); _ = touchButtonStrafeRightResource
        var touchButtonBlankResourceImage IImage = imageCache.GetWithKey(touchButtonBlankResource.TouchButtonResource_().RESOURCE); _ = touchButtonBlankResourceImage
        var touchButtonGenericActionResourceImage IImage = imageCache.GetWithKey(touchButtonGenericActionResource.TouchButtonResource_().RESOURCE); _ = touchButtonGenericActionResourceImage
        var touchButtonStartResourceImage IImage = imageCache.GetWithKey(touchButtonStartResource.TouchButtonResource_().RESOURCE); _ = touchButtonStartResourceImage
        var touchButtonUpResourceImage IImage = imageCache.GetWithKey(touchButtonUpResource.TouchButtonResource_().RESOURCE); _ = touchButtonUpResourceImage
        var touchButtonDownResourceImage IImage = imageCache.GetWithKey(touchButtonDownResource.TouchButtonResource_().RESOURCE); _ = touchButtonDownResourceImage
        var touchButtonTurnLeftResourceImage IImage = imageCache.GetWithKey(touchButtonTurnLeftResource.TouchButtonResource_().RESOURCE); _ = touchButtonTurnLeftResourceImage
        var touchButtonTurnRightResourceImage IImage = imageCache.GetWithKey(touchButtonTurnRightResource.TouchButtonResource_().RESOURCE); _ = touchButtonTurnRightResourceImage
        var touchButtonStrafeLeftResourceImage IImage = imageCache.GetWithKey(touchButtonStrafeLeftResource.TouchButtonResource_().RESOURCE); _ = touchButtonStrafeLeftResourceImage
        var touchButtonStrafeRightResourceImage IImage = imageCache.GetWithKey(touchButtonStrafeRightResource.TouchButtonResource_().RESOURCE); _ = touchButtonStrafeRightResourceImage
        var nullAnimationFactory *NullAnimationFactory = NullAnimationFactory_GetFactoryInstance(); _ = nullAnimationFactory
        if OperatingSystemFactory_GetInstance().GetOperatingSystemInstance().IsOverScan() {
            imageCache.GetWithKey(TouchButtonStartResource_GetInstance().TouchButtonResource_().HINT)
        }
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        var imageCompleteUtil IImageCompleteUtil = ImageCompleteUtil_GetInstance(); _ = imageCompleteUtil
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonBlankResourceImage, touchButtonBlankResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonBlankResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonBlankResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonBlankResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonGenericActionResourceImage, touchButtonGenericActionResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonGenericActionResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonGenericActionResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonGenericActionResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonStartResourceImage, touchButtonStartResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonStartResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonStartResourceImage, AnimationBehaviorFactory_GetInstance()))
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonUpResourceImage, touchButtonUpResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonUpResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonUpResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonUpResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonDownResourceImage, touchButtonDownResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonDownResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonDownResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonDownResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonTurnLeftResourceImage, touchButtonTurnLeftResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonTurnLeftResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonTurnLeftResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonTurnLeftResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonTurnRightResourceImage, touchButtonTurnRightResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonTurnRightResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonTurnRightResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonTurnRightResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonStrafeLeftResourceImage, touchButtonStrafeLeftResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonStrafeLeftResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonStrafeLeftResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonStrafeLeftResource.TouchButtonResource_().HINT, nullAnimationFactory)
        progressCanvas.AddPortion(portion, loadingString, Java_PostIncrement(&index))
        imageCompleteUtil.WaitForImage(touchButtonStrafeRightResourceImage, touchButtonStrafeRightResource.TouchButtonResource_().RESOURCE)
        this.Add(touchButtonStrafeRightResource.TouchButtonResource_().RESOURCE, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonStrafeRightResourceImage, AnimationBehaviorFactory_GetInstance()))
        this.Add(touchButtonStrafeRightResource.TouchButtonResource_().HINT, nullAnimationFactory)
        if OperatingSystemFactory_GetInstance().GetOperatingSystemInstance().IsOverScan() {
            var touchButtonStartHintResource IImage = imageCache.GetWithKey(TouchButtonStartResource_GetInstance().TouchButtonResource_().HINT); _ = touchButtonStartHintResource
            this.Add(touchButtonStartResource.TouchButtonResource_().HINT, OneRowSpriteIndexedAnimationFactory_CreateFactory(touchButtonStartHintResource, AnimationBehaviorFactory_GetInstance()))
        } else {
            this.Add(touchButtonStartResource.TouchButtonResource_().HINT, nullAnimationFactory)
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

