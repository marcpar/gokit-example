# Hex to RGB

A go lang tool for transforming hex to awesome human readable color.

note:
I use gokit for a microservice framework

## Installation

To run the service theres a makefile involve.

Requirements:
   * docker
   * docker-compose
   * build-essentials

installation to run make
```bash

sudo apt-get install build-essential
```    

there are 2 dockerfile since I dont have registry to push the development to improve developer expirience:

   - dockerfile.dev
      here watcher is installed for developer exp
   - dockerfile
      its using scratch image and alpine as a builder for the final product


Use the Force!! sorry its _**make**_


```bash
# to buid images
make build
# to run developer mode
make dev
# to run test
make test
# to cleanup
make clean
# to push to registry (not currently active)
make push


```

## Usage

requirements

```bash
## you can run without a docker 
##    it uses go module 

cd hex_to_rgb
go mod download
go build
go run cmd/main.go

```

## logging development

```docker
docker-compose logs -f hex_to_rgb 
```

## metrics

Metrics are available in grafana but you have to use docker-compose up  -d 
  - [grafana - visualization port :3000](localhost:3000)
  - [cadvisor - for profiling docker images port:8082](localhost:8082)
  - [prometheus - monitoring port :9090](localhost:9090)
  


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)