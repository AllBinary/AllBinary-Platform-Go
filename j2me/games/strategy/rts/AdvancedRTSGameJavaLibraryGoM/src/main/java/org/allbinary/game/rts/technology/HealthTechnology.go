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

type HealthTechnology struct {
    Level int
    MaxLevel int
}

// Implements RTSInterface


func (this *HealthTechnology) initFields() {
    
    
}


func NewHealthTechnology(maxLevel int) *HealthTechnology {
    this := &HealthTechnology{}
    this.initFields()
        this.MaxLevel = maxLevel
    return this
}


func (this *HealthTechnology) IsCompleted() bool {
        return false
}


func (this *HealthTechnology) GetPercentComplete() int {
        return 0
}


func (this *HealthTechnology) GetLevel() int {
        return this.Level
}


func (this *HealthTechnology) GetMaxLevel() int {
        return this.MaxLevel
}


func (this *HealthTechnology) GetCost() int {
        return 0
}


func (this *HealthTechnology) Downgrade() {
}


func (this *HealthTechnology) GetDowngradeCost() int {
        return 0
}


func (this *HealthTechnology) IsDowngradeable() bool {
        return false
}


func (this *HealthTechnology) GetUpgradeCost() int {
        return (this.GetLevel() * this.GetLevel()) * 1000
}


func (this *HealthTechnology) IsUpgradeable() bool {
        if this.GetLevel() < this.GetMaxLevel() {
            return true
        } else {
            return false
        }
}


func (this *HealthTechnology) Upgrade() {
        this.Level++
        var unitTechnologySingletonFactory *UnitTechnologySingletonFactory = UnitTechnologySingletonFactory_GetInstance(); _ = unitTechnologySingletonFactory
        unitTechnologySingletonFactory.SetHealth(unitTechnologySingletonFactory.GetHealth() + (this.GetLevel() * 100))
}


func (this *HealthTechnology) IsSelfUpgradeable() bool {
        return false
}

