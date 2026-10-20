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

type CanvasImageJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonSeps *CommonSeps
    ImageStrings *ImageStrings
    ImageProcessorInput *ImageProcessorInput
    Result *BufferedImage
    GenerateJButton *JButton
    HeightReductionTextField *JTextField
    HeightReductionTextField1 *JTextField
    JLabel4 *JLabel
    JLabel5 *JLabel
    JLabel6 *JLabel
    JLabel7 *JLabel
    JLabel8 *JLabel
    JPanel1 *JPanel
    JPanel2 *JPanel
    WidthReductionTextField *JTextField
    WidthReductionTextField1 *JTextField
    WriteOverOriginalJCheckBox *JCheckBox
}

// Implements ImageProcessedVisitor


func (this *CanvasImageJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.ImageStrings = ImageStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewCanvasImageJPanel(imageProcessorInput *ImageProcessorInput) *CanvasImageJPanel {
    this := &CanvasImageJPanel{}
        this.JPanel = *NewJPanel()
        this.initFields()
        this.InitComponents()
        this.ImageProcessorInput = imageProcessorInput
    return this
}


func (this *CanvasImageJPanel) Process() {
        /* Java anonymous class body: implement it as a named type
            

        func (this *CanvasImageJPanel) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                        }
                    }()
                    var canvasTrimUtil *CanvasTrimUtil = CanvasTrimUtil_GetInstance(); _ = canvasTrimUtil
                    var canvasExpandUtil *CanvasExpandUtil = CanvasExpandUtil_GetInstance(); _ = canvasExpandUtil
                    var leftReduction int = fmt.Sprint(this.WidthReductionTextField1.GetText()); _ = leftReduction
                    var topReduction int = fmt.Sprint(this.HeightReductionTextField1.GetText()); _ = topReduction
                    var widthReduction int = fmt.Sprint(this.WidthReductionTextField.GetText()); _ = widthReduction
                    var heightReduction int = fmt.Sprint(this.HeightReductionTextField.GetText()); _ = heightReduction
                    if leftReduction <= 0 || topReduction <= 0 || widthReduction <= 0 || heightReduction <= 0 {
                        canvasTrimUtil.Process(this.GetImageProcessorInput(), -leftReduction, -topReduction, -widthReduction, -heightReduction, this)
                    } else {
                        canvasExpandUtil.Process(this.GetImageProcessorInput(), leftReduction, topReduction, this)
                    }
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewThread().Start()
}


func (this *CanvasImageJPanel) InitComponents() {
        this.WriteOverOriginalJCheckBox = NewJCheckBox()
        this.JPanel1 = /* Java anonymous class body: implement it as a named type
            

        func (this *CanvasImageJPanel) Paint(graphics IGraphics) {
                if this.Result != nil {
                    graphics.DrawImage(this.Result, 0, 0, this.Result.GetWidth(nil), this.Result.GetHeight(nil), nil)
                }
        }
*/ NewJPanel()
        this.GenerateJButton = NewJButton()
        this.JLabel4 = NewJLabel()
        this.JPanel2 = NewJPanel()
        this.JLabel5 = NewJLabel()
        this.WidthReductionTextField = NewJTextField()
        this.JLabel6 = NewJLabel()
        this.HeightReductionTextField = NewJTextField()
        this.WidthReductionTextField1 = NewJTextField()
        this.JLabel7 = NewJLabel()
        this.JLabel8 = NewJLabel()
        this.HeightReductionTextField1 = NewJTextField()
        this.WriteOverOriginalJCheckBox.SetSelected(true)
        this.WriteOverOriginalJCheckBox.SetText("Write Over Original")
        this.WriteOverOriginalJCheckBox.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *CanvasImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.WriteOverOriginalJCheckBoxActionPerformed(evt)
        }
*/ NewActionListener())
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 0, Short_MAX_VALUE))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGap(0, 196, Short_MAX_VALUE))
        this.GenerateJButton.SetText("Generate")
        this.GenerateJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *CanvasImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.GenerateJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel4.SetText("Results:")
        this.JLabel5.SetText("Right:")
        this.WidthReductionTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.WidthReductionTextField.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *CanvasImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.WidthReductionTextFieldActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel6.SetText("Bottom")
        this.HeightReductionTextField.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.WidthReductionTextField1.SetText(CommonPhoneStrings_GetInstance().ZERO)
        this.WidthReductionTextField1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *CanvasImageJPanel) ActionPerformed(evt *ActionEvent) {
                this.WidthReductionTextField1ActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel7.SetText("Top")
        this.JLabel8.SetText("Left:")
        this.HeightReductionTextField1.SetText(CommonPhoneStrings_GetInstance().ZERO)
        var jPanel2Layout *GroupLayout = NewGroupLayout(this.JPanel2); _ = jPanel2Layout
        this.JPanel2.SetLayout(jPanel2Layout)
        jPanel2Layout.SetHorizontalGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel2Layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.JLabel8).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.WidthReductionTextField1, javax_Swing.GroupLayout.PREFERRED_SIZE, 41, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel5).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.WidthReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 41, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel7).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.HeightReductionTextField1, javax_Swing.GroupLayout.PREFERRED_SIZE, 46, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel6).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.HeightReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 46, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap(140, Short_MAX_VALUE)))
        jPanel2Layout.SetVerticalGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, jPanel2Layout.CreateSequentialGroup().AddGap(0, 16, Short_MAX_VALUE).AddGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel6).AddComponent(this.HeightReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel8).AddComponent(this.WidthReductionTextField1, javax_Swing.GroupLayout.PREFERRED_SIZE, 22, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel7).AddComponent(this.HeightReductionTextField1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddGroup(jPanel2Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel5).AddComponent(this.WidthReductionTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, 22, javax_Swing.GroupLayout.PREFERRED_SIZE)))))
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JPanel2, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddGap(78, 78, 78)).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JLabel4).AddGroup(layout.CreateSequentialGroup().AddComponent(this.WriteOverOriginalJCheckBox).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.GenerateJButton))).AddGap(0, 0, Short_MAX_VALUE))).AddContainerGap()))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JPanel2, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.GenerateJButton).AddComponent(this.WriteOverOriginalJCheckBox)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JLabel4).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.JPanel1, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
}


func (this *CanvasImageJPanel) WriteOverOriginalJCheckBoxActionPerformed(evt *ActionEvent) {
}


func (this *CanvasImageJPanel) GenerateJButtonActionPerformed(evt *ActionEvent) {
        this.Process()
}


func (this *CanvasImageJPanel) WidthReductionTextFieldActionPerformed(evt *ActionEvent) {
}


func (this *CanvasImageJPanel) WidthReductionTextField1ActionPerformed(evt *ActionEvent) {
}


func (this *CanvasImageJPanel) GetImageProcessorInput() *ImageProcessorInput {
        return this.ImageProcessorInput
}


func (this *CanvasImageJPanel) SetImageProcessorInput(imageProcessorInput *ImageProcessorInput) {
        this.ImageProcessorInput = imageProcessorInput
}


// throws IOException
func (this *CanvasImageJPanel) Visit(generatedBufferedImage *BufferedImage, name string, index int) {
        this.LogUtil.PutF(this.CommonStrings.VISIT, this, this.CommonStrings.VISIT)
        this.Result = generatedBufferedImage
        this.GetParent().Repaint()
        var file *File = this.ImageProcessorInput.GetFiles()[index]; _ = file
        if !this.WriteOverOriginalJCheckBox.IsSelected() {
            var filePath string = file.GetAbsolutePath(); _ = filePath
            var extensionIndex int = strings.Index(filePath, string(this.ImageStrings.PNG_EXTENSION)); _ = extensionIndex
            filePath = fmt.Sprint(NewStringMaker().Append(filePath[0:extensionIndex]).Append(this.CommonSeps.UNDERSCORE).Append(name).Append(this.ImageStrings.PNG_EXTENSION))
            file = NewFile(filePath)
        }
        var isWritten bool = ImageIO_Write(any(this.Result).(*RenderedImage), this.ImageStrings.PNG, file); _ = isWritten
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("File: ").Append(StringUtil_GetInstance().ToString(file)).Append(" Wrote: ").Appendboolean(isWritten)), this, this.CommonStrings.RUN)
}

