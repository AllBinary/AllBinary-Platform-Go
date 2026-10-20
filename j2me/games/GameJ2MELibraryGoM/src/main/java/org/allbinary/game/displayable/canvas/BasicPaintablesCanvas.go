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

type BasicPaintablesCanvas struct {
    GameCommandCanvas
    PaintableArray []IPaintable
    ColorFillPaintable IColorFillBasePaintable
}


func (this *BasicPaintablesCanvas) initFields() {
    
    
}


func NewBasicPaintablesCanvas(cmdListener CommandListener, name string, allBinaryGameLayerManager IAllBinaryGameLayerManager, paintableArray []IPaintable) *BasicPaintablesCanvas {
    this := &BasicPaintablesCanvas{}
        this.GameCommandCanvas = *NewGameCommandCanvas(cmdListener, name, allBinaryGameLayerManager.GetBackgroundBasicColor(), allBinaryGameLayerManager.GetForegroundBasicColor())
        this.initFields()
        this.PaintableArray = paintableArray
        this.ColorFillPaintable = ColorFillPaintableFactory_GetInstance().GetInstance(allBinaryGameLayerManager.GetBackgroundBasicColor(), false)
    return this
}


func (this *BasicPaintablesCanvas) Paint(graphics IGraphics) {
        this.ColorFillPaintable.Paint(graphics)
        graphics.SetColor(this.ForegroundColor)
        for index := int(len(this.PaintableArray) - 1); index >= 0; index-- {
            this.PaintableArray[index].Paint(graphics)
        }
        this.GameCommandCanvas.Paint(graphics)
}

