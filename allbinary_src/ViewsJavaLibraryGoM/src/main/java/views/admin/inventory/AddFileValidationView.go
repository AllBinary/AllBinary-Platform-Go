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

type AddFileValidationView struct {
    InventoryItemView
    LogUtil *LogUtil
    BasicItemData *BasicItemData
}

// Implements ValidationComponentInterface

var AddFileValidationView_ADDPRODUCT string = "Add Product"
var AddFileValidationView_NEXTSTEP string = fmt.Sprint(CommonLabels_GetInstance().NEXT) + " Step"

func (this *AddFileValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
}


func NewAddFileValidationView(transformInfoInterface TransformInfoInterface) *AddFileValidationView {
    this := &AddFileValidationView{}
        this.InventoryItemView = *NewInventoryItemView(transformInfoInterface)
        this.initFields()
    return this
}


func (this *AddFileValidationView) IsValid_1() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Exception in validation", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var command string = any(this.GetRequestHashMap().Get(InventoryItemView_GLOBALS2.ADMINCOMMAND)).(string); _ = command
            if StringValidationUtil_GetInstance().IsEmpty(command) || (strings.Compare(command, AddFileValidationView_ADDPRODUCT) != 0 && strings.Compare(command, AddFileValidationView_NEXTSTEP) != 0) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("Invalid AdminCommand: ")
                    stringBuffer.Append(command)
                    stringBuffer.Append(" must be: ")
                    stringBuffer.Append(AddFileValidationView_ADDPRODUCT)
                    stringBuffer.Append(" or ")
                    stringBuffer.Append(AddFileValidationView_NEXTSTEP)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var basicItemValidation *BasicItemValidation = NewBasicItemValidation(this.ItemInterface); _ = basicItemValidation
            if basicItemValidation.IsValid() == false {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF(basicItemValidation.ValidationInfo(), this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
            this.ItemInterface.SetCategory(storeFrontInterface.GetCategoryPath() + this.ItemInterface.GetCategory())
            var fullCategory string = fmt.Sprint(NewStringMaker().Append(InventoryItemView_URLGLOBALS.GetWebappPath()).Append(storeFrontInterface.GetCurrentHostNamePath()).Append(this.ItemInterface.GetCategory())); _ = fullCategory
            var categoryFile *AbFile = AbFile_CreateAbFile(fullCategory); _ = categoryFile
            if !categoryFile.IsDirectory() {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Category Does Not Exist: " + fullCategory, this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            if InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.ItemInterface.GetId()) != nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Item Already Exists", this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var imageFileItemObject any = this.GetRequestHashMap().Get(this.BasicItemData.IMAGE); _ = imageFileItemObject
            if HttpFileUploadUtil_GetInstance().IsValid(imageFileItemObject) {
                var fileItem *FileItem = any(imageFileItemObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                HttpFileUploadUtil_Log(fileItem)
                if this.IsValid(fileName, size) == false {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of AddFileValidationView.isValid
func (this *AddFileValidationView) IsValid(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.IsValid_1()
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[int64](args[1], pass) {
            return this.IsValid_2(Java_Arg[string](args[0]), Java_Arg[int64](args[1]))
        }
    }
    panic("No Java overload of AddFileValidationView.isValid for the arguments")
}


func (this *AddFileValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *AddFileValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *AddFileValidationView) ValidationInfo_1() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Getting Validation Info"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var command string = any(this.GetRequestHashMap().Get(InventoryItemView_GLOBALS2.ADMINCOMMAND)).(string); _ = command
            if StringValidationUtil_GetInstance().IsEmpty(command) || (strings.Compare(command, AddFileValidationView_ADDPRODUCT) != 0 && strings.Compare(command, AddFileValidationView_NEXTSTEP) != 0) {
                tryControl1, tryResult1 = 1, StringUtil_GetInstance().EMPTY_STRING; return
            }
            var basicItemValidation *BasicItemValidation = NewBasicItemValidation(this.ItemInterface); _ = basicItemValidation
            if basicItemValidation.IsValid() == false {
                stringBuffer.Append(basicItemValidation.ValidationInfo())
                stringBuffer.Append("<br/>")
            }
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
            var fullCategory string = fmt.Sprint(NewStringMaker().Append(InventoryItemView_URLGLOBALS.GetWebappPath()).Append(storeFrontInterface.GetCurrentHostNamePath()).Append(this.ItemInterface.GetCategory())); _ = fullCategory
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Category: " + fullCategory, this, "validationInfo()")
            }
            var categoryFile *AbFile = AbFile_CreateAbFile(fullCategory); _ = categoryFile
            if !categoryFile.IsDirectory() {
                stringBuffer.Append("Category ")
                stringBuffer.Append(this.ItemInterface.GetCategory())
                stringBuffer.Append(" does not exist.<br />")
            }
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Existing Item With MoneyException", this, "validationInfo()")
                }
                    }
                }()
                if InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.ItemInterface.GetId()) != nil {
                    stringBuffer.Append("Id is already in use. Refresh the page for a valid id.<br/>")
                }
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
            var anyType any = this.GetRequestHashMap().Get(this.BasicItemData.IMAGE); _ = anyType
            if HttpFileUploadUtil_GetInstance().IsValid(anyType) {
                var fileItem *FileItem = any(anyType).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var fileItemFieldName string = fileItem.GetFieldName(); _ = fileItemFieldName
                var size int64 = fileItem.GetSize(); _ = size
                this.ValidationInfo(stringBuffer, fileName, fileItemFieldName, size)
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of AddFileValidationView.validationInfo
func (this *AddFileValidationView) ValidationInfo(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.ValidationInfo_1()
        }
        if len(args) == 4 && Java_IsArg[*StringMaker](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[int64](args[3], pass) {
            return this.ValidationInfo_2(Java_Arg[*StringMaker](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[int64](args[3]))
        }
    }
    panic("No Java overload of AddFileValidationView.validationInfo for the arguments")
}


// throws Exception
func (this *AddFileValidationView) IsValid_2(fileName string, size int64) bool {
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        if size > int64(fileData.MINIMAGEFILESIZE) {
            if size < int64(fileData.MAXIMAGEFILESIZE) {
                if !StringValidationUtil_GetInstance().IsValidRequired(fileName, fileData.MINLEN, fileData.MAXLEN) {
                    return false
                } else {
                    var pathData *AbPathData = AbPathData_GetInstance(); _ = pathData
                    var extension string = pathData.GetExtension(fileName); _ = extension
                    var uploadMedia *UploadMediaSingleton = UploadMediaSingleton_GetInstance(); _ = uploadMedia
                    if !uploadMedia.IsWriterSupported(extension) && !uploadMedia.IsReaderSupported(extension) {
                        return false
                    }
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("File Size To Large: ")
                    stringBuffer.Appendlong(size)
                    stringBuffer.Append(">")
                    stringBuffer.Appendint(fileData.MAXIMAGEFILESIZE)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                }
                return false
            }
            this.ProcessImageFiles()
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("File Size To Small: ")
                stringBuffer.Appendlong(size)
                stringBuffer.Append(">")
                stringBuffer.Appendint(fileData.MINIMAGEFILESIZE)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
            }
            return false
        }
        return true
}


func (this *AddFileValidationView) ValidationInfo_2(stringBuffer *StringMaker, fileName string, fileItemFieldName string, size int64) {
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        if size > int64(fileData.MINIMAGEFILESIZE) {
            if strings.Compare(fileItemFieldName, this.BasicItemData.IMAGE) == 0 {
                if size < int64(fileData.MAXIMAGEFILESIZE) {
                    if !StringValidationUtil_GetInstance().IsValidRequired(fileName, fileData.MINLEN, fileData.MAXLEN) {
                        stringBuffer.Append("FileName must be >")
                        stringBuffer.Appendint(fileData.MINLEN)
                        stringBuffer.Append(" and <")
                        stringBuffer.Appendint(fileData.MAXLEN)
                        stringBuffer.Append("<br/>")
                    } else {
                        var pathData *AbPathData = AbPathData_GetInstance(); _ = pathData
                        var extension string = pathData.GetExtension(fileName); _ = extension
                        var uploadMedia *UploadMediaSingleton = UploadMediaSingleton_GetInstance(); _ = uploadMedia
                        if !uploadMedia.IsWriterSupported(extension) && !uploadMedia.IsReaderSupported(extension) {
                            stringBuffer.Append("Image type: ")
                            stringBuffer.Append(extension)
                            stringBuffer.Append(" not supported.<br />")
                        }
                    }
                } else {
                    stringBuffer.Append("Image File Is To Large. ")
                    stringBuffer.Appendint(fileData.MINIMAGEFILESIZE)
                    stringBuffer.Append(" < > ")
                    stringBuffer.Appendint(fileData.MAXIMAGEFILESIZE)
                    stringBuffer.Append("<br/>")
                }
            }
        } else {
            stringBuffer.Append("Image File Is To Small. ")
            stringBuffer.Appendint(fileData.MINIMAGEFILESIZE)
            stringBuffer.Append(" < > ")
            stringBuffer.Appendint(fileData.MAXIMAGEFILESIZE)
            stringBuffer.Append("<br/>")
        }
}

