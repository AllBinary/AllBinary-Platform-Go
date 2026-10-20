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

type InventoryRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    ItemInterface ItemInterface
}


func (this *InventoryRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
}


func NewInventoryRequestHelper(propertiesHashMap *HashMap, pageContext *PageContext) *InventoryRequestHelper {
    this := &InventoryRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        var inventoryTag IHelperTag = any(propertiesHashMap.Get(AbTagData_PARENT)).(IHelperTag); _ = inventoryTag
        var itemView IInventoryItemView = InventoryItemViewParentTagHelper_GetInstance().GetInventoryItemView(inventoryTag); _ = itemView
        this.ItemInterface = itemView.GetItemInterface()
    return this
}


// throws Exception
func (this *InventoryRequestHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add item to Inventory"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "insert()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var dataMappingInterface TableMappingInterface = any(this.GetItemInterface()).(TableMappingInterface); _ = dataMappingInterface
            var values IBasicArrayList = dataMappingInterface.ToVector(); _ = values
            InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().Insert(values)
            var success string = "Successfully Added Product"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "insert()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InventoryRequestHelper) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to remove item with from Inventory"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.TagStrings.DO_START_TAG, e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var dataMappingInterface TableMappingInterface = any(this.GetItemInterface()).(TableMappingInterface); _ = dataMappingInterface
            var id string = any(dataMappingInterface.GetKey()).(string); _ = id
            InventoryEntityFactory_GetInstance().GetInventoryEntityInstance().DeleteWhere(this.BasicItemData.ID, id)
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Successfully Removed the item with ")
            stringBuffer.Append(this.BasicItemData.ID)
            stringBuffer.Append("=")
            stringBuffer.Append(id)
            stringBuffer.Append(" from to the Inventory table")
            var success string = fmt.Sprint(stringBuffer); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, this.TagStrings.DO_START_TAG)
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InventoryRequestHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed To Update Item"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            this.Delete()
            this.Insert()
            var success string = "Item Successfully Updated"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "update()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InventoryRequestHelper) GetItemInterface() ItemInterface {
        return this.ItemInterface
}

