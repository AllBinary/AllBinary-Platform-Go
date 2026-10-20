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

type VectorAnimationInterfaceFactory struct {
    CurrentPoints [][][]int
    BasicColor *BasicColor
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements AnimationInterfaceFactoryInterface


func (this *VectorAnimationInterfaceFactory) initFields() {
    
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    
}


func NewVectorAnimationInterfaceFactory(currentPoints [][][]int, basicColor *BasicColor, animationBehaviorFactory IAnimationBehaviorFactory) *VectorAnimationInterfaceFactory {
    this := &VectorAnimationInterfaceFactory{}
    this.initFields()
        this.CurrentPoints = currentPoints
        this.SetBasicColorP(basicColor)
        this.AnimationBehaviorFactory = animationBehaviorFactory
    return this
}


// throws Exception
func (this *VectorAnimationInterfaceFactory) GetInstance(instanceId int) IAnimation {
        return NewVectorAnimation(this.CurrentPoints, this.GetBasicColorP(), this.AnimationBehaviorFactory.GetOrCreateInstance())
}


func (this *VectorAnimationInterfaceFactory) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}


func (this *VectorAnimationInterfaceFactory) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *VectorAnimationInterfaceFactory) SetInitialScale(scaleProperties *ScaleProperties) {
}

