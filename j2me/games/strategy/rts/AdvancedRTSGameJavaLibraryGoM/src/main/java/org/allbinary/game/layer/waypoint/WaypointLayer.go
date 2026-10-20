/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type WaypointLayer struct {
    AdvancedRTSGameLayer
}

// The Java class WaypointLayer has subclasses so its variables are this interface
type IWaypointLayer interface {
    IAdvancedRTSGameLayer
    WaypointLayer_() *WaypointLayer
    Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface)
    Paint(graphics IGraphics)
    ProcessTick(allBinaryLayerManager IAllBinaryLayerManager)
    Damage(damage int, damageType int)
    GetDamage(damageType int) int
    SetDestroyed(destroyed bool)
    GetType() int
    IsSelfUpgradeable() bool
}

func (this *WaypointLayer) WaypointLayer_() *WaypointLayer {
    return this
}



func NewWaypointLayer(remoteInfo *RemoteInfo, parentLayer PathFindingLayerInterface, groupInterface []*Group, rootName string, name string, animationInterfaceFactoryInterface AnimationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface ProceduralAnimationInterfaceFactoryInterface, rectangle *Rectangle, x int, y int) *WaypointLayer {
    this := &WaypointLayer{}
        this.AdvancedRTSGameLayer = *NewAdvancedRTSGameLayer(remoteInfo, parentLayer, AdvancedRTSProperties_CreateProperties(), groupInterface, rootName, name, NewHealth(1000), NullRTSFormInputFactory_GetInstance(), animationInterfaceFactoryInterface, animationInterfaceFactoryInterface, animationInterfaceFactoryInterface, animationInterfaceFactoryInterface, animationInterfaceFactoryInterface, proceduralAnimationInterfaceFactoryInterface, rectangle, x, y, NewTileLayerPositionIntoViewPosition())
        this.SetCollidableInferface(NewCollidableWaypointBehavior(true))
        this.GetWaypointBehavior().SetWaypoint(NewWaypoint(this, AttackSound_GetInstance()))
        this.SetAnimationInterface(this.IndexedButShouldBeRotationAnimationInterface)
    return this
}


func (this *WaypointLayer) Construct(rtsPlayerLayerInterface RTSPlayerLayerInterface) {
        this.PercentCompleteP = 100
        this.InitVisibility(rtsPlayerLayerInterface)
}


func (this *WaypointLayer) Paint(graphics IGraphics) {
        if this.IsVisible() {
            this.AdvancedRTSGameLayer.Paint(graphics)
        }
}


func (this *WaypointLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "processTick", e)
                }
            }()
            this.IndexedButShouldBeRotationAnimationInterface.NextFrame()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *WaypointLayer) Damage(damage int, damageType int) {
}


// throws Exception
func (this *WaypointLayer) GetDamage(damageType int) int {
        return 0
}


// throws Exception
func (this *WaypointLayer) SetDestroyed(destroyed bool) {
        this.AdvancedRTSGameLayer.SetDestroyed(destroyed)
        if this.IsDestroyed() {
            WaypointCellPositionHistory_GetInstance().RemoveLayer(this)
        }
}


func (this *WaypointLayer) GetType() int {
        return WaypointLayer_GetStaticType()
}


func WaypointLayer_GetStaticType() int {
        return 4
}


func (this *WaypointLayer) IsSelfUpgradeable() bool {
        return false
}

