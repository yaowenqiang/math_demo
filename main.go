package main

import (
    "fmt"
    "log"
)

/*
const (
    None = 0
    EnSuite = 1
    SeaView = 2
    StairFreeAccess 4
    Breakfast = 8
    //etc
)
*/
const (
    //None = 0x0
    //EnSuite = 0x1
    //SeaView = 0x2
    //StairFreeAccess = 0x4
    //Breakfast = 0x8
    //0x10
    //0x20
    //0x40
    //0x80
    //etc
)

const (
    None = 0
    EnSuite = 1
    SeaView = 1 << 1
    StairFreeAccess = 1 << 2
    Breakfast = 1 << 3
    //etc
)

func main() {
    f := 1.2
    b := 15000000000000
    c := 15e6
    d := -1.5e-6
    g := 0.125
    fmt.Println(f)
    fmt.Println(b)
    fmt.Printf("%f\n", f)
    fmt.Printf("%.20f\n", f)
    fmt.Printf("%.20f\n", g)
    fmt.Printf("%f\n", c)
    fmt.Printf("%f\n", d)

    waterMolecule := 0.00000000000000000000000002992
    fmt.Printf("%f\n", waterMolecule)
    fmt.Printf("%.26f\n", waterMolecule)
    fmt.Printf("%.29f\n", waterMolecule)
    fmt.Printf("%+v\n", waterMolecule)
    /*
    fmt.Printf("%+v\n", 0.0/0.0)
    fmt.Printf("%+v\n", 1.0/0.0)
    fmt.Printf("%+v\n", -1.0/0.0)
    */
    var zero float64
    zero = 0.0

    fmt.Printf("%+v\n", -1.0/zero)
    fmt.Printf("%+v\n", 1.0/zero)
    fmt.Printf("%+v\n", 0/zero)

    var infinity float64
    infinity = -1/zero

    fmt.Printf("%+v\n", infinity)
    fmt.Printf("%+v\n", 1/infinity)
    fmt.Printf("%+v\n", infinity/infinity)
    fmt.Printf("%+v\n", infinity*infinity)


    var float1, float2, float3, float4 float64
    float1 = 1.07
    float2 = 0.07
    float3 = float1 + float2
    float4 = 1.14

    fmt.Printf("%.20f\n", float1)
    fmt.Printf("%.20f\n", float2)
    fmt.Printf("%.20f\n", float1 + float2)
    fmt.Printf("%.20f\n", float3)
    fmt.Printf("%.20f\n", float4)
    fmt.Printf("%+v\n", float3 == float4)


    float5 := 0.25
    float6 := 0.2
    float7 := 0.05
    float8 := float6 + float7

    fmt.Printf("%.20f\n", float5)
    fmt.Printf("%.20f\n", float8)

    //true
    fmt.Printf("%+v\n", float8 == float5)

    float9 := 0.05
    float10 := float9

    //true
    fmt.Printf("%+v\n", float9 == float10)

    float11 := 0.05
    float12 := float64(float32(float11))

    fmt.Printf("%.20f\n", float11)
    fmt.Printf("%.20f\n", float12)
    //false
    fmt.Printf("%+v\n", float11 == float12)


    //Never use float point as a map key
    map1 := make(map[float64]int)
    map1[float11] = 1
    map1[float12] = 1

    fmt.Println(map1)
    fmt.Println(log.Flags())
    fmt.Println(log.Ldate)
    fmt.Println(log.Ltime)
    fmt.Println(log.Lmicroseconds)
    fmt.Println(log.Llongfile)
    fmt.Println(log.Lshortfile)
    fmt.Println(log.LUTC)
    fmt.Println(log.Lmsgprefix)
    fmt.Println(log.LstdFlags)
    log.Println("hello world")
    //log.SetFlags(log.Llongfile|log.LUTC|log.Ltime)
    //log.SetFlags(log.Flags() & (~log.Ldate))
    //log.SetFlags(log.Flags() & (^log.Ldate))
    if((log.Flags() & log.Ldate) != 0) {
        log.Println("lDate flags is set")
    } else {
        log.Println("lDate flags isn't set")
    }
    log.SetFlags(log.Flags() & (^log.Ldate))
    if((log.Flags() & log.Ldate) != 0) {
        log.Println("lDate flags is set")
    } else {
        log.Println("lDate flags isn't set")
    }
    log.Println("long file")







}
