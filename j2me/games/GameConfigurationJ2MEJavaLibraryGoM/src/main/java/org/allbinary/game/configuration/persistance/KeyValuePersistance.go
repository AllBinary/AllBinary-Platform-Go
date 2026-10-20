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

type KeyValuePersistance struct {
    BasicPersitance
    TsUtil *TsUtil
}

// The Java class KeyValuePersistance has subclasses so its variables are this interface
type IKeyValuePersistance interface {
    IBasicPersitance
    KeyValuePersistance_() *KeyValuePersistance
    LoadAll(abeClientInformation AbeClientInformationInterface)
    LoadAllSize(abeClientInformation AbeClientInformationInterface, size int)
    Save(abeClientInformation AbeClientInformationInterface, hashtable *ABHashtable[any, any])
    Get(index int) *ABHashtable[any, any]
}

func (this *KeyValuePersistance) KeyValuePersistance_() *KeyValuePersistance {
    return this
}


func (this *KeyValuePersistance) initFields() {
    this.TsUtil = TsUtil_GetInstance()
}


func NewKeyValuePersistance(recordId string) *KeyValuePersistance {
    this := &KeyValuePersistance{}
        this.BasicPersitance = *NewBasicPersitance(recordId)
        this.initFields()
    return this
}


// throws Exception
func (this *KeyValuePersistance) LoadAll(abeClientInformation AbeClientInformationInterface) {
        this.LoadAllSize(abeClientInformation, 1)
}


// throws Exception
func (this *KeyValuePersistance) LoadAllSize(abeClientInformation AbeClientInformationInterface, size int) {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            if recordStore != nil {
                PreLogUtil_Put(this.PersistanceStrings.CLOSING_RECORDSTORE, this, this.PersistanceStrings.LOAD_ALL)
                recordStore.CloseRecordStore()
            }
            }()
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
            var recordEnum RecordEnumeration = recordStore.EnumerateRecords(NullRecordFilter_NULL_RECORD_FILTER, NullRecordComparator_NULL_RECORD_COMPARATOR, true); _ = recordEnum
            var hashtable *ABHashtable[string, string]; _ = hashtable
            var name string; _ = name
            var value string; _ = value
            var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var recordAsBytes []int8; _ = recordAsBytes
            var byteArrayInputStream *ByteArrayInputStream; _ = byteArrayInputStream
            var inputStream IDataInputStream; _ = inputStream
            for recordEnum.HasNextElement() {
                var id int = recordEnum.NextRecordId(); _ = id
                stringBuffer.Delete(0, stringBuffer.Length())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer.Append(this.PersistanceStrings.LOADING_ID).Appendint(id)), this, this.PersistanceStrings.LOAD_ALL)
                recordAsBytes = this.TsUtil.GetRecord(recordStore, id)
                if recordAsBytes != nil {
                    byteArrayInputStream = NewByteArrayInputStream(recordAsBytes)
                    inputStream = NewDataInputStream(byteArrayInputStream)
                    hashtable = NewABHashtable[string, string]()
                    for index := int(0); index < size; index++ {
                        name = inputStream.ReadUTF()
                        inputStream.ReadUTF()
                        value = inputStream.ReadUTF()
                        hashtable.Put(name, value)
                    }
                    this.ValueList.Add(hashtable)
                    this.IdList.Add(smallIntegerSingletonFactory.GetAt(id))
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *KeyValuePersistance) Save(abeClientInformation AbeClientInformationInterface, hashtable *ABHashtable[any, any]) {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            if recordStore != nil {
                PreLogUtil_Put(this.PersistanceStrings.CLOSING_RECORDSTORE, this, this.CommonStrings.SAVE)
                recordStore.CloseRecordStore()
            }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.PersistanceStrings.SAVING).Append(StringUtil_GetInstance().ToString(hashtable))), this, this.CommonStrings.SAVE)
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
            var byteArrayOutputStream *strings.Builder = &strings.Builder{}; _ = byteArrayOutputStream
            var outputStream IDataOutputStream = NewDataOutputStream(byteArrayOutputStream); _ = outputStream
            var value string; _ = value
            var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
            var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
            var anyType any; _ = anyType
            var size int = len(objectArray); _ = size
            for index := int(0); index < size; index++ {
                outputStream.WriteUTF(any(objectArray[index]).(string))
                outputStream.WriteUTF(commonSeps.EQUALS)
                anyType = any(objectArray[index]).(any)
                value = any(hashtable.Get(anyType)).(string)
                outputStream.WriteUTF(value)
            }
            var tsUtil *TsUtil = TsUtil_GetInstance(); _ = tsUtil
            var savedGameBytes []int8 = tsUtil.GetByteArray(byteArrayOutputStream.String()); _ = savedGameBytes
            recordStore.AddRecord(savedGameBytes, 0, len(savedGameBytes))
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *KeyValuePersistance) Get(index int) *ABHashtable[any, any] {
        var hashtable *ABHashtable[any, any] = any(this.ValueList.BasicArrayList_().ObjectArray[index]).(*ABHashtable[any, any]); _ = hashtable
        return hashtable
}

