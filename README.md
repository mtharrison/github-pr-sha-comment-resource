# github-pr-comment-resource

A resource type for [Concourse CI](https://concourse-ci.org/) to trigger builds from Github PR comments. Also allows parameters to be provided in comments.

 ![master](https://github.com/mtharrison/github-pr-comment-resource/workflows/Master/badge.svg?branch=master) ![release](https://img.shields.io/github/v/release/mtharrison/github-pr-comment-resource)
 
 ---
 
 ## Usage
 
 ```yaml
resource_types:
- name: github-pr-comment-resource
  type: docker-image
  source:
    repository: mtharrison/github-pr-comment-resource
    tag: v0.2.0

resources:
- name: deployment-trigger
  type: github-pr-comment-resource
  icon: github
  source:
    repository: "golang/go"
    access_token: "[...]"
    v3_endpoint: "[...]"                          # optional - for github enterprise users
    regex: "^deploy (\\w+) to (\\w+) please$"     # optional
    
jobs:
- name: deployment-test
  plan:
  - get: deployment-trigger
    trigger: true
  - task: echo
    config:
      image_resource:
        type: registry-image
        source:
          repository: alpine
      inputs:
      - name: deployment-trigger
      platform: linux
      run:
        path: /bin/sh
        args:
        - -c
        - |
          apk add jq &> /dev/null
          cat deployment-trigger/comment.json | jq

 ```
