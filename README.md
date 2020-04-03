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
- Navigate into the repo and ```cd``` into ```main```.
- Run ```go get .```
- After running the above go command run ```go build -o paraglider .```
- If the above commands run successfully, you will find the paraglider binary file generated.
- You can run paraglider using ```./paraglider -config </path/to/config> <start|stop>```

## Configuring paraglider

By default Paraglider expects to find the config file at ```/etc/paraglider/glider.yaml```
However it is possible to provide a sample configuration file uaing ```-config``` option

Given below is a sample configuration file for Paraglider:

```
frontend:
  bind: 166.1.101.2:4444
  privateip: 172.19.0.4
  backends:
    - 172.19.0.2:5555
    - 172.19.0.3:5555
    - 172.19.0.6:5555
```

In the above sample yaml, ```frontend``` is the root object which contains the following properties:
- ```bind```: The adress to which the load balancer will bind it self. The coresponding interface should be attached to the
  network to which the users/clients are connected to, ie, the users should be able to reach this IP.
  Optionally this can be same as the private IP.
  
- ```privateip```: This is he ip which should be reachable by the target backends. In other words this IP should 
  correspond to the interface which is connected to the network to which the target backends are also connected. 
  Optionally as aleady mentioned, this IP can be same as the ```bind``` IP.
  
- ```backends```: Listof target backends. Backends should be in the format ```<IP>:<PORY>

### Starting and Stoppomg Paraglider

Paraglider can be started using ```./paraglider [-config] [config file] start ```
It can be stopped using ```./paraglider [-config] [config_file] stop```

## How Paraglider works 

