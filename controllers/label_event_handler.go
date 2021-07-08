package controllers

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MapToOwnerByLabel returns an event handler will convert events on a resource to requests on
// another resource whose name if found in a given label
// it maps the namespace to a request on the "owner" (or "associated") resource
// (if the label exists)
func MapToOwnerByLabel(namespace, label string) func(object client.Object) []reconcile.Request {
	return func(obj client.Object) []reconcile.Request {
		if name, exists := obj.GetLabels()[label]; exists {
			return []reconcile.Request{
				{
					NamespacedName: types.NamespacedName{
						Namespace: namespace,
						Name:      name,
					},
				},
			}
		}
		// the obj was not a namespace or it did not have the required label.
		return []reconcile.Request{}
	}
}
