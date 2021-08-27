// Code generated by "./generator ./com.deepin.daemon.authenticate"; DO NOT EDIT.

package authenticate

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockAuthenticate struct {
	MockInterfaceAuthenticate // interface com.deepin.daemon.Authenticate
	proxy.MockObject
}

type MockInterfaceAuthenticate struct {
	mock.Mock
}

// method Authenticate

func (v *MockInterfaceAuthenticate) GoAuthenticate(flags dbus.Flags, ch chan *dbus.Call, username string, authFlags int32, appType int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, authFlags, appType)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) Authenticate(flags dbus.Flags, username string, authFlags int32, appType int32) (string, error) {
	mockArgs := v.Called(flags, username, authFlags, appType)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetLimits

func (v *MockInterfaceAuthenticate) GoGetLimits(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) GetLimits(flags dbus.Flags, username string) (string, error) {
	mockArgs := v.Called(flags, username)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method PreOneKeyLogin

func (v *MockInterfaceAuthenticate) GoPreOneKeyLogin(flags dbus.Flags, ch chan *dbus.Call, flag int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, flag)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) PreOneKeyLogin(flags dbus.Flags, flag int32) (string, error) {
	mockArgs := v.Called(flags, flag)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal LimitUpdated

func (v *MockInterfaceAuthenticate) ConnectLimitUpdated(cb func(username string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property SupportEncrypts s

func (v *MockInterfaceAuthenticate) SupportEncrypts() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property FrameworkState i

func (v *MockInterfaceAuthenticate) FrameworkState() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportedFlags i

func (v *MockInterfaceAuthenticate) SupportedFlags() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockFingerprint struct {
	MockInterfaceFingerprint // interface com.deepin.daemon.Authenticate.Fingerprint
	proxy.MockObject
}

type MockInterfaceFingerprint struct {
	mock.Mock
}

// method Claim

func (v *MockInterfaceFingerprint) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) Claim(flags dbus.Flags, username string, claimed bool) error {
	mockArgs := v.Called(flags, username, claimed)

	return mockArgs.Error(0)
}

// method DeleteAllFingers

func (v *MockInterfaceFingerprint) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) DeleteAllFingers(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method DeleteFinger

func (v *MockInterfaceFingerprint) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) DeleteFinger(flags dbus.Flags, username string, finger string) error {
	mockArgs := v.Called(flags, username, finger)

	return mockArgs.Error(0)
}

// method Enroll

func (v *MockInterfaceFingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) Enroll(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// method ListFingers

func (v *MockInterfaceFingerprint) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) ListFingers(flags dbus.Flags, username string) ([]string, error) {
	mockArgs := v.Called(flags, username)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetDefaultDevice

func (v *MockInterfaceFingerprint) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) SetDefaultDevice(flags dbus.Flags, device string) error {
	mockArgs := v.Called(flags, device)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *MockInterfaceFingerprint) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceFingerprint) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Verify

func (v *MockInterfaceFingerprint) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) Verify(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceFingerprint) ConnectEnrollStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceFingerprint) ConnectVerifyStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Touch

func (v *MockInterfaceFingerprint) ConnectTouch(cb func(id string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property DefaultDevice s

func (v *MockInterfaceFingerprint) DefaultDevice() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Devices s

func (v *MockInterfaceFingerprint) Devices() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockUKey struct {
	MockInterfaceUkey // interface com.deepin.daemon.Authenticate.UKey
	proxy.MockObject
}

type MockInterfaceUkey struct {
	mock.Mock
}

// method ConstructVerification

func (v *MockInterfaceUkey) GoConstructVerification(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, useDefaultService bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username, useDefaultService)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) ConstructVerification(flags dbus.Flags, serviceName string, username string, useDefaultService bool) (string, error) {
	mockArgs := v.Called(flags, serviceName, username, useDefaultService)

	ret0, ok := mockArgs.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetPINLength

func (v *MockInterfaceUkey) GoGetPINLength(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, useDefaultDevice bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username, useDefaultDevice)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) GetPINLength(flags dbus.Flags, serviceName string, username string, useDefaultDevice bool) (int32, error) {
	mockArgs := v.Called(flags, serviceName, username, useDefaultDevice)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetUserList

func (v *MockInterfaceUkey) GoGetUserList(flags dbus.Flags, ch chan *dbus.Call, serviceName string, useDefaultDevice bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, useDefaultDevice)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) GetUserList(flags dbus.Flags, serviceName string, useDefaultDevice bool) ([]string, error) {
	mockArgs := v.Called(flags, serviceName, useDefaultDevice)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetDefaultDevice

func (v *MockInterfaceUkey) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetDefaultDevice(flags dbus.Flags, device string) error {
	mockArgs := v.Called(flags, device)

	return mockArgs.Error(0)
}

// method SetPin

func (v *MockInterfaceUkey) GoSetPin(flags dbus.Flags, ch chan *dbus.Call, id string, pin string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id, pin)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetPin(flags dbus.Flags, id string, pin string) error {
	mockArgs := v.Called(flags, id, pin)

	return mockArgs.Error(0)
}

// method SetSessionPath

func (v *MockInterfaceUkey) GoSetSessionPath(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetSessionPath(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method StartVerify

func (v *MockInterfaceUkey) GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) StartVerify(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceUkey) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) StopVerify(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// signal VerifyResult

func (v *MockInterfaceUkey) ConnectVerifyResult(cb func(id string, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal State

func (v *MockInterfaceUkey) ConnectState(cb func(id string, state int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property ValidDevices s

func (v *MockInterfaceUkey) ValidDevices() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DefaultDevice s

func (v *MockInterfaceUkey) DefaultDevice() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
