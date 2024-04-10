// Copyright (c) 2020, NVIDIA CORPORATION.  All rights reserved.
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

package nvml

// nvml.EventSetCreate()
func (l *library) EventSetCreate() (EventSet, Return) {
	var Set nvmlEventSet
	ret := nvmlEventSetCreate(&Set)
	return Set, ret
}

// nvml.EventSetWait()
func (l *library) EventSetWait(Set EventSet, Timeoutms uint32) (EventData, Return) {
	return Set.Wait(Timeoutms)
}

func (Set nvmlEventSet) Wait(Timeoutms uint32) (EventData, Return) {
	var Data EventData
	ret := nvmlEventSetWait(Set, &Data, Timeoutms)
	return Data, ret
}

// nvml.EventSetFree()
func (l *library) EventSetFree(Set EventSet) Return {
	return Set.Free()
}

func (Set nvmlEventSet) Free() Return {
	return nvmlEventSetFree(Set)
}
