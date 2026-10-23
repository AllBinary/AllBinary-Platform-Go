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

type UpdateFileValidationView struct {
    DownloadableInventoryItemView
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}

// Implements ValidationComponentInterface


func (this *UpdateFileValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewUpdateFileValidationView(transformInfoInterface TransformInfoInterface) *UpdateFileValidationView {
    this := &UpdateFileValidationView{}
        this.DownloadableInventoryItemView = *NewDownloadableInventoryItemView(transformInfoInterface)
        this.initFields()
    return this
}


func (this *UpdateFileValidationView) IsValid_1() bool {
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
            var command string = any(this.GetRequestHashMap().Get(DownloadableInventoryItemView_GLOBALS2.ADMINCOMMAND)).(string); _ = command
            if StringValidationUtil_GetInstance().IsEmpty(command) || strings.Compare(command, this.CommonStrings.UPDATE) != 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Invalid AdminCommand=" + command, this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 bool) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.Put("Existing Item With MoneyException", this, this.CommonStrings.IS_VALID, e)
                }
                    }
                }()
                var inventoryEntity *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntity
                this.ItemInterface = inventoryEntity.GetItem(this.Id)
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
            if this.ItemInterface == nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Item Does Not Exist", this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            if NewBasicItemValidation(this.ItemInterface).IsValid() == false {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("BasicItem is not valid", this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
            var fullCategory string = fmt.Sprint(NewStringMaker().Append(DownloadableInventoryItemView_URLGLOBALS.GetWebappPath()).Append(storeFrontInterface.GetCurrentHostNamePath()).Append(this.ItemInterface.GetCategory())); _ = fullCategory
            var categoryFile *AbFile = AbFile_CreateAbFile(fullCategory); _ = categoryFile
            if !categoryFile.IsDirectory() {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Category Does Not Exist: " + fullCategory, this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var downloadableFileObject any = this.GetRequestHashMap().Get(DownloadItemData_FILE); _ = downloadableFileObject
            if HttpFileUploadUtil_GetInstance().IsValid(downloadableFileObject) {
                var fileItem *FileItem = any(downloadableFileObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                HttpFileUploadUtil_Log(fileItem)
                if this.IsValid(fileName, size) == false {
                    tryControl1, tryResult1 = 1, false; return
                } else {
                    this.ProcessFile(fileItem)
                    this.GetRequestHashMap().Put(DownloadItemData_SIZE, fmt.Sprint(size))
                }
            } else {
                this.GetRequestHashMap().Remove(DownloadItemData_FILE)
            }
            var licenseFileObject any = this.GetRequestHashMap().Get(DownloadItemData_LICENSE_FILE); _ = licenseFileObject
            if HttpFileUploadUtil_GetInstance().IsValid(licenseFileObject) {
                var fileItem *FileItem = any(licenseFileObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                HttpFileUploadUtil_Log(fileItem)
                if this.IsValid(fileName, size) == false {
                } else {
                    this.ProcessFile(fileItem)
                }
            } else {
                this.GetRequestHashMap().Remove(DownloadItemData_LICENSE_FILE)
            }
            var extraFilesObject any = this.GetRequestHashMap().Get(DownloadItemData_EXTRA_FILES); _ = extraFilesObject
            if HttpFileUploadUtil_GetInstance().IsValid(extraFilesObject) {
                var fileItem *FileItem = any(extraFilesObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                HttpFileUploadUtil_Log(fileItem)
                if this.IsValid(fileName, size) == false {
                } else {
                    this.ProcessFile(fileItem)
                    this.Unzip(fileItem)
                }
            } else {
                this.GetRequestHashMap().Remove(DownloadItemData_EXTRA_FILES)
            }
            var hashMap *HashMap = this.GetRequestHashMap(); _ = hashMap
            hashMap.Put(this.BasicItemData.ID, this.ItemInterface.GetId())
            var licenseObject any = hashMap.Get(DownloadItemData_LICENSE_FILE); _ = licenseObject
            if licenseObject != nil {
                hashMap.Put(DownloadItemData_LICENSE_FILE, (any(licenseObject).(*FileItem)).GetName())
            } else {
                hashMap.Put(DownloadItemData_LICENSE_FILE, StringUtil_GetInstance().EMPTY_STRING)
            }
            var fileObject any = hashMap.Get(DownloadItemData_FILE); _ = fileObject
            if fileObject != nil {
                hashMap.Put(DownloadItemData_FILE, (any(fileObject).(*FileItem)).GetName())
            } else {
                hashMap.Put(DownloadItemData_FILE, StringUtil_GetInstance().EMPTY_STRING)
            }
            this.DownloadableItem = NewDownloadableItem(hashMap)
            var downloadableItemValidation *DownloadableItemValidation = NewDownloadableItemValidation(this.DownloadableItem); _ = downloadableItemValidation
            if downloadableItemValidation.IsValid() == false {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("DownloadableItem is not valid", this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of UpdateFileValidationView.isValid
func (this *UpdateFileValidationView) IsValid(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.IsValid_1()
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[int64](args[1], pass) {
            return this.IsValid_2(Java_Arg[string](args[0]), Java_Arg[int64](args[1]))
        }
    }
    panic("No Java overload of UpdateFileValidationView.isValid for the arguments")
}


func (this *UpdateFileValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *UpdateFileValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *UpdateFileValidationView) ValidationInfo_1() string {
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
            var command string = any(this.GetRequestHashMap().Get(DownloadableInventoryItemView_GLOBALS2.ADMINCOMMAND)).(string); _ = command
            if StringValidationUtil_GetInstance().IsEmpty(command) || strings.Compare(command, this.CommonStrings.UPDATE) != 0 {
                tryControl1, tryResult1 = 1, StringUtil_GetInstance().EMPTY_STRING; return
            }
            if this.ItemInterface == nil {
                var ITEM_NOT_FOUND string = "Item Does Not Exist"; _ = ITEM_NOT_FOUND
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF(ITEM_NOT_FOUND, this, "validationInfo()")
                }
                tryControl1, tryResult1 = 1, ITEM_NOT_FOUND; return
            }
            var basicItemValidation *BasicItemValidation = NewBasicItemValidation(this.ItemInterface); _ = basicItemValidation
            if basicItemValidation.IsValid() == false {
                stringBuffer.Append(basicItemValidation.ValidationInfo())
                stringBuffer.Append("<br/>")
            }
            var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
            var fullCategory string = fmt.Sprint(NewStringMaker().Append(DownloadableInventoryItemView_URLGLOBALS.GetWebappPath()).Append(storeFrontInterface.GetCurrentHostNamePath()).Append(this.ItemInterface.GetCategory())); _ = fullCategory
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Category: " + fullCategory, this, "validationInfo()")
            }
            var categoryFile *AbFile = AbFile_CreateAbFile(fullCategory); _ = categoryFile
            if !categoryFile.IsDirectory() {
                stringBuffer.Append("Category ")
                stringBuffer.Append(this.ItemInterface.GetCategory())
                stringBuffer.Append(" does not exist.")
                stringBuffer.Append("<br />")
            }
            var downloadableFileObject any = this.GetRequestHashMap().Get(DownloadItemData_FILE); _ = downloadableFileObject
            if HttpFileUploadUtil_GetInstance().IsValid(downloadableFileObject) {
                var fileItem *FileItem = any(downloadableFileObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                this.ValidationInfo(stringBuffer, fileName, size)
            } else {
                stringBuffer.Append("Download File Missing")
                stringBuffer.Append("<br/>")
            }
            var licenseFileObject any = this.GetRequestHashMap().Get(DownloadItemData_LICENSE_FILE); _ = licenseFileObject
            if HttpFileUploadUtil_GetInstance().IsValid(licenseFileObject) {
                var fileItem *FileItem = any(licenseFileObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                this.ValidationInfo(stringBuffer, fileName, size)
            } else {
            }
            var extraFilesObject any = this.GetRequestHashMap().Get(DownloadItemData_EXTRA_FILES); _ = extraFilesObject
            if HttpFileUploadUtil_GetInstance().IsValid(extraFilesObject) {
                var fileItem *FileItem = any(extraFilesObject).(*FileItem); _ = fileItem
                var fileName string = fileItem.GetName(); _ = fileName
                var size int64 = fileItem.GetSize(); _ = size
                this.ValidationInfo(stringBuffer, fileName, size)
            } else {
            }
            if this.DownloadableItem != nil {
                var downloadableItemValidation *DownloadableItemValidation = NewDownloadableItemValidation(this.DownloadableItem); _ = downloadableItemValidation
                if downloadableItemValidation.IsValid() == false {
                    var log string = downloadableItemValidation.ValidationInfo(); _ = log
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        this.LogUtil.PutF(log, this, this.CommonStrings.IS_VALID)
                    }
                    stringBuffer.Append(log)
                    stringBuffer.Append("<br/>")
                }
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of UpdateFileValidationView.validationInfo
func (this *UpdateFileValidationView) ValidationInfo(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.ValidationInfo_1()
        }
        if len(args) == 3 && Java_IsArg[*StringMaker](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[int64](args[2], pass) {
            return this.ValidationInfo_2(Java_Arg[*StringMaker](args[0]), Java_Arg[string](args[1]), Java_Arg[int64](args[2]))
        }
    }
    panic("No Java overload of UpdateFileValidationView.validationInfo for the arguments")
}


// throws Exception
func (this *UpdateFileValidationView) IsValid_2(fileName string, size int64) bool {
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        if size > int64(fileData.MINDOWNLOADABLEFILESIZE) {
            if size < int64(fileData.MAXDOWNLOADABLEFILESIZE) {
                if !StringValidationUtil_GetInstance().IsValidRequired(fileName, fileData.MINLEN, fileData.MAXLEN) {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("File name is not valid: ")
                        stringBuffer.Append(fileName)
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                    }
                    return false
                } else {
                }
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("File Size To Large: ")
                    stringBuffer.Appendlong(size)
                    stringBuffer.Append(">")
                    stringBuffer.Appendint(fileData.MAXDOWNLOADABLEFILESIZE)
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
                }
                return false
            }
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("File Size To Small: ")
                stringBuffer.Appendlong(size)
                stringBuffer.Append(">")
                stringBuffer.Appendint(fileData.MINDOWNLOADABLEFILESIZE)
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.IS_VALID)
            }
            return false
        }
        return true
}


func (this *UpdateFileValidationView) ValidationInfo_2(stringBuffer *StringMaker, fileName string, size int64) {
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        if size > int64(fileData.MINDOWNLOADABLEFILESIZE) {
            if size < int64(fileData.MAXDOWNLOADABLEFILESIZE) {
                if !StringValidationUtil_GetInstance().IsValidRequired(fileName, fileData.MINLEN, fileData.MAXLEN) {
                    stringBuffer.Append("File Name must be more than ")
                    stringBuffer.Appendint(fileData.MINLEN)
                    stringBuffer.Append(" and less than ")
                    stringBuffer.Appendint(fileData.MAXLEN)
                    stringBuffer.Append(" characters in length: ")
                    stringBuffer.Append(fileName)
                    stringBuffer.Append("<br/>")
                } else {
                }
            } else {
                stringBuffer.Append("File Is Not The Right Size. ")
                stringBuffer.Appendint(fileData.MINDOWNLOADABLEFILESIZE)
                stringBuffer.Append("< >")
                stringBuffer.Appendint(fileData.MAXDOWNLOADABLEFILESIZE)
                stringBuffer.Append("<br/>")
            }
        }
}

