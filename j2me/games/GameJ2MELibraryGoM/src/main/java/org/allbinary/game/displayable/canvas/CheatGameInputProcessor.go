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

type CheatGameInputProcessor struct {
    PlayerGameInput
    PROCESS_GAME string
    GameCanvas IAllBinaryGameCanvas
    GameKeyFactory *GameKeyFactory
}

var CheatGameInputProcessor_InputProcessor SpecialGameInputInterface = SpecialGameInputFactory_NO_SPECIAL_GAME_INPUT

func (this *CheatGameInputProcessor) initFields() {
    this.PROCESS_GAME = "processGame"
    
    this.GameKeyFactory = GameKeyFactory_GetInstance()
}


func NewCheatGameInputProcessor(gameCanvas IAllBinaryGameCanvas) *CheatGameInputProcessor {
    this := &CheatGameInputProcessor{}
        this.PlayerGameInput = *NewPlayerGameInput(NewBasicArrayListD(), NewBasicArrayListD(), -1)
        this.initFields()
        this.GameCanvas = gameCanvas
    return this
}


// synchronized
func (this *CheatGameInputProcessor) Update() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.UPDATE, e)
                }
            }()
            var gameLayerManager IAllBinaryGameLayerManager = this.GameCanvas.GetLayerManager(); _ = gameLayerManager
            var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
            var size int = list.Size(); _ = size
            for index := int(0); index < size; index++ {
                var anyType any = list.BasicArrayList_().ObjectArray[index]; _ = anyType
                var key int = GameKeyEventUtil_GetKey(anyType); _ = key
                if key == this.GameKeyFactory.LEVEL_DOWN.GetId() {
                    CheatGameInputProcessor_InputProcessor.StrafeLeft()
                    break
                } else if key == this.GameKeyFactory.LEVEL_UP.GetId() {
                    CheatGameInputProcessor_InputProcessor.StrafeRight()
                    break
                }
                if key == this.GameKeyFactory.LEFT.GetId() {
                    CheatGameInputProcessor_InputProcessor.Left()
                    break
                } else if key == this.GameKeyFactory.RIGHT.GetId() {
                    CheatGameInputProcessor_InputProcessor.Right()
                    break
                } else if key == this.GameKeyFactory.UP.GetId() {
                    CheatGameInputProcessor_InputProcessor.Up()
                    break
                } else if key == this.GameKeyFactory.DOWN.GetId() {
                    CheatGameInputProcessor_InputProcessor.Down()
                    break
                } else if key == this.GameKeyFactory.KEY_NUM1.GetId() {
                    CheatGameInputProcessor_InputProcessor.Special1(gameLayerManager, GameKeyEvent_NONE)
                    break
                } else if key == this.GameKeyFactory.KEY_NUM3.GetId() {
                    CheatGameInputProcessor_InputProcessor.Special2(gameLayerManager, GameKeyEvent_NONE)
                    break
                }
            }
            list.Clear()
            return
        }(); tryControl1 == 1 {
            return
        }
}

