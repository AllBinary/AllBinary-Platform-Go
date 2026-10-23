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

type SubscriberUser struct {
    User
    Subscription *Subscription
}


func (this *SubscriberUser) initFields() {
    
}




// Java overloads of new SubscriberUser
func NewSubscriberUser(args ...any) *SubscriberUser {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewSubscriberUser_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewSubscriberUser_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new SubscriberUser for the arguments")
}
func NewSubscriberUser_1() *SubscriberUser {
    this := &SubscriberUser{}
        this.User = *NewUser()
        this.initFields()
        this.Subscription = NewSubscription()
    return this
}


func NewSubscriberUser_2(userHashMap *HashMap) *SubscriberUser {
    this := &SubscriberUser{}
        this.User = *NewUser(userHashMap)
        this.initFields()
        this.Subscription = NewSubscription()
    return this
}


func (this *SubscriberUser) ValidateSession(weblisketSession WeblisketSessionInterface) {
        this.User.ValidateSession(weblisketSession)
        this.UpdateSession(weblisketSession)
}


func (this *SubscriberUser) UpdateSession(weblisketSession WeblisketSessionInterface) {
        weblisketSession.SetStoreName(this.GetPermissions())
}


func (this *SubscriberUser) IsSessionValid() bool {
        return this.Subscription.IsSubscribed()
}

