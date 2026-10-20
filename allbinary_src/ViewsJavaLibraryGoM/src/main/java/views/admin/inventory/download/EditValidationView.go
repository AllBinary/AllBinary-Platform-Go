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

type EditValidationView struct {
    DownloadableInventoryItemView
    DownloadItemId string
}

// Implements ValidationComponentInterface


func (this *EditValidationView) initFields() {
    
}


func NewEditValidationView(transformInfoInterface TransformInfoInterface) *EditValidationView {
    this := &EditValidationView{}
        this.DownloadableInventoryItemView = *NewDownloadableInventoryItemView(transformInfoInterface, StringUtil_GetInstance().EMPTY_STRING)
        this.initFields()
        this.GetFormData()
    return this
}


// throws Exception
func (this *EditValidationView) GetFormData() {
        var hashMap *HashMap = NewRequestParams(this.Request).ToHashMap(); _ = hashMap
        if hashMap == nil {
            panic(errors.New(fmt.Sprint("No Request Params Found")))
        }
        this.Id = any(hashMap.Get(this.BasicItemData.ID)).(string)
        this.DownloadItemId = any(hashMap.Get(DownloadItemData_ID)).(string)
}


// throws Exception
func (this *EditValidationView) IsValid() bool {
        if this.Id == nil || !StringValidationUtil_GetInstance().IsNumber(this.Id) {
            return false
        }
        this.ItemInterface = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.Id)
        if this.ItemInterface == nil {
            return false
        }
        var downloadable int = func() int { value, _ := strconv.Atoi(this.ItemInterface.GetDownloads()); return value }(); _ = downloadable
        if downloadable != 0 {
            var downloadItemsEntity *DownloadItemsEntity = DownloadItemsEntityFactory_GetInstance().GetDownloadItemsEntityInstance(); _ = downloadItemsEntity
            var vector *BasicArrayList = downloadItemsEntity.GetForItem(this.Id, this.DownloadItemId); _ = vector
            if vector.Size() != 1 {
                return false
            }
            this.DownloadableItem = any(vector.Get(0)).(*DownloadableItem)
        } else {
            return false
        }
        return true
}


func (this *EditValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *EditValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws MoneyException
func (this *EditValidationView) ValidationInfo() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        if this.Id == nil || !StringValidationUtil_GetInstance().IsNumber(this.Id) {
            stringBuffer.Append("Id is not valid.<br />")
        }
        if InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().GetItem(this.Id) == nil {
            stringBuffer.Append("Item does not exist.<br />")
        }
        var downloadable int = func() int { value, _ := strconv.Atoi(this.ItemInterface.GetDownloads()); return value }(); _ = downloadable
        if downloadable != 0 {
            if this.DownloadableItem == nil {
                stringBuffer.Append("DownloadableItem does not exist for item.<br />")
            }
        } else {
            stringBuffer.Append("Item should not have DownloadableItem.<br />")
        }
        return fmt.Sprint(stringBuffer)
}

