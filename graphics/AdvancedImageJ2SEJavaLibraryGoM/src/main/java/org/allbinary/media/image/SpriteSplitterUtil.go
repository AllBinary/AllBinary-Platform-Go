/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type SpriteSplitterUtil struct {
    LogUtil *LogUtil
    CommonSeps *CommonSeps
    CommonStrings *CommonStrings
    ANIMATIONS_LABELS []string
    DIRECTIONAL_ANIMATIONS string
    HORIZONTAL_ANIMATIONS string
    HORIZONTAL_SPRITE string
    ROW_NAME []string
    _ROW string
}

var SpriteSplitterUtil_Instance *SpriteSplitterUtil = NewSpriteSplitterUtil()

func (this *SpriteSplitterUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ANIMATIONS_LABELS = []string{"Row Total:", "Column Total:", "Column Total:", "Row Total:", "Column Total:", "Row Total:"}
    this.DIRECTIONAL_ANIMATIONS = "Directional DLRU"
    this.HORIZONTAL_ANIMATIONS = "Horizontal Animations"
    this.HORIZONTAL_SPRITE = "Horizontal A=R=Y F=C=X"
    this.ROW_NAME = []string{"idle", "walk", "attack", "defeat"}
    this._ROW = "_row"
}

func NewSpriteSplitterUtil() *SpriteSplitterUtil {
    this := &SpriteSplitterUtil{}
    this.initFields()
    return this
}


func SpriteSplitterUtil_GetInstance() *SpriteSplitterUtil {
        return SpriteSplitterUtil_Instance
}


// throws Exception
func (this *SpriteSplitterUtil) Process(imageProcessorInput *ImageProcessorInput, totalFrames int, totalAnimations int, widthReduction int, heightReduction int, increaseWidth int, increaseHeight int, spriteType string, visitor ImageProcessedVisitor) {
        var bufferedImage *BufferedImage; _ = bufferedImage
        var generatedBufferedImageArray [][]*BufferedImage; _ = generatedBufferedImageArray
        var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
        for index := int(0); index < len(bufferedImageArray); index++ {
            bufferedImage = bufferedImageArray[index]
            this.LogUtil.PutF(spriteType, this, this.CommonStrings.RUN)
            if spriteType == this.HORIZONTAL_SPRITE {
                var cellHeight int = bufferedImage.GetHeight() / totalAnimations; _ = cellHeight
                var cellWidth int = bufferedImage.GetWidth() / totalFrames; _ = cellWidth
                var columns int = totalFrames; _ = columns
                var rows int = totalAnimations; _ = rows
                generatedBufferedImageArray = func() [][]*BufferedImage { array := make([][]*BufferedImage, rows); for index := range array { array[index] = make([]*BufferedImage, columns) }; return array }()
                this.LogUtil.PutF("Processing Individual Cells columns: " + fmt.Sprint(columns) + " rows: " + fmt.Sprint(rows), this, this.CommonStrings.RUN)
                this.LogUtil.PutF("Processing Individual Cells cellHeight: " + fmt.Sprint(cellHeight) + " cellWidth: " + fmt.Sprint(cellWidth), this, this.CommonStrings.RUN)
                var imageUtil *ImageUtil = ImageUtil_GetInstance(); _ = imageUtil
                var nameEnding string = nil; _ = nameEnding
                var x int = 0; _ = x
                var y int = 0; _ = y
                for index2 := int(0); index2 < rows; index2++ {
                    y = cellHeight * index2
                    for index3 := int(0); index3 < columns; index3++ {
                        x = cellWidth * index3
                        generatedBufferedImageArray[index2][index3] = bufferedImage.GetSubimage(x + widthReduction, y + heightReduction, cellWidth - (widthReduction * 2), cellHeight - (heightReduction * 2))
                        if increaseWidth != 0 || increaseHeight != 0 {
                            generatedBufferedImageArray[index2][index3] = imageUtil.CreateBufferedImageAllowTranslate(generatedBufferedImageArray[index2][index3], cellWidth + increaseWidth, cellHeight + increaseHeight, false, true)
                        }
                        nameEnding = fmt.Sprint(NewStringMaker().Appendint(index2).Append(this.CommonSeps.UNDERSCORE).Appendint(index3))
                        visitor.Visit(generatedBufferedImageArray[index2][index3], nameEnding, index)
                    }
                }
                this.LogUtil.PutF("Processing Rows from Cells", this, this.CommonStrings.RUN)
                var imageUnifierProperties *ImageUnifierProperties = NewImageUnifierProperties(); _ = imageUnifierProperties
                imageUnifierProperties.SetRows(fmt.Sprint(1))
                imageUnifierProperties.SetColumns(fmt.Sprint(columns))
                var imageUnifierCell *ImageUnifierCell = NewImageUnifierCell(fmt.Sprint(cellWidth - (2 * widthReduction)), fmt.Sprint(cellHeight - (2 * heightReduction))); _ = imageUnifierCell
                imageUnifierProperties.SetImageUnifierCell(imageUnifierCell)
                for index2 := int(0); index2 < rows; index2++ {
                    y = cellHeight * index2
                    var tempBufferedImageArray []*BufferedImage = make([]*BufferedImage, columns); _ = tempBufferedImageArray
                    for index3 := int(0); index3 < columns; index3++ {
                        tempBufferedImageArray[index3] = generatedBufferedImageArray[index2][index3]
                        nameEnding = fmt.Sprint(NewStringMaker().Appendint(index2).Append(this.CommonSeps.UNDERSCORE).Appendint(1).Append(this._ROW))
                    }
                    var generatedBufferedImage *BufferedImage = ImageUnifierUtil_GetInstance().GetImage(tempBufferedImageArray, imageUnifierProperties); _ = generatedBufferedImage
                    visitor.Visit(generatedBufferedImage, nameEnding, index)
                }
            } else if spriteType == this.HORIZONTAL_ANIMATIONS {
            } else if spriteType == this.DIRECTIONAL_ANIMATIONS {
                var totalFrames2 int = 4; _ = totalFrames2
                var cellHeight int = bufferedImage.GetHeight() / totalFrames2; _ = cellHeight
                var cellWidth int = bufferedImage.GetWidth() / totalAnimations; _ = cellWidth
                var columns int = totalAnimations; _ = columns
                var rows int = totalFrames2; _ = rows
                generatedBufferedImageArray = func() [][]*BufferedImage { array := make([][]*BufferedImage, rows); for index := range array { array[index] = make([]*BufferedImage, columns) }; return array }()
                this.LogUtil.PutF("Processing Individual Cells for each Direction", this, this.CommonStrings.RUN)
                var nameEnding string = nil; _ = nameEnding
                var x int = 0; _ = x
                var y int = 0; _ = y
                for index2 := int(0); index2 < rows; index2++ {
                    y = cellHeight * index2
                    for index3 := int(0); index3 < columns; index3++ {
                        x = cellWidth * index3
                        generatedBufferedImageArray[index2][index3] = bufferedImage.GetSubimage(x + widthReduction, y + heightReduction, cellWidth - (widthReduction * 2), cellHeight - (heightReduction * 2))
                        if increaseWidth != 0 || increaseHeight != 0 {
                            generatedBufferedImageArray[index2][index3] = ImageUtil_GetInstance().CreateBufferedImageAllowTranslate(generatedBufferedImageArray[index2][index3], cellWidth + increaseWidth, cellHeight + increaseHeight, false, true)
                        }
                        nameEnding = fmt.Sprint(NewStringMaker().Append(this.CommonStrings.DIRECTION_NAME[index2]).Append(this.CommonSeps.UNDERSCORE).Appendint(index3))
                        visitor.Visit(generatedBufferedImageArray[index2][index3], nameEnding, index)
                    }
                }
                var imageUnifierProperties *ImageUnifierProperties = NewImageUnifierProperties(); _ = imageUnifierProperties
                imageUnifierProperties.SetRows(fmt.Sprint(1))
                imageUnifierProperties.SetColumns(fmt.Sprint(columns))
                var imageUnifierCell *ImageUnifierCell = NewImageUnifierCell(fmt.Sprint(cellWidth - (2 * widthReduction)) + increaseWidth, fmt.Sprint(cellHeight - (2 * heightReduction) + increaseHeight)); _ = imageUnifierCell
                imageUnifierProperties.SetImageUnifierCell(imageUnifierCell)
                this.LogUtil.PutF("Processing Rows from Cells for each Direction", this, this.CommonStrings.RUN)
                for index2 := int(0); index2 < rows; index2++ {
                    y = cellHeight * index2
                    var tempBufferedImageArray []*BufferedImage = make([]*BufferedImage, columns); _ = tempBufferedImageArray
                    for index3 := int(0); index3 < columns; index3++ {
                        tempBufferedImageArray[index3] = generatedBufferedImageArray[index2][index3]
                        nameEnding = fmt.Sprint(NewStringMaker().Append(this.CommonStrings.DIRECTION_NAME[index2]).Append(this.CommonSeps.UNDERSCORE).Appendint(1).Append(this._ROW))
                    }
                    var generatedBufferedImage *BufferedImage = ImageUnifierUtil_GetInstance().GetImage(tempBufferedImageArray, imageUnifierProperties); _ = generatedBufferedImage
                    visitor.Visit(generatedBufferedImage, nameEnding, index)
                }
            }
        }
}

