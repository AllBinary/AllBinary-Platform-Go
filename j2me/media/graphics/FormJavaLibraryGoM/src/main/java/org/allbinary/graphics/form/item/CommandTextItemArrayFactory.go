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

type CommandTextItemArrayFactory struct {
    List IBasicArrayList
    VisitorInterface IVisitor
}


func (this *CommandTextItemArrayFactory) initFields() {
    this.List = NewBasicArrayListD()
    
}


func NewCommandTextItemArrayFactory(visitorInterface IVisitor) *CommandTextItemArrayFactory {
    this := &CommandTextItemArrayFactory{}
    this.initFields()
        this.VisitorInterface = visitorInterface
    return this
}


func (this *CommandTextItemArrayFactory) GetInstance(vector any, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) []IABCustomItem {
        var vectorUtil *VectorUtil = VectorUtil_GetInstance(); _ = vectorUtil
        var size int = vectorUtil.GetSize(vector); _ = size
        this.List.Clear()
        var textItem *CommandTextItem; _ = textItem
        var priorityLimit int = 7; _ = priorityLimit
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var isLargeEnoughDisplayForBigMenu bool = (displayInfo.IsPortrait() && displayInfo.GetLastHeight() >= ScreenInfo_GetInstance().MEDIUM_WIDTH) || (!displayInfo.IsPortrait() && displayInfo.GetLastWidth() >= ScreenInfo_GetInstance().MEDIUM_WIDTH); _ = isLargeEnoughDisplayForBigMenu
        if !isLargeEnoughDisplayForBigMenu && size > 3 {
            priorityLimit = 3
        }
        var command ICommand; _ = command
        for index := int(0); index < size; index++ {
            command = any(vectorUtil.ElementAt(vector, index)).(ICommand)
            if command.GetPriority() < priorityLimit {
                var aBoolean bool = any(this.VisitorInterface.Visit(command)).(bool); _ = aBoolean
                if aBoolean {
                    textItem = NewCommandTextItem(command, ImageItem_LAYOUT_DEFAULT, StringUtil_GetInstance().EMPTY_STRING, backgroundBasicColor, foregroundBasicColor)
                    this.List.Add(textItem)
                }
            }
        }
        var textItemArray []IABCustomItem = make([]IABCustomItem, this.List.Size()); _ = textItemArray
        var size2 int = len(textItemArray); _ = size2
        for index := int(0); index < size2; index++ {
            textItemArray[index] = any(this.List.BasicArrayList_().ObjectArray[index]).(IABCustomItem)
        }
        return textItemArray
}

