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

type AbeClassLoader struct {
    ClassLoader
    LogUtil *LogUtil
    AbCryptUtil *AbCryptUtil
    LOAD_CLASS string
    ENCRYPTED_EXTENSION string
    Key string
}

var AbeClassLoader_Classes Map = StdUtil_GetInstance().CreateHashMap()
var AbeClassLoader_PATH string

func (this *AbeClassLoader) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbCryptUtil = AbCryptUtil_GetInstance()
    this.LOAD_CLASS = "loadClass"
    this.ENCRYPTED_EXTENSION = fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + "abc"
    
}


func NewAbeClassLoader(parent *ClassLoader, key string) *AbeClassLoader {
    this := &AbeClassLoader{}
        this.ClassLoader = *NewClassLoader(parent)
        this.initFields()
        this.Key = key
        AbeClassLoader_PATH = fmt.Sprint(org_Allbinary.Globals.URLGLOBALS.GetWebappPath()) + "WEB-INF/classes/"
    return this
}


func (this *AbeClassLoader) FindLoadedClass1(name string) reflect.Type {
        return any(AbeClassLoader_Classes.Get(name)).(reflect.Type)
}


// throws ClassNotFoundException
// synchronized
func (this *AbeClassLoader) LoadClass(name string) reflect.Type {
        return this.LoadClass(name, false)
}


// throws ClassNotFoundException
// Java overloaded method: Go needs a different name
// synchronized
func (this *AbeClassLoader) LoadClass(name string, resolve bool) reflect.Type {
        var loadedWith string = "findLoadedClass1"; _ = loadedWith
        var logBuffer *LogBuffer = NewLogBuffer(); _ = logBuffer
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 reflect.Type) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADERERROR) {
                logBuffer.LogAll()
                this.LogUtil.Put("Failure loading: " + name + "\nwith: " + loadedWith, this, this.LOAD_CLASS, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var myClass reflect.Type = this.FindLoadedClass1(name); _ = myClass
            if myClass == nil {
                if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 reflect.Type) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logBuffer.Add(LogFactory_GetInstance("Failed to Load Class: " + name + "\nwith: " + loadedWith, this, this.LOAD_CLASS, e))
                        }
                    }()
                    loadedWith = "super.loadClass"
                    var normalClass reflect.Type = this.ClassLoader.LoadClass(name, resolve); _ = normalClass
                    if normalClass != nil {
                        tryControl2, tryResult2 = 1, normalClass; return
                    }
                    return
                }(); tryControl2 == 1 {
                    tryControl1, tryResult1 = 1, tryResult2; return
                }
                var classBytes []int8 = this.LoadClassBytesFromFile(name); _ = classBytes
                if classBytes == nil {
                    panic(errors.New(fmt.Sprint("My Bytes not loaded for: " + name + "\nwith: " + "loadClassBytes")))
                }
                myClass = this.DefineClass(name, classBytes, 0, len(classBytes))
                if myClass == nil {
                    panic(errors.New(fmt.Sprint("My Class Not Defineable for: " + name + "\nwith: " + "loadClassBytes")))
                }
                AbeClassLoader_Classes.Put(name, myClass)
            } else {
                logBuffer.Add(LogFactory_GetInstanceF("Already Loaded: " + name + "\nwith: " + loadedWith, this, this.LOAD_CLASS))
            }
            this.ResolveClass(myClass)
            tryControl1, tryResult1 = 1, myClass; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AbeClassLoader) LoadClassBytesFromFile(name string) []int8 {
        var in *FileInputStream = nil; _ = in
        var cname string = StringUtil_GetInstance().EMPTY_STRING; _ = cname
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADERERROR) {
                this.LogUtil.Put("(Before LogBuffer Output) Failure loading Encrypted: " + name + " File: " + cname, "AbeClassLoader", "loadClassBytes", e)
            }
            StreamUtil_GetInstance().Close(in)
                }
            }()
            cname = fmt.Sprint(AbeClassLoader_PATH) + strings.ReplaceAll(name, string('.'), string(AbPathData_GetInstance().SEPARATORCHAR)) + this.ENCRYPTED_EXTENSION
            in = NewFileInputStream(cname)
            var decrypted []int8 = this.AbCryptUtil.Decrypt(in, this.Key); _ = decrypted
            tryControl1, tryResult1 = 1, decrypted; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


// throws ClassNotFoundException
// synchronized
func (this *AbeClassLoader) OldLoadClass(name string, resolve bool) reflect.Type {
        var loadedWith string = "findLoadedClass1"; _ = loadedWith
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 reflect.Type) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADERERROR) {
                this.LogUtil.Put("Failure loading: " + name, "AbeClassLoader", this.LOAD_CLASS, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var myClass reflect.Type = this.FindLoadedClass1(name); _ = myClass
            if myClass == nil {
                if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 reflect.Type) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADER) {
                        this.LogUtil.PutF("Failed Loaded Class: " + name + " with: " + loadedWith, this, this.LOAD_CLASS)
                    }
                        }
                    }()
                    loadedWith = "super.loadClass"
                    var normalClass reflect.Type = this.ClassLoader.LoadClass(name, resolve); _ = normalClass
                    if normalClass != nil {
                        tryControl2, tryResult2 = 1, normalClass; return
                    }
                    return
                }(); tryControl2 == 1 {
                    tryControl1, tryResult1 = 1, tryResult2; return
                }
                loadedWith = "loadClassBytes"
                var classBytes []int8 = this.LoadClassBytesFromFile(name); _ = classBytes
                if classBytes == nil {
                    panic(errors.New(fmt.Sprint("Bytes not loaded for: " + name)))
                }
                myClass = this.DefineClass(name, classBytes, 0, len(classBytes))
                if myClass == nil {
                    panic(errors.New(fmt.Sprint("Class Not Defineable for: " + name)))
                }
                AbeClassLoader_Classes.Put(name, myClass)
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().LOADER) {
                    this.LogUtil.PutF("Already Loaded: " + name + " with: " + loadedWith, this, this.LOAD_CLASS)
                }
            }
            this.ResolveClass(myClass)
            tryControl1, tryResult1 = 1, myClass; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

