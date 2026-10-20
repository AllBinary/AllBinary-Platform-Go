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

type CollidableUnitBehavior struct {
    CollidableRTSBehavior
    LayerPartialCellPositionsUtil *LayerPartialCellPositionsUtil
    SimpleSteeringVisitor *SimpleSteeringVisitor
}


func (this *CollidableUnitBehavior) initFields() {
    this.LayerPartialCellPositionsUtil = LayerPartialCellPositionsUtil_GetInstance()
    this.SimpleSteeringVisitor = NewSimpleSteeringVisitor()
}


func NewCollidableUnitBehavior(collidable bool) *CollidableUnitBehavior {
    this := &CollidableUnitBehavior{}
        this.CollidableRTSBehavior = *NewCollidableRTSBehavior(collidable)
        this.initFields()
    return this
}


// throws Exception
func (this *CollidableUnitBehavior) CollideNone(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        this.Chase(ownerLayer, collidableInterfaceCompositeInterface)
        var rtsLayer IAdvancedRTSGameLayer = any(collidableInterfaceCompositeInterface).(IAdvancedRTSGameLayer); _ = rtsLayer
        if rtsLayer.GetType() == UnitLayer_GetStaticType() {
            this.CollideUnit(ownerLayer, any(rtsLayer).(*UnitLayer))
        }
}


// throws Exception
func (this *CollidableUnitBehavior) Steer(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) bool {
        var ownerUnitLayer *UnitLayer = any(ownerLayer).(*UnitLayer); _ = ownerUnitLayer
        var unitLayer *UnitLayer = any(collidableInterfaceCompositeInterface).(*UnitLayer); _ = unitLayer
        var angleInfo2 *AngleInfo = unitLayer.GetRotationAnimationInterface().GetAngleInfoP(); _ = angleInfo2
        var angleInfo *AngleInfo = ownerUnitLayer.GetRotationAnimationInterface().GetAngleInfoP(); _ = angleInfo
        var angle int = int(angleInfo.GetAngle()) - int(angleInfo2.GetAngle()); _ = angle
        if angle < 90 || angle > 270 {
            return true
        } else if angle >= 90 || angle <= 270 {
            ownerUnitLayer.GetGameKeyEventList().Add(GameKeyEventFactory_GetInstance().GetInstanceForKey(ownerUnitLayer, Canvas_RIGHT))
            return false
        }
        panic(errors.New(fmt.Sprint("Error")))
}


// throws Exception
func (this *CollidableUnitBehavior) Chase(ownerLayer ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        var rtsLayer IAdvancedRTSGameLayer = any(collidableInterfaceCompositeInterface).(IAdvancedRTSGameLayer); _ = rtsLayer
        if rtsLayer.GetType() == UnitLayer_GetStaticType() {
            if !this.SimpleSteeringVisitor.GetList().Contains(collidableInterfaceCompositeInterface) {
                this.SimpleSteeringVisitor.GetList().Add(collidableInterfaceCompositeInterface)
            }
            var ownerUnitLayer *UnitLayer = any(ownerLayer).(*UnitLayer); _ = ownerUnitLayer
            var list IBasicArrayList = ownerUnitLayer.GetUnitWaypointBehavior().GetSteeringVisitorList(); _ = list
            if !list.Contains(this.SimpleSteeringVisitor) {
                list.Add(this.SimpleSteeringVisitor)
            }
        }
}


// throws Exception
func (this *CollidableUnitBehavior) CollideUnit(ownerLayer ICollidableCompositeLayer, unitLayer *UnitLayer) {
        var ownerUnitLayer *UnitLayer = any(ownerLayer).(*UnitLayer); _ = ownerUnitLayer
        var partialPositionList IBasicArrayList = UnitLayer_GetPartialpositionlist(); _ = partialPositionList
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(ownerUnitLayer.AllBinaryGameLayerManagerP).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        var basicGeographicMap IBasicGeographicMap = geographicMapInterface; _ = basicGeographicMap
        var dropCellPositionHistory *DropCellPositionHistory = DropCellPositionHistory_GetInstance(); _ = dropCellPositionHistory
        var tiledLayerUtil *TiledLayerUtil = TiledLayerUtil_GetInstance(); _ = tiledLayerUtil
        var tiledLayer IAllBinaryTiledLayer = basicGeographicMap.GetAllBinaryTiledLayer(); _ = tiledLayer
        if ownerLayer.GetXP() < unitLayer.GetXP() && ownerLayer.GetX2() > unitLayer.GetXP() {
            var diff int = ownerLayer.GetWidth() + 1; _ = diff
            this.LayerPartialCellPositionsUtil.GetAllDXY(basicGeographicMap, ownerLayer, -diff, 0, partialPositionList)
            if !dropCellPositionHistory.AnyCellPositionWithDrop(partialPositionList) {
                var x int = unitLayer.GetXP() - diff; _ = x
                var y int = ownerLayer.GetYP(); _ = y
                x = tiledLayerUtil.KeepOnMapX(tiledLayer, x, ownerLayer.GetWidth())
                y = tiledLayerUtil.KeepOnMapY(tiledLayer, y, ownerLayer.GetHeight())
                ownerLayer.SetPosition(x, y, ownerLayer.GetZP())
            }
            this.LayerPartialCellPositionsUtil.GetAllDXY(basicGeographicMap, unitLayer, diff, 0, partialPositionList)
            if !dropCellPositionHistory.AnyCellPositionWithDrop(partialPositionList) {
                var x int = unitLayer.GetXP() + diff; _ = x
                var y int = unitLayer.GetYP(); _ = y
                x = tiledLayerUtil.KeepOnMapX(tiledLayer, x, unitLayer.GetWidth())
                y = tiledLayerUtil.KeepOnMapY(tiledLayer, y, unitLayer.GetHeight())
                unitLayer.SetPosition(x, y, unitLayer.GetZP())
            }
        }
        if ownerLayer.GetYP() < unitLayer.GetYP() && ownerLayer.GetY2() > unitLayer.GetYP() {
            var diff int = ownerLayer.GetHeight() + 1; _ = diff
            this.LayerPartialCellPositionsUtil.GetAllDXY(basicGeographicMap, ownerLayer, 0, -diff, partialPositionList)
            if !dropCellPositionHistory.AnyCellPositionWithDrop(partialPositionList) {
                var x int = ownerLayer.GetXP(); _ = x
                var y int = unitLayer.GetYP() - diff; _ = y
                x = tiledLayerUtil.KeepOnMapX(tiledLayer, x, ownerLayer.GetWidth())
                y = tiledLayerUtil.KeepOnMapY(tiledLayer, y, ownerLayer.GetHeight())
                ownerLayer.SetPosition(x, y, ownerLayer.GetZP())
            }
            this.LayerPartialCellPositionsUtil.GetAllDXY(basicGeographicMap, unitLayer, 0, diff, partialPositionList)
            if !dropCellPositionHistory.AnyCellPositionWithDrop(partialPositionList) {
                var x int = unitLayer.GetXP(); _ = x
                var y int = unitLayer.GetYP() + diff; _ = y
                x = tiledLayerUtil.KeepOnMapX(tiledLayer, x, unitLayer.GetWidth())
                y = tiledLayerUtil.KeepOnMapY(tiledLayer, y, unitLayer.GetHeight())
                unitLayer.SetPosition(x, y, unitLayer.GetZP())
            }
        }
}

type SimpleSteeringVisitor struct {
    SteeringVisitor
}


func NewSimpleSteeringVisitor() *SimpleSteeringVisitor {
    this := &SimpleSteeringVisitor{}
    this.SteeringVisitor = *NewSteeringVisitor()
    return this
}
    

    func (this *CollidableUnitBehavior) Visit(anyType any) any {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                logUtil.Put(commonStrings.EXCEPTION, this, "visit", e)
                tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_OBJECT; return
                    }
                }()
                var ownerLayer ICollidableCompositeLayer = any(anyType).(ICollidableCompositeLayer); _ = ownerLayer
                if this.GetList().Size() > 0 {
                    var allbinaryLayer ICollidableCompositeLayer = any(this.GetList().Get(0)).(ICollidableCompositeLayer); _ = allbinaryLayer
                    var clear_ bool = this.Steer(ownerLayer, allbinaryLayer); _ = clear_
                    if clear_ {
                        this.GetList().Clear()
                        tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_OBJECT; return
                    }
                    tryControl1, tryResult1 = 1, BooleanFactory_GetInstance().TRUE; return
                }
                tryControl1, tryResult1 = 1, NullUtil_GetInstance().NULL_OBJECT; return
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
            panic("unreachable")
    }

