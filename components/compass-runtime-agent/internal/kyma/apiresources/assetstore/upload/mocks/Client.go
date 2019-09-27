// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "kyma-project.io/compass-runtime-agent/internal/apperrors"
import mock "github.com/stretchr/testify/mock"
import upload "kyma-project.io/compass-runtime-agent/internal/kyma/apiresources/assetstore/upload"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Upload provides a mock function with given fields: fileName, contents
func (_m *Client) Upload(fileName string, contents []byte) (upload.UploadedFile, apperrors.AppError) {
	ret := _m.Called(fileName, contents)

	var r0 upload.UploadedFile
	if rf, ok := ret.Get(0).(func(string, []byte) upload.UploadedFile); ok {
		r0 = rf(fileName, contents)
	} else {
		r0 = ret.Get(0).(upload.UploadedFile)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, []byte) apperrors.AppError); ok {
		r1 = rf(fileName, contents)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
