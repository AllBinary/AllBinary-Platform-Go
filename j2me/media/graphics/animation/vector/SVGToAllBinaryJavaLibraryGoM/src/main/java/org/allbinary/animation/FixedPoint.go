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

type FixedPoint struct {
    ONE int
}

var FixedPoint_Instance *FixedPoint = NewFixedPoint()

func (this *FixedPoint) initFields() {
    this.ONE = 1 << 16
}

func NewFixedPoint() *FixedPoint {
    this := &FixedPoint{}
    this.initFields()
    return this
}


func FixedPoint_GetInstance() *FixedPoint {
        return FixedPoint_Instance
}


func (this *FixedPoint) ToFixed_1(val float64) int {
        return int((val * 65536.0))
}

// Java overloads of FixedPoint.toFixed
func (this *FixedPoint) ToFixed(args ...any) int {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[float64](args[0], pass) {
            return this.ToFixed_1(Java_Arg[float64](args[0]))
        }
        if len(args) == 1 && Java_IsArg[float32](args[0], pass) {
            return this.ToFixed_2(Java_Arg[float32](args[0]))
        }
    }
    panic("No Java overload of FixedPoint.toFixed for the arguments")
}


func (this *FixedPoint) ToFixed_2(val float32) int {
        return int((val * 65536.0))
}


func (this *FixedPoint) Mul(a int, b int) int {
        return int(((int64(a) * int64(b)) >> 16))
}


func (this *FixedPoint) Div(a int, b int) int {
        return int((((int64(a)) << 16) / int64(b)))
}


func (this *FixedPoint) ToInt(fixed int) int {
        return fixed >> 16
}

