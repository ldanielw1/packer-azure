// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package response

import (
	"github.com/MSOpenTech/packer-azure/packer/builder/azure/driver_restapi/response/model"
	"io"
)

func ParseStorageService(body io.ReadCloser) (*model.StorageService, error) {
	data, err := toModel(body, &model.StorageService{})

	if err != nil {
		return nil, err
	}

	m := data.(*model.StorageService)

	return m, nil
}
