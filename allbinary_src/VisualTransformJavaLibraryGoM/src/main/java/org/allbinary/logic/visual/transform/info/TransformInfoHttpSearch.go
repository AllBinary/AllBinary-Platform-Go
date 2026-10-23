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

type TransformInfoHttpSearch struct {
    TransformInfoHttp
}



func NewTransformInfoHttpSearch(searchRequest *SearchRequest) *TransformInfoHttpSearch {
    this := &TransformInfoHttpSearch{}
        this.TransformInfoHttp = *NewTransformInfoHttp(searchRequest.GetPropertiesHashMap(), searchRequest.GetPageContext())
        if searchRequest.GetStoreFront().GetName() != "" {
            this.SetStoreName(searchRequest.GetStoreFront().GetName())
        } else {
            this.SetStoreName(StringUtil_GetInstance().EMPTY_STRING)
        }
        var dataFileName string = searchRequest.GetXslFile(); _ = dataFileName
        this.SetTemplateFile(dataFileName)
    return this
}


// throws Exception
func (this *TransformInfoHttpSearch) GetPath() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(TransformInfoHttp_URLGLOBALS.GetMainPath())
        stringBuffer.Append(TransformInfoHttp_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(this.GetStoreName())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        return fmt.Sprint(stringBuffer)
}


// throws Exception
func (this *TransformInfoHttpSearch) GetTemplateFilePath() IAbPath {
        return NewAbPath(this.GetPath(), this.GetTemplateFile())
}


// throws Exception
func (this *TransformInfoHttpSearch) GetObjectConfigFilePath() IAbPath {
        return NewAbPath(this.GetPath(), this.GetObjectConfigFile())
}


// throws Exception
func (this *TransformInfoHttpSearch) GetDataFilePath() IAbPath {
        return NewAbPath(this.GetPath(), this.GetDataFile())
}

