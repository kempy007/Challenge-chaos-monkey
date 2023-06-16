

operator-sdk init --domain cypherpunk.io --repo github.com/kempy007/Challenge-chaos-monkey --plugins go.kubebuilder.io/v4-alpha --skip-go-version-check

operator-sdk create api --group=core --version=v1 --kind=ChaosMonkey --controller=true --resource=true