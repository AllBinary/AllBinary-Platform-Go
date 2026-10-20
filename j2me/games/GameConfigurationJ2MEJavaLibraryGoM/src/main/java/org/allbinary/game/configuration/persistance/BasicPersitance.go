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

type BasicPersitance struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PersistanceStrings *PersistanceStrings
    PlatformRecordIdUtil *PlatformRecordIdUtil
    RecordId string
    ValueList IBasicArrayList
    IdList IBasicArrayList
}

// The Java class BasicPersitance has subclasses so its variables are this interface
type IBasicPersitance interface {
    BasicPersitance_() *BasicPersitance
    DeleteRecoreStore(abeClientInformation AbeClientInformationInterface)
    DeleteAll(abeClientInformation AbeClientInformationInterface)
    Delete(abeClientInformation AbeClientInformationInterface, deleteId int)
    GetRecordId(abeClientInformation AbeClientInformationInterface) string
    GetList() IBasicArrayList
    GetIds() IBasicArrayList
    Clear()
}

func (this *BasicPersitance) BasicPersitance_() *BasicPersitance {
    return this
}


func (this *BasicPersitance) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.PersistanceStrings = PersistanceStrings_GetInstance()
    this.PlatformRecordIdUtil = PlatformRecordIdUtil_GetInstance()
    
    this.ValueList = NewBasicArrayListD()
    this.IdList = NewBasicArrayListD()
}


func NewBasicPersitance(recordId string) *BasicPersitance {
    this := &BasicPersitance{}
    this.initFields()
        this.RecordId = recordId
    return this
}


// throws Exception
func (this *BasicPersitance) DeleteRecoreStore(abeClientInformation AbeClientInformationInterface) {
        RecordStore_DeleteRecordStore(this.GetRecordId(abeClientInformation))
}


// throws Exception
func (this *BasicPersitance) DeleteAll(abeClientInformation AbeClientInformationInterface) {
        var size int = this.IdList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var integer int = any(this.IdList.BasicArrayList_().ObjectArray[index]).(int); _ = integer
            this.Delete(abeClientInformation, int(integer))
        }
        this.Clear()
}


// throws Exception
func (this *BasicPersitance) Delete(abeClientInformation AbeClientInformationInterface, deleteId int) {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            if recordStore != nil {
                PreLogUtil_Put(this.PersistanceStrings.CLOSING_RECORDSTORE, this, this.CommonStrings.Delete)
                recordStore.CloseRecordStore()
            }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.PersistanceStrings.DELETING_WITH_ID).Appendint(deleteId)), this, this.CommonStrings.Delete)
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
            recordStore.DeleteRecord(deleteId)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BasicPersitance) GetRecordId(abeClientInformation AbeClientInformationInterface) string {
        return this.PlatformRecordIdUtil.GetRecordId(abeClientInformation, this.RecordId)
}


func (this *BasicPersitance) GetList() IBasicArrayList {
        return this.ValueList
}


func (this *BasicPersitance) GetIds() IBasicArrayList {
        return this.IdList
}


func (this *BasicPersitance) Clear() {
        this.ValueList.Clear()
        this.IdList.Clear()
}

