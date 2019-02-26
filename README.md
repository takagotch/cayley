### cayley
---
https://github.com/cayleygraph/cayley

https://github.com/cayleygraph/cayley/tree/master/docs

```go
g.V().Has("<name>","Casablanca")
  .Out("</film/film/starring>").Out("</film/performance/actor>")
  .Out("<name>").All()
  
var filmToActor = g.Morphism().Out("</film/film/starring>").Out("</film/performance/actor>")

g.V().Has("<name>", "Casablanca").Follow(filmToActor).Out("<name>").All()


graph.Vertex().GetLimit(5)
graph.Vertex("Humphrey Bogart").All()
g.V("Humphrey Bogart").All()
g.V("Humphrey Bogart").In("<name>").All()
g.V("Casablanca").In("<name>").All()
g.V().Has("<name>", "Casablanca").All()

g.V("<dani>").Tag("source").Out("<follows>").Tag("target").All()
```

```
[
{
  "source": "node1",
  "target": "node2"
},
{
  "source": "node1",
  "target": "node3"
}
]
```

```sh
./cayley http -c cayley_overview.yml
listening on :64210, web interface at http://localhost:64210

:d subject predicate object .
2 + 2
x = 2 * 8
x
graph.Vertex("<dani>").All()
graph.Vertex("<dani>").Out("<follows>")All()

./cayley http -i ./data/30kmoviedata.nq.gz -d memstore --host=:64210
./cayley load -c cayley_overview.yml -i data/testdata.nq
./cayley load -c cayley_overview.yml -i data/testdata.nq --alsologtostderr=true
./cayley conv -i dataset.nq.gz -o dataset.pq.gz
,/cayley repl -c cayley_overview.yml
:a subject predicate object label .
:d subject predicate object .
```


