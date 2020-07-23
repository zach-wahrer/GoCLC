module github.com/zachtheclimber/GoCLC

go 1.13

require client v1.0.0

replace client => ./client

require server v1.0.0

replace server => ./server

require goclctest v1.0.0

replace goclctest => ./goclctest
