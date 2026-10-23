/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type TransformInfoObjectConfigAndManipulatorFactoryBase struct {
    CommonStrings *CommonStrings
}

// Implements TransformInfoObjectConfigAndManipulatorFactoryInterface

// The Java class TransformInfoObjectConfigAndManipulatorFactoryBase has subclasses so its variables are this interface
type ITransformInfoObjectConfigAndManipulatorFactoryBase interface {
    TransformInfoObjectConfigAndManipulatorFactoryBase_() *TransformInfoObjectConfigAndManipulatorFactoryBase
    GetInstance(args ...any) TransformInfoObjectConfigInterface
}

func (this *TransformInfoObjectConfigAndManipulatorFactoryBase) TransformInfoObjectConfigAndManipulatorFactoryBase_() *TransformInfoObjectConfigAndManipulatorFactoryBase {
    return this
}


func (this *TransformInfoObjectConfigAndManipulatorFactoryBase) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewTransformInfoObjectConfigAndManipulatorFactoryBase() *TransformInfoObjectConfigAndManipulatorFactoryBase {
    this := &TransformInfoObjectConfigAndManipulatorFactoryBase{}
    this.initFields()
    return this
}


// throws Exception
func (this *TransformInfoObjectConfigAndManipulatorFactoryBase) GetInstance_1(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, objectConfigFileAbPath IAbPath) TransformInfoObjectConfigInterface {
        panic(errors.New(fmt.Sprint()))
}

// Java overloads of TransformInfoObjectConfigAndManipulatorFactoryBase.getInstance
func (this *TransformInfoObjectConfigAndManipulatorFactoryBase) GetInstance(args ...any) TransformInfoObjectConfigInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[IAbPath](args[2], pass) {
            return this.GetInstance_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[IAbPath](args[2]))
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) {
            return this.GetInstance_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]))
        }
        if len(args) == 3 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[TransformInfoInterface](args[1], pass) && Java_IsArg[*Document](args[2], pass) {
            return this.GetInstance_3(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[TransformInfoInterface](args[1]), Java_Arg[*Document](args[2]))
        }
    }
    panic("No Java overload of TransformInfoObjectConfigAndManipulatorFactoryBase.getInstance for the arguments")
}


// throws Exception
func (this *TransformInfoObjectConfigAndManipulatorFactoryBase) GetInstance_2(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) TransformInfoObjectConfigInterface {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *TransformInfoObjectConfigAndManipulatorFactoryBase) GetInstance_3(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface, document *Document) TransformInfoObjectConfigInterface {
        panic(errors.New(fmt.Sprint()))
}

