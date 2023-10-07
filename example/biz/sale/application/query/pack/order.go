//
// Copyright 2023 Bytedance Ltd. and/or its affiliates
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

package pack

import (
	"github.com/bytedance/dddfirework/example/biz/sale/infrastructure/po"
	"github.com/bytedance/dddfirework/example/common/dto/sale"
)

func MakeSimpleOrder(po *po.OrderPO) *sale.SimpleOrder {
	return &sale.SimpleOrder{
		ID:          po.ID,
		UserID:      po.User,
		TotalAmount: po.TotalAmount,
	}
}

func MakeSimpleOrderList(pos []*po.OrderPO) []*sale.SimpleOrder {
	result := make([]*sale.SimpleOrder, 0)
	for _, p := range pos {
		result = append(result, MakeSimpleOrder(p))
	}
	return result
}
