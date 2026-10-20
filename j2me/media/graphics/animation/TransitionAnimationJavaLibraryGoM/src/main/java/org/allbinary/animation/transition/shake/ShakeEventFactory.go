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

type ShakeEventFactory struct {
    MICRO *ShakeEvent
    SMALL *ShakeEvent
    MEDIUM *ShakeEvent
    LARGE *ShakeEvent
}

var ShakeEventFactory_Instance *ShakeEventFactory = NewShakeEventFactory()

func (this *ShakeEventFactory) initFields() {
    this.MICRO = NewShakeEvent()
    this.SMALL = NewShakeEvent()
    this.MEDIUM = NewShakeEvent()
    this.LARGE = NewShakeEvent()
}

func NewShakeEventFactory() *ShakeEventFactory {
    this := &ShakeEventFactory{}
    this.initFields()
    return this
}


func ShakeEventFactory_GetInstance() *ShakeEventFactory {
        return ShakeEventFactory_Instance
}


func (this *ShakeEventFactory) Init() {
}

