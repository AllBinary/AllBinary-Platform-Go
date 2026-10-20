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

type Group struct {
    Name string
    TeamId int16
    String string
}

// Implements GroupInterface


func (this *Group) initFields() {
    this.Name = StringUtil_GetInstance().EMPTY_STRING
    
    this.String = StringUtil_GetInstance().EMPTY_STRING
}


func NewGroup(teamName string, teamId int16) *Group {
    this := &Group{}
    this.initFields()
        this.TeamId = teamId
        this.SetName(teamName)
    return this
}


func (this *Group) GetGroupName() string {
        return this.Name
}


func (this *Group) SetName(name string) {
        this.Name = name
        var groupCommonFactory *GroupCommonFactory = GroupCommonFactory_GetInstance(); _ = groupCommonFactory
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(groupCommonFactory.GROUP_NAME_LABEL)
        stringBuffer.Append(this.Name)
        stringBuffer.Append(groupCommonFactory.ID_LABEL)
        stringBuffer.Appendshort(this.TeamId)
        this.String = fmt.Sprint(stringBuffer)
}


func (this *Group) GetGroupId() int16 {
        return this.TeamId
}


func (this *Group) ToString() string {
        return this.String
}

func (this *Group) String() string {
    return this.ToString()
}

