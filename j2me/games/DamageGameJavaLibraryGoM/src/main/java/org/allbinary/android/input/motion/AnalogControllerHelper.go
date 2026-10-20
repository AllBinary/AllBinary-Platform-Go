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

type AnalogControllerHelper struct {
    ReduceTimeFactor int
    SCALE_VALUE int
    RightTimeElapsedHelper ITimeDelayHelper
    LeftTimeElapsedHelper ITimeDelayHelper
    DownTimeElapsedHelper ITimeDelayHelper
    UpTimeElapsedHelper ITimeDelayHelper
}


func (this *AnalogControllerHelper) initFields() {
    
    this.SCALE_VALUE = AnalogControllerConfigurationFactory_GetInstance().SCALE_VALUE
    this.RightTimeElapsedHelper = NewTimeDelayHelper(0)
    this.LeftTimeElapsedHelper = NewTimeDelayHelper(0)
    this.DownTimeElapsedHelper = NewTimeDelayHelper(0)
    this.UpTimeElapsedHelper = NewTimeDelayHelper(0)
}


func NewAnalogControllerHelper(reduceTimeFactor int) *AnalogControllerHelper {
    this := &AnalogControllerHelper{}
    this.initFields()
        this.ReduceTimeFactor = reduceTimeFactor
    return this
}


// throws Exception
func (this *AnalogControllerHelper) Right(collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer, xAnalogValue int) {
        this.RightTimeElapsedHelper.TimeDelayHelper_().Delay = ((this.SCALE_VALUE - xAnalogValue) >> this.ReduceTimeFactor)
        if this.RightTimeElapsedHelper.IsTimeTNT() {
            collidableDestroyableDamageableLayer.Right()
        }
}


// throws Exception
func (this *AnalogControllerHelper) Left(collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer, xAnalogValue int) {
        this.LeftTimeElapsedHelper.TimeDelayHelper_().Delay = ((this.SCALE_VALUE + xAnalogValue) >> this.ReduceTimeFactor)
        if this.LeftTimeElapsedHelper.IsTimeTNT() {
            collidableDestroyableDamageableLayer.Left()
        }
}


// throws Exception
func (this *AnalogControllerHelper) Up(collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer, xAnalogValue int) {
        this.UpTimeElapsedHelper.TimeDelayHelper_().Delay = ((this.SCALE_VALUE - xAnalogValue) >> this.ReduceTimeFactor)
        if this.UpTimeElapsedHelper.IsTimeTNT() {
            collidableDestroyableDamageableLayer.Up()
        }
}


// throws Exception
func (this *AnalogControllerHelper) Down(collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer, xAnalogValue int) {
        this.DownTimeElapsedHelper.TimeDelayHelper_().Delay = ((this.SCALE_VALUE + xAnalogValue) >> this.ReduceTimeFactor)
        if this.DownTimeElapsedHelper.IsTimeTNT() {
            collidableDestroyableDamageableLayer.Down()
        }
}

