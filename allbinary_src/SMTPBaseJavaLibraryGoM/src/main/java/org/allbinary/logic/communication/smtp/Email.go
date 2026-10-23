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

type Email struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Msg *MimeMessage
    Properties IProperties
    Bs *strings.Builder
    IsDebug bool
}

// Implements EmailInterface

var Email_DEBUG string = "mail.debug"
var Email_SMTP_HOST string = "mail.smtp.host"
var Email_SMTP_PORT string = "mail.smtp.port"
var Email_SMTP_USER string = "mail.smtp.user"
var Email_SMTP_RETURN_ADDRESS string = "mail.from"
var Email_SMTP_LOCAL_HOST string = "mail.smtp.localhost"

func (this *Email) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
}




// Java overloads of new Email
func NewEmail(args ...any) *Email {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 7 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) && Java_IsArg[string](args[5], pass) && Java_IsArg[string](args[6], pass) {
            return NewEmail_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]), Java_Arg[string](args[5]), Java_Arg[string](args[6]))
        }
        if len(args) == 8 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) && Java_IsArg[string](args[5], pass) && Java_IsArg[string](args[6], pass) && Java_IsArg[bool](args[7], pass) {
            return NewEmail_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]), Java_Arg[string](args[5]), Java_Arg[string](args[6]), Java_Arg[bool](args[7]))
        }
    }
    panic("No Java overload of new Email for the arguments")
}
func NewEmail_1(from string, to string, subject string, server string, textBody string, htmlAttachment string, contentBase string) *Email {
    this := &Email{}
    this.initFields()
        this.IsDebug = false
        this.Init(from, to, subject, server, textBody, htmlAttachment, contentBase)
    return this
}


func NewEmail_2(from string, to string, subject string, server string, textBody string, htmlAttachment string, contentBase string, isDebug bool) *Email {
    this := &Email{}
    this.initFields()
        this.IsDebug = isDebug
        this.Init(from, to, subject, server, textBody, htmlAttachment, contentBase)
    return this
}


// throws Exception
func (this *Email) Init_1(from string, to string, subject string, server string, textBody string, htmlAttachment string, contentBase string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "emailConstructor", e)
            }
            panic(e)
                }
            }()
            var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
            var mimeBodyParts []*MimeBodyPart = make([]*MimeBodyPart, func() int { if !stringValidationUtil.IsEmpty(htmlAttachment) { return 2 }; return 1 }()); _ = mimeBodyParts
            mimeBodyParts[0] = NewMimeBodyPart()
            mimeBodyParts[0].SetText(textBody)
            if !stringValidationUtil.IsEmpty(htmlAttachment) {
                var internetHeaders *InternetHeaders = NewInternetHeaders(); _ = internetHeaders
                internetHeaders.AddHeader("Content-Type", "text/html")
                if !stringValidationUtil.IsEmpty(contentBase) {
                    internetHeaders.AddHeader("Content-Base", contentBase)
                }
                mimeBodyParts[1] = NewMimeBodyPart(internetHeaders, Java_GetBytes(htmlAttachment))
            }
            this.Init(server, any(nil).(Authenticator), []*InternetAddress{NewInternetAddress(from)}, []*InternetAddress{NewInternetAddress(to)}, any(nil).([]*InternetAddress), any(nil).([]*InternetAddress), subject, mimeBodyParts)
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of Email.init
func (this *Email) Init(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 7 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[string](args[4], pass) && Java_IsArg[string](args[5], pass) && Java_IsArg[string](args[6], pass) {
            this.Init_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]), Java_Arg[string](args[3]), Java_Arg[string](args[4]), Java_Arg[string](args[5]), Java_Arg[string](args[6]))
            return
        }
        if len(args) == 8 && Java_IsArg[string](args[0], pass) && Java_IsArg[Authenticator](args[1], pass) && Java_IsArg[[]*InternetAddress](args[2], pass) && Java_IsArg[[]*InternetAddress](args[3], pass) && Java_IsArg[[]*InternetAddress](args[4], pass) && Java_IsArg[[]*InternetAddress](args[5], pass) && Java_IsArg[string](args[6], pass) && Java_IsArg[[]*MimeBodyPart](args[7], pass) {
            this.Init_2(Java_Arg[string](args[0]), Java_Arg[Authenticator](args[1]), Java_Arg[[]*InternetAddress](args[2]), Java_Arg[[]*InternetAddress](args[3]), Java_Arg[[]*InternetAddress](args[4]), Java_Arg[[]*InternetAddress](args[5]), Java_Arg[string](args[6]), Java_Arg[[]*MimeBodyPart](args[7]))
            return
        }
    }
    panic("No Java overload of Email.init for the arguments")
}


// throws MessagingException
func (this *Email) Init_2(server string, authenticator Authenticator, addresses_from []*InternetAddress, addresses_to []*InternetAddress, addresses_cc []*InternetAddress, addresses_bcc []*InternetAddress, subject string, mimeBodyParts []*MimeBodyPart) {
        this.Properties = NewProperties()
        this.Properties.Put(Email_SMTP_HOST, server)
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put("Continuing on Exception: Unable to get HostName", this, "init()", e)
            }
            this.Properties.Put(Email_SMTP_LOCAL_HOST, "FakeHostName")
                }
            }()
            var hostName string = StringUtil_GetInstance().EMPTY_STRING; _ = hostName
            if StringValidationUtil_GetInstance().IsEmpty(hostName) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                    this.LogUtil.PutF("Unable to get HostName so using fake", this, "init()")
                }
                this.Properties.Put(Email_SMTP_LOCAL_HOST, "FakeHostName")
            }
            return
        }(); tryControl1 == 1 {
            return
        }
        if this.IsDebug {
            this.Properties.Put(Email_DEBUG, "true")
        }
        var session *Session = Session_GetInstance(this.Properties, authenticator); _ = session
        if this.IsDebug {
            session.SetDebug(true)
            this.Bs = &strings.Builder{}
            var printStream *PrintStream = this.Bs; _ = printStream
            session.SetDebugOut(printStream)
        }
        this.Msg = NewMimeMessage(session)
        this.Msg.AddFrom(addresses_from)
        this.Msg.SetRecipients(Message_RecipientType.TO, addresses_to)
        this.Msg.SetRecipients(Message_RecipientType.CC, addresses_cc)
        this.Msg.SetRecipients(Message_RecipientType.BCC, addresses_bcc)
        this.Msg.SetSubject(subject)
        var mimeMultipart *MimeMultipart = NewMimeMultipart(); _ = mimeMultipart
        for i := int(0); i < len(mimeBodyParts); i++ {
            mimeMultipart.AddBodyPart(mimeBodyParts[i])
        }
        this.Msg.SetContent(mimeMultipart)
}


func (this *Email) GetMimeMessage() *MimeMessage {
        return this.Msg
}


func (this *Email) GetDebugInfo() string {
        if this.IsDebug {
            return "More Logging Before This To Standard Out: " + fmt.Sprint(this.Bs.String())
        } else {
            return "Debugging Is Off"
        }
}


// throws Exception
func (this *Email) Log() string {
        return "\nHashMap: " + fmt.Sprint(this.ToHashMap())
}


// throws Exception
func (this *Email) ToHashMap() *HashMap {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HashMap) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toHashMap()", e)
            }
            panic(e)
                }
            }()
            var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
            hashMap.Put(EmailData_SERVER, any(this.Properties.Get(Email_SMTP_HOST)).(string))
            var addresses []*Address = this.Msg.GetFrom(); _ = addresses
            if addresses != nil {
                for index := int(0); index < len(addresses); index++ {
                    hashMap.Put(EmailData_FROM, fmt.Sprint(addresses[index]))
                }
            }
            addresses = this.Msg.GetRecipients(Message_RecipientType.TO)
            if addresses != nil {
                for index := int(0); index < len(addresses); index++ {
                    hashMap.Put(EmailData_TO, fmt.Sprint(addresses[index]))
                }
            }
            addresses = this.Msg.GetRecipients(Message_RecipientType.CC)
            if addresses != nil {
                for index := int(0); index < len(addresses); index++ {
                    hashMap.Put(EmailData_CC, fmt.Sprint(addresses[index]))
                }
            }
            addresses = this.Msg.GetRecipients(Message_RecipientType.BCC)
            if addresses != nil {
                for index := int(0); index < len(addresses); index++ {
                    hashMap.Put(EmailData_BCC, fmt.Sprint(addresses[index]))
                }
            }
            hashMap.Put(EmailData_SUBJECT, this.Msg.GetSubject())
            var mimeMultipart *MimeMultipart = any(this.Msg.GetContent()).(*MimeMultipart); _ = mimeMultipart
            if mimeMultipart != nil {
                for index := int(0); index < mimeMultipart.GetCount(); index++ {
                    var content string = any(mimeMultipart.GetBodyPart(index).GetContent()).(string); _ = content
                    hashMap.Put(EmailData_CONTENT, content)
                }
            }
            tryControl1, tryResult1 = 1, hashMap; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *Email) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNameValueNodes(document, EmailData_NAME, this.ToHashMap()); _ = node
        return node
}


// throws Exception
func (this *Email) ToXmlDoc() *Document {
        panic(errors.New(fmt.Sprint(CommonStrings_GetInstance().NOT_IMPLEMENTED)))
}

