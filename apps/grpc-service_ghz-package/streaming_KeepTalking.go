package main

import (
    "fmt"
    "os"

    "github.com/bojand/ghz/printer"
    "github.com/bojand/ghz/runner"
)

func main() {
    report, err := runner.Run(
        "greet.Greeter.KeepTalking",
        "localhost:5022",
        runner.WithProtoFile("../grpc-service/Protos/greet.proto", []string{}),
        runner.WithDataFromFile("data.json"),
        runner.WithInsecure(true),
        runner.WithTotalRequests(100),
        runner.WithStreamCallCount(2), // 1 initial + streaming
        runner.WithTimeout(0),
        runner.WithCountErrors(true),
        // runner.WithStreamCallDuration(100000000), //Note that in server streaming calls the cancellation will result in call cancelled error.
        // runner.WithRPS(0),
        // runner.WithStreamCallCount(10),
    )

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	file, err := os.OpenFile("./streaming_KeepTalking.html", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

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