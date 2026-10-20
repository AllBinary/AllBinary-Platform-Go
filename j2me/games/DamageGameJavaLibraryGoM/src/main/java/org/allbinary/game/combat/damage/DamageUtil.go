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

type DamageUtil struct {
    LogUtil *LogUtil
    COLLIDE string
}

var DamageUtil_Instance *DamageUtil = NewDamageUtil()

func (this *DamageUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.COLLIDE = "collide"
}


func DamageUtil_GetInstance() *DamageUtil {
        return DamageUtil_Instance
}


func NewDamageUtil() *DamageUtil {
    this := &DamageUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *DamageUtil) Process(damageableInterface DamageableInterface, damageableInterface2 DamageableInterface) {
        damageableInterface.Damage(damageableInterface2.GetDamage(0), 0)
}


// throws Exception
func (this *DamageUtil) DebugDamage(collidableInterfaceCompositeInterface1 ICollidableCompositeLayer, collidableInterfaceCompositeInterface ICollidableCompositeLayer) {
        var collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer = any(collidableInterfaceCompositeInterface).(ICollidableDestroyableDamageableLayer); _ = collidableDestroyableDamageableLayer
        var damage int = collidableDestroyableDamageableLayer.GetDamage(0); _ = damage
        if damage > 20000 {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Collision ")
            stringBuffer.Append(fmt.Sprint(collidableInterfaceCompositeInterface1))
            stringBuffer.Append(" with: ")
            stringBuffer.Append(fmt.Sprint(collidableInterfaceCompositeInterface))
            stringBuffer.Append(" damage: ")
            stringBuffer.Appendint(damage)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), collidableInterfaceCompositeInterface1, this.COLLIDE)
        }
}

