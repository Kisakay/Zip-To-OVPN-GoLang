# Zip-To-OVPN-GoLang
 Zip-To-OVPN rewrite in GoLang (For started Learning GoLang)

This repository was created to provide a convenient and efficient tool for transforming .zip files into .ovpn files, simplifying the process of installing and connecting to a Virtual Private Network (VPN) using OpenVPN. By combining the necessary configuration files into a single .ovpn file, this program makes the installation and connection to the VPN more accessible and user-friendly, especially for those who are not familiar with manual .ovpn file configuration.

## Tranform .zip to .ovpn

### Before
```sh
➜  ovpn ls
ca.crt  client.crt  client.key  openvpn.ovpn
```

### After
```
➜  ovpn ls
ca.crt  client.crt  client.key  openvpn.ovpn combined.ovpn
```
<br>

# FAQ

### What is .ovpn file ?
* An . OVPN file is an [OpenVPN](https://openvpn.com) pro file. An OVPN file's content includes virtual private network access configurations and client program settings for the OpenVPN software
### Why this repo are created ?
* In summary, the main purpose of this repository is to offer a practical and effective solution that easily converts .zip files to .ovpn files, thereby streamlining the setup and usage of a Virtual Private Network with OpenVPN.


**Developed in GoLang by Anaïs Saraiva aka Kisakay with <3**
