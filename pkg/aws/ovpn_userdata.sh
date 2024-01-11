#!/bin/bash
sudo yum update -y
sudo yum install -y openvpn easy-rsa

# Set up Easy RSA for managing SSL
mkdir -p ~/easy-rsa
ln -s /usr/share/easy-rsa/3/* ~/easy-rsa/
cd ~/easy-rsa
./easyrsa init-pki
./easyrsa build-ca nopass
./easyrsa gen-req server nopass
./easyrsa sign-req server server
./easyrsa gen-dh

# Copy necessary keys and certificates
mkdir -p /etc/openvpn/server
cp pki/ca.crt pki/issued/server.crt pki/private/server.key pki/dh.pem /etc/openvpn/server

# Generate server.conf file for OpenVPN
cat <<EOF > /etc/openvpn/server/server.conf
port 1194
proto udp
dev tun
ca ca.crt
cert server.crt
key server.key
dh dh.pem
server 10.8.0.0 255.255.255.0
ifconfig-pool-persist ipp.txt
push "redirect-gateway def1 bypass-dhcp"
push "dhcp-option DNS 8.8.8.8"
push "dhcp-option DNS 8.8.4.4"
keepalive 10 120
cipher AES-256-CBC
user nobody
group nobody
persist-key
persist-tun
status openvpn-status.log
verb 3
EOF

# Enable IP forwarding and NAT
echo 'net.ipv4.ip_forward = 1' | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
sudo service iptables save

# Start OpenVPN
sudo systemctl start openvpn@server
sudo systemctl enable openvpn@server
