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

type BasicUriResolver struct {
    LogUtil *LogUtil
    UriResolverStrings *URIResolverStrings
    Extension string
}

// Implements *URIResolver


func (this *BasicUriResolver) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.UriResolverStrings = URIResolverStrings_GetInstance()
    
}


func NewBasicUriResolver(extension string) *BasicUriResolver {
    this := &BasicUriResolver{}
    this.initFields()
        this.Extension = extension
    return this
}


func (this *BasicUriResolver) GetExtension() string {
        return this.Extension
}


// throws TransformerException
func (this *BasicUriResolver) Resolve(href string, base string) *Source {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Source) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetMainPath())
            stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
            stringBuffer.Append(href)
            var abPath IAbPath = any(NewAbFilePath(fmt.Sprint(stringBuffer))).(IAbPath); _ = abPath
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XMLLOGGING) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.UriResolverStrings.ATTEMPT)
                stringBuffer.Append(href)
                stringBuffer.Append(this.UriResolverStrings.BASE)
                stringBuffer.Append(base)
                stringBuffer.Append(this.UriResolverStrings.NEW_PATH)
                stringBuffer.Append(fmt.Sprint(abPath))
                stringBuffer.Append(this.UriResolverStrings.NOTE)
                stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
                stringBuffer.Append(this.UriResolverStrings.URL_GLOBAL)
                stringBuffer.Append(this.UriResolverStrings.REQUIRED_EXTENSION)
                stringBuffer.Append(this.Extension)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.UriResolverStrings.RESOLVE)
            }
            tryControl1, tryResult1 = 1, NewStreamSource(NewCryptFileReader(TransformInfoTemplateData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoTemplateData_GetInstance().ENCRYPTED_EXTENSION).GetInputStream(abPath)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *BasicUriResolver) ToString() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "BasicUriResolver - Does not work without webapp path should be changed"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetMainPath())
            stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
            stringBuffer.Append(this.UriResolverStrings.IMPORT_URL)
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

func (this *BasicUriResolver) String() string {
    return this.ToString()
}

