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

type ARectangleFilledAnimation struct {
    RectangleFilledAnimation
    X int
    Y int
}


func (this *ARectangleFilledAnimation) initFields() {
    
    
}


func NewARectangleFilledAnimation() *ARectangleFilledAnimation {
    this := &ARectangleFilledAnimation{}
        this.RectangleFilledAnimation = *NewRectangleFilledAnimation(0, 0, BasicColorFactory_GetInstance().BLACK)
        this.initFields()
    return this
}


func (this *ARectangleFilledAnimation) NextFrame() {
}


func (this *ARectangleFilledAnimation) PaintXY(graphics IGraphics, unusedX int, unusedY int) {
        this.RectangleFilledAnimation.PaintXY(graphics, this.X, this.Y)
}

