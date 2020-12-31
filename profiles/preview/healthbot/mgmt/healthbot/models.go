// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package healthbot

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/healthbot/mgmt/2020-12-08/healthbot"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type IdentityType = original.IdentityType

const (
	Application     IdentityType = original.Application
	Key             IdentityType = original.Key
	ManagedIdentity IdentityType = original.ManagedIdentity
	User            IdentityType = original.User
)

type SkuName = original.SkuName

const (
	C0 SkuName = original.C0
	F0 SkuName = original.F0
	S1 SkuName = original.S1
)

type AvailableOperations = original.AvailableOperations
type AvailableOperationsIterator = original.AvailableOperationsIterator
type AvailableOperationsPage = original.AvailableOperationsPage
type BaseClient = original.BaseClient
type BotResponseList = original.BotResponseList
type BotResponseListIterator = original.BotResponseListIterator
type BotResponseListPage = original.BotResponseListPage
type BotsClient = original.BotsClient
type BotsCreateFuture = original.BotsCreateFuture
type BotsDeleteFuture = original.BotsDeleteFuture
type Error = original.Error
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorError = original.ErrorError
type HealthBot = original.HealthBot
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type Properties = original.Properties
type Resource = original.Resource
type Sku = original.Sku
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource
type UpdateParameters = original.UpdateParameters
type ValidationResult = original.ValidationResult

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailableOperationsIterator(page AvailableOperationsPage) AvailableOperationsIterator {
	return original.NewAvailableOperationsIterator(page)
}
func NewAvailableOperationsPage(cur AvailableOperations, getNextPage func(context.Context, AvailableOperations) (AvailableOperations, error)) AvailableOperationsPage {
	return original.NewAvailableOperationsPage(cur, getNextPage)
}
func NewBotResponseListIterator(page BotResponseListPage) BotResponseListIterator {
	return original.NewBotResponseListIterator(page)
}
func NewBotResponseListPage(cur BotResponseList, getNextPage func(context.Context, BotResponseList) (BotResponseList, error)) BotResponseListPage {
	return original.NewBotResponseListPage(cur, getNextPage)
}
func NewBotsClient(subscriptionID string) BotsClient {
	return original.NewBotsClient(subscriptionID)
}
func NewBotsClientWithBaseURI(baseURI string, subscriptionID string) BotsClient {
	return original.NewBotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
