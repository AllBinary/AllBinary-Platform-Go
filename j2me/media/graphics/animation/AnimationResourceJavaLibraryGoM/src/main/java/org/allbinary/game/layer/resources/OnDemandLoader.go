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

type OnDemandLoader struct {
    OnDemandResourcesFactory *OnDemandResourcesFactory
    GameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory GameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory
    ResourceAnimationInterfaceFactoryInterfaceFactory IBaseResourceAnimationInterfaceFactoryInterfaceFactory
}


func (this *OnDemandLoader) initFields() {
    
    
    
}


func NewOnDemandLoader(onDemandResourcesFactory *OnDemandResourcesFactory, gameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory GameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory, resourceAnimationInterfaceFactoryInterfaceFactory IBaseResourceAnimationInterfaceFactoryInterfaceFactory) *OnDemandLoader {
    this := &OnDemandLoader{}
    this.initFields()
        this.OnDemandResourcesFactory = onDemandResourcesFactory
        this.GameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory = gameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory
        this.ResourceAnimationInterfaceFactoryInterfaceFactory = resourceAnimationInterfaceFactoryInterfaceFactory
    return this
}


// throws Exception
func (this *OnDemandLoader) Init(imageCache IImageCache, level int, loadingString string) {
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        if Features_GetInstance().IsFeature(MainFeatureFactory_GetInstance().LOAD_ALL) {
            if this.ResourceAnimationInterfaceFactoryInterfaceFactory.IsInitialized() {
                return
            }
            this.InitAll(imageCache, loadingString)
        } else if level != resourceLoadingLevelFactory.LOAD_ALL.GetLevel() {
            var layerPlacementInterfaceBasicArrayListFactoryInterface LayerPlacementInterfaceBasicArrayListFactoryInterface = this.GameLayerPlacementInterfaceBasicArrayListFactoryInterfaceFactory.GetInstance(level); _ = layerPlacementInterfaceBasicArrayListFactoryInterface
            this.InitLevel(layerPlacementInterfaceBasicArrayListFactoryInterface, imageCache, level, loadingString)
        }
}


// throws Exception
func (this *OnDemandLoader) InitAll(imageCache IImageCache, loadingString string) {
        var index int = 0; _ = index
        var portion int = 120; _ = portion
        var onDemandResources *OnDemandResources; _ = onDemandResources
        for onDemandIndex := int(this.OnDemandResourcesFactory.Size()); Java_PreDecrement(&onDemandIndex) >= 0;  {
            onDemandResources = this.OnDemandResourcesFactory.Remove(onDemandIndex)
            if onDemandResources != OnDemandResources_NULL_ON_DEMAND_RESOURCES {
                index = onDemandResources.InitAt(imageCache, this.ResourceAnimationInterfaceFactoryInterfaceFactory, portion, loadingString, index)
            }
        }
}


// throws Exception
func (this *OnDemandLoader) InitLevel(layerPlacementInterfaceBasicArrayListFactoryInterface LayerPlacementInterfaceBasicArrayListFactoryInterface, imageCache IImageCache, level int, loadingString string) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var portion int = 120; _ = portion
        var index int = 0; _ = index
        var onDemandArray []int = layerPlacementInterfaceBasicArrayListFactoryInterface.GetOnDemandResources(); _ = onDemandArray
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(commonStrings.START).Appendint(len(onDemandArray))), this, "initLevel")
        var onDemandResources *OnDemandResources; _ = onDemandResources
        for onDemandIndex := int(len(onDemandArray)); Java_PreDecrement(&onDemandIndex) >= 0;  {
            onDemandResources = this.OnDemandResourcesFactory.Remove(onDemandArray[onDemandIndex])
            if onDemandResources != OnDemandResources_NULL_ON_DEMAND_RESOURCES {
                index = onDemandResources.InitAt(imageCache, this.ResourceAnimationInterfaceFactoryInterfaceFactory, portion, loadingString, index)
            }
        }
}

