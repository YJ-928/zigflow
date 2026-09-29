/*
 * Copyright 2025 - 2026 Zigflow authors <https://github.com/zigflow/zigflow/graphs/contributors>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package runworkflowinput checks that a `run.workflow` child receives the
// declared `input`, evaluated in the parent, rather than the parent's own
// input.
package runworkflowinput

import (
	"testing"

	"github.com/zigflow/zigflow/tests/e2e/utils"
)

var testCase = utils.TestCase{
	Name:         "run-workflow-input",
	WorkflowPath: "workflow.yaml",
	Test: func(t *testing.T, test *utils.TestCase) {
		utils.RunToCompletionNamed[map[string]any](
			t,
			"run-workflow-input", "placeOrder",
			map[string]any{"orderId": 7},
			map[string]any{
				"greeted": float64(42),
				// Only on the parent's input, so the child must not see it.
				"orderId": nil,
			},
		)
	},
}

func init() {
	utils.AddTestCase(&testCase)
}
