# Python Voice Chat
This is an example Python application, that implements client-server voice chat capabilities using TCP and UDP sockets.

## Setup
### Windows
#### Dependancy Installation:
- ``pip install -r requirements.txt``

If you have problems with windows dependancies handling, you can use [miniconda](https://docs.conda.io/en/latest/miniconda.html) open-source package management system and environment management system for ``pyaudio`` installation.   

### Linux/Mac
#### Dependancy Installation:
- ``sudo apt install -y portaudio19-dev``
- ``sudo apt install -y pyaudio``
- ``pip install -r requirements.txt``

## Running 
To run a TCP version of application, run ``python server-tcp.py`` and ``python client-tcp.py``.
To run a UDP version of application, run ``python server-udp.py`` and ``python client-udp.py``.

## Usage
- Run ``server-tcp.py`` or ``server-udp.py`` specifying the port you want to bind to.
- If you intend to use this program across the internet, ensure you have port forwarding that is forwarding the port the server is running on to the server's machine local IP (the IP displayed on the server program) and the correct port.
- Clients can connect across the internet by entering your public IP (as long as you have port forwarding to your machine) and the port the machine is running on or in the same network by entering the IP displayed on the server.
- If the client displays ``"Connected to Server"``, you can now communicate with others in the same server by speaking into a connected microphone.

## Requirements
- Python 3
- PyAudio
- Socket Module (standard library)
- Threading Module (standard library)

## Contributing
Since this is a simple project, this repository is unlikely to be majorily changed, however if you wish to contribute with bug fixes/new features/code improvements, pull requests are welcome. Issues are also welcome if you want to discuss or raise an issue.

# Reference

Based on [TomPrograms](https://github.com/TomPrograms) [Python-Voice-Chat](https://github.com/TomPrograms/Python-Voice-Chat) project.

## License
[MIT](https://choosealicense.com/licenses/mit/)
