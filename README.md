# Paraglider

Paraglider is a golang based tool to convert any linux host into a TCP loadbaancer using iptable rules. Paraglider can be used
in situations where you do not want to use a heavy weight, full blown load balancer, but a minimalistic, lightweight 
TCP load balancer to load balance between given backends. One such usecase for Paraglider can be low end IOT projects where
the device which needs to serve as a load balancer is low end, less powerfull.

## Quick start

The following sections will show how to get quickly started with Paraglider.

### Installing paraglider from binaries

- Download prebuild Paraglider binary from here.
- Unzip the tar file using ```tar xvf Paraglider-0.1.1```
- Once unmpressed, you will find the paraglider binary file.
- You can run paraglider using ```./paraglider -config </path/to/config> <start|stop>
