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

type InGameMenuBehavior struct {
    BaseMenuBehavior
}

var InGameMenuBehavior_InstanceC *InGameMenuBehavior = NewInGameMenuBehavior()

func NewInGameMenuBehavior() *InGameMenuBehavior {
    this := &InGameMenuBehavior{}
    this.BaseMenuBehavior = *NewBaseMenuBehavior()
    return this
}


func InGameMenuBehavior_GetInstance() *InGameMenuBehavior {
        return InGameMenuBehavior_InstanceC
}


// throws Exception
func (this *InGameMenuBehavior) OnDisplayChangeEvent(allBinaryGameCanvas IAllBinaryGameCanvas, displayChangeEvent *DisplayChangeEvent) {
        allBinaryGameCanvas.UpdateMenuFromEvent(displayChangeEvent)
}


func (this *InGameMenuBehavior) InitSpecialPaint(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.SetNonBotPaintableP(NewGameCanvasNonBotPaintable(allBinaryGameCanvas))
}


// throws Exception
func (this *InGameMenuBehavior) InitMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.InitMenu2()
}


// throws Exception
func (this *InGameMenuBehavior) UpdateMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.InitMenu2()
}


// throws Exception
func (this *InGameMenuBehavior) PopupMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.PopupMenu2()
}


func (this *InGameMenuBehavior) CloseMenu(allBinaryGameCanvas IAllBinaryGameCanvas) {
        allBinaryGameCanvas.CloseMenu2()
}

