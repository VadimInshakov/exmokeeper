Sample:

    export EXMO_PUBLIC="your public key"
    export EXMO_SECRET="your secret key"
    
    go build
    
    ./exmokeeper -buysymbol BTC -sellsymbol RUB -offset 0 -limit 10000