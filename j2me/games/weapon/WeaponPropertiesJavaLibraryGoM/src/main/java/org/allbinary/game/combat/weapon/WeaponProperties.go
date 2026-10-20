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

type WeaponProperties struct {
    SimpleWeaponProperties
    ReloadTime int64
    TargetingTime int64
    Speed *BasicDecimal
    MAX int64
    ZERO int16
}

var WeaponProperties_NULL_WEAPON_PROPERTIES *WeaponProperties = NewWeaponProperties(0, 0, 0, 0, int16(0))
var WeaponProperties_MessageSent bool = false
var WeaponProperties_DAMAGE string = "Damage: "
var WeaponProperties_RANGE string = "Range: "
var WeaponProperties_RELOAD string = "Reload: "

func (this *WeaponProperties) initFields() {
    
    
    this.Speed = BasicDecimal_ZERO_BIGDECIMAL
    this.MAX = 10240
    this.ZERO = 0
}


func NewWeaponProperties(reloadTime int64, targetingTime int64, speed int64, damage int, dissipation int16) *WeaponProperties {
    this := &WeaponProperties{}
    this.SimpleWeaponProperties = *NewSimpleWeaponProperties()
    this.initFields()
        if speed < this.MAX && speed != 0 && !WeaponProperties_MessageSent {
            var MESSAGE string = "Danger Danger Danger: Speed probably to slow if using 1 degree calculations as velocity for a single axis could be below 1024: "; _ = MESSAGE
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(MESSAGE).Appendlong(speed)), this, commonStrings.CONSTRUCTOR)
            WeaponProperties_MessageSent = true
        }
        this.SetReloadTime(reloadTime)
        this.SetTargetingTime(targetingTime)
        this.SetDamage(damage)
        this.SetDissipation(dissipation)
        this.SetSpeed(NewBasicDecimal(speed))
        if int(dissipation) != int(this.ZERO) {
            var unscaledDamage int64 = this.Speed.GetUnscaled() * int64(damage); _ = unscaledDamage
            var scaledDissipation int = int(dissipation) * this.Speed.GetScaledFactorValue(); _ = scaledDissipation
            var value int64 = (unscaledDamage / int64(scaledDissipation)); _ = value
            this.SetRange(int((value * 9)) / 10)
        }
    return this
}


func (this *WeaponProperties) SetReloadTime(reloadTime int64) {
        this.ReloadTime = reloadTime
}


func (this *WeaponProperties) GetReloadTime() int64 {
        return this.ReloadTime
}


func (this *WeaponProperties) SetTargetingTime(targetingTime int64) {
        this.TargetingTime = targetingTime
}


func (this *WeaponProperties) GetTargetingTime() int64 {
        return this.TargetingTime
}


func (this *WeaponProperties) GetSpeed() *BasicDecimal {
        return this.Speed
}


func (this *WeaponProperties) SetSpeed(speed *BasicDecimal) {
        this.Speed = speed
}


func (this *WeaponProperties) GetDamageI(range_ int) int {
        return this.GetDamage() - ((int(this.GetDissipation()) * range_) / this.Speed.GetScaled())
}


func (this *WeaponProperties) ToStringArray() []string {
        var index int = 0; _ = index
        var stringArray []string = make([]string, 3); _ = stringArray
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringArray[Java_PostIncrement(&index)] = fmt.Sprint(stringBuffer.Append(WeaponProperties_DAMAGE).Appendint(this.GetDamage()))
        stringBuffer.Delete(0, stringBuffer.Length())
        stringArray[Java_PostIncrement(&index)] = fmt.Sprint(stringBuffer.Append(WeaponProperties_RANGE).Appendint(this.GetRange()))
        stringBuffer.Delete(0, stringBuffer.Length())
        stringArray[Java_PostIncrement(&index)] = fmt.Sprint(stringBuffer.Append(WeaponProperties_RELOAD).Appendlong(this.GetReloadTime()))
        return stringArray
}


func (this *WeaponProperties) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(WeaponProperties_DAMAGE).Appendint(this.GetDamage())
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(WeaponProperties_RANGE).Appendint(this.GetRange())
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(WeaponProperties_RELOAD).Appendlong(this.GetReloadTime())
        return fmt.Sprint(stringBuffer)
}

func (this *WeaponProperties) String() string {
    return this.ToString()
}

