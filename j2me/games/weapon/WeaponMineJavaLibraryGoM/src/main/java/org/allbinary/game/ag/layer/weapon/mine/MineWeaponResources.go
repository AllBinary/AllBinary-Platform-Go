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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type MineWeaponResources struct {
    BasicGameResources
    DROP_TEXT_RESOURCE string
    DROP_RESOURCE string
    DROP_ICON_RESOURCE string
}

var MineWeaponResources_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *MineWeaponResources) initFields() {
    this.DROP_TEXT_RESOURCE = "/mine_drop_text_20_by_20.png"
    this.DROP_RESOURCE = "/mine_drop_20_by_20.png"
    this.DROP_ICON_RESOURCE = "/mine_drop_icon_20_by_20.png"
}


func MineWeaponResources_GetInstance() *MineWeaponResources {
        if MineWeaponResources_Instance == NullUtil_GetInstance().NULL_OBJECT {
            MineWeaponResources_Instance = NewMineWeaponResources()
        }
        return any(MineWeaponResources_Instance).(*MineWeaponResources)
}


func NewMineWeaponResources() *MineWeaponResources {
    this := &MineWeaponResources{}
    this.BasicGameResources = *NewBasicGameResources()
    this.initFields()
        var ROOT string = "/mine"; _ = ROOT
        var SMALL string = "_20_by_20.png"; _ = SMALL
        var MEDIUM string = SMALL; _ = MEDIUM
        var SIZE_FOUR string = SMALL; _ = SIZE_FOUR
        var SIZE_FIVE string = SMALL; _ = SIZE_FIVE
        var SIZE_SIX string = SMALL; _ = SIZE_SIX
        var SIZE []string = []string{SMALL, MEDIUM, SIZE_FOUR, SIZE_FIVE, SIZE_SIX}; _ = SIZE
        this.BasicGameResources.Init(ROOT, SIZE)
    return this
}

