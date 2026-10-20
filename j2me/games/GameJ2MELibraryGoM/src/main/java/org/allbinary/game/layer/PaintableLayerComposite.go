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

type PaintableLayerComposite struct {
    Paintable
    PaintableArray []ILayer
}


func (this *PaintableLayerComposite) initFields() {
    
}


func NewPaintableLayerComposite(paintableArray []ILayer) *PaintableLayerComposite {
    this := &PaintableLayerComposite{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.PaintableArray = paintableArray
    return this
}


func (this *PaintableLayerComposite) Paint(graphics IGraphics) {
        for index := int(len(this.PaintableArray)); Java_PreDecrement(&index) >= 0;  {
            this.PaintableArray[index].Paint(graphics)
        }
}


func (this *PaintableLayerComposite) PaintThreed(graphics IGraphics) {
        var paintableInterface PaintableInterface; _ = paintableInterface
        for index := int(len(this.PaintableArray)); Java_PreDecrement(&index) >= 0;  {
            paintableInterface = any(this.PaintableArray[index]).(PaintableInterface)
            paintableInterface.PaintThreed(graphics)
        }
}

