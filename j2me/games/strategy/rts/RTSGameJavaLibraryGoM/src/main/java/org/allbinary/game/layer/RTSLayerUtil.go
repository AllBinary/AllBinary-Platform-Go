/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type RTSLayerUtil struct {
    LogUtil *LogUtil
    GroupCommonFactory *GroupCommonFactory
    MAX_RELOAD_TIME int64
}

var RTSLayerUtil_Instance *RTSLayerUtil = NewRTSLayerUtil()

func (this *RTSLayerUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.GroupCommonFactory = GroupCommonFactory_GetInstance()
    this.MAX_RELOAD_TIME = int64(Integer_MAX_VALUE) / 100000
}

func NewRTSLayerUtil() *RTSLayerUtil {
    this := &RTSLayerUtil{}
    this.initFields()
    return this
}


func RTSLayerUtil_GetInstance() *RTSLayerUtil {
        return RTSLayerUtil_Instance
}


func (this *RTSLayerUtil) CreateWeaponProperties(weaponProperties *WeaponProperties, currentLevel int, lastLevel int) *WeaponProperties {
        var reloadTime int64 = weaponProperties.GetReloadTime(); _ = reloadTime
        if lastLevel > currentLevel {
            reloadTime = reloadTime + 50
        } else if lastLevel < currentLevel {
            reloadTime = reloadTime - 50
        }
        var newWeaponProperties *WeaponProperties = NewWeaponProperties(reloadTime, weaponProperties.GetTargetingTime(), weaponProperties.GetSpeed().GetUnscaled(), weaponProperties.GetDamage() / lastLevel * currentLevel, weaponProperties.GetDissipation()); _ = newWeaponProperties
        return newWeaponProperties
}


func (this *RTSLayerUtil) Downgrade(rtsLayer IRTSLayer) {
        rtsLayer.SetLevel(rtsLayer.GetLevel() - 1)
        var partInterfaceArray []PartInterface = rtsLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var partInterface IBasicWeaponPart; _ = partInterface
        for index := int(size); Java_PreDecrement(&index) >= 0;  {
            partInterface = any(partInterfaceArray[index]).(IBasicWeaponPart)
            partInterface.SetWeaponProperties(this.CreateWeaponProperties(partInterface.GetWeaponProperties(), rtsLayer.GetLevel(), rtsLayer.GetLevel() + 1))
        }
        rtsLayer.Select()
}


func (this *RTSLayerUtil) Upgrade(rtsLayer IRTSLayer) {
        rtsLayer.SetLevel(rtsLayer.GetLevel() + 1)
        var partInterfaceArray []PartInterface = rtsLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var partInterface IBasicWeaponPart; _ = partInterface
        for index := int(size); Java_PreDecrement(&index) >= 0;  {
            partInterface = any(partInterfaceArray[index]).(IBasicWeaponPart)
            partInterface.SetWeaponProperties(this.CreateWeaponProperties(partInterface.GetWeaponProperties(), rtsLayer.GetLevel(), rtsLayer.GetLevel() - 1))
        }
        rtsLayer.Select()
}


func (this *RTSLayerUtil) GetCostExponential(cost int64) int64 {
        return ((cost * cost * cost) / (cost * 1000))
}


func (this *RTSLayerUtil) GetWeaponPropertiesCost(weaponProperties *WeaponProperties) int {
        var cost int64 = (int64(weaponProperties.GetDamage() + weaponProperties.GetRange()) + ((this.MAX_RELOAD_TIME / weaponProperties.GetReloadTime()) >> 1)); _ = cost
        cost = this.GetCostExponential(cost)
        return int(cost)
}


func (this *RTSLayerUtil) GetCost(rtsLayer IRTSLayer) int {
        var total int = 0; _ = total
        var partInterfaceArray []PartInterface = rtsLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var partInterface IBasicWeaponPart; _ = partInterface
        var weaponProperties *WeaponProperties; _ = weaponProperties
        for index := int(size - 1); index >= 0; index-- {
            partInterface = any(partInterfaceArray[index]).(IBasicWeaponPart)
            weaponProperties = partInterface.GetWeaponProperties()
            total += this.GetWeaponPropertiesCost(weaponProperties)
        }
        return total
}


func (this *RTSLayerUtil) GetDowngradeCost(rtsLayer IRTSLayer) int {
        var partInterfaceArray []PartInterface = rtsLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var downgradeCost int = 0; _ = downgradeCost
        var partInterface IBasicWeaponPart; _ = partInterface
        var weaponProperties *WeaponProperties; _ = weaponProperties
        var downgradeWeaponCost int; _ = downgradeWeaponCost
        var currentWeaponCost int; _ = currentWeaponCost
        for index := int(size - 1); index >= 0; index-- {
            partInterface = any(partInterfaceArray[index]).(IBasicWeaponPart)
            weaponProperties = partInterface.GetWeaponProperties()
            downgradeWeaponCost = this.GetWeaponPropertiesCost(this.CreateWeaponProperties(weaponProperties, rtsLayer.GetLevel() - 1, rtsLayer.GetLevel()))
            currentWeaponCost = this.GetWeaponPropertiesCost(weaponProperties)
            downgradeCost += (currentWeaponCost - downgradeWeaponCost)
        }
        downgradeCost = downgradeCost * 9 / 10
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Total Cost: ").Appendint(downgradeCost)), this, "getDowngradeCost")
        return downgradeCost
}


func (this *RTSLayerUtil) GetUpgradeCost(rtsLayer IRTSLayer) int {
        var upgradeCost int = 0; _ = upgradeCost
        var partInterfaceArray []PartInterface = rtsLayer.GetPartInterfaceArray(); _ = partInterfaceArray
        var size int = len(partInterfaceArray); _ = size
        var partInterface IBasicWeaponPart; _ = partInterface
        var weaponProperties *WeaponProperties; _ = weaponProperties
        var upgradedWeaponCost int; _ = upgradedWeaponCost
        var currentWeaponCost int; _ = currentWeaponCost
        for index := int(size - 1); index >= 0; index-- {
            partInterface = any(partInterfaceArray[index]).(IBasicWeaponPart)
            weaponProperties = partInterface.GetWeaponProperties()
            upgradedWeaponCost = this.GetWeaponPropertiesCost(this.CreateWeaponProperties(weaponProperties, rtsLayer.GetLevel() + 1, rtsLayer.GetLevel()))
            currentWeaponCost = this.GetWeaponPropertiesCost(weaponProperties)
            upgradeCost += (upgradedWeaponCost - currentWeaponCost)
        }
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Total Cost: ").Appendint(upgradeCost)), this, "getUpgradeCost")
        return upgradeCost
}


// throws Exception
func (this *RTSLayerUtil) DestroyAndClear(list IBasicArrayList) {
        var rtsLayer IRTSLayer; _ = rtsLayer
        for index := int(list.Size() - 1); index >= 0; index-- {
            rtsLayer = any(list.Get(index)).(IRTSLayer)
            rtsLayer.SetDestroyed(true)
        }
        list.Clear()
}


func (this *RTSLayerUtil) CreateFakeRTSLayerHashtable(baseRaceTrackGeographicMap IBaseRaceTrackGeographicMap) *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = hashtable
        var layerManager *FakeLayerManager = NewFakeLayerManager(GameInfo_NONE); _ = layerManager
        layerManager.SetGeographicMapInterface([]IBasicGeographicMap{baseRaceTrackGeographicMap})
        hashtable.Put(AllBinaryGameLayerManager_ID, layerManager)
        hashtable.Put(DirectionFactory_GetInstance().NAME, DirectionFactory_GetInstance().LEFT)
        hashtable.Put(this.GroupCommonFactory.ID, BasicGroupFactory_GetInstance().NONE_ARRAY)
        return hashtable
}

