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

type BasicGroupFactory struct {
    NAME string
    NO_SUCH_GROUP string
    GOOD *Group
    ENEMY *Group
    ENEMY_ON_LEVEL *Group
    NONE *Group
    GOOD_ARRAY []*Group
    ENEMY_ARRAY []*Group
    ENEMY_ON_LEVEL_ARRAY []*Group
    NONE_ARRAY []*Group
}

var BasicGroupFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *BasicGroupFactory) initFields() {
    this.NAME = "GROUP_NAME"
    this.NO_SUCH_GROUP = "No Such Group: "
    this.GOOD = NewGroup("Good Guys", int16(0))
    this.ENEMY = NewGroup("Bad Guys", int16(1))
    this.ENEMY_ON_LEVEL = NewGroup("Bad Guys on Level", int16(2))
    this.NONE = NewGroup("Not On A Team", int16(3))
    this.GOOD_ARRAY = []*Group{this.GOOD}
    this.ENEMY_ARRAY = []*Group{this.ENEMY}
    this.ENEMY_ON_LEVEL_ARRAY = []*Group{this.ENEMY, this.ENEMY_ON_LEVEL}
    this.NONE_ARRAY = []*Group{this.NONE}
}

func NewBasicGroupFactory() *BasicGroupFactory {
    this := &BasicGroupFactory{}
    this.initFields()
    return this
}


func BasicGroupFactory_GetInstance() *BasicGroupFactory {
        if BasicGroupFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            BasicGroupFactory_Instance = NewBasicGroupFactory()
        }
        return any(BasicGroupFactory_Instance).(*BasicGroupFactory)
}

