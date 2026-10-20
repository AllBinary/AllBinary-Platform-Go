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

type InputFeatureFactory struct {
    REMOVE_DUPLICATE_KEY_PRESSES *InputFeature
    INPUT_MAPPING *InputFeature
    MULTI_KEY_PRESS *InputFeature
    SINGLE_KEY_REPEAT_PRESS *InputFeature
    SINGLE_KEY_PRESS *InputFeature
}

var InputFeatureFactory_Instance *InputFeatureFactory = NewInputFeatureFactory()

func (this *InputFeatureFactory) initFields() {
    this.REMOVE_DUPLICATE_KEY_PRESSES = NewInputFeature("Remove Duplicate Key Presses")
    this.INPUT_MAPPING = NewInputFeature("Input Mapping")
    this.MULTI_KEY_PRESS = NewInputFeature("Multi Key Press")
    this.SINGLE_KEY_REPEAT_PRESS = NewInputFeature("Single Key Repeat Press")
    this.SINGLE_KEY_PRESS = NewInputFeature("Single Key Press")
}


func InputFeatureFactory_GetInstance() *InputFeatureFactory {
        return InputFeatureFactory_Instance
}


func NewInputFeatureFactory() *InputFeatureFactory {
    this := &InputFeatureFactory{}
    this.initFields()
    return this
}


func (this *InputFeatureFactory) IsSingleKeyProcessing() bool {
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(this.SINGLE_KEY_REPEAT_PRESS) || features.IsFeature(this.SINGLE_KEY_PRESS) {
            return true
        } else {
            return false
        }
}

