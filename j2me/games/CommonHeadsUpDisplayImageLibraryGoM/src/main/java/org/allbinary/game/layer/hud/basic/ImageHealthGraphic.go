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

type ImageHealthGraphic struct {
    HealthGraphic
    Image []IImage
}


func (this *ImageHealthGraphic) initFields() {
    
}


func NewImageHealthGraphic(animationInterface AnimationInterface, healthInterface HealthInterface, location int, direction int) *ImageHealthGraphic {
    this := &ImageHealthGraphic{}
        this.HealthGraphic = *NewHealthGraphic(animationInterface, healthInterface, location, direction)
        this.initFields()
        this.Image = make([]IImage, this.HealthInterface.GetMaxHealth())
        var imageCacheFactory *ImageCacheFactory = ImageCacheFactory_GetInstance(); _ = imageCacheFactory
        for numOfHealth := int(0); numOfHealth <= this.Max; numOfHealth++ {
            this.Image[numOfHealth] = imageCacheFactory.Get(this, numOfHealth + 1 * 10, 10)
        }
        for numOfHealth := int(0); numOfHealth <= this.Max; numOfHealth++ {
            for index := int(0); index < numOfHealth; index++ {
                this.AnimationInterface.Paint(this.Image[numOfHealth].GetGraphics(), numOfHealth * 10, 0)
            }
        }
    return this
}


func (this *ImageHealthGraphic) Paint(graphics IGraphics) {
        graphics.DrawImage(this.Image[this.Max], this.X, this.Y, Anchor_TOP_LEFT)
}

