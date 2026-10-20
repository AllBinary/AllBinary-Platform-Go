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

type CompleteMotionGestureInputEventListener struct {
}

// Implements CompleteMotionGestureInputEventListenerInterface

// The Java class CompleteMotionGestureInputEventListener has subclasses so its variables are this interface
type ICompleteMotionGestureInputEventListener interface {
    CompleteMotionGestureInputEventListener_() *CompleteMotionGestureInputEventListener
    OnEvent(eventObject IAllBinaryEventObject)
    OnCompleteMotionGestureInputEvent(completeMotionGestureInputEvent ICompleteMotionGestureInputEvent)
}

func (this *CompleteMotionGestureInputEventListener) CompleteMotionGestureInputEventListener_() *CompleteMotionGestureInputEventListener {
    return this
}


func NewCompleteMotionGestureInputEventListener() *CompleteMotionGestureInputEventListener {
    this := &CompleteMotionGestureInputEventListener{}
    return this
}


func (this *CompleteMotionGestureInputEventListener) OnEvent(eventObject IAllBinaryEventObject) {
}


func (this *CompleteMotionGestureInputEventListener) OnCompleteMotionGestureInputEvent(completeMotionGestureInputEvent ICompleteMotionGestureInputEvent) {
}

