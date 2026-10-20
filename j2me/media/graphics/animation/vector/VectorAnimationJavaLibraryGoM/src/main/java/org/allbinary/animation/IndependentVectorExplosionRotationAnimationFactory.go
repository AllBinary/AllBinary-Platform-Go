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

type IndependentVectorExplosionRotationAnimationFactory struct {
    FramePoints [][][]int
    VectorInfo *VectorInfo
    BasicColor *BasicColor
    VectorAnimationFactoryInterface VectorAnimationFactoryInterface
    VectorExplosionGenerator *VectorExplosionGenerator
}

// Implements ProceduralAnimationInterfaceFactoryInterface


func (this *IndependentVectorExplosionRotationAnimationFactory) initFields() {
    this.FramePoints = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY_ARRAY
    
    
    
    this.VectorExplosionGenerator = VectorExplosionGenerator_GetInstance()
}


func NewIndependentVectorExplosionRotationAnimationFactory(vectorInfo *VectorInfo, basicColor *BasicColor, vectorAnimationFactoryInterface VectorAnimationFactoryInterface) *IndependentVectorExplosionRotationAnimationFactory {
    this := &IndependentVectorExplosionRotationAnimationFactory{}
    this.initFields()
        this.VectorInfo = vectorInfo
        this.BasicColor = basicColor
        this.VectorAnimationFactoryInterface = vectorAnimationFactoryInterface
        this.Init()
    return this
}


// throws Exception
func (this *IndependentVectorExplosionRotationAnimationFactory) Init() {
        this.FramePoints = VectorRotationGenerator_GetInstance().GetInstance(this.VectorInfo)
}


// throws Exception
func (this *IndependentVectorExplosionRotationAnimationFactory) GetInstanceAnimation(animationInterface IAnimation) IAnimation {
        var frame int = 0; _ = frame
        if animationInterface != nil {
            var rotationAnimationInterface IIndexedAnimation = any(animationInterface).(IIndexedAnimation); _ = rotationAnimationInterface
            frame = rotationAnimationInterface.GetFrame()
        }
        var points [][][]int = this.VectorExplosionGenerator.GetInstance(this.FramePoints[frame], 6, this.VectorExplosionGenerator.ROTATION); _ = points
        return this.VectorAnimationFactoryInterface.GetInstance(points, this.BasicColor)
}

