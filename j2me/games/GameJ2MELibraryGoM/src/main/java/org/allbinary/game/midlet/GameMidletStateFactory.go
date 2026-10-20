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

type GameMidletStateFactory struct {
    LogUtil *LogUtil
    CurrentGameState *GameState
}

var GameMidletStateFactory_INSTANCE *GameMidletStateFactory = NewGameMidletStateFactory()

func (this *GameMidletStateFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CurrentGameState = GameStateFactory_GetInstance().NO_GAME_STATE
}


func NewGameMidletStateFactory() *GameMidletStateFactory {
    this := &GameMidletStateFactory{}
    this.initFields()
    return this
}


func GameMidletStateFactory_GetInstance() *GameMidletStateFactory {
        return GameMidletStateFactory_INSTANCE
}


func (this *GameMidletStateFactory) GetCurrentGameState() *GameState {
        return this.CurrentGameState
}


func (this *GameMidletStateFactory) SetCurrentGameState(currentGameState *GameState) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().CURRENT).Append(" GameState: ").Append(StringUtil_GetInstance().ToString(currentGameState))), this, "setCurrentGameState")
        this.CurrentGameState = currentGameState
}

