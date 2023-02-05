package handler

import (
	"log"

	"github.com/bondyra/wtf/internal/config"
	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type Handler interface {
	Execute(c config.Config)
}

type ConfigHandler struct {
	Args []string
}

type QueryHandler struct {
	Query string
}

func Dupa(x string) string {
	return x + "dupa"
}

func (q QueryHandler) Execute(c config.Config) {
	// PARSE
	// convert .Query into an intermediary tree using command, service config and config
	// PULL
	// data is pulled in parallel based on tree from remote
	// OUTPUT
	// tree is traversed once again with fresh data and json is generated
	// VISUALIZE (optional, true by default)
	// output json is converted to graphviz
}

func visualize() {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	n, err := graph.CreateNode("n")
	if err != nil {
		log.Fatal(err)
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		log.Fatal(err)
	}
	m.SetShape(cgraph.RectangleShape)
	m.SetImage("assets/aws-ec2-icon-423x512-iaajemnx.png")
	m.SetImagePos("tc")
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		log.Fatal(err)
	}
	e.SetLabel("e")
	if err := g.RenderFilename(graph, graphviz.SVG, "graph.svg"); err != nil {
		log.Fatal(err)
	}
}
