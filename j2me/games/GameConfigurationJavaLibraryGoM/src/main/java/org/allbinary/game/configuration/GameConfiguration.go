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

type GameConfiguration struct {
    Name string
    DefaultValue int
    Value int
    MinValue int
    MaxValue int
    Modifiable bool
}

// The Java class GameConfiguration has subclasses so its variables are this interface
type IGameConfiguration interface {
    GameConfiguration_() *GameConfiguration
    SetValue(value int)
    GetValue() int
    SetMaxValue(maxValue int)
    GetMaxValue() int
    SetMinValue(minValue int)
    GetMinValue() int
    SetName(name string)
    GetName() string
    SetDefaultValue(defaultValue int)
    GetDefaultValue() int
    SetDefault()
    ToString() string
    SetModifiable(modifiable bool)
    IsModifiable() bool
}

func (this *GameConfiguration) GameConfiguration_() *GameConfiguration {
    return this
}

var GameConfiguration_MIN_VALUE int = int(Integer_MIN_VALUE)

func (this *GameConfiguration) initFields() {
    this.Name = StringUtil_GetInstance().EMPTY_STRING
    this.DefaultValue = SmallIntegerSingletonFactory_GetInstance().GetAt(0)
    this.Value = GameConfiguration_MIN_VALUE
    this.MinValue = this.DefaultValue
    this.MaxValue = this.DefaultValue
    this.Modifiable = BooleanFactory_GetInstance().TRUE
}


func NewGameConfiguration(name string, defaultValue int, minValue int, maxValue int) *GameConfiguration {
    this := &GameConfiguration{}
    this.initFields()
        this.SetName(name)
        this.SetDefaultValue(defaultValue)
        this.Value = this.GetDefaultValue()
        this.SetMinValue(minValue)
        this.SetMaxValue(maxValue)
        GameConfigurationSingleton_GetInstance().GetHashtable().Put(this.GetName(), this)
    return this
}


// throws Exception
func (this *GameConfiguration) SetValue(value int) {
        if this.Value != value {
            var oldValue int = this.Value; _ = oldValue
            this.Value = value
            GameFeatureEventHandler_GetInstance().FireEvent(NewGameFeatureEvent(this, fmt.Sprint(NewStringMaker().Append(this.Name).Append(" value: from: ").Appendint(int(oldValue)).Append(" to ").Appendint(int(this.Value)))))
        }
}


func (this *GameConfiguration) GetValue() int {
        return this.Value
}


func (this *GameConfiguration) SetMaxValue(maxValue int) {
        this.MaxValue = maxValue
}


func (this *GameConfiguration) GetMaxValue() int {
        return this.MaxValue
}


func (this *GameConfiguration) SetMinValue(minValue int) {
        this.MinValue = minValue
}


func (this *GameConfiguration) GetMinValue() int {
        return this.MinValue
}


func (this *GameConfiguration) SetName(name string) {
        this.Name = name
}


func (this *GameConfiguration) GetName() string {
        return this.Name
}


func (this *GameConfiguration) SetDefaultValue(defaultValue int) {
        this.DefaultValue = defaultValue
}


func (this *GameConfiguration) GetDefaultValue() int {
        return this.DefaultValue
}


// throws Exception
func (this *GameConfiguration) SetDefault() {
        this.SetValue(this.GetDefaultValue())
}


func (this *GameConfiguration) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Name: ")
        stringBuffer.Append(this.Name)
        stringBuffer.Append(" Min: ")
        stringBuffer.Appendint(int(this.GetMinValue()))
        stringBuffer.Append(" Max: ")
        stringBuffer.Appendint(int(this.GetMaxValue()))
        stringBuffer.Append(" Value: ")
        stringBuffer.Appendint(int(this.GetValue()))
        stringBuffer.Append(" Default: ")
        stringBuffer.Appendint(int(this.GetDefaultValue()))
        stringBuffer.Append(" Modifiable: ")
        stringBuffer.Appendboolean(this.IsModifiable())
        return fmt.Sprint(stringBuffer)
}

func (this *GameConfiguration) String() string {
    return this.ToString()
}


func (this *GameConfiguration) SetModifiable(modifiable bool) {
        this.Modifiable = modifiable
}


func (this *GameConfiguration) IsModifiable() bool {
        return this.Modifiable
}

