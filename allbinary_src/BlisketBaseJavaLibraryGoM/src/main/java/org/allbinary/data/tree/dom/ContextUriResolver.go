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

type ContextUriResolver struct {
    LogUtil *LogUtil
    UriResolverStrings *URIResolverStrings
    BasicURIResolver *BasicUriResolver
}

// Implements *URIResolver


func (this *ContextUriResolver) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.UriResolverStrings = URIResolverStrings_GetInstance()
    
}


func NewContextUriResolver(basicURIResolver *BasicUriResolver) *ContextUriResolver {
    this := &ContextUriResolver{}
    this.initFields()
        this.BasicURIResolver = basicURIResolver
    return this
}


// throws TransformerException
func (this *ContextUriResolver) Resolve(href string, base string) *Source {
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
            stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().INSTALLPATH)
            stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
            stringBuffer.Append(href)
            var fileAbPath IAbPath = NewAbFilePath(fmt.Sprint(stringBuffer)); _ = fileAbPath
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XMLLOGGING) {
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(this.UriResolverStrings.ATTEMPT)
                stringBuffer.Append(href)
                stringBuffer.Append(this.UriResolverStrings.BASE)
                stringBuffer.Append(base)
                stringBuffer.Append(this.UriResolverStrings.NEW_PATH)
                stringBuffer.Append(fmt.Sprint(fileAbPath))
                stringBuffer.Append(this.UriResolverStrings.NOTE)
                stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
                stringBuffer.Append(this.UriResolverStrings.URL_GLOBAL)
                stringBuffer.Append(this.UriResolverStrings.REQUIRED_EXTENSION)
                stringBuffer.Append(this.BasicURIResolver.GetExtension())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.UriResolverStrings.RESOLVE)
            }
            tryControl1, tryResult1 = 1, NewStreamSource(NewCryptFileReader(TransformInfoTemplateData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoTemplateData_GetInstance().ENCRYPTED_EXTENSION).GetInputStream(fileAbPath)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ContextUriResolver) ToString() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "ContextUriResolver - Does not work without webapp path should be changed"; return
                }
            }()
            tryControl1, tryResult1 = 1, URLGLOBALS.GetMainPath() + FREEBLISKET_PATH_GLOBALS.GetInstance().INSTALLPATH + this.UriResolverStrings.IMPORT_URL; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

func (this *ContextUriResolver) String() string {
    return this.ToString()
}

