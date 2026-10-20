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

type LoadGameForm struct {
    CommandForm
    MeUtil *MEUtil
    AreChoices bool
}


func (this *LoadGameForm) initFields() {
    this.MeUtil = MEUtil_GetInstance()
    
}


func NewLoadGameForm(commandListener CommandListener, title string, backgrounBasicColor *BasicColor, foregroundBasicColor *BasicColor) *LoadGameForm {
    this := &LoadGameForm{}
        this.CommandForm = *NewCommandForm(commandListener, title, backgrounBasicColor, foregroundBasicColor)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.Update()
        this.InitCommands(commandListener)
    return this
}


func (this *LoadGameForm) InitCommands(cmdListener CommandListener) {
        var gameCommandsFactory *GameCommandsFactory = GameCommandsFactory_GetInstance(); _ = gameCommandsFactory
        this.RemoveAllCommands()
        this.AddCommand(gameCommandsFactory.LOAD_FILE)
        this.AddCommand(gameCommandsFactory.DELETE_FILE)
        this.AddCommand(gameCommandsFactory.CLOSE_AND_SHOW_GAME_CANVAS)
        this.SetCommandListener(cmdListener)
}


// throws Exception
func (this *LoadGameForm) Update() {
        var list IBasicArrayList = GamePersistanceSingleton_GetInstance().GetIds(); _ = list
        this.DeleteAll()
        if list.Size() > 0 {
            this.Add(list, "Select A Saved Game", Choice_EXCLUSIVE)
            this.AreChoices = true
        } else {
            this.MeUtil.AppendItem(this, NewStringItem("No Saved Games", StringUtil_GetInstance().EMPTY_STRING, Item_PLAIN))
            this.AreChoices = false
        }
        this.CommandForm.Update()
}


func (this *LoadGameForm) Add(list IBasicArrayList, name string, option int) {
        var choiceGroup IChoiceGroup = this.GetChoiceGroup(list, name, option); _ = choiceGroup
        if list.Size() > 0 {
            choiceGroup.SetSelectedIndex(0, true)
        }
        this.MeUtil.AppendItem(this, choiceGroup)
}


func (this *LoadGameForm) GetChoiceGroup(list IBasicArrayList, name string, option int) IChoiceGroup {
        var METHOD_NAME string = "getChoiceGroup"; _ = METHOD_NAME
        var ADDING_CHOICE string = "Adding Choice: "; _ = ADDING_CHOICE
        var choiceGroup IChoiceGroup = NewChoiceGroup(name, option, StringUtil_GetInstance().GetArrayInstance(), NullImage_NULL_IMAGE_ARRAY); _ = choiceGroup
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = list.BasicArrayList_().ObjectArray[index]; _ = anyType
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(ADDING_CHOICE).Append(fmt.Sprint(anyType))), this, METHOD_NAME)
            choiceGroup.Append(fmt.Sprint(anyType), NullImage_NULL_IMAGE)
        }
        return choiceGroup
}


func (this *LoadGameForm) GetSelectedId() int {
        if this.AreChoices {
            var choiceGroup IChoiceGroup = any(this.Get(0)).(IChoiceGroup); _ = choiceGroup
            return choiceGroup.GetSelectedIndex()
        } else {
            return -1
        }
}


func (this *LoadGameForm) GetSelectedText() int {
        if this.AreChoices {
            var choiceGroup IChoiceGroup = any(this.Get(0)).(IChoiceGroup); _ = choiceGroup
            var value int = int(fmt.Sprint(choiceGroup.GetString(choiceGroup.GetSelectedIndex()))); _ = value
            return value
        } else {
            return -1
        }
}

