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

type MainGameDemoStatePaintable struct {
    Paintable
    OwnershipPaintableInterface IPaintable
    MenuPaintableInterface IPaintable
}


func (this *MainGameDemoStatePaintable) initFields() {
    
    
}


func NewMainGameDemoStatePaintable(ownershipPaintableInterface IPaintable, menuPaintableInterface IPaintable) *MainGameDemoStatePaintable {
    this := &MainGameDemoStatePaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
        this.OwnershipPaintableInterface = ownershipPaintableInterface
        this.MenuPaintableInterface = menuPaintableInterface
    return this
}


func (this *MainGameDemoStatePaintable) Paint(graphics IGraphics) {
        this.MenuPaintableInterface.Paint(graphics)
        this.OwnershipPaintableInterface.Paint(graphics)
}


func (this *MainGameDemoStatePaintable) GetOwnershipPaintableInterface() IPaintable {
        return this.OwnershipPaintableInterface
}


func (this *MainGameDemoStatePaintable) GetMenuPaintableInterface() IPaintable {
        return this.MenuPaintableInterface
}

