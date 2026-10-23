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

type OverrideStoreTransformInfoObjectConfig struct {
    GenericStoreTransformInfoObjectConfig
}





// Java overloads of new OverrideStoreTransformInfoObjectConfig
func NewOverrideStoreTransformInfoObjectConfig(args ...any) *OverrideStoreTransformInfoObjectConfig {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) {
            return NewOverrideStoreTransformInfoObjectConfig_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[*Document](args[2], pass) {
            return NewOverrideStoreTransformInfoObjectConfig_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[*Document](args[2]))
        }
        if len(args) == 4 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) {
            return NewOverrideStoreTransformInfoObjectConfig_3(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]))
        }
    }
    panic("No Java overload of new OverrideStoreTransformInfoObjectConfig for the arguments")
}
func NewOverrideStoreTransformInfoObjectConfig_1(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) *OverrideStoreTransformInfoObjectConfig {
    this := &OverrideStoreTransformInfoObjectConfig{}
        this.GenericStoreTransformInfoObjectConfig = *NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface)
    return this
}


func NewOverrideStoreTransformInfoObjectConfig_2(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, document *Document) *OverrideStoreTransformInfoObjectConfig {
    this := &OverrideStoreTransformInfoObjectConfig{}
        this.GenericStoreTransformInfoObjectConfig = *NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface, document)
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


func NewOverrideStoreTransformInfoObjectConfig_3(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, name string, type_ string) *OverrideStoreTransformInfoObjectConfig {
    this := &OverrideStoreTransformInfoObjectConfig{}
        this.GenericStoreTransformInfoObjectConfig = *NewGenericStoreTransformInfoObjectConfig(abeClientInformation, transformInfoInterface, name, type_)
        this.SetDocument(this.Generate(this.ToXmlDoc()))
    return this
}


// throws Exception
func (this *OverrideStoreTransformInfoObjectConfig) Generate(objectConfigDocument *Document) *Document {
        var newObjectConfigDocument *Document = this.GenericStoreTransformInfoObjectConfig.Generate(objectConfigDocument); _ = newObjectConfigDocument
        return newObjectConfigDocument
}

