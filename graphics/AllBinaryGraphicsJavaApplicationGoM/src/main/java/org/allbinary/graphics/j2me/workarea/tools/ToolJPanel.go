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

type ToolJPanel struct {
    JPanel
    ArcJButton *JButton
    CharJButton *JButton
    ColorJList *JList
    ColorJPanel *JPanel
    DottedArcJButton *JButton
    DottedFillArcJButton *JButton
    DottedFillRectJButton *JButton
    DottedFillRountRectJButton *JButton
    DottedLineJButton *JButton
    DottedRectButton *JButton
    FillArcJButton *JButton
    FillRectJButton *JButton
    FillRoundRectJButton *JButton
    ImageJButton *JButton
    LineJButton *JButton
    RectJButton *JButton
    RoundRectJButton *JButton
    StringJButton *JButton
}


func (this *ToolJPanel) initFields() {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
}


func NewToolJPanel() *ToolJPanel {
    this := &ToolJPanel{}
    this.JPanel = *NewJPanel()
    this.initFields()
        this.InitComponents()
    return this
}


func (this *ToolJPanel) InitComponents() {
        this.LineJButton = NewJButton()
        this.RectJButton = NewJButton()
        this.ArcJButton = NewJButton()
        this.StringJButton = NewJButton()
        this.CharJButton = NewJButton()
        this.RoundRectJButton = NewJButton()
        this.ImageJButton = NewJButton()
        this.FillArcJButton = NewJButton()
        this.FillRectJButton = NewJButton()
        this.FillRoundRectJButton = NewJButton()
        this.DottedLineJButton = NewJButton()
        this.DottedRectButton = NewJButton()
        this.DottedArcJButton = NewJButton()
        this.DottedFillRectJButton = NewJButton()
        this.DottedFillArcJButton = NewJButton()
        this.DottedFillRountRectJButton = NewJButton()
        this.ColorJPanel = NewJPanel()
        this.ColorJList = NewJList()
        this.SetLayout(NewGridLayout(17, 1))
        this.LineJButton.SetFont(NewFont("Dialog", 0, 10))
        this.LineJButton.SetText(LinesGraphicItem_GetStaticName())
        this.LineJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MouseClicked(evt IMouseEvent) {
                this.LineJButtonMouseClicked(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.LineJButton)
        this.RectJButton.SetFont(NewFont("Dialog", 0, 10))
        this.RectJButton.SetText("Rect")
        this.RectJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MouseClicked(evt IMouseEvent) {
                this.RectJButtonMouseClicked(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.RectJButton)
        this.ArcJButton.SetFont(NewFont("Dialog", 0, 10))
        this.ArcJButton.SetText("Arc")
        this.ArcJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.ArcJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.ArcJButton)
        this.StringJButton.SetFont(NewFont("Dialog", 0, 10))
        this.StringJButton.SetText("String")
        this.StringJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.StringJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.StringJButton)
        this.CharJButton.SetFont(NewFont("Dialog", 0, 10))
        this.CharJButton.SetText("Char")
        this.CharJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.CharJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.CharJButton)
        this.RoundRectJButton.SetFont(NewFont("Dialog", 0, 10))
        this.RoundRectJButton.SetText("Round Rect")
        this.RoundRectJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.RoundRectJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.RoundRectJButton)
        this.ImageJButton.SetFont(NewFont("Dialog", 0, 10))
        this.ImageJButton.SetText("Image")
        this.ImageJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.ImageJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.ImageJButton)
        this.FillArcJButton.SetFont(NewFont("Dialog", 0, 10))
        this.FillArcJButton.SetText("Fill Arc")
        this.FillArcJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.FillArcJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.FillArcJButton)
        this.FillRectJButton.SetFont(NewFont("Dialog", 0, 10))
        this.FillRectJButton.SetText("Fill Rect")
        this.FillRectJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.FillRectJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.FillRectJButton)
        this.FillRoundRectJButton.SetFont(NewFont("Dialog", 0, 10))
        this.FillRoundRectJButton.SetText("Fill Rnd Rect")
        this.FillRoundRectJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.FillRoundRectJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.FillRoundRectJButton)
        this.DottedLineJButton.SetFont(NewFont("Dialog", 0, 10))
        this.DottedLineJButton.SetText("Dot Line")
        this.DottedLineJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.DottedLineJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.DottedLineJButton)
        this.DottedRectButton.SetFont(NewFont("Dialog", 0, 10))
        this.DottedRectButton.SetText("Dot Rect")
        this.DottedRectButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.DottedRectButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.DottedRectButton)
        this.DottedArcJButton.SetFont(NewFont("Dialog", 0, 10))
        this.DottedArcJButton.SetText("Dot Arc")
        this.DottedArcJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.DottedArcJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.DottedArcJButton)
        this.DottedFillRectJButton.SetFont(NewFont("Dialog", 0, 10))
        this.DottedFillRectJButton.SetText("Dot Fill Rect")
        this.DottedFillRectJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.DottedFillRectJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.DottedFillRectJButton)
        this.DottedFillArcJButton.SetFont(NewFont("Dialog", 0, 10))
        this.DottedFillArcJButton.SetText("Dot Fill Arc")
        this.DottedFillArcJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.DottedFillArcJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.DottedFillArcJButton)
        this.DottedFillRountRectJButton.SetFont(NewFont("Dialog", 0, 10))
        this.DottedFillRountRectJButton.SetText("Dot Fill Rnd Rect")
        this.DottedFillRountRectJButton.AddMouseListener(/* Java anonymous class body: implement it as a named type
            

        func (this *ToolJPanel) MousePressed(evt IMouseEvent) {
                this.DottedFillRountRectJButtonMousePressed(evt)
        }
*/ NewMouseAdapter())
        this.Add(this.DottedFillRountRectJButton)
        this.ColorJPanel.SetLayout(NewGridLayout(1, 1))
        this.ColorJPanel.Add(this.ColorJList)
        this.Add(this.ColorJPanel)
}


func (this *ToolJPanel) DottedFillRountRectJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.DottedFillRountRectJButton.GetText())
}


func (this *ToolJPanel) DottedFillArcJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.DottedFillArcJButton.GetText())
}


func (this *ToolJPanel) DottedFillRectJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.DottedFillRectJButton.GetText())
}


func (this *ToolJPanel) DottedArcJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.DottedArcJButton.GetText())
}


func (this *ToolJPanel) DottedRectButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.DottedRectButton.GetText())
}


func (this *ToolJPanel) DottedLineJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.DottedLineJButton.GetText())
}


func (this *ToolJPanel) FillRoundRectJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.FillRoundRectJButton.GetText())
}


func (this *ToolJPanel) FillRectJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.FillRectJButton.GetText())
}


func (this *ToolJPanel) FillArcJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.FillArcJButton.GetText())
}


func (this *ToolJPanel) ImageJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.ImageJButton.GetText())
}


func (this *ToolJPanel) RoundRectJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.RoundRectJButton.GetText())
}


func (this *ToolJPanel) CharJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.CharJButton.GetText())
}


func (this *ToolJPanel) StringJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.StringJButton.GetText())
}


func (this *ToolJPanel) ArcJButtonMousePressed(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.ArcJButton.GetText())
}


func (this *ToolJPanel) RectJButtonMouseClicked(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.RectJButton.GetText())
}


func (this *ToolJPanel) LineJButtonMouseClicked(evt IMouseEvent) {
        ToolFactory_GetInstance().SetSelectedTool(this.LineJButton.GetText())
}

