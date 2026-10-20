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

type MirrorSpriteImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageStrings *ImageStrings
    ImageProcessorInput *ImageProcessorInput
    Result *BufferedImage
    GenerateJButton *JButton
    HorizontalJCheckBox *JCheckBox
    ImageHeightIsFrameSizeJCheckBox *JCheckBox
    JLabel2 *JLabel
    JLayeredPane1 *JLayeredPane
    JPanel1 *JPanel
    JPanel2 *JPanel
    JPanel3 *JPanel
    VerticleJCheckBox *JCheckBox
    WriteOverOriginalJCheckBox *JCheckBox
}

// Implements ImageProcessorInputCompositeInterface


func (this *MirrorSpriteImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageStrings = ImageStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewMirrorSpriteImageJPanel(imageProcessorInput *ImageProcessorInput) *MirrorSpriteImageJPanel {
    this := &MirrorSpriteImageJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.LogUtil.PutF("Starting", this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
    return this
}


func (this *MirrorSpriteImageJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *MirrorSpriteImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var generatedBufferedImageArray []*BufferedImage; _ = generatedBufferedImageArray
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
                    for index := int(0); index < len(bufferedImageArray); index++ {
                        generatedBufferedImageArray = MirrorImageUtil_GetInstance().GetImages(bufferedImageArray[index], this.VerticleJCheckBox.IsSelected(), this.HorizontalJCheckBox.IsSelected())
                        this.Result = BufferedImageUtil2_GetInstance().CreateSpriteImage(generatedBufferedImageArray)
                        this.GetParent().Repaint()
                        var file *File = imageProcessorInput.GetFiles()[index]; _ = file
                        if !this.WriteOverOriginalJCheckBox.IsSelected() {
                            var filePath string = file.GetAbsolutePath(); _ = filePath
                            var extensionIndex int = strings.Index(filePath, string(this.ImageStrings.PNG_EXTENSION)); _ = extensionIndex
                            filePath = filePath[0:extensionIndex] + "_mirror" + fmt.Sprint(this.ImageStrings.PNG_EXTENSION)
                            logUtil.PutF("Renamed File: " + filePath, this, this.CommonStrings.RUN)
                            file = NewFile(filePath)
                        }
                        var isWritten bool = ImageIO_Write(any(this.Result).(*RenderedImage), this.ImageStrings.PNG, file); _ = isWritten
                        logUtil.PutF("File: " + fmt.Sprint(file) + " Wrote: " + fmt.Sprint(isWritten), this, this.CommonStrings.RUN)
                    }
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *MirrorSpriteImageJPanel) InitComponents() {
        this.JLayeredPane1 = NewJLayeredPane()
        this.JPanel1 = /* Java anonymous class body: implement it as a named type
            

        func (this *MirrorSpriteImageJPanel) Paint(graphics IGraphics) {
                if this.Result != nil {
                    graphics.DrawImage(this.Result, 0, 0, this.Result.GetWidth(nil), this.Result.GetHeight(nil), nil)
                }
        }
*/ NewJPanel()
        this.JPanel2 = NewJPanel()
        this.ImageHeightIsFrameSizeJCheckBox = NewJCheckBox()
        this.HorizontalJCheckBox = NewJCheckBox()
        this.VerticleJCheckBox = NewJCheckBox()
        this.WriteOverOriginalJCheckBox = NewJCheckBox()
        this.GenerateJButton = NewJButton()
        this.JPanel3 = NewJPanel()
        this.JLabel2 = NewJLabel()
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(0, 507, Short_MAX_VALUE))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(0, 230, Short_MAX_VALUE))
        this.ImageHeightIsFrameSizeJCheckBox.SetSelected(true)
        this.ImageHeightIsFrameSizeJCheckBox.SetText("Image Height Is Frame Size")
        this.ImageHeightIsFrameSizeJCheckBox.SetEnabled(false)
        this.ImageHeightIsFrameSizeJCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MirrorSpriteImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.ImageHeightIsFrameSizeJCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        this.HorizontalJCheckBox.SetSelected(true)
        this.HorizontalJCheckBox.SetText("Horizontal")
        this.HorizontalJCheckBox.SetEnabled(false)
        this.VerticleJCheckBox.SetText("Verticle")
        this.VerticleJCheckBox.SetEnabled(false)
        this.WriteOverOriginalJCheckBox.SetText("Write Over Original")
        this.GenerateJButton.SetText("Generate")
        this.GenerateJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MirrorSpriteImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.GenerateJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var jPanel2Layout *GroupLayout = NewGroupLayout(this.JPanel2); _ = jPanel2Layout
        this.JPanel2.SetLayout(jPanel2Layout)
        jPanel2Layout.SetHorizontalGroup(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel2Layout.CreateSequentialGroup().AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(this.VerticleJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.UNRELATED).Add(this.HorizontalJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.ImageHeightIsFrameSizeJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.WriteOverOriginalJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.GenerateJButton)))
        jPanel2Layout.SetVerticalGroup(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel2Layout.CreateSequentialGroup().Add(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.ImageHeightIsFrameSizeJCheckBox, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(this.HorizontalJCheckBox).Add(this.VerticleJCheckBox).Add(this.WriteOverOriginalJCheckBox).Add(this.GenerateJButton, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, 25, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE)).AddContainerGap()))
        this.JLabel2.SetText("Results:")
        var jPanel3Layout *GroupLayout = NewGroupLayout(this.JPanel3); _ = jPanel3Layout
        this.JPanel3.SetLayout(jPanel3Layout)
        jPanel3Layout.SetHorizontalGroup(jPanel3Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel3Layout.CreateSequentialGroup().Add(this.JLabel2).AddContainerGap(448, Short_MAX_VALUE)))
        jPanel3Layout.SetVerticalGroup(jPanel3Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(this.JLabel2))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(org_Jdesktop.Layout.GroupLayout.TRAILING, layout.CreateSequentialGroup().AddContainerGap().Add(this.JPanel3, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddContainerGap()).Add(this.JPanel1, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(layout.CreateSequentialGroup().Add(this.JPanel2, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(layout.CreateSequentialGroup().Add(this.JPanel2, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, 25, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.JPanel3, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.JPanel1, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
}


func (this *MirrorSpriteImageJPanel) GenerateJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}


func (this *MirrorSpriteImageJPanel) ImageHeightIsFrameSizeJCheckBoxActionPerformed(evt *ActionEvent) {
}


func (this *MirrorSpriteImageJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *MirrorSpriteImageJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}

