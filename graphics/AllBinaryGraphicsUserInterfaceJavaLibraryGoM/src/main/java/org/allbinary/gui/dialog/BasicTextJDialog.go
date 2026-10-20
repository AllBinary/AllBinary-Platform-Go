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

type BasicTextJDialog struct {
    JDialog
    CloseListenerInterface CloseJDialogListenerInterface
    JButton1 *JButton
    TextJLabel *JLabel
}


func (this *BasicTextJDialog) initFields() {
    
    
    
}




// Java overloads of new BasicTextJDialog
func NewBasicTextJDialog(args ...any) *BasicTextJDialog {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[*Frame](args[0], pass) && Java_IsArg[bool](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewBasicTextJDialog_1(Java_Arg[*Frame](args[0]), Java_Arg[bool](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewBasicTextJDialog_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new BasicTextJDialog for the arguments")
}
func NewBasicTextJDialog_1(parent *Frame, modal bool, text string) *BasicTextJDialog {
    this := &BasicTextJDialog{}
        this.JDialog = *NewJDialog(parent, modal)
        this.initFields()
        this.InitComponents()
        this.SetText(text)
    return this
}


func (this *BasicTextJDialog) SetText(text string) {
        this.TextJLabel.SetText(text)
}


func (this *BasicTextJDialog) AddCloseListener(closeListenerInterface CloseJDialogListenerInterface) {
        this.CloseListenerInterface = closeListenerInterface
}


func NewBasicTextJDialog_2(text string) *BasicTextJDialog {
    this := &BasicTextJDialog{}
    this.JDialog = *NewJDialog()
    this.initFields()
        this.InitComponents()
        this.TextJLabel.SetText(text)
    return this
}


func (this *BasicTextJDialog) InitComponents() {
        this.JButton1 = NewJButton()
        this.TextJLabel = NewJLabel()
        this.SetDefaultCloseOperation(javax_Swing.WindowConstants.DISPOSE_ON_CLOSE)
        this.JButton1.SetText("OK")
        this.JButton1.AddActionListener(/* Java anonymous class body: implement it as a named type
            

        func (this *BasicTextJDialog) ActionPerformed(evt *ActionEvent) {
                this.JButton1ActionPerformed(evt)
        }
*/ NewActionListener())
        this.TextJLabel.SetText("No Text Set")
        var layout *GroupLayout = NewGroupLayout(this.GetContentPane()); _ = layout
        this.GetContentPane().SetLayout(layout)
        layout.SetHorizontalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddGap(80, 80, 80).AddComponent(this.JButton1)).AddGroup(layout.CreateSequentialGroup().AddContainerGap().AddComponent(this.TextJLabel, javax_Swing.GroupLayout.DEFAULT_SIZE, 184, Short_MAX_VALUE))).AddContainerGap()))
        layout.SetVerticalGroup(layout.CreateParallelGroup(javax_Swing.GroupLayout.Alignment.LEADING).AddGroup(layout.CreateSequentialGroup().AddGap(20, 20, 20).AddComponent(this.TextJLabel, javax_Swing.GroupLayout.PREFERRED_SIZE, 20, javax_Swing.GroupLayout.PREFERRED_SIZE).AddGap(30, 30, 30).AddComponent(this.JButton1)))
        this.Pack()
}


func (this *BasicTextJDialog) JButton1ActionPerformed(evt *ActionEvent) {
        this.SetVisible(false)
        if this.CloseListenerInterface != nil {
            this.CloseListenerInterface.OnClose()
        }
}


func BasicTextJDialog_Main(args []string) {
        java_Awt.EventQueue.InvokeLater(/* Java anonymous class body: implement it as a named type
            

        func (this *BasicTextJDialog) Run() {
                NewBasicTextJDialog(NewJFrame(), true, "Test").SetVisible(true)
        }
*/ NewARunnable())
}

