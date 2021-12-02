package webhook

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/validate-v1-mapping,mutating=false,failurePolicy=fail,groups=*,resources=contextsystemmappings,verbs=create;update,versions=v1,name=webhook.beta.example.com,admissionReviewVersions=v1,sideEffects=none

// podValidator validates Pods
type PodValidator struct {
	Client  client.Client
	decoder *admission.Decoder
}

// podValidator admits a pod if a specific annotation exists.
func (v *PodValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	logger := log.FromContext(ctx)
	//contextsystemmapping := &betav1.ContextSystemMappings{}

	logger.Info(">>>>>>>>>>>>>>")
	return admission.Denied(fmt.Sprintf("Test"))

	//err := v.decoder.Decode(req, contextsystemmapping)
	//if err != nil {
	//	return admission.Errored(http.StatusBadRequest, err)
	//}

	//logger.Info(">>>>>>>>>>>>>>")
	////logger.Info("%v", contextsystemmapping)
	//logger.Info(">>>>>>>>>>>>>>")
	//return admission.Denied(fmt.Sprintf("Test"))
	//var contexts ContextList
	//if err := v.Client.List(ctx, &contexts, client.InNamespace(req.Namespace)); err != nil {
	//return admission.Denied(fmt.Sprintf("Context with id %d does not exist", contextsystemmapping.Spec.ContextID))
	//}

	//for _, k := range contexts.Items {
	//fmt.Println(k.Spec)
	//}

	//var systems SystemList
	//if err := v.Client.List(ctx, &systems, client.InNamespace(req.Namespace)); err != nil {
	//logMsg := fmt.Sprintf("System with id %d does not exist", contextsystemmapping.Spec.SystemID)
	//return admission.Denied(logMsg)
	//}

	//for _, k := range systems.Items {
	//fmt.Println(k.Spec)
	//}

	//return admission.Allowed("")
}

// podValidator implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (v *PodValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
