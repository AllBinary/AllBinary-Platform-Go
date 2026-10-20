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

type FeaturedResourceRelativeRelationshipFactory struct {
    FeaturedResourceFactory
    BasicArrayListUtil *BasicArrayListUtil
}

var FeaturedResourceRelativeRelationshipFactory_INSTANCE *FeaturedResourceRelativeRelationshipFactory = NewFeaturedResourceRelativeRelationshipFactory()

func (this *FeaturedResourceRelativeRelationshipFactory) initFields() {
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
}


func NewFeaturedResourceRelativeRelationshipFactory() *FeaturedResourceRelativeRelationshipFactory {
    this := &FeaturedResourceRelativeRelationshipFactory{}
    this.FeaturedResourceFactory = *NewFeaturedResourceFactory()
    this.initFields()
    return this
}


func FeaturedResourceRelativeRelationshipFactory_GetInstance() *FeaturedResourceRelativeRelationshipFactory {
        return FeaturedResourceRelativeRelationshipFactory_INSTANCE
}


// throws Exception
func (this *FeaturedResourceRelativeRelationshipFactory) GetRelativeRelationshipList(resource string, layer IAllBinaryLayer) IBasicArrayList {
        var featureReleaseList IBasicArrayList = this.GetList(); _ = featureReleaseList
        var size int = this.GetList().Size(); _ = size
        for index := int(0); index < size; index++ {
            var featureInterface ResourceRelativeRelationshipFactoryInterface = any(featureReleaseList.BasicArrayList_().ObjectArray[index]).(ResourceRelativeRelationshipFactoryInterface); _ = featureInterface
            if featureInterface.IsFeature() {
                var list IBasicArrayList = featureInterface.GetResourceRelativeRelationshipList(resource); _ = list
                if list != this.BasicArrayListUtil.GetImmutableInstance() {
                    return this.Duplicate(list, layer)
                }
            }
        }
        panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Not available for current feature selection or Resource: ").Append(resource)))))
}


// throws Exception
func (this *FeaturedResourceRelativeRelationshipFactory) Duplicate(list IBasicArrayList, layer IAllBinaryLayer) IBasicArrayList {
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        var newList IBasicArrayList = NewBasicArrayListD(); _ = newList
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var relativeRelationship IRelativeRelationship = any(list.BasicArrayList_().ObjectArray[index]).(IRelativeRelationship); _ = relativeRelationship
            newList.Add(NewRelativeLayerRelationship(layer, pointFactory.CreateXY(relativeRelationship.GetX(), relativeRelationship.GetY()), BasicArrayListUtil_GetInstance().GetImmutableInstance()))
        }
        return newList
}

