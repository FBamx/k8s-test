package mycrd

import "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

var (
	HandlerMap = map[string]admission.Handler{
		"/mutate-test-io-elisii-com-v1-mycrd":   &Mutating{},
		"/validate-test-io-elisii-com-v1-mycrd": &Validating{},
	}
)
