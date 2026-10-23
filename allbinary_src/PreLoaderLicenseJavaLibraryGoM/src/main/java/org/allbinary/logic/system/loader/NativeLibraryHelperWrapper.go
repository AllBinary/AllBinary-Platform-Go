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

type NativeLibraryHelperWrapper struct {
    LibraryName string
    Loader *ClassLoader
    AnyType any
}

// The Java class NativeLibraryHelperWrapper has subclasses so its variables are this interface
type INativeLibraryHelperWrapper interface {
    NativeLibraryHelperWrapper_() *NativeLibraryHelperWrapper
    Load()
    Unload()
    ToString() string
    GetLibraryName() string
    SetLibraryName(libraryName string)
}

func (this *NativeLibraryHelperWrapper) NativeLibraryHelperWrapper_() *NativeLibraryHelperWrapper {
    return this
}


func (this *NativeLibraryHelperWrapper) initFields() {
    
    
    
}


func NewNativeLibraryHelperWrapper(libraryName string) *NativeLibraryHelperWrapper {
    this := &NativeLibraryHelperWrapper{}
    this.initFields()
        this.SetLibraryName(libraryName)
    return this
}


// throws Exception
func (this *NativeLibraryHelperWrapper) Load() {
        var params []any = make([]any, 1); _ = params
        params[0] = this.GetLibraryName()
        var classes []reflect.Type = make([]reflect.Type, 1); _ = classes
        classes[0] = reflect.TypeOf(this.GetLibraryName())
        this.Loader = NewNativeLibraryClassLoader(reflect.TypeOf(this).GetClassLoader())
        var myClass reflect.Type = this.Loader.LoadClass("dynamic.NativeLibraryHelper"); _ = myClass
        var constructor *Constructor = myClass.GetConstructor(classes); _ = constructor
        this.AnyType = constructor.NewInstance(params)
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var method *Method = reflect.TypeOf(this.AnyType).GetMethod(commonStrings.LOAD, nil); _ = method
        method.Invoke(this.AnyType, nil)
}


func (this *NativeLibraryHelperWrapper) Unload() {
        this.AnyType = nil
        this.Loader = nil
        runtime.GC()
        runtime.GC()
}


func (this *NativeLibraryHelperWrapper) ToString() string {
        return "Library: " + fmt.Sprint(this.GetLibraryName())
}

func (this *NativeLibraryHelperWrapper) String() string {
    return this.ToString()
}


func (this *NativeLibraryHelperWrapper) GetLibraryName() string {
        return this.LibraryName
}


func (this *NativeLibraryHelperWrapper) SetLibraryName(libraryName string) {
        this.LibraryName = libraryName
}

