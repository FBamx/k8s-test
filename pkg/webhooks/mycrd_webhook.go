package webhooks

import "k8s-test/pkg/webhooks/mycrd"

func init() {
	addHandler(mycrd.HandlerMap)
}
