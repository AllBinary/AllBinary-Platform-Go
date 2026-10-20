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

type ColorAtActionScriptConditionJPanel struct {
    JPanel
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ColorAtActionScriptConditionInterface ColorAtActionScriptConditionInterface
    BlueJLabel *JLabel
    ColorAtActionJButton *JButton
    ColorAtActionJDialog *JDialog
    ColorAtXJTextField *JTextField
    ColorAtYJTextField *JTextField
    ColorRangeJLabel *JLabel
    GreenJLabel *JLabel
    JLabel1 *JLabel
    LocationJLabel *JLabel
    MaxBlueJTextField *JTextField
    MaxGreenJTextField *JTextField
    MaxJLabel *JLabel
    MaxRedJTextField *JTextField
    MinBlueJTextField *JTextField
    MinGreenJTextField *JTextField
    MinJLabel *JLabel
    MinRedJTextField *JTextField
    OkJButton *JButton
    RedJLabel *JLabel
    TitleJLabel *JLabel
    XJLabel *JLabel
    YJLabel *JLabel
}


func (this *ColorAtActionScriptConditionJPanel) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewColorAtActionScriptConditionJPanel(colorAtActionScriptConditionInterface ColorAtActionScriptConditionInterface) *ColorAtActionScriptConditionJPanel {
    this := &ColorAtActionScriptConditionJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.InitComponents()
        this.ColorAtActionScriptConditionInterface = colorAtActionScriptConditionInterface
        this.Set()
    return this
}


func (this *ColorAtActionScriptConditionJPanel) Set() {
        var colorRangeInterface ColorRangeInterface = this.ColorAtActionScriptConditionInterface.GetColorRangeInterface(); _ = colorRangeInterface
        this.GetMinRedJTextField().SetText(fmt.Sprint(colorRangeInterface.GetMinRed()))
        this.GetMaxRedJTextField().SetText(fmt.Sprint(colorRangeInterface.GetMaxRed()))
        this.GetMinGreenJTextField().SetText(fmt.Sprint(colorRangeInterface.GetMinGreen()))
        this.GetMaxGreenJTextField().SetText(fmt.Sprint(colorRangeInterface.GetMaxGreen()))
        this.GetMinBlueJTextField().SetText(fmt.Sprint(colorRangeInterface.GetMinBlue()))
        this.GetMaxBlueJTextField().SetText(fmt.Sprint(colorRangeInterface.GetMaxBlue()))
        this.GetColorAtXJTextField().SetText(fmt.Sprint(this.ColorAtActionScriptConditionInterface.GetPoint().X))
        this.GetColorAtYJTextField().SetText(fmt.Sprint(this.ColorAtActionScriptConditionInterface.GetPoint().Y))
        this.ColorAtActionScriptConditionInterface.Log()
}


func (this *ColorAtActionScriptConditionJPanel) Update() {
        var colorRangeInterface ColorRangeInterface = this.ColorAtActionScriptConditionInterface.GetColorRangeInterface(); _ = colorRangeInterface
        colorRangeInterface.SetMinRed(int(fmt.Sprint(this.GetMinRedJTextField().GetText())))
        colorRangeInterface.SetMaxRed(int(fmt.Sprint(this.GetMaxRedJTextField().GetText())))
        colorRangeInterface.SetMinGreen(int(fmt.Sprint(this.GetMinGreenJTextField().GetText())))
        colorRangeInterface.SetMaxGreen(int(fmt.Sprint(this.GetMaxGreenJTextField().GetText())))
        colorRangeInterface.SetMinBlue(int(fmt.Sprint(this.GetMinBlueJTextField().GetText())))
        colorRangeInterface.SetMaxBlue(int(fmt.Sprint(this.GetMaxBlueJTextField().GetText())))
        var xString string = this.GetColorAtXJTextField().GetText(); _ = xString
        if !StringValidationUtil_GetInstance().IsEmpty(xString) && StringValidationUtil_GetInstance().IsNumber(xString) {
            this.ColorAtActionScriptConditionInterface.GetPoint().X = func() int { value, _ := strconv.Atoi(xString); return value }()
        }
        var yString string = this.GetColorAtYJTextField().GetText(); _ = yString
        if !StringValidationUtil_GetInstance().IsEmpty(yString) && StringValidationUtil_GetInstance().IsNumber(yString) {
            this.ColorAtActionScriptConditionInterface.GetPoint().Y = func() int { value, _ := strconv.Atoi(xString); return value }()
        }
        this.ColorAtActionScriptConditionInterface.Log()
}


func (this *ColorAtActionScriptConditionJPanel) GetColorAtXJTextField() *JTextField {
        return this.ColorAtXJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetColorAtXJTextField(colorAtXJTextField *JTextField) {
        this.ColorAtXJTextField = colorAtXJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetColorAtYJTextField() *JTextField {
        return this.ColorAtYJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetColorAtYJTextField(colorAtYJTextField *JTextField) {
        this.ColorAtYJTextField = colorAtYJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetMaxBlueJTextField() *JTextField {
        return this.MaxBlueJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetMaxBlueJTextField(maxBlueJTextField *JTextField) {
        this.MaxBlueJTextField = maxBlueJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetMaxGreenJTextField() *JTextField {
        return this.MaxGreenJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetMaxGreenJTextField(maxGreenJTextField *JTextField) {
        this.MaxGreenJTextField = maxGreenJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetMaxRedJTextField() *JTextField {
        return this.MaxRedJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetMaxRedJTextField(maxRedJTextField *JTextField) {
        this.MaxRedJTextField = maxRedJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetMinBlueJTextField() *JTextField {
        return this.MinBlueJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetMinBlueJTextField(minBlueJTextField *JTextField) {
        this.MinBlueJTextField = minBlueJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetMinGreenJTextField() *JTextField {
        return this.MinGreenJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetMinGreenJTextField(minGreenJTextField *JTextField) {
        this.MinGreenJTextField = minGreenJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetMinRedJTextField() *JTextField {
        return this.MinRedJTextField
}


func (this *ColorAtActionScriptConditionJPanel) SetMinRedJTextField(minRedJTextField *JTextField) {
        this.MinRedJTextField = minRedJTextField
}


func (this *ColorAtActionScriptConditionJPanel) GetYJLabel() *JLabel {
        return this.YJLabel
}


func (this *ColorAtActionScriptConditionJPanel) SetYJLabel(yJLabel *JLabel) {
        this.YJLabel = yJLabel
}


func (this *ColorAtActionScriptConditionJPanel) GetColorAtActionJDialog() *JDialog {
        return this.ColorAtActionJDialog
}


func (this *ColorAtActionScriptConditionJPanel) SetColorAtActionJDialog(colorAtActionJDialog *JDialog) {
        this.ColorAtActionJDialog = colorAtActionJDialog
}


func (this *ColorAtActionScriptConditionJPanel) InitComponents() {
        this.ColorAtActionJDialog = NewJDialog()
        this.TitleJLabel = NewJLabel()
        this.LocationJLabel = NewJLabel()
        this.ColorAtXJTextField = NewJTextField()
        this.XJLabel = NewJLabel()
        this.YJLabel = NewJLabel()
        this.ColorAtYJTextField = NewJTextField()
        this.ColorRangeJLabel = NewJLabel()
        this.RedJLabel = NewJLabel()
        this.MinJLabel = NewJLabel()
        this.MaxJLabel = NewJLabel()
        this.GreenJLabel = NewJLabel()
        this.BlueJLabel = NewJLabel()
        this.MinRedJTextField = NewJTextField()
        this.MaxRedJTextField = NewJTextField()
        this.MaxGreenJTextField = NewJTextField()
        this.MinBlueJTextField = NewJTextField()
        this.MaxBlueJTextField = NewJTextField()
        this.OkJButton = NewJButton()
        this.MinGreenJTextField = NewJTextField()
        this.ColorAtActionJButton = NewJButton()
        this.JLabel1 = NewJLabel()
        this.ColorAtActionJDialog.SetMinimumSize(NewDimension(350, 250))
        this.TitleJLabel.SetText("              Color At Action Options            ")
        this.LocationJLabel.SetText("Location:")
        this.XJLabel.SetText("X:")
        this.YJLabel.SetText("Y:")
        this.ColorRangeJLabel.SetText("Color Range")
        this.RedJLabel.SetText("Red:")
        this.MinJLabel.SetText("Minimum")
        this.MaxJLabel.SetText("Maximum")
        this.GreenJLabel.SetText("Green:")
        this.BlueJLabel.SetText("Blue:")
        this.OkJButton.SetText("OK")
        this.OkJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ColorAtActionScriptConditionJPanel) ActionPerformed(evt *ActionEvent) {
                this.OkJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        var colorAtActionJDialogLayout *GroupLayout = NewGroupLayout(this.ColorAtActionJDialog.GetContentPane()); _ = colorAtActionJDialogLayout
        this.ColorAtActionJDialog.GetContentPane().SetLayout(colorAtActionJDialogLayout)
        colorAtActionJDialogLayout.SetHorizontalGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.OkJButton).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING, false).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.BlueJLabel).AddComponent(this.GreenJLabel).AddComponent(this.RedJLabel).AddComponent(this.ColorRangeJLabel)).AddGap(29, 29, 29)).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, colorAtActionJDialogLayout.CreateSequentialGroup().AddComponent(this.LocationJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.XJLabel).AddGap(18, 18, 18))).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.MinBlueJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 40, Short_MAX_VALUE).AddComponent(this.MinGreenJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 40, Short_MAX_VALUE).AddComponent(this.MinJLabel, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddComponent(this.MinRedJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 40, Short_MAX_VALUE).AddComponent(this.ColorAtXJTextField, javax_Swing.GroupLayout.DEFAULT_SIZE, 40, Short_MAX_VALUE)).AddGap(16, 16, 16))).AddGap(6, 6, 6).AddComponent(this.YJLabel).AddGap(24, 24, 24).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.MaxJLabel, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).AddGroup(javax_Swing.GroupLayout.Alignment.LEADING, colorAtActionJDialogLayout.CreateSequentialGroup().AddGap(1, 1, 1).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING, false).AddComponent(this.MaxBlueJTextField, javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.MaxRedJTextField, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 43, Short_MAX_VALUE).AddComponent(this.MaxGreenJTextField, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.PREFERRED_SIZE, 43, javax_Swing.GroupLayout.PREFERRED_SIZE))).AddComponent(this.ColorAtYJTextField, javax_Swing.GroupLayout.Alignment.LEADING, javax_Swing.GroupLayout.DEFAULT_SIZE, 44, Short_MAX_VALUE)).AddGap(22, 22, 22)).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, colorAtActionJDialogLayout.CreateSequentialGroup().AddComponent(this.TitleJLabel).AddGap(33, 33, 33)))))
        colorAtActionJDialogLayout.SetVerticalGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddContainerGap().AddComponent(this.TitleJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddGap(15, 15, 15).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.XJLabel).AddComponent(this.LocationJLabel).AddComponent(this.ColorAtXJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.YJLabel)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.MinJLabel).AddComponent(this.ColorRangeJLabel)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.MinRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.RedJLabel)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.GreenJLabel).AddComponent(this.MinGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(colorAtActionJDialogLayout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.BlueJLabel).AddComponent(this.MinBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 17, Short_MAX_VALUE).AddComponent(this.OkJButton)).AddGroup(colorAtActionJDialogLayout.CreateSequentialGroup().AddGap(14, 14, 14).AddComponent(this.ColorAtYJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.MaxJLabel).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.MaxRedJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.MaxGreenJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.MaxBlueJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE))).AddContainerGap()))
        this.SetPreferredSize(NewDimension(100, 0))
        this.ColorAtActionJButton.SetText("Edit")
        this.ColorAtActionJButton.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ColorAtActionScriptConditionJPanel) ActionPerformed(evt *ActionEvent) {
                this.ColorAtActionJButtonActionPerformed(evt)
        }
*/ NewActionListener())
        this.JLabel1.SetText("Color At:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 15, Short_MAX_VALUE).AddComponent(this.ColorAtActionJButton)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.ColorAtActionJButton)))
}


func (this *ColorAtActionScriptConditionJPanel) OkJButtonActionPerformed(evt *ActionEvent) {
        this.Update()
        this.GetColorAtActionJDialog().SetVisible(false)
}


func (this *ColorAtActionScriptConditionJPanel) ColorAtActionJButtonActionPerformed(evt *ActionEvent) {
        this.GetColorAtActionJDialog().SetVisible(true)
}

