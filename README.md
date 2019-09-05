# Mattermost Marketplace ![CircleCI branch](https://img.shields.io/circleci/project/github/mattermost/mattermost-marketplace/master.svg)

The Mattermost Marketplace is a collection of plugins for use with [Mattermost](https://github.com/mattermost/mattermost-server). This repository houses the stateless HTTP service that will eventually run at marketplace.integrations.mattermost.com. It is meant to be queried by the Mattermost server to enable plugin discovery by system administrators.

Although Mattermost hosts the marketplace as an AWS Lambda function backed by S3 and CloudFront, the core feature set is designed for use in any hosting environment, enabling private, self-hosted collections of plugins.

Read more about the [Mattermost Plugins Marketplace Architecture](https://docs.google.com/document/d/1tVj0eNwMdIIGn8YoTs-cYz9NYvXjqx6bqWH-wa-yDLk/edit).

## Other Resources

This repository houses the open-source components of the Mattermost Plugins Marketplce. Other resources are linked below:

- [Mattermost the server and user interface](https://github.com/mattermost/mattermost-server)

## Get Involved

- [Join the discussion on ~Plugins Marketplace](https://community.mattermost.com/core/channels/plugins-marketplace)

## Developing

### Environment Setup

1. Install [Go](https://golang.org/doc/install)

### Running

Simply run the following:

```
$ make run-server
```

### Testing

Running all tests:

```
$ make test
```

### Deploying as a Lambda Function

In addition to running as a standalone server, the Marketplace is also designed to run as a Lambda function, compiling the `plugins.json` database into the binary for immediate access without further configuration.

Simply run the following:

```
$ make deploy-lambda
```

To iterate quickly after the Cloud Formation stack is up, simply run:

```
$ make deploy-lambda-fast
```
