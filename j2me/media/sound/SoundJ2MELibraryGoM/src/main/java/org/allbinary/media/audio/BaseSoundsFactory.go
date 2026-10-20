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

type BaseSoundsFactory struct {
    Initialized bool
}

// Implements SoundsFactoryInterface

// The Java class BaseSoundsFactory has subclasses so its variables are this interface
type IBaseSoundsFactory interface {
    BaseSoundsFactory_() *BaseSoundsFactory
    Init()
    GetSoundInterfaceArray() []ISound
    SetInitialized(initialized bool)
    IsInitialized() bool
}

func (this *BaseSoundsFactory) BaseSoundsFactory_() *BaseSoundsFactory {
    return this
}


func (this *BaseSoundsFactory) initFields() {
    
}

func NewBaseSoundsFactory() *BaseSoundsFactory {
    this := &BaseSoundsFactory{}
    this.initFields()
    return this
}


func (this *BaseSoundsFactory) Init() {
}


// throws Exception
func (this *BaseSoundsFactory) GetSoundInterfaceArray() []ISound {
        return make([]ISound, 0)
}


func (this *BaseSoundsFactory) SetInitialized(initialized bool) {
        this.Initialized = initialized
}


func (this *BaseSoundsFactory) IsInitialized() bool {
        return this.Initialized
}

