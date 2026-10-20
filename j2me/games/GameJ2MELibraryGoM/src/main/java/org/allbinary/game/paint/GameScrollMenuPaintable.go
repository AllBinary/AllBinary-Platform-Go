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

type GameScrollMenuPaintable struct {
    BasicGameDemoPaintable
}

// Implements ColorChangeListener



func NewGameScrollMenuPaintable(mainDemoStatePaintable IPaintable, ownershipPaintable IPaintable, helpPaintableInterface IPaintable, basicColor *BasicColor) *GameScrollMenuPaintable {
    this := &GameScrollMenuPaintable{}
        this.BasicGameDemoPaintable = *NewBasicGameDemoPaintable(mainDemoStatePaintable, ownershipPaintable, helpPaintableInterface)
        this.SetColor(basicColor)
        ColorChangeEventHandler_GetInstance().AddListenerInterface(this)
    return this
}


func (this *GameScrollMenuPaintable) SetState(state int) {
        this.SetCurrentStatePaintable(this.GetMainDemoStatePaintable())
}


func (this *GameScrollMenuPaintable) OnEvent(eventObject IAllBinaryEventObject) {
}


func (this *GameScrollMenuPaintable) SetColor(basicColor *BasicColor) {
        var mainGameDemoStatePaintable *MainGameDemoStatePaintable = any(this.GetMainDemoStatePaintable()).(*MainGameDemoStatePaintable); _ = mainGameDemoStatePaintable
        var formPaintable *FormPaintable = any(mainGameDemoStatePaintable.GetMenuPaintableInterface()).(*FormPaintable); _ = formPaintable
        var scrollSelectionForm IScrollSelectionForm = any(formPaintable.GetForm()).(IScrollSelectionForm); _ = scrollSelectionForm
        mainGameDemoStatePaintable.GetOwnershipPaintableInterface().SetBasicColorP(basicColor)
        scrollSelectionForm.SetButtonBasicColor(basicColor)
        var item IABTextItem; _ = item
        var stringComponent *ABStringComponent; _ = stringComponent
        for index := int(scrollSelectionForm.Size() - 1); index >= 0; index-- {
            item = any(scrollSelectionForm.Get(index)).(IABTextItem)
            stringComponent = item.GetLabelStringComponent()
            if stringComponent != nil {
                stringComponent.SetForegroundBasicColor(basicColor)
            }
        }
}

