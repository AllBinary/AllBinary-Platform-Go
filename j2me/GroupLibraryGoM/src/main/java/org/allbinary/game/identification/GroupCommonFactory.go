/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type GroupCommonFactory struct {
    ID int
    GROUP_NAME_LABEL string
    ID_LABEL string
    ADDING_LABEL string
    TO_LABEL string
}

var GroupCommonFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GroupCommonFactory) initFields() {
    this.ID = SmallIntegerSingletonFactory_GetInstance().GetAt(10)
    this.GROUP_NAME_LABEL = "Group Name: "
    this.ID_LABEL = " Id: "
    this.ADDING_LABEL = "Adding: "
    this.TO_LABEL = " to: "
}

func NewGroupCommonFactory() *GroupCommonFactory {
    this := &GroupCommonFactory{}
    this.initFields()
    return this
}


func GroupCommonFactory_GetInstance() *GroupCommonFactory {
        if GroupCommonFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GroupCommonFactory_Instance = NewGroupCommonFactory()
        }
        return any(GroupCommonFactory_Instance).(*GroupCommonFactory)
}

