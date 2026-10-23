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

type AbeFactory struct {
    LogUtil *LogUtil
    UseCustomLoader bool
}

var AbeFactory_Instance *AbeFactory = NewAbeFactory()

func (this *AbeFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.UseCustomLoader = false
}


func AbeFactory_GetInstance() *AbeFactory {
        return AbeFactory_Instance
}


func NewAbeFactory() *AbeFactory {
    this := &AbeFactory{}
    this.initFields()
    return this
}


// throws LicensingException
// synchronized
func (this *AbeFactory) GetInstance_1(abeClientInformation AbeClientInformationInterface, className string) any {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADERERROR) {
                this.LogUtil.Put("Failure for: " + className, this, "getInstance(classname)", e)
            }
            panic(e)
                }
            }()
            tryControl1, tryResult1 = 1, this.GetClass(abeClientInformation, className).NewInstance(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of AbeFactory.getInstance
func (this *AbeFactory) GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetInstance_1(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 4 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[[]reflect.Type](args[2], pass) && Java_IsArg[[]any](args[3], pass) {
            return this.GetInstance_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[string](args[1]), Java_Arg[[]reflect.Type](args[2]), Java_Arg[[]any](args[3]))
        }
    }
    panic("No Java overload of AbeFactory.getInstance for the arguments")
}


// throws LicensingException
// synchronized
func (this *AbeFactory) GetInstance_2(abeClientInformation AbeClientInformationInterface, className string, classes []reflect.Type, params []any) any {
        var constructor *Constructor = nil; _ = constructor
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADERERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Failure for: ")
                stringBuffer.Append(className)
                stringBuffer.Append(CommonSeps_GetInstance().SPACE)
                stringBuffer.Append(ConstructorUtil_View(constructor, CommonSeps_GetInstance().NEW_LINE))
                this.LogUtil.Put(fmt.Sprint(stringBuffer), this, "getInstance(className,params)", e)
            }
            panic(e)
                }
            }()
            var parent *ClassLoader = WebappClassLoaderInfo_GetLoader(); _ = parent
            if this.UseCustomLoader {
                var loader *ClassLoader = NewAbeClassLoader(parent, AbKeys_GetInstance().GetKey(abeClientInformation, className)); _ = loader
                var myClass reflect.Type = loader.LoadClass(className); _ = myClass
                constructor = myClass.GetConstructor(classes)
                tryControl1, tryResult1 = 1, constructor.NewInstance(params); return
            } else {
                var myClass reflect.Type = parent.LoadClass(className); _ = myClass
                constructor = myClass.GetConstructor(classes)
                tryControl1, tryResult1 = 1, constructor.NewInstance(params); return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
// synchronized
func (this *AbeFactory) GetClass(abeClientInformation AbeClientInformationInterface, className string) reflect.Type {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 reflect.Type) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADERERROR) {
                this.LogUtil.Put("Failure for: " + className, this, "getClass(className)", e)
            }
            panic(e)
                }
            }()
            var parent *ClassLoader = WebappClassLoaderInfo_GetLoader(); _ = parent
            if this.UseCustomLoader {
                var loader *ClassLoader = NewAbeClassLoader(parent, AbKeys_GetInstance().GetKey(abeClientInformation, className)); _ = loader
                var c reflect.Type = loader.LoadClass(className); _ = c
                tryControl1, tryResult1 = 1, c; return
            } else {
                var c reflect.Type = parent.LoadClass(className); _ = c
                tryControl1, tryResult1 = 1, c; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

