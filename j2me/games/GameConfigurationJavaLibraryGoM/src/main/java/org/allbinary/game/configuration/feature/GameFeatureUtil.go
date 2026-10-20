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

type GameFeatureUtil struct {
    LogUtil *LogUtil
    TsUtil *TsUtil
    EnumerationUtil *EnumerationUtil
    ON_GAME_FEATURE_CHANGE string
    GAME_FEATURE_CHANGED string
}

var GameFeatureUtil_Instance *GameFeatureUtil = NewGameFeatureUtil()

func (this *GameFeatureUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TsUtil = TsUtil_GetInstance()
    this.EnumerationUtil = EnumerationUtil_GetInstance()
    this.ON_GAME_FEATURE_CHANGE = "onGameFeatureChange"
    this.GAME_FEATURE_CHANGED = "Game Feature Changed: "
}

func NewGameFeatureUtil() *GameFeatureUtil {
    this := &GameFeatureUtil{}
    this.initFields()
    return this
}


func GameFeatureUtil_GetInstance() *GameFeatureUtil {
        return GameFeatureUtil_Instance
}


// throws Exception
func (this *GameFeatureUtil) SetDefault(choiceGroup IChoiceGroup) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var METHOD_NAME string = "setDefault"; _ = METHOD_NAME
        var SELECTED_ARRAY_RETURN string = "selectedArray_return["; _ = SELECTED_ARRAY_RETURN
        var SELECTED_SEP string = "] = "; _ = SELECTED_SEP
        var selectedArray_return []bool = make([]bool, choiceGroup.Size()); _ = selectedArray_return
        var total int = choiceGroup.GetSelectedFlags(selectedArray_return); _ = total
        stringBuffer.Append("Multiple Total Choices: ")
        stringBuffer.Appendint(total)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
        var features *Features = Features_GetInstance(); _ = features
        for index := int(0); index < len(selectedArray_return); index++ {
            var isSelected bool = selectedArray_return[index]; _ = isSelected
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(SELECTED_ARRAY_RETURN)
            stringBuffer.Appendint(index)
            stringBuffer.Append(SELECTED_SEP)
            stringBuffer.Appendboolean(isSelected)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
            var selectedChoiceLabel string = choiceGroup.GetString(index); _ = selectedChoiceLabel
            var gameFeature IFeature = Feature_GetInstance(selectedChoiceLabel); _ = gameFeature
            if features.IsDefault(gameFeature) {
                features.Add(gameFeature)
                choiceGroup.SetSelectedIndex(index, true)
            } else {
                features.Remove(gameFeature)
                choiceGroup.SetSelectedIndex(index, false)
            }
        }
}


// throws Exception
func (this *GameFeatureUtil) UpdateMultiple(choiceGroup IChoiceGroup) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var METHOD_NAME string = "updateMultiple"; _ = METHOD_NAME
        var SELECTED_ARRAY_RETURN string = "selectedArray_return["; _ = SELECTED_ARRAY_RETURN
        var SELECTED_SEP string = "] = "; _ = SELECTED_SEP
        var selectedArray_return []bool = make([]bool, choiceGroup.Size()); _ = selectedArray_return
        var total int = choiceGroup.GetSelectedFlags(selectedArray_return); _ = total
        stringBuffer.Append("Multiple Total Choices: ")
        stringBuffer.Appendint(total)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
        var features *Features = Features_GetInstance(); _ = features
        for index := int(0); index < len(selectedArray_return); index++ {
            var isSelected bool = selectedArray_return[index]; _ = isSelected
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(SELECTED_ARRAY_RETURN)
            stringBuffer.Appendint(index)
            stringBuffer.Append(SELECTED_SEP)
            stringBuffer.Appendboolean(isSelected)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
            var selectedChoiceLabel string = choiceGroup.GetString(index); _ = selectedChoiceLabel
            var gameFeature IFeature = Feature_GetInstance(selectedChoiceLabel); _ = gameFeature
            if !isSelected && features.IsFeature(gameFeature) {
                features.Remove(gameFeature)
            } else if isSelected && !features.IsFeature(gameFeature) {
                features.Add(gameFeature)
            }
        }
}


// throws Exception
func (this *GameFeatureUtil) UpdateExclusiveForChoiceGroup(choiceGroup IChoiceGroup) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var METHOD_NAME string = "updateExclusive"; _ = METHOD_NAME
        var SELECTED_ARRAY_RETURN string = "selectedArray_return["; _ = SELECTED_ARRAY_RETURN
        var SELECTED_SEP string = "] = "; _ = SELECTED_SEP
        var selectedArray_return []bool = make([]bool, choiceGroup.Size()); _ = selectedArray_return
        var total int = choiceGroup.GetSelectedFlags(selectedArray_return); _ = total
        stringBuffer.Append("Exclusive Total Choices: 1==")
        stringBuffer.Appendint(total)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
        for index := int(0); index < len(selectedArray_return); index++ {
            var isSelected bool = selectedArray_return[index]; _ = isSelected
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append(SELECTED_ARRAY_RETURN)
            stringBuffer.Appendint(index)
            stringBuffer.Append(SELECTED_SEP)
            stringBuffer.Appendboolean(isSelected)
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, METHOD_NAME)
            if isSelected {
                var selectedChoiceLabel string = choiceGroup.GetString(index); _ = selectedChoiceLabel
                this.UpdateExclusive(selectedChoiceLabel)
            }
        }
}


func (this *GameFeatureUtil) IsExclusive(itemLabel string) bool {
        var enumeration Enumeration = GameFeatureChoiceGroups_GetExclusiveInstance().Get().Keys(); _ = enumeration
        var name string; _ = name
        for this.EnumerationUtil.HasMoreElements(enumeration) {
            name = any(this.EnumerationUtil.NextElement(enumeration)).(string)
            if this.TsUtil.CompareTo(itemLabel, name) == 0 {
                return true
            }
        }
        return false
}


// throws Exception
func (this *GameFeatureUtil) UpdateExclusive(selectedChoiceLabel string) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().ITEM_LABEL).Append(selectedChoiceLabel)), this, "updateExclusive")
        var gameFeature IFeature = Feature_GetInstance(selectedChoiceLabel); _ = gameFeature
        var hashtable *ABHashtable[any, any] = GameFeatureChoiceGroups_GetExclusiveInstance().Get(); _ = hashtable
        var objectArray []any = HashtableUtil_GetInstance().GetKeysAsArray(hashtable); _ = objectArray
        var basicArrayList IBasicArrayList; _ = basicArrayList
        var anyType any; _ = anyType
        var size int = len(objectArray); _ = size
        for index := int(0); index < size; index++ {
            anyType = any(objectArray[index]).(any)
            basicArrayList = any(hashtable.Get(anyType)).(IBasicArrayList)
            this.UpdateExclusiveList(gameFeature, basicArrayList)
        }
}


// throws Exception
func (this *GameFeatureUtil) UpdateExclusiveList(gameFeature IFeature, list IBasicArrayList) {
        if list.Contains(gameFeature) {
            var features *Features = Features_GetInstance(); _ = features
            var addIndex int = list.IndexOf(gameFeature); _ = addIndex
            var size int = list.Size(); _ = size
            for index := int(0); index < size; index++ {
                if addIndex != index {
                    features.Remove(any(list.BasicArrayList_().ObjectArray[index]).(IFeature))
                } else {
                    features.Add(any(list.BasicArrayList_().ObjectArray[index]).(IFeature))
                }
            }
        }
}

