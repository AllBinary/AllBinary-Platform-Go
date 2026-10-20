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

type RootStoreThemeCategoryProperties struct {
    RootStoreCategoryProperties
}





// Java overloads of new RootStoreThemeCategoryProperties
func NewRootStoreThemeCategoryProperties(args ...any) *RootStoreThemeCategoryProperties {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewRootStoreThemeCategoryProperties_1(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) {
            return NewRootStoreThemeCategoryProperties_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[IAbPath](args[1]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewRootStoreThemeCategoryProperties_3(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*Node](args[1]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewRootStoreThemeCategoryProperties_4(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*HashMap](args[1]))
        }
    }
    panic("No Java overload of new RootStoreThemeCategoryProperties for the arguments")
}
func NewRootStoreThemeCategoryProperties_1(transformInfoInterface TransformInfoInterface) *RootStoreThemeCategoryProperties {
    this := &RootStoreThemeCategoryProperties{}
        this.RootStoreCategoryProperties = *NewRootStoreCategoryProperties(transformInfoInterface)
    return this
}


func NewRootStoreThemeCategoryProperties_2(transformInfoInterface TransformInfoInterface, abPath IAbPath) *RootStoreThemeCategoryProperties {
    this := &RootStoreThemeCategoryProperties{}
        this.RootStoreCategoryProperties = *NewRootStoreCategoryProperties(transformInfoInterface, abPath)
    return this
}


func NewRootStoreThemeCategoryProperties_3(transformInfoInterface TransformInfoInterface, node *Node) *RootStoreThemeCategoryProperties {
    this := &RootStoreThemeCategoryProperties{}
        this.RootStoreCategoryProperties = *NewRootStoreCategoryProperties(transformInfoInterface, node)
    return this
}


func NewRootStoreThemeCategoryProperties_4(transformInfoInterface TransformInfoInterface, categoryPropertiesHashMap *HashMap) *RootStoreThemeCategoryProperties {
    this := &RootStoreThemeCategoryProperties{}
        this.RootStoreCategoryProperties = *NewRootStoreCategoryProperties(transformInfoInterface, categoryPropertiesHashMap)
    return this
}


// throws Exception
func (this *RootStoreThemeCategoryProperties) InitPath() {
        var transformInfoHttpStoreInterface TransformInfoHttpStoreInterface = any(this.TransformInfoInterface).(TransformInfoHttpStoreInterface); _ = transformInfoHttpStoreInterface
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(transformInfoHttpStoreInterface.GetStoreName()); _ = storeFrontInterface
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append(RootStoreCategoryProperties_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(storeFrontInterface.GetName())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        stringBuffer.Append(RootStoreCategoryProperties_FREEBLISKET_PATH_GLOBALS.GetInstance().THEMEPATH)
        var abPath IAbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = abPath
        var httpServletRequest *HttpServletRequest = any(transformInfoHttpStoreInterface.GetPageContext().GetRequest()).(*HttpServletRequest); _ = httpServletRequest
        this.WebAppAbPath = NewAbPath(fmt.Sprint(httpServletRequest.GetContextPath()) + fmt.Sprint(abPath), StringUtil_GetInstance().EMPTY_STRING)
        this.SetRootFilePath(NewAbPath(fmt.Sprint(RootStoreCategoryProperties_URLGLOBALS.GetMainPath()) + fmt.Sprint(abPath), StringUtil_GetInstance().EMPTY_STRING))
}

