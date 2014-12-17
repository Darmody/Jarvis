package main

import (
    
    "github.com/go-martini/martini"
    //"github.com/yosssi/ace"
    "github.com/yosssi/martini-acerender"
)

func main() {

    m := martini.Classic()
    m.Use(acerender.Renderer(nil))    

    m.Get("/", func(r acerender.Render) {

      r.HTML(200, "base:inner", map[string]string{"Msg": "喻小伟送我域名: jarvisair.io"}, nil)
    })

    m.Run()
}
