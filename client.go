package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"
    "github.com/mrphil2105/ntpSimulation/proto"
)

func calcTimeDiff(c ntpSimulation.NtpClient) {
    var t [4]time.Time

    for i := 0; i < 2; i++ {
        t1 := timestamppb.Now()
        t2, err := c.GetTime(context.Background(), &ntpSimulation.SendTime{Time: t1})

        if err != nil {
            log.Fatalf("could not SendTime: %v", err)
        }

        t[i*2] = t1.AsTime()
        t[i*2+1] = t2.Time.AsTime()
    }

    scew := t[3].Sub(t[0]) - t[2].Sub(t[1])
    log.Printf("\nt1: %v\nt2: %v\nt3: %v\nt4: %v\nscew: %v", t[0], t[1], t[2], t[3], scew)
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
