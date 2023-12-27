/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	"github.com/xianfuhui/fabric/common/channelconfig"
)

func TestOrdererConfigInterface(t *testing.T) {
	_ = channelconfig.Orderer(&Orderer{})
}
