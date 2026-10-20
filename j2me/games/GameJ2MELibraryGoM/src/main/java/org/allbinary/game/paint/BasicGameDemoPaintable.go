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

type BasicGameDemoPaintable struct {
    StatePaintable
    MainDemoStatePaintable IPaintable
    OwnershipPaintable IPaintable
    HelpPaintableInterface IPaintable
    CurrentStatePaintable IPaintable
}

// The Java class BasicGameDemoPaintable has subclasses so its variables are this interface
type IBasicGameDemoPaintable interface {
    IStatePaintable
    BasicGameDemoPaintable_() *BasicGameDemoPaintable
    SetState(state int)
    Paint(graphics IGraphics)
    GetHelpPaintableInterface() IPaintable
    SetCurrentStatePaintable(currentStatePaintable IPaintable)
    GetCurrentStatePaintable() IPaintable
    GetMainDemoStatePaintable() IPaintable
}

func (this *BasicGameDemoPaintable) BasicGameDemoPaintable_() *BasicGameDemoPaintable {
    return this
}


func (this *BasicGameDemoPaintable) initFields() {
    
    
    
    this.CurrentStatePaintable = NullPaintable_GetInstance()
}


func NewBasicGameDemoPaintable(mainDemoStatePaintable IPaintable, ownershipPaintable IPaintable, helpPaintableInterface IPaintable) *BasicGameDemoPaintable {
    this := &BasicGameDemoPaintable{}
    this.StatePaintable = *NewStatePaintable()
    this.initFields()
        this.MainDemoStatePaintable = mainDemoStatePaintable
        this.SetCurrentStatePaintable(NullPaintable_GetInstance())
        this.OwnershipPaintable = ownershipPaintable
        this.HelpPaintableInterface = helpPaintableInterface
    return this
}


func (this *BasicGameDemoPaintable) SetState(state int) {
        if state == 0 {
            this.SetCurrentStatePaintable(this.GetMainDemoStatePaintable())
        } else if state == 1 {
            this.SetCurrentStatePaintable(this.HelpPaintableInterface)
        } else {
            this.SetCurrentStatePaintable(this.OwnershipPaintable)
        }
}


func (this *BasicGameDemoPaintable) Paint(graphics IGraphics) {
        this.GetCurrentStatePaintable().Paint(graphics)
}


func (this *BasicGameDemoPaintable) GetHelpPaintableInterface() IPaintable {
        return this.HelpPaintableInterface
}


func (this *BasicGameDemoPaintable) SetCurrentStatePaintable(currentStatePaintable IPaintable) {
        this.CurrentStatePaintable = currentStatePaintable
}


func (this *BasicGameDemoPaintable) GetCurrentStatePaintable() IPaintable {
        return this.CurrentStatePaintable
}


func (this *BasicGameDemoPaintable) GetMainDemoStatePaintable() IPaintable {
        return this.MainDemoStatePaintable
}

