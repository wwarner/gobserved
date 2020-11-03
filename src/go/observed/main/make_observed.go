package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Type    string
	Package string
	Output  string
}

func main() {
	d := &data{}
	flag.StringVar(&d.Type, "type", "", "The type we're observing")
	flag.StringVar(&d.Package, "package", "", "The package in which the generated code will be compiled")
	flag.StringVar(&d.Output, "output", "", "The file in which the generated source will be written")
	flag.Parse()

	t := template.Must(template.New("observer").Parse(observerTemplate))
	f, err := os.Create(d.Output)
	if err != nil {
		panic(err)
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	t.Execute(f, d)
}

var observerTemplate = `package {{.Package}}

import (
	"sync"
)

type Observed{{.Type}} struct {
	mu          sync.Mutex
	Subscribers map[chan <- *{{.Type}}]struct{}
}

func  NewObserved{{.Type}}() (rval *Observed{{.Type}}) {
	return &Observed{{.Type}}{
		mu: sync.Mutex{},
		Subscribers: map[chan <- *{{.Type}}]struct{}{},
	}
}

func (o *Observed{{.Type}}) Subscribe(c chan <- *{{.Type}}) {
	defer o.mu.Unlock()
	o.mu.Lock()
	o.Subscribers[c] = struct{}{}
}

func (o *Observed{{.Type}}) Unsubscribe(c chan <- *{{.Type}}) {
	defer o.mu.Unlock()
	o.mu.Lock()
	delete(o.Subscribers, c)
}

func (o *Observed{{.Type}}) Notify(v *{{.Type}}) {
	for c, _ :=range o.Subscribers {
		c <- v
	}
	return
}
`
