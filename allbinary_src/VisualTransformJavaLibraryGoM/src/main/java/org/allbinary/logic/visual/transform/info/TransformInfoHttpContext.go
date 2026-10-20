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

type TransformInfoHttpContext struct {
    TransformInfoHttp
    LogUtil *LogUtil
}


func (this *TransformInfoHttpContext) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewTransformInfoHttpContext(databaseHashMap *HashMap, propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttpContext {
    this := &TransformInfoHttpContext{}
        this.TransformInfoHttp = *NewTransformInfoHttp(databaseHashMap, propertiesHashMap, pageContext)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Database HashMap: ")
            stringBuffer.Append(fmt.Sprint(databaseHashMap))
            stringBuffer.Append("\nProperties HashMap: ")
            stringBuffer.Append(fmt.Sprint(propertiesHashMap))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.SetStoreName(StringUtil_GetInstance().EMPTY_STRING)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoHttpContext(propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoHttpContext {
    this := &TransformInfoHttpContext{}
        this.TransformInfoHttp = *NewTransformInfoHttp(propertiesHashMap, pageContext)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Properties HashMap: " + fmt.Sprint(propertiesHashMap), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.SetStoreName(StringUtil_GetInstance().EMPTY_STRING)
    return this
}


// throws Exception
func (this *TransformInfoHttpContext) GetPath() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(TransformInfoHttp_URLGLOBALS.GetMainPath())
        stringBuffer.Append(TransformInfoHttp_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(TransformInfoHttp_FREEBLISKET_PATH_GLOBALS.GetInstance().INSTALLPATH)
        stringBuffer.Append(this.GetStoreName())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        return fmt.Sprint(stringBuffer)
}


// throws Exception
func (this *TransformInfoHttpContext) GetTemplateFilePath() *AbPath {
        return NewAbPath(this.GetPath(), this.GetTemplateFile())
}


// throws Exception
func (this *TransformInfoHttpContext) GetObjectConfigFilePath() *AbPath {
        return NewAbPath(this.GetPath(), this.GetObjectConfigFile())
}


// throws Exception
func (this *TransformInfoHttpContext) GetDataFilePath() *AbPath {
        return NewAbPath(this.GetPath(), this.GetDataFile())
}

