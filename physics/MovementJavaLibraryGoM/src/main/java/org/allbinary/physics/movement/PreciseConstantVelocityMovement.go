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

type PreciseConstantVelocityMovement struct {
    BasicConstantVelocityMovement
    AccumulatedX int
    AccumulatedY int
    AccumulatedZ int
    FactorValue int
}


func (this *PreciseConstantVelocityMovement) initFields() {
    
    
    
    this.FactorValue = ScaleFactorFactory_GetInstance().DEFAULT_SCALE_VALUE
}


func NewPreciseConstantVelocityMovement(basicDecimal *BasicDecimal, velocityProperties IBasicVelocityProperties) *PreciseConstantVelocityMovement {
    this := &PreciseConstantVelocityMovement{}
        this.BasicConstantVelocityMovement = *NewBasicConstantVelocityMovement(basicDecimal, velocityProperties)
        this.initFields()
    return this
}


func (this *PreciseConstantVelocityMovement) Init(speedBasicDecimal *BasicDecimal, angle int, otherAngle int) {
}


// throws Exception
func (this *PreciseConstantVelocityMovement) Process(layer IAllBinaryGameLayer) {
        var mathUtil *MathUtil = MathUtil_GetInstance(); _ = mathUtil
        var velocityProperties IBasicVelocityProperties = this.GetVelocityProperties(); _ = velocityProperties
        this.AccumulatedX = int((int64(this.AccumulatedX) + velocityProperties.GetVelocityXBasicDecimalP().GetUnscaled()))
        this.AccumulatedY = int((int64(this.AccumulatedY) + velocityProperties.GetVelocityYBasicDecimalP().GetUnscaled()))
        this.AccumulatedZ = int((int64(this.AccumulatedZ) + velocityProperties.GetVelocityZBasicDecimalP().GetUnscaled()))
        layer.MoveDXYZ(this.AccumulatedX / this.FactorValue, this.AccumulatedY / this.FactorValue, this.AccumulatedZ / this.FactorValue)
        if mathUtil.Abs(this.AccumulatedX) > this.FactorValue {
            this.AccumulatedX = 0
        }
        if mathUtil.Abs(this.AccumulatedY) > this.FactorValue {
            this.AccumulatedY = 0
        }
        if mathUtil.Abs(this.AccumulatedZ) > this.FactorValue {
            this.AccumulatedZ = 0
        }
}

