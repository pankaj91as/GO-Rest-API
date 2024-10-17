# Environment Setup
### Install Multipass
```
snap install multipass
```

### Setup Network
```
nmcli connection add type bridge con-name knet ifname knet ipv4.method manual ipv4.addresses 10.10.10.1/24
```

### Launch Instance
```
sudo multipass launch lts --name kmaster --memory 2G --disk 20G --cpus 2 --network name=knet,mode=manual
sudo multipass launch lts --name knode-1 --memory 1G --disk 10G --cpus 1 --network name=knet,mode=manual
sudo multipass launch lts --name knode-2 --memory 1G --disk 10G --cpus 1 --network name=knet,mode=manual
sudo multipass launch lts --name ansible-server --memory 4G --disk 20G --cpus 2 --network name=knet,mode=manual
sudo multipass launch lts --name jenkins-server --memory 4G --disk 20G --cpus 2 --network name=knet,mode=manual
```

## Ansible Installation
```
$ sudo apt update
$ sudo apt install software-properties-common
$ sudo add-apt-repository --yes --update ppa:ansible/ansible
$ sudo apt install ansible
```

## Jenkins Installation
```
sudo wget -O /usr/share/keyrings/jenkins-keyring.asc \
  https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key

echo "deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc]" \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null
  
sudo apt-get update
sudo apt-get install jenkins
```

## Docker Installation
#### Add Docker's official GPG key:
```
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
```

#### Add the repository to Apt sources:
```
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
```

#### Install the Docker packages:

```
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```
### Add docker user into jenkins group
```
sudo usermod -a -G docker jenkins
```