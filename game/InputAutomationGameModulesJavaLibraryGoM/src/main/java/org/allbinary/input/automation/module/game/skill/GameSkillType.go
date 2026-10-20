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

type GameSkillType struct {
    SPELL_GAME_SKILL_TYPE string
    Name string
}

var GameSkillType_MANA *GameSkillType = NewGameSkillType("Mana")
var GameSkillType_REPLENISH *GameSkillType = NewGameSkillType("Replenish")
var GameSkillType_TIMED *GameSkillType = NewGameSkillType("Timed")
var GameSkillType_ATTACK *GameSkillType = NewGameSkillType("Attack")
var GameSkillType_TRANSPORT *GameSkillType = NewGameSkillType("Transport")

func (this *GameSkillType) initFields() {
    this.SPELL_GAME_SKILL_TYPE = "SpellGameSkillType: "
    
}


func NewGameSkillType(name string) *GameSkillType {
    this := &GameSkillType{}
    this.initFields()
        this.SetName(name)
    return this
}


func (this *GameSkillType) GetName() string {
        return this.Name
}


func (this *GameSkillType) SetName(name string) {
        this.Name = name
}


func (this *GameSkillType) ToString() string {
        return this.SPELL_GAME_SKILL_TYPE + this.GetName()
}

func (this *GameSkillType) String() string {
    return this.ToString()
}

