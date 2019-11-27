// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/internal/resource (interfaces: Connection)

package views

import (
	k8s "github.com/derailed/k9s/internal/k8s"
	pegomock "github.com/petergtz/pegomock"
	v1 "k8s.io/api/core/v1"
	version "k8s.io/apimachinery/pkg/version"
	disk "k8s.io/client-go/discovery/cached/disk"
	dynamic "k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	versioned "k8s.io/metrics/pkg/client/clientset/versioned"
	"reflect"
	"time"
)

type MockConnection struct {
	fail func(message string, callerSkip ...int)
}

func NewMockConnection(options ...pegomock.Option) *MockConnection {
	mock := &MockConnection{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockConnection) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockConnection) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockConnection) CachedDiscovery() (*disk.CachedDiscoveryClient, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CachedDiscovery", params, []reflect.Type{reflect.TypeOf((**disk.CachedDiscoveryClient)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *disk.CachedDiscoveryClient
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*disk.CachedDiscoveryClient)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) CanI(_param0 string, _param1 string, _param2 []string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CanI", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) CheckListNSAccess() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CheckListNSAccess", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockConnection) CheckNSAccess(_param0 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CheckNSAccess", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockConnection) Config() *k8s.Config {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Config", params, []reflect.Type{reflect.TypeOf((**k8s.Config)(nil)).Elem()})
	var ret0 *k8s.Config
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*k8s.Config)
		}
	}
	return ret0
}

func (mock *MockConnection) CurrentNamespaceName() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentNamespaceName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) DialOrDie() kubernetes.Interface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DialOrDie", params, []reflect.Type{reflect.TypeOf((*kubernetes.Interface)(nil)).Elem()})
	var ret0 kubernetes.Interface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(kubernetes.Interface)
		}
	}
	return ret0
}

func (mock *MockConnection) DynDialOrDie() dynamic.Interface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DynDialOrDie", params, []reflect.Type{reflect.TypeOf((*dynamic.Interface)(nil)).Elem()})
	var ret0 dynamic.Interface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(dynamic.Interface)
		}
	}
	return ret0
}

func (mock *MockConnection) FetchNodes() (*v1.NodeList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchNodes", params, []reflect.Type{reflect.TypeOf((**v1.NodeList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.NodeList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.NodeList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) HasMetrics() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasMetrics", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockConnection) IsNamespaced(_param0 string) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("IsNamespaced", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockConnection) MXDial() (*versioned.Clientset, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MXDial", params, []reflect.Type{reflect.TypeOf((**versioned.Clientset)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *versioned.Clientset
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*versioned.Clientset)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) NSDialOrDie() dynamic.NamespaceableResourceInterface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NSDialOrDie", params, []reflect.Type{reflect.TypeOf((*dynamic.NamespaceableResourceInterface)(nil)).Elem()})
	var ret0 dynamic.NamespaceableResourceInterface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(dynamic.NamespaceableResourceInterface)
		}
	}
	return ret0
}

func (mock *MockConnection) NodePods(_param0 string) (*v1.PodList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NodePods", params, []reflect.Type{reflect.TypeOf((**v1.PodList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.PodList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.PodList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) RestConfigOrDie() *rest.Config {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RestConfigOrDie", params, []reflect.Type{reflect.TypeOf((**rest.Config)(nil)).Elem()})
	var ret0 *rest.Config
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*rest.Config)
		}
	}
	return ret0
}

func (mock *MockConnection) ServerVersion() (*version.Info, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ServerVersion", params, []reflect.Type{reflect.TypeOf((**version.Info)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *version.Info
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*version.Info)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) SupportsRes(_param0 string, _param1 []string) (string, bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsRes", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 bool
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(bool)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockConnection) SupportsResource(_param0 string) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsResource", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockConnection) SwitchContextOrDie(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SwitchContextOrDie", params, []reflect.Type{})
}

func (mock *MockConnection) ValidNamespaces() ([]v1.Namespace, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConnection().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ValidNamespaces", params, []reflect.Type{reflect.TypeOf((*[]v1.Namespace)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []v1.Namespace
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]v1.Namespace)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockConnection) VerifyWasCalledOnce() *VerifierMockConnection {
	return &VerifierMockConnection{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockConnection) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockConnection {
	return &VerifierMockConnection{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockConnection) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockConnection {
	return &VerifierMockConnection{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockConnection) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockConnection {
	return &VerifierMockConnection{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockConnection struct {
	mock                   *MockConnection
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockConnection) CachedDiscovery() *MockConnection_CachedDiscovery_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CachedDiscovery", params, verifier.timeout)
	return &MockConnection_CachedDiscovery_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_CachedDiscovery_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_CachedDiscovery_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_CachedDiscovery_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) CanI(_param0 string, _param1 string, _param2 []string) *MockConnection_CanI_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CanI", params, verifier.timeout)
	return &MockConnection_CanI_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_CanI_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_CanI_OngoingVerification) GetCapturedArguments() (string, string, []string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockConnection_CanI_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) CheckListNSAccess() *MockConnection_CheckListNSAccess_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CheckListNSAccess", params, verifier.timeout)
	return &MockConnection_CheckListNSAccess_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_CheckListNSAccess_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_CheckListNSAccess_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_CheckListNSAccess_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) CheckNSAccess(_param0 string) *MockConnection_CheckNSAccess_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CheckNSAccess", params, verifier.timeout)
	return &MockConnection_CheckNSAccess_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_CheckNSAccess_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_CheckNSAccess_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockConnection_CheckNSAccess_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) Config() *MockConnection_Config_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Config", params, verifier.timeout)
	return &MockConnection_Config_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_Config_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_Config_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_Config_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) CurrentNamespaceName() *MockConnection_CurrentNamespaceName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentNamespaceName", params, verifier.timeout)
	return &MockConnection_CurrentNamespaceName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_CurrentNamespaceName_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_CurrentNamespaceName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_CurrentNamespaceName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) DialOrDie() *MockConnection_DialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DialOrDie", params, verifier.timeout)
	return &MockConnection_DialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_DialOrDie_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_DialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_DialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) DynDialOrDie() *MockConnection_DynDialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DynDialOrDie", params, verifier.timeout)
	return &MockConnection_DynDialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_DynDialOrDie_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_DynDialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_DynDialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) FetchNodes() *MockConnection_FetchNodes_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchNodes", params, verifier.timeout)
	return &MockConnection_FetchNodes_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_FetchNodes_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_FetchNodes_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_FetchNodes_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) HasMetrics() *MockConnection_HasMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasMetrics", params, verifier.timeout)
	return &MockConnection_HasMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_HasMetrics_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_HasMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_HasMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) IsNamespaced(_param0 string) *MockConnection_IsNamespaced_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "IsNamespaced", params, verifier.timeout)
	return &MockConnection_IsNamespaced_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_IsNamespaced_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_IsNamespaced_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockConnection_IsNamespaced_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) MXDial() *MockConnection_MXDial_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MXDial", params, verifier.timeout)
	return &MockConnection_MXDial_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_MXDial_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_MXDial_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_MXDial_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) NSDialOrDie() *MockConnection_NSDialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NSDialOrDie", params, verifier.timeout)
	return &MockConnection_NSDialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_NSDialOrDie_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_NSDialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_NSDialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) NodePods(_param0 string) *MockConnection_NodePods_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodePods", params, verifier.timeout)
	return &MockConnection_NodePods_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_NodePods_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_NodePods_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockConnection_NodePods_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) RestConfigOrDie() *MockConnection_RestConfigOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RestConfigOrDie", params, verifier.timeout)
	return &MockConnection_RestConfigOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_RestConfigOrDie_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_RestConfigOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_RestConfigOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) ServerVersion() *MockConnection_ServerVersion_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ServerVersion", params, verifier.timeout)
	return &MockConnection_ServerVersion_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_ServerVersion_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_ServerVersion_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_ServerVersion_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockConnection) SupportsRes(_param0 string, _param1 []string) *MockConnection_SupportsRes_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsRes", params, verifier.timeout)
	return &MockConnection_SupportsRes_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_SupportsRes_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_SupportsRes_OngoingVerification) GetCapturedArguments() (string, []string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockConnection_SupportsRes_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) SupportsResource(_param0 string) *MockConnection_SupportsResource_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsResource", params, verifier.timeout)
	return &MockConnection_SupportsResource_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_SupportsResource_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_SupportsResource_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockConnection_SupportsResource_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) SwitchContextOrDie(_param0 string) *MockConnection_SwitchContextOrDie_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SwitchContextOrDie", params, verifier.timeout)
	return &MockConnection_SwitchContextOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_SwitchContextOrDie_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_SwitchContextOrDie_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockConnection_SwitchContextOrDie_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockConnection) ValidNamespaces() *MockConnection_ValidNamespaces_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ValidNamespaces", params, verifier.timeout)
	return &MockConnection_ValidNamespaces_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockConnection_ValidNamespaces_OngoingVerification struct {
	mock              *MockConnection
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockConnection_ValidNamespaces_OngoingVerification) GetCapturedArguments() {
}

func (c *MockConnection_ValidNamespaces_OngoingVerification) GetAllCapturedArguments() {
}
