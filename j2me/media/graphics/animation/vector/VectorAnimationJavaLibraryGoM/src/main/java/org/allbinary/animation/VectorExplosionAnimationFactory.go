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

type VectorExplosionAnimationFactory struct {
    VectorExplosionGenerator *VectorExplosionGenerator
    BasicColor *BasicColor
    AnimationBehaviorFactory IAnimationBehaviorFactory
}

// Implements ProceduralAnimationInterfaceFactoryInterface


func (this *VectorExplosionAnimationFactory) initFields() {
    this.VectorExplosionGenerator = VectorExplosionGenerator_GetInstance()
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    
}


func NewVectorExplosionAnimationFactory(basicColor *BasicColor, animationBehaviorFactory IAnimationBehaviorFactory) *VectorExplosionAnimationFactory {
    this := &VectorExplosionAnimationFactory{}
    this.initFields()
        this.AnimationBehaviorFactory = animationBehaviorFactory
        this.SetBasicColorP(basicColor)
    return this
}


// throws Exception
func (this *VectorExplosionAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        var vectorRotationAnimationInterface VectorAnimationInterface = any(animationInterface).(VectorAnimationInterface); _ = vectorRotationAnimationInterface
        var frame int = vectorRotationAnimationInterface.GetFrame(); _ = frame
        var framePoints [][]int = vectorRotationAnimationInterface.GetPoints(frame); _ = framePoints
        var points [][][]int = this.VectorExplosionGenerator.GetInstance(framePoints, 6, this.VectorExplosionGenerator.ROTATION); _ = points
        return NewVectorRotationAnimation(points, this.GetBasicColorP(), this.AnimationBehaviorFactory.GetOrCreateInstance())
}


func (this *VectorExplosionAnimationFactory) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}


func (this *VectorExplosionAnimationFactory) GetBasicColorP() *BasicColor {
        return this.BasicColor
}

