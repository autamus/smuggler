# Smuggler

![img/smuggler.png](img/smuggler.png)

Smuggler is a web server that helps to get build artifacts out of containers. Specifically,
if you run Smuggler as a service during a build in a Continuous Integration workflow, you
can upload files back to the host (the CI worker) from your build. Smuggler is intended
to be run as a GitHub Action, but can be extended to other use cases.
