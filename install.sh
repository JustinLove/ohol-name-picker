#!/bin/sh

apt-get update
apt-get install -y zip vim git curl

echo "Downloading Go"
curl --silent https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz > /tmp/go.tar.gz
echo "Extracting Go"
tar -xvzf /tmp/go.tar.gz --directory /usr/lib
echo "Setting Go environment variables"
if [ ! -e /home/vagrant/go ]; then
  mkdir /home/vagrant/go
fi
chmod -R 777 /home/vagrant/go
echo 'export GOROOT="/usr/lib/go"' >> /home/vagrant/.profile
echo 'export GOPATH="/home/vagrant/go"' >> /home/vagrant/.profile
echo 'export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"' >> /home/vagrant/.profile
