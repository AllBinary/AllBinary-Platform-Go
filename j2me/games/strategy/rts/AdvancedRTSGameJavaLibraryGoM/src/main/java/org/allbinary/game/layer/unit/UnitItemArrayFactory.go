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

type UnitItemArrayFactory struct {
    CommonSeps *CommonSeps
    BasicColor *BasicColor
    Resources []IBasicGameResources
}

// Implements ItemArraySingletonFactoryInterface

var UnitItemArrayFactory_ITEMS []IABCustomItem = CustomItemUtil_GetInstance().CUSTOM_ITEM_ARRAY

func (this *UnitItemArrayFactory) initFields() {
    this.CommonSeps = CommonSeps_GetInstance()
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    
}


func NewUnitItemArrayFactory(resources []IBasicGameResources) *UnitItemArrayFactory {
    this := &UnitItemArrayFactory{}
    this.initFields()
        this.Resources = resources
    return this
}


// throws Exception
func (this *UnitItemArrayFactory) CreateFlagItem(flagResources *FlagGameResources, layerInterfaceFactoryInterface LayerInterfaceFactoryInterface) IABCustomItem {
        var name string = flagResources.NAME; _ = name
        var index int = strings.Index(name, string(this.CommonSeps.SPACE)); _ = index
        if index >= 0 {
            name = name[index:]
        }
        return this.CreateNamedFlagItem(flagResources, name, layerInterfaceFactoryInterface)
}


// throws Exception
func (this *UnitItemArrayFactory) CreateNamedFlagItem(flagResources *FlagGameResources, name string, layerInterfaceFactoryInterface LayerInterfaceFactoryInterface) IABCustomItem {
        var IMAGE_CACHE IImageCache = GameFeatureImageCacheFactory_GetInstance(); _ = IMAGE_CACHE
        var image IImage = IMAGE_CACHE.GetWithKey(flagResources.RESOURCE_ICON); _ = image
        var item IABCustomItem = NewLayerInterfaceFactoryImageItem(name, image, ImageItem_LAYOUT_DEFAULT, flagResources.NAME, this.BasicColor, []IAnimation{NewRTSLayerTextAnimation(RTSGameStrings_GetInstance().DRAGGABLE, image)}, layerInterfaceFactoryInterface); _ = item
        return item
}


// throws Exception
func (this *UnitItemArrayFactory) GetItems() []IABCustomItem {
        return UnitItemArrayFactory_ITEMS
}


func (this *UnitItemArrayFactory) GetResources() []IBasicGameResources {
        return this.Resources
}


func (this *UnitItemArrayFactory) IndexOf(basicGameResources IBasicGameResources) int {
        for index := int(len(this.Resources) - 1); index >= 0; index-- {
            if this.Resources[index] == basicGameResources {
                return index
            }
        }
        return -1
}

