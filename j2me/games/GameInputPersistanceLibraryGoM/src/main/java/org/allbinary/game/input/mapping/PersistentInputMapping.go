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
type PersistentInputMapping struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    EnumerationUtil *EnumerationUtil
    InputMapping *InputToGameKeyMapping
    InputMappingEventListenerInterface InputMappingEventListenerInterface
    InputPersistance *InputPersistance
    InputMappingEvent *InputMappingEvent
}

// The Java class PersistentInputMapping has subclasses so its variables are this interface
type IPersistentInputMapping interface {
    PersistentInputMapping_() *PersistentInputMapping
    GetTotalMapped() int
    IsDefaultNew() bool
    GetDefault() *InputToGameKeyMapping
    SetDefault(abeClientInformation AbeClientInformationInterface)
    Update(abeClientInformation AbeClientInformationInterface)
    Save(abeClientInformation AbeClientInformationInterface)
    Init(abeClientInformation AbeClientInformationInterface)
    SetInputMappingEventListenerInterface(inputMappingEventListenerInterface InputMappingEventListenerInterface)
    GetInputMappingEventListenerInterface() InputMappingEventListenerInterface
    GetInputMapping() *InputToGameKeyMapping
    IsDelete(input IInput) bool
    IsSystemInput(input IInput) bool
}

func (this *PersistentInputMapping) PersistentInputMapping_() *PersistentInputMapping {
    return this
}

var PersistentInputMapping_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *PersistentInputMapping) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.EnumerationUtil = EnumerationUtil_GetInstance()
    this.InputMapping = NewInputToGameKeyMapping()
    this.InputMappingEventListenerInterface = NullInputMappingEventListener_NULL_INPUT_MAPPING_EVENT_LISTENER
    
    this.InputMappingEvent = NewInputMappingEvent(this)
}


func PersistentInputMapping_GetNullInstance() IPersistentInputMapping {
        if PersistentInputMapping_Instance == NullUtil_GetInstance().NULL_OBJECT {
            PersistentInputMapping_Instance = NewPersistentInputMapping(GamePersistanceStrings_GetInstance().SAVED_INPUT_CONFIGURATION_RECORD_ID)
        }
        return any(PersistentInputMapping_Instance).(IPersistentInputMapping)
}


func NewPersistentInputMapping(name string) *PersistentInputMapping {
    this := &PersistentInputMapping{}
    this.initFields()
        this.InputPersistance = NewInputPersistance(name)
    return this
}


func (this *PersistentInputMapping) GetTotalMapped() int {
        return this.GetInputMapping().GetHashtable().Size()
}


func (this *PersistentInputMapping) IsDefaultNew() bool {
        return false
}


func (this *PersistentInputMapping) GetDefault() *InputToGameKeyMapping {
        return InputToGameKeyMapping_GetNullInstance()
}


// throws Exception
func (this *PersistentInputMapping) SetDefault(abeClientInformation AbeClientInformationInterface) {
        this.GetInputMapping().RemoveAll()
        this.GetInputMapping().AddMapping(this.GetDefault())
        this.Save(abeClientInformation)
}


// throws Exception
func (this *PersistentInputMapping) Update(abeClientInformation AbeClientInformationInterface) {
        this.InputPersistance.DeleteAll(abeClientInformation)
        this.Save(abeClientInformation)
}


// throws Exception
func (this *PersistentInputMapping) Save(abeClientInformation AbeClientInformationInterface) {
        this.InputPersistance.Save(abeClientInformation, this.GetInputMapping().GetHashtable())
        this.InputMappingEvent.SetInputToGameKeyMapping(this.GetInputMapping())
        if this.GetInputMappingEventListenerInterface() != nil {
            this.GetInputMappingEventListenerInterface().OnInputMappingEvent(this.InputMappingEvent)
        }
}


// throws Exception
func (this *PersistentInputMapping) Init(abeClientInformation AbeClientInformationInterface) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.INIT)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
            this.InputPersistance.DeleteRecoreStore(abeClientInformation)
            this.SetDefault(abeClientInformation)
            this.InputPersistance.LoadAll(abeClientInformation)
                }
            }()
            this.InputPersistance.LoadAll(abeClientInformation)
            return
        }(); tryControl1 == 1 {
            return
        }
        var list IBasicArrayList = this.InputPersistance.GetList(); _ = list
        var totalMappedTo int = 0; _ = totalMappedTo
        var size int = list.Size(); _ = size
        var hashtable *ABHashtable[any, any]; _ = hashtable
        var enumeration Enumeration; _ = enumeration
        var mappedToInput IInput; _ = mappedToInput
        var gameActionInput IInput; _ = gameActionInput
        for index := int(0); index < size; index++ {
            hashtable = any(list.BasicArrayList_().ObjectArray[index]).(*ABHashtable[any, any])
            enumeration = hashtable.Keys()
            for this.EnumerationUtil.HasMoreElements(enumeration) {
                mappedToInput = any(this.EnumerationUtil.NextElement(enumeration)).(IInput)
                gameActionInput = any(hashtable.Get(mappedToInput)).(IInput)
                totalMappedTo++
                this.GetInputMapping().Add(gameActionInput, mappedToInput)
            }
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("End - Total Loaded Keys Mapped: ")
        stringBuffer.Appendint(this.GetTotalMapped())
        stringBuffer.Append(" to: ")
        stringBuffer.Appendint(totalMappedTo)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.INIT)
}


func (this *PersistentInputMapping) SetInputMappingEventListenerInterface(inputMappingEventListenerInterface InputMappingEventListenerInterface) {
        this.InputMappingEventListenerInterface = inputMappingEventListenerInterface
}


func (this *PersistentInputMapping) GetInputMappingEventListenerInterface() InputMappingEventListenerInterface {
        return this.InputMappingEventListenerInterface
}


func (this *PersistentInputMapping) GetInputMapping() *InputToGameKeyMapping {
        return this.InputMapping
}


// throws Exception
func (this *PersistentInputMapping) IsDelete(input IInput) bool {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


// throws Exception
func (this *PersistentInputMapping) IsSystemInput(input IInput) bool {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}

