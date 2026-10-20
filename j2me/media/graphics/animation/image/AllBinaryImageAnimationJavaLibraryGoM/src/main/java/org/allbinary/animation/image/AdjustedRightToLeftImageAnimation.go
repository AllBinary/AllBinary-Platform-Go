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

type AdjustedRightToLeftImageAnimation struct {
    RightToLeftImageAnimation
    Dx int
    Dy int
}


func (this *AdjustedRightToLeftImageAnimation) initFields() {
    
    
}


func NewAdjustedRightToLeftImageAnimation(image IImage, dx int, dy int, animationBehavior IAnimationBehavior) *AdjustedRightToLeftImageAnimation {
    this := &AdjustedRightToLeftImageAnimation{}
        this.RightToLeftImageAnimation = *NewRightToLeftImageAnimation(image, animationBehavior)
        this.initFields()
        this.Init(dx, dy)
    return this
}


// throws Exception
func (this *AdjustedRightToLeftImageAnimation) Init(dx int, dy int) {
        this.Dx = dx
        this.Dy = dy
}


func (this *AdjustedRightToLeftImageAnimation) PaintXY(graphics IGraphics, x int, y int) {
        this.RightToLeftImageAnimation.PaintXY(graphics, x + this.Dx, y + this.Dy)
}

