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

type GameInput struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GameKeyEventList IBasicArrayList
    RemovalGameKeyEventList IBasicArrayList
    IsRemoveDuplicateKeyPresses bool
}

// The Java class GameInput has subclasses so its variables are this interface
type IGameInput interface {
    GameInput_() *GameInput
    Add(gameKeyEvent *GameKeyEvent)
    AddForRemoval(gameKeyEvent *GameKeyEvent)
    IsForRemoval(gameKeyEvent *GameKeyEvent) bool
    Clear()
    RemoveNonAIInputGameKeyEvents()
    Update()
    GetRemovalGameKeyEventList() IBasicArrayList
    GetGameKeyEventList() IBasicArrayList
}

func (this *GameInput) GameInput_() *GameInput {
    return this
}


func (this *GameInput) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewGameInput(gameKeyEventList IBasicArrayList, removalGameKeyEventList IBasicArrayList) *GameInput {
    this := &GameInput{}
    this.initFields()
        this.GameKeyEventList = gameKeyEventList
        this.RemovalGameKeyEventList = removalGameKeyEventList
        this.IsRemoveDuplicateKeyPresses = Features_GetInstance().IsFeature(InputFeatureFactory_GetInstance().REMOVE_DUPLICATE_KEY_PRESSES)
    return this
}


// synchronized
func (this *GameInput) Add(gameKeyEvent *GameKeyEvent) {
        if this.IsRemoveDuplicateKeyPresses && this.GameKeyEventList.Contains(gameKeyEvent) {
            return
        }
        if gameKeyEvent != nil {
            this.GameKeyEventList.Add(gameKeyEvent)
        } else {
            this.LogUtil.PutF("Danger Passed Null GameKeyEvent", this, this.CommonStrings.ADD)
        }
}


// synchronized
func (this *GameInput) AddForRemoval(gameKeyEvent *GameKeyEvent) {
        this.RemovalGameKeyEventList.Add(gameKeyEvent)
}


// synchronized
func (this *GameInput) IsForRemoval(gameKeyEvent *GameKeyEvent) bool {
        return this.RemovalGameKeyEventList.Contains(gameKeyEvent)
}


// synchronized
func (this *GameInput) Clear() {
        this.GameKeyEventList.Clear()
}


// synchronized
func (this *GameInput) RemoveNonAIInputGameKeyEvents() {
        var list IBasicArrayList = this.GameKeyEventList; _ = list
        var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
        for index := int(list.Size()); Java_PreDecrement(&index) >= 0;  {
            gameKeyEvent = any(list.BasicArrayList_().ObjectArray[index]).(*GameKeyEvent)
            if gameKeyEvent.GetSourceId() != 1 {
                list.RemoveAt(index)
            }
        }
}


// synchronized
func (this *GameInput) Update() {
        var removeList IBasicArrayList = this.RemovalGameKeyEventList; _ = removeList
        var list IBasicArrayList = this.GameKeyEventList; _ = list
        var size int = removeList.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = removeList.BasicArrayList_().ObjectArray[index]; _ = anyType
            for index2 := int(list.Size()); Java_PreDecrement(&index2) >= 0;  {
                if list.BasicArrayList_().ObjectArray[index2] == anyType {
                    list.RemoveAt(index2)
                }
            }
        }
        removeList.Clear()
}


func (this *GameInput) GetRemovalGameKeyEventList() IBasicArrayList {
        return this.RemovalGameKeyEventList
}


func (this *GameInput) GetGameKeyEventList() IBasicArrayList {
        return this.GameKeyEventList
}

