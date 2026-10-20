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

type MirrorImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageStrings *ImageStrings
    ImageProcessorInput *ImageProcessorInput
    Result *BufferedImage
    ButtonGroup *ButtonGroup
    GenerateJButton *JButton
    HorizontalJCheckBox *JCheckBox
    JLabel2 *JLabel
    JLayeredPane1 *JLayeredPane
    JPanel1 *JPanel
    JPanel2 *JPanel
    JPanel3 *JPanel
    RenameJRadioButton *JRadioButton
    VerticleJCheckBox *JCheckBox
    WriteOverOriginalJRadioButton *JRadioButton
}

// Implements ImageProcessorInputCompositeInterface


func (this *MirrorImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageStrings = ImageStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewMirrorImageJPanel(imageProcessorInput *ImageProcessorInput) *MirrorImageJPanel {
    this := &MirrorImageJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.LogUtil.PutF("Starting", this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
    return this
}


func (this *MirrorImageJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *MirrorImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var imageProcessorInput *ImageProcessorInput = this.GetImageProcessorInput(); _ = imageProcessorInput
                    var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
                    for index := int(0); index < len(bufferedImageArray); index++ {
                        this.Result = MirrorImageUtil_GetInstance().GetImage(bufferedImageArray[index], this.VerticleJCheckBox.IsSelected(), this.HorizontalJCheckBox.IsSelected())
                        var file *File = imageProcessorInput.GetFiles()[index]; _ = file
                        if this.WriteOverOriginalJRadioButton.IsSelected() {
                        } else if this.RenameJRadioButton.IsSelected() {
                            var filePath string = file.GetAbsolutePath(); _ = filePath
                            var extensionIndex int = strings.Index(filePath, string(this.ImageStrings.PNG_EXTENSION)); _ = extensionIndex
                            filePath = filePath[0:extensionIndex] + "_mirror" + fmt.Sprint(this.ImageStrings.PNG_EXTENSION)
                            logUtil.PutF("Renamed File: " + filePath, this, this.CommonStrings.RUN)
                            file = NewFile(filePath)
                        }
                        var isWritten bool = ImageIO_Write(any(this.Result).(*RenderedImage), this.ImageStrings.PNG, file); _ = isWritten
                        logUtil.PutF("File: " + fmt.Sprint(file) + " Wrote: " + fmt.Sprint(isWritten), this, this.CommonStrings.RUN)
                        this.GetParent().Repaint()
                    }
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *MirrorImageJPanel) InitComponents() {
        this.JLayeredPane1 = NewJLayeredPane()
        this.ButtonGroup = NewButtonGroup()
        this.JPanel1 = /* Java anonymous class body: implement it as a named type
            

        func (this *MirrorImageJPanel) Paint(graphics IGraphics) {
                if this.Result != nil {
                    graphics.DrawImage(this.Result, 0, 0, this.Result.GetWidth(nil), this.Result.GetHeight(nil), nil)
                }
        }
*/ NewJPanel()
        this.JPanel2 = NewJPanel()
        this.HorizontalJCheckBox = NewJCheckBox()
        this.VerticleJCheckBox = NewJCheckBox()
        this.GenerateJButton = NewJButton()
        this.RenameJRadioButton = NewJRadioButton()
        this.WriteOverOriginalJRadioButton = NewJRadioButton()
        this.JPanel3 = NewJPanel()
        this.JLabel2 = NewJLabel()
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(0, 522, Short_MAX_VALUE))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(0, 230, Short_MAX_VALUE))
        this.HorizontalJCheckBox.SetSelected(true)
        this.HorizontalJCheckBox.SetText("Horizontal")
        this.HorizontalJCheckBox.SetEnabled(false)
        this.VerticleJCheckBox.SetText("Verticle")
        this.VerticleJCheckBox.SetEnabled(false)
        this.GenerateJButton.SetText("Generate")
        this.GenerateJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MirrorImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.GenerateJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.RenameJRadioButton.SetText("Save As filename_mirror")
        this.RenameJRadioButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *MirrorImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.RenameJRadioButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.WriteOverOriginalJRadioButton.SetText("Write Over Original")
        var jPanel2Layout *GroupLayout = NewGroupLayout(this.JPanel2); _ = jPanel2Layout
        this.JPanel2.SetLayout(jPanel2Layout)
        jPanel2Layout.SetHorizontalGroup(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel2Layout.CreateSequentialGroup().AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(this.VerticleJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.UNRELATED).Add(this.HorizontalJCheckBox).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.WriteOverOriginalJRadioButton).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.UNRELATED).Add(this.RenameJRadioButton).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.GenerateJButton).Add(37, 37, 37)))
        jPanel2Layout.SetVerticalGroup(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel2Layout.CreateSequentialGroup().Add(jPanel2Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.BASELINE).Add(this.HorizontalJCheckBox).Add(this.VerticleJCheckBox).Add(this.WriteOverOriginalJRadioButton).Add(this.RenameJRadioButton).Add(this.GenerateJButton)).AddContainerGap()))
        this.JLabel2.SetText("Results:")
        var jPanel3Layout *GroupLayout = NewGroupLayout(this.JPanel3); _ = jPanel3Layout
        this.JPanel3.SetLayout(jPanel3Layout)
        jPanel3Layout.SetHorizontalGroup(jPanel3Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(jPanel3Layout.CreateSequentialGroup().Add(this.JLabel2).AddContainerGap(433, Short_MAX_VALUE)))
        jPanel3Layout.SetVerticalGroup(jPanel3Layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(this.JLabel2))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(org_Jdesktop.Layout.GroupLayout.TRAILING, layout.CreateSequentialGroup().AddContainerGap().Add(this.JPanel3, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddContainerGap()).Add(this.JPanel1, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(layout.CreateSequentialGroup().Add(this.JPanel2, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(layout.CreateSequentialGroup().Add(this.JPanel2, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, 25, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.JPanel3, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddPreferredGap(org_Jdesktop.Layout.LayoutStyle.RELATED).Add(this.JPanel1, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
}


func (this *MirrorImageJPanel) GenerateJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}


func (this *MirrorImageJPanel) RenameJRadioButtonActionPerformed(evt *ActionEvent) {
}


func (this *MirrorImageJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *MirrorImageJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}

