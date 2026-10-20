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

type BlankModuleConfigurationJPanel struct {
    JPanel
    JLabel1 *JLabel
}


func (this *BlankModuleConfigurationJPanel) initFields() {
    
}


func NewBlankModuleConfigurationJPanel() *BlankModuleConfigurationJPanel {
    this := &BlankModuleConfigurationJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *BlankModuleConfigurationJPanel) InitComponents() {
        this.JLabel1 = NewJLabel()
        this.JLabel1.SetText("No Configuration Available")
        var layout *GroupLayout = NewGroupLayout(this); _ = layout
        this.SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(javax_Swing.GroupLayout.Alignment.TRAILING, layout.CreateSequentialGroup().AddContainerGap(231, Short_MAX_VALUE).AddComponent(this.JLabel1).AddGap(225, 225, 225)))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.JLabel1).AddContainerGap(362, Short_MAX_VALUE)))
}

