/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type MultipassWaypointPathRunnable struct {
    WaypointPathRunnableBase
    LogUtil *LogUtil
    BasicArrayListUtil *BasicArrayListUtil
    MultipassState *MultipassState
    Done bool
    List IBasicArrayList
    PathFindingInfo any
    FIRST_RUNNABLE Runnable
    SECOND_RUNNABLE Runnable
    END_RUNNABLE Runnable
    ALREADY_ENDED_RUNNABLE Runnable
    CurrentPassRunnable Runnable
}


func (this *MultipassWaypointPathRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicArrayListUtil = BasicArrayListUtil_GetInstance()
    this.MultipassState = NewMultipassState()
    this.Done = false
    this.List = this.BasicArrayListUtil.GetImmutableInstance()
    this.PathFindingInfo = NullUtil_GetInstance().NULL_OBJECT
    this.FIRST_RUNNABLE = NewFirstRunnable(this)
    this.SECOND_RUNNABLE = NewSecondRunnable(this)
    this.END_RUNNABLE = NewEndRunnable(this)
    this.ALREADY_ENDED_RUNNABLE = /* Java anonymous class body: implement it as a named type
        

    func (this *MultipassWaypointPathRunnable) Run() {
            panic(errors.New(fmt.Sprint()))
    }
*/ NewARunnable()
    this.CurrentPassRunnable = this.FIRST_RUNNABLE
}


func NewMultipassWaypointPathRunnable() *MultipassWaypointPathRunnable {
    this := &MultipassWaypointPathRunnable{}
    this.WaypointPathRunnableBase = *NewWaypointPathRunnableBase()
    this.initFields()
    return this
}


func (this *MultipassWaypointPathRunnable) SetRunning(isRunning bool) {
        this.RunningP = isRunning
        if this.RunningP {
            this.Reset()
            this.Done = false
        }
}


func (this *MultipassWaypointPathRunnable) Run() {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
            this.SetRunning(false)
                }
            }()
            this.CurrentPassRunnable.Run()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MultipassWaypointPathRunnable) Reset2() {
        this.MultipassState.Step = 0
        this.MultipassState.Iteration = 0
        this.MultipassState.Iteration2 = 0
        this.PathFindingInfo = NullUtil_GetInstance().NULL_OBJECT
}


func (this *MultipassWaypointPathRunnable) Finish() {
        this.Reset2()
        this.CurrentPassRunnable = this.ALREADY_ENDED_RUNNABLE
        this.Done = true
}


func (this *MultipassWaypointPathRunnable) IsDone() bool {
        return this.Done
}


func (this *MultipassWaypointPathRunnable) Reset() {
        this.Reset2()
        this.CurrentPassRunnable = this.FIRST_RUNNABLE
        this.Done = false
}


func (this *MultipassWaypointPathRunnable) ProcessFirstRunnable() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
            this.SetRunning(false)
            this.Finish()
                }
            }()
            this.PathFindingLayer.GetWaypointRunnableLogHelper().Start(this.PathFindingLayer)
            this.Reset2()
            var geographicMapCellPosition IGeographicMapCellPosition = this.PathFindingLayer.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
            if geographicMapCellPosition == nil {
                panic(errors.New(fmt.Sprint("Should never be running here")))
            }
            this.PathFindingInfo = this.TargetPathFindingLayer.GetWaypointBehavior().GetWaypoint().GetPathFindingInfo(geographicMapCellPosition)
            var localPathFindingInfo *PathFindingInfo = any(this.PathFindingInfo).(*PathFindingInfo); _ = localPathFindingInfo
            this.List = this.TargetPathFindingLayer.GetWaypointBehavior().GetWaypoint().GetPathsList(geographicMapCellPosition, localPathFindingInfo, this.MultipassState)
            if this.List != this.BasicArrayListUtil.GetImmutableInstance() {
                this.END_RUNNABLE.Run()
            } else {
                this.CurrentPassRunnable = this.SECOND_RUNNABLE
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MultipassWaypointPathRunnable) ProcessSecondRunnable() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
            this.SetRunning(false)
            this.Finish()
                }
            }()
            var geographicMapCellPosition IGeographicMapCellPosition = this.PathFindingLayer.GetCurrentGeographicMapCellPosition(); _ = geographicMapCellPosition
            var localPathFindingInfo *PathFindingInfo = any(this.PathFindingInfo).(*PathFindingInfo); _ = localPathFindingInfo
            this.List = this.TargetPathFindingLayer.GetWaypointBehavior().GetWaypoint().GetPathsList(geographicMapCellPosition, localPathFindingInfo, this.MultipassState)
            if this.List != this.BasicArrayListUtil.GetImmutableInstance() {
                this.END_RUNNABLE.Run()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *MultipassWaypointPathRunnable) ProcessEndRunnable() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
            this.SetRunning(false)
                }
            }()
            var waypointBehavior IWaypointBehaviorBase = this.PathFindingLayer.GetWaypointBehavior(); _ = waypointBehavior
            waypointBehavior.SetWaypointPathsList(this.List)
            this.PathFindingLayer.GetWaypointRunnableLogHelper().End(this.PathFindingLayer)
            return
        }(); tryControl1 == 1 {
            return
        }
        this.Finish()
}

type FirstRunnable struct {
    ARunnable
    MultipassWaypointPathRunnable *MultipassWaypointPathRunnable
}


func (this *FirstRunnable) initFields() {
        
}
    

    func NewMultipassWaypointPathRunnable(multipassWaypointPathRunnable *MultipassWaypointPathRunnable) *MultipassWaypointPathRunnable {
        this := &MultipassWaypointPathRunnable{}
        this.ARunnable = *NewARunnable()
        this.initFields()
            this.MultipassWaypointPathRunnable = multipassWaypointPathRunnable
        return this
    }
    

    func (this *MultipassWaypointPathRunnable) Run() {
            this.MultipassWaypointPathRunnable.ProcessFirstRunnable()
    }

//    private boolean first = true;
type SecondRunnable struct {
    ARunnable
    MultipassWaypointPathRunnable *MultipassWaypointPathRunnable
}


func (this *SecondRunnable) initFields() {
        
}
    

    func NewMultipassWaypointPathRunnable(multipassWaypointPathRunnable *MultipassWaypointPathRunnable) *MultipassWaypointPathRunnable {
        this := &MultipassWaypointPathRunnable{}
        this.ARunnable = *NewARunnable()
        this.initFields()
            this.MultipassWaypointPathRunnable = multipassWaypointPathRunnable
        return this
    }
    

    func (this *MultipassWaypointPathRunnable) Run() {
            this.MultipassWaypointPathRunnable.ProcessSecondRunnable()
    }

type EndRunnable struct {
    ARunnable
    MultipassWaypointPathRunnable *MultipassWaypointPathRunnable
}


func (this *EndRunnable) initFields() {
        
}
    

    func NewMultipassWaypointPathRunnable(multipassWaypointPathRunnable *MultipassWaypointPathRunnable) *MultipassWaypointPathRunnable {
        this := &MultipassWaypointPathRunnable{}
        this.ARunnable = *NewARunnable()
        this.initFields()
            this.MultipassWaypointPathRunnable = multipassWaypointPathRunnable
        return this
    }
    

    func (this *MultipassWaypointPathRunnable) Run() {
            this.MultipassWaypointPathRunnable.ProcessEndRunnable()
    }

