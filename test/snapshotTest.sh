#!/bin/sh

curl -F "image=@test.jpg;Filename=2019-10-1.jpg;type=image/jpeg" http://127.0.0.1:8080/snapshot
