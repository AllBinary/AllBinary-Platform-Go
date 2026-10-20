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

type BaseDropCellPositionGenerator struct {
}

// Implements DropCellPositionGeneratorInterface

// The Java class BaseDropCellPositionGenerator has subclasses so its variables are this interface
type IBaseDropCellPositionGenerator interface {
    BaseDropCellPositionGenerator_() *BaseDropCellPositionGenerator
    GetName() string
    Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap)
    IsDropAllowedAt(geographicMapCellPosition IGeographicMapCellPosition) bool
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
}

func (this *BaseDropCellPositionGenerator) BaseDropCellPositionGenerator_() *BaseDropCellPositionGenerator {
    return this
}


func NewBaseDropCellPositionGenerator() *BaseDropCellPositionGenerator {
    this := &BaseDropCellPositionGenerator{}
    return this
}


func (this *BaseDropCellPositionGenerator) GetName() string {
        return reflect.TypeOf(this).String()
}


// throws Exception
func (this *BaseDropCellPositionGenerator) Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap) {
}


// throws Exception
func (this *BaseDropCellPositionGenerator) IsDropAllowedAt(geographicMapCellPosition IGeographicMapCellPosition) bool {
        return true
}


// throws Exception
func (this *BaseDropCellPositionGenerator) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
}

