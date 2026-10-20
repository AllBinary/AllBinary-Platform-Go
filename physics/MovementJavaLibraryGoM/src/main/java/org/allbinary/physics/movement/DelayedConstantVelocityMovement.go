/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
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

type DelayedConstantVelocityMovement struct {
    BasicConstantVelocityMovement
    TimeDelayHelper ITimeDelayHelper
}


func (this *DelayedConstantVelocityMovement) initFields() {
    this.TimeDelayHelper = NewTimeDelayHelper(60)
}


func NewDelayedConstantVelocityMovement(basicDecimal *BasicDecimal, velocityProperties IBasicVelocityProperties) *DelayedConstantVelocityMovement {
    this := &DelayedConstantVelocityMovement{}
        this.BasicConstantVelocityMovement = *NewBasicConstantVelocityMovement(basicDecimal, velocityProperties)
        this.initFields()
    return this
}


func (this *DelayedConstantVelocityMovement) Init(speedBasicDecimal *BasicDecimal, angle int, otherAngle int) {
}


// throws Exception
func (this *DelayedConstantVelocityMovement) Process(layer IAllBinaryGameLayer) {
        if this.TimeDelayHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
            this.BasicConstantVelocityMovement.Process(layer)
        }
}

