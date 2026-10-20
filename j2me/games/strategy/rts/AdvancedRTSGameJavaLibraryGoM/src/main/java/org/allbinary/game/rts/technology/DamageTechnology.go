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

type DamageTechnology struct {
    Level int
    MaxLevel int
}

// Implements RTSInterface


func (this *DamageTechnology) initFields() {
    
    
}


func NewDamageTechnology(maxLevel int) *DamageTechnology {
    this := &DamageTechnology{}
    this.initFields()
        this.MaxLevel = maxLevel
    return this
}


func (this *DamageTechnology) IsCompleted() bool {
        return false
}


func (this *DamageTechnology) GetPercentComplete() int {
        return 0
}


func (this *DamageTechnology) GetLevel() int {
        return this.Level
}


func (this *DamageTechnology) GetMaxLevel() int {
        return this.MaxLevel
}


func (this *DamageTechnology) GetCost() int {
        return 0
}


func (this *DamageTechnology) Downgrade() {
}


func (this *DamageTechnology) GetDowngradeCost() int {
        return 0
}


func (this *DamageTechnology) IsDowngradeable() bool {
        return false
}


func (this *DamageTechnology) GetUpgradeCost() int {
        return (this.GetLevel() * this.GetLevel()) * 1000
}


func (this *DamageTechnology) IsUpgradeable() bool {
        if this.GetLevel() < this.GetMaxLevel() {
            return true
        } else {
            return false
        }
}


func (this *DamageTechnology) Upgrade() {
        this.Level++
        var unitTechnologySingletonFactory *UnitTechnologySingletonFactory = UnitTechnologySingletonFactory_GetInstance(); _ = unitTechnologySingletonFactory
        unitTechnologySingletonFactory.SetDamage(unitTechnologySingletonFactory.GetDamage() + (this.GetLevel() * 100))
}


func (this *DamageTechnology) IsSelfUpgradeable() bool {
        return false
}

