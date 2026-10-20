/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type SimulatedBuildingPropertiesFactory struct {
    BuildingPropertiesFactory
}

var SimulatedBuildingPropertiesFactory_InstanceC *SimulatedBuildingPropertiesFactory = NewSimulatedBuildingPropertiesFactory()

func NewSimulatedBuildingPropertiesFactory() *SimulatedBuildingPropertiesFactory {
    this := &SimulatedBuildingPropertiesFactory{}
    this.BuildingPropertiesFactory = *NewBuildingPropertiesFactory()
    return this
}


func SimulatedBuildingPropertiesFactory_GetInstance() *SimulatedBuildingPropertiesFactory {
        return SimulatedBuildingPropertiesFactory_InstanceC
}


func (this *SimulatedBuildingPropertiesFactory) GetEfficiencyPerLevel(buildingLayer *BuildingLayer) int {
        return 0
}


func (this *SimulatedBuildingPropertiesFactory) GetDamageFloaters(buildingLayer *BuildingLayer) IDamageFloaters {
        var damageFloaters IDamageFloaters = DamageFloaters_GetInstance(); _ = damageFloaters
        this.DamageFloatersPaintableInterface = damageFloaters
        return damageFloaters
}


// throws Exception
func (this *SimulatedBuildingPropertiesFactory) GetHealthBar(buildingLayer *BuildingLayer) IPaintable {
        return NullPaintable_GetInstance()
}


func (this *SimulatedBuildingPropertiesFactory) GetHashtable() *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}


func (this *SimulatedBuildingPropertiesFactory) GetTrackingEvent(buildingLayer *BuildingLayer) *TrackingEvent {
        return NewTrackingEvent(TrackingEventHandler_GetInstance())
}

