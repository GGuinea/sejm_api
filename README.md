PL
API Polskiego Sejmu

---

EN
Polish parliament API

```
go get github.com/gguinea/sejm_api

import (
	"github.com/gguinea/sejm_api/api"
)

func main() {
	client := api.NewClient("10")
	client.ListClubs()
}

```

https://api.sejm.gov.pl/pages.html
