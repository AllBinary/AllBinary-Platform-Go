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

type InputMappingHelpPaintable struct {
    HelpPaintable
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    StringUtil *StringUtil
    BasicColorFactory *BasicColorFactory
    NONE *GameKey
    EMPTY_STRING string
    GameInputMappingArray []*GameInputMapping
    KeyMappingArray []IBasicArrayList
    ActionBasicColor []*BasicColor
    InputBasicColorArray [][]*BasicColor
    SelectedBasicColor *BasicColor
    UpdateMyFontProperties *UpdateMyFontProperties
}

var InputMappingHelpPaintable_AND string = " and "
var InputMappingHelpPaintable_SEP string = ", "
var InputMappingHelpPaintable_MORE_THAN_TWO_IN_LIST_AND string = ", and "

func (this *InputMappingHelpPaintable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.BasicColorFactory = BasicColorFactory_GetInstance()
    this.NONE = GameKeyFactory_GetInstance().NONE
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
    
    this.KeyMappingArray = make([]IBasicArrayList, 0)
    this.ActionBasicColor = BasicColorUtil_GetInstance().ZERO_ARRAY
    this.InputBasicColorArray = func() [][]*BasicColor { array := make([][]*BasicColor, 0); for index := range array { array[index] = make([]*BasicColor, 0) }; return array }()
    
    this.UpdateMyFontProperties = NewUpdateMyFontProperties()
}


func NewInputMappingHelpPaintable(gameInputMappingArray []*GameInputMapping, backgroundBasicColor *BasicColor, basicColor *BasicColor) *InputMappingHelpPaintable {
    this := &InputMappingHelpPaintable{}
        this.HelpPaintable = *NewHelpPaintable("Input Mapping", backgroundBasicColor, basicColor)
        this.initFields()
        this.GameInputMappingArray = gameInputMappingArray
        this.Update(this.NONE, this.NONE)
        if backgroundBasicColor == this.BasicColorFactory.WHITE || basicColor == this.BasicColorFactory.WHITE {
            this.SelectedBasicColor = this.BasicColorFactory.RED
        } else {
            this.SelectedBasicColor = this.BasicColorFactory.WHITE
        }
    return this
}


func (this *InputMappingHelpPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        var updateMyFontProperties *UpdateMyFontProperties = NewUpdateMyFontProperties(); _ = updateMyFontProperties
        updateMyFontProperties.CharHeight = font.GetHeight()
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var inputInfo []string = this.InputInfo; _ = inputInfo
        var size int = len(inputInfo); _ = size
        updateMyFontProperties.ActionStringArray = make([]string, size)
        updateMyFontProperties.KeymappingBeginWidthArray = make([]int, size)
        updateMyFontProperties.ActionStringDeltaXArray = make([]int, size)
        updateMyFontProperties.InputDeltaXArray = func() [][]int { array := make([][]int, size); for index := range array { array[index] = make([]int, 0) }; return array }()
        updateMyFontProperties.SepArray = func() [][]string { array := make([][]string, size); for index := range array { array[index] = make([]string, 0) }; return array }()
        updateMyFontProperties.SepWidthArray = func() [][]int { array := make([][]int, size); for index := range array { array[index] = make([]int, 0) }; return array }()
        var list IBasicArrayList; _ = list
        var keyMappings string; _ = keyMappings
        var actionString string; _ = actionString
        var input IInput; _ = input
        var size2 int = 0; _ = size2
        for index := int(0); index < size; index++ {
            list = this.KeyMappingArray[index]
            keyMappings = this.Get(list)
            stringMaker.Delete(0, stringMaker.Length())
            actionString = fmt.Sprint(stringMaker.Append(inputInfo[index]).Append(this.CommonSeps.COLON).Append(this.CommonSeps.SPACE).Append(this.CommonSeps.SPACE))
            updateMyFontProperties.ActionStringArray[index] = actionString
            stringMaker.Delete(0, stringMaker.Length())
            updateMyFontProperties.KeymappingBeginWidthArray[index] = (font.StringWidth(fmt.Sprint(stringMaker.Append(updateMyFontProperties.ActionStringArray[index]).Append(keyMappings))) >> 1)
            updateMyFontProperties.ActionStringDeltaXArray[index] = font.StringWidth(actionString)
            size2 = list.Size()
            updateMyFontProperties.InputDeltaXArray[index] = make([]int, size2)
            updateMyFontProperties.SepArray[index] = make([]string, size2)
            updateMyFontProperties.SepWidthArray[index] = make([]int, size2)
            for index2 := int(0); index2 < size2; index2++ {
                input = any(list.BasicArrayList_().ObjectArray[index2]).(IInput)
                updateMyFontProperties.InputDeltaXArray[index][index2] = font.StringWidth(input.GetName())
                updateMyFontProperties.SepArray[index][index2] = this.EMPTY_STRING
                if index2 + 1 < list.Size() {
                    if list.Size() == 2 {
                        updateMyFontProperties.SepArray[index][index2] = InputMappingHelpPaintable_AND
                    } else {
                        if index2 + 2 == list.Size() {
                            updateMyFontProperties.SepArray[index][index2] = InputMappingHelpPaintable_MORE_THAN_TWO_IN_LIST_AND
                        } else {
                            updateMyFontProperties.SepArray[index][index2] = InputMappingHelpPaintable_SEP
                        }
                    }
                }
                if updateMyFontProperties.SepArray[index][index2] != this.EMPTY_STRING {
                    updateMyFontProperties.SepWidthArray[index][index2] = font.StringWidth(updateMyFontProperties.SepArray[index][index2])
                }
            }
        }
        this.UpdateMyFontProperties = updateMyFontProperties
        this.HelpPaintable.UpdateMeasurement(graphics)
}


func (this *InputMappingHelpPaintable) Update(selectedGameKey *GameKey, selectedInput IInput) {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(CommonLabels_GetInstance().START_LABEL).Append("selected GameKey: ").Append(this.StringUtil.ToString(selectedGameKey)).Append(" Input: ").Append(this.StringUtil.ToString(selectedInput))), this, this.CommonStrings.UPDATE)
        var gameKeyMapping IPersistentInputMapping = PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance(); _ = gameKeyMapping
        var size int = len(this.GameInputMappingArray); _ = size
        var keyInfo []string = make([]string, size); _ = keyInfo
        var keyMappingArray []IBasicArrayList = make([]IBasicArrayList, size); _ = keyMappingArray
        var actionBasicColor []*BasicColor = make([]*BasicColor, size); _ = actionBasicColor
        var inputBasicColorArray [][]*BasicColor = make([][]*BasicColor, size); _ = inputBasicColorArray
        var gameInputMapping *GameInputMapping; _ = gameInputMapping
        var gameKey *GameKey; _ = gameKey
        var list IBasicArrayList; _ = list
        for index := int(0); index < size; index++ {
            gameInputMapping = this.GameInputMappingArray[index]
            gameKey = gameInputMapping.GetGameKey()
            list = gameKeyMapping.GetInputMapping().GetMappedInput(gameKey)
            var size2 int = list.Size(); _ = size2
            inputBasicColorArray[index] = make([]*BasicColor, size2)
            for index2 := int(0); index2 < size2; index2++ {
                inputBasicColorArray[index][index2] = this.BasicColor
            }
            if gameKey == selectedGameKey {
                stringMaker.Delete(0, stringMaker.Length())
                this.LogUtil.PutF(fmt.Sprint(stringMaker.Append("Found: selected GameKey: ").Append(this.StringUtil.ToString(selectedGameKey))), this, this.CommonStrings.UPDATE)
                actionBasicColor[index] = this.SelectedBasicColor
                var indexOfSelectedInput int = list.IndexOf(selectedInput); _ = indexOfSelectedInput
                if indexOfSelectedInput >= 0 {
                    stringMaker.Delete(0, stringMaker.Length())
                    this.LogUtil.PutF(fmt.Sprint(stringMaker.Append("Found: selected Input: ").Append(this.StringUtil.ToString(selectedInput))), this, this.CommonStrings.UPDATE)
                    inputBasicColorArray[index][indexOfSelectedInput] = this.SelectedBasicColor
                }
            } else {
                actionBasicColor[index] = this.BasicColor
            }
            keyInfo[index] = gameInputMapping.GetName()
            keyMappingArray[index] = list
        }
        this.KeyMappingArray = keyMappingArray
        this.ActionBasicColor = actionBasicColor
        this.InputBasicColorArray = inputBasicColorArray
        this.HelpPaintable.SetInputInfoP(keyInfo)
        this.MyFontProcessor = this.UpdateMyFontProcessor
}


func (this *InputMappingHelpPaintable) Get(keyList IBasicArrayList) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var key IInput; _ = key
        var size int = keyList.Size(); _ = size
        for index := int(0); index < size; index++ {
            key = any(keyList.BasicArrayList_().ObjectArray[index]).(IInput)
            stringBuffer.Append(key.GetName())
            if index + 1 < keyList.Size() {
                if keyList.Size() == 2 {
                    stringBuffer.Append(InputMappingHelpPaintable_AND)
                } else {
                    if index + 2 == keyList.Size() {
                        stringBuffer.Append(InputMappingHelpPaintable_MORE_THAN_TWO_IN_LIST_AND)
                    } else {
                        stringBuffer.Append(InputMappingHelpPaintable_SEP)
                    }
                }
            }
        }
        return fmt.Sprint(stringBuffer)
}


func (this *InputMappingHelpPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var halfWidth int = this.DisplayInfo.GetLastHalfWidth(); _ = halfWidth
        graphics.SetColor(this.BasicColor.IntValue())
        graphics.DrawString(this.Title, halfWidth - this.TitleBeginWidth, this.UpdateMyFontProperties.CharHeight, this.Anchor)
        var inputInfo []string = this.InputInfo; _ = inputInfo
        var size int = len(inputInfo); _ = size
        var y int = 0; _ = y
        var deltaX int = 0; _ = deltaX
        var size2 int = 0; _ = size2
        var actionString string; _ = actionString
        var input IInput; _ = input
        var list IBasicArrayList; _ = list
        var sep string; _ = sep
        var beginWidth int; _ = beginWidth
        for index := int(0); index < size; index++ {
            y = (index + 3) * this.UpdateMyFontProperties.CharHeight
            deltaX = 0
            list = this.KeyMappingArray[index]
            size2 = list.Size()
            beginWidth = this.UpdateMyFontProperties.KeymappingBeginWidthArray[index]
            graphics.SetColor(int(this.ActionBasicColor[index]))
            actionString = this.UpdateMyFontProperties.ActionStringArray[index]
            graphics.DrawString(actionString, halfWidth - beginWidth + deltaX, y, this.Anchor)
            deltaX += this.UpdateMyFontProperties.ActionStringDeltaXArray[index]
            for index2 := int(0); index2 < size2; index2++ {
                input = any(list.BasicArrayList_().ObjectArray[index2]).(IInput)
                graphics.SetColor(int(this.InputBasicColorArray[index][index2]))
                graphics.DrawString(input.GetName(), halfWidth - beginWidth + deltaX, y, this.Anchor)
                deltaX += this.UpdateMyFontProperties.InputDeltaXArray[index][index2]
                sep = this.UpdateMyFontProperties.SepArray[index][index2]
                graphics.SetColor(this.BasicColor.IntValue())
                graphics.DrawString(sep, halfWidth - beginWidth + deltaX, y, this.Anchor)
                deltaX += this.UpdateMyFontProperties.SepWidthArray[index][index2]
            }
        }
}

type UpdateMyFontProperties struct {
    ActionStringArray []string
    KeymappingBeginWidthArray []int
    ActionStringDeltaXArray []int
    InputDeltaXArray [][]int
    SepArray [][]string
    SepWidthArray [][]int
    CharHeight int
}


func (this *UpdateMyFontProperties) initFields() {
        this.ActionStringArray = StringUtil_GetInstance().GetArrayInstance()
        this.KeymappingBeginWidthArray = NullUtil_GetInstance().NULL_INT_ARRAY
        this.ActionStringDeltaXArray = NullUtil_GetInstance().NULL_INT_ARRAY
        this.InputDeltaXArray = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY
        this.SepArray = func() [][]string { array := make([][]string, 0); for index := range array { array[index] = make([]string, 0) }; return array }()
        this.SepWidthArray = NullUtil_GetInstance().NULL_INT_ARRAY_ARRAY
        
}

func NewUpdateMyFontProperties() *UpdateMyFontProperties {
    this := &UpdateMyFontProperties{}
    this.initFields()
    return this
}

