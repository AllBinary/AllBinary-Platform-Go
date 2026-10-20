/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type AdjustCustomTextAnimation struct {
    CustomTextAnimation
    Dx int
    Dy int
}


func (this *AdjustCustomTextAnimation) initFields() {
    
    
}


func NewAdjustCustomTextAnimation(text string, fontSize int, dx int, dy int, animationBehavior IAnimationBehavior) *AdjustCustomTextAnimation {
    this := &AdjustCustomTextAnimation{}
        this.CustomTextAnimation = *NewCustomTextAnimation(text, fontSize, animationBehavior)
        this.initFields()
        this.Dx = dx
        this.Dy = dy
    return this
}


func (this *AdjustCustomTextAnimation) GetDx() int {
        return this.Dx
}


func (this *AdjustCustomTextAnimation) SetDx(dx int) {
        this.Dx = dx
}


func (this *AdjustCustomTextAnimation) GetDy() int {
        return this.Dy
}


func (this *AdjustCustomTextAnimation) SetDy(dy int) {
        this.Dy = dy
}


func (this *AdjustCustomTextAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.CustomTextAnimation.PaintXY(graphics, x + this.Dx, y + this.Dy)
}

