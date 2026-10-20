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

type GameFeatureFormUtil struct {
    LogUtil *LogUtil
    MeUtil *MEUtil
}

var GameFeatureFormUtil_Instance *GameFeatureFormUtil = NewGameFeatureFormUtil()

func (this *GameFeatureFormUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.MeUtil = MEUtil_GetInstance()
}

func NewGameFeatureFormUtil() *GameFeatureFormUtil {
    this := &GameFeatureFormUtil{}
    this.initFields()
    return this
}


func GameFeatureFormUtil_GetInstance() *GameFeatureFormUtil {
        return GameFeatureFormUtil_Instance
}


func (this *GameFeatureFormUtil) GetChoiceGroup(hashtable *ABHashtable[any, any], name string, option int) IChoiceGroup {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var choiceGroup IChoiceGroup = NewChoiceGroup(name, option, StringUtil_GetInstance().GetArrayInstance(), NullImage_NULL_IMAGE_ARRAY); _ = choiceGroup
        var list IBasicArrayList = any(hashtable.Get(any(name).(any))).(IBasicArrayList); _ = list
        var features *Features = Features_GetInstance(); _ = features
        var ADD_CHOICE string = ": Adding Choice: "; _ = ADD_CHOICE
        var GET_CHOICE_GROUP string = "getChoiceGroup"; _ = GET_CHOICE_GROUP
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var gameFeature IFeature = any(list.BasicArrayList_().ObjectArray[index]).(IFeature); _ = gameFeature
            stringMaker.Delete(0, stringMaker.Length())
            this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(name).Append(ADD_CHOICE).Append(fmt.Sprint(gameFeature))), this, GET_CHOICE_GROUP)
            choiceGroup.Append(fmt.Sprint(gameFeature), NullImage_NULL_IMAGE)
            if features.IsFeature(gameFeature) {
                choiceGroup.SetSelectedIndex(index, true)
            }
        }
        return choiceGroup
}


func (this *GameFeatureFormUtil) AddChoiceGroup(form ICommandForm, hashtable *ABHashtable[any, any], option int) {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var ADDING_CHOICE_GROUP string = "Adding Choice Group: "; _ = ADDING_CHOICE_GROUP
        var ADD_CHOICE_GROUP string = "addChoiceGroup"; _ = ADD_CHOICE_GROUP
        var size int = hashtable.Size(); _ = size
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
        var name string; _ = name
        for index := int(0); index < size; index++ {
            name = any(objectArray[index]).(string)
            stringMaker.Delete(0, stringMaker.Length())
            this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(ADDING_CHOICE_GROUP).Append(name)), this, ADD_CHOICE_GROUP)
            this.MeUtil.AppendItem(form, this.GetChoiceGroup(hashtable, name, option))
        }
}


// throws Exception
func (this *GameFeatureFormUtil) SetDefault(form ICommandForm) {
        var size int = form.Size(); _ = size
        for index := int(0); index < size; index++ {
            var item IItem = form.Get(index); _ = item
            if func() bool { _, ok := any(item).(*GameConfigurationGauge); return ok }() {
                GameConfigurationUtil_GetInstance().SetDefault(any(item).(*GameConfigurationGauge))
            } else if func() bool { _, ok := any(item).(IChoiceGroup); return ok }() {
                GameFeatureUtil_GetInstance().SetDefault(any(item).(IChoiceGroup))
            }
        }
}

