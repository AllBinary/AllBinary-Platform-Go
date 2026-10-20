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

type GameInputMappingInstructionsCanvas struct {
    GameCommandCanvas
    DrawStringUtil *DrawStringUtil
    TITLE string
    Instructions []string
    ColorFillPaintable IColorFillBasePaintable
    Anchor int
}

var GameInputMappingInstructionsCanvas_DISPLAY ICommand = NewCommand("Help", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
var GameInputMappingInstructionsCanvas_CLOSE ICommand = NewCommand("Close", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)
var GameInputMappingInstructionsCanvas_NAME string = "GameInputMappingInstructionsCanvas"

func (this *GameInputMappingInstructionsCanvas) initFields() {
    this.DrawStringUtil = DrawStringUtil_GetInstance()
    this.TITLE = "Input Instructions"
    this.Instructions = []string{"Add Input Mapping:", "1. Select the desired action by generating a currently defined input for a given action.", "(Example: Press the '1' key to select the Fire action)", "2. Make the newly desired input for the selected action.", "(Example: Press 'f' to add it to the selected action)", "Note: If the input is already mapped to another action then it will not be added. You", "will need to deleted from the other action before adding it to another.", "Warning: Not all keys map for a given platform and show as Unknown.", StringUtil_GetInstance().EMPTY_STRING, "Remove Input Mapping:", "1. Select the desired action by generating a currently defined input for a given action.", "(Example: Press the key '1' to select the Fire action)", "2. Select a mapped input for the selected action.", "(Example: Press the key '1' to select the '1' key input for the selected Fire action)", "3. Press the Delete Key."}
    
    this.Anchor = Anchor_TOP_LEFT
}


func NewGameInputMappingInstructionsCanvas(commandListener CommandListener, allBinaryGameLayerManager IAllBinaryGameLayerManager) *GameInputMappingInstructionsCanvas {
    this := &GameInputMappingInstructionsCanvas{}
        this.GameCommandCanvas = *NewGameCommandCanvas(commandListener, GameInputMappingInstructionsCanvas_NAME, allBinaryGameLayerManager.GetBackgroundBasicColor(), allBinaryGameLayerManager.GetForegroundBasicColor())
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.ColorFillPaintable = ColorFillPaintableFactory_GetInstance().GetInstance(allBinaryGameLayerManager.GetBackgroundBasicColor(), false)
    return this
}


func (this *GameInputMappingInstructionsCanvas) InitCommands(cmdListener CommandListener) {
        this.RemoveAllCommands()
        this.AddCommand(GameInputMappingInstructionsCanvas_CLOSE)
        this.SetCommandListener(cmdListener)
}


func (this *GameInputMappingInstructionsCanvas) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        this.ColorFillPaintable.Paint(graphics)
        var halfWidth int = this.DisplayInfo.GetLastHalfWidth(); _ = halfWidth
        var beginWidth int = (graphics.GetFont().StringWidth(this.TITLE) >> 1); _ = beginWidth
        graphics.SetColor(this.ForegroundColor)
        graphics.DrawString(this.TITLE, halfWidth - beginWidth, this.FontHeight, this.Anchor)
        this.DrawStringUtil.DrawCenterStrings(graphics, this.Instructions, this.DisplayInfo.GetLastWidth(), this.FontHeight, halfWidth, 3 * this.FontHeight)
        this.GameCommandCanvas.Paint(graphics)
}

