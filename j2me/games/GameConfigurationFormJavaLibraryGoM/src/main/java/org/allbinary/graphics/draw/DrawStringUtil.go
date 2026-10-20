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

type DrawStringUtil struct {
    EMPTY_STRING string
    Anchor int
}

var DrawStringUtil_Instance *DrawStringUtil = NewDrawStringUtil()

func (this *DrawStringUtil) initFields() {
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
    this.Anchor = Anchor_TOP_LEFT
}

func NewDrawStringUtil() *DrawStringUtil {
    this := &DrawStringUtil{}
    this.initFields()
    return this
}


func DrawStringUtil_GetInstance() *DrawStringUtil {
        return DrawStringUtil_Instance
}


func (this *DrawStringUtil) DrawCenterString(graphics IGraphics, string_ string, offset int, length int, x int, y int) {
        var width int = (graphics.GetFont().SubstringWidth(string_, offset, length) >> 1); _ = width
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Exception: stringLength: ").Appendint(len(string_)).Append(" offset: ").Appendint(offset).Append(" currentLength: ").Appendint(length)), this, "drawCenterString")
                }
            }()
            graphics.DrawSubstring(string_, offset, length, x - width, y, this.Anchor)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DrawStringUtil) DrawCenterStrings(graphics IGraphics, stringArray []string, maxWidth int, charHeight int, x int, y int) {
        var extraLines int = 0; _ = extraLines
        for index := int(0); index < len(stringArray); index++ {
            var string_ string = stringArray[index]; _ = string_
            if string_ != this.EMPTY_STRING {
                var width int = graphics.GetFont().StringWidth(string_); _ = width
                var minTotalLines int = 1; _ = minTotalLines
                if width > maxWidth {
                    minTotalLines = (width / maxWidth) + 1
                }
                var linePortion int = len(string_) / minTotalLines; _ = linePortion
                var offset int = 0; _ = offset
                var currentLength int = linePortion; _ = currentLength
                var size int = len(string_); _ = size
                for offset < size {
                    if offset + currentLength > size {
                        currentLength = size - offset
                    }
                    if offset + currentLength != size {
                        for currentLength > 0 && rune(stringArray[index][offset + currentLength]) != ' ' {
                            currentLength--
                        }
                        if currentLength <= 4 {
                            currentLength = linePortion
                        }
                    }
                    this.DrawCenterString(graphics, string_, offset, currentLength, x, y + ((index + Java_PostIncrement(&extraLines)) * charHeight))
                    offset = offset + currentLength
                    currentLength = linePortion
                }
                extraLines--
            }
        }
}

