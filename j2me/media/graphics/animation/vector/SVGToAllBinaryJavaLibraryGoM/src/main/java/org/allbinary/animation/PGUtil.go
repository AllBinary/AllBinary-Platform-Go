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

type PGUtil struct {
}

var PGUtil_Instance *PGUtil = NewPGUtil()

func NewPGUtil() *PGUtil {
    this := &PGUtil{}
    return this
}


func PGUtil_GetInstance() *PGUtil {
        return PGUtil_Instance
}


func (this *PGUtil) FlattenCubicBezier(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, outX []int, outY []int, count []int, maxPoints int) {
        var fixedPoint *FixedPoint = FixedPoint_GetInstance(); _ = fixedPoint
        var steps int = 10; _ = steps
        for i := int(1); i <= steps; i++ {
            var t int = fixedPoint.Div(fixedPoint.ToFixed(float64(i)), fixedPoint.ToFixed(steps)); _ = t
            var mt int = fixedPoint.ONE - t; _ = mt
            var mt3 int = fixedPoint.Mul(fixedPoint.Mul(mt, mt), mt); _ = mt3
            var mt2t3 int = 3 * fixedPoint.Mul(fixedPoint.Mul(mt, mt), t); _ = mt2t3
            var mtt23 int = 3 * fixedPoint.Mul(fixedPoint.Mul(mt, t), t); _ = mtt23
            var t3 int = fixedPoint.Mul(fixedPoint.Mul(t, t), t); _ = t3
            var x int = fixedPoint.Mul(mt3, x0) + fixedPoint.Mul(mt2t3, x1) + fixedPoint.Mul(mtt23, x2) + fixedPoint.Mul(t3, x3); _ = x
            var y int = fixedPoint.Mul(mt3, y0) + fixedPoint.Mul(mt2t3, y1) + fixedPoint.Mul(mtt23, y2) + fixedPoint.Mul(t3, y3); _ = y
            if count[0] < maxPoints {
                outX[count[0]] = x
                outY[count[0]] = y
                count[0]++
            }
        }
}

