package mycrd

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

//+kubebuilder:webhook:path=/validate-test-io-elisii-com-v1-mycrd,mutating=false,failurePolicy=fail,sideEffects=None,groups=test.io.elisii.com,resources=mycrds,verbs=create;update,versions=v1,name=vmycrd.kb.io,admissionReviewVersions=v1

type Validating struct {
	Client  client.Client
	decoder *admission.Decoder
}

var _ admission.Handler = &Validating{}

func (v *Validating) Handle(ctx context.Context, req admission.Request) admission.Response {
	return admission.Allowed("ok")
}

var _ admission.DecoderInjector = &Validating{}

func (v *Validating) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
