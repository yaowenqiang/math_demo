package main

import (
    "fmt"
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


    //var float1, float2, float3, float4 float64
    float1 := 1.06
    float2 := 0.06
    float3 := float1 + float2
    float4 := 1.12

    fmt.Printf("%.10f\n", float1)
    fmt.Printf("%.10f\n", float2)
    fmt.Printf("%.10f\n", float1 + float2)
    fmt.Printf("%.10f\n", float3)
    fmt.Printf("%.10f\n", float4)

}
