# LOGFAIRY

Client to handle [Graylog2](https://www.graylog.org/) features throught API build on top of [cobra](https://github.com/spf13/cobra).

## Prerequisites

In order to user LOGFAIRY the user used to call the commands must be enable to call the actions. See `Give permissions` to know more about it.

## Give permissions

In order to request a token to use against graylog2 server the user needs `users:tokenlist`, `users:tokencreate` and `users:tokenremove` permissions. To achieve this create a role to and assign it to the user that will make the requests.

```shell
$ curl -v -XPOST -u ADMIN:PASSWORD \
  -H 'Content-Type: application/json' \
  'http://<host>:<port>/api/roles' \
  -d "{\
    'read_only': false, \
    'permissions': ['users:tokenlist', 'users:tokencreate', 'users:tokenremove'], \
    'name': 'token access', \
    'description': 'Permission to query for token, mandatatory to use api' \
    }"
```

To handle streams the user must master it, this means being able to create it and have access to all the straems in the node, the user needs `streams:create`, `streams:read`,`streams:edit` and `streams:changestate` privileges.

```shell
$ curl -v -XPOST -u ADMIN:PASSWORD \
  -H 'Content-Type: application/json' \
  'http://<host>:<port>/api/roles' \
  -d "{\
    'read_only': false, \
    'permissions': ['streams:create', 'streams:read', 'streams:edit', 'streams:changestate'], \
    'name': 'stream master', \
    'description': 'Permission to master streams' \
    }"
```

In the same way to master dashboards the user must have the following permissions: `dashboards:create`, `dashboards:read` and `dashboards:edit`.

```shell
$ curl -v -XPOST -u ADMIN:PASSWORD \
  -H 'Content-Type: application/json' \
  'http://<host>:<port>/api/roles' \
  -d "{\
    'read_only': false, \
    'permissions': ['dashboards:create', 'dashboards:read', 'dashboards:edit'], \
    'name': 'dashboard master', \
    'description': 'Permission to master dashboards' \
    }"
```

for more information about creating users and assigning roles [read the doc](http://docs.graylog.org/en/2.1/pages/users_and_roles/system_users.html).

## Installing

```shell
go install github.com/uniplaces/logfairy
```

then add logfairy to your $PATH.

## Run the commands

LOGFAIRY reads from the environment `GRAYLOG_USERNAME` and `GRAYLOG_PASSWORD` and check for a config file in the the directory where it runs in order to find the `base_url` of the server and the `timeout` for the request. The config searched config file is named `graylog.yaml` and looks like this:

```yaml
client:
  base_url: "<graylog-host>"
  timeout: 5
```

### List of commands

```shell
logfairy goal is to standardize the way streams, dashboard and widget are created

Usage:
  logfairy [command]

Available Commands:
  bulk        bulk allows to create a set of streams, dashboards and widgets
  dashboard   handle dashboards actions
  help        Help about any command
  stream      handle stream actions
  widget      handle widgets actions

Flags:
  -h, --help   help for logfairy
```

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct and the process for submitting pull requests to us.

## Authors

Made with :heart: at [uniplaces](www.uniplaces.com)

## Release notes

| Release date | Description | Release notes         |
| ------------ | ----------- | --------------------- |
| <dd/mm/YY>   | < desc >    | [Release notes](link) |

## Licence

Copyright 2016 UNIPLACES

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
