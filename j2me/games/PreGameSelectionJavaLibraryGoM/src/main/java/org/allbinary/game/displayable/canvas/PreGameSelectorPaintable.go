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

type PreGameSelectorPaintable struct {
    InitUpdatePaintable
    Title string
    CanvasDrawLineString *CanvasDrawLineString
    PreGameScrollSelectionForm *PreGameScrollSelectionForm
}


func (this *PreGameSelectorPaintable) initFields() {
    
    
    
}


func NewPreGameSelectorPaintable(title string, preGameScrollSelectionForm *PreGameScrollSelectionForm) *PreGameSelectorPaintable {
    this := &PreGameSelectorPaintable{}
    this.InitUpdatePaintable = *NewInitUpdatePaintable()
    this.initFields()
        this.Title = title
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        this.CanvasDrawLineString = NewCanvasDrawLineString(displayInfo.GetLastHalfWidth(), 0)
        this.PreGameScrollSelectionForm = preGameScrollSelectionForm
    return this
}


func (this *PreGameSelectorPaintable) Update() {
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        this.CanvasDrawLineString = NewCanvasDrawLineString(displayInfo.GetLastHalfWidth(), 0)
}


func (this *PreGameSelectorPaintable) Paint(graphics IGraphics) {
        graphics.SetColor(BasicColorFactory_GetInstance().WHITE.IntValue())
        this.CanvasDrawLineString.Paint(graphics, this.Title, 0)
        this.GetPreGameSelectionForm().Paint(graphics)
}


func (this *PreGameSelectorPaintable) GetPreGameSelectionForm() *PreGameScrollSelectionForm {
        return this.PreGameScrollSelectionForm
}

