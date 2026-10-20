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

type FeaturedAnimationInterfaceFactoryInterfaceFactory struct {
    FeaturedResourceFactory
}

var FeaturedAnimationInterfaceFactoryInterfaceFactory_INSTANCE *FeaturedAnimationInterfaceFactoryInterfaceFactory = NewFeaturedAnimationInterfaceFactoryInterfaceFactory()


func NewFeaturedAnimationInterfaceFactoryInterfaceFactory() *FeaturedAnimationInterfaceFactoryInterfaceFactory {
    this := &FeaturedAnimationInterfaceFactoryInterfaceFactory{}
    this.FeaturedResourceFactory = *NewFeaturedResourceFactory()
    return this
}


func FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance() *FeaturedAnimationInterfaceFactoryInterfaceFactory {
        return FeaturedAnimationInterfaceFactoryInterfaceFactory_INSTANCE
}


// throws Exception
func (this *FeaturedAnimationInterfaceFactoryInterfaceFactory) GetRectangle(resource string) *Rectangle {
        return this.GetRectangleXY(resource, 0, 0)
}


// throws Exception
func (this *FeaturedAnimationInterfaceFactoryInterfaceFactory) GetRectangleXY(resource string, x int, y int) *Rectangle {
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var list IBasicArrayList = this.GetList(); _ = list
        var scale int = int(GameConfigurationCentral_GetInstance().SCALE.GetValue()); _ = scale
        var size int = this.GetList().Size(); _ = size
        for index := int(0); index < size; index++ {
            var featureInterface FeatureResourceAnimationInterfaceFactoryInterface = any(list.BasicArrayList_().ObjectArray[index]).(FeatureResourceAnimationInterfaceFactoryInterface); _ = featureInterface
            if featureInterface.IsFeature() {
                var rectangle *Rectangle = featureInterface.GetRectangle(resource); _ = rectangle
                if rectangle != RectangleFactory_SINGLETON {
                    return NewRectangle(pointFactory.CreateXY(x, y), ((rectangle.GetWidth() * scale) >> 1), ((rectangle.GetHeight() * scale) >> 1))
                }
            }
        }
        panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("No rectangle available for current feature selection or Resource: ").Append(resource)))))
}


// throws Exception
func (this *FeaturedAnimationInterfaceFactoryInterfaceFactory) GetProcedural(resource string) ProceduralAnimationInterfaceFactoryInterface {
        return any(this.GetBasicAnimationInterfaceFactoryInstance(resource)).(ProceduralAnimationInterfaceFactoryInterface)
}


// throws Exception
func (this *FeaturedAnimationInterfaceFactoryInterfaceFactory) Get(resource string) AnimationInterfaceFactoryInterface {
        return any(this.GetBasicAnimationInterfaceFactoryInstance(resource)).(AnimationInterfaceFactoryInterface)
}


// throws Exception
func (this *FeaturedAnimationInterfaceFactoryInterfaceFactory) GetBasicAnimationInterfaceFactoryInstance(resource string) BasicAnimationInterfaceFactoryInterface {
        var list IBasicArrayList = this.GetList(); _ = list
        var resourceTypeAvailableList IBasicArrayList = NewBasicArrayListD(); _ = resourceTypeAvailableList
        var size int = this.GetList().Size(); _ = size
        var featureInterface FeatureResourceAnimationInterfaceFactoryInterface; _ = featureInterface
        var animationInterfaceFactoryInterface BasicAnimationInterfaceFactoryInterface; _ = animationInterfaceFactoryInterface
        for index := int(0); index < size; index++ {
            featureInterface = any(list.BasicArrayList_().ObjectArray[index]).(FeatureResourceAnimationInterfaceFactoryInterface)
            if featureInterface.IsFeature() {
                resourceTypeAvailableList.Add(featureInterface)
                animationInterfaceFactoryInterface = featureInterface.GetBasicAnimationInterfaceFactoryInstance(resource)
                if animationInterfaceFactoryInterface != NullAnimationFactory_NULL_NOT_FOR_USE_ANIMATION_FACTORY {
                    return animationInterfaceFactoryInterface
                }
            }
        }
        if resourceTypeAvailableList.Size() > 0 {
            var NO_ANIMATION_AVAILABLE_FROM string = "No animation available from: "; _ = NO_ANIMATION_AVAILABLE_FROM
            var FOR_FACTORIES string = " factories: "; _ = FOR_FACTORIES
            var FOR_RESOURCE string = " for Resource: "; _ = FOR_RESOURCE
            var HAS_KEY string = " has: "; _ = HAS_KEY
            var RESOURCES_LABEL string = " resources "; _ = RESOURCES_LABEL
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(NO_ANIMATION_AVAILABLE_FROM)
            stringBuffer.Appendint(resourceTypeAvailableList.Size())
            stringBuffer.Append(CommonSeps_GetInstance().FORWARD_SLASH)
            stringBuffer.Appendint(size)
            stringBuffer.Append(FOR_FACTORIES)
            var size2 int = resourceTypeAvailableList.Size(); _ = size2
            for index := int(0); index < size2; index++ {
                stringBuffer.Append(fmt.Sprint(resourceTypeAvailableList.Get(index)))
                stringBuffer.Append(CommonSeps_GetInstance().COMMA_SEP)
            }
            stringBuffer.Append(FOR_RESOURCE)
            stringBuffer.Append(resource)
            for index := int(0); index < size; index++ {
                stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
                featureInterface = any(list.BasicArrayList_().ObjectArray[index]).(FeatureResourceAnimationInterfaceFactoryInterface)
                if featureInterface.IsFeature() {
                    stringBuffer.Append(StringUtil_GetInstance().ToString(featureInterface))
                    stringBuffer.Append(HAS_KEY)
                    stringBuffer.Appendint(featureInterface.GetHashtable().Size())
                    stringBuffer.Append(RESOURCES_LABEL)
                }
            }
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
        } else {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            for index := int(0); index < size; index++ {
                featureInterface = any(list.BasicArrayList_().ObjectArray[index]).(FeatureResourceAnimationInterfaceFactoryInterface)
                stringBuffer.Append(fmt.Sprint(featureInterface))
                stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            }
            var result string = fmt.Sprint(stringBuffer); _ = result
            stringBuffer.Delete(0, stringBuffer.Length())
            panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer.Append("No feature resource type available for Resource: ").Append(resource).Append(" Resource Factories Available: ").Append(result)))))
        }
}

