#!/bin/bash

curl http://localhost:8080/loginJson -d "`cat login.json`" -H "Content-Type: application/json"
echo
curl http://localhost:8080/loginJson -d "{\"User\": \"test\"}" -H "Content-Type: application/json"
echo
curl http://localhost:8080/loginXML -d "`cat login.xml`" -H "Content-Type: application/xml"
echo
curl http://localhost:8080/loginForm -d "user=test&password=123"
echo

