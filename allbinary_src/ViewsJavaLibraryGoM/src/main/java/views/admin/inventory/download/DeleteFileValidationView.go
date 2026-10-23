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

type DeleteFileValidationView struct {
    DownloadableInventoryItemView
    DownloadItemId string
}

// Implements ValidationComponentInterface


func (this *DeleteFileValidationView) initFields() {
    
}


func NewDeleteFileValidationView(transformInfoInterface TransformInfoInterface) *DeleteFileValidationView {
    this := &DeleteFileValidationView{}
        this.DownloadableInventoryItemView = *NewDownloadableInventoryItemView(transformInfoInterface, StringUtil_GetInstance().EMPTY_STRING)
        this.initFields()
        this.GetFormData()
    return this
}


// throws Exception
func (this *DeleteFileValidationView) GetFormData() {
        var hashMap *HashMap = NewRequestParams(this.Request).ToHashMap(); _ = hashMap
        if hashMap == nil {
            panic(errors.New(fmt.Sprint("No Request Params Found")))
        }
        this.Id = any(hashMap.Get(this.BasicItemData.ID)).(string)
        this.DownloadItemId = any(hashMap.Get(DownloadItemData_ID)).(string)
}


func (this *DeleteFileValidationView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if this.Id == nil || !StringValidationUtil_GetInstance().IsNumber(this.Id) {
                tryControl1, tryResult1 = 1, false; return
            }
            var inventoryEntity *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntity
            this.ItemInterface = inventoryEntity.GetItem(this.Id)
            if this.ItemInterface == nil {
                tryControl1, tryResult1 = 1, false; return
            }
            var downloadItemsEntity *DownloadItemsEntity = DownloadItemsEntityFactory_GetInstance().GetDownloadItemsEntityInstance(); _ = downloadItemsEntity
            var vector IBasicArrayList = downloadItemsEntity.GetForItem(this.Id, this.DownloadItemId); _ = vector
            if vector.Size() != 1 {
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DeleteFileValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *DeleteFileValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *DeleteFileValidationView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, "Unkown Validation Error.<br />"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if this.Id == nil || !StringValidationUtil_GetInstance().IsNumber(this.Id) {
                stringBuffer.Append("Id is not valid.<br />")
            }
            if InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.Id) == nil {
                stringBuffer.Append("Item does not exist.<br />")
            }
            var downloadItemsEntity *DownloadItemsEntity = DownloadItemsEntityFactory_GetInstance().GetDownloadItemsEntityInstance(); _ = downloadItemsEntity
            var vector IBasicArrayList = downloadItemsEntity.GetForItem(this.Id, this.DownloadItemId); _ = vector
            if vector.Size() == 0 {
                stringBuffer.Append("DownloadItem does not exist.<br />")
            }
            if vector.Size() > 1 {
                stringBuffer.Append("To Many DownloadItems exist.<br />")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

