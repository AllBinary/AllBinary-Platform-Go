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

type RepairTechnology struct {
    Level int
    MaxLevel int
}

// Implements RTSInterface


func (this *RepairTechnology) initFields() {
    
    
}


func NewRepairTechnology(maxLevel int) *RepairTechnology {
    this := &RepairTechnology{}
    this.initFields()
        this.MaxLevel = maxLevel
    return this
}


func (this *RepairTechnology) IsCompleted() bool {
        return false
}


func (this *RepairTechnology) GetPercentComplete() int {
        return 0
}


func (this *RepairTechnology) GetLevel() int {
        return this.Level
}


func (this *RepairTechnology) GetMaxLevel() int {
        return this.MaxLevel
}


func (this *RepairTechnology) GetCost() int {
        return 0
}


func (this *RepairTechnology) Downgrade() {
}


func (this *RepairTechnology) GetDowngradeCost() int {
        return 0
}


func (this *RepairTechnology) IsDowngradeable() bool {
        return false
}


func (this *RepairTechnology) GetUpgradeCost() int {
        return (this.GetLevel() * this.GetLevel()) * 1000
}


func (this *RepairTechnology) IsUpgradeable() bool {
        if this.GetLevel() < this.GetMaxLevel() {
            return true
        } else {
            return false
        }
}


func (this *RepairTechnology) Upgrade() {
        this.Level++
        var unitTechnologySingletonFactory *UnitTechnologySingletonFactory = UnitTechnologySingletonFactory_GetInstance(); _ = unitTechnologySingletonFactory
        unitTechnologySingletonFactory.SetRepair(unitTechnologySingletonFactory.GetRepair() + (this.GetLevel() * 10))
}


func (this *RepairTechnology) IsSelfUpgradeable() bool {
        return false
}

