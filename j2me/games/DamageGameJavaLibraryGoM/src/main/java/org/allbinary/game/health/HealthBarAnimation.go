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

type HealthBarAnimation struct {
    Animation
    BasicColorFactory *BasicColorFactory
    X2 int
    Thickness int
    Location int
    AllbinaryLayer IAllBinaryLayer
}

// The Java class HealthBarAnimation has subclasses so its variables are this interface
type IHealthBarAnimation interface {
    IAnimation
    HealthBarAnimation_() *HealthBarAnimation
    OnHealthChange(newX2 int)
    GetY() int
}

func (this *HealthBarAnimation) HealthBarAnimation_() *HealthBarAnimation {
    return this
}


func (this *HealthBarAnimation) initFields() {
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    
    
    
    
}


func NewHealthBarAnimation(layerInterface IAllBinaryLayer, location int) *HealthBarAnimation {
    this := &HealthBarAnimation{}
    this.Animation = *NewAnimation()
    this.initFields()
        this.AllbinaryLayer = layerInterface
        this.Location = location
        var basicHudFactory *BasicHudFactory = BasicHudFactory_GetInstance(); _ = basicHudFactory
        if this.Location != basicHudFactory.TOPLEFT && this.Location != basicHudFactory.BOTTOMLEFT {
            panic(errors.New(fmt.Sprint("Location Not Valid")))
        }
        if layerInterface.GetWidth() > 40 {
            this.Thickness = 3
        } else if layerInterface.GetWidth() > 20 {
            this.Thickness = 2
        } else {
            this.Thickness = 1
        }
    return this
}


func (this *HealthBarAnimation) OnHealthChange(newX2 int) {
        this.X2 = newX2
        var quarter int = (this.AllbinaryLayer.GetWidth() >> 2); _ = quarter
        if this.X2 > quarter * 3 {
            this.BasicColor = this.BasicColorFactory.GREEN
            this.ColorP = int(this.BasicColor)
        } else if this.X2 > quarter * 2 {
            this.BasicColor = this.BasicColorFactory.YELLOW
            this.ColorP = int(this.BasicColor)
        } else if this.X2 > quarter {
            this.BasicColor = this.BasicColorFactory.ORANGE
            this.ColorP = int(this.BasicColor)
        } else {
            this.BasicColor = this.BasicColorFactory.RED
            this.ColorP = int(this.BasicColor)
        }
}


func (this *HealthBarAnimation) GetY() int {
        var basicHudFactory *BasicHudFactory = BasicHudFactory_GetInstance(); _ = basicHudFactory
        if this.Location == basicHudFactory.TOPLEFT {
            return this.AllbinaryLayer.GetViewPosition().GetY() - 4
        } else if this.Location == basicHudFactory.BOTTOMLEFT {
            return this.AllbinaryLayer.GetViewPosition().GetY2() + 4
        }
        return -1
}

