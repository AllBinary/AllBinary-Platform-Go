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

type AnimationDamageFloaters struct {
    DamageFloaters
    LogUtil *LogUtil
    LayerInterface IAllBinaryLayer
    AnimationInterfaceArray []IIndexedAnimation
    CircularIndexUtil *CircularIndexUtil
    Dx int
    Dy int
}


func (this *AnimationDamageFloaters) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
}


func NewAnimationDamageFloaters(layerInterface IAllBinaryLayer, animationInterfaceArray []IIndexedAnimation, dx int, dy int) *AnimationDamageFloaters {
    this := &AnimationDamageFloaters{}
    this.DamageFloaters = *NewDamageFloaters()
    this.initFields()
        this.AnimationInterfaceArray = animationInterfaceArray
        for index := int(len(this.AnimationInterfaceArray) - 1); index >= 0; index-- {
            this.AnimationInterfaceArray[index].SetFrame(this.AnimationInterfaceArray[index].GetSize() - 1)
        }
        this.LayerInterface = layerInterface
        this.CircularIndexUtil = CircularIndexUtil_CreateInstance(len(this.AnimationInterfaceArray))
        this.Dx = dx
        this.Dy = dy
    return this
}


func (this *AnimationDamageFloaters) Add(damage int) {
        var i int = this.CircularIndexUtil.GetIndex(); _ = i
        this.AnimationInterfaceArray[i].SetFrame(0)
        this.CircularIndexUtil.Next()
}


func (this *AnimationDamageFloaters) Paint(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            var canvasStrings *CanvasStrings = CanvasStrings_GetInstance(); _ = canvasStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, canvasStrings.PAINT, e)
                }
            }()
            var viewPosition IViewPositionBase = this.LayerInterface.GetViewPosition(); _ = viewPosition
            var x int = viewPosition.GetX(); _ = x
            var y int = viewPosition.GetY(); _ = y
            for index := int(0); index < len(this.AnimationInterfaceArray); index++ {
                var animationInterface IIndexedAnimation = this.AnimationInterfaceArray[index]; _ = animationInterface
                if animationInterface.GetFrame() < animationInterface.GetAnimationSize() - 1 {
                    var delta int = animationInterface.GetFrame() * 20; _ = delta
                    animationInterface.PaintXY(graphics, x + this.Dx, y - delta + this.Dy)
                    animationInterface.NextFrame()
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

