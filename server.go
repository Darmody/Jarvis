package main

import (
    
    "github.com/go-martini/martini"
    //"github.com/yosssi/ace"
    //"github.com/yosssi/martini-acerender"
    "github.com/martini-contrib/render"
)

func main() {

    m := martini.Classic()
    
    //加载render中间件
    m.Use(render.Renderer(render.Options{
      Extensions: []string{".tmpl", ".html"},
    }))

    m.Get("/", func(r render.Render) {

      r.HTML(200, "index", nil)
    })

    m.Run()
}
