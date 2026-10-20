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

type DemoCanvasBasicStartInputProcessor struct {
    BasicMenuInputProcessor
}



func NewDemoCanvasBasicStartInputProcessor(gameKeyEventList IBasicArrayList, gameCanvas IRunnableCanvas) *DemoCanvasBasicStartInputProcessor {
    this := &DemoCanvasBasicStartInputProcessor{}
        this.BasicMenuInputProcessor = *NewBasicMenuInputProcessor(gameKeyEventList, -1, gameCanvas)
    return this
}


// throws Exception
func (this *DemoCanvasBasicStartInputProcessor) ProcessInputList() int {
        var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
        var size int = list.Size(); _ = size
        var gameKeyEvent *GameKeyEvent; _ = gameKeyEvent
        for index := int(0); index < size; index++ {
            gameKeyEvent = any(list.BasicArrayList_().ObjectArray[index]).(*GameKeyEvent)
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Start GameKey: ").Append(StringUtil_GetInstance().ToString(gameKeyEvent))), this, GameInputStrings_GetInstance().PROCESS_INPUT)
        }
        this.GetCanvas().GetCustomCommandListener().CommandAction(GameCommandsFactory_GetInstance().START_COMMAND, NullCanvas_NULL_CANVAS)
        return size
}

