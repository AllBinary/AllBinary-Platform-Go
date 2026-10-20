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

type NullIndexedAnimation struct {
    IndexedAnimation
}

// The Java class NullIndexedAnimation has subclasses so its variables are this interface
type INullIndexedAnimation interface {
    IIndexedAnimation
    NullIndexedAnimation_() *NullIndexedAnimation
    GetAnimationSize() int
}

func (this *NullIndexedAnimation) NullIndexedAnimation_() *NullIndexedAnimation {
    return this
}



func NewNullIndexedAnimation(animationBehavior IAnimationBehavior) *NullIndexedAnimation {
    this := &NullIndexedAnimation{}
        this.IndexedAnimation = *NewIndexedAnimation(animationBehavior)
    return this
}


// throws Exception
func (this *NullIndexedAnimation) GetAnimationSize() int {
        return 0
}

