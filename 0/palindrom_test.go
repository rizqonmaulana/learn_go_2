package main

import "testing"

func TestLoop(t *testing.T){
    loop1 := loop(1,100)
    if  loop1 != 18{
        t.Error("Result value must be 18. your test result =", loop1 )
    }

    loop2 := loop(1,1000)
    if loop2 != 108{
        t.Error("Result value must be 108. your test result =", loop2 )
    }

    loop3:=loop(1,9999)
    if loop3 != 198{
        t.Error("Result value must be 198. your test result =", loop3 )
    }
}

func TestErrorCheck(t *testing.T){
    error1 := errorCheck(1,100)
    if error1 != "" {
        t.Error("Result value must be nil. your test result =", error1)
    }

    error2 := errorCheck(0,100)
    if error2 != "int value cannot lower than 1" {
        t.Error("Result value must be 'int value cannot lower than 1'. your test result = '"+ error2 + "'")
    }

    error3 := errorCheck(1,1000000001)
    if error3 != "int value cannot higher than 10^9" {
        t.Error("Result value must be 'int value cannot higher than 10^9'. your test result = '"+ error3 + "'")
    }

    error4 := errorCheck(100,10)
    if error4 != "int2 value cannot lower than int1 value" {
        t.Error("Result value must be 'int2 value cannot lower than int1 value'. your test result = '"+ error4 + "'")
    }

    // error1 := errorCheck(1,100)
    // if error1 != nil {
    //     t.Error("Result value must be nil. your test result =", error1)
    // }
}