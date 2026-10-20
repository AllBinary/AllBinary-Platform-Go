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

type NullCanvas struct {
    Canvas
}

var NullCanvas_NULL_CANVAS *NullCanvas = NewNullCanvas()


func NewNullCanvas() *NullCanvas {
    this := &NullCanvas{}
    this.Canvas = *NewCanvas()
    return this
}


func (this *NullCanvas) SetFullScreenMode(mode bool) {
}


func (this *NullCanvas) SizeChanged(w int, h int) {
}


func (this *NullCanvas) GetCommandStack() *Stack {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullCanvas) IsCommand(command ICommand) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullCanvas) AddCommand(command ICommand) {
}


func (this *NullCanvas) RemoveCommand(command ICommand) {
}


func (this *NullCanvas) RemoveAllCommands() {
}


func (this *NullCanvas) SetCommandListener(l CommandListener) {
}


func (this *NullCanvas) GetCustomCommandListener() CommandListener {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullCanvas) IsPaused() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullCanvas) RemovePauseCommand() {
}


func (this *NullCanvas) Pause() {
}


func (this *NullCanvas) UnPause() {
}


// throws Exception
func (this *NullCanvas) Process() {
}


func (this *NullCanvas) Paint(graphics IGraphics) {
}


func (this *NullCanvas) Destroy() {
}


func (this *NullCanvas) SetPaused(isPaused bool) {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullCanvas) PointerDragged(x int, y int) {
}


func (this *NullCanvas) PointerPressed(x int, y int) {
}


func (this *NullCanvas) PointerReleased(x int, y int) {
}

