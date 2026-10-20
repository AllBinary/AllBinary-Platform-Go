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

type ViewSwapper struct {
    RootViewId int
    Activity *Activity
    CurrentMainViewId int
    MainViewArray []*View
    TopViewArray []*View
}

var ViewSwapper_NULL_VIEW_ARRAY []*View = make([]*View, 0)

func (this *ViewSwapper) initFields() {
    
    
    this.CurrentMainViewId = -1
    this.MainViewArray = ViewSwapper_NULL_VIEW_ARRAY
    this.TopViewArray = ViewSwapper_NULL_VIEW_ARRAY
}


func NewViewSwapper(activity *Activity, rootViewId int) *ViewSwapper {
    this := &ViewSwapper{}
    this.initFields()
        this.Activity = activity
        this.RootViewId = rootViewId
    return this
}


func (this *ViewSwapper) SetMainViews(viewArray []*View) {
        this.MainViewArray = viewArray
}


func (this *ViewSwapper) SetTopViewArray(topViewArray []*View) {
        this.TopViewArray = topViewArray
}


func (this *ViewSwapper) SetMainView(id int) {
        var viewGroup *ViewGroup = any(this.Activity.FindViewById(this.RootViewId)).(*ViewGroup); _ = viewGroup
        if id != this.CurrentMainViewId {
            for index := int(0); index < len(this.TopViewArray); index++ {
                viewGroup.RemoveView(this.TopViewArray[index])
            }
            for index := int(0); index < len(this.MainViewArray); index++ {
                if id != this.MainViewArray[index].GetId() {
                    viewGroup.RemoveView(this.MainViewArray[index])
                }
            }
            for index := int(0); index < len(this.MainViewArray); index++ {
                if id == this.MainViewArray[index].GetId() {
                    viewGroup.AddView(this.MainViewArray[index])
                    this.CurrentMainViewId = id
                }
            }
            for index := int(0); index < len(this.TopViewArray); index++ {
                viewGroup.AddView(this.TopViewArray[index])
            }
        }
}

