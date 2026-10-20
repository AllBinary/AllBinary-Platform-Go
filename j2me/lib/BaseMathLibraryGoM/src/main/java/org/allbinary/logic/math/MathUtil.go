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

type MathUtil struct {
    ACCURACY int
}

var MathUtil_Instance *MathUtil = NewMathUtil()

func (this *MathUtil) initFields() {
    this.ACCURACY = 100
}

func NewMathUtil() *MathUtil {
    this := &MathUtil{}
    this.initFields()
    return this
}


func MathUtil_GetInstance() *MathUtil {
        return MathUtil_Instance
}


func (this *MathUtil) GetTotalDigits(digits int) int {
        var total int = 0; _ = total
        for digits > 0 {
            digits = digits / 10
            total++
        }
        return total
}


func (this *MathUtil) Sqrt(value int) int {
        if value <= 1 {
            return value
        }
        var accumulated int = value >> 1; _ = accumulated
        var result int = (accumulated + (value / accumulated)) >> 1; _ = result
        for {
            accumulated = result
            result = (accumulated + (value / accumulated)) >> 1
            if !(accumulated > result) {
                break
            }
        }
        return accumulated
}


func (this *MathUtil) Sqrtd(x int) float64 {
        if x == 0 {
            return float64(0.0)
        }
        if x < 0 {
            return float64(-1.0)
        }
        var result float64 = float64(x); _ = result
        for index := int(0); index < this.ACCURACY; index++ {
            result = (result + (float64(x) / result)) / 2.0
        }
        return result
}


func (this *MathUtil) Abs(value int) int {
        return func() int { if (value < 0) { return -value }; return value }()
}


func (this *MathUtil) Abslong(value int64) int64 {
        return func() int64 { if (value < 0) { return -value }; return value }()
}


func (this *MathUtil) Min(value int, value2 int) int {
        return func() int { if (value <= value2) { return value }; return value2 }()
}


func (this *MathUtil) Max(value int, value2 int) int {
        return func() int { if (value >= value2) { return value }; return value2 }()
}
