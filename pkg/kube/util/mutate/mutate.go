package mutate

import (
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	qjv1a1 "code.cloudfoundry.org/quarks-job/pkg/kube/apis/quarksjob/v1alpha1"
	bdv1 "code.cloudfoundry.org/quarks-operator/pkg/kube/apis/boshdeployment/v1alpha1"
	qsv1a1 "code.cloudfoundry.org/quarks-operator/pkg/kube/apis/quarkssecret/v1alpha1"
	qstsv1a1 "code.cloudfoundry.org/quarks-operator/pkg/kube/apis/quarksstatefulset/v1alpha1"
)

// BoshDeploymentMutateFn returns MutateFn which mutates BoshDeployment including:
// - labels, annotations
// - spec
func BoshDeploymentMutateFn(boshDeployment *bdv1.BOSHDeployment) controllerutil.MutateFn {
	updated := boshDeployment.DeepCopy()
	return func() error {
		boshDeployment.Labels = updated.Labels
		boshDeployment.Annotations = updated.Annotations
		boshDeployment.Spec = updated.Spec
		return nil
	}
}

// QuarksStatefulSetMutateFn returns MutateFn which mutates QuarksStatefulSet including:
// - labels, annotations
// - spec
func QuarksStatefulSetMutateFn(qSts *qstsv1a1.QuarksStatefulSet) controllerutil.MutateFn {
	updated := qSts.DeepCopy()
	return func() error {
		qSts.Labels = updated.Labels
		qSts.Annotations = updated.Annotations
		qSts.Spec = updated.Spec
		return nil
	}
}

// StatefulSetMutateFn returns MutateFn which mutates StatefulSet including:
// - labels, annotations
// - spec
func StatefulSetMutateFn(sfs *appsv1.StatefulSet) controllerutil.MutateFn {
	updated := sfs.DeepCopy()
	return func() error {
		sfs.Labels = updated.Labels
		sfs.Annotations = updated.Annotations
		sfs.Spec.Replicas = updated.Spec.Replicas
		sfs.Spec.Template = updated.Spec.Template
		sfs.Spec.UpdateStrategy = updated.Spec.UpdateStrategy
		return nil
	}
}

// QuarksJobMutateFn returns MutateFn which mutates QuarksJob including:
// - annotations and trigger strategy if empty
// - labels
// - spec.output, spec.Template, spec.updateOnConfigChange
func QuarksJobMutateFn(qJob *qjv1a1.QuarksJob) controllerutil.MutateFn {
	updated := qJob.DeepCopy()
	return func() error {
		qJob.Labels = updated.Labels
		// Does not reset Annotations
		if qJob.ObjectMeta.Annotations == nil {
			qJob.ObjectMeta.Annotations = updated.ObjectMeta.Annotations
		}
		// Does not reset Spec.Trigger.Strategy
		if len(qJob.Spec.Trigger.Strategy) == 0 {
			qJob.Spec.Trigger.Strategy = updated.Spec.Trigger.Strategy
		}
		qJob.Spec.Output = updated.Spec.Output
		qJob.Spec.Template = updated.Spec.Template
		qJob.Spec.UpdateOnConfigChange = updated.Spec.UpdateOnConfigChange
		return nil
	}
}

// QuarksSecretMutateFn returns MutateFn which mutates QuarksSecret including:
// - labels, annotations
// - spec
func QuarksSecretMutateFn(qSec *qsv1a1.QuarksSecret) controllerutil.MutateFn {
	updated := qSec.DeepCopy()
	return func() error {
		qSec.Labels = updated.Labels
		qSec.Annotations = updated.Annotations
		qSec.Spec = updated.Spec

		return nil
	}
}

// SecretMutateFn returns MutateFn which mutates Secret including:
// - labels, annotations
// - stringData
func SecretMutateFn(s *corev1.Secret) controllerutil.MutateFn {
	updated := s.DeepCopy()
	return func() error {
		s.Labels = updated.Labels
		s.Annotations = updated.Annotations
		for key, data := range updated.StringData {
			// Update once one of data has been changed
			oriData, ok := s.Data[key]
			if ok && reflect.DeepEqual(string(oriData), data) {
				continue
			} else {
				s.StringData = updated.StringData
				break
			}
		}
		return nil
	}
}

// ServiceMutateFn returns MutateFn which mutates Service including:
// - labels, annotations
// - spec.ports, spec.selector
func ServiceMutateFn(svc *corev1.Service) controllerutil.MutateFn {
	updated := svc.DeepCopy()
	return func() error {
		svc.Labels = updated.Labels
		svc.Annotations = updated.Annotations
		// Should keep the existing ClusterIP
		svc.Spec.Ports = updated.Spec.Ports
		svc.Spec.Selector = updated.Spec.Selector
		return nil
	}
}
