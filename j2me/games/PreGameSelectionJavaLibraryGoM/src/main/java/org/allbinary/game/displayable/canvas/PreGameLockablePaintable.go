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

type PreGameLockablePaintable struct {
    LockablePaintable
    HalfWidth int
    PaintableForm IScrollCurrentSelectionForm
}


func (this *PreGameLockablePaintable) initFields() {
    
    
}


func NewPreGameLockablePaintable(paintableForm IScrollCurrentSelectionForm, lockedIndex int) *PreGameLockablePaintable {
    this := &PreGameLockablePaintable{}
        this.LockablePaintable = *NewLockablePaintable(lockedIndex)
        this.initFields()
        this.PaintableForm = paintableForm
        this.HalfWidth = GameFeatureImageCacheFactory_GetInstance().GetWithKey(LockedResources_GetInstance().LOCKED_DEMO_GAME_FEATURE_RESOURCE).GetWidth() / 2
    return this
}


func (this *PreGameLockablePaintable) Paint(graphics IGraphics, currentIndex int, x int, y int) {
        if LockedUtil_GetInstance().IsLocked() && this.GetPaintableForm().GetSelectedIndex() >= this.GetLockedIndex() {
            this.GetAnimation().PaintXY(graphics, this.GetPaintableForm().GetDx() - this.HalfWidth, this.GetPaintableForm().GetDy() + 34)
        }
}


func (this *PreGameLockablePaintable) GetPaintableForm() IScrollCurrentSelectionForm {
        return this.PaintableForm
}

