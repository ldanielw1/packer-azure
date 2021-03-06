// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.

package request

import (
	"fmt"
)

func (m *Manager) DeleteRole(serviceName, vmName string) *Data {

	uri := fmt.Sprintf("https://management.core.windows.net/%s/services/hostedservices/%s/deployments/%s/roles/%s", m.SubscrId, serviceName, vmName, vmName)

	data := &Data{
		Verb: "DELETE",
		Uri:  uri,
	}

	return data
}
