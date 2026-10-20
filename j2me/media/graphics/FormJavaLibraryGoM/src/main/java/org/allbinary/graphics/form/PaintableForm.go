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

type PaintableForm struct {
    ABCustomForm
    CanvasStrings *CanvasStrings
    Rectangle *Rectangle
    X int
    Y int
    FormType *FormType
}

// The Java class PaintableForm has subclasses so its variables are this interface
type IPaintableForm interface {
    IABCustomForm
    PaintableForm_() *PaintableForm
    Init(rectangle *Rectangle, formType *FormType)
    GetWidth() int
}

func (this *PaintableForm) PaintableForm_() *PaintableForm {
    return this
}

var PaintableForm_NULL_PAINTABLE_FORM any = NullUtil_GetInstance().NULL_OBJECT

func (this *PaintableForm) initFields() {
    this.CanvasStrings = CanvasStrings_GetInstance()
    this.Rectangle = RectangleFactory_SINGLETON
    
    
    this.FormType = FormTypeFactory_GetInstance().NULL_FORM_TYPE
}


func PaintableForm_GetNullPaintableForm() IPaintableForm {
        if PaintableForm_NULL_PAINTABLE_FORM == NullUtil_GetInstance().NULL_OBJECT {
            PaintableForm_NULL_PAINTABLE_FORM = NewPaintableForm(StringUtil_GetInstance().EMPTY_STRING, make([]IABCustomItem, 0), BasicColorFactory_GetInstance().BLACK, BasicColorFactory_GetInstance().WHITE)
        }
        return any(PaintableForm_NULL_PAINTABLE_FORM).(IPaintableForm)
}


func NewPaintableForm(title string, items []IABCustomItem, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *PaintableForm {
    this := &PaintableForm{}
        this.ABCustomForm = *NewABCustomForm(title, items, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.SetSelectedIndex(0)
    return this
}


// throws Exception
func (this *PaintableForm) Init(rectangle *Rectangle, formType *FormType) {
        this.Rectangle = rectangle
        var point IGPoint = this.Rectangle.GetPoint(); _ = point
        this.X = point.GetX()
        this.Y = point.GetY()
        this.FormType = formType
}


func (this *PaintableForm) GetWidth() int {
        return this.Rectangle.GetWidth()
}

