# gangway

gangway is an open source tool for automating the building of Docker containers
through an easy use to API.

## Definitions

gangway using the concept of definitions to build containers. A simple recipe might
look like:

```
source: git@github.com:bobbytables/gangway.git
dockerfile: Dockerfile-build
environment:
  COMMIT_SHA: c7d219401d8e773f1b1190b493887da6f94f2bbe
  COMMIT_AUTHOR: "Robert Ross"
tag: gangway/gangway:latest
```
