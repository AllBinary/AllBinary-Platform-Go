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

type BasicLayerProcessor struct {
    LogUtil *LogUtil
    List IBasicArrayList
}

// The Java class BasicLayerProcessor has subclasses so its variables are this interface
type IBasicLayerProcessor interface {
    BasicLayerProcessor_() *BasicLayerProcessor
    Add(layerInterface IAllBinaryLayer)
    Process(allBinaryLayerManager IAllBinaryLayerManager)
    GetList() IBasicArrayList
}

func (this *BasicLayerProcessor) BasicLayerProcessor_() *BasicLayerProcessor {
    return this
}

var BasicLayerProcessor_NULL_LAYER_PROCESSOR IBasicLayerProcessor = NewBasicLayerProcessor()

func (this *BasicLayerProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.List = NewBasicArrayListD()
}


func NewBasicLayerProcessor() *BasicLayerProcessor {
    this := &BasicLayerProcessor{}
    this.initFields()
    return this
}


func (this *BasicLayerProcessor) Add(layerInterface IAllBinaryLayer) {
        if !this.List.Contains(layerInterface) {
            this.List.Add(layerInterface)
        }
}


// throws Exception
func (this *BasicLayerProcessor) Process(allBinaryLayerManager IAllBinaryLayerManager) {
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
}


func (this *BasicLayerProcessor) GetList() IBasicArrayList {
        return this.List
}

