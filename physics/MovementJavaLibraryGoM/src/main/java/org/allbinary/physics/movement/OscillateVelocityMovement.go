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

type OscillateVelocityMovement struct {
    BasicConstantVelocityMovement
    OscillateMin int
    OscillateMax int
    Oscillate int
    OscillatePositive bool
}


func (this *OscillateVelocityMovement) initFields() {
    
    
    
    
}


func NewOscillateVelocityMovement(speedBasicDecimal *BasicDecimal, oscillateMin int, oscillateMax int) *OscillateVelocityMovement {
    this := &OscillateVelocityMovement{}
        this.BasicConstantVelocityMovement = *NewBasicConstantVelocityMovement(speedBasicDecimal, NewVelocityProperties(int(speedBasicDecimal.GetUnscaled()), int(speedBasicDecimal.GetUnscaled())))
        this.initFields()
        this.OscillateMin = oscillateMin
        this.OscillateMax = oscillateMax
    return this
}


// throws Exception
func (this *OscillateVelocityMovement) Process(layer IAllBinaryGameLayer) {
        this.BasicConstantVelocityMovement.Process(layer)
        if this.Oscillate < this.OscillateMin {
            this.OscillatePositive = true
        } else if this.Oscillate > this.OscillateMax {
            this.OscillatePositive = false
        }
        if this.OscillatePositive {
            this.Oscillate++
            this.GetVelocityProperties().GetVelocityXBasicDecimalP().Addint(2000)
        } else {
            this.Oscillate--
            this.GetVelocityProperties().GetVelocityXBasicDecimalP().Subtractint(2000)
        }
}

