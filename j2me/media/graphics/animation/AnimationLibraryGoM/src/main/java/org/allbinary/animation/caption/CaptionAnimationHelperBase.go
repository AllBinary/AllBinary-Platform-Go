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

type CaptionAnimationHelperBase struct {
    Animation
}

// The Java class CaptionAnimationHelperBase has subclasses so its variables are this interface
type ICaptionAnimationHelperBase interface {
    IAnimation
    CaptionAnimationHelperBase_() *CaptionAnimationHelperBase
    IsShowing() bool
    Tick()
    Update(message string, basicColor *BasicColor)
    PaintXY(graphics IGraphics, x int, y int)
}

func (this *CaptionAnimationHelperBase) CaptionAnimationHelperBase_() *CaptionAnimationHelperBase {
    return this
}

var CaptionAnimationHelperBase_INSTANCE ICaptionAnimationHelperBase = NewCaptionAnimationHelperBase()

func NewCaptionAnimationHelperBase() *CaptionAnimationHelperBase {
    this := &CaptionAnimationHelperBase{}
    this.Animation = *NewAnimation()
    return this
}


func (this *CaptionAnimationHelperBase) IsShowing() bool {
        return false
}


func (this *CaptionAnimationHelperBase) Tick() {
}


func (this *CaptionAnimationHelperBase) Update(message string, basicColor *BasicColor) {
}


func (this *CaptionAnimationHelperBase) PaintXY(graphics IGraphics, x int, y int) {
}

