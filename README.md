# grpcli
A generic grpc client based on grpcurl.


## Requires:

```
import "google.golang.org/grpc/reflection"

func main(){
    server := grpc.NewServer()
    
    reflection.Register(server)
}
```

## Usage:

```
import "github.com/gogoods/grpcli"

func main(){
    cli := grpcli.NewClient()
    
    reply, cost, err := cli.Invoke(context.Background(),
    		"10.40.212.34:40004",
    		"Account",
    		"DescUserInfo",
    		`{"UserID":1532}`,
    	)
}

```