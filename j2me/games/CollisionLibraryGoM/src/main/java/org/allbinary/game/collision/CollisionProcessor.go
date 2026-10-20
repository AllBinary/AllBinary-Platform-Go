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

type CollisionProcessor struct {
}

// The Java class CollisionProcessor has subclasses so its variables are this interface
type ICollisionProcessor interface {
    CollisionProcessor_() *CollisionProcessor
    IsCollision(myLayer IAllBinaryLayer, myLayer2 IAllBinaryLayer) bool
}

func (this *CollisionProcessor) CollisionProcessor_() *CollisionProcessor {
    return this
}


func NewCollisionProcessor() *CollisionProcessor {
    this := &CollisionProcessor{}
    return this
}


func (this *CollisionProcessor) IsCollision(myLayer IAllBinaryLayer, myLayer2 IAllBinaryLayer) bool {
        return false
}

