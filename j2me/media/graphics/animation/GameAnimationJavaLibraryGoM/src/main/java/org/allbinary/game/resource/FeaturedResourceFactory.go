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

type FeaturedResourceFactory struct {
    LogUtil *LogUtil
    List IBasicArrayList
    CommonStrings *CommonStrings
    ANIMATION_FEATURES string
    IMAGE_LABEL string
    IMAGE_GRAPHICS_ARRAY string
    IMAGE_GRAPHICS_ROTATION string
    SPRITE_QUARTER string
    SPRITE_FULL string
    IS_LOADING_LEVEL_LABEL string
    IS_FEATURE string
    GAME_FEATURE_CONTROLLED string
}

// The Java class FeaturedResourceFactory has subclasses so its variables are this interface
type IFeaturedResourceFactory interface {
    FeaturedResourceFactory_() *FeaturedResourceFactory
    Init(level int)
    Clear()
    Add(featureInterface GameFeatureControlledInterface)
    GetList() IBasicArrayList
}

func (this *FeaturedResourceFactory) FeaturedResourceFactory_() *FeaturedResourceFactory {
    return this
}


func (this *FeaturedResourceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.List = NewBasicArrayListD()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ANIMATION_FEATURES = "Animation Features: Vector: "
    this.IMAGE_LABEL = " Image: "
    this.IMAGE_GRAPHICS_ARRAY = "Image Array: "
    this.IMAGE_GRAPHICS_ROTATION = "Image Rotate: "
    this.SPRITE_QUARTER = " Sprite Quarter: "
    this.SPRITE_FULL = " Sprite Full: "
    this.IS_LOADING_LEVEL_LABEL = " isLoadingLevel "
    this.IS_FEATURE = " isFeature: "
    this.GAME_FEATURE_CONTROLLED = "GameFeatureControlledInterface: "
}


func NewFeaturedResourceFactory() *FeaturedResourceFactory {
    this := &FeaturedResourceFactory{}
    this.initFields()
    return this
}


// throws Exception
func (this *FeaturedResourceFactory) Init(level int) {
        var size int = this.List.Size(); _ = size
        for index := int(0); index < size; index++ {
            var featureInterface GameFeatureControlledInterface = any(this.List.BasicArrayList_().ObjectArray[index]).(GameFeatureControlledInterface); _ = featureInterface
            var isLoadingLevel bool = featureInterface.IsLoadingLevel(level); _ = isLoadingLevel
            var isFeature bool = featureInterface.IsFeature(); _ = isFeature
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.GAME_FEATURE_CONTROLLED).Append(fmt.Sprint(featureInterface)).Append(this.IS_LOADING_LEVEL_LABEL).Append(ResourceLoadingLevelFactory_GetInstance().GetLevelString(level)).Append(CommonSeps_GetInstance().COLON_SEP).Appendboolean(isLoadingLevel).Append(this.IS_FEATURE).Appendboolean(isFeature)), this, this.CommonStrings.INIT)
            if isLoadingLevel && isFeature {
                featureInterface.Init(level)
            }
        }
        var features *Features = Features_GetInstance(); _ = features
        var graphicsFeatureFactory *GraphicsFeatureFactory = GraphicsFeatureFactory_GetInstance(); _ = graphicsFeatureFactory
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.ANIMATION_FEATURES)
        stringBuffer.Appendboolean(features.IsFeature(graphicsFeatureFactory.VECTOR_GRAPHICS))
        stringBuffer.Append(this.IMAGE_LABEL)
        stringBuffer.Appendboolean(features.IsFeature(graphicsFeatureFactory.IMAGE_GRAPHICS))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.INIT)
        if features.IsFeature(graphicsFeatureFactory.IMAGE_GRAPHICS) {
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(this.IMAGE_GRAPHICS_ARRAY)
            stringBuffer.Appendboolean(features.IsFeature(graphicsFeatureFactory.IMAGE_TO_ARRAY_GRAPHICS))
            stringBuffer.Append(this.IMAGE_GRAPHICS_ROTATION)
            stringBuffer.Appendboolean(features.IsFeature(graphicsFeatureFactory.IMAGE_TO_ARRAY_GRAPHICS))
            stringBuffer.Append(this.SPRITE_QUARTER)
            stringBuffer.Appendboolean(features.IsFeature(graphicsFeatureFactory.SPRITE_QUARTER_ROTATION_GRAPHICS))
            stringBuffer.Append(this.SPRITE_FULL)
            stringBuffer.Appendboolean(features.IsFeature(graphicsFeatureFactory.SPRITE_FULL_GRAPHICS))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.INIT)
        }
}


func (this *FeaturedResourceFactory) Clear() {
        this.List.Clear()
}


func (this *FeaturedResourceFactory) Add(featureInterface GameFeatureControlledInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(StringUtil_GetInstance().ToString(featureInterface))), this, this.CommonStrings.ADD)
        this.List.Add(featureInterface)
}


func (this *FeaturedResourceFactory) GetList() IBasicArrayList {
        return this.List
}

