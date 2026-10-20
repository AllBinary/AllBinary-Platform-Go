/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type ModifierBaseProcessor struct {
}

// The Java class ModifierBaseProcessor has subclasses so its variables are this interface
type IModifierBaseProcessor interface {
    ModifierBaseProcessor_() *ModifierBaseProcessor
    UpdateColor(imageModifierUtil *ImageModifierUtil, originalImage IImage, imageToShow IImage, index int, basicColor *BasicColor)
    UpdateAlpha(imageModifierUtil *ImageModifierUtil, originalImage IImage, imageToShow IImage, index int, alpha int)
}

func (this *ModifierBaseProcessor) ModifierBaseProcessor_() *ModifierBaseProcessor {
    return this
}

var ModifierBaseProcessor_Instance IModifierBaseProcessor = NewModifierBaseProcessor()

func NewModifierBaseProcessor() *ModifierBaseProcessor {
    this := &ModifierBaseProcessor{}
    return this
}


func ModifierBaseProcessor_GetInstance() IModifierBaseProcessor {
        return ModifierBaseProcessor_Instance
}


func (this *ModifierBaseProcessor) UpdateColor(imageModifierUtil *ImageModifierUtil, originalImage IImage, imageToShow IImage, index int, basicColor *BasicColor) {
}


func (this *ModifierBaseProcessor) UpdateAlpha(imageModifierUtil *ImageModifierUtil, originalImage IImage, imageToShow IImage, index int, alpha int) {
}

