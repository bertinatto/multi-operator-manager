//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationMenuSpec) DeepCopyInto(out *ApplicationMenuSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationMenuSpec.
func (in *ApplicationMenuSpec) DeepCopy() *ApplicationMenuSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationMenuSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLIDownloadLink) DeepCopyInto(out *CLIDownloadLink) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLIDownloadLink.
func (in *CLIDownloadLink) DeepCopy() *CLIDownloadLink {
	if in == nil {
		return nil
	}
	out := new(CLIDownloadLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleCLIDownload) DeepCopyInto(out *ConsoleCLIDownload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleCLIDownload.
func (in *ConsoleCLIDownload) DeepCopy() *ConsoleCLIDownload {
	if in == nil {
		return nil
	}
	out := new(ConsoleCLIDownload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleCLIDownload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleCLIDownloadList) DeepCopyInto(out *ConsoleCLIDownloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleCLIDownload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleCLIDownloadList.
func (in *ConsoleCLIDownloadList) DeepCopy() *ConsoleCLIDownloadList {
	if in == nil {
		return nil
	}
	out := new(ConsoleCLIDownloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleCLIDownloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleCLIDownloadSpec) DeepCopyInto(out *ConsoleCLIDownloadSpec) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]CLIDownloadLink, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleCLIDownloadSpec.
func (in *ConsoleCLIDownloadSpec) DeepCopy() *ConsoleCLIDownloadSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleCLIDownloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleExternalLogLink) DeepCopyInto(out *ConsoleExternalLogLink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleExternalLogLink.
func (in *ConsoleExternalLogLink) DeepCopy() *ConsoleExternalLogLink {
	if in == nil {
		return nil
	}
	out := new(ConsoleExternalLogLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleExternalLogLink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleExternalLogLinkList) DeepCopyInto(out *ConsoleExternalLogLinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleExternalLogLink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleExternalLogLinkList.
func (in *ConsoleExternalLogLinkList) DeepCopy() *ConsoleExternalLogLinkList {
	if in == nil {
		return nil
	}
	out := new(ConsoleExternalLogLinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleExternalLogLinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleExternalLogLinkSpec) DeepCopyInto(out *ConsoleExternalLogLinkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleExternalLogLinkSpec.
func (in *ConsoleExternalLogLinkSpec) DeepCopy() *ConsoleExternalLogLinkSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleExternalLogLinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleLink) DeepCopyInto(out *ConsoleLink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleLink.
func (in *ConsoleLink) DeepCopy() *ConsoleLink {
	if in == nil {
		return nil
	}
	out := new(ConsoleLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleLink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleLinkList) DeepCopyInto(out *ConsoleLinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleLink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleLinkList.
func (in *ConsoleLinkList) DeepCopy() *ConsoleLinkList {
	if in == nil {
		return nil
	}
	out := new(ConsoleLinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleLinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleLinkSpec) DeepCopyInto(out *ConsoleLinkSpec) {
	*out = *in
	out.Link = in.Link
	if in.ApplicationMenu != nil {
		in, out := &in.ApplicationMenu, &out.ApplicationMenu
		*out = new(ApplicationMenuSpec)
		**out = **in
	}
	if in.NamespaceDashboard != nil {
		in, out := &in.NamespaceDashboard, &out.NamespaceDashboard
		*out = new(NamespaceDashboardSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleLinkSpec.
func (in *ConsoleLinkSpec) DeepCopy() *ConsoleLinkSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleLinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotification) DeepCopyInto(out *ConsoleNotification) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotification.
func (in *ConsoleNotification) DeepCopy() *ConsoleNotification {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleNotification) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationList) DeepCopyInto(out *ConsoleNotificationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleNotification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationList.
func (in *ConsoleNotificationList) DeepCopy() *ConsoleNotificationList {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleNotificationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleNotificationSpec) DeepCopyInto(out *ConsoleNotificationSpec) {
	*out = *in
	if in.Link != nil {
		in, out := &in.Link, &out.Link
		*out = new(Link)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleNotificationSpec.
func (in *ConsoleNotificationSpec) DeepCopy() *ConsoleNotificationSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleNotificationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePlugin) DeepCopyInto(out *ConsolePlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePlugin.
func (in *ConsolePlugin) DeepCopy() *ConsolePlugin {
	if in == nil {
		return nil
	}
	out := new(ConsolePlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsolePlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginBackend) DeepCopyInto(out *ConsolePluginBackend) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ConsolePluginService)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginBackend.
func (in *ConsolePluginBackend) DeepCopy() *ConsolePluginBackend {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginI18n) DeepCopyInto(out *ConsolePluginI18n) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginI18n.
func (in *ConsolePluginI18n) DeepCopy() *ConsolePluginI18n {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginI18n)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginList) DeepCopyInto(out *ConsolePluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsolePlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginList.
func (in *ConsolePluginList) DeepCopy() *ConsolePluginList {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsolePluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginProxy) DeepCopyInto(out *ConsolePluginProxy) {
	*out = *in
	in.Endpoint.DeepCopyInto(&out.Endpoint)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginProxy.
func (in *ConsolePluginProxy) DeepCopy() *ConsolePluginProxy {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginProxyEndpoint) DeepCopyInto(out *ConsolePluginProxyEndpoint) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ConsolePluginProxyServiceConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginProxyEndpoint.
func (in *ConsolePluginProxyEndpoint) DeepCopy() *ConsolePluginProxyEndpoint {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginProxyEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginProxyServiceConfig) DeepCopyInto(out *ConsolePluginProxyServiceConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginProxyServiceConfig.
func (in *ConsolePluginProxyServiceConfig) DeepCopy() *ConsolePluginProxyServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginProxyServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginService) DeepCopyInto(out *ConsolePluginService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginService.
func (in *ConsolePluginService) DeepCopy() *ConsolePluginService {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsolePluginSpec) DeepCopyInto(out *ConsolePluginSpec) {
	*out = *in
	in.Backend.DeepCopyInto(&out.Backend)
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = make([]ConsolePluginProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.I18n = in.I18n
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsolePluginSpec.
func (in *ConsolePluginSpec) DeepCopy() *ConsolePluginSpec {
	if in == nil {
		return nil
	}
	out := new(ConsolePluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleQuickStart) DeepCopyInto(out *ConsoleQuickStart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleQuickStart.
func (in *ConsoleQuickStart) DeepCopy() *ConsoleQuickStart {
	if in == nil {
		return nil
	}
	out := new(ConsoleQuickStart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleQuickStart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleQuickStartList) DeepCopyInto(out *ConsoleQuickStartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleQuickStart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleQuickStartList.
func (in *ConsoleQuickStartList) DeepCopy() *ConsoleQuickStartList {
	if in == nil {
		return nil
	}
	out := new(ConsoleQuickStartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleQuickStartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleQuickStartSpec) DeepCopyInto(out *ConsoleQuickStartSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Prerequisites != nil {
		in, out := &in.Prerequisites, &out.Prerequisites
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]ConsoleQuickStartTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NextQuickStart != nil {
		in, out := &in.NextQuickStart, &out.NextQuickStart
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessReviewResources != nil {
		in, out := &in.AccessReviewResources, &out.AccessReviewResources
		*out = make([]authorizationv1.ResourceAttributes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleQuickStartSpec.
func (in *ConsoleQuickStartSpec) DeepCopy() *ConsoleQuickStartSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleQuickStartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleQuickStartTask) DeepCopyInto(out *ConsoleQuickStartTask) {
	*out = *in
	if in.Review != nil {
		in, out := &in.Review, &out.Review
		*out = new(ConsoleQuickStartTaskReview)
		**out = **in
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(ConsoleQuickStartTaskSummary)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleQuickStartTask.
func (in *ConsoleQuickStartTask) DeepCopy() *ConsoleQuickStartTask {
	if in == nil {
		return nil
	}
	out := new(ConsoleQuickStartTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleQuickStartTaskReview) DeepCopyInto(out *ConsoleQuickStartTaskReview) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleQuickStartTaskReview.
func (in *ConsoleQuickStartTaskReview) DeepCopy() *ConsoleQuickStartTaskReview {
	if in == nil {
		return nil
	}
	out := new(ConsoleQuickStartTaskReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleQuickStartTaskSummary) DeepCopyInto(out *ConsoleQuickStartTaskSummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleQuickStartTaskSummary.
func (in *ConsoleQuickStartTaskSummary) DeepCopy() *ConsoleQuickStartTaskSummary {
	if in == nil {
		return nil
	}
	out := new(ConsoleQuickStartTaskSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSample) DeepCopyInto(out *ConsoleSample) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSample.
func (in *ConsoleSample) DeepCopy() *ConsoleSample {
	if in == nil {
		return nil
	}
	out := new(ConsoleSample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleSample) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleContainerImportSource) DeepCopyInto(out *ConsoleSampleContainerImportSource) {
	*out = *in
	out.Service = in.Service
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleContainerImportSource.
func (in *ConsoleSampleContainerImportSource) DeepCopy() *ConsoleSampleContainerImportSource {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleContainerImportSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleContainerImportSourceService) DeepCopyInto(out *ConsoleSampleContainerImportSourceService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleContainerImportSourceService.
func (in *ConsoleSampleContainerImportSourceService) DeepCopy() *ConsoleSampleContainerImportSourceService {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleContainerImportSourceService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleGitImportSource) DeepCopyInto(out *ConsoleSampleGitImportSource) {
	*out = *in
	out.Repository = in.Repository
	out.Service = in.Service
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleGitImportSource.
func (in *ConsoleSampleGitImportSource) DeepCopy() *ConsoleSampleGitImportSource {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleGitImportSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleGitImportSourceRepository) DeepCopyInto(out *ConsoleSampleGitImportSourceRepository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleGitImportSourceRepository.
func (in *ConsoleSampleGitImportSourceRepository) DeepCopy() *ConsoleSampleGitImportSourceRepository {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleGitImportSourceRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleGitImportSourceService) DeepCopyInto(out *ConsoleSampleGitImportSourceService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleGitImportSourceService.
func (in *ConsoleSampleGitImportSourceService) DeepCopy() *ConsoleSampleGitImportSourceService {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleGitImportSourceService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleList) DeepCopyInto(out *ConsoleSampleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleSample, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleList.
func (in *ConsoleSampleList) DeepCopy() *ConsoleSampleList {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleSampleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleSource) DeepCopyInto(out *ConsoleSampleSource) {
	*out = *in
	if in.GitImport != nil {
		in, out := &in.GitImport, &out.GitImport
		*out = new(ConsoleSampleGitImportSource)
		**out = **in
	}
	if in.ContainerImport != nil {
		in, out := &in.ContainerImport, &out.ContainerImport
		*out = new(ConsoleSampleContainerImportSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleSource.
func (in *ConsoleSampleSource) DeepCopy() *ConsoleSampleSource {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleSampleSpec) DeepCopyInto(out *ConsoleSampleSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleSampleSpec.
func (in *ConsoleSampleSpec) DeepCopy() *ConsoleSampleSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleSampleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleYAMLSample) DeepCopyInto(out *ConsoleYAMLSample) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleYAMLSample.
func (in *ConsoleYAMLSample) DeepCopy() *ConsoleYAMLSample {
	if in == nil {
		return nil
	}
	out := new(ConsoleYAMLSample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleYAMLSample) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleYAMLSampleList) DeepCopyInto(out *ConsoleYAMLSampleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConsoleYAMLSample, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleYAMLSampleList.
func (in *ConsoleYAMLSampleList) DeepCopy() *ConsoleYAMLSampleList {
	if in == nil {
		return nil
	}
	out := new(ConsoleYAMLSampleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConsoleYAMLSampleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleYAMLSampleSpec) DeepCopyInto(out *ConsoleYAMLSampleSpec) {
	*out = *in
	out.TargetResource = in.TargetResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleYAMLSampleSpec.
func (in *ConsoleYAMLSampleSpec) DeepCopy() *ConsoleYAMLSampleSpec {
	if in == nil {
		return nil
	}
	out := new(ConsoleYAMLSampleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceDashboardSpec) DeepCopyInto(out *NamespaceDashboardSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceDashboardSpec.
func (in *NamespaceDashboardSpec) DeepCopy() *NamespaceDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceDashboardSpec)
	in.DeepCopyInto(out)
	return out
}