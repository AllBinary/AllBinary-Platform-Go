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

type ImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CanvasStrings *CanvasStrings
    IsDrawn bool
    BufferedImage *BufferedImage
}


func (this *ImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CanvasStrings = CanvasStrings_GetInstance()
    this.IsDrawn = false
    
}


func NewImageJPanel(bufferedImage *BufferedImage) *ImageJPanel {
    this := &ImageJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
        this.BufferedImage = bufferedImage
        this.SetSize(this.BufferedImage.GetWidth(nil), this.BufferedImage.GetHeight(nil))
    return this
}


func (this *ImageJPanel) Paint(graphics IGraphics) {
        if !this.IsDrawn {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CanvasStrings.PAINT)
            graphics.DrawImage(this.BufferedImage, 0, 0, this.BufferedImage.GetWidth(nil), this.BufferedImage.GetHeight(nil), nil)
            this.IsDrawn = true
        }
}


func (this *ImageJPanel) InitComponents() {
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 400, Short_MAX_VALUE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 300, Short_MAX_VALUE))
}

