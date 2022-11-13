# Portfolio Risk Management

Portfolio Risk Management API:
- Portfolio & Risk/Return API
- Mobile Portfolio Management Application

## Introduction

This monolithic API is meant to combine the user system, portfolio storage, and risk management system into a single API. This backend will provide data to serve the frontend as well as other web components in the future.

This API is designed to be run in a Docker Container with a few helper tools.

The API follows standard RESTful API designs.
- GET
- POST (create)
- PUT (replace)
- PATCH (update)
- DELETE


## Installation

There are a few ways to use this... the preferred is using a Containerd deployment (see the Docker File). However, you can also run this locally.

### Local

To run locally, you'll need to setup a MySQL instance.

1. Download the desired release from github
2. Create a database, user and password in MySQL for the API to utilize
3. Rename .env.example and edit filling in the values as appropriate
4. Run main.go

### FAQ

1. How do I start the API automatically on boot?

    - Use a systemd service file, or utilize Docker/Containerd/Kubernetes.
        - For systemd, see the `systemd` directory for an example service file or visit [this link](https://www.digitalocean.com/community/tutorials/how-to-use-systemctl-to-manage-systemd-services-and-units) for more information.

2. How do I update the API?

    - Download the latest release, stop the API, replace the binary, and start the API.
    - If using Docker/Containerd/Kubernetes, follow the standard procedure for updating a container. The container is ephemeral, so data will not be lost.

3. How do I update the API without downtime?

    - Use a load balancer that supports zero downtime deployments.  This is not a requirement, but is recommended.
    - Use Kubernetes and utilize the rolling update strategy.
