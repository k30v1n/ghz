package main

import (
    "fmt"
    "os"

    "github.com/bojand/ghz/printer"
    "github.com/bojand/ghz/runner"
)

func main() {
    report, err := runner.Run(
        "greet.Greeter.SayHello",
        "localhost:5022",
        runner.WithProtoFile("../grpc-service/Protos/greet.proto", []string{}),
        runner.WithDataFromFile("data.json"),
        runner.WithInsecure(true),
    )

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	file, err := os.OpenFile("./unary_SayHello.html", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    printer := printer.ReportPrinter{
        Out:    file,
        Report: report,
    }

    printer.Print("html")
}