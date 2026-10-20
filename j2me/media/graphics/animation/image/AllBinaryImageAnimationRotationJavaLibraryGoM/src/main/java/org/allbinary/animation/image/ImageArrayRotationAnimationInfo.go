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

type ImageArrayRotationAnimationInfo struct {
    ImageArray []IImage
    AngleIncrement int
    TotalAngle int
    Dx int
    Dy int
}


func (this *ImageArrayRotationAnimationInfo) initFields() {
    
    
    
    
    
}


func NewImageArrayRotationAnimationInfo(imageArray []IImage, angleIncrement int, totalAngle int, dx int, dy int) *ImageArrayRotationAnimationInfo {
    this := &ImageArrayRotationAnimationInfo{}
    this.initFields()
        this.ImageArray = imageArray
        this.AngleIncrement = angleIncrement
        this.TotalAngle = totalAngle
        this.Dx = dx
        this.Dy = dy
    return this
}


func (this *ImageArrayRotationAnimationInfo) SetImageArray(imageArray []IImage) {
        this.ImageArray = imageArray
}


func (this *ImageArrayRotationAnimationInfo) GetImageArray() []IImage {
        return this.ImageArray
}


func (this *ImageArrayRotationAnimationInfo) GetAngleInfoP() *AngleInfo {
        return AngleInfo_GetInstance(int16(this.AngleIncrement))
}


func (this *ImageArrayRotationAnimationInfo) SetTotalAngle(totalAngle int) {
        this.TotalAngle = totalAngle
}


func (this *ImageArrayRotationAnimationInfo) GetTotalAngle() int {
        return this.TotalAngle
}


func (this *ImageArrayRotationAnimationInfo) SetDx(dx int) {
        this.Dx = dx
}


func (this *ImageArrayRotationAnimationInfo) GetDx() int {
        return this.Dx
}


func (this *ImageArrayRotationAnimationInfo) SetDy(dy int) {
        this.Dy = dy
}


func (this *ImageArrayRotationAnimationInfo) GetDy() int {
        return this.Dy
}

