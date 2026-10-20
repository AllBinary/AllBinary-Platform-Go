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

type HighScoreNamePersistanceSingleton struct {
    LogUtil *LogUtil
    TsUtil *TsUtil
    CommonStrings *CommonStrings
    PlatformRecordIdUtil *PlatformRecordIdUtil
    RECORD_ID string
    Name string
    NameBasicArrayList IBasicArrayList
}

var HighScoreNamePersistanceSingleton_SINGLETON *HighScoreNamePersistanceSingleton = NewHighScoreNamePersistanceSingleton()

func (this *HighScoreNamePersistanceSingleton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.PlatformRecordIdUtil = PlatformRecordIdUtil_GetInstance()
    this.RECORD_ID = "_SN"
    this.Name = StringUtil_GetInstance().EMPTY_STRING
    this.NameBasicArrayList = NewBasicArrayListD()
}

func NewHighScoreNamePersistanceSingleton() *HighScoreNamePersistanceSingleton {
    this := &HighScoreNamePersistanceSingleton{}
    this.initFields()
    return this
}


func HighScoreNamePersistanceSingleton_GetInstance() *HighScoreNamePersistanceSingleton {
        return HighScoreNamePersistanceSingleton_SINGLETON
}


func (this *HighScoreNamePersistanceSingleton) Clear() {
        this.Name = StringUtil_GetInstance().EMPTY_STRING
}


// throws Exception
func (this *HighScoreNamePersistanceSingleton) DeleteAll(abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo) {
        var size int = this.NameBasicArrayList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var integer int = any(this.NameBasicArrayList.BasicArrayList_().ObjectArray[index]).(int); _ = integer
            this.Delete(abeClientInformation, gameInfo, int(integer))
        }
        this.Clear()
}


func (this *HighScoreNamePersistanceSingleton) GetRecordId(abeClientInformation AbeClientInformationInterface) string {
        return this.PlatformRecordIdUtil.GetRecordId(abeClientInformation, this.RECORD_ID)
}


// throws Exception
func (this *HighScoreNamePersistanceSingleton) Delete(abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo, deleteId int) {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            if recordStore != nil {
                PreLogUtil_Put("Closing RecordStore", this, this.CommonStrings.Delete)
                recordStore.CloseRecordStore()
            }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Deleting: ").Appendint(deleteId)), this, this.CommonStrings.Delete)
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
            recordStore.DeleteRecord(deleteId)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *HighScoreNamePersistanceSingleton) GetIds() IBasicArrayList {
        return this.NameBasicArrayList
}


func (this *HighScoreNamePersistanceSingleton) Load(abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo) string {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.Save(abeClientInformation, gameInfo, this.Name)
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.EXCEPTION_LABEL).Append(ExceptionUtil_GetInstance().GetStackTrace(e))), this, this.CommonStrings.LOAD)
                }
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.LOAD, e)
                    }
                }()
                if recordStore != nil {
                    PreLogUtil_Put("Closing RecordStore", this, this.CommonStrings.LOAD)
                    recordStore.CloseRecordStore()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
            }()
            if this.Name == StringUtil_GetInstance().EMPTY_STRING {
                var LOADING_ID string = "Loading id: "; _ = LOADING_ID
                recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
                var recordEnum RecordEnumeration = recordStore.EnumerateRecords(NullRecordFilter_NULL_RECORD_FILTER, NullRecordComparator_NULL_RECORD_COMPARATOR, true); _ = recordEnum
                var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
                var recordAsBytes []int8; _ = recordAsBytes
                var byteArrayInputStream *ByteArrayInputStream; _ = byteArrayInputStream
                var inputStream IDataInputStream; _ = inputStream
                for recordEnum.HasNextElement() {
                    var id int = recordEnum.NextRecordId(); _ = id
                    this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(LOADING_ID).Appendint(id)), this, this.CommonStrings.LOAD)
                    recordAsBytes = this.TsUtil.GetRecord(recordStore, id)
                    byteArrayInputStream = NewByteArrayInputStream(recordAsBytes)
                    inputStream = NewDataInputStream(byteArrayInputStream)
                    for inputStream.Available() > 0 {
                        this.Name = inputStream.ReadUTF()
                    }
                    this.NameBasicArrayList.Add(smallIntegerSingletonFactory.GetAt(id))
                }
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return this.Name
}


func (this *HighScoreNamePersistanceSingleton) Save(abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo, name string) {
        var recordStore IRecordStore = NullRecordStore_NULL_RECORD_STORE; _ = recordStore
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.SAVE, e)
                }
            if tryControl2 := func() (tryControl2 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.SAVE, e)
                    }
                }()
                if recordStore != nil {
                    PreLogUtil_Put("Closing RecordStore", this, this.CommonStrings.SAVE)
                    recordStore.CloseRecordStore()
                }
                return
            }(); tryControl2 == 1 {
                tryControl1 = 1; return
            }
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Saving: ").Append(name)), this, this.CommonStrings.SAVE)
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
            var byteArrayOutputStream *strings.Builder = &strings.Builder{}; _ = byteArrayOutputStream
            var outputStream IDataOutputStream = NewDataOutputStream(byteArrayOutputStream); _ = outputStream
            outputStream.WriteUTF(name)
            var savedGameBytes []int8 = byteArrayOutputStream.ToByteArray(); _ = savedGameBytes
            recordStore.AddRecord(savedGameBytes, 0, len(savedGameBytes))
            this.Name = name
            return
        }(); tryControl1 == 1 {
            return
        }
}

