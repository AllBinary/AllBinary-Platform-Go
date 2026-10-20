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

//GameInputPersistanceLibrary
type InputPersistance struct {
    BasicPersitance
    TsUtil *TsUtil
    HashtableUtil *HashtableUtil
}


func (this *InputPersistance) initFields() {
    this.TsUtil = TsUtil_GetInstance()
    this.HashtableUtil = HashtableUtil_GetInstance()
}


func NewInputPersistance(name string) *InputPersistance {
    this := &InputPersistance{}
        this.BasicPersitance = *NewBasicPersitance(name)
        this.initFields()
    return this
}


// throws Exception
func (this *InputPersistance) LoadAll(abeClientInformation AbeClientInformationInterface) {
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
            var gameActionInputId int64; _ = gameActionInputId
            var inputId int64; _ = inputId
            var gameActionInput IInput; _ = gameActionInput
            var input IInput; _ = input
            var hashtable *ABHashtable[any, any]; _ = hashtable
            var gameKeyFactory *GameKeyMappingFactory = GameKeyMappingFactory_GetInstance(); _ = gameKeyFactory
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            var inputFactory *InputFactory = InputFactory_GetInstance(); _ = inputFactory
            var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
            var stdUtil *StdUtil = StdUtil_GetInstance(); _ = stdUtil
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
                    hashtable = stdUtil.CreateHashtable()
                    var value int; _ = value
                    for inputStream.Available() > 0 {
                        var gameActionInputIdAsString string = inputStream.ReadUTF(); _ = gameActionInputIdAsString
                        value = func() int { value, _ := strconv.Atoi(gameActionInputIdAsString); return value }()
                        gameActionInputId = int64(value)
                        inputStream.ReadUTF()
                        value = func() int { value, _ := strconv.Atoi(inputStream.ReadUTF()); return value }()
                        inputId = int64(value)
                        gameActionInput = gameKeyFactory.GetGameKey(int(gameActionInputId))
                        input = inputFactory.GetInstanceById(int(inputId))
                        if input == inputFactory.NO_INPUT || gameActionInput == nil {
                            stringBuffer.Delete(0, stringBuffer.Length())
                            if input == inputFactory.NO_INPUT {
                                stringBuffer.Append(this.PersistanceStrings.ERROR_LOADING_ID)
                                stringBuffer.Appendlong(inputId)
                                stringBuffer.Append(this.PersistanceStrings.GAME_ACTION_INPUT)
                                stringBuffer.Appendlong(gameActionInputId)
                                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, this.PersistanceStrings.LOAD_ALL)
                            }
                            if gameActionInput == nil {
                                stringBuffer.Delete(0, stringBuffer.Length())
                                stringBuffer.Append(this.PersistanceStrings.ERROR_LOADING)
                                stringBuffer.Appendlong(gameActionInputId)
                                stringBuffer.Append(this.PersistanceStrings.ID)
                                stringBuffer.Appendlong(inputId)
                                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, this.PersistanceStrings.LOAD_ALL)
                            }
                        } else {
                        }
                        hashtable.Put(input, gameActionInput)
                    }
                    this.ValueList.Add(hashtable)
                    this.IdList.Add(smallIntegerSingletonFactory.GetAt(id))
                } else {
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *InputPersistance) Save(abeClientInformation AbeClientInformationInterface, hashtable *ABHashtable[any, any]) {
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
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            PreLogUtil_Put(fmt.Sprint(stringBuffer.Append(this.PersistanceStrings.SAVING).Append(StringUtil_GetInstance().ToString(hashtable))), this, this.CommonStrings.SAVE)
            recordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true)
            var byteArrayOutputStream *strings.Builder = &strings.Builder{}; _ = byteArrayOutputStream
            var outputStream IDataOutputStream = NewDataOutputStream(byteArrayOutputStream); _ = outputStream
            var gameActionInput IInput; _ = gameActionInput
            var list IBasicArrayList; _ = list
            var input IInput; _ = input
            var savedGameBytes []int8; _ = savedGameBytes
            var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
            var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
            var inputObjectArray []any = this.HashtableUtil.GetKeysAsArray(hashtable); _ = inputObjectArray
            var size int = len(inputObjectArray); _ = size
            for index := int(0); index < size; index++ {
                gameActionInput = any(inputObjectArray[index]).(IInput)
                list = any(hashtable.Get(any(inputObjectArray[index]).(any))).(IBasicArrayList)
                for index2 := int(0); index2 < list.Size(); index2++ {
                    var gameActionInputIdAsString string = fmt.Sprint(smallIntegerSingletonFactory.GetAt(gameActionInput.GetId())); _ = gameActionInputIdAsString
                    outputStream.WriteUTF(gameActionInputIdAsString)
                    outputStream.WriteUTF(commonSeps.EQUALS)
                    input = any(list.BasicArrayList_().ObjectArray[index2]).(IInput)
                    var inputIdAsString string = fmt.Sprint(smallIntegerSingletonFactory.GetAt(input.GetId())); _ = inputIdAsString
                    outputStream.WriteUTF(inputIdAsString)
                }
            }
            savedGameBytes = byteArrayOutputStream.ToByteArray()
            recordStore.AddRecord(savedGameBytes, 0, len(savedGameBytes))
            return
        }(); tryControl1 == 1 {
            return
        }
}

