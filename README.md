# GoLang HTTP Server for Remote WOL Requesting from an CSV Computer List

A HTTP server who sends a Wake On LAN package on an HTTP request.

### Docker Image

[![](https://images.microbadger.com/badges/version/dabondi/go-rest-wol.svg)](https://hub.docker.com/repository/docker/dabondi/go-rest-wol "https://hub.docker.com/repository/docker/dabondi/go-rest-wol") [![](https://images.microbadger.com/badges/image/dabondi/go-rest-wol.svg)](https://hub.docker.com/repository/docker/dabondi/go-rest-wol "https://hub.docker.com/repository/docker/dabondi/go-rest-wol")

### Simple bootstrap UI for easy usage

![Screenshot](https://github.com/daBONDi/go-rest-wol/raw/master/screenshot.PNG)

### Simple REST API to let a machine wake someone up

/api/wakeup/computer/**&lt;hostname&gt;** -  Returns a JSON object

```json
{
  "success":true,
  "message":"Succesfully Wakeup Computer Computer1 with Mac 64-07-2D-BB-BB-BF on Broadcast IP 192.168.10.254:9",
  "error":null
}
```

## Enviroment Variables

| Variable Name | Description |
| ------------- | ------------------------------------------------------------------------------- |
| WOLHTTPPORT   | Define the port on which the webserver will listen to **(Default: 8080)**       |
| WOLFILE       | Path to the CSV file containing the list of hosts **(Default: .\computer.csv)** |


## Commandline arguments

| Commandline argument | Example          | Description                                                                            |
| -------------------- | ---------------- | -------------------------------------------------------------------------------------- |
| --port               | --port 80        | Define the port on which the webserver will listen to **(Default: 8080)**              |
| --file               | --file comp.csv  | Path to the CSV file containing the list of hosts **(Default: .\computer.csv)**        |

## Computer list file CSV layout

### Columns
__&lt;name of the computer&gt;__,__&lt;mac address of the computer&gt;__,__&lt;broadcast ip to send the magic packet&gt;__


### Example
```csv
name,mac,ip
Computer1,64-07-2D-BB-BB-BF,192.168.10.254:9
Computer2,2D-F2-3D-06-17-00,192.168.10.254:9
Computer3,FF-B3-95-62-1C-DD,192.168.10.254:9
```

## Docker

**Docker Image:** dabondi/go-rest-wol

```
docker build -t go-rest-wol .
docker run go-rest-wol
```
If you want to run it on a different port (i.e.: 6969) and also want to provide the CSV file on your host:

```
docker run -p 6969:8080 -v $(pwd)/externall-file-on-host.csv:/app/computer.csv dabondi/go-rest-wol
```

If you want to run the WOL Webserver Process in the Webserver on a different Port:

```
# Used if you run in Network Host Mode
docker run -e "WOLHTTPPORT=9090" -p 9090:9090 -v $(pwd)/externall-file-on-host.csv:/app/computer.csv dabondi/go-rest-wol
```

This was a good exercise to learn Golang (and refresh my Docker skills).

Thx https://github.com/sabhiram/go-wol for the WOL code, sorry that I stole it from you, because I got no clue how I can inject it into my program. :-(

**If you have any good ideas, I'm open for pull requests.**
