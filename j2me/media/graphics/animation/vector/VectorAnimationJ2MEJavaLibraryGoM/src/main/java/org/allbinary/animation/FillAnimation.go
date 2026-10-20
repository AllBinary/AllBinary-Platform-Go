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

type FillAnimation struct {
    IndexedAnimation
    PixelBuffer []int
    Width int
    Height int
}


func (this *FillAnimation) initFields() {
    
    
    
}


func NewFillAnimation(pixelBuffer []int, width int, height int, animationBehavior IAnimationBehavior) *FillAnimation {
    this := &FillAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
        this.initFields()
        this.Width = 100
        this.Height = 100
        this.PixelBuffer = pixelBuffer
    return this
}


func (this *FillAnimation) SetBasicColorP(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetBasicColorP() == nil || this.GetBasicColorP().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.IndexedAnimation.SetBasicColorP(basicColor)
        if changed {
            this.UpdateModifiers()
        }
}


func (this *FillAnimation) ChangeBasicColor(basicColor *BasicColor) {
        var changed bool = false; _ = changed
        if this.GetChangeBasicColor() == nil || this.GetChangeBasicColor().IntValue() != basicColor.IntValue() {
            changed = true
        }
        this.IndexedAnimation.ChangeBasicColor(basicColor)
        if changed {
            this.UpdateModifiers()
        }
}


func (this *FillAnimation) SetAlpha(alpha int) {
        var changed bool = false; _ = changed
        if this.AlphaP != alpha {
            changed = true
        }
        this.IndexedAnimation.SetAlpha(alpha)
        if changed {
            this.UpdateModifiers()
        }
}


func (this *FillAnimation) SetScale(scaleX float32, scaleY float32) {
}


func (this *FillAnimation) SetMaxScale(maxScaleX float32, maxScaleY float32) {
}


func (this *FillAnimation) UpdateModifiers() {
}


func (this *FillAnimation) Paint(g IGraphics) {
        g.DrawRGB(this.PixelBuffer, 0, this.Width, 0, 0, this.Width, this.Height, true)
}

