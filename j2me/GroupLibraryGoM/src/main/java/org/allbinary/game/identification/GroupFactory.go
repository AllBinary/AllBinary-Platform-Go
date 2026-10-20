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

type GroupFactory struct {
    LogUtil *LogUtil
    NULL_GROUP_ARRAY []*Group
    List IBasicArrayList
    Index int
}

var GroupFactory_Instance *GroupFactory = NewGroupFactory()

func (this *GroupFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NULL_GROUP_ARRAY = make([]*Group, 0)
    this.List = NewBasicArrayListD()
    this.Index = 0
}

func NewGroupFactory() *GroupFactory {
    this := &GroupFactory{}
    this.initFields()
    return this
}


func GroupFactory_GetInstance() *GroupFactory {
        return GroupFactory_Instance
}


func (this *GroupFactory) GetNextGroup() *Group {
        var group *Group = any(this.List.BasicArrayList_().ObjectArray[this.Index]).(*Group); _ = group
        this.Index++
        return group
}


func (this *GroupFactory) GetNextGroupByName(name string) *Group {
        var group *Group = this.GetNextGroup(); _ = group
        group.SetName(name)
        return group
}


func (this *GroupFactory) Init(groups int16, nameArray []string) {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var TEAM string = "Team "; _ = TEAM
        this.List.Clear()
        this.Index = 0
        var size int = this.List.Size(); _ = size
        for size < int(groups) {
            var name string = stringUtil.EMPTY_STRING; _ = name
            if size < len(nameArray) {
                name = nameArray[size]
            } else {
                stringMaker.Delete(0, stringMaker.Length())
                name = fmt.Sprint(stringMaker.Append(TEAM).Appendint(size))
            }
            this.List.Add(NewGroup(name, int16((size + 3))))
            size++
        }
}

