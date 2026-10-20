/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type SpecialGameInput struct {
}

// Implements SpecialGameInputInterface

// The Java class SpecialGameInput has subclasses so its variables are this interface
type ISpecialGameInput interface {
    SpecialGameInput_() *SpecialGameInput
    Up()
    Down()
    Right()
    Left()
    StrafeLeft()
    StrafeRight()
    Fire(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special1(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special2(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special3(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special4(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Special5(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    Paint(graphics IGraphics)
    PaintThreed(graphics IGraphics)
}

func (this *SpecialGameInput) SpecialGameInput_() *SpecialGameInput {
    return this
}


func NewSpecialGameInput() *SpecialGameInput {
    this := &SpecialGameInput{}
    return this
}


// throws Exception
func (this *SpecialGameInput) Up() {
}


// throws Exception
func (this *SpecialGameInput) Down() {
}


// throws Exception
func (this *SpecialGameInput) Right() {
}


// throws Exception
func (this *SpecialGameInput) Left() {
}


// throws Exception
func (this *SpecialGameInput) StrafeLeft() {
}


// throws Exception
func (this *SpecialGameInput) StrafeRight() {
}


// throws Exception
func (this *SpecialGameInput) Fire(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *SpecialGameInput) Special1(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *SpecialGameInput) Special2(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *SpecialGameInput) Special3(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *SpecialGameInput) Special4(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *SpecialGameInput) Special5(layerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


func (this *SpecialGameInput) Paint(graphics IGraphics) {
}


func (this *SpecialGameInput) PaintThreed(graphics IGraphics) {
}

