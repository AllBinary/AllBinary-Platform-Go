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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type MultiKeyPressesAI struct {
    KeyPressesAI
    ToggleTimeHelper ITimeDelayHelper
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
}


func (this *MultiKeyPressesAI) initFields() {
    this.ToggleTimeHelper = NewTimeDelayHelper(600)
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
}


func NewMultiKeyPressesAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *MultiKeyPressesAI {
    this := &MultiKeyPressesAI{}
        this.KeyPressesAI = *NewKeyPressesAI(hashtable, ownerLayerInterface, gameInput)
        this.initFields()
        this.Update()
    return this
}


// throws Exception
func (this *MultiKeyPressesAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
}


// throws Exception
func (this *MultiKeyPressesAI) Update() {
        var gameInput IGameInput = this.GetGameInput(); _ = gameInput
        var keys []int = this.Keys; _ = keys
        var size int = len(keys); _ = size
        if this.IsOn() {
            for index := int(0); index < size; index++ {
                var key int = int(keys[index]); _ = key
                if key != -1 {
                    gameInput.Add(this.GameKeyEventFactory.GetInstanceForKey(this, key))
                }
            }
        } else {
            for index := int(0); index < size; index++ {
                var key int = int(keys[index]); _ = key
                if key != -1 {
                    gameInput.AddForRemoval(this.GameKeyEventFactory.GetInstanceForKey(this, key))
                }
            }
        }
}


// throws Exception
func (this *MultiKeyPressesAI) Toggle() {
        if this.ToggleTimeHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            this.KeyPressesAI.Toggle()
            this.Update()
        }
}


// throws Exception
func (this *MultiKeyPressesAI) Disable() {
        if this.ToggleTimeHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            this.KeyPressesAI.Disable()
            this.Update()
        }
}


// throws Exception
func (this *MultiKeyPressesAI) Enable() {
        if this.ToggleTimeHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            this.KeyPressesAI.Enable()
            this.Update()
        }
}

