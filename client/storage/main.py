#!/usr/bin/env python
   
import socket 

TCP_IP = '104.155.226.23'
TCP_PORT = 9977
BUFFER_SIZE = 1024
MESSAGE = "Hello, penlook!"
  
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))
s.send(MESSAGE)
data = s.recv(BUFFER_SIZE)
s.close()
  
print "received data:", data
