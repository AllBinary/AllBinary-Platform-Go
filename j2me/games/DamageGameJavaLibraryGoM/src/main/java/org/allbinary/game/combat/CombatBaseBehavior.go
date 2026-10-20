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

type CombatBaseBehavior struct {
    DamageableBaseBehavior *DamageableBaseBehavior
    DestroyableBaseBehavior IDestroyableBaseBehavior
    ReadyForExplosion bool
}

var CombatBaseBehavior_NULL_COMBAT_BEHAVIOR *CombatBaseBehavior = NewCombatBaseBehavior(DamageableBaseBehavior_GetInstance(), DestroyableBaseBehavior_GetInstance())

func (this *CombatBaseBehavior) initFields() {
    
    
    
}


func NewCombatBaseBehavior(damageableBaseBehavior *DamageableBaseBehavior, destroyableBaseBehavior IDestroyableBaseBehavior) *CombatBaseBehavior {
    this := &CombatBaseBehavior{}
    this.initFields()
        this.DamageableBaseBehavior = damageableBaseBehavior
        this.DestroyableBaseBehavior = destroyableBaseBehavior
    return this
}


func (this *CombatBaseBehavior) GetDamageableBaseBehavior() *DamageableBaseBehavior {
        return this.DamageableBaseBehavior
}


func (this *CombatBaseBehavior) GetDestroyableBaseBehavior() IDestroyableBaseBehavior {
        return this.DestroyableBaseBehavior
}


func (this *CombatBaseBehavior) IsReadyForExplosion() bool {
        return this.ReadyForExplosion
}


func (this *CombatBaseBehavior) SetReadyForExplosion(readyForExplosion bool) {
        this.ReadyForExplosion = readyForExplosion
}

