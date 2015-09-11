Backend JayDiApps mobile apps
=============================

# To run on container

docker build -t jdbackend .
docker run --publish 8001:8080 jdbackend

# To run on EC2

sudo apt-get update
sudo apt-get install build-essential golang git
mkdir $HOME/go
export GOPATH=$HOME/go
go get github.com/juandiegoh/jaydiapps
cd ~/go/src/github.com/juandiegoh/jaydiapps
go build

sudo vim /etc/init/jaydiapps.conf

`description "start and stop the go program 'jaydiapps'"

start on filesystem or runlevel [2345]
stop on runlevel [!2345]

env USER='ubuntu'
env APP_DIR='/home/ubuntu/go/src/github.com/juandiegoh/jaydiapps/'
env APP_EXEC='jaydiapps'

exec start-stop-daemon --start --chuid ${USER} --chdir ${APP_DIR} --exec ${APP_DIR}${APP_EXEC} >> /home/ubuntu/go/logs/app.log 2>&1`

# create logs folder
```sh
cd /home/ubuntu/; mkdir logs
sudo start jaydiapps
sudo service jaydiapps stop
```

# Rotate logs
Create file /etc/logrotate.d/golang

```
/home/ubuntu/go/logs/*.log {
        daily
        missingok
        rotate 7
        compress
        notifempty
        nocreate
}
```

## Redirect :80 to :8080
### install nginx
sudo apt-get install nginx

Proxy :80 -> :8080

modify /etc/nginx/nginx.conf
`user www-data;
worker_processes 4;
pid /run/nginx.pid;

events {
	worker_connections 768;
	# multi_accept on;
}

http {
	##
	# Logging Settings
	##

	access_log /var/log/nginx/access.log;
	error_log /var/log/nginx/error.log debug;

	server {
                listen 80;
                # The host name to respond to
                server_name jaydiapps;

                location / {
                        proxy_pass http://localhost:8080/;
                        proxy_set_header X-Real-IP  $remote_addr;
                        proxy_set_header X-Forwarded-For $remote_addr;
                        proxy_set_header Host $host;
                        proxy_set_header X-Real-Port $server_port;
                        proxy_set_header X-Real-Scheme $scheme;
                }
        }
}`

to start nginx
```sh
sudo nginx
```

to reaload when modified conf
```sh
sudo nginx -s reaload
```

to stop nginx

```sh
sudo nginx -s stop
```
