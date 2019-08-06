##
# ATAGO
# 
# All in one
# Template
# Assembly of
# GO lang
#
# @author chixm
##

ATAGO is an Web Server Template with very often used libraries.
The ATAGO is assembled to quickly make an web service.
 
It contains

html layout
static file server (files like css and javascript files)
database
websocket
keyvalue store(redis)

So, making a copy of The ATAGO will easily brings you to start developing your web service. 

#build for Windows
To build executable binary, execute build.bat (go command is required). 

# cert key was created by command below
openssl req -x509 -nodes -new -keyout cm.key -out cm.crt -days 999