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

type TransformInfoHttpStore struct {
    TransformInfoHttp
    LogUtil *LogUtil
}

// Implements TransformInfoHttpStoreInterface

// The Java class TransformInfoHttpStore has subclasses so its variables are this interface
type ITransformInfoHttpStore interface {
    ITransformInfoHttp
    TransformInfoHttpStore_() *TransformInfoHttpStore
    GetPath() string
    GetTemplateFilePath() IAbPath
    GetObjectConfigFilePath() IAbPath
    GetDataFilePath() IAbPath
}

func (this *TransformInfoHttpStore) TransformInfoHttpStore_() *TransformInfoHttpStore {
    return this
}


func (this *TransformInfoHttpStore) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new TransformInfoHttpStore
func NewTransformInfoHttpStore(args ...any) *TransformInfoHttpStore {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return NewTransformInfoHttpStore_1(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 3 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) && Java_IsArg[*PageContext](args[2], pass) {
            return NewTransformInfoHttpStore_2(Java_Arg[*HashMap](args[0]), Java_Arg[*HashMap](args[1]), Java_Arg[*PageContext](args[2]))
        }
        if len(args) == 3 && Java_IsArg[StoreFrontInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) && Java_IsArg[*PageContext](args[2], pass) {
            return NewTransformInfoHttpStore_3(Java_Arg[StoreFrontInterface](args[0]), Java_Arg[*HashMap](args[1]), Java_Arg[*PageContext](args[2]))
        }
    }
    panic("No Java overload of new TransformInfoHttpStore for the arguments")
}
func NewTransformInfoHttpStore_1(propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttpStore {
    this := &TransformInfoHttpStore{}
        this.TransformInfoHttp = *NewTransformInfoHttp(propertiesHashMap, pageContext, true)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func NewTransformInfoHttpStore_2(databaseHashMap *HashMap, propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttpStore {
    this := &TransformInfoHttpStore{}
        this.TransformInfoHttp = *NewTransformInfoHttp(databaseHashMap, propertiesHashMap, pageContext)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Database HashMap: ")
            stringBuffer.Append(fmt.Sprint(databaseHashMap))
            stringBuffer.Append("\nProperties HashMap: ")
            stringBuffer.Append(fmt.Sprint(propertiesHashMap))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "Constructor(HashMap, HashMap , PageContext)")
        }
        this.SetStoreName(any(databaseHashMap.Get(StoreFrontData_GetInstance().NAME)).(string))
    return this
}


func NewTransformInfoHttpStore_3(storeFrontInterface StoreFrontInterface, propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttpStore {
    this := &TransformInfoHttpStore{}
        this.TransformInfoHttp = *NewTransformInfoHttp(propertiesHashMap, pageContext)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.SetStoreName(storeFrontInterface.GetName())
    return this
}


// throws Exception
func (this *TransformInfoHttpStore) GetPath() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(TransformInfoHttp_URLGLOBALS.GetMainPath())
        stringBuffer.Append(TransformInfoHttp_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(this.GetStoreName())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        return fmt.Sprint(stringBuffer)
}


// throws Exception
func (this *TransformInfoHttpStore) GetTemplateFilePath() IAbPath {
        return NewAbPath(this.GetPath(), this.GetTemplateFile())
}


// throws Exception
func (this *TransformInfoHttpStore) GetObjectConfigFilePath() IAbPath {
        return NewAbPath(this.GetPath(), this.GetObjectConfigFile())
}


// throws Exception
func (this *TransformInfoHttpStore) GetDataFilePath() IAbPath {
        return NewAbPath(this.GetPath(), this.GetDataFile())
}

