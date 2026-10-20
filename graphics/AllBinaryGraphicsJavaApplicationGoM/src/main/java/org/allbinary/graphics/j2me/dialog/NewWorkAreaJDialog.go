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

type NewWorkAreaJDialog struct {
    JDialog
    Parent *MyFrame
    Dimension *Dimension
    XSizeJTextField *JTextField
    YSizeJTextField *JTextField
    NameJTextField *JTextField
}


func (this *NewWorkAreaJDialog) initFields() {
    
    
    
    
    
}


func NewNewWorkAreaJDialog(parent *MyFrame, modal bool, dimension *Dimension, newName string) *NewWorkAreaJDialog {
    this := &NewWorkAreaJDialog{}
        this.JDialog = *NewJDialog(any(parent).(*Frame), modal)
        this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            this.InitComponents()
            this.Parent = parent
            this.XSizeJTextField = NewJTextField("12")
            this.YSizeJTextField = NewJTextField("12")
            this.NameJTextField = NewJTextField(newName)
            this.Dimension = dimension
            var submitButton *JButton = NewJButton("Ok"); _ = submitButton
            submitButton.AddActionListener(/* Java anonymous class body: implement it as a named type
                

            func (this *NewWorkAreaJDialog) ActionPerformed(evt *ActionEvent) {
                    if tryControl1 := func() (tryControl1 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                            }
                        }()
                        this.DisposeNewDialog()
                        return
                    }(); tryControl1 == 1 {
                        return
                    }
            }
*/ NewActionListener())
            this.SetSize(150, 100)
            this.GetContentPane().SetLayout(NewGridLayout(5, 2))
            this.GetContentPane().Add(NewJLabel("Please enter"))
            this.GetContentPane().Add(NewJLabel(" the size."))
            this.GetContentPane().Add(NewJLabel("Name:"))
            this.GetContentPane().Add(this.NameJTextField)
            this.GetContentPane().Add(NewJLabel("X: "))
            this.GetContentPane().Add(this.XSizeJTextField)
            this.GetContentPane().Add(NewJLabel("Y: "))
            this.GetContentPane().Add(this.YSizeJTextField)
            this.GetContentPane().Add(submitButton)
            this.Show()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// throws Exception
func (this *NewWorkAreaJDialog) DisposeNewDialog() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var canvasWidth int = int(this.XSizeJTextField.GetText()); _ = canvasWidth
            var canvasHeight int = int(this.YSizeJTextField.GetText()); _ = canvasHeight
            var newName string = this.NameJTextField.GetText(); _ = newName
            if newName == "" {
                panic(errors.New(fmt.Sprint("No Name", this, "disposeNewDialog")))
            }
            var workAreaJPanel *WorkAreaJPanel = NewWorkAreaJPanel(newName, this.Dimension, int(canvasWidth), int(canvasHeight)); _ = workAreaJPanel
            workAreaJPanel.SetLayout(NewGridLayout(1, 1))
            this.Parent.Add(any(workAreaJPanel).(WorkAreaJPanelInterface))
            this.Dispose()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *NewWorkAreaJDialog) InitComponents() {
        this.AddWindowListener(/* Java anonymous class body: implement it as a named type
            

        func (this *NewWorkAreaJDialog) WindowClosing(evt *WindowEvent) {
                this.CloseDialog(evt)
        }
*/ NewWindowAdapter())
        this.Pack()
}


func (this *NewWorkAreaJDialog) CloseDialog(evt *WindowEvent) {
        this.SetVisible(false)
        this.Dispose()
}


func NewWorkAreaJDialog_Main(args []string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            NewNewWorkAreaJDialog(NewMyFrame(), true, NewDimension(100, 100), "testing").Show()
            return
        }(); tryControl1 == 1 {
            return
        }
}

