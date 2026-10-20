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

type Health struct {
    Health int
    MaxHealth int
    HealthListenerInterface HealthListenerInterface
}

// Implements HealthInterface

var Health_NULL_HEALTH *Health = NewHealth(0)
var Health_NAME string = "Health Resource"
var Health_HEALTH_LABEL string = "Health: "
var Health_MAX_HEALTH_LABEL string = " Max: "

func (this *Health) initFields() {
    
    
    this.HealthListenerInterface = NullHealthListener_NULL_HEALTH_LISTENER
}


func NewHealth(maxHealth int) *Health {
    this := &Health{}
    this.initFields()
        this.SetMaxHealth(maxHealth)
        this.SetHealth(maxHealth)
    return this
}


func (this *Health) Heal(ahealth int) {
        var newHealth int = this.GetHealth() + ahealth; _ = newHealth
        if newHealth > this.GetMaxHealth() {
            this.SetHealth(this.GetMaxHealth())
        } else {
            this.SetHealth(newHealth)
        }
}


func (this *Health) HealMax() {
        this.SetHealth(this.GetMaxHealth())
}


func (this *Health) IsDamaged() bool {
        if this.GetHealth() != this.GetMaxHealth() {
            return true
        } else {
            return false
        }
}


func (this *Health) Damage(ahealth int) {
        if ahealth < 0 {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("******* Trying to heal when damaging. Damage: ")
            stringBuffer.Appendint(ahealth)
            stringBuffer.Append(CommonSeps_GetInstance().SPACE)
            stringBuffer.Append(this.ToString())
            ForcedLogUtil_Log(fmt.Sprint(stringBuffer), this)
            return
        }
        var health int = this.GetHealth() - ahealth; _ = health
        this.SetHealth(health)
}


func (this *Health) IsAlive() bool {
        if this.GetHealth() <= 0 {
            return false
        } else {
            return true
        }
}


func (this *Health) GetMaxHealth() int {
        return this.MaxHealth
}


func (this *Health) SetMaxHealth(maxHealth int) {
        this.MaxHealth = maxHealth
}


func (this *Health) SetHealth(health int) {
        if health > this.GetMaxHealth() {
            this.Health = this.GetMaxHealth()
        } else if health < 0 {
            this.Health = 0
        } else {
            this.Health = health
        }
        if this.HealthListenerInterface != nil {
            this.HealthListenerInterface.OnHealthChange()
        }
}


func (this *Health) GetHealth() int {
        return this.Health
}


func (this *Health) AddListener(healthGraphic HealthListenerInterface) {
        this.HealthListenerInterface = healthGraphic
}


func (this *Health) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(Health_HEALTH_LABEL)
        stringBuffer.Appendint(this.Health)
        stringBuffer.Append(Health_MAX_HEALTH_LABEL)
        stringBuffer.Appendint(this.MaxHealth)
        return fmt.Sprint(stringBuffer)
}

func (this *Health) String() string {
    return this.ToString()
}

