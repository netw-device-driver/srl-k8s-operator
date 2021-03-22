/*
	Copyright 2021 Wim Henderickx.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package controllers

import (
	"math"
	"math/rand"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const maxBackOffCount = 10

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// actionResult is an interface that encapsulates the result of a Reconcile
// call, as returned by the action corresponding to the current state.
type actionResult interface {
	Result() (reconcile.Result, error)
	Dirty() bool
}

// actionError is a result indicating that an error occurred while attempting
// to advance the current action, and that reconciliation should be retried.
type actionError struct {
	err error
}

func (r actionError) Result() (result reconcile.Result, err error) {
	err = r.err
	return
}

func (r actionError) Dirty() bool {
	return false
}

// actionContinue is a result indicating that the current action is still
// in progress, and that the resource should remain in the same state
// waiting for the result.
type actionContinue struct {
	delay time.Duration
}

func (r actionContinue) Result() (result reconcile.Result, err error) {
	result.RequeueAfter = r.delay
	// Set Requeue true as well as RequeueAfter in case the delay is 0.
	if result.RequeueAfter == 0 {
		result.Requeue = true
	}
	result.Requeue = true
	return
}

func (r actionContinue) Dirty() bool {
	return false
}

// actionUpdate is a result indicating that the current action is still
// in progress, and that the resource should remain in the same provisioning
// state but write new Status data.
type actionUpdate struct {
	delay time.Duration
}

func (r actionUpdate) Result() (result reconcile.Result, err error) {
	result.RequeueAfter = r.delay
	// Set Requeue true as well as RequeueAfter in case the delay is 0.

	if result.RequeueAfter == 0 {
		result.Requeue = true
	}
	result.Requeue = true
	return
}

func (r actionUpdate) Dirty() bool {
	return true
}

// actionComplete is a result indicating that the current action has completed,
// and that the resource should transition to the next state.
type actionComplete struct {
}

func (r actionComplete) Result() (result reconcile.Result, err error) {
	return
}

func (r actionComplete) Dirty() bool {
	return true
}

// actionFailed is a result indicating that the current action has failed,
// and that the resource should be marked as in error.
type actionFailed struct {
	dirty      bool
	errorCount int
}

func (r actionFailed) Result() (result reconcile.Result, err error) {
	result.Requeue = true
	result.RequeueAfter = calculateBackoff(r.errorCount)
	return
}

func (r actionFailed) Dirty() bool {
	return r.dirty
}

// Distribution sample for errorCount values:
// 1  [1m, 2m]
// 2  [2m, 4m]
// 3  [4m, 8m]
// 4  [8m, 16m]
// 5  [16m, 32m]
// 6  [32m, 1h4m]
// 7  [1h4m, 2h8m]
// 8  [2h8m, 4h16m]
// 9  [4h16m, 8h32m]
func calculateBackoff(errorCount int) time.Duration {

	if errorCount > maxBackOffCount {
		errorCount = maxBackOffCount
	}

	base := math.Exp2(float64(errorCount))
	/* #nosec */
	backOff := base - (rand.Float64() * base * 0.5)
	backOffDuration := time.Minute * time.Duration(backOff)
	return backOffDuration
}
