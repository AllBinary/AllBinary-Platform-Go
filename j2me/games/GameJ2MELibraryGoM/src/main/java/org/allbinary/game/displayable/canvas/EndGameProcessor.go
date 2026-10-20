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

type EndGameProcessor struct {
    Processor
    GameStateFactory *GameStateFactory
    GameCanvas IAllBinaryGameCanvas
    WAIT int64
}


func (this *EndGameProcessor) initFields() {
    this.GameStateFactory = GameStateFactory_GetInstance()
    
    this.WAIT = 5000
}


func NewEndGameProcessor(gameCanvas IAllBinaryGameCanvas) *EndGameProcessor {
    this := &EndGameProcessor{}
    this.Processor = *NewProcessor()
    this.initFields()
        this.GameCanvas = gameCanvas
    return this
}


// throws Exception
func (this *EndGameProcessor) Process() {
        if this.GameCanvas.IsHighScoreSubmitted() {
            if this.GameCanvas.GetGameStateTimeHelper().IsElapsedTNT(this.WAIT) {
                if this.GameCanvas.GetGameState() == this.GameStateFactory.SHOW_END_RESULT_GAME_STATE {
                    var highScoresBase IHighScoresHelperBase = any(this.GameCanvas.AllBinaryGameCanvas_().HighScoresHelper).(IHighScoresHelperBase); _ = highScoresBase
                    var highScores IHighScores = highScoresBase.GetSelectedHighScores(); _ = highScores
                    if highScores != NullHighScoresSingletonFactory_GetInstance() {
                        highScoresBase.SelectHighScores()
                        this.GameCanvas.GetRealHighScoresPaintable().SetHighScores(highScores)
                        this.GameCanvas.SetGameState(this.GameStateFactory.SHOW_HIGH_SCORE_GAME_STATE)
                        this.GameCanvas.SetHighScoresPaintable(this.GameCanvas.GetRealHighScoresPaintable())
                    }
                } else if this.GameCanvas.GetGameState() == this.GameStateFactory.SHOW_HIGH_SCORE_GAME_STATE {
                    this.GameCanvas.SetGameState(this.GameStateFactory.SHOW_END_RESULT_GAME_STATE)
                    this.GameCanvas.SetHighScoresPaintable(NullPaintable_GetInstance())
                }
            }
        }
}

