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

type GameInputProcessor struct {
}

// The Java class GameInputProcessor has subclasses so its variables are this interface
type IGameInputProcessor interface {
    GameInputProcessor_() *GameInputProcessor
    ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    ProcessReleasedEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent)
    ProcessAnalog(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent, analogValue int)
    Process(allbinaryLayerManager IAllBinaryLayerManager, keyAsInteger int)
    ProcessReleased(allbinaryLayerManager IAllBinaryLayerManager, keyAsInteger int)
}

func (this *GameInputProcessor) GameInputProcessor_() *GameInputProcessor {
    return this
}

var GameInputProcessor_Instance IGameInputProcessor = NewGameInputProcessor()


func GameInputProcessor_GetInstance() IGameInputProcessor {
        return GameInputProcessor_Instance
}


func NewGameInputProcessor() *GameInputProcessor {
    this := &GameInputProcessor{}
    return this
}


// throws Exception
func (this *GameInputProcessor) ProcessEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *GameInputProcessor) ProcessReleasedEvent(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent) {
}


// throws Exception
func (this *GameInputProcessor) ProcessAnalog(allbinaryLayerManager IAllBinaryLayerManager, gameKeyEvent *GameKeyEvent, analogValue int) {
}


// throws Exception
func (this *GameInputProcessor) Process(allbinaryLayerManager IAllBinaryLayerManager, keyAsInteger int) {
}


// throws Exception
func (this *GameInputProcessor) ProcessReleased(allbinaryLayerManager IAllBinaryLayerManager, keyAsInteger int) {
}

