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

type AdvancedPlayerOwnedRTSLayers struct {
    PlayerOwnedRTSLayers
    UnitsList IBasicArrayList
    WaypointsList IBasicArrayList
    BuildingList IBasicArrayList
    BaseList IBasicArrayList
    GarrisonList IBasicArrayList
    HouseList IBasicArrayList
    LabList IBasicArrayList
    FactoryList IBasicArrayList
}


func (this *AdvancedPlayerOwnedRTSLayers) initFields() {
    this.UnitsList = NewBasicArrayListD()
    this.WaypointsList = NewBasicArrayListD()
    this.BuildingList = NewBasicArrayListD()
    this.BaseList = NewBasicArrayListD()
    this.GarrisonList = NewBasicArrayListD()
    this.HouseList = NewBasicArrayListD()
    this.LabList = NewBasicArrayListD()
    this.FactoryList = NewBasicArrayListD()
}

func NewAdvancedPlayerOwnedRTSLayers() *AdvancedPlayerOwnedRTSLayers {
    this := &AdvancedPlayerOwnedRTSLayers{}
    this.PlayerOwnedRTSLayers = *NewPlayerOwnedRTSLayers()
    this.initFields()
    return this
}


func (this *AdvancedPlayerOwnedRTSLayers) AddUnit(rtsLayer IRTSLayer) {
        this.UnitsList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) AddWaypoint(rtsLayer IRTSLayer) {
        this.WaypointsList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) AddBase(rtsLayer IRTSLayer) {
        this.BaseList.Add(rtsLayer)
        this.BuildingList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) AddGarrison(rtsLayer IRTSLayer) {
        this.GarrisonList.Add(rtsLayer)
        this.BuildingList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) AddHouse(rtsLayer IRTSLayer) {
        this.HouseList.Add(rtsLayer)
        this.BuildingList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) AddLab(rtsLayer IRTSLayer) {
        this.LabList.Add(rtsLayer)
        this.BuildingList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) AddFactory(rtsLayer IRTSLayer) {
        this.FactoryList.Add(rtsLayer)
        this.BuildingList.Add(rtsLayer)
        this.RtsLayerList.Add(rtsLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) RemoveUnit(gameLayer IAllBinaryGameLayer) {
        this.UnitsList.Remove(gameLayer)
        this.RtsLayerList.Remove(gameLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) RemoveWaypoint(gameLayer IAllBinaryGameLayer) {
        this.WaypointsList.Remove(gameLayer)
        this.RtsLayerList.Remove(gameLayer)
}


func (this *AdvancedPlayerOwnedRTSLayers) GetUnitsList() IBasicArrayList {
        return this.UnitsList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetWaypointsList() IBasicArrayList {
        return this.WaypointsList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetBaseList() IBasicArrayList {
        return this.BaseList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetGarrisonList() IBasicArrayList {
        return this.GarrisonList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetHouseList() IBasicArrayList {
        return this.HouseList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetLabList() IBasicArrayList {
        return this.LabList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetFactoryList() IBasicArrayList {
        return this.FactoryList
}


func (this *AdvancedPlayerOwnedRTSLayers) GetTotalBuildings() int {
        return this.BaseList.Size() + this.GarrisonList.Size() + this.HouseList.Size() + this.LabList.Size() + this.FactoryList.Size()
}


func (this *AdvancedPlayerOwnedRTSLayers) RemoveBuilding(layerInterface IAllBinaryGameLayer) {
        this.BuildingList.Add(layerInterface)
        if this.GetBaseList().Remove(layerInterface) {
            this.RtsLayerList.Remove(layerInterface)
        } else if this.GetFactoryList().Remove(layerInterface) {
            this.RtsLayerList.Remove(layerInterface)
        } else if this.GetGarrisonList().Remove(layerInterface) {
            this.RtsLayerList.Remove(layerInterface)
        } else if this.GetLabList().Remove(layerInterface) {
            this.RtsLayerList.Remove(layerInterface)
        } else if this.GetHouseList().Remove(layerInterface) {
            this.RtsLayerList.Remove(layerInterface)
        }
}


func (this *AdvancedPlayerOwnedRTSLayers) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.PlayerOwnedRTSLayers.ToString())
        stringBuffer.Append(" Bases: ")
        stringBuffer.Appendint(this.GetBaseList().Size())
        stringBuffer.Append(" Factory: ")
        stringBuffer.Appendint(this.GetFactoryList().Size())
        stringBuffer.Append(" Garrison: ")
        stringBuffer.Appendint(this.GetGarrisonList().Size())
        stringBuffer.Append(" Labs: ")
        stringBuffer.Appendint(this.GetLabList().Size())
        stringBuffer.Append(" Houses: ")
        stringBuffer.Appendint(this.GetHouseList().Size())
        stringBuffer.Append(" Waypoints: ")
        stringBuffer.Appendint(this.GetWaypointsList().Size())
        stringBuffer.Append(" Units: ")
        stringBuffer.Appendint(this.GetUnitsList().Size())
        return fmt.Sprint(stringBuffer)
}

func (this *AdvancedPlayerOwnedRTSLayers) String() string {
    return this.ToString()
}


func (this *AdvancedPlayerOwnedRTSLayers) GetBuildingList() IBasicArrayList {
        return this.BuildingList
}

