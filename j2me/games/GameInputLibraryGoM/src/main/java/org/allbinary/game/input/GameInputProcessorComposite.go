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

type GameInputProcessorComposite struct {
    PlayerGameInput
    Name string
    GameInputProcessorInterface GameInputProcessorInterface
    NAME_LABEL string
}


func (this *GameInputProcessorComposite) initFields() {
    
    
    this.NAME_LABEL = " GameInputProcessorInterface: "
}


func NewGameInputProcessorComposite(name string, gameInputProcessorInterface GameInputProcessorInterface) *GameInputProcessorComposite {
    this := &GameInputProcessorComposite{}
        this.PlayerGameInput = *NewPlayerGameInput(NewBasicArrayListD(), NewBasicArrayListD(), -1)
        this.initFields()
        this.Name = name
        this.GameInputProcessorInterface = gameInputProcessorInterface
    return this
}


// synchronized
func (this *GameInputProcessorComposite) Update() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.UPDATE, e)
                }
            }()
            var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
            this.GameInputProcessorInterface.OnInput(list)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameInputProcessorComposite) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(this.PlayerGameInput.ToString()).Append(this.NAME_LABEL).Append(this.Name))
}

func (this *GameInputProcessorComposite) String() string {
    return this.ToString()
}

