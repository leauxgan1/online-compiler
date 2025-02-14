# Online Compiler
A website which validates, compiles, and runs user provided source code in their desired language and runs it in a containerized environment before presenting the result.

# Supported languages
- Golang
- C
- Ziglang (Stretch goal)

# Supported features
 - Compiling with various compiler optimizations
 - Console window with compiler warnings and errors

# Architecture

(1) <-> (2) <-> (3)

1) Frontend client
	- Contains a text area for programs to be written
    - Upon clicking the 'Compile' button, the code is sent to the backend server to be propagated to a container and compiled, returning any compiler errors but not the result of the program
    - Upon clicking the 'Run' button, the code is sent to the backend server to be propagated to a container, compiled, and then ran, returning any compiler errors and the result of the program being run.
    - The client performs long polling to await the response from the server
2) Backend server
    - Maintains the endpoints of '/compile' and '/run'
        - The '/compile' endpoint anticipates a POST request from the client with the request body containing a 'code' field, sending the code to a container to be compiled using the respective language's compile command
        - The '/run' endpoint performs the above only with the compile command followed by a run command, capturing the result of the program and sending it back to the long-polling user's client.
3) Kubernetes cluster 
    - A number of containers are set up with the available compilers, awaiting requests from the server to be provided code to be either compiled or compiled and ran.
    - Each container does not maintain state and merely returns the result of compilation or execution of a provided program
    - 

# Deployment
1. Start up backend and client servers
2. Start up message queue and load balancers for backend 
3. Start up containers for running compiled code

