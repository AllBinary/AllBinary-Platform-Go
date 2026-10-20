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

type BufferedImageUtil2 struct {
    ImageUtil *ImageUtil
}

var BufferedImageUtil2_Instance *BufferedImageUtil2 = NewBufferedImageUtil2()

func (this *BufferedImageUtil2) initFields() {
    this.ImageUtil = ImageUtil_GetInstance()
}

func NewBufferedImageUtil2() *BufferedImageUtil2 {
    this := &BufferedImageUtil2{}
    this.initFields()
    return this
}


func BufferedImageUtil2_GetInstance() *BufferedImageUtil2 {
        return BufferedImageUtil2_Instance
}


func (this *BufferedImageUtil2) CreateSpriteImage(bufferedImageArray []*BufferedImage) *BufferedImage {
        var columns int = len(bufferedImageArray); _ = columns
        var max_ int = columns; _ = max_
        var rows int = 0; _ = rows
        if len(bufferedImageArray) < columns {
            columns = len(bufferedImageArray)
        }
        rows = (len(bufferedImageArray) / columns)
        if len(bufferedImageArray) % columns != 0 {
            rows++
        }
        var bufferedImage *BufferedImage = this.ImageUtil.Create(bufferedImageArray[0].GetWidth(nil) * columns, bufferedImageArray[0].GetHeight(nil) * rows); _ = bufferedImage
        var g *Graphics2D = bufferedImage.CreateGraphics(); _ = g
        var columnIndex int = 0; _ = columnIndex
        var rowIndex int = 0; _ = rowIndex
        for index := int(0); index < len(bufferedImageArray); index++ {
            if index / max_ != 0 && index % max_ == 0 {
                rowIndex++
                columnIndex = 0
            }
            g.DrawImage(bufferedImageArray[index], bufferedImageArray[index].GetWidth(nil) * columnIndex, bufferedImageArray[index].GetHeight(nil) * rowIndex, bufferedImageArray[index].GetWidth(nil), bufferedImageArray[index].GetHeight(nil), nil)
            columnIndex++
        }
        g.Dispose()
        return bufferedImage
}

