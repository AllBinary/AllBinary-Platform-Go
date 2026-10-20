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

type RotateJDialog struct {
    JDialog
    LogUtil *LogUtil
    HashMap *HashMap
    AngleIncrementJTextField *JTextField
    JButton1 *JButton
    JLabel1 *JLabel
    JLabel2 *JLabel
    JPanel1 *JPanel
    TotalRotationJTextField *JTextField
}


func (this *RotateJDialog) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.HashMap = StdUtil_GetInstance().CreateHashMap()
    
    
    
    
    
    
}


func NewRotateJDialog(parent *Frame, modal bool) *RotateJDialog {
    this := &RotateJDialog{}
        this.JDialog = *NewJDialog(parent, modal)
        this.initFields()
        this.InitComponents()
    return this
}


func (this *RotateJDialog) InitComponents() {
        this.AngleIncrementJTextField = NewJTextField()
        this.JLabel1 = NewJLabel()
        this.TotalRotationJTextField = NewJTextField()
        this.JLabel2 = NewJLabel()
        this.JPanel1 = NewJPanel()
        this.JButton1 = NewJButton()
        this.SetDefaultCloseOperation(javax_Swing.WindowConstants.DISPOSE_ON_CLOSE)
        this.AngleIncrementJTextField.SetText("36")
        this.JLabel1.SetText("Angle Increment:")
        this.TotalRotationJTextField.SetText("360")
        this.JLabel2.SetText("Total Angular Rotation:")
        this.JButton1.SetText("OK")
        this.JButton1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *RotateJDialog) ActionPerformed(evt *ActionEvent) {
                this.JButton1ActionPerformed(evt)
        }
*/ NewActionListener())
        var jPanel1Layout *GroupLayout = NewGroupLayout(this.JPanel1); _ = jPanel1Layout
        this.JPanel1.SetLayout(jPanel1Layout)
        jPanel1Layout.SetHorizontalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(jPanel1Layout.CreateSequentialGroup().AddGap(62, 62, 62).AddComponent(this.JButton1).AddContainerGap(62, Short_MAX_VALUE)))
        jPanel1Layout.SetVerticalGroup(jPanel1Layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.JButton1))
        var layout *GroupLayout = NewGroupLayout(this.GetContentPane()); _ = layout
        this.GetContentPane().SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.TRAILING).AddComponent(this.JLabel1).AddComponent(this.JLabel2)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.UNRELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddComponent(this.TotalRotationJTextField, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, 31, Short_MAX_VALUE).AddComponent(this.AngleIncrementJTextField, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.PREFERRED_SIZE, 31, javax_Swing.GroupLayout.PREFERRED_SIZE)).AddContainerGap()).AddComponent(this.JPanel1, javax_Swing.GroupLayout.Alignment.TRAILING, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.AngleIncrementJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel1)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED).AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.BASELINE).AddComponent(this.TotalRotationJTextField, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddComponent(this.JLabel2)).AddPreferredGap(javax_Swing.LayoutStyle.ComponentPlacement.RELATED, 16, Short_MAX_VALUE).AddComponent(this.JPanel1, javax_Swing.GroupLayout.PREFERRED_SIZE, javax_Swing.GroupLayout.DEFAULT_SIZE, javax_Swing.GroupLayout.PREFERRED_SIZE).AddContainerGap()))
        this.Pack()
}


func (this *RotateJDialog) JButton1ActionPerformed(evt *ActionEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.EXCEPTION, this, "jButton1ActionPerformed")
                }
            }()
            this.HashMap.Put(fmt.Sprint(0), fmt.Sprint(this.AngleIncrementJTextField.GetText()))
            this.HashMap.Put(fmt.Sprint(1), fmt.Sprint(this.TotalRotationJTextField.GetText()))
            (any(this.GetParent()).(HashMapJDialogListenerInterface)).OnHashMap(this.HashMap)
            this.SetVisible(false)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func RotateJDialog_Main(args []string) {
        java_Awt.EventQueue.InvokeLater(/* Java anonymous class body: implement it as a named type
            

        func (this *RotateJDialog) Run() {
                var dialog *RotateJDialog = NewRotateJDialog(NewJFrame(), true); _ = dialog
                dialog.AddWindowListener(/* Java anonymous class body: implement it as a named type
                    

                func (this *RotateJDialog) WindowClosing(e *WindowEvent) {
                        os.Exit(0)
                }
* / NewWindowAdapter())
                dialog.SetVisible(true)
        }
*/ NewARunnable())
}

