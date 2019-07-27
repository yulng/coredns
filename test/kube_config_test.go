package test

// kubeconfig is a fake kubernetes config that we inject during the readme test.
const kubeconfig = `
apiVersion: v1
kind: Config
clusters:
- name: test
  cluster:
    insecure-skip-tls-verify: true
    server: https://127.0.0.1:443
contexts:
- context:
    cluster: test
    user: admin
  name: test
current-context: test
users:
- name: test
  user:
    password: test
username: test
`
