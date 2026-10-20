/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type LayerInterfaceFactoryImageItem struct {
    ABCustomImageItem
    AnimationArray []IAnimation
    LayerInterfaceFactoryInterface LayerInterfaceFactoryInterface
}


func (this *LayerInterfaceFactoryImageItem) initFields() {
    
    
}


func NewLayerInterfaceFactoryImageItem(label string, img IImage, layout int, altText string, basicColor *BasicColor, animationArray []IAnimation, layerInterfaceFactoryInterface LayerInterfaceFactoryInterface) *LayerInterfaceFactoryImageItem {
    this := &LayerInterfaceFactoryImageItem{}
        this.ABCustomImageItem = *NewABCustomImageItem(label, img, layout, altText, basicColor, 0)
        this.initFields()
        this.AnimationArray = animationArray
        this.LayerInterfaceFactoryInterface = layerInterfaceFactoryInterface
    return this
}


func (this *LayerInterfaceFactoryImageItem) GetLayerInterfaceFactoryInterface() LayerInterfaceFactoryInterface {
        return this.LayerInterfaceFactoryInterface
}


func (this *LayerInterfaceFactoryImageItem) PaintXY(graphics IGraphics, x int, y int) {
        this.ABCustomImageItem.PaintXY(graphics, x, y)
        for index := int(len(this.AnimationArray)); Java_PreDecrement(&index) >= 0;  {
            this.AnimationArray[index].PaintXY(graphics, x, y)
        }
}

