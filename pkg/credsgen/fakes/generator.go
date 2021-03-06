// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/quarks-operator/pkg/credsgen"
)

type FakeGenerator struct {
	GenerateCertificateStub        func(string, credsgen.CertificateGenerationRequest) (credsgen.Certificate, error)
	generateCertificateMutex       sync.RWMutex
	generateCertificateArgsForCall []struct {
		arg1 string
		arg2 credsgen.CertificateGenerationRequest
	}
	generateCertificateReturns struct {
		result1 credsgen.Certificate
		result2 error
	}
	generateCertificateReturnsOnCall map[int]struct {
		result1 credsgen.Certificate
		result2 error
	}
	GenerateCertificateSigningRequestStub        func(credsgen.CertificateGenerationRequest) ([]byte, []byte, error)
	generateCertificateSigningRequestMutex       sync.RWMutex
	generateCertificateSigningRequestArgsForCall []struct {
		arg1 credsgen.CertificateGenerationRequest
	}
	generateCertificateSigningRequestReturns struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	generateCertificateSigningRequestReturnsOnCall map[int]struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	GeneratePasswordStub        func(string, credsgen.PasswordGenerationRequest) string
	generatePasswordMutex       sync.RWMutex
	generatePasswordArgsForCall []struct {
		arg1 string
		arg2 credsgen.PasswordGenerationRequest
	}
	generatePasswordReturns struct {
		result1 string
	}
	generatePasswordReturnsOnCall map[int]struct {
		result1 string
	}
	GenerateRSAKeyStub        func(string) (credsgen.RSAKey, error)
	generateRSAKeyMutex       sync.RWMutex
	generateRSAKeyArgsForCall []struct {
		arg1 string
	}
	generateRSAKeyReturns struct {
		result1 credsgen.RSAKey
		result2 error
	}
	generateRSAKeyReturnsOnCall map[int]struct {
		result1 credsgen.RSAKey
		result2 error
	}
	GenerateSSHKeyStub        func(string) (credsgen.SSHKey, error)
	generateSSHKeyMutex       sync.RWMutex
	generateSSHKeyArgsForCall []struct {
		arg1 string
	}
	generateSSHKeyReturns struct {
		result1 credsgen.SSHKey
		result2 error
	}
	generateSSHKeyReturnsOnCall map[int]struct {
		result1 credsgen.SSHKey
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGenerator) GenerateCertificate(arg1 string, arg2 credsgen.CertificateGenerationRequest) (credsgen.Certificate, error) {
	fake.generateCertificateMutex.Lock()
	ret, specificReturn := fake.generateCertificateReturnsOnCall[len(fake.generateCertificateArgsForCall)]
	fake.generateCertificateArgsForCall = append(fake.generateCertificateArgsForCall, struct {
		arg1 string
		arg2 credsgen.CertificateGenerationRequest
	}{arg1, arg2})
	fake.recordInvocation("GenerateCertificate", []interface{}{arg1, arg2})
	fake.generateCertificateMutex.Unlock()
	if fake.GenerateCertificateStub != nil {
		return fake.GenerateCertificateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateCertificateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGenerator) GenerateCertificateCallCount() int {
	fake.generateCertificateMutex.RLock()
	defer fake.generateCertificateMutex.RUnlock()
	return len(fake.generateCertificateArgsForCall)
}

func (fake *FakeGenerator) GenerateCertificateCalls(stub func(string, credsgen.CertificateGenerationRequest) (credsgen.Certificate, error)) {
	fake.generateCertificateMutex.Lock()
	defer fake.generateCertificateMutex.Unlock()
	fake.GenerateCertificateStub = stub
}

func (fake *FakeGenerator) GenerateCertificateArgsForCall(i int) (string, credsgen.CertificateGenerationRequest) {
	fake.generateCertificateMutex.RLock()
	defer fake.generateCertificateMutex.RUnlock()
	argsForCall := fake.generateCertificateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGenerator) GenerateCertificateReturns(result1 credsgen.Certificate, result2 error) {
	fake.generateCertificateMutex.Lock()
	defer fake.generateCertificateMutex.Unlock()
	fake.GenerateCertificateStub = nil
	fake.generateCertificateReturns = struct {
		result1 credsgen.Certificate
		result2 error
	}{result1, result2}
}

func (fake *FakeGenerator) GenerateCertificateReturnsOnCall(i int, result1 credsgen.Certificate, result2 error) {
	fake.generateCertificateMutex.Lock()
	defer fake.generateCertificateMutex.Unlock()
	fake.GenerateCertificateStub = nil
	if fake.generateCertificateReturnsOnCall == nil {
		fake.generateCertificateReturnsOnCall = make(map[int]struct {
			result1 credsgen.Certificate
			result2 error
		})
	}
	fake.generateCertificateReturnsOnCall[i] = struct {
		result1 credsgen.Certificate
		result2 error
	}{result1, result2}
}

func (fake *FakeGenerator) GenerateCertificateSigningRequest(arg1 credsgen.CertificateGenerationRequest) ([]byte, []byte, error) {
	fake.generateCertificateSigningRequestMutex.Lock()
	ret, specificReturn := fake.generateCertificateSigningRequestReturnsOnCall[len(fake.generateCertificateSigningRequestArgsForCall)]
	fake.generateCertificateSigningRequestArgsForCall = append(fake.generateCertificateSigningRequestArgsForCall, struct {
		arg1 credsgen.CertificateGenerationRequest
	}{arg1})
	fake.recordInvocation("GenerateCertificateSigningRequest", []interface{}{arg1})
	fake.generateCertificateSigningRequestMutex.Unlock()
	if fake.GenerateCertificateSigningRequestStub != nil {
		return fake.GenerateCertificateSigningRequestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.generateCertificateSigningRequestReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeGenerator) GenerateCertificateSigningRequestCallCount() int {
	fake.generateCertificateSigningRequestMutex.RLock()
	defer fake.generateCertificateSigningRequestMutex.RUnlock()
	return len(fake.generateCertificateSigningRequestArgsForCall)
}

func (fake *FakeGenerator) GenerateCertificateSigningRequestCalls(stub func(credsgen.CertificateGenerationRequest) ([]byte, []byte, error)) {
	fake.generateCertificateSigningRequestMutex.Lock()
	defer fake.generateCertificateSigningRequestMutex.Unlock()
	fake.GenerateCertificateSigningRequestStub = stub
}

func (fake *FakeGenerator) GenerateCertificateSigningRequestArgsForCall(i int) credsgen.CertificateGenerationRequest {
	fake.generateCertificateSigningRequestMutex.RLock()
	defer fake.generateCertificateSigningRequestMutex.RUnlock()
	argsForCall := fake.generateCertificateSigningRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenerator) GenerateCertificateSigningRequestReturns(result1 []byte, result2 []byte, result3 error) {
	fake.generateCertificateSigningRequestMutex.Lock()
	defer fake.generateCertificateSigningRequestMutex.Unlock()
	fake.GenerateCertificateSigningRequestStub = nil
	fake.generateCertificateSigningRequestReturns = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGenerator) GenerateCertificateSigningRequestReturnsOnCall(i int, result1 []byte, result2 []byte, result3 error) {
	fake.generateCertificateSigningRequestMutex.Lock()
	defer fake.generateCertificateSigningRequestMutex.Unlock()
	fake.GenerateCertificateSigningRequestStub = nil
	if fake.generateCertificateSigningRequestReturnsOnCall == nil {
		fake.generateCertificateSigningRequestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 []byte
			result3 error
		})
	}
	fake.generateCertificateSigningRequestReturnsOnCall[i] = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGenerator) GeneratePassword(arg1 string, arg2 credsgen.PasswordGenerationRequest) string {
	fake.generatePasswordMutex.Lock()
	ret, specificReturn := fake.generatePasswordReturnsOnCall[len(fake.generatePasswordArgsForCall)]
	fake.generatePasswordArgsForCall = append(fake.generatePasswordArgsForCall, struct {
		arg1 string
		arg2 credsgen.PasswordGenerationRequest
	}{arg1, arg2})
	fake.recordInvocation("GeneratePassword", []interface{}{arg1, arg2})
	fake.generatePasswordMutex.Unlock()
	if fake.GeneratePasswordStub != nil {
		return fake.GeneratePasswordStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generatePasswordReturns
	return fakeReturns.result1
}

func (fake *FakeGenerator) GeneratePasswordCallCount() int {
	fake.generatePasswordMutex.RLock()
	defer fake.generatePasswordMutex.RUnlock()
	return len(fake.generatePasswordArgsForCall)
}

func (fake *FakeGenerator) GeneratePasswordCalls(stub func(string, credsgen.PasswordGenerationRequest) string) {
	fake.generatePasswordMutex.Lock()
	defer fake.generatePasswordMutex.Unlock()
	fake.GeneratePasswordStub = stub
}

func (fake *FakeGenerator) GeneratePasswordArgsForCall(i int) (string, credsgen.PasswordGenerationRequest) {
	fake.generatePasswordMutex.RLock()
	defer fake.generatePasswordMutex.RUnlock()
	argsForCall := fake.generatePasswordArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGenerator) GeneratePasswordReturns(result1 string) {
	fake.generatePasswordMutex.Lock()
	defer fake.generatePasswordMutex.Unlock()
	fake.GeneratePasswordStub = nil
	fake.generatePasswordReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGenerator) GeneratePasswordReturnsOnCall(i int, result1 string) {
	fake.generatePasswordMutex.Lock()
	defer fake.generatePasswordMutex.Unlock()
	fake.GeneratePasswordStub = nil
	if fake.generatePasswordReturnsOnCall == nil {
		fake.generatePasswordReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.generatePasswordReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGenerator) GenerateRSAKey(arg1 string) (credsgen.RSAKey, error) {
	fake.generateRSAKeyMutex.Lock()
	ret, specificReturn := fake.generateRSAKeyReturnsOnCall[len(fake.generateRSAKeyArgsForCall)]
	fake.generateRSAKeyArgsForCall = append(fake.generateRSAKeyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GenerateRSAKey", []interface{}{arg1})
	fake.generateRSAKeyMutex.Unlock()
	if fake.GenerateRSAKeyStub != nil {
		return fake.GenerateRSAKeyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateRSAKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGenerator) GenerateRSAKeyCallCount() int {
	fake.generateRSAKeyMutex.RLock()
	defer fake.generateRSAKeyMutex.RUnlock()
	return len(fake.generateRSAKeyArgsForCall)
}

func (fake *FakeGenerator) GenerateRSAKeyCalls(stub func(string) (credsgen.RSAKey, error)) {
	fake.generateRSAKeyMutex.Lock()
	defer fake.generateRSAKeyMutex.Unlock()
	fake.GenerateRSAKeyStub = stub
}

func (fake *FakeGenerator) GenerateRSAKeyArgsForCall(i int) string {
	fake.generateRSAKeyMutex.RLock()
	defer fake.generateRSAKeyMutex.RUnlock()
	argsForCall := fake.generateRSAKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenerator) GenerateRSAKeyReturns(result1 credsgen.RSAKey, result2 error) {
	fake.generateRSAKeyMutex.Lock()
	defer fake.generateRSAKeyMutex.Unlock()
	fake.GenerateRSAKeyStub = nil
	fake.generateRSAKeyReturns = struct {
		result1 credsgen.RSAKey
		result2 error
	}{result1, result2}
}

func (fake *FakeGenerator) GenerateRSAKeyReturnsOnCall(i int, result1 credsgen.RSAKey, result2 error) {
	fake.generateRSAKeyMutex.Lock()
	defer fake.generateRSAKeyMutex.Unlock()
	fake.GenerateRSAKeyStub = nil
	if fake.generateRSAKeyReturnsOnCall == nil {
		fake.generateRSAKeyReturnsOnCall = make(map[int]struct {
			result1 credsgen.RSAKey
			result2 error
		})
	}
	fake.generateRSAKeyReturnsOnCall[i] = struct {
		result1 credsgen.RSAKey
		result2 error
	}{result1, result2}
}

func (fake *FakeGenerator) GenerateSSHKey(arg1 string) (credsgen.SSHKey, error) {
	fake.generateSSHKeyMutex.Lock()
	ret, specificReturn := fake.generateSSHKeyReturnsOnCall[len(fake.generateSSHKeyArgsForCall)]
	fake.generateSSHKeyArgsForCall = append(fake.generateSSHKeyArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GenerateSSHKey", []interface{}{arg1})
	fake.generateSSHKeyMutex.Unlock()
	if fake.GenerateSSHKeyStub != nil {
		return fake.GenerateSSHKeyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateSSHKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGenerator) GenerateSSHKeyCallCount() int {
	fake.generateSSHKeyMutex.RLock()
	defer fake.generateSSHKeyMutex.RUnlock()
	return len(fake.generateSSHKeyArgsForCall)
}

func (fake *FakeGenerator) GenerateSSHKeyCalls(stub func(string) (credsgen.SSHKey, error)) {
	fake.generateSSHKeyMutex.Lock()
	defer fake.generateSSHKeyMutex.Unlock()
	fake.GenerateSSHKeyStub = stub
}

func (fake *FakeGenerator) GenerateSSHKeyArgsForCall(i int) string {
	fake.generateSSHKeyMutex.RLock()
	defer fake.generateSSHKeyMutex.RUnlock()
	argsForCall := fake.generateSSHKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGenerator) GenerateSSHKeyReturns(result1 credsgen.SSHKey, result2 error) {
	fake.generateSSHKeyMutex.Lock()
	defer fake.generateSSHKeyMutex.Unlock()
	fake.GenerateSSHKeyStub = nil
	fake.generateSSHKeyReturns = struct {
		result1 credsgen.SSHKey
		result2 error
	}{result1, result2}
}

func (fake *FakeGenerator) GenerateSSHKeyReturnsOnCall(i int, result1 credsgen.SSHKey, result2 error) {
	fake.generateSSHKeyMutex.Lock()
	defer fake.generateSSHKeyMutex.Unlock()
	fake.GenerateSSHKeyStub = nil
	if fake.generateSSHKeyReturnsOnCall == nil {
		fake.generateSSHKeyReturnsOnCall = make(map[int]struct {
			result1 credsgen.SSHKey
			result2 error
		})
	}
	fake.generateSSHKeyReturnsOnCall[i] = struct {
		result1 credsgen.SSHKey
		result2 error
	}{result1, result2}
}

func (fake *FakeGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateCertificateMutex.RLock()
	defer fake.generateCertificateMutex.RUnlock()
	fake.generateCertificateSigningRequestMutex.RLock()
	defer fake.generateCertificateSigningRequestMutex.RUnlock()
	fake.generatePasswordMutex.RLock()
	defer fake.generatePasswordMutex.RUnlock()
	fake.generateRSAKeyMutex.RLock()
	defer fake.generateRSAKeyMutex.RUnlock()
	fake.generateSSHKeyMutex.RLock()
	defer fake.generateSSHKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGenerator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ credsgen.Generator = new(FakeGenerator)
