// Code generated by "./generator ./system/com.deepin.daemon.authenticate.BiometricsDriver"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package BiometricsDriver

import "errors"
import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type BiometricsDriver interface {
	biometricsDriver // interface com.deepin.daemon.authenticate.BiometricsDriver
	proxy.Object
}

type objectBiometricsDriver struct {
	interfaceBiometricsDriver // interface com.deepin.daemon.authenticate.BiometricsDriver
	proxy.ImplObject
}

func NewBiometricsDriver(conn *dbus.Conn, serviceName string, path dbus.ObjectPath) (BiometricsDriver, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectBiometricsDriver)
	obj.ImplObject.Init_(conn, serviceName, path)
	return obj, nil
}

type biometricsDriver interface {
	SetInterfaceName_(name string)
	GoDelete(flags dbus.Flags, ch chan *dbus.Call, chara string) *dbus.Call
	Delete(flags dbus.Flags, chara string) error
	GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, chara string, type0 int32, action string) *dbus.Call
	EnrollStart(flags dbus.Flags, chara string, type0 int32, action string) (dbus.UnixFD, error)
	GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call, action string) *dbus.Call
	EnrollStop(flags dbus.Flags, action string) error
	GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, charas []string, action string) *dbus.Call
	VerifyStart(flags dbus.Flags, charas []string, action string) (dbus.UnixFD, error)
	GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call, action string) *dbus.Call
	VerifyStop(flags dbus.Flags, action string) error
	ConnectVerifyStatus(cb func(action string, verifyStatusCode int32, message string)) (dbusutil.SignalHandlerId, error)
	ConnectEnrollStatus(cb func(action string, enrollStatusCode int32, message string)) (dbusutil.SignalHandlerId, error)
	List() proxy.PropStringArray
	CharaType() proxy.PropInt32
	Claim() proxy.PropBool
}

type interfaceBiometricsDriver struct{}

func (v *interfaceBiometricsDriver) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (v *interfaceBiometricsDriver) SetInterfaceName_(name string) {
	v.GetObject_().SetExtra("customIfc", name)
}

func (v *interfaceBiometricsDriver) GetInterfaceName_() string {
	ifcName, _ := v.GetObject_().GetExtra("customIfc")
	ifcNameStr, _ := ifcName.(string)
	return ifcNameStr
}

// method Delete

func (v *interfaceBiometricsDriver) GoDelete(flags dbus.Flags, ch chan *dbus.Call, chara string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, chara)
}

func (v *interfaceBiometricsDriver) Delete(flags dbus.Flags, chara string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), chara).Done).Err
}

// method EnrollStart

func (v *interfaceBiometricsDriver) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, chara string, type0 int32, action string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStart", flags, ch, chara, type0, action)
}

func (*interfaceBiometricsDriver) StoreEnrollStart(call *dbus.Call) (fd dbus.UnixFD, err error) {
	err = call.Store(&fd)
	return
}

func (v *interfaceBiometricsDriver) EnrollStart(flags dbus.Flags, chara string, type0 int32, action string) (dbus.UnixFD, error) {
	return v.StoreEnrollStart(
		<-v.GoEnrollStart(flags, make(chan *dbus.Call, 1), chara, type0, action).Done)
}

// method EnrollStop

func (v *interfaceBiometricsDriver) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call, action string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStop", flags, ch, action)
}

func (v *interfaceBiometricsDriver) EnrollStop(flags dbus.Flags, action string) error {
	return (<-v.GoEnrollStop(flags, make(chan *dbus.Call, 1), action).Done).Err
}

// method VerifyStart

func (v *interfaceBiometricsDriver) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, charas []string, action string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStart", flags, ch, charas, action)
}

func (*interfaceBiometricsDriver) StoreVerifyStart(call *dbus.Call) (fd dbus.UnixFD, err error) {
	err = call.Store(&fd)
	return
}

func (v *interfaceBiometricsDriver) VerifyStart(flags dbus.Flags, charas []string, action string) (dbus.UnixFD, error) {
	return v.StoreVerifyStart(
		<-v.GoVerifyStart(flags, make(chan *dbus.Call, 1), charas, action).Done)
}

// method VerifyStop

func (v *interfaceBiometricsDriver) GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call, action string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStop", flags, ch, action)
}

func (v *interfaceBiometricsDriver) VerifyStop(flags dbus.Flags, action string) error {
	return (<-v.GoVerifyStop(flags, make(chan *dbus.Call, 1), action).Done).Err
}

// signal VerifyStatus

func (v *interfaceBiometricsDriver) ConnectVerifyStatus(cb func(action string, verifyStatusCode int32, message string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var action string
		var verifyStatusCode int32
		var message string
		err := dbus.Store(sig.Body, &action, &verifyStatusCode, &message)
		if err == nil {
			cb(action, verifyStatusCode, message)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal EnrollStatus

func (v *interfaceBiometricsDriver) ConnectEnrollStatus(cb func(action string, enrollStatusCode int32, message string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var action string
		var enrollStatusCode int32
		var message string
		err := dbus.Store(sig.Body, &action, &enrollStatusCode, &message)
		if err == nil {
			cb(action, enrollStatusCode, message)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property List as

func (v *interfaceBiometricsDriver) List() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "List",
	}
}

// property CharaType i

func (v *interfaceBiometricsDriver) CharaType() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "CharaType",
	}
}

// property Claim b

func (v *interfaceBiometricsDriver) Claim() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Claim",
	}
}