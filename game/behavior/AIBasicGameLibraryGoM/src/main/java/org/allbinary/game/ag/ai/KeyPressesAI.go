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

type KeyPressesAI struct {
    BasicAI
    On bool
    Keys []int
}

// The Java class KeyPressesAI has subclasses so its variables are this interface
type IKeyPressesAI interface {
    IBasicAI
    KeyPressesAI_() *KeyPressesAI
    ProcessAI(allBinaryLayerManager IAllBinaryLayerManager)
    Toggle()
    Disable()
    Enable()
    SetOn(on bool)
    IsOn() bool
}

func (this *KeyPressesAI) KeyPressesAI_() *KeyPressesAI {
    return this
}


func (this *KeyPressesAI) initFields() {
    this.On = true
    
}


func NewKeyPressesAI(hashtable *ABHashtable[any, any], ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *KeyPressesAI {
    this := &KeyPressesAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
        this.Keys = make([]int, hashtable.Size())
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
        var size int = len(objectArray); _ = size
        for index := int(0); index < size; index++ {
            this.Keys[index] = any(hashtable.Get(any(objectArray[index]).(any))).(int)
        }
    return this
}


// throws Exception
func (this *KeyPressesAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.On {
            for index := int(0); index < len(this.Keys); index++ {
                this.BasicAI.ProcessKeyAI(int(this.Keys[index]))
            }
        }
}


// throws Exception
func (this *KeyPressesAI) Toggle() {
        if this.IsOn() {
            this.SetOn(false)
        } else {
            this.SetOn(true)
        }
}


// throws Exception
func (this *KeyPressesAI) Disable() {
        this.SetOn(false)
}


// throws Exception
func (this *KeyPressesAI) Enable() {
        this.SetOn(true)
}


func (this *KeyPressesAI) SetOn(on bool) {
        this.On = on
}


func (this *KeyPressesAI) IsOn() bool {
        return this.On
}

