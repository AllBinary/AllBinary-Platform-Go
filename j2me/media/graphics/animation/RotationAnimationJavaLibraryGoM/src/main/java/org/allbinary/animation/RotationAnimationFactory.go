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

type RotationAnimationFactory struct {
    TotalAngle int16
    AngleIncrement int16
}

// Implements AnimationInterfaceFactoryInterface, ProceduralAnimationInterfaceFactoryInterface


func (this *RotationAnimationFactory) initFields() {
    
    
}


func RotationAnimationFactory_CreateFactory() *RotationAnimationFactory {
        var totalAngle int16 = AngleFactory_GetInstance().TOTAL_ANGLE; _ = totalAngle
        return NewRotationAnimationFactory(totalAngle, int16((int(totalAngle) / GameConfigurationCentral_GetInstance().GetGameControlFidelity())))
}


func NewRotationAnimationFactory(totalAngle int16, angleIncrement int16) *RotationAnimationFactory {
    this := &RotationAnimationFactory{}
    this.initFields()
        this.TotalAngle = totalAngle
        this.AngleIncrement = angleIncrement
    return this
}


// throws Exception
func (this *RotationAnimationFactory) GetInstance(instanceId int) IAnimation {
        return NullRotationAnimation_CreateTotalAngle(AngleInfo_GetInstance(this.AngleIncrement), this.TotalAngle, AnimationBehavior_GetInstance())
}


// throws Exception
func (this *RotationAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        return NullRotationAnimation_CreateTotalAngle(AngleInfo_GetInstance(this.AngleIncrement), this.TotalAngle, AnimationBehavior_GetInstance())
}


func (this *RotationAnimationFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

