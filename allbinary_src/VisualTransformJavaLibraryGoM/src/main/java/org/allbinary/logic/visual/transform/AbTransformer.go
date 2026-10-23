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

type AbTransformer struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TransformInfoInterface TransformInfoInterface
    InputStream IInputStream
    UriResolver *URIResolver
}

// Implements BasicTransformerInterface

// The Java class AbTransformer has subclasses so its variables are this interface
type IAbTransformer interface {
    AbTransformer_() *AbTransformer
    GetTransformInfoInterface() TransformInfoInterface
    SetInputStream(inputStream IInputStream)
    SetURIResolver(uriResolver *URIResolver)
    GetInputStream() IInputStream
    GetURIResolver() *URIResolver
    GetStreamSource() *StreamSource
    Translate(xmlDocumentStr string) string
}

func (this *AbTransformer) AbTransformer_() *AbTransformer {
    return this
}


func (this *AbTransformer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewAbTransformer(transformInfoInterface TransformInfoInterface) *AbTransformer {
    this := &AbTransformer{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
    return this
}


func (this *AbTransformer) GetTransformInfoInterface() TransformInfoInterface {
        return this.TransformInfoInterface
}


func (this *AbTransformer) SetInputStream(inputStream IInputStream) {
        this.InputStream = inputStream
}


func (this *AbTransformer) SetURIResolver(uriResolver *URIResolver) {
        this.UriResolver = uriResolver
}


func (this *AbTransformer) GetInputStream() IInputStream {
        return this.InputStream
}


func (this *AbTransformer) GetURIResolver() *URIResolver {
        return this.UriResolver
}


func (this *AbTransformer) GetStreamSource() *StreamSource {
        return NewStreamSource(this.GetInputStream())
}


// throws Exception
func (this *AbTransformer) Translate(xmlDocumentStr string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.TransformInfoInterface.Log(), this, "translate(document)", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XMLLOGGING) {
                this.LogUtil.PutF("xml: \n" + xmlDocumentStr, this, "translate(String xmlDocumentStr)")
            }
            if this.GetURIResolver() == nil {
                panic(errors.New(fmt.Sprint("No URIResolver")))
            }
            var result string = XslHelper_GetInstance().Translate(this.GetURIResolver(), NewStreamSource(this.GetInputStream()), NewStreamSource(NewStringBufferInputStream(xmlDocumentStr))).(IOutputStream); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().XSLLOGGING) {
                this.LogUtil.PutF("translated xml: " + result, this, "translate(String xmlDocumentStr)")
            }
            var transformInfoObjectConfigGeneratorInterface ITransformInfoObjectConfigGenerator = TransformInfoObjectConfigGeneratorFactory_GetInstance().GetInstance(this.GetTransformInfoInterface()); _ = transformInfoObjectConfigGeneratorInterface
            tryControl1, tryResult1 = 1, transformInfoObjectConfigGeneratorInterface.Process(result); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

