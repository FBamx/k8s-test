package mycrd

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

//+kubebuilder:webhook:path=/mutate-test-io-elisii-com-v1-mycrd,mutating=true,failurePolicy=fail,sideEffects=None,groups=test.io.elisii.com,resources=mycrds,verbs=create;update,versions=v1,name=mmycrd.kb.io,admissionReviewVersions=v1

type Mutating struct {
	Client  client.Client
	decoder *admission.Decoder
}

var _ admission.Handler = &Mutating{}

func (m *Mutating) Handle(ctx context.Context, req admission.Request) admission.Response {
	return admission.Allowed("ok")
}

var _ admission.DecoderInjector = &Mutating{}

func (m *Mutating) InjectDecoder(d *admission.Decoder) error {
	m.decoder = d
	return nil
}
