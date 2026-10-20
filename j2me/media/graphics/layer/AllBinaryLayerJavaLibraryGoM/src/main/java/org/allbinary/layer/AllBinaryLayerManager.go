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

type AllBinaryLayerManager struct {
    LayerManager
    LayerManagerEventHandler *LayerManagerEventHandler
    CreateLayerManagerEvent *LayerManagerEvent
    DeleteLayerManagerEvent *LayerManagerEvent
    BasicLayerProcessorArray []ILayerProcessor
}

// The Java class AllBinaryLayerManager has subclasses so its variables are this interface
type IAllBinaryLayerManager interface {
    ILayerManager
    AllBinaryLayerManager_() *AllBinaryLayerManager
    GetLayerProcessorArray() []ILayerProcessor
    SetLayerProcessorArray(layerProcessorArray []ILayerProcessor)
    Log()
    Append(layerInterface IAllBinaryLayer)
    AppendAt(layerInterface IAllBinaryLayer, index int)
    AppendProcessors(layerInterface IAllBinaryLayer)
    Remove(layerInterface IAllBinaryLayer)
    Process()
    Cleanup()
}

func (this *AllBinaryLayerManager) AllBinaryLayerManager_() *AllBinaryLayerManager {
    return this
}


func (this *AllBinaryLayerManager) initFields() {
    this.LayerManagerEventHandler = LayerManagerEventHandler_GetInstance()
    this.CreateLayerManagerEvent = NewLayerManagerEvent(this, this.LayerManagerEventHandler.CREATE)
    this.DeleteLayerManagerEvent = NewLayerManagerEvent(this, this.LayerManagerEventHandler.DELETE)
    this.BasicLayerProcessorArray = make([]ILayerProcessor, 0)
}


func NewAllBinaryLayerManager() *AllBinaryLayerManager {
    this := &AllBinaryLayerManager{}
        this.LayerManager = *NewLayerManager(LayerManagerNoDebug_GetInstance())
        this.initFields()
    return this
}


func (this *AllBinaryLayerManager) GetLayerProcessorArray() []ILayerProcessor {
        return this.BasicLayerProcessorArray
}


func (this *AllBinaryLayerManager) SetLayerProcessorArray(layerProcessorArray []ILayerProcessor) {
        this.BasicLayerProcessorArray = layerProcessorArray
}


func (this *AllBinaryLayerManager) Log() {
        var size int = len(this.BasicLayerProcessorArray); _ = size
        for index := int(0); index < size; index++ {
            var layerProcessorInterface ILayerProcessor = this.BasicLayerProcessorArray[index]; _ = layerProcessorInterface
        }
}


// throws Exception
func (this *AllBinaryLayerManager) Append(layerInterface IAllBinaryLayer) {
        if !this.Contains(layerInterface) {
            this.AppendProcessors(layerInterface)
            this.LayerManager.Append(layerInterface)
        }
}


// throws Exception
func (this *AllBinaryLayerManager) AppendAt(layerInterface IAllBinaryLayer, index int) {
        if !this.Contains(layerInterface) {
            this.AppendProcessors(layerInterface)
            this.LayerManager.AppendAt(layerInterface, index)
        }
}


// throws Exception
func (this *AllBinaryLayerManager) AppendProcessors(layerInterface IAllBinaryLayer) {
        this.CreateLayerManagerEvent.SetLayerInterface(layerInterface)
        this.LayerManagerEventHandler.FireEvent(this.CreateLayerManagerEvent)
        var layerProcessorInterface ILayerProcessor; _ = layerProcessorInterface
        for index := int(len(this.BasicLayerProcessorArray)); Java_PreDecrement(&index) >= 0;  {
            layerProcessorInterface = this.BasicLayerProcessorArray[index]
            if layerProcessorInterface.IsProcessorLayer(layerInterface) {
                layerProcessorInterface.GetLayerManager().Append(layerInterface)
            }
        }
}


// throws Exception
// synchronized
func (this *AllBinaryLayerManager) Remove(layerInterface IAllBinaryLayer) {
        this.DeleteLayerManagerEvent.SetLayerInterface(layerInterface)
        this.LayerManagerEventHandler.FireDeleteEvent(this.DeleteLayerManagerEvent)
        var layerProcessorInterface ILayerProcessor; _ = layerProcessorInterface
        for index := int(len(this.BasicLayerProcessorArray)); Java_PreDecrement(&index) >= 0;  {
            layerProcessorInterface = this.BasicLayerProcessorArray[index]
            layerProcessorInterface.GetLayerManager().Remove(layerInterface)
        }
        this.LayerManager.Remove(layerInterface)
}


// throws Exception
func (this *AllBinaryLayerManager) Process() {
        var layerProcessorInterface ILayerProcessor; _ = layerProcessorInterface
        var size int = len(this.BasicLayerProcessorArray); _ = size
        for index := int(0); index < size; index++ {
            layerProcessorInterface = this.BasicLayerProcessorArray[index]
            layerProcessorInterface.Process(this)
        }
}


// throws Exception
func (this *AllBinaryLayerManager) Cleanup() {
        var layerProcessorInterface ILayerProcessor; _ = layerProcessorInterface
        var size int = len(this.BasicLayerProcessorArray); _ = size
        for index := int(0); index < size; index++ {
            layerProcessorInterface = this.BasicLayerProcessorArray[index]
            layerProcessorInterface.GetLayerManager().Cleanup()
        }
        this.LayerManager.Cleanup()
}

