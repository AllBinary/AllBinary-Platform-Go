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

type GameDifficultyOptions struct {
    CommandForm
    MeUtil *MEUtil
    List IBasicArrayList
}


func (this *GameDifficultyOptions) initFields() {
    this.MeUtil = MEUtil_GetInstance()
    
}


func NewGameDifficultyOptions(commandListener CommandListener, title string, list IBasicArrayList, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *GameDifficultyOptions {
    this := &GameDifficultyOptions{}
        this.CommandForm = *NewCommandForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.List = list
        this.Update()
        this.InitCommands(commandListener)
        var gameDifficultyFactory *GameDifficultyFactory = GameDifficultyFactory_GetInstance(); _ = gameDifficultyFactory
        this.SetSelectedId(gameDifficultyFactory.GetLevel())
    return this
}


func (this *GameDifficultyOptions) Open() {
        this.CommandForm.Open()
}


// throws Exception
func (this *GameDifficultyOptions) Close() {
        this.CommandForm.Close()
        this.Save()
}


// throws Exception
func (this *GameDifficultyOptions) Save() {
        var gameDifficultyFactory *GameDifficultyFactory = GameDifficultyFactory_GetInstance(); _ = gameDifficultyFactory
        gameDifficultyFactory.SetLevel(this.GetSelectedId())
}


func (this *GameDifficultyOptions) InitCommands(cmdListener CommandListener) {
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        this.RemoveAllCommands()
        this.AddCommand(gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS)
        this.SetCommandListener(cmdListener)
}


// throws Exception
func (this *GameDifficultyOptions) Update() {
        this.DeleteAll()
        this.Add(this.List, "Difficulty", Choice_EXCLUSIVE)
        this.CommandForm.Update()
}


func (this *GameDifficultyOptions) Add(list IBasicArrayList, name string, option int) {
        var choiceGroup IChoiceGroup = this.GetChoiceGroup(list, name, option); _ = choiceGroup
        if list.Size() > 0 {
            choiceGroup.SetSelectedIndex(0, true)
        }
        this.MeUtil.AppendItem(this, choiceGroup)
}


func (this *GameDifficultyOptions) GetChoiceGroup(list IBasicArrayList, name string, option int) IChoiceGroup {
        var METHOD_NAME string = "addChoiceGroup"; _ = METHOD_NAME
        var NAME string = "Adding Choice: "; _ = NAME
        var choiceGroup IChoiceGroup = NewChoiceGroup(name, option, StringUtil_GetInstance().GetArrayInstance(), NullImage_NULL_IMAGE_ARRAY); _ = choiceGroup
        var size int = list.Size(); _ = size
        var anyType any; _ = anyType
        for index := int(0); index < size; index++ {
            anyType = list.BasicArrayList_().ObjectArray[index]
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(NAME).Append(fmt.Sprint(anyType))), this, METHOD_NAME)
            choiceGroup.Append(fmt.Sprint(anyType), NullImage_NULL_IMAGE)
        }
        return choiceGroup
}


func (this *GameDifficultyOptions) SetSelectedId(id int) {
        var choiceGroup IChoiceGroup = any(this.Get(0)).(IChoiceGroup); _ = choiceGroup
        choiceGroup.SetSelectedIndex(id, true)
}


func (this *GameDifficultyOptions) GetSelectedId() int {
        var choiceGroup IChoiceGroup = any(this.Get(0)).(IChoiceGroup); _ = choiceGroup
        return choiceGroup.GetSelectedIndex()
}


func (this *GameDifficultyOptions) GetSelectedText() int {
        var choiceGroup IChoiceGroup = any(this.Get(0)).(IChoiceGroup); _ = choiceGroup
        var value int = int(fmt.Sprint(choiceGroup.GetString(choiceGroup.GetSelectedIndex()))); _ = value
        return value
}

