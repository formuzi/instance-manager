### EKS Managed Node Group (alpha-1)

You can also provision EKS managed node groups by submitting a spec with a different provisioner.

```yaml
apiVersion: instancemgr.keikoproj.io/v1alpha1
kind: InstanceGroup
metadata:
  name: hello-world
  namespace: instance-manager
spec:
  strategy:
    type: managed
  provisioner: eks-managed
  eks-managed:
    maxSize: 6
    minSize: 3
    configuration:
      clusterName: my-eks-cluster
      labels:
        example.label.com/label: some-value
      volSize: 20
      nodeRole: arn:aws:iam::012345678910:role/basic-eks-role
      amiType: AL2_x86_64
      instanceType: m5.large
      keyPairName: my-ec2-key-pair
      securityGroups:
      - sg-04adb6343b07c7914
      subnets:
      - subnet-0bf9bc85fd80af561
      - subnet-0130025d2673de5e4
      - subnet-01a5c28e074c46580
      tags:
      - key: my-ec2-tag
        value: some-value
```