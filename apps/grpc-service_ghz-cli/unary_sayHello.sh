
ghz --insecure --async \
    --proto ../grpc-service/Protos/greet.proto \
    --call greet.Greeter.SayHello \
    --data '{ "name" : "Someone {{.RequestNumber}}" }' \
    --output unary_sayHello.html --format html  \
    localhost:5022