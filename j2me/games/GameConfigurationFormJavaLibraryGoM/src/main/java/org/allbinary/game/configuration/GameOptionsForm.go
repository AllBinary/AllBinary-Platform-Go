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

type GameOptionsForm struct {
    CommandForm
    MeUtil *MEUtil
}


func (this *GameOptionsForm) initFields() {
    this.MeUtil = MEUtil_GetInstance()
}


func NewGameOptionsForm(commandListener CommandListener, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *GameOptionsForm {
    this := &GameOptionsForm{}
        this.CommandForm = *NewCommandForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.AddConfiguration()
        var gameFeatureFormUtil *GameFeatureFormUtil = GameFeatureFormUtil_GetInstance(); _ = gameFeatureFormUtil
        gameFeatureFormUtil.AddChoiceGroup(this, GameFeatureChoiceGroups_GetExclusiveInstance().Get(), Choice_EXCLUSIVE)
        gameFeatureFormUtil.AddChoiceGroup(this, GameFeatureChoiceGroups_GetMultipleInstance().Get(), Choice_MULTIPLE)
        this.InitCommands(commandListener)
        this.SetItemStateListener(NewGameFeatureItemStateListener(this))
        this.AddTextFieldsIfSimulated()
    return this
}


// throws Exception
func (this *GameOptionsForm) CloseClientInformation(abeClientInformation AbeClientInformationInterface) {
        this.CommandForm.Close()
        this.Save(abeClientInformation)
}


func (this *GameOptionsForm) AddTextFieldsIfSimulated() {
        var key string = OrientationData_GetInstance().ORIENTATION_SENSOR_INPUT; _ = key
        var hashtable *ABHashtable[any, any] = GameFeatureChoiceGroups_GetExclusiveInstance().Get(); _ = hashtable
        if hashtable != nil {
            var listCanBeNull any = hashtable.Get(key); _ = listCanBeNull
            if listCanBeNull != nil {
                var list IBasicArrayList = any(listCanBeNull).(IBasicArrayList); _ = list
                if list.Contains(SensorFeatureFactory_GetInstance().SIMULATED_ORIENTATION_SENSORS) {
                    this.AddTextFields()
                }
            }
        }
}


func (this *GameOptionsForm) AddTextFields() {
        var hashtable *ABHashtable[any, any] = GameConfigurationTextInput_GetHashtable(); _ = hashtable
        var size int = hashtable.Size(); _ = size
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
        var gameConfigurationTextInput *GameConfigurationTextInput; _ = gameConfigurationTextInput
        var textField ITextField; _ = textField
        for index := int(0); index < size; index++ {
            gameConfigurationTextInput = any(hashtable.Get(any(objectArray[index]).(any))).(*GameConfigurationTextInput)
            textField = NewTextField(gameConfigurationTextInput.GetLabel(), gameConfigurationTextInput.GetText(), 30, TextField_ANY)
            this.MeUtil.AppendItem(this, textField)
        }
}


func (this *GameOptionsForm) AddConfiguration() {
        var METHOD_NAME string = "addConfiguration"; _ = METHOD_NAME
        var NAME string = "Name: "; _ = NAME
        var list IBasicArrayList = GameConfigurationSingleton_GetInstance().GetOptionsBasicArrayList(); _ = list
        var GAUGE_CHANGE ICommand = MyCommandsFactory_GetInstance().GAUGE_CHANGE; _ = GAUGE_CHANGE
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var size int = list.Size(); _ = size
        var gameConfiguration IGameConfiguration; _ = gameConfiguration
        var gauge *GameConfigurationGauge; _ = gauge
        for index := int(0); index < size; index++ {
            gameConfiguration = any(list.BasicArrayList_().ObjectArray[index]).(IGameConfiguration)
            stringMaker.Delete(0, stringMaker.Length())
            this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(NAME).Append(fmt.Sprint(gameConfiguration))), this, METHOD_NAME)
            gauge = NewGameConfigurationGauge(gameConfiguration)
            gauge.SetDefaultCommand(GAUGE_CHANGE)
            gauge.SetItemCommandListener(NewGameFeatureItemCommandListener(this))
            this.MeUtil.AppendItem(this, gauge)
        }
}


func (this *GameOptionsForm) InitCommands(cmdListener CommandListener) {
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        this.RemoveAllCommands()
        this.AddCommand(gameCommandsFactory.CLOSE_OPTIONS)
        this.AddCommand(gameCommandsFactory.DEFAULT_OPTIONS)
        this.SetCommandListener(cmdListener)
}


// throws Exception
func (this *GameOptionsForm) Save(abeClientInformation AbeClientInformationInterface) {
        var size int = this.Size(); _ = size
        var item IItem; _ = item
        for index := int(0); index < size; index++ {
            item = this.Get(index)
            if func() bool { _, ok := any(item).(*GameConfigurationGauge); return ok }() {
                GameConfigurationUtil_GetInstance().Update(any(item).(*GameConfigurationGauge))
            } else if func() bool { _, ok := any(item).(ITextField); return ok }() {
                GameConfigurationTextInput_Update(any(item).(ITextField))
            }
        }
        GameConfigurationUtil_GetInstance().UpdateCompetitionValue()
        var hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable(); _ = hashtable
        var SCALE IGameConfiguration = GameConfigurationCentral_GetInstance().SCALE; _ = SCALE
        hashtable.Put(SCALE.GetName(), fmt.Sprint(SCALE.GetValue()))
        var keyValuePersistance IKeyValuePersistance = GameConfigurationPersistanceSingleton_GetInstance(); _ = keyValuePersistance
        keyValuePersistance.Clear()
        keyValuePersistance.LoadAll(abeClientInformation)
        var list IBasicArrayList = keyValuePersistance.GetIds(); _ = list
        keyValuePersistance.Save(abeClientInformation, hashtable)
        var size2 int = list.Size(); _ = size2
        var integer int; _ = integer
        for index := int(0); index < size2; index++ {
            integer = any(list.BasicArrayList_().ObjectArray[index]).(int)
            keyValuePersistance.Delete(abeClientInformation, int(integer))
        }
}

