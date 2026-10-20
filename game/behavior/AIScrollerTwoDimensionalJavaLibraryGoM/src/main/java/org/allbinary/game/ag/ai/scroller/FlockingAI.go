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

type FlockingAI struct {
    BasicAI
    AllowedDistance int
}


func (this *FlockingAI) initFields() {
    
}


func NewFlockingAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *FlockingAI {
    this := &FlockingAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        var allowedDistance int = any(hashtable.Get(BasePatrolAI_MAX_DISTANCE)).(int); _ = allowedDistance
        this.AllowedDistance = allowedDistance.IntValue()
    return this
}


// throws Exception
func (this *FlockingAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        var maxDistance int = 0; _ = maxDistance
        var farAllbinaryLayer IAllBinaryLayer = AllBinaryLayer_NULL_ALLBINARY_LAYER; _ = farAllbinaryLayer
        var groupLayerManagerListener *GroupLayerManagerListener = GroupLayerManagerListener_GetInstance(); _ = groupLayerManagerListener
        var list IBasicArrayList = groupLayerManagerListener.GetList(BasicGroupFactory_GetInstance().ENEMY); _ = list
        var size int = list.Size(); _ = size
        var allBinaryLayer IAllBinaryLayer; _ = allBinaryLayer
        var ownerLayerInterface IAllBinaryLayer; _ = ownerLayerInterface
        for index := int(0); index < size; index++ {
            allBinaryLayer = any(list.Get(index)).(IAllBinaryLayer)
            if allBinaryLayer.GetType() != WeaponLayer_GetStaticType() {
                var distance int = this.GetXYDistance(allBinaryLayer); _ = distance
                if distance > maxDistance {
                    maxDistance = distance
                    farAllbinaryLayer = allBinaryLayer
                    if distance > this.AllowedDistance {
                        break
                    }
                }
            }
        }
        if maxDistance > this.AllowedDistance {
            if farAllbinaryLayer != AllBinaryLayer_NULL_ALLBINARY_LAYER {
                ownerLayerInterface = this.GetOwnerLayerInterface()
                if farAllbinaryLayer.GetXP() < ownerLayerInterface.GetXP() {
                    this.ProcessKeyAI(Canvas_LEFT)
                } else if farAllbinaryLayer.GetXP() > ownerLayerInterface.GetXP() {
                    this.ProcessKeyAI(Canvas_RIGHT)
                } else if farAllbinaryLayer.GetYP() > ownerLayerInterface.GetYP() {
                    this.ProcessKeyAI(Canvas_UP)
                } else if farAllbinaryLayer.GetYP() < ownerLayerInterface.GetYP() {
                    this.ProcessKeyAI(Canvas_DOWN)
                }
            }
        }
}


func (this *FlockingAI) GetXYDistance(allBinaryLayer IAllBinaryLayer) int {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var xTotalDistance int = (allBinaryLayer.GetXP() + allBinaryLayer.GetHalfWidth()) - (ownerLayerInterface.GetXP() + ownerLayerInterface.GetHalfWidth()); _ = xTotalDistance
        var yTotalDistance int = (allBinaryLayer.GetYP() + allBinaryLayer.GetHalfHeight()) - (ownerLayerInterface.GetYP() + ownerLayerInterface.GetHalfHeight()); _ = yTotalDistance
        var mathUtil *MathUtil = MathUtil_GetInstance(); _ = mathUtil
        var totalDistance int = mathUtil.Abs(xTotalDistance) + mathUtil.Abs(yTotalDistance); _ = totalDistance
        return totalDistance
}

