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

type PropertiesJPanel struct {
    JPanel
    RotateJButton *JButton
    CopyVerticalJButton *JButton
    LeftJButton *JButton
    RotateJPanel *JPanel
    AngleJTextField *JTextField
    WorkAreaPropertiesJPanel *JPanel
    ModifierJPanel *JPanel
    UpJButton *JButton
    JPanel8 *JPanel
    DownJButton *JButton
    JPanel7 *JPanel
    RightJButton *JButton
    CopyHorizontalJButton *JButton
}


func (this *PropertiesJPanel) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewPropertiesJPanel() *PropertiesJPanel {
    this := &PropertiesJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *PropertiesJPanel) Set(panel *JPanel) {
        this.WorkAreaPropertiesJPanel.Add(panel)
        this.Repaint()
}


func (this *PropertiesJPanel) InitComponents() {
        this.WorkAreaPropertiesJPanel = NewJPanel()
        this.ModifierJPanel = NewJPanel()
        this.CopyVerticalJButton = NewJButton()
        this.CopyHorizontalJButton = NewJButton()
        this.RotateJPanel = NewJPanel()
        this.RotateJButton = NewJButton()
        this.AngleJTextField = NewJTextField()
        this.JPanel7 = NewJPanel()
        this.UpJButton = NewJButton()
        this.DownJButton = NewJButton()
        this.JPanel8 = NewJPanel()
        this.LeftJButton = NewJButton()
        this.RightJButton = NewJButton()
        this.SetLayout(NewGridLayout(2, 1))
        this.WorkAreaPropertiesJPanel.SetLayout(NewGridLayout())
        this.Add(this.WorkAreaPropertiesJPanel)
        this.ModifierJPanel.SetLayout(NewGridLayout(6, 0))
        this.CopyVerticalJButton.SetFont(NewFont("Dialog", 0, 10))
        this.CopyVerticalJButton.SetText("Copy Vertical")
        this.ModifierJPanel.Add(this.CopyVerticalJButton)
        this.CopyHorizontalJButton.SetFont(NewFont("Dialog", 0, 10))
        this.CopyHorizontalJButton.SetText("Copy Horizontal")
        this.ModifierJPanel.Add(this.CopyHorizontalJButton)
        this.RotateJPanel.SetLayout(NewGridLayout(1, 2))
        this.RotateJButton.SetFont(NewFont("Dialog", 0, 10))
        this.RotateJButton.SetText("Rotate")
        this.RotateJPanel.Add(this.RotateJButton)
        this.AngleJTextField.SetText("Angle")
        this.RotateJPanel.Add(this.AngleJTextField)
        this.ModifierJPanel.Add(this.RotateJPanel)
        this.JPanel7.SetLayout(NewGridLayout(1, 2))
        this.UpJButton.SetFont(NewFont("Dialog", 0, 10))
        this.UpJButton.SetText("Up")
        this.JPanel7.Add(this.UpJButton)
        this.DownJButton.SetFont(NewFont("Dialog", 0, 10))
        this.DownJButton.SetText("Down")
        this.JPanel7.Add(this.DownJButton)
        this.ModifierJPanel.Add(this.JPanel7)
        this.JPanel8.SetLayout(NewGridLayout(1, 2))
        this.LeftJButton.SetFont(NewFont("Dialog", 0, 10))
        this.LeftJButton.SetText("Left")
        this.JPanel8.Add(this.LeftJButton)
        this.RightJButton.SetFont(NewFont("Dialog", 0, 10))
        this.RightJButton.SetText("Right")
        this.JPanel8.Add(this.RightJButton)
        this.ModifierJPanel.Add(this.JPanel8)
        this.Add(this.ModifierJPanel)
}

