# tRPC-Go HelloWorld Protocol Implementation & Deployment

## Project Overview
This project demonstrates how to build and publish a basic `tRPC-Go` echo service using a HelloWorld protocol. It includes steps for code generation with `protobuf`, service creation on the `123` platform, and pipeline setup using `vepc` and `rick`.

## Getting Started

### 1. Create a Git Repository
Create a new git repo using `vepc`:
```bash
vepc new video_app_short_video/hello_alice --desc="Simple Echo Service"
```
### 2. Clone Locally：
```
git clone git@git.code.oa.com:video_app_short_video/hello_alice.git
```
### 3. Initialize Go Module：
```
1. cd hello_alice  
2. go mod init git.code.oa.com/video_app_short_video/hello_alice
```

## Protocol & Stub Setup

### Define Protobuf Interface

- Go to the rick platform and create a .proto file.

- App Name: video_app_short_video

- Service Name: hello_alice

- Service name in option should match the actual server name to avoid default postfixes like _greeter.

### Generate Service Code
- On the 123 platform, generate tRPC-Go stub/service code using:

    - TRPC-Go Stub Mod (for basic stubs)

    - TRPC-Go Service Generator (for full code template)

### Push Stub Code
After downloading the generated stub:

- Replace existing stub folder

- Clean up go.mod:
    ```
    replace git.code.oa.com/... => ./stub/git.code.oa.com/...
    git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter v0.0.0-0001010100000-0000000
    ```
- Run: 
    ```
    go get -v git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter@v1.1.2
    ```
## CI/CD with Vepc

### Create Pipeline

```
vepc create --rick-id=20547 --git-path=video_app_short_video/hello_alice
```
Avoid enabling unit test by default if not needed.

## Service Deployment
- Deploy your service to the test environment using the 123 platform.

- Application Name: video_app_short_video

- Service Name: hello_alice

## Development Workflow
- Modify and push your code locally.

- Re-generate protobuf interfaces if changes were made.

- Trigger pipeline build and publish.

- Use Polaris name service (IP and port) for endpoint testing.

## FAQ

- **Q:** What if the service can't find the stub?
  - **A:** Ensure the service and proto names match; double-check `rick` and `123` platform configurations.
- **Q:** How to clean stub conflicts?
  - **A:** Regenerate stubs, delete local `stub/` and reset `go.mod`.

## Maintainers
Alice He ((hehy1113@gmail.com))