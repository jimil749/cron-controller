/*
Copyright 2021.

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

package v1

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous hasn't finished
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	//+kubebuilder:validation:MinLength=0

	// Schedule in Cron format
	Schedule string `json:"schedule"`

	//+kubebuilder:validation:Minimum=0

	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	Suspend *bool `json:"suspend,omitempty"`

	// JobTemplate specifies the job that will be created when executing a CronJob
	JobTemplate batchv1beta1.JobTemplateSpec `json:"jobTemplate"`

	//+kubebuilder:validation:Minimum=0

	// SuccessfulJobsHistoryLimit number of successful finished jobs to retain.
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	//+kubebuilder:validation:Minimum=0

	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of pointers to currently running jobs
	Active []corev1.ObjectReference `json:"active,omitempty"`

	// Information when was the last time the job was successfully scheduled.
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
