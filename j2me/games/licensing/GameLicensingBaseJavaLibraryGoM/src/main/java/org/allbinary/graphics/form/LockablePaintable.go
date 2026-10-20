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

type LockablePaintable struct {
    ItemPaintable
    Animation IAnimation
    LockedIndex int
}

// The Java class LockablePaintable has subclasses so its variables are this interface
type ILockablePaintable interface {
    IItemPaintable
    LockablePaintable_() *LockablePaintable
    Paint(graphics IGraphics, currentIndex int, x int, y int)
    GetLockedIndex() int
    GetAnimation() IAnimation
}

func (this *LockablePaintable) LockablePaintable_() *LockablePaintable {
    return this
}


func (this *LockablePaintable) initFields() {
    this.Animation = FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(LockedResources_GetInstance().LOCKED_DEMO_GAME_FEATURE_RESOURCE).GetInstance(0)
    
}


func NewLockablePaintable(lockedIndex int) *LockablePaintable {
    this := &LockablePaintable{}
    this.ItemPaintable = *NewItemPaintable()
    this.initFields()
        this.LockedIndex = lockedIndex
    return this
}


func (this *LockablePaintable) Paint(graphics IGraphics, currentIndex int, x int, y int) {
        if LockedUtil_GetInstance().IsLocked() && currentIndex >= this.GetLockedIndex() {
            this.GetAnimation().PaintXY(graphics, x, y)
        }
}


func (this *LockablePaintable) GetLockedIndex() int {
        return this.LockedIndex
}


func (this *LockablePaintable) GetAnimation() IAnimation {
        return this.Animation
}

