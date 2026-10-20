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

type InputAutomationTypeJPanel struct {
    JPanel
    InputAutomationTypeJComboBox *JComboBox
    JLabel1 *JLabel
}


func (this *InputAutomationTypeJPanel) initFields() {
    
    
}


func NewInputAutomationTypeJPanel() *InputAutomationTypeJPanel {
    this := &InputAutomationTypeJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *InputAutomationTypeJPanel) GetInputAutomationTypeJComboBox() *JComboBox {
        return this.InputAutomationTypeJComboBox
}


func (this *InputAutomationTypeJPanel) SetInputAutomationTypeJComboBox(inputAutomationTypeJComboBox *JComboBox) {
        this.InputAutomationTypeJComboBox = inputAutomationTypeJComboBox
}


func (this *InputAutomationTypeJPanel) InitComponents() {
        this.InputAutomationTypeJComboBox = NewJComboBox()
        this.JLabel1 = NewJLabel()
        this.InputAutomationTypeJComboBox.SetModel(NewDefaultComboBoxModel([]string{"DirectX", "DirectX 8", "Low Level", "Java Robot"}))
        this.JLabel1.SetText("Input Type:")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddComponent(this.JLabel1).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddComponent(this.InputAutomationTypeJComboBox, 0, 98, Short_MAX_VALUE)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.JLabel1).AddComponent(this.InputAutomationTypeJComboBox, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE)))
}

