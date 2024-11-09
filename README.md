# Contributing

See the [Contributing.md](./.github/CONTRIBUTING.md) for more
information on how to contribute.

# Working Group Charter

This repository is maintained by [App Runtime
Platform](https://github.com/cloudfoundry/community/blob/main/toc/working-groups/app-runtime-platform.md)
under `Networking` area.

## Healthchecker release

This repository is a [BOSH](https://github.com/cloudfoundry/bosh)
release for `healthchecker` that is a go executable designed to perform
TCP/HTTP based health checks of processes managed by `monit` in BOSH
releases. Since the version of `monit` included in BOSH does not support
specific tcp/http health checks, we designed this utility to perform
health checking and restart processes if they become unreachable.

# Docs

-   [How to use](./docs/01-how-to-use.md)

> \[!IMPORTANT\]
>
> Content in this file is managed by the [CI task
> `sync-readme`](https://github.com/cloudfoundry/wg-app-platform-runtime-ci/blob/main/shared/tasks/sync-readme/metadata.yml)
> and is generated by CI following a convention.
