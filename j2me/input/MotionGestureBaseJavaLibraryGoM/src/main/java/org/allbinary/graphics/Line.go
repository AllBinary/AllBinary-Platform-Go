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

type Line struct {
    P1 IGPoint
    P2 IGPoint
    DeltaX float64
    DeltaY float64
}


func (this *Line) initFields() {
    
    
    
    
}


func NewLine(p1 IGPoint, p2 IGPoint) *Line {
    this := &Line{}
    this.initFields()
        this.P1 = PointFactory_GetInstance().ZERO_ZERO
        this.P2 = this.P1
        this.SetP1(p1)
        this.SetP2(p2)
    return this
}


func (this *Line) GetGradient() float64 {
        return this.GetDeltaY() / this.GetDeltaX()
}


func (this *Line) Update() {
        this.DeltaX = float64(this.GetP1().GetX()) - float64(this.GetP2().GetX())
        this.DeltaY = float64(this.GetP1().GetY()) - float64(this.GetP2().GetY())
}


func (this *Line) GetDeltaX() float64 {
        return this.DeltaX
}


func (this *Line) GetDeltaY() float64 {
        return this.DeltaY
}


func (this *Line) GetP1() IGPoint {
        return this.P1
}


func (this *Line) GetP2() IGPoint {
        return this.P2
}


func (this *Line) SetP1(p1 IGPoint) {
        this.P1 = p1
        this.Update()
}


func (this *Line) SetP2(p2 IGPoint) {
        this.P2 = p2
        this.Update()
}

