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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type UpDownVectorAI struct {
    BasicAI
    VelocityInterface IBasicVelocityProperties
    Index int
}


func (this *UpDownVectorAI) initFields() {
    
    this.Index = 0
}


func NewUpDownVectorAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *UpDownVectorAI {
    this := &UpDownVectorAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        var velocityInterfaceCompositeInterface VelocityInterfaceCompositeInterface = any(this.GetOwnerLayerInterface()).(VelocityInterfaceCompositeInterface); _ = velocityInterfaceCompositeInterface
        this.VelocityInterface = velocityInterfaceCompositeInterface.GetVelocityProperties()
        this.VelocityInterface.GetVelocityYBasicDecimalP().Setint(0)
    return this
}


// throws Exception
func (this *UpDownVectorAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var y int = this.GetOwnerLayerInterface().GetHeight() * 270; _ = y
        if this.Index < 5 {
            var basicDecimal *BasicDecimal = this.VelocityInterface.GetVelocityYBasicDecimalP(); _ = basicDecimal
            basicDecimal.Setint(0)
            basicDecimal.Addint(y)
            this.Index++
        } else if this.Index < 10 {
            var basicDecimal *BasicDecimal = this.VelocityInterface.GetVelocityYBasicDecimalP(); _ = basicDecimal
            basicDecimal.Setint(0)
            basicDecimal.Subtractint(y)
            this.Index++
        } else {
            this.Index = 0
            this.ProcessAI(allBinaryLayerManager)
        }
}

