timezone
===

helper for timezone

### QuickStart

````
import "github.com/liujianping/timezone"


//! add your business timezone
timezone.AddZone("BZ-1", 1)
timezone.AddZone("BZ-2", 2)

//! del your business timezone
timezone.DelZone("BZ-1")
timezone.DelZone("BZ-2")

//! get my timezone location
timezone.GetZone("BZ-1")

//! set time in dest zone

now := time.Now()

t1, _ := timezone.Zone("BZ-2", now)

//! or

t2 := now.In(timezone.GetZone("BZ-2"))

//! check t1 == t2
fmt.Println(t1.Equal(t2))

//! time utc
t4 := now.UTC()

//! time convert hour zone
t5 := timezone.ZoneHour(4, now)

//! check now == t5
fmt.Println(now.Equal(t5))
````

of course, you can use the system timezone either.

````
import "github.com/liujianping/timezone"

now := time.Now()

t1, _ := timezone.Zone("Asia/Beijing", now)

t2 := now.In(timezone.GetZone("Asia/Beijing"))

````
