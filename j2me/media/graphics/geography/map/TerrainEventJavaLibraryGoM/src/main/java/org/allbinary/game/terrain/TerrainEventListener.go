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

type TerrainEventListener struct {
    List IBasicArrayList
}

// Implements TerrainEventListenerInterface


func (this *TerrainEventListener) initFields() {
    this.List = NewBasicArrayListD()
}

func NewTerrainEventListener() *TerrainEventListener {
    this := &TerrainEventListener{}
    this.initFields()
    return this
}


func (this *TerrainEventListener) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *TerrainEventListener) OnTerrainEvent(terrainEvent *TerrainEvent) {
        this.List.Clear()
        this.List.Add(terrainEvent)
}


func (this *TerrainEventListener) GetList() IBasicArrayList {
        return this.List
}

