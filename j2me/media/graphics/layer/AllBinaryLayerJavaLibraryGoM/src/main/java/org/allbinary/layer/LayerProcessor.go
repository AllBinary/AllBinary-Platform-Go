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

type LayerProcessor struct {
    CommonStrings *CommonStrings
    LayerManager ILayerManager
}

// Implements LayerProcessorInterface

// The Java class LayerProcessor has subclasses so its variables are this interface
type ILayerProcessor interface {
    LayerProcessor_() *LayerProcessor
    ProcessAt(allBinaryLayerManager IAllBinaryLayerManager, layerInterface IAllBinaryLayer, index int)
    IsProcessorLayer(layerInterface IAllBinaryLayer) bool
    GetLayerManager() ILayerManager
    Process(allBinaryLayerManager IAllBinaryLayerManager)
}

func (this *LayerProcessor) LayerProcessor_() *LayerProcessor {
    return this
}


func (this *LayerProcessor) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewLayerProcessor(layerManager ILayerManager) *LayerProcessor {
    this := &LayerProcessor{}
    this.initFields()
        this.LayerManager = layerManager
    return this
}


// throws Exception
func (this *LayerProcessor) ProcessAt(allBinaryLayerManager IAllBinaryLayerManager, layerInterface IAllBinaryLayer, index int) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *LayerProcessor) IsProcessorLayer(layerInterface IAllBinaryLayer) bool {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *LayerProcessor) GetLayerManager() ILayerManager {
        return this.LayerManager
}


// throws Exception
func (this *LayerProcessor) Process(allBinaryLayerManager IAllBinaryLayerManager) {
        var layerManager ILayerManager = this.GetLayerManager(); _ = layerManager
        var size int = layerManager.GetSize(); _ = size
        for index := int(0); index < size; index++ {
            this.ProcessAt(allBinaryLayerManager, any(layerManager.GetLayerAt(index)).(IAllBinaryLayer), index)
        }
}

