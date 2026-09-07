# Provider Vcd

`provider-vcd` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Vcd API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/arkilasystems/provider-vcd):
```
up ctp provider install xpkg.upbound.io/arkilasystems/provider-vcd:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-vcd
spec:
  package: xpkg.upbound.io/arkilasystems/provider-vcd:v0.1.0
EOF
```

You can see the API reference [here](https://doc.crds.dev/github.com/arkilasystems/provider-vcd).

## Cluster-scoped vs. namespaced resources

`provider-vcd` supports [Crossplane v2](https://docs.crossplane.io/latest/whats-new/)-style
namespaced managed resources alongside the original cluster-scoped API, and continues to work
against Crossplane v1 clusters using the cluster-scoped API only.

- **Cluster-scoped** (legacy): group `vcd.upbound.io`, e.g. `vcd.upbound.io/v1alpha1` for resources
  and `vcd.upbound.io/v1beta1` for `ProviderConfig`. Works on both Crossplane v1 and v2.
- **Namespaced**: group `vcd.m.upbound.io`, e.g. `vcd.m.upbound.io/v1alpha1` for resources and
  `vcd.m.upbound.io/v1beta1` for `ProviderConfig`/`ClusterProviderConfig`. Requires Crossplane v2.

Both APIs are generated from the same set of vCD resources and installed by the same provider
package — pick whichever scope fits your cluster and Crossplane version, or use both side by side.

Namespaced managed resources reference a provider config via `spec.providerConfigRef`, which
defaults to `kind: ClusterProviderConfig` (cluster-scoped, matching today's single shared
`ProviderConfig` model) but can instead target `kind: ProviderConfig` for a config local to the
resource's own namespace, e.g.:

```yaml
spec:
  providerConfigRef:
    kind: ProviderConfig # or ClusterProviderConfig (default)
    name: default
```

See [`examples/`](examples/) for a cluster-scoped and `*-namespaced.yaml` example for every
supported resource, and [`examples/providerconfig/`](examples/providerconfig/) for
`ProviderConfig`, `ClusterProviderConfig`, and a namespaced `ProviderConfig` example.

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/arkilasystems/provider-vcd/issues).
