# github-pr-sha-comment-resource

[![master](https://github.com/mtharrison/github-pr-sha-comment-resource/actions/workflows/master.yml/badge.svg)](https://github.com/mtharrison/github-pr-sha-comment-resource/actions/workflows/master.yml) [![release](https://github.com/mtharrison/github-pr-sha-comment-resource/actions/workflows/release.yml/badge.svg)](https://github.com/mtharrison/github-pr-sha-comment-resource/actions/workflows/release.yml) ![GitHub release](https://img.shields.io/github/v/release/mtharrison/github-pr-sha-comment-resource) ![report](https://goreportcard.com/badge/github.com/mtharrison/github-pr-sha-comment-resource)

A resource type for [Concourse CI](https://concourse-ci.org/) to comment on PRs based on the HEAD sha of a given git repo (useful for linking PRs back to merges)
 
---
 
## Usage
 
 To use register the resource type using the public [Docker image](https://hub.docker.com/repository/docker/mtharrison/github-pr-sha-comment-resource) `mtharrison/github-pr-sha-comment-resource`.

 ```yaml
 resource_types:
  - name: github-pr-sha-comment-resource
    type: docker-image
    source:
      repository: mtharrison/github-pr-sha-comment-resource
      tag: v0.2.0
 ```
 Then create a new resource using this resource type. You'll need to provide the `repository`, `access_token`, optionally a `v3_endpoint` if you're using Github Enterprise, otherwise this will default to the public Github API.
 
 ```yaml
resources:
  - name: github-pr-comment
    type: github-pr-sha-comment
    icon: github
    source:
      repository: golang/go
      access_token: '[...]'
      v3_endpoint: '[...]'
 ```
 Finally you can use the resource to post a comment to a PR
 ```yaml
jobs:
  - name: master-job-example
    plan:
      - get: some-repo
        trigger: true
      - put: github-pr-comment
        params:
            dir: master-job-example
            comment: "Triggered master build - [click to open pipeline](\${ATC_EXTERNAL_URL}/builds/\${BUILD_ID})"
 ```
