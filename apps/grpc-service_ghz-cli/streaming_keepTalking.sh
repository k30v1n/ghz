
ghz --insecure --async \
    --proto ../grpc-service/Protos/greet.proto \
    --call greet.Greeter.KeepTalking \
    --data '{ "name" : "Someone {{.RequestNumber}}" }' \
    --total 1000 --concurrency 20 \
    --timeout 0 --stream-call-count 2 --count-errors \
    --output streaming_keepTalking.html --format html  \
    localhost:5022