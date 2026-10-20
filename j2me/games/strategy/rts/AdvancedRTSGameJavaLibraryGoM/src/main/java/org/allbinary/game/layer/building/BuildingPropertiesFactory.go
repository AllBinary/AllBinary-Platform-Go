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

type BuildingPropertiesFactory struct {
    DamageFloatersPaintableInterface IPaintable
}

// The Java class BuildingPropertiesFactory has subclasses so its variables are this interface
type IBuildingPropertiesFactory interface {
    BuildingPropertiesFactory_() *BuildingPropertiesFactory
    GetEfficiencyPerLevel(buildingLayer *BuildingLayer) int
    GetDamageFloaters(buildingLayer *BuildingLayer) IDamageFloaters
    GetHealthBar(buildingLayer *BuildingLayer) IPaintable
    GetHashtable() *ABHashtable[any, any]
    GetTrackingEvent(buildingLayer *BuildingLayer) *TrackingEvent
}

func (this *BuildingPropertiesFactory) BuildingPropertiesFactory_() *BuildingPropertiesFactory {
    return this
}

var BuildingPropertiesFactory_Instance IBuildingPropertiesFactory = NewBuildingPropertiesFactory()

func (this *BuildingPropertiesFactory) initFields() {
    this.DamageFloatersPaintableInterface = NullPaintable_GetInstance()
}

func NewBuildingPropertiesFactory() *BuildingPropertiesFactory {
    this := &BuildingPropertiesFactory{}
    this.initFields()
    return this
}


func BuildingPropertiesFactory_GetInstance() IBuildingPropertiesFactory {
        return BuildingPropertiesFactory_Instance
}


func (this *BuildingPropertiesFactory) GetEfficiencyPerLevel(buildingLayer *BuildingLayer) int {
        return 10000 / buildingLayer.GetMaxLevel() + 10000 % buildingLayer.GetMaxLevel()
}


func (this *BuildingPropertiesFactory) GetDamageFloaters(buildingLayer *BuildingLayer) IDamageFloaters {
        var damageFloaters IDamageFloaters = DamageFloaters_GetInstance(); _ = damageFloaters
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().DAMAGE_FLOATERS) {
            damageFloaters = NewPtsDamageFloaters(buildingLayer)
            this.DamageFloatersPaintableInterface = damageFloaters
        } else {
            damageFloaters = NewDamageFloaters()
        }
        return damageFloaters
}


// throws Exception
func (this *BuildingPropertiesFactory) GetHealthBar(buildingLayer *BuildingLayer) IPaintable {
        var healthBar IPaintable = NullPaintable_GetInstance(); _ = healthBar
        if Features_GetInstance().IsFeature(GameFeatureFactory_GetInstance().HEALTH_BARS) {
            healthBar = NewHealthBar(buildingLayer, buildingLayer.GetHealthInterface(), NewHealthBarTwodAnimation(any(buildingLayer).(IAllBinaryLayer), BasicHudFactory_GetInstance().BOTTOMLEFT), -1)
        }
        return healthBar
}


func (this *BuildingPropertiesFactory) GetHashtable() *ABHashtable[any, any] {
        return StdUtil_GetInstance().CreateHashtable()
}


func (this *BuildingPropertiesFactory) GetTrackingEvent(buildingLayer *BuildingLayer) *TrackingEvent {
        return NewTrackingEvent(buildingLayer)
}

