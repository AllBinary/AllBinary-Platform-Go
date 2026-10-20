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

type NullRotationAnimation struct {
    RotationAnimation
}

// The Java class NullRotationAnimation has subclasses so its variables are this interface
type INullRotationAnimation interface {
    IRotationAnimation
    NullRotationAnimation_() *NullRotationAnimation
    GetAnimationSize() int
}

func (this *NullRotationAnimation) NullRotationAnimation_() *NullRotationAnimation {
    return this
}



func NullRotationAnimation_CreateTotalAngle(angleInfo *AngleInfo, totalAngle int16, animationBehavior IAnimationBehavior) INullRotationAnimation {
        return NewNullRotationAnimation(angleInfo, CircularIndexUtil_CreateInstance(int(totalAngle) / int(angleInfo.GetAngleIncrementInfo().GetAngleIncrement())), animationBehavior)
}


func NullRotationAnimation_CreateQuarter(animationBehavior IAnimationBehavior) INullRotationAnimation {
        return NewNullRotationAnimation(AngleInfo_GetInstance(AngleFactory_GetInstance().QUARTER_TOTAL_ANGLE), CircularIndexUtil_CreateInstance(4), animationBehavior)
}


func NewNullRotationAnimation(angleInfo *AngleInfo, circularIndexUtil *CircularIndexUtil, animationBehavior IAnimationBehavior) *NullRotationAnimation {
    this := &NullRotationAnimation{}
        this.RotationAnimation = *NewRotationAnimation(angleInfo, circularIndexUtil, animationBehavior)
    return this
}


// throws Exception
func (this *NullRotationAnimation) GetAnimationSize() int {
        return 0
}

