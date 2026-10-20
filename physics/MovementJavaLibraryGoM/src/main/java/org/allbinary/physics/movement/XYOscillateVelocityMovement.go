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

type XYOscillateVelocityMovement struct {
    BasicConstantVelocityMovement
    OscillateMin int
    OscillateMax int
    Oscillate int
    OscillatePositive bool
}


func (this *XYOscillateVelocityMovement) initFields() {
    
    
    
    
}


func NewXYOscillateVelocityMovement(speedBasicDecimal *BasicDecimal, oscillateMin int, oscillateMax int) *XYOscillateVelocityMovement {
    this := &XYOscillateVelocityMovement{}
        this.BasicConstantVelocityMovement = *NewBasicConstantVelocityMovement(speedBasicDecimal, NewVelocityProperties(int(speedBasicDecimal.GetUnscaled()), int(speedBasicDecimal.GetUnscaled())))
        this.initFields()
        this.OscillateMin = oscillateMin
        this.OscillateMax = oscillateMax
    return this
}


// throws Exception
func (this *XYOscillateVelocityMovement) Process(layer IAllBinaryGameLayer) {
        this.BasicConstantVelocityMovement.Process(layer)
        if this.Oscillate < this.OscillateMin {
            this.OscillatePositive = true
            this.GetVelocityProperties().GetVelocityYBasicDecimalP().Addint(3000)
        } else if this.Oscillate > this.OscillateMax {
            this.OscillatePositive = false
            this.GetVelocityProperties().GetVelocityYBasicDecimalP().Subtractint(3000)
        }
        if this.OscillatePositive {
            this.Oscillate++
            this.GetVelocityProperties().GetVelocityXBasicDecimalP().Addint(5000)
        } else {
            this.Oscillate--
            this.GetVelocityProperties().GetVelocityXBasicDecimalP().Subtractint(5000)
        }
}

