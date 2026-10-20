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

//GameInputNoPersistanceLibrary
type InputPersistance struct {
    BasicPersitance
    HashtableUtil *HashtableUtil
}


func (this *InputPersistance) initFields() {
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
        var recordStore IRecordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true); _ = recordStore
        var recordEnum RecordEnumeration = recordStore.EnumerateRecords(NullRecordFilter_NULL_RECORD_FILTER, NullRecordComparator_NULL_RECORD_COMPARATOR, true); _ = recordEnum
        var hashtable *ABHashtable[any, any]; _ = hashtable
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var id int; _ = id
        for recordEnum.HasNextElement() {
            id = recordEnum.NextRecordId()
            stringBuffer.Delete(0, stringBuffer.Length())
            this.LogUtil.PutF(fmt.Sprint(stringBuffer.Append(this.PersistanceStrings.LOADING_ID).Appendint(id)), this, this.PersistanceStrings.LOAD_ALL)
            hashtable = StdUtil_GetInstance().CreateHashtable()
            this.ValueList.Add(hashtable)
            this.IdList.Add(smallIntegerSingletonFactory.GetAt(id))
        }
        recordStore.CloseRecordStore()
}


// throws Exception
func (this *InputPersistance) Save(abeClientInformation AbeClientInformationInterface, hashtable *ABHashtable[any, any]) {
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.PersistanceStrings.NOT_SAVING).Append(StringUtil_GetInstance().ToString(hashtable))), this, this.CommonStrings.SAVE)
        var recordStore IRecordStore = RecordStore_OpenRecordStore(this.GetRecordId(abeClientInformation), true); _ = recordStore
        var gameActionInput IInput; _ = gameActionInput
        var list IBasicArrayList; _ = list
        var input IInput; _ = input
        var inputObjectArray []any = this.HashtableUtil.GetKeysAsArray(hashtable); _ = inputObjectArray
        var size int = len(inputObjectArray); _ = size
        for index := int(0); index < size; index++ {
            gameActionInput = any(inputObjectArray[index]).(IInput)
            list = any(hashtable.Get(inputObjectArray[index])).(IBasicArrayList)
            for index2 := int(0); index2 < list.Size(); index2++ {
                input = any(list.Get(index2)).(IInput)
            }
        }
        recordStore.CloseRecordStore()
}

