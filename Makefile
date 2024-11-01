TAG:= v1alpha1

build:
        docker build -t whalepower/whalepower:${TAG} .

run:
        docker run -itd --name whalepower -e PORT=8888 -p 8888:8888 whalepower/whalepower:${TAG}

rm:
        docker stop whalepower && docker rm whalepower