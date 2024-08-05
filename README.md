# RimzBuster C2 Framework

## Overview
The RimzBuster C2 Framework is a sophisticated Command and Control (C2) system designed to demonstrate and analyze the potential cybersecurity vulnerabilities in modern vehicle systems. This framework leverages the On-Board Diagnostics (OBD-II) interface and the Controller Area Network (CAN) bus to remotely control various vehicle functions. Developed in Golang, it comprises a server component for macOS and a client component for Linux, ensuring cross-platform compatibility and robustness.

![Animated demonstration of the project](https://i.ibb.co/fr7rq55/rimz.gif)

![image](https://github.com/user-attachments/assets/4bdbf274-3e88-48e5-837b-2c6ac67d85dd)

## Key Features

### Session Management

**sessions -l**: List all active sessions, providing details like session IDs, usernames, operating systems, time zones, and local times.

**sessions -i <id>**: To connect to a specific client session.

### Command Execution and Control

**accelerate**: Increases vehicle speed to the maximum limit.

**leftsignal and rightsignal**: Activate the respective vehicle indicators.

**hazard**: Activates the hazard lights.

**dooropen and doorclose**: Controls the locking mechanisms of the vehicle doors.

### Data Transmission and File Management

**upload <localpath> <remotedir>**: Transfers files to the vehicle's system.
**download <remotepath> <localpath>**: Retrieves files from the vehicle's system.

### Monitoring and Diagnostics

**sniff**: Captures CAN packet data for monitoring vehicle communication.
**canreplay**: Replays CAN packet data to the vehicle.
**cansend**: Sends custom CAN packet data to the vehicle.

### System and Network Security

**exit**: Terminates the session, ensuring secure disconnection.

## Setup Guide

### Server and Client Configuration

Clone the Repository:

```bash
git clone https://github.com/vengeance1337/Rimzbuster-C2.git
```

### Server Setup:

Modify the IP address and port in the server.go file.
Compile the server:
```bash
go build -o server
```

### Client Setup:

Set the server's IP address and port in the client.go file.
Compile the client:

```bash
go build -o client
```

### Testing and Verification:

Start the server:
```bash
./server
```

### Start the client:
```bash
./client
```

Verify the connection via server logs.

## Installing ICSim on Kali Linux

### Prerequisites:

To effectively test the RimzBuster C2 Framework, you need to install the ICSim tool on a Kali Linux environment. ICSim simulates a vehicle's CAN bus, providing a platform to test various commands and monitor CAN traffic. Follow the steps below to set up ICSim and its dependencies.

### Steps to Install ICSim

Clone the ICSim Repository

Open a terminal on your Kali Linux system and clone the ICSim repository from GitHub:

```bash
git clone https://github.com/zombieCraig/ICSim.git
```

Navigate to the ICSim directory:

```bash
cd ICSim
```

### Install Necessary Dependencies

ICSim requires several libraries, including libsdl2-dev and can-utils. Install these dependencies using the following command:

```bash
sudo apt-get install libsdl2-dev libsdl2-image-dev can-utils
```

Build ICSim

Compile the ICSim source code by running:

```bash
make
```

This command will generate the icsim and controls binaries within the ICSim directory.

Configure Virtual CAN Interfaces

### Load the vcan Kernel Module

Load the vcan module, which supports virtual CAN interfaces:

```bash
sudo modprobe vcan
```

### Create Virtual CAN Interfaces

Create two virtual CAN interfaces, vcan0 and vcan1:

```bash
sudo ip link add dev vcan0 type vcan
sudo ip link add dev vcan1 type vcan
```

### Bring Up the Interfaces

Activate the virtual CAN interfaces:

```bash
sudo ip link set up vcan0
sudo ip link set up vcan1
```

### Verify Installation

To ensure everything is set up correctly, verify the availability of can-utils and test the virtual CAN interfaces.

Check for utilities like candump, cansend, and canplayer by typing their names in the terminal.


### Test the virtual CAN setup using:

```bash
candump vcan0
```
