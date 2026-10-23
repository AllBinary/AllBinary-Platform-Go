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

type BasicTransformer struct {
    AbTransformer
    LogUtil *LogUtil
    StreamUtil *StreamUtil
    AbeClientInformation AbeClientInformationInterface
}

// The Java class BasicTransformer has subclasses so its variables are this interface
type IBasicTransformer interface {
    IAbTransformer
    BasicTransformer_() *BasicTransformer
    SetEncryptedTemplateFileAsInputStream(file *AbFile)
    SetEncryptedTemplateAsInputStream()
    SetTemplateAsInputStream()
}

func (this *BasicTransformer) BasicTransformer_() *BasicTransformer {
    return this
}


func (this *BasicTransformer) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StreamUtil = StreamUtil_GetInstance()
    
}


func NewBasicTransformer(abeClientInformation AbeClientInformationInterface, transformInfoInterface TransformInfoInterface) *BasicTransformer {
    this := &BasicTransformer{}
        this.AbTransformer = *NewAbTransformer(transformInfoInterface)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.SetTemplateAsInputStream()
    return this
}


// throws Exception
func (this *BasicTransformer) SetEncryptedTemplateFileAsInputStream(file *AbFile) {
        var outputStream *strings.Builder = nil; _ = outputStream
        var inputStream IAbFileInputStream = nil; _ = inputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Failed to get Encrypted File: ")
                stringBuffer.Append(file.GetPath())
                stringBuffer.Append(" \nName:")
                stringBuffer.Append(file.GetName())
                this.LogUtil.Put(fmt.Sprint(stringBuffer), this, "setEncryptedTemplateFileAsInputStream()", e)
            }
            panic(e)
                }
            this.StreamUtil.Close(outputStream)
            this.StreamUtil.Close(inputStream)
            }()
            outputStream = &strings.Builder{}
            inputStream = NewAbFileInputStream(file)
            outputStream = any(this.StreamUtil.Get(inputStream, outputStream, make([]int8, 16384))).(*strings.Builder)
            var abCrypt *AbCrypt = NewAbCrypt(KeySpecFactory_DESEDE, AbKeys_GetInstance().GetKey(this.AbeClientInformation, file.GetAbsolutePath())); _ = abCrypt
            var decrypted []int8 = abCrypt.Decrypt(outputStream.ToByteArray()); _ = decrypted
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Decrypted Template: \n" + fmt.Sprint(string(decrypted)), this, "setEncryptedTemplateFileAsInputStream(file)")
            }
            this.SetInputStream(any(NewByteArrayInputStream(decrypted)).(IInputStream))
            this.SetURIResolver(any(NewBasicUriResolver(TransformInfoTemplateData_GetInstance().ENCRYPTED_EXTENSION)).(*URIResolver))
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *BasicTransformer) SetEncryptedTemplateAsInputStream() {
        var outputStream *strings.Builder = nil; _ = outputStream
        var inputStream *ByteArrayInputStream = nil; _ = inputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put("Failed to set with template: " + fmt.Sprint(this.GetTransformInfoInterface().GetTemplateFilePath()), this, "setEncryptedTemplateFileAsInputStream()", e)
            }
            panic(e)
                }
            this.StreamUtil.Close(outputStream)
            this.StreamUtil.Close(inputStream)
            }()
            outputStream = &strings.Builder{}
            inputStream = NewByteArrayInputStream(Java_GetBytes(this.GetTransformInfoInterface().GetTemplate()))
            outputStream = any(this.StreamUtil.Get(inputStream, outputStream, make([]int8, 16384))).(*strings.Builder)
            var abCrypt *AbCrypt = NewAbCrypt(KeySpecFactory_DESEDE, AbKeys_GetInstance().GetKey(this.AbeClientInformation, this.GetTransformInfoInterface().GetName())); _ = abCrypt
            var decrypted []int8 = abCrypt.Decrypt(outputStream.ToByteArray()); _ = decrypted
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF("Decrypted Template: \n" + fmt.Sprint(decrypted), this, "setEncryptedTemplateAsInputStream()")
            }
            this.SetInputStream(any(NewByteArrayInputStream(decrypted)).(IInputStream))
            this.SetURIResolver(any(NewBasicUriResolver(TransformInfoTemplateData_GetInstance().ENCRYPTED_EXTENSION)).(*URIResolver))
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *BasicTransformer) SetTemplateAsInputStream() {
        var inputStream IAbFileInputStream = nil; _ = inputStream
        var outputStream *strings.Builder = nil; _ = outputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put("Failed to get Template Data", this, "setTemplateAsInputStream()", e)
            }
            panic(e)
                }
            this.StreamUtil.Close(outputStream)
            this.StreamUtil.Close(inputStream)
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                this.LogUtil.PutF(this.GetTransformInfoInterface().Log(), this, "setTemplateAsInputStream()")
            }
            var transformInfoTemplateData *TransformInfoTemplateData = TransformInfoTemplateData_GetInstance(); _ = transformInfoTemplateData
            var pathUtil *PathUtil = PathUtil_GetInstance(); _ = pathUtil
            if this.GetTransformInfoInterface().GetTemplateFilePath() != nil {
                var extension string = pathUtil.GetExtensionWithAbPath(this.GetTransformInfoInterface().GetTemplateFilePath()); _ = extension
                var filePath string = pathUtil.GetWithoutExtensionWithAbPath(this.GetTransformInfoInterface().GetTemplateFilePath()); _ = filePath
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                    this.LogUtil.PutF(this.GetTransformInfoInterface().Log(), this, "setTemplateAsInputStream()")
                }
                if strings.Compare(extension, transformInfoTemplateData.UNCRYPTED_EXTENSION) == 0 {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(filePath)
                    stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
                    stringBuffer.Append(transformInfoTemplateData.ENCRYPTED_EXTENSION)
                    var encFile *AbFile = AbFile_CreateAbFile(fmt.Sprint(stringBuffer)); _ = encFile
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                        this.LogUtil.PutF(this.GetTransformInfoInterface().Log(), this, "setTemplateAsInputStream()")
                    }
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                        stringBuffer = NewStringMaker()
                        stringBuffer.Append("Encrypted Template File isFile=")
                        stringBuffer.Appendboolean(encFile.IsFile())
                        stringBuffer.Append("\nEncTemplateFilePath: ")
                        stringBuffer.Append(encFile.GetPath())
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "setTemplateAsInputStream()")
                    }
                    if encFile.IsFile() {
                        this.SetEncryptedTemplateFileAsInputStream(encFile)
                        tryControl1 = 1; return
                    } else {
                        var file *AbFile = AbFile_CreateAbFileFromAbPath(this.GetTransformInfoInterface().GetTemplateFilePath()); _ = file
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                            this.LogUtil.PutF(this.GetTransformInfoInterface().Log(), this, "setTemplateAsInputStream()")
                        }
                        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                            this.LogUtil.PutF("Template File isFile=" + fmt.Sprint(file.IsFile()), this, "setTemplateAsInputStream()")
                        }
                        if file.IsFile() {
                            inputStream = NewAbFileInputStream(file)
                            outputStream = &strings.Builder{}
                            outputStream = any(this.StreamUtil.Get(inputStream, outputStream, make([]int8, 16384))).(*strings.Builder)
                            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
                                this.LogUtil.PutF("Template: " + outputStream.String(), this, "getTemplateAsInputStream()")
                            }
                            this.SetInputStream(any(NewByteArrayInputStream(outputStream.ToByteArray())).(IInputStream))
                            this.SetURIResolver(any(NewBasicUriResolver(transformInfoTemplateData.UNCRYPTED_EXTENSION)).(*URIResolver))
                            tryControl1 = 1; return
                        }
                    }
                } else if strings.Compare(extension, transformInfoTemplateData.ENCRYPTED_EXTENSION) == 0 {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append(filePath)
                    stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
                    stringBuffer.Append(transformInfoTemplateData.ENCRYPTED_EXTENSION)
                    var file *AbFile = AbFile_CreateAbFile(fmt.Sprint(stringBuffer)); _ = file
                    if file.IsFile() {
                        this.SetEncryptedTemplateFileAsInputStream(file)
                    }
                } else {
                    panic(errors.New(fmt.Sprint("View Template File Type Is Not Recognized: " + fmt.Sprint(this.GetTransformInfoInterface().GetTemplateFilePath()))))
                }
            } else {
                this.SetEncryptedTemplateAsInputStream()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
        panic(errors.New(fmt.Sprint("Error setTemplateAsInputStream()")))
}

