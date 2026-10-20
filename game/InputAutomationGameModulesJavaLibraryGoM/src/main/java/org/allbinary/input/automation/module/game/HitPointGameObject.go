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

type HitPointGameObject struct {
    MeteredGameObject
    Damaged bool
    CloseDeath bool
    HP_LABEL string
    IS_DAMAGED_LABEL string
    IS_CLOSE_DEATH_LABEL string
}

// Implements HitPointGameObjectInterface


func (this *HitPointGameObject) initFields() {
    this.Damaged = false
    this.CloseDeath = false
    this.HP_LABEL = "HP: "
    this.IS_DAMAGED_LABEL = "isDamaged: "
    this.IS_CLOSE_DEATH_LABEL = "isCloseDeath: "
}


func NewHitPointGameObject(bufferedImage []*BufferedImage, minX int, maxX int, y int) *HitPointGameObject {
    this := &HitPointGameObject{}
        this.MeteredGameObject = *NewMeteredGameObject(bufferedImage, minX, maxX, y)
        this.initFields()
        if float64(this.GetRatios()[len(bufferedImage) - 1]) < float64(1.0) {
            this.SetDamaged(true)
        } else {
            this.SetDamaged(false)
        }
        if float64(this.GetRatios()[len(bufferedImage) - 1]) > float64(0.25) {
            this.CloseDeath = false
        } else {
            this.CloseDeath = true
        }
    return this
}


func (this *HitPointGameObject) IsMoreThan(aFloat float32) bool {
        if float64(this.GetRatios()[len(this.GetRatios()) - 1]) > float64(aFloat) {
            return true
        } else {
            return false
        }
}


func (this *HitPointGameObject) IsCloseDeath() bool {
        return this.CloseDeath
}


func (this *HitPointGameObject) SetCloseDeath(closeDeath bool) {
        this.CloseDeath = closeDeath
}


func (this *HitPointGameObject) IsDamaged() bool {
        return this.Damaged
}


func (this *HitPointGameObject) SetDamaged(damaged bool) {
        this.Damaged = damaged
}


func (this *HitPointGameObject) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.HP_LABEL)
        stringBuffer.Append(this.MeteredGameObject.ToString())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(this.IS_DAMAGED_LABEL)
        stringBuffer.Appendboolean(this.IsDamaged())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(this.IS_CLOSE_DEATH_LABEL)
        stringBuffer.Appendboolean(this.IsCloseDeath())
        return fmt.Sprint(stringBuffer)
}

func (this *HitPointGameObject) String() string {
    return this.ToString()
}

