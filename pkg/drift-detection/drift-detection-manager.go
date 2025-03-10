// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2023. projectsveltos.io. All rights reserved.

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
package driftdetection

var driftDetectionYAML = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: projectsveltos
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: drift-detection-manager
  namespace: projectsveltos
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: drift-detection-manager-role
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - create
  - get
  - list
  - update
  - watch
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
- apiGroups:
  - lib.projectsveltos.io
  resources:
  - debuggingconfigurations
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - lib.projectsveltos.io
  resources:
  - resourcesummaries
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - lib.projectsveltos.io
  resources:
  - resourcesummaries/finalizers
  verbs:
  - update
- apiGroups:
  - lib.projectsveltos.io
  resources:
  - resourcesummaries/status
  verbs:
  - get
  - patch
  - update
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: drift-detection-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: drift-detection-manager-role
subjects:
- kind: ServiceAccount
  name: drift-detection-manager
  namespace: projectsveltos
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: drift-detection-manager
  name: drift-detection-manager
  namespace: projectsveltos
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: drift-detection-manager
  template:
    metadata:
      annotations:
        kubectl.kubernetes.io/default-container: manager
      labels:
        control-plane: drift-detection-manager
    spec:
      containers:
      - args:
        - --diagnostics-address=:8443
        - --v=5
        - --cluster-namespace=
        - --cluster-name=
        - --cluster-type=
        - --current-cluster=managed-cluster
        - --run-mode=do-not-send-updates
        - --version=v0.50.0
        command:
        - /manager
        image: docker.io/projectsveltos/drift-detection-manager@sha256:a5a460f8bb5aa514aeb8a40475bd8d8aa9464a809bb58f3778a42936a9f2129c
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /healthz
            port: healthz
            scheme: HTTP
          initialDelaySeconds: 15
          periodSeconds: 20
        name: manager
        ports:
        - containerPort: 8443
          name: metrics
          protocol: TCP
        - containerPort: 9440
          name: healthz
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /readyz
            port: healthz
            scheme: HTTP
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          limits:
            cpu: 500m
            memory: 512Mi
          requests:
            cpu: 10m
            memory: 128Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - ALL
      securityContext:
        runAsNonRoot: true
      serviceAccountName: drift-detection-manager
      terminationGracePeriodSeconds: 10
`)
