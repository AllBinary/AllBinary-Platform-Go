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

type DestroyedEvent struct {
    AllBinaryEventObject
    AllBinaryGameCanvas IAllBinaryGameCanvas
    AllBinaryLayerInterface IAllBinaryLayer
}


func (this *DestroyedEvent) initFields() {
    this.AllBinaryGameCanvas = NullGameCanvas_GetInstance()
    this.AllBinaryLayerInterface = AllBinaryLayer_NULL_ALLBINARY_LAYER
}


func NewDestroyedEvent(combatGameCanvas IAllBinaryGameCanvas) *DestroyedEvent {
    this := &DestroyedEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(DestroyedEventHandler_GetInstance())
        this.initFields()
        this.SetAllBinaryGameCanvas(combatGameCanvas)
    return this
}


func (this *DestroyedEvent) GetAllBinaryGameCanvas() IAllBinaryGameCanvas {
        return this.AllBinaryGameCanvas
}


func (this *DestroyedEvent) GetLayerInterface() IAllBinaryLayer {
        return this.AllBinaryLayerInterface
}


func (this *DestroyedEvent) SetAllBinaryGameCanvas(allBinaryGameCanvas IAllBinaryGameCanvas) {
        this.AllBinaryGameCanvas = allBinaryGameCanvas
}


func (this *DestroyedEvent) SetLayerInterfaceForCircularStaticPool(allBinaryLayerInterface IAllBinaryLayer) {
        this.AllBinaryLayerInterface = allBinaryLayerInterface
}


func (this *DestroyedEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("DestroyedEvent: \n")
        stringBuffer.Append("AllBinaryLayerInterface: ")
        stringBuffer.Append(fmt.Sprint(this.AllBinaryLayerInterface))
        stringBuffer.Append("\nAllBinaryGameCanvas: ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.GetAllBinaryGameCanvas()))
        return fmt.Sprint(stringBuffer)
}

func (this *DestroyedEvent) String() string {
    return this.ToString()
}

