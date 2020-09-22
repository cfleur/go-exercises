package main
import(
	"fmt"
)

func main() {
type ByteSize int64
 const (
     _ = iota 
     KB ByteSize = 1<<(10*iota)
     MB
     GB
)

  fmt.Println(KB, MB, GB)      
}