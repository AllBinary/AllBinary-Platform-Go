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

type TransformInfo struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbeClientInformation AbeClientInformationInterface
    IsChild bool
    Name string
    StoreName string
    ObjectFileName string
    AnyType any
    ObjectConfigFileName string
    ObjectConfigInterface TransformInfoObjectConfigInterface
    Template string
    TemplateFileName string
    DataFileName string
    Data string
}

// Implements TransformInfoInterface

// The Java class TransformInfo has subclasses so its variables are this interface
type ITransformInfo interface {
    TransformInfo_() *TransformInfo
    Override(hashMap *HashMap)
    Log() string
    SetName(value string)
    SetStoreName(value string)
    GetStoreName() string
    SetObject(anyType any)
    SetObjectFile(value string)
    SetTemplate(value string)
    SetTemplateFile(value string)
    SetObjectConfigFileName(value string)
    SetObjectConfigFile(value string)
    SetObjectConfig(value string)
    SetDataFile(value string)
    SetData(value string)
    GetName() string
    GetObject() any
    GetObjectFile() string
    GetTemplate() string
    GetTemplateFile() string
    GetTemplateFilePath() IAbPath
    GetObjectConfigFilePath() IAbPath
    GetDataFilePath() IAbPath
    GetObjectConfigFile() string
    GetData() string
    GetDataDocument() *Document
    GetDataFile() string
    ToHashMap() *HashMap
    ToVector() IBasicArrayList
    GetKey() any
    SetChild()
    IsChild() bool
    GetObjectConfigInterface() TransformInfoObjectConfigInterface
    SetObjectConfigInterface(objectConfigInterface TransformInfoObjectConfigInterface)
}

func (this *TransformInfo) TransformInfo_() *TransformInfo {
    return this
}


func (this *TransformInfo) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    this.IsChild = false
    
    
    
    
    
    
    
    
    
    this.Data = StringUtil_GetInstance().EMPTY_STRING
}




// Java overloads of new TransformInfo
func NewTransformInfo(args ...any) *TransformInfo {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewTransformInfo_1()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewTransformInfo_2(Java_Arg[string](args[0]))
        }
        if len(args) == 5 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) {
            return NewTransformInfo_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]))
        }
    }
    panic("No Java overload of new TransformInfo for the arguments")
}
func NewTransformInfo_1() *TransformInfo {
    this := &TransformInfo{}
    this.initFields()
    return this
}


func NewTransformInfo_2(name string) *TransformInfo {
    this := &TransformInfo{}
    this.initFields()
        this.SetName(name)
    return this
}


func NewTransformInfo_3(name string, objectFileName string, objectConfigFileName string, templateFileName string, dataFileName string) *TransformInfo {
    this := &TransformInfo{}
    this.initFields()
        this.SetName(name)
        this.SetObjectFile(objectFileName)
        this.SetObjectConfigFile(objectConfigFileName)
        this.SetTemplateFile(templateFileName)
        this.SetDataFile(dataFileName)
    return this
}


// throws Exception
func (this *TransformInfo) Override(hashMap *HashMap) {
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        var aName string = any(hashMap.Get(transformInfoData.NAME)).(string); _ = aName
        if !stringValidationUtil.IsEmpty(aName) {
            this.Name = aName
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("TransformInfo properties overridden for: ")
            stringBuffer.Append(this.GetName())
            stringBuffer.Append("\n properties: ")
            stringBuffer.Append(fmt.Sprint(hashMap))
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "override(HashMap hashMap)")
        }
        var aStoreName string = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = aStoreName
        if !stringValidationUtil.IsEmpty(aStoreName) {
            this.StoreName = aStoreName
        }
        var aObjectFileName string = any(hashMap.Get(transformInfoData.OBJECTFILENAME)).(string); _ = aObjectFileName
        if !stringValidationUtil.IsEmpty(aObjectFileName) {
            this.ObjectFileName = aObjectFileName
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(URLGLOBALS.GetMainPath())
        stringBuffer.Append(FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(this.StoreName)
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        var fileAbPath IAbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = fileAbPath
        var aObject string = any(hashMap.Get(transformInfoData.OBJECT)).(string); _ = aObject
        if !stringValidationUtil.IsEmpty(aObject) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("TransformInfo override view Object for: " + this.GetName(), this, "override(HashMap hashMap)")
            }
            var fileDataString string = NewCryptFileReader(TransformInfoObjectData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoObjectData_GetInstance().ENCRYPTED_EXTENSION).Get(NewAbPath(fmt.Sprint(fileAbPath), aObject)); _ = fileDataString
            if fileDataString != "" {
                this.AnyType = fileDataString
            }
        }
        if !stringValidationUtil.IsEmpty(any(hashMap.Get(transformInfoData.OBJECTCONFIG)).(string)) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("TransformInfo override ObjectConfig for: " + this.GetName(), this, "override(HashMap hashMap)")
            }
            this.SetObjectConfig(any(hashMap.Get(transformInfoData.OBJECTCONFIG)).(string))
        }
        var objectConfigFileName string = any(hashMap.Get(transformInfoData.OBJECTCONFIGFILENAME)).(string); _ = objectConfigFileName
        if !stringValidationUtil.IsEmpty(objectConfigFileName) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("TransformInfo override ObjectConfig File for: " + this.GetName(), this, "override(HashMap hashMap)")
            }
            this.SetObjectConfigFile(objectConfigFileName)
        }
        var type_ string = any(hashMap.Get(OutputTypeData_GetInstance().NAME)).(string); _ = type_
        if !stringValidationUtil.IsEmpty(type_) {
            if this.GetObjectConfigInterface() == nil {
                this.SetObjectConfigInterface(NewTransformInfoObjectConfig(this))
            }
            this.GetObjectConfigInterface().SetOutputTypeName(type_)
        }
        var aTemplateFileName string = any(hashMap.Get(transformInfoData.TEMPLATEFILENAME)).(string); _ = aTemplateFileName
        if !stringValidationUtil.IsEmpty(aTemplateFileName) {
            this.TemplateFileName = aTemplateFileName
        }
        var aTemplate string = any(hashMap.Get(transformInfoData.TEMPLATE)).(string); _ = aTemplate
        if !stringValidationUtil.IsEmpty(aTemplate) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("TransformInfo override Template for: " + this.GetName(), this, "override(HashMap hashMap)")
            }
            var fileDataString string = NewCryptFileReader(TransformInfoTemplateData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoTemplateData_GetInstance().ENCRYPTED_EXTENSION).Get(NewAbPath(fmt.Sprint(fileAbPath), aTemplate)); _ = fileDataString
            if !stringValidationUtil.IsEmpty(fileDataString) {
                this.Template = fileDataString
            }
        }
        var aDataFileName string = any(hashMap.Get(transformInfoData.DATAFILENAME)).(string); _ = aDataFileName
        if !stringValidationUtil.IsEmpty(aDataFileName) {
            this.DataFileName = aDataFileName
        }
        var aData string = any(hashMap.Get(transformInfoData.DATA)).(string); _ = aData
        if !stringValidationUtil.IsEmpty(aData) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("TransformInfo override Data for: " + this.GetName(), this, "override(HashMap hashMap)")
            }
            var fileDataString string = NewCryptFileReader(TransformInfoDataData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoDataData_GetInstance().ENCRYPTED_EXTENSION).Get(NewAbPath(fmt.Sprint(fileAbPath), aData)); _ = fileDataString
            if !stringValidationUtil.IsEmpty(fileDataString) {
                this.Data = fileDataString
            }
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            stringBuffer = NewStringMaker()
            stringBuffer.Append("Name: ")
            stringBuffer.Append(aName)
            stringBuffer.Append("\nTemp Object File: ")
            stringBuffer.Append(aObject)
            stringBuffer.Append("\nTemp Object Config File: ")
            stringBuffer.Append(fmt.Sprint(this.GetObjectConfigFilePath()))
            stringBuffer.Append("\nTemp Template File: ")
            stringBuffer.Append(aTemplate)
            stringBuffer.Append("\nTemp Data File: ")
            stringBuffer.Append(aData)
            stringBuffer.Append("\nStore Name: ")
            stringBuffer.Append(aStoreName)
            stringBuffer.Append("\nObjectFile: ")
            stringBuffer.Append(this.GetObjectFile())
            stringBuffer.Append("->TransformInfo")
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "TransformInfoInterface(HashMap)")
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORY) {
            this.LogUtil.PutF(this.Log(), this, "override")
        }
}


func (this *TransformInfo) Log() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Name: ")
        stringBuffer.Append(this.Name)
        stringBuffer.Append("\nObject File: ")
        stringBuffer.Append(this.ObjectFileName)
        stringBuffer.Append("\nObject Config File Path: ")
        stringBuffer.Append(this.GetObjectConfigFile())
        stringBuffer.Append("\nTemplate File: ")
        stringBuffer.Append(this.TemplateFileName)
        stringBuffer.Append("\nData File: ")
        stringBuffer.Append(this.DataFileName)
        stringBuffer.Append("\nStore Name: ")
        stringBuffer.Append(this.StoreName)
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append(this.GetObjectFile())
        stringBuffer.Append("->TransformInfo")
        stringBuffer.Append("TransformInfoInterface(HashMap)")
        return fmt.Sprint(stringBuffer)
}


func (this *TransformInfo) SetName(value string) {
        this.Name = value
}


func (this *TransformInfo) SetStoreName(value string) {
        this.StoreName = value
}


func (this *TransformInfo) GetStoreName() string {
        return this.StoreName
}


func (this *TransformInfo) SetObject(anyType any) {
        this.AnyType = anyType
}


func (this *TransformInfo) SetObjectFile(value string) {
        this.ObjectFileName = value
}


func (this *TransformInfo) SetTemplate(value string) {
        this.Template = value
}


func (this *TransformInfo) SetTemplateFile(value string) {
        this.TemplateFileName = value
}


// throws Exception
func (this *TransformInfo) SetObjectConfigFileName(value string) {
        this.ObjectConfigFileName = value
}


// throws Exception
func (this *TransformInfo) SetObjectConfigFile(value string) {
        this.ObjectConfigFileName = value
        this.SetObjectConfigInterface(TransformInfoObjectConfigAndManipulatorFactory_GetInstance().GetInstance(this.AbeClientInformation, this, this.GetObjectConfigFilePath()))
}


// throws Exception
func (this *TransformInfo) SetObjectConfig(value string) {
        var document *Document = DomDocumentHelper_Create(value); _ = document
        this.SetObjectConfigInterface(TransformInfoObjectConfigAndManipulatorFactory_GetInstance().GetInstance(this.AbeClientInformation, this, document))
}


func (this *TransformInfo) SetDataFile(value string) {
        this.DataFileName = value
}


func (this *TransformInfo) SetData(value string) {
        this.Data = value
}


func (this *TransformInfo) GetName() string {
        return this.Name
}


func (this *TransformInfo) GetObject() any {
        return this.AnyType
}


func (this *TransformInfo) GetObjectFile() string {
        return this.ObjectFileName
}


func (this *TransformInfo) GetTemplate() string {
        return this.Template
}


func (this *TransformInfo) GetTemplateFile() string {
        return this.TemplateFileName
}


// throws Exception
func (this *TransformInfo) GetTemplateFilePath() IAbPath {
        return NewAbFilePath(this.GetTemplateFile())
}


// throws Exception
func (this *TransformInfo) GetObjectConfigFilePath() IAbPath {
        return NewAbFilePath(this.GetObjectConfigFile())
}


// throws Exception
func (this *TransformInfo) GetDataFilePath() IAbPath {
        return NewAbFilePath(this.GetDataFile())
}


func (this *TransformInfo) GetObjectConfigFile() string {
        return this.ObjectConfigFileName
}


func (this *TransformInfo) GetData() string {
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        var dataFileString string = this.GetDataFile(); _ = dataFileString
        if !stringValidationUtil.IsEmpty(dataFileString) {
            if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                    this.LogUtil.Put("Could Not Load Data from: " + fmt.Sprint(this.GetDataFile()), this, "getData()", e)
                }
                    }
                }()
                var fileData string = NewCryptFileReader(TransformInfoDataData_GetInstance().UNCRYPTED_EXTENSION, TransformInfoDataData_GetInstance().ENCRYPTED_EXTENSION).Get(this.GetDataFilePath()); _ = fileData
                if !stringValidationUtil.IsEmpty(this.Data) {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        this.LogUtil.PutF("Data File overriding existing data: " + fmt.Sprint(this.GetDataFilePath()), this, "getData()")
                    }
                }
                this.Data = fileData
                return
            }(); tryControl1 == 1 {
                return tryResult1
            }
        }
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Could Not Preview Data: " + this.Data, this, "getData()", e)
            }
                }
            }()
            if !stringValidationUtil.IsEmpty(this.Data) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                    this.LogUtil.PutF("Data: " + XmlDocumentHelper_ToString(DomDocumentHelper_Create(this.Data)), this, "getData()")
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("Name: ")
                    stringBuffer.Append(this.Name)
                    stringBuffer.Append("\nEmpty Data For: ")
                    stringBuffer.Append(dataFileString)
                    stringBuffer.Append("\nPath: ")
                    stringBuffer.Append(this.GetDataFilePath().ToFileSystemString())
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "getData()")
                    }
                }
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return this.Data
}


// throws Exception
func (this *TransformInfo) GetDataDocument() *Document {
        var localData string = this.GetData(); _ = localData
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if stringValidationUtil.IsEmpty(localData) {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.PutF("No Data So Creating New Document", this, "getDataDocument()")
            }
            return DomDocumentHelper_Create()
        } else {
            var document *Document = DomDocumentHelper_Create(localData); _ = document
            return document
        }
}


func (this *TransformInfo) GetDataFile() string {
        return this.DataFileName
}


func (this *TransformInfo) ToHashMap() *HashMap {
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(transformInfoData.NAME, this.Name)
        hashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
        hashMap.Put(transformInfoData.OBJECTFILENAME, this.ObjectFileName)
        hashMap.Put(transformInfoData.OBJECT, this.AnyType)
        hashMap.Put(transformInfoData.OBJECTCONFIGFILENAME, this.GetObjectConfigFile())
        hashMap.Put(transformInfoData.OBJECTCONFIG, Encoder_Encode(Java_GetBytes(fmt.Sprint(this.GetObjectConfigInterface()))))
        hashMap.Put(transformInfoData.TEMPLATEFILENAME, this.TemplateFileName)
        hashMap.Put(transformInfoData.TEMPLATE, this.Template)
        hashMap.Put(transformInfoData.DATAFILENAME, this.DataFileName)
        hashMap.Put(transformInfoData.DATA, Encoder_Encode(Java_GetBytes(this.Data)))
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        hashMap.Put(EntryData_GetInstance().LASTMODIFIED, time)
        return hashMap
}


func (this *TransformInfo) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.Name)
        vector.Add(this.StoreName)
        vector.Add(this.ObjectFileName)
        vector.Add(this.AnyType)
        vector.Add(this.GetObjectConfigFile())
        vector.Add(Encoder_Encode(Java_GetBytes(fmt.Sprint(this.GetObjectConfigInterface()))))
        vector.Add(this.TemplateFileName)
        vector.Add(this.Template)
        vector.Add(this.DataFileName)
        vector.Add(Encoder_Encode(Java_GetBytes(this.Data)))
        var calendar *Calendar = Calendar_GetInstance(); _ = calendar
        var time string = fmt.Sprint(int64(calendar.GetTimeInMillis())); _ = time
        vector.Add(time)
        vector.Add(time)
        return vector
}


func (this *TransformInfo) GetKey() any {
        return any(this.GetName()).(any)
}


func (this *TransformInfo) SetChild() {
        this.IsChild = true
}


func (this *TransformInfo) IsChild() bool {
        return this.IsChild
}


func (this *TransformInfo) GetObjectConfigInterface() TransformInfoObjectConfigInterface {
        return this.ObjectConfigInterface
}


func (this *TransformInfo) SetObjectConfigInterface(objectConfigInterface TransformInfoObjectConfigInterface) {
        this.ObjectConfigInterface = objectConfigInterface
}

