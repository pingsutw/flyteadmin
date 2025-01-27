// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NamespaceMappingConfiguration is an autogenerated mock type for the NamespaceMappingConfiguration type
type NamespaceMappingConfiguration struct {
	mock.Mock
}

type NamespaceMappingConfiguration_GetNamespaceTemplate struct {
	*mock.Call
}

func (_m NamespaceMappingConfiguration_GetNamespaceTemplate) Return(_a0 string) *NamespaceMappingConfiguration_GetNamespaceTemplate {
	return &NamespaceMappingConfiguration_GetNamespaceTemplate{Call: _m.Call.Return(_a0)}
}

func (_m *NamespaceMappingConfiguration) OnGetNamespaceTemplate() *NamespaceMappingConfiguration_GetNamespaceTemplate {
	c := _m.On("GetNamespaceTemplate")
	return &NamespaceMappingConfiguration_GetNamespaceTemplate{Call: c}
}

func (_m *NamespaceMappingConfiguration) OnGetNamespaceTemplateMatch(matchers ...interface{}) *NamespaceMappingConfiguration_GetNamespaceTemplate {
	c := _m.On("GetNamespaceTemplate", matchers...)
	return &NamespaceMappingConfiguration_GetNamespaceTemplate{Call: c}
}

// GetNamespaceTemplate provides a mock function with given fields:
func (_m *NamespaceMappingConfiguration) GetNamespaceTemplate() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
