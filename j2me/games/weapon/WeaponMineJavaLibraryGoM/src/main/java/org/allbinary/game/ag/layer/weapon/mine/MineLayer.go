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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type MineLayer struct {
    SimpleWeaponLayer
    ViewPositionEventHandler *ViewPositionEventHandler
}

var MineLayer_NAME string = "MineLayer"

func (this *MineLayer) initFields() {
    this.ViewPositionEventHandler = ViewPositionEventHandler_GetInstance()
}


func NewMineLayer(viewPosition IViewPositionBase) *MineLayer {
    this := &MineLayer{}
        this.SimpleWeaponLayer = *NewSimpleWeaponLayer(MineLayer_NAME, RemoteInfo_REMOTE_INFO, -1, NoMovementFactory_GetInstance().GetMovmentInstance(), FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().Get(MineWeaponResources_GetInstance().RESOURCE).GetInstance(0), SimpleWeaponLayer_CreateDestroyed(), NewRectangle(PointFactory_GetInstance().ZERO_ZERO, 10, 10), viewPosition)
        this.initFields()
    return this
}


func (this *MineLayer) InitXYZ(x int, y int, z int) {
        this.SetPosition(x + 5, y + 5, z)
        this.ViewPositionEventHandler.AddListener(any(this).(IAllBinaryLayer))
        SecondaryPlayerQueueFactory_GetInstance().Add(DropWeaponSound_GetInstance())
}


func (this *MineLayer) Damage(damage int, damageType int) {
        this.TotalDamage = this.GetInitDamage() + 1
}


func (this *MineLayer) GetDamage(damageType int) int {
        this.SimpleWeaponLayer.GetDamage(damageType)
        return this.GetInitDamage()
}


func (this *MineLayer) SetDestroyed(destroyed bool) {
        this.SimpleWeaponLayer.SetDestroyed(destroyed)
        if this.IsDestroyed() {
            this.ViewPositionEventHandler.RemoveListener(this)
        }
}

