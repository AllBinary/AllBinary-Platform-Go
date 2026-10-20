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

type LayerManagerEvent struct {
    AllBinaryEventObject
    LayerInterface IAllBinaryLayer
    Crud string
}


func (this *LayerManagerEvent) initFields() {
    this.LayerInterface = AllBinaryLayer_NULL_ALLBINARY_LAYER
    
}


func NewLayerManagerEvent(layerManager ILayerManager, crud string) *LayerManagerEvent {
    this := &LayerManagerEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(layerManager)
        this.initFields()
        this.Crud = crud
    return this
}


func (this *LayerManagerEvent) GetLayerInterface() IAllBinaryLayer {
        return this.LayerInterface
}


func (this *LayerManagerEvent) SetLayerInterface(layerInterface IAllBinaryLayer) {
        this.LayerInterface = layerInterface
}


func (this *LayerManagerEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("LayerManagerEvent: \n")
        stringBuffer.Append("LayerInterface: ")
        stringBuffer.Append(fmt.Sprint(this.GetLayerInterface()))
        stringBuffer.Append("\nCRUD: ")
        stringBuffer.Append(this.Crud)
        return fmt.Sprint(stringBuffer)
}

func (this *LayerManagerEvent) String() string {
    return this.ToString()
}

