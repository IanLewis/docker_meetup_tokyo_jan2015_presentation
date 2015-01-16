#!/bin/sh

set -e

# Install docker
curl -sSL https://get.docker.com/ubuntu/ | sudo sh

# Install mercurial
apt-get install -y mercurial

# Install go
cd /opt
wget --no-verbose https://storage.googleapis.com/golang/go1.4.1.linux-amd64.tar.gz
tar xzf go1.4.1.linux-amd64.tar.gz

cat <<EOF > /etc/environment
PATH="/opt/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games"
GOROOT=/opt/go
GOPATH=/opt/workspace
EOF

# Create the go workspace
mkdir -p /opt/workspace

# Export environment so we can use it.
export PATH="/opt/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games"
export GOROOT=/opt/go
export GOPATH=/opt/workspace

cd $GOPATH
# Get the present tool
# TODO: I'd like to use the git repo but it seems not to work right now.
# go get github.com/golang/tools/cmd/present
echo "Installing cmd/present..."
go get code.google.com/p/go.tools/cmd/present

# Get the docker client library
echo "Installing go-dockerclient..."
go get github.com/fsouza/go-dockerclient

# Pull the nginx docker image.
docker pull nginx

# Run the tool so that the presentation can be viewed.
cat <<EOF > /etc/init.d/present
#!/bin/sh
cd $GOPATH/src/github.com/IanMLewis/docker_meetup_slides
nohup $GOPATH/bin/present -http="0.0.0.0:3999" -orighost="localhost" > /var/log/present.out.log 2> $GOPATH/present.err.log < /dev/null &
EOF

sudo chmod +x /etc/init.d/present
sudo update-rc.d present defaults 

# Execute the present script to start the app.
/etc/init.d/present
