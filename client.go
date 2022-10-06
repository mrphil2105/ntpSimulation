package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"
    "github.com/mrphil2105/ntpSimulation/proto"
)

func getSingleTimeStamp(c ntpSimulation.NtpClient) (time.Time, time.Time){
    t1 := timestamppb.Now()
    t2, err := c.GetTime(context.Background(), &ntpSimulation.SendTime{Time: t1})
    if err != nil {
        log.Fatalf("could not SendTime: %v", err)
    }
    return t1.AsTime(), t2.Time.AsTime()
}

func calcTimeDiff(c ntpSimulation.NtpClient) {
    t1, t2 := getSingleTimeStamp(c)
    t3, t4 := getSingleTimeStamp(c)
    scew := t4.Sub(t1) - t3.Sub(t2)

    log.Printf("\nt1: %v\nt2: %v\nt3: %v\nt4: %v\nscew: %v", t1, t2, t3, t4, scew)
}

func main() {
    conn, err := grpc.Dial(":50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    c := ntpSimulation.NewNtpClient(conn)

    calcTimeDiff(c)
    conn.Close()
}
