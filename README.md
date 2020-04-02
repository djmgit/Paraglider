# Paraglider

Paraglider is a golang based tool to convert any linux host into a TCP loadbaancer using iptable rules. Paraglider can be used
in situations where you do not want to use a heavy weight, full blown load balancer, but a minimalistic, lightweight 
TCP load balancer to load balance between given backends. One such usecase for Paraglider can be low end IOT projects where
the device which needs to serve as a load balancer is low end, less powerfull.

## Quick start

The following sections will show how to get quickly started with Paraglider.

### Using prebuilt Paraglider binary

- Download prebuild Paraglider binary from here.
- Unzip the tar file using ```tar xvf Paraglider-0.1.1```
- Once unmpressed, you will find the paraglider binary file.
- You can run paraglider using ```./paraglider -config </path/to/config> <start|stop>```

### Building binary from source

- Clone/Download this repo into your go path.
- Navigate into the repo and ```cd``` into main
- Run ```go get .```
- After running the above go command run ```go build -o paraglider .```
- If the above commands run successfully, you will find the paraglider binary file generated.
- You can run paraglider using ```./paraglider -config </path/to/config> <start|stop>```
