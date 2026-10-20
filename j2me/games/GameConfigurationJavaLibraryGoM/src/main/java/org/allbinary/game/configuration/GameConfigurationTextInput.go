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

type GameConfigurationTextInput struct {
    Label string
    Text string
}

var GameConfigurationTextInput_Hashtable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable()
var GameConfigurationTextInput_SIMULATOR_IP *GameConfigurationTextInput = NewGameConfigurationTextInput("IP: ", "192.168.1.3")
var GameConfigurationTextInput_SIMULATOR_PORT *GameConfigurationTextInput = NewGameConfigurationTextInput("Port: ", "8010")

func (this *GameConfigurationTextInput) initFields() {
    this.Label = StringUtil_GetInstance().EMPTY_STRING
    this.Text = StringUtil_GetInstance().EMPTY_STRING
}


func NewGameConfigurationTextInput(label string, text string) *GameConfigurationTextInput {
    this := &GameConfigurationTextInput{}
    this.initFields()
        this.SetText(text)
        this.SetLabel(label)
        GameConfigurationTextInput_Hashtable.Put(this.GetLabel(), this)
    return this
}


func GameConfigurationTextInput_Update(textField ITextField) {
        var anyType any = any(textField.GetLabel()).(any); _ = anyType
        var gameConfigurationTextInput *GameConfigurationTextInput = any(GameConfigurationTextInput_Hashtable.Get(anyType)).(*GameConfigurationTextInput); _ = gameConfigurationTextInput
        gameConfigurationTextInput.SetText(textField.GetString())
}


func (this *GameConfigurationTextInput) SetText(text string) {
        this.Text = text
}


func (this *GameConfigurationTextInput) GetText() string {
        return this.Text
}


func (this *GameConfigurationTextInput) SetLabel(label string) {
        this.Label = label
}


func (this *GameConfigurationTextInput) GetLabel() string {
        return this.Label
}


func GameConfigurationTextInput_GetHashtable() *ABHashtable[any, any] {
        return GameConfigurationTextInput_Hashtable
}

