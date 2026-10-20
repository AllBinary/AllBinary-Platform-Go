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

type GameSkill struct {
    GameSkillType *GameSkillType
    Extra string
    Time int
    Vector IBasicArrayList
    GAME_SKILL_LABEL string
    EXTRA_LABEL string
    TIME_LABEL string
}

// The Java class GameSkill has subclasses so its variables are this interface
type IGameSkill interface {
    GameSkill_() *GameSkill
    GetGameSkillType() *GameSkillType
    SetGameSkillType(gameSkillType *GameSkillType)
    GetExtra() string
    SetExtra(extra string)
    GetTime() int
    SetTime(time int)
    AddProperty(gameSkillPropertyInterface GameSkillPropertyInterface)
    GetProperties() IBasicArrayList
    ToString() string
}

func (this *GameSkill) GameSkill_() *GameSkill {
    return this
}

var GameSkill_HashMap *HashMap = StdUtil_GetInstance().CreateHashMap()

func (this *GameSkill) initFields() {
    
    
    
    
    this.GAME_SKILL_LABEL = "GameSkill: "
    this.EXTRA_LABEL = "Extra: "
    this.TIME_LABEL = "Time: "
}


func NewGameSkill(gameSkillType *GameSkillType, extra string, time int) *GameSkill {
    this := &GameSkill{}
    this.initFields()
        this.SetGameSkillType(gameSkillType)
        this.SetExtra(extra)
        this.SetTime(time)
        this.Vector = NewBasicArrayListD()
        GameSkill_HashMap.Put(fmt.Sprint(NewStringMaker().Append(this.GetGameSkillType().GetName()).Append(CommonSeps_GetInstance().SPACE).Append(extra)), this)
    return this
}


func (this *GameSkill) GetGameSkillType() *GameSkillType {
        return this.GameSkillType
}


func (this *GameSkill) SetGameSkillType(gameSkillType *GameSkillType) {
        this.GameSkillType = gameSkillType
}


func (this *GameSkill) GetExtra() string {
        return this.Extra
}


func (this *GameSkill) SetExtra(extra string) {
        this.Extra = extra
}


func (this *GameSkill) GetTime() int {
        return this.Time
}


func (this *GameSkill) SetTime(time int) {
        this.Time = time
}


func (this *GameSkill) AddProperty(gameSkillPropertyInterface GameSkillPropertyInterface) {
        this.Vector.Add(gameSkillPropertyInterface)
}


func (this *GameSkill) GetProperties() IBasicArrayList {
        return this.Vector
}


func GameSkill_GetInstance(gameSkillType *GameSkillType, extra string) IGameSkill {
        return any(GameSkill_HashMap.Get(fmt.Sprint(NewStringMaker().Append(gameSkillType.GetName()).Append(CommonSeps_GetInstance().SPACE).Append(extra)))).(IGameSkill)
}


func (this *GameSkill) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.GAME_SKILL_LABEL)
        stringBuffer.Append(fmt.Sprint(this.GetGameSkillType()))
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(this.EXTRA_LABEL)
        stringBuffer.Append(this.GetExtra())
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(this.TIME_LABEL)
        stringBuffer.Appendlong(int64(this.GetTime()))
        return fmt.Sprint(stringBuffer)
}

func (this *GameSkill) String() string {
    return this.ToString()
}

