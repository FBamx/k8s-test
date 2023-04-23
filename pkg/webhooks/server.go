package webhooks

import (
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	HandlerMap = map[string]admission.Handler{}
)

func addHandler(m map[string]admission.Handler) {
	for path, handler := range m {
		if len(path) == 0 {
			klog.Warningf("path is empty")
			continue
		}
		if _, find := HandlerMap[path]; find {
			klog.V(1).Infof("conflicting webhook builder path %v in handler map", path)
		}
		HandlerMap[path] = handler
	}
}

func SetupWithManager(mgr ctrl.Manager) error {
	server := mgr.GetWebhookServer()
	for path, handler := range HandlerMap {
		server.Register(path, &webhook.Admission{Handler: handler})
	}
	return nil
}
