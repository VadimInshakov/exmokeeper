Tool for displaying the real value of investments on the EXMO platform after all trades 

Sample:

    export EXMO_PUBLIC="your public key"
    export EXMO_SECRET="your secret key"
    
    go build
    
    ./exmokeeper -buysymbol BTC -sellsymbol RUB -offset 0 -limit 10000
 
 Output:
    
     For BTC_RUB pair you:
     bought 1.19492792 BTC for 651332.1764 RUB
     sold 1.12556334 BTC for 620284.8899 RUB
 
     YOU HAVE NOW: 0.06936458 BTC
     YOU SPENT TOTAL: 31047.28645 RUB`
