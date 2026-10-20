/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type BaseMenuBehavior struct {
}

// The Java class BaseMenuBehavior has subclasses so its variables are this interface
type IBaseMenuBehavior interface {
    BaseMenuBehavior_() *BaseMenuBehavior
    OnDisplayChangeEvent(allBinaryGameCanvas IAllBinaryGameCanvas, displayChangeEvent *DisplayChangeEvent)
    InitSpecialPaint(allBinaryGameCanvas IAllBinaryGameCanvas)
    InitMenu(allBinaryGameCanvas IAllBinaryGameCanvas)
    UpdateMenu(allBinaryGameCanvas IAllBinaryGameCanvas)
    PopupMenu(allBinaryGameCanvas IAllBinaryGameCanvas)
    CloseMenu(allBinaryGameCanvas IAllBinaryGameCanvas)
}

func (this *BaseMenuBehavior) BaseMenuBehavior_() *BaseMenuBehavior {
    return this
}

var BaseMenuBehavior_Instance IBaseMenuBehavior = NewBaseMenuBehavior()

func NewBaseMenuBehavior() *BaseMenuBehavior {
    this := &BaseMenuBehavior{}
    return this
}


func BaseMenuBehavior_GetInstance() IBaseMenuBehavior {
        return BaseMenuBehavior_Instance
}


// throws Exception
func (this *BaseMenuBehavior) OnDisplayChangeEvent(allBinaryGameCanvas IAllBinaryGameCanvas, displayChangeEvent *DisplayChangeEvent) {
}


func (this *BaseMenuBehavior) InitSpecialPaint(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.SetNonBotPaintableP(NullPaintable_GetInstance())
}


// throws Exception
func (this *BaseMenuBehavior) InitMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *BaseMenuBehavior) UpdateMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


// throws Exception
func (this *BaseMenuBehavior) PopupMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
}


func (this *BaseMenuBehavior) CloseMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
}

