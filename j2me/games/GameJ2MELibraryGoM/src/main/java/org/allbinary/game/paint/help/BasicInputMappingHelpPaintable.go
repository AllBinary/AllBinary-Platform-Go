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

type BasicInputMappingHelpPaintable struct {
    HelpPaintable
    GameInputMappingArray []*GameInputMapping
    NONE *GameKey
}


func (this *BasicInputMappingHelpPaintable) initFields() {
    
    this.NONE = GameKeyFactory_GetInstance().NONE
}


func NewBasicInputMappingHelpPaintable(gameInputMappingArray []*GameInputMapping, basicColor *BasicColor) *BasicInputMappingHelpPaintable {
    this := &BasicInputMappingHelpPaintable{}
        this.HelpPaintable = *NewHelpPaintable("Help Screen", BasicColorFactory_GetInstance().BLACK, basicColor)
        this.initFields()
        this.GameInputMappingArray = gameInputMappingArray
        this.Update(this.NONE, this.NONE)
    return this
}


func (this *BasicInputMappingHelpPaintable) Update(selectedGameKey *GameKey, selectedInput IInput) {
        var gameKeyMapping IPersistentInputMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance(); _ = gameKeyMapping
        var EQUAL string = " = "; _ = EQUAL
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var size int = len(this.GameInputMappingArray); _ = size
        var keyInfo []string = make([]string, size); _ = keyInfo
        var gameInputMapping *GameInputMapping; _ = gameInputMapping
        var gameKey *GameKey; _ = gameKey
        var list IBasicArrayList; _ = list
        for index := int(0); index < size; index++ {
            gameInputMapping = this.GameInputMappingArray[index]
            gameKey = gameInputMapping.GetGameKey()
            list = gameKeyMapping.GetInputMapping().GetMappedInput(gameKey)
            stringMaker.Delete(0, stringMaker.Length())
            keyInfo[index] = fmt.Sprint(stringMaker.Append(gameInputMapping.GetName()).Append(EQUAL).Append(this.Get(list)))
        }
        this.HelpPaintable.SetInputInfoP(keyInfo)
}


func (this *BasicInputMappingHelpPaintable) Get(keyList IBasicArrayList) string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var AND string = "and "; _ = AND
        var key IInput; _ = key
        var size int = keyList.Size(); _ = size
        for index := int(0); index < size; index++ {
            key = any(keyList.BasicArrayList_().ObjectArray[index]).(IInput)
            stringBuffer.Append(key.GetName())
            if index + 1 < keyList.Size() {
                if keyList.Size() == 2 {
                    stringBuffer.Append(commonSeps.SPACE).Append(AND)
                } else {
                    if index + 2 == keyList.Size() {
                        stringBuffer.Append(commonSeps.COMMA_SEP)
                        stringBuffer.Append(AND)
                    } else {
                        stringBuffer.Append(commonSeps.COMMA_SEP)
                    }
                }
            }
        }
        return fmt.Sprint(stringBuffer)
}

