// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeBuildpackRepository struct {
	FindByNameStub        func(name string) (buildpack models.Buildpack, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.Buildpack
		result2 error
	}
	findByNameReturnsOnCall map[int]struct {
		result1 models.Buildpack
		result2 error
	}
	FindByNameAndStackStub        func(name, stack string) (buildpack models.Buildpack, apiErr error)
	findByNameAndStackMutex       sync.RWMutex
	findByNameAndStackArgsForCall []struct {
		name  string
		stack string
	}
	findByNameAndStackReturns struct {
		result1 models.Buildpack
		result2 error
	}
	findByNameAndStackReturnsOnCall map[int]struct {
		result1 models.Buildpack
		result2 error
	}
	ListBuildpacksStub        func(func(models.Buildpack) bool) error
	listBuildpacksMutex       sync.RWMutex
	listBuildpacksArgsForCall []struct {
		arg1 func(models.Buildpack) bool
	}
	listBuildpacksReturns struct {
		result1 error
	}
	listBuildpacksReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(name string, position *int, enabled *bool, locked *bool) (createdBuildpack models.Buildpack, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name     string
		position *int
		enabled  *bool
		locked   *bool
	}
	createReturns struct {
		result1 models.Buildpack
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 models.Buildpack
		result2 error
	}
	DeleteStub        func(buildpackGUID string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		buildpackGUID string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(buildpack models.Buildpack) (updatedBuildpack models.Buildpack, apiErr error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		buildpack models.Buildpack
	}
	updateReturns struct {
		result1 models.Buildpack
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 models.Buildpack
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildpackRepository) FindByName(name string) (buildpack models.Buildpack, apiErr error) {
	fake.findByNameMutex.Lock()
	ret, specificReturn := fake.findByNameReturnsOnCall[len(fake.findByNameArgsForCall)]
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindByName", []interface{}{name})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findByNameReturns.result1, fake.findByNameReturns.result2
}

func (fake *FakeBuildpackRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeBuildpackRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeBuildpackRepository) FindByNameReturns(result1 models.Buildpack, result2 error) {
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) FindByNameReturnsOnCall(i int, result1 models.Buildpack, result2 error) {
	fake.FindByNameStub = nil
	if fake.findByNameReturnsOnCall == nil {
		fake.findByNameReturnsOnCall = make(map[int]struct {
			result1 models.Buildpack
			result2 error
		})
	}
	fake.findByNameReturnsOnCall[i] = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) FindByNameAndStack(name string, stack string) (buildpack models.Buildpack, apiErr error) {
	fake.findByNameAndStackMutex.Lock()
	ret, specificReturn := fake.findByNameAndStackReturnsOnCall[len(fake.findByNameAndStackArgsForCall)]
	fake.findByNameAndStackArgsForCall = append(fake.findByNameAndStackArgsForCall, struct {
		name  string
		stack string
	}{name, stack})
	fake.recordInvocation("FindByNameAndStack", []interface{}{name, stack})
	fake.findByNameAndStackMutex.Unlock()
	if fake.FindByNameAndStackStub != nil {
		return fake.FindByNameAndStackStub(name, stack)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findByNameAndStackReturns.result1, fake.findByNameAndStackReturns.result2
}

func (fake *FakeBuildpackRepository) FindByNameAndStackCallCount() int {
	fake.findByNameAndStackMutex.RLock()
	defer fake.findByNameAndStackMutex.RUnlock()
	return len(fake.findByNameAndStackArgsForCall)
}

func (fake *FakeBuildpackRepository) FindByNameAndStackArgsForCall(i int) (string, string) {
	fake.findByNameAndStackMutex.RLock()
	defer fake.findByNameAndStackMutex.RUnlock()
	return fake.findByNameAndStackArgsForCall[i].name, fake.findByNameAndStackArgsForCall[i].stack
}

func (fake *FakeBuildpackRepository) FindByNameAndStackReturns(result1 models.Buildpack, result2 error) {
	fake.FindByNameAndStackStub = nil
	fake.findByNameAndStackReturns = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) FindByNameAndStackReturnsOnCall(i int, result1 models.Buildpack, result2 error) {
	fake.FindByNameAndStackStub = nil
	if fake.findByNameAndStackReturnsOnCall == nil {
		fake.findByNameAndStackReturnsOnCall = make(map[int]struct {
			result1 models.Buildpack
			result2 error
		})
	}
	fake.findByNameAndStackReturnsOnCall[i] = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) ListBuildpacks(arg1 func(models.Buildpack) bool) error {
	fake.listBuildpacksMutex.Lock()
	ret, specificReturn := fake.listBuildpacksReturnsOnCall[len(fake.listBuildpacksArgsForCall)]
	fake.listBuildpacksArgsForCall = append(fake.listBuildpacksArgsForCall, struct {
		arg1 func(models.Buildpack) bool
	}{arg1})
	fake.recordInvocation("ListBuildpacks", []interface{}{arg1})
	fake.listBuildpacksMutex.Unlock()
	if fake.ListBuildpacksStub != nil {
		return fake.ListBuildpacksStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.listBuildpacksReturns.result1
}

func (fake *FakeBuildpackRepository) ListBuildpacksCallCount() int {
	fake.listBuildpacksMutex.RLock()
	defer fake.listBuildpacksMutex.RUnlock()
	return len(fake.listBuildpacksArgsForCall)
}

func (fake *FakeBuildpackRepository) ListBuildpacksArgsForCall(i int) func(models.Buildpack) bool {
	fake.listBuildpacksMutex.RLock()
	defer fake.listBuildpacksMutex.RUnlock()
	return fake.listBuildpacksArgsForCall[i].arg1
}

func (fake *FakeBuildpackRepository) ListBuildpacksReturns(result1 error) {
	fake.ListBuildpacksStub = nil
	fake.listBuildpacksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildpackRepository) ListBuildpacksReturnsOnCall(i int, result1 error) {
	fake.ListBuildpacksStub = nil
	if fake.listBuildpacksReturnsOnCall == nil {
		fake.listBuildpacksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listBuildpacksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildpackRepository) Create(name string, position *int, enabled *bool, locked *bool) (createdBuildpack models.Buildpack, apiErr error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name     string
		position *int
		enabled  *bool
		locked   *bool
	}{name, position, enabled, locked})
	fake.recordInvocation("Create", []interface{}{name, position, enabled, locked})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name, position, enabled, locked)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeBuildpackRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBuildpackRepository) CreateArgsForCall(i int) (string, *int, *bool, *bool) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name, fake.createArgsForCall[i].position, fake.createArgsForCall[i].enabled, fake.createArgsForCall[i].locked
}

func (fake *FakeBuildpackRepository) CreateReturns(result1 models.Buildpack, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) CreateReturnsOnCall(i int, result1 models.Buildpack, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 models.Buildpack
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) Delete(buildpackGUID string) (apiErr error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		buildpackGUID string
	}{buildpackGUID})
	fake.recordInvocation("Delete", []interface{}{buildpackGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(buildpackGUID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeBuildpackRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeBuildpackRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].buildpackGUID
}

func (fake *FakeBuildpackRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildpackRepository) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildpackRepository) Update(buildpack models.Buildpack) (updatedBuildpack models.Buildpack, apiErr error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		buildpack models.Buildpack
	}{buildpack})
	fake.recordInvocation("Update", []interface{}{buildpack})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(buildpack)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *FakeBuildpackRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeBuildpackRepository) UpdateArgsForCall(i int) models.Buildpack {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].buildpack
}

func (fake *FakeBuildpackRepository) UpdateReturns(result1 models.Buildpack, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) UpdateReturnsOnCall(i int, result1 models.Buildpack, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 models.Buildpack
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildpackRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	fake.findByNameAndStackMutex.RLock()
	defer fake.findByNameAndStackMutex.RUnlock()
	fake.listBuildpacksMutex.RLock()
	defer fake.listBuildpacksMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildpackRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.BuildpackRepository = new(FakeBuildpackRepository)
