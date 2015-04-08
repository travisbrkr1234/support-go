#!/usr/bin/env bash

echo "Installing mysql server"
apt-get install mysql-server --assume-yes > /dev/null

# Setup mysql user and database for vagrant user

echo "Install Go 1.4.2"
echo "Installing dependency: git"
apt-get install git --assume-yes > /dev/null

echo "Downloading and unpacking go1.4.2.linux-amd64"
curl -s https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar xz -C /usr/local

echo "Setting GOPATH"
cat <<PROFILE >> /home/vagrant/.profile
# Setup Go
export GOPATH=\$HOME/go
export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin
PROFILE

# Setup go workspace
echo "Creating go workspace"
su - vagrant -c 'mkdir -p $GOPATH/src $GOPATH/bin $GOPATH/pkg'

echo "Linking vagrant directory to workbench"
# Init project directory
su - vagrant -c 'mkdir -p $GOPATH/src/github.com/travisbrkr1234 && ln -s /vagrant $GOPATH/src/github.com/travisbrkr1234/support-go'

cat <<PROFILE >> ~vagrant/.profile
# Change directory to project
cd \$GOPATH/src/github.com/travisbrkr1234/support-go
PROFILE
