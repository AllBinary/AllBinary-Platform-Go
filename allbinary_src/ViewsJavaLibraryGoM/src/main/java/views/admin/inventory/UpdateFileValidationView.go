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
    InventoryItemView
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface

var UpdateFileValidationView_UPDATEPRODUCT string = CommonStrings_GetInstance().UPDATE

func (this *UpdateFileValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewUpdateFileValidationView(transformInfoInterface TransformInfoInterface) *UpdateFileValidationView {
    this := &UpdateFileValidationView{}
        this.InventoryItemView = *NewInventoryItemView(transformInfoInterface)
        this.initFields()
    return this
}


func (this *UpdateFileValidationView) IsValid_1() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.Put("Exception in validation", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var command string = any(this.GetRequestHashMap().Get(InventoryItemView_GLOBALS2.ADMINCOMMAND)).(string); _ = command
            if command == "" || strings.Compare(command, UpdateFileValidationView_UPDATEPRODUCT) != 0 {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Invalid AdminCommand=" + command, this, this.CommonStrings.IS_VALID)
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
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(InventoryItemView_URLGLOBALS.GetWebappPath())
            stringBuffer.Append(storeFrontInterface.GetCurrentHostNamePath())
            stringBuffer.Append(this.ItemInterface.GetCategory())
            var fullCategory string = fmt.Sprint(stringBuffer); _ = fullCategory
            var categoryFile *AbFile = AbFile_CreateAbFile(fullCategory); _ = categoryFile
            if !categoryFile.IsDirectory() {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("Category Does Not Exist: " + fullCategory, this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            var inventoryEntity *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntity
            if inventoryEntity.GetItem(this.ItemInterface.GetId()) == nil {
                tryControl1, tryResult1 = 1, false; return
            }
            var anyType any = this.GetRequestHashMap().Get(this.BasicItemData.IMAGE); _ = anyType
            if HttpFileUploadUtil_GetInstance().IsValid(anyType) {
                var fileItem *FileItem = any(anyType).(*FileItem); _ = fileItem
                var size int64 = fileItem.GetSize(); _ = size
                var fileName string = fileItem.GetName(); _ = fileName
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


// throws Exception
func (this *UpdateFileValidationView) ValidationInfo_1() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var command string = any(this.GetRequestHashMap().Get(InventoryItemView_GLOBALS2.ADMINCOMMAND)).(string); _ = command
        if command == "" || strings.Compare(command, UpdateFileValidationView_UPDATEPRODUCT) != 0 {
            return CommonSeps_GetInstance().SPACE
        }
        stringBuffer.Append(NewBasicItemValidation(this.ItemInterface).ValidationInfo())
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
        var fullCategory string = any(InventoryItemView_URLGLOBALS.GetWebappPath()).(string) + storeFrontInterface.GetCurrentHostNamePath() + this.ItemInterface.GetCategory(); _ = fullCategory
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Category: " + fullCategory, this, "validationInfo()")
        }
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Existing Item With MoneyException", this, "validationInfo()")
            }
                }
            }()
            if InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.ItemInterface.GetId()) == nil {
                stringBuffer.Append("Item does not exist.<br>")
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        var anyType any = this.GetRequestHashMap().Get(this.BasicItemData.IMAGE); _ = anyType
        if HttpFileUploadUtil_GetInstance().IsValid(anyType) {
            var fileItem *FileItem = any(anyType).(*FileItem); _ = fileItem
            var size int64 = fileItem.GetSize(); _ = size
            var fileName string = fileItem.GetName(); _ = fileName
            var fileItemFieldName string = fileItem.GetFieldName(); _ = fileItemFieldName
            this.ValidationInfo(stringBuffer, fileName, fileItemFieldName, size)
        }
        return fmt.Sprint(stringBuffer)
}

// Java overloads of UpdateFileValidationView.validationInfo
func (this *UpdateFileValidationView) ValidationInfo(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.ValidationInfo_1()
        }
        if len(args) == 4 && Java_IsArg[*StringMaker](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[int64](args[3], pass) {
            return this.ValidationInfo_2(Java_Arg[*StringMaker](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[int64](args[3]))
        }
    }
    panic("No Java overload of UpdateFileValidationView.validationInfo for the arguments")
}


// throws Exception
func (this *UpdateFileValidationView) IsValid_2(fileName string, size int64) bool {
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        if size > int64(fileData.MINIMAGEFILESIZE) {
            if size < int64(fileData.MAXIMAGEFILESIZE) {
                if !StringValidationUtil_GetInstance().IsValidRequired(fileName, fileData.MINLEN, fileData.MAXLEN) {
                    return false
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
        }
        return true
}


func (this *UpdateFileValidationView) ValidationInfo_2(stringBuffer *StringMaker, fileName string, fileItemFieldName string, size int64) {
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        if size > int64(fileData.MINIMAGEFILESIZE) {
            if strings.Compare(fileItemFieldName, this.BasicItemData.IMAGE) == 0 {
                if size < int64(fileData.MAXIMAGEFILESIZE) {
                    if !StringValidationUtil_GetInstance().IsValidRequired(fileName, fileData.MINLEN, fileData.MAXLEN) {
                        stringBuffer.Append("FileName must be >")
                        stringBuffer.Appendint(fileData.MINLEN)
                        stringBuffer.Append("and <")
                        stringBuffer.Appendint(fileData.MAXLEN)
                        stringBuffer.Append("<br />")
                    }
                } else {
                    stringBuffer.Append("Image File Is Not The Right Size. ")
                    stringBuffer.Appendint(fileData.MINIMAGEFILESIZE)
                    stringBuffer.Append("< > ")
                    stringBuffer.Appendint(fileData.MAXIMAGEFILESIZE)
                    stringBuffer.Append("<br />")
                }
            }
        }
}

