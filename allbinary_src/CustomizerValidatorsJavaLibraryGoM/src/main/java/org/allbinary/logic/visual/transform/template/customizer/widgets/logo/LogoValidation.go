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

type LogoValidation struct {
    Validation
    LogUtil *LogUtil
    LogoFileItem *FileItem
    LogoAbPath IAbPath
    LogoFile string
    IsFile bool
}

// Implements DomNodeInterface


func (this *LogoValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
}




// Java overloads of new LogoValidation
func NewLogoValidation(args ...any) *LogoValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewLogoValidation_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewLogoValidation_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewLogoValidation_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new LogoValidation for the arguments")
}
func NewLogoValidation_1() *LogoValidation {
    this := &LogoValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.IsFile = false
        this.LogoAbPath = nil
        this.LogoFile = nil
    return this
}


func NewLogoValidation_2(node *Node) *LogoValidation {
    this := &LogoValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.IsFile = false
        this.LogoAbPath = NewAbPath(DomSearchHelper_GetNode(LogoData_GetInstance().IMAGEPATH, node.GetChildNodes()).GetNodeValue(), StringUtil_GetInstance().EMPTY_STRING)
        this.LogoFile = DomSearchHelper_GetNode(LogoData_GetInstance().IMAGEFILENAME, node.GetChildNodes()).GetNodeValue()
    return this
}


func NewLogoValidation_3(hashMap *HashMap) *LogoValidation {
    this := &LogoValidation{}
    this.Validation = *NewValidation()
    this.initFields()
        this.IsFile = true
        this.GetFormData(hashMap)
    return this
}


// throws Exception
func (this *LogoValidation) GetFormData(hashMap *HashMap) {
        this.LogoFileItem = any(hashMap.Get(LogoData_GetInstance().IMAGE)).(*FileItem)
        this.LogoAbPath = NewAbPath(any(hashMap.Get(LogoData_GetInstance().IMAGEPATH)).(string), StringUtil_GetInstance().EMPTY_STRING)
        if this.LogoFileItem != nil {
            this.LogoFile = HttpRequestUtil_GetInstance().GenerateFileName(this.LogoFileItem.GetName())
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("HashMap: ")
            stringBuffer.Append(fmt.Sprint(hashMap))
            stringBuffer.Append("\nLogoFileName: ")
            stringBuffer.Append(this.LogoFile)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "getFormData()")
        }
}


func (this *LogoValidation) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            var fileData *FileData = FileData_GetInstance(); _ = fileData
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("LogoValidation", this, this.CommonStrings.IS_VALID)
            }
            if !StringValidationUtil_GetInstance().IsValidRequired(fmt.Sprint(this.LogoAbPath), 0, 512) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("logoPath is invalid: " + fmt.Sprint(this.LogoAbPath), this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            if !StringValidationUtil_GetInstance().IsValidRequired(this.LogoFile, fileData.MINLEN, fileData.MAXLEN) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("logoFile is invalid: " + this.LogoFile, this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            if this.IsFile {
                if this.LogoFileItem == nil {
                    tryControl1, tryResult1 = 1, false; return
                }
                if this.LogoFileItem.GetSize() > fileData.MAXIMAGEFILESIZE || this.LogoFileItem.GetSize() < fileData.MINIMAGEFILESIZE {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        this.LogUtil.PutF("Image File Is Not The Right Size. " + fmt.Sprint(fileData.MINIMAGEFILESIZE) + "< > " + fmt.Sprint(fileData.MAXIMAGEFILESIZE), this, this.CommonStrings.IS_VALID)
                    }
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Logo Validation: " + fmt.Sprint(valid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *LogoValidation) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Logo is not valid.")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *LogoValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *LogoValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *LogoValidation) ToHashMap() *HashMap {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(LogoData_GetInstance().IMAGE, stringUtil.EMPTY_STRING)
        if this.LogoFile != "" {
            hashMap.Put(LogoData_GetInstance().IMAGEPATH, fmt.Sprint(this.LogoAbPath))
            hashMap.Put(LogoData_GetInstance().IMAGEFILENAME, this.LogoFile)
        } else {
            hashMap.Put(LogoData_GetInstance().IMAGEPATH, nil)
            hashMap.Put(LogoData_GetInstance().IMAGEFILENAME, this.LogoFile)
        }
        return hashMap
}


// throws Exception
func (this *LogoValidation) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        return ModDomHelper_CreateNameValueNodes(document, LogoData_GetInstance().NAME, hashMap)
}


// throws Exception
func (this *LogoValidation) ProcessLogoFile() {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Processing logoFile: ")
            stringBuffer.Append(fmt.Sprint(this.LogoAbPath))
            stringBuffer.Append(this.LogoFile)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "processLogoFile()")
        }
        NewLogoImageFileUtil(this.LogoAbPath, this.LogoFile).SaveFiles(this.LogoFileItem)
}

