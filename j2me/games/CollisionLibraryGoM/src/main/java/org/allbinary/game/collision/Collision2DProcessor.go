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

type Collision2DProcessor struct {
    CollisionProcessor
}


func NewCollision2DProcessor() *Collision2DProcessor {
    this := &Collision2DProcessor{}
    this.CollisionProcessor = *NewCollisionProcessor()
    return this
}


func (this *Collision2DProcessor) IsCollision(myLayer IAllBinaryLayer, myLayer2 IAllBinaryLayer) bool {
        if myLayer2.GetXP() >= myLayer.GetX2() || myLayer2.GetYP() >= myLayer.GetY2() || myLayer2.GetX2() <= myLayer.GetXP() || myLayer2.GetY2() <= myLayer.GetYP() {
            return false
        } else {
            return true
        }
}

