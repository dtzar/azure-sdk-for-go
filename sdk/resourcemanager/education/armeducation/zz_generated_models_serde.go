//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeducation

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type GrantDetailProperties.
func (g GrantDetailProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allocatedBudget", g.AllocatedBudget)
	populateTimeRFC3339(objectMap, "effectiveDate", g.EffectiveDate)
	populateTimeRFC3339(objectMap, "expirationDate", g.ExpirationDate)
	populate(objectMap, "offerCap", g.OfferCap)
	populate(objectMap, "offerType", g.OfferType)
	populate(objectMap, "status", g.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GrantDetailProperties.
func (g *GrantDetailProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "allocatedBudget":
			err = unpopulate(val, "AllocatedBudget", &g.AllocatedBudget)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, "EffectiveDate", &g.EffectiveDate)
			delete(rawMsg, key)
		case "expirationDate":
			err = unpopulateTimeRFC3339(val, "ExpirationDate", &g.ExpirationDate)
			delete(rawMsg, key)
		case "offerCap":
			err = unpopulate(val, "OfferCap", &g.OfferCap)
			delete(rawMsg, key)
		case "offerType":
			err = unpopulate(val, "OfferType", &g.OfferType)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &g.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LabProperties.
func (l LabProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "budgetPerStudent", l.BudgetPerStudent)
	populate(objectMap, "description", l.Description)
	populate(objectMap, "displayName", l.DisplayName)
	populateTimeRFC3339(objectMap, "effectiveDate", l.EffectiveDate)
	populateTimeRFC3339(objectMap, "expirationDate", l.ExpirationDate)
	populate(objectMap, "invitationCode", l.InvitationCode)
	populate(objectMap, "maxStudentCount", l.MaxStudentCount)
	populate(objectMap, "status", l.Status)
	populate(objectMap, "totalAllocatedBudget", l.TotalAllocatedBudget)
	populate(objectMap, "totalBudget", l.TotalBudget)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LabProperties.
func (l *LabProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "budgetPerStudent":
			err = unpopulate(val, "BudgetPerStudent", &l.BudgetPerStudent)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &l.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &l.DisplayName)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, "EffectiveDate", &l.EffectiveDate)
			delete(rawMsg, key)
		case "expirationDate":
			err = unpopulateTimeRFC3339(val, "ExpirationDate", &l.ExpirationDate)
			delete(rawMsg, key)
		case "invitationCode":
			err = unpopulate(val, "InvitationCode", &l.InvitationCode)
			delete(rawMsg, key)
		case "maxStudentCount":
			err = unpopulate(val, "MaxStudentCount", &l.MaxStudentCount)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &l.Status)
			delete(rawMsg, key)
		case "totalAllocatedBudget":
			err = unpopulate(val, "TotalAllocatedBudget", &l.TotalAllocatedBudget)
			delete(rawMsg, key)
		case "totalBudget":
			err = unpopulate(val, "TotalBudget", &l.TotalBudget)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StudentLabProperties.
func (s StudentLabProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "budget", s.Budget)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populateTimeRFC3339(objectMap, "effectiveDate", s.EffectiveDate)
	populateTimeRFC3339(objectMap, "expirationDate", s.ExpirationDate)
	populate(objectMap, "labScope", s.LabScope)
	populate(objectMap, "role", s.Role)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StudentLabProperties.
func (s *StudentLabProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "budget":
			err = unpopulate(val, "Budget", &s.Budget)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &s.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &s.DisplayName)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, "EffectiveDate", &s.EffectiveDate)
			delete(rawMsg, key)
		case "expirationDate":
			err = unpopulateTimeRFC3339(val, "ExpirationDate", &s.ExpirationDate)
			delete(rawMsg, key)
		case "labScope":
			err = unpopulate(val, "LabScope", &s.LabScope)
			delete(rawMsg, key)
		case "role":
			err = unpopulate(val, "Role", &s.Role)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &s.SubscriptionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StudentProperties.
func (s StudentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "budget", s.Budget)
	populateTimeRFC3339(objectMap, "effectiveDate", s.EffectiveDate)
	populate(objectMap, "email", s.Email)
	populateTimeRFC3339(objectMap, "expirationDate", s.ExpirationDate)
	populate(objectMap, "firstName", s.FirstName)
	populate(objectMap, "lastName", s.LastName)
	populate(objectMap, "role", s.Role)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "subscriptionAlias", s.SubscriptionAlias)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	populateTimeRFC3339(objectMap, "subscriptionInviteLastSentDate", s.SubscriptionInviteLastSentDate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StudentProperties.
func (s *StudentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "budget":
			err = unpopulate(val, "Budget", &s.Budget)
			delete(rawMsg, key)
		case "effectiveDate":
			err = unpopulateTimeRFC3339(val, "EffectiveDate", &s.EffectiveDate)
			delete(rawMsg, key)
		case "email":
			err = unpopulate(val, "Email", &s.Email)
			delete(rawMsg, key)
		case "expirationDate":
			err = unpopulateTimeRFC3339(val, "ExpirationDate", &s.ExpirationDate)
			delete(rawMsg, key)
		case "firstName":
			err = unpopulate(val, "FirstName", &s.FirstName)
			delete(rawMsg, key)
		case "lastName":
			err = unpopulate(val, "LastName", &s.LastName)
			delete(rawMsg, key)
		case "role":
			err = unpopulate(val, "Role", &s.Role)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "subscriptionAlias":
			err = unpopulate(val, "SubscriptionAlias", &s.SubscriptionAlias)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &s.SubscriptionID)
			delete(rawMsg, key)
		case "subscriptionInviteLastSentDate":
			err = unpopulateTimeRFC3339(val, "SubscriptionInviteLastSentDate", &s.SubscriptionInviteLastSentDate)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
