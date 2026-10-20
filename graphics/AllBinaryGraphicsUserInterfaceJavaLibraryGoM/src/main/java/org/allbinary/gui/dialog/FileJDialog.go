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

type FileJDialog struct {
    JDialog
    LogUtil *LogUtil
    FinishedListenerInterface FileJDialogListenerInterface
    JFileChooser1 *JFileChooser
}


func (this *FileJDialog) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new FileJDialog
func NewFileJDialog(args ...any) *FileJDialog {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewFileJDialog_1()
        }
        if len(args) == 2 && Java_IsArg[*Frame](args[0], pass) && Java_IsArg[bool](args[1], pass) {
            return NewFileJDialog_2(Java_Arg[*Frame](args[0]), Java_Arg[bool](args[1]))
        }
    }
    panic("No Java overload of new FileJDialog for the arguments")
}
func NewFileJDialog_1() *FileJDialog {
    this := &FileJDialog{}
        this.JDialog = *NewJDialog()
        this.initFields()
        this.InitComponents()
        this.GetJFileChooser1().SetMultiSelectionEnabled(true)
    return this
}


func NewFileJDialog_2(parent *Frame, modal bool) *FileJDialog {
    this := &FileJDialog{}
        this.JDialog = *NewJDialog(parent, modal)
        this.initFields()
        this.InitComponents()
        this.GetJFileChooser1().SetMultiSelectionEnabled(true)
    return this
}


func (this *FileJDialog) AddFinishedListener(finishedListenerInterface FileJDialogListenerInterface) {
        this.FinishedListenerInterface = finishedListenerInterface
}


func (this *FileJDialog) InitComponents() {
        this.JFileChooser1 = NewJFileChooser()
        this.SetDefaultCloseOperation(javax_Swing.WindowConstants.DISPOSE_ON_CLOSE)
        this.JFileChooser1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *FileJDialog) ActionPerformed(evt *ActionEvent) {
                this.JFileChooser1ActionPerformed(evt)
        }
*/ NewActionListener())
        var layout *GroupLayout = NewGroupLayout(this.GetContentPane()); _ = layout
        this.GetContentPane().SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(org_Jdesktop.Layout.GroupLayout.TRAILING, layout.CreateSequentialGroup().AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE).Add(this.JFileChooser1, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddContainerGap()))
        layout.SetVerticalGroup(layout.CreateParallelGroup(org_Jdesktop.Layout.GroupLayout.LEADING).Add(layout.CreateSequentialGroup().AddContainerGap().Add(this.JFileChooser1, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE, 357, org_Jdesktop.Layout.GroupLayout.PREFERRED_SIZE).AddContainerGap(org_Jdesktop.Layout.GroupLayout.DEFAULT_SIZE, Short_MAX_VALUE)))
        this.Pack()
}


func (this *FileJDialog) JFileChooser1ActionPerformed(evt *ActionEvent) {
        this.LogUtil.PutF("Starting Action Command: " + fmt.Sprint(evt.GetActionCommand()), this, "ActionPerformed")
        var files []*File = this.GetJFileChooser1().GetSelectedFiles(); _ = files
        if evt.GetActionCommand() == this.GetJFileChooser1().CANCEL_SELECTION {
            this.SetVisible(false)
        }
        if evt.GetActionCommand() == this.GetJFileChooser1().APPROVE_SELECTION {
            if files != nil && len(files) > 0 {
                this.FinishedListenerInterface.OnFiles(files)
            }
        }
}


func FileJDialog_Main(args []string) {
        java_Awt.EventQueue.InvokeLater(/* Java anonymous class body: implement it as a named type
            

        func (this *FileJDialog) Run() {
                NewFileJDialog(NewJFrame(), true).SetVisible(true)
        }
*/ NewARunnable())
}


func (this *FileJDialog) GetJFileChooser1() *JFileChooser {
        return this.JFileChooser1
}


func (this *FileJDialog) SetJFileChooser1(jFileChooser1 *JFileChooser) {
        this.JFileChooser1 = jFileChooser1
}

